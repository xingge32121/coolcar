/*
 * @Author: Wujiahuo
 * @Date: 2022-08-26 14:50:15
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-21 11:38:02
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\auth\dao\mongodb.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package dao

import (
	"context"
	"coolcar/shared/id"
	mong "coolcar/shared/mongodb"
	"coolcar/shared/mongodb/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const OpenIdField = "open_id"

// Mongo defines a mongo dao.
type Mongo struct {
	col      *mongo.Collection
	newObjID func() primitive.ObjectID
}

// Mongo create new dao
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("account"),
		newObjID: primitive.NewObjectID,
	}
}

func (m *Mongo) ResolveAccountId(c context.Context, open_id string) (id.AccountId, error) {
	insertedId := m.newObjID()
	res := m.col.FindOneAndUpdate(c, bson.M{
		OpenIdField: open_id,
	}, mong.SetOnInsert(bson.M{
		mong.IDFieldName: insertedId,
		OpenIdField:      open_id,
	}), options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))
	err := res.Err()
	if err != nil {
		return "", fmt.Errorf("connot FindOneAndUpdate %v", err)
	}

	var row mong.IDField
	err = res.Decode(&row)

	if err != nil {
		return "", fmt.Errorf("connot Decode %v", err)
	}

	return objid.ToAccountID(row.ID), nil

}

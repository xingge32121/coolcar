/*
 * @Author: Wujiahuo
 * @Date: 2022-09-15 16:26:19
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-21 11:24:09
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\rental\trip\dao\mongodb.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mong "coolcar/shared/mongodb"
	"coolcar/shared/mongodb/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField      = "trip"
	accountIdField = tripField + ".accountid"
)

// Mongo defines a mongo dao.
type Mongo struct {
	col      *mongo.Collection
	newObjId func() primitive.ObjectID
}

// Mongo create new dao
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("trip"),
		newObjId: primitive.NewObjectID,
	}
}

type TripRecord struct {
	mong.IDField       `bson:"inline"`
	mong.UpdateAtField `bson:"inline"`
	Trip               *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mong.NewObjId
	r.UpdateAt = mong.UpdateAt()

	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (m *Mongo) GetTrip(c context.Context, id id.Id, acccount id.AccountId) (*TripRecord, error) {
	objId, err := objid.FromID(id)
	if err != nil {
		return nil, fmt.Errorf("invaild id:%v", err)
	}
	res := m.col.FindOne(c, bson.M{
		mong.IDFieldName: objId,
		accountIdField:   acccount,
	})
	err = res.Err()
	if err != nil {
		return nil, err
	}
	var trip TripRecord
	ok := res.Decode(&trip)
	if ok != nil {
		return nil, fmt.Errorf("cannot Decode trip")
	}
	return &trip, nil
}

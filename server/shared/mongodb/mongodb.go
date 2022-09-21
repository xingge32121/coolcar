/*
 * @Author: Wujiahuo
 * @Date: 2022-08-29 10:47:08
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-16 16:24:43
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\shared\mongodb\mongodb.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package mong

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IDFieldName       = "_id"
	UpdateAtFieldName = "updateat"
)

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}
type UpdateAtField struct {
	UpdateAt int64 `bson:"updateat"`
}

var NewObjId = primitive.NewObjectID()

var UpdateAt = func() int64 {
	return time.Now().UnixNano()
}

// Set returns a $set update document.
func Set(m interface{}) interface{} {
	return bson.M{
		"$set": m,
	}
}

// SetOnInsert returns a $setOnInsert update document.
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

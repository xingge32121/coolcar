/*
 * @Author: Wujiahuo
 * @Date: 2022-09-16 16:34:28
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-20 16:04:30
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\rental\trip\dao\mongodb_test.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"coolcar/shared/mongodb/objid"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateTrip(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))

	if err != nil {
		t.Fatalf("cannot connet mongodb %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	acc := id.AccountId("1112222c")
	r, err := m.CreateTrip(context.Background(), &rentalpb.Trip{
		AccountId: acc.String(),
		CarId:     "ab1c",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  123,
				Longitude: 287,
			},
			PoiName:  "电信大楼",
			FeeCent:  1200,
			KmDriven: 38,
		},
		Status: rentalpb.TripStatus_FINISHED,
	})

	if err != nil {
		t.Errorf("cannot create Trip %v", err)
	}
	t.Errorf("get date %v", r)

	trip, err := m.GetTrip(context.Background(), id.Id(objid.ToTripID(r.ID)), acc)
	if err != nil {
		t.Errorf("cannot get Trip %v", err)
	}

	t.Errorf("get date %v but %v", r, trip)
}

/*
 * @Author: Wujiahuo
 * @Date: 2022-08-26 15:58:55
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-21 11:39:45
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\auth\dao\mongodb_test.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package dao

import (
	"context"
	"coolcar/shared/id"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountId(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))

	if err != nil {
		t.Fatalf("cannot connet mongodb %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	acc := id.AccountId("123")
	accountId, err := m.ResolveAccountId(c, acc.String())

	if err != nil {
		t.Fatalf("cannot ResolveAccountId accountId %v", err)
	} else {
		var open_id = "63086f79fa230cacfb73102e"

		if accountId.String() != open_id {
			t.Fatalf("revice  %q but get %q", accountId, open_id)
		}
	}

}

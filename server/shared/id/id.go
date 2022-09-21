/*
 * @Author: Wujiahuo
 * @Date: 2022-09-16 16:49:39
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-20 16:48:36
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\shared\id\id.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package id

type Id string

func (i Id) String() string {
	return string(i)
}

type AccountId string

func (a AccountId) String() string {
	return string(a)
}

type TripId string

func (t TripId) String() string {
	return string(t)
}

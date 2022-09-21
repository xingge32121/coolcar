/*
 * @Author: Wujiahuo
 * @Date: 2022-09-05 15:48:37
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-06 11:35:56
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\shared\context\main.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func to1(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("to1 is over")
			return
		default:
			fmt.Println("to1 : ", n)
			n++
			time.Sleep(time.Second)
		}
	}
}
func to2(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("to2 is over")
			return
		default:
			fmt.Println("to2 : ", n)
			n++
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	// 设置为6秒后context结束
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	go to1(ctx)
	// go to2(ctx)
	n := 1
	for {
		select {
		case <-time.Tick(time.Second):
			if n == 9 {
				return
			}
			fmt.Println("number :", n)
			n++
		}
	}

}

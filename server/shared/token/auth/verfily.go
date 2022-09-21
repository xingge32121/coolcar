/*
 * @Author: Wujiahuo
 * @Date: 2022-09-02 10:34:21
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-05 11:44:37
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\shared\token\auth\verfily.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package auth

import (
	"crypto/rsa"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)
/**
 * @description: 
 * @return {*}
 */
type JWTTokenGenVerfily struct {
	PublicKey *rsa.PublicKey
}

/**
 * @description:
 * @param {string} accessToken
 * @return {*}
 */
func (v *JWTTokenGenVerfily) Verfily(accessToken string) (string, error) {
	clm := jwt.StandardClaims{}
	t, err := jwt.ParseWithClaims(accessToken, &clm, func(t *jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot Parse token:%v", err)
	}
	ok := t.Valid
	if !ok {
		log.Fatal("cannot Valid token", "")
		return "", nil
	}
	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("cannot Valid Claims token:%v", err)
	}
	return clm.Subject, nil
}

/*
 * @Author: Wujiahuo
 * @Date: 2022-08-19 11:09:20
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-21 11:41:56
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\auth\auth\auth.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"log"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	authpb.UnimplementedAuthServiceServer
	OppenIdResolve OppenIdResolve
	TokenGenerate  TokenGenerate
	TimeExp        time.Duration
	Mongo          *dao.Mongo
}

type OppenIdResolve interface {
	Resolve(code string) (string, error)
}

type TokenGenerate interface {
	GenerateToken(accountId string, exp time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	OppenId, err := s.OppenIdResolve.Resolve(req.Code)
	if err != nil {
		log.Fatal("canont create OppenId:", err)
	}
	accountId, err := s.Mongo.ResolveAccountId(c, OppenId)
	if err != nil {
		s.Logger.Info("recive code", zap.String("code", OppenId))
		return nil, nil
	}
	token, err := s.TokenGenerate.GenerateToken(accountId.String(), s.TimeExp)
	if err != nil {
		s.Logger.Info("TokenGenerate code", zap.String("code", token))
		return nil, nil
	}
	return &authpb.LoginResponse{
		AccessToken: token,
		ExpiresIn:   int32(s.TimeExp.Seconds()),
	}, nil
}

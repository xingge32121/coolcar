/*
 * @Author: Wujiahuo
 * @Date: 2022-09-06 16:20:51
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-09 16:10:47
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\rental\main.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip"
	"coolcar/shared/server"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap %c", err)
	}
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:       "rental",
		Tcp:        ":8082",
		PublicFIle: "../shared/token/public.key",
		Logger:     logger,
		RegisterFunc: func(s *grpc.Server) {
			rentalpb.RegisterTripServiceServer(s, &trip.Service{
				Logger: logger,
			})
		},
	}))
}
func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}

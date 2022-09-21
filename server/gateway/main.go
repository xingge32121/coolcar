/*
 * @Author: Wujiahuo
 * @Date: 2022-08-19 15:13:48
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-09 15:56:54
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\gateway\main.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/server"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap %c", err)
	}
	logger.Sugar().Infof("grpc geteway started at 0")
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:  true,
				UseEnumNumbers: true,
			},
		},
	))
	logger.Sugar().Infof("grpc geteway started at 1")
	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "rental",
			addr:         "localhost:8082",
			registerFunc: rentalpb.RegisterTripServiceHandlerFromEndpoint,
		},
	}

	for _, v := range serverConfig {
		err := v.registerFunc(
			c,
			mux,
			v.addr,
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
		)
		if err != nil {
			logger.Sugar().Fatalf("cannot start grpc gateway: %v +%s", err, v.name)
		}
	}
	logger.Sugar().Infof("grpc geteway started at 2")
	addr := ":8080"
	logger.Sugar().Infof("grpc geteway started at 3")
	logger.Sugar().Fatal(http.ListenAndServe(addr, mux))
}

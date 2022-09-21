package main

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
	"coolcar/auth/dao"
	"coolcar/auth/token"
	"coolcar/auth/wechat"
	"coolcar/shared/server"
	"io/ioutil"
	"log"
	"os"
	"time"

	"coolcar/auth/auth"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	logger, err := server.NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap %c", err)
	}
	priKey, err := os.Open("auth/private.key")
	if err != nil {
		log.Fatalf("cannot Open private %v", err)
	}
	priByts, err := ioutil.ReadAll(priKey)
	if err != nil {
		log.Fatalf("cannot ReadAll private %v", err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priByts)
	if err != nil {
		log.Fatalf("cannot create private %v", err)
	}
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		log.Fatalf("cannot connect mongodb %v", err)
	}
	logger.Sugar().Fatal(server.RunGRPCServer(&server.GRPCConfig{
		Name:   "auth",
		Tcp:    ":8081",
		Logger: logger,
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{
				OppenIdResolve: &wechat.Service{
					AppId:  "wxbee9abd628f92df4",
					Secret: "0c6e253e5be5a57b3f548c9e7b707576",
				},
				Logger:        logger,
				Mongo:         dao.NewMongo(mc.Database("coolcar")),
				TimeExp:       10 * time.Second,
				TokenGenerate: token.NewJWTTokenGen("coolcar/auth", privKey),
			})
		},
	}))
}

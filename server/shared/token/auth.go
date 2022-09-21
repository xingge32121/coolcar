/*
 * @Author: Wujiahuo
 * @Date: 2022-09-07 14:53:39
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-16 16:51:15
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\shared\token\auth.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package token

import (
	"context"
	"coolcar/shared/id"
	"coolcar/shared/token/auth"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHead = "authorization"
	BearerFiled       = "Bearer "
)

func Interceptor(pubkeyFile string) (grpc.UnaryServerInterceptor, error) {
	pubkeys, err := os.Open(pubkeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannit open pubkeyFile %v", err)
	}

	pubbyts, err := ioutil.ReadAll(pubkeys)

	if err != nil {
		return nil, fmt.Errorf("cannot readall pubKeyFile %v", err)
	}

	pubkey, err := jwt.ParseRSAPublicKeyFromPEM(pubbyts)

	if err != nil {
		return nil, fmt.Errorf("cannot parse  pubKeyFile %v", err)
	}
	i := &interceptor{
		verfily: &auth.JWTTokenGenVerfily{
			PublicKey: pubkey,
		},
	}
	return i.HandleReq, nil
}

type tokenVerifier interface {
	Verfily(string) (string, error)
}
type interceptor struct {
	verfily tokenVerifier
}

func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unimplemented, "")
	}
	aId, err := i.verfily.Verfily(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)

	}
	return handler(ContextWithAccountId(ctx, id.AccountId(aId)), req)
}

func tokenFromContext(ctx context.Context) (string, error) {
	Unimplemented := status.Error(codes.Unimplemented, "")
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", Unimplemented
	}
	token := ""
	for _, v := range m[authorizationHead] {
		if strings.Contains(v, BearerFiled) {
			token = v[len(BearerFiled):]
		}
	}
	if token == "" {
		return "", Unimplemented
	}
	return token, nil
}

type accountIdKey struct {
}

// ContextWithAccountId creates a context with given account
func ContextWithAccountId(ctx context.Context, aid id.AccountId) context.Context {
	return context.WithValue(ctx, &accountIdKey{}, aid)
}

//AccountIdFromContxt gets account id from context
// Returns Unauthenticated error if no account id is ava
func AccountIdFromContxt(ctx context.Context) (id.AccountId, error) {
	v := ctx.Value(accountIdKey{})
	aid, ok := v.(id.AccountId)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}

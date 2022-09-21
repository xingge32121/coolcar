/*
 * @Author: Wujiahuo
 * @Date: 2022-09-06 15:55:55
 * @LastEditors: OBKoro1
 * @LastEditTime: 2022-09-15 17:15:32
 * @FilePath: \wxc:\Users\Administrator\go\src\coolcar\server\rental\trip\trip.go
 * @Description:
 * Copyright (c) 2022 by Wujiahuo/bzlrobot, All Rights Reserved.
 */
package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/token"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	rentalpb.UnimplementedTripServiceServer
	Logger *zap.Logger
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (res *rentalpb.TripEntity, err error) {

	aid, err := token.AccountIdFromContxt(c)

	if err != nil {
		return nil, err
	}
	// s.Logger.Warn("cannot create trip", zap.Error(err), zap.String("account_id", aid.string()))
	s.Logger.Info("create Trip", zap.String("start", req.String()), zap.String("account_id", string(aid)))
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

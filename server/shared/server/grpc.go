package server

import (
	"coolcar/shared/token"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Name         string
	Tcp          string
	PublicFIle   string
	Logger       *zap.Logger
	RegisterFunc func(*grpc.Server)
}

func RunGRPCServer(g *GRPCConfig) error {
	nameFiled := zap.String("name", g.Name)
	lit, err := net.Listen("tcp", g.Tcp)
	if err != nil {
		g.Logger.Fatal("cannot link zap %c", nameFiled, zap.Error(err))
	}
	var opts []grpc.ServerOption
	if g.PublicFIle != "" {
		in, err := token.Interceptor(g.PublicFIle)
		if err != nil {
			g.Logger.Fatal("cannot Interceptor token  %c", nameFiled, zap.Error(err))
		}
		opts = append(opts, grpc.UnaryInterceptor(in))
	}

	s := grpc.NewServer(opts...)
	g.RegisterFunc(s)
	g.Logger.Info("server started at ", nameFiled, zap.String("addr", g.Tcp))
	// 对外开始服务
	return s.Serve(lit)

}

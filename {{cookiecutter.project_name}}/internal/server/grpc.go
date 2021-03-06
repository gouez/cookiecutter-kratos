package server

import (
	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.app_name}}/v1"
	"{{cookiecutter.module_name}}/internal/conf"
	"{{cookiecutter.module_name}}/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, {{cookiecutter.app_name}} *service.{{cookiecutter.app_name.capitalize()}}Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.Register{{cookiecutter.app_name.capitalize()}}Server(srv, {{cookiecutter.app_name}})
	return srv
}

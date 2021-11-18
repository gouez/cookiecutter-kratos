package service

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.app_name}}/v1"
	"{{cookiecutter.module_name}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// {{cookiecutter.app_name.capitalize()}}Service is a {{cookiecutter.app_name}} service.
type {{cookiecutter.app_name.capitalize()}}Service struct {
	v1.Unimplemented{{cookiecutter.app_name.capitalize()}}Server

	uc  *biz.{{cookiecutter.app_name.capitalize()}}Usecase
	log *log.Helper
}

// New{{cookiecutter.app_name.capitalize()}}Service new a {{cookiecutter.app_name}} service.
func New{{cookiecutter.app_name.capitalize()}}Service(uc *biz.{{cookiecutter.app_name.capitalize()}}Usecase, logger log.Logger) *{{cookiecutter.app_name.capitalize()}}Service {
	return &{{cookiecutter.app_name.capitalize()}}Service{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements {{cookiecutter.app_name}}.{{cookiecutter.app_name}}Server
func (s *{{cookiecutter.app_name.capitalize()}}Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())

	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

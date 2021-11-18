package data

import (
	"context"
	"{{cookiecutter.module_name}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type {{cookiecutter.app_name}}Repo struct {
	data *Data
	log  *log.Helper
}

// New{{cookiecutter.app_name.capitalize()}}Repo .
func New{{cookiecutter.app_name.capitalize()}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.app_name.capitalize()}}Repo {
	return &{{cookiecutter.app_name}}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{cookiecutter.app_name}}Repo) Create{{cookiecutter.app_name.capitalize()}}(ctx context.Context, g *biz.{{cookiecutter.app_name.capitalize()}}) error {
	return nil
}

func (r *{{cookiecutter.app_name}}Repo) Update{{cookiecutter.app_name.capitalize()}}(ctx context.Context, g *biz.{{cookiecutter.app_name.capitalize()}}) error {
	return nil
}

package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type {{cookiecutter.app_name.capitalize()}} struct {
	Hello string
}

type {{cookiecutter.app_name.capitalize()}}Repo interface {
	Create{{cookiecutter.app_name.capitalize()}}(context.Context, *{{cookiecutter.app_name.capitalize()}}) error
	Update{{cookiecutter.app_name.capitalize()}}(context.Context, *{{cookiecutter.app_name.capitalize()}}) error
}

type {{cookiecutter.app_name.capitalize()}}Usecase struct {
	repo {{cookiecutter.app_name.capitalize()}}Repo
	log  *log.Helper
}

func New{{cookiecutter.app_name.capitalize()}}Usecase(repo {{cookiecutter.app_name.capitalize()}}Repo, logger log.Logger) *{{cookiecutter.app_name.capitalize()}}Usecase {
	return &{{cookiecutter.app_name.capitalize()}}Usecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *{{cookiecutter.app_name.capitalize()}}Usecase) Create(ctx context.Context, g *{{cookiecutter.app_name.capitalize()}}) error {
	return uc.repo.Create{{cookiecutter.app_name.capitalize()}}(ctx, g)
}

func (uc *{{cookiecutter.app_name.capitalize()}}Usecase) Update(ctx context.Context, g *{{cookiecutter.app_name.capitalize()}}) error {
	return uc.repo.Update{{cookiecutter.app_name.capitalize()}}(ctx, g)
}

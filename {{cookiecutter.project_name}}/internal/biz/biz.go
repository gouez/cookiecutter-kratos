package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(New{{cookiecutter.app_name.capitalize()}}Usecase)

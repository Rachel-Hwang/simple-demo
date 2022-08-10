// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"helloworld/app/consumer/internal/biz"
	"helloworld/app/consumer/internal/conf"
	"helloworld/app/consumer/internal/data"
	"helloworld/app/consumer/internal/server"
	"helloworld/app/consumer/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

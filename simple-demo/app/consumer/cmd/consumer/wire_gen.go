// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/app/consumer/internal/biz"
	"helloworld/app/consumer/internal/conf"
	"helloworld/app/consumer/internal/data"
	"helloworld/app/consumer/internal/server"
	"helloworld/app/consumer/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger,db)
	if err != nil {
		return nil, nil, err
	}
	consumerRepo := data.NewProviderClient(dataData, logger)
	consumerUsecase := biz.NewConsumerUsecase(consumerRepo, logger)
	consumerService := service.NewConsumerService(consumerUsecase)
	grpcServer := server.NewGRPCServer(confServer, consumerService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

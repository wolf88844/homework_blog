// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"blog/internal/user/v1/biz"
	"blog/internal/user/v1/conf"
	"blog/internal/user/v1/data"
	"blog/internal/user/v1/server"
	"blog/internal/user/v1/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userServiceService := service.NewUserServiceService(userUseCase, logger)
	httpServer := server.NewHttpServer(confServer, userServiceService)
	grpcServer := server.NewGrpcServer(confServer, userServiceService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}

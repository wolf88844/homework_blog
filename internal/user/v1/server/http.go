package server

import (
	v1 "blog/api/user/v1"
	"blog/internal/user/v1/conf"
	"blog/internal/user/v1/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHttpServer(c *conf.Server,user *service.UserServiceService) *http.Server{
	var opts []http.ServerOption
	if c.Http.Network!=""{
		opts=append(opts,http.Network(c.Http.Network))
	}
	if c.Http.Addr!=""{
		opts=append(opts,http.Address(c.Http.Addr))
	}
	if c.Http.Timeout!=nil{
		opts=append(opts,http.Timeout(c.Http.Timeout.AsDuration()))
	}
	m:=http.Middleware(
		middleware.Chain(
			logging.Server(log.DefaultLogger),
			recovery.Recovery(),
			),
		)
	srv:=http.NewServer(opts...)
	srv.HandlePrefix("/",v1.NewUserServiceHandler(user,m))
	return srv
}

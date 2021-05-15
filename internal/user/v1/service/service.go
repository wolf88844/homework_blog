package service

import (
	pb "blog/api/user/v1"
	"blog/internal/user/v1/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet=wire.NewSet(NewUserServiceService)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
	logger *log.Helper
	u *biz.UserUseCase
}


package service

import (
	pb "blog/api/user/v1"
	"blog/internal/user/v1/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)



func NewUserServiceService(u *biz.UserUseCase,logger log.Logger) *UserServiceService {
	return &UserServiceService{
		logger: log.NewHelper("user",logger),
		u: u,
	}
}

func (s *UserServiceService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	s.logger.Infof("input data %v",req)
	err:=s.u.Create(ctx,&biz.User{
		Name: req.Name,
		Age:req.Age,
		Sex: req.Sex,
	})
	return &pb.CreateUserReply{}, err
}
func (s *UserServiceService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	s.logger.Infof("input data %v",req)
	err:=s.u.Update(ctx,req.Id,&biz.User{
		Name: req.Name,
		Age: req.Age,
		Sex: req.Sex,
	})
	return &pb.UpdateUserReply{}, err
}
func (s *UserServiceService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	s.logger.Infof("input data %v",req)
	err:=s.u.Delete(ctx,req.Id)
	return &pb.DeleteUserReply{}, err
}
func (s *UserServiceService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u,err:=s.u.Get(ctx,req.Id)
	if err!=nil{
		return nil,err
	}
	return &pb.GetUserReply{User: &pb.User{Id: u.Id,Name: u.Name,Age: u.Age,Sex: u.Sex}}, nil
}
func (s *UserServiceService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	us,err:=s.u.List(ctx)
	reply:=&pb.ListUserReply{}
	for _,u:=range us{
		reply.User=append(reply.User,&pb.User{
			Id: u.Id,
			Name: u.Name,
			Age: u.Age,
			Sex: u.Sex,
		})
	}
	return reply, err
}

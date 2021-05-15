package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type User struct {
	Id int64
	Name string
	Age uint32
	Sex uint32
	CreatedAt time.Time
	UpdateAt time.Time
}

type UserRepo interface {
	GetUser(ctx context.Context,id int64) (*User,error)
	ListUser(ctx context.Context) ([]*User,error)
	CreateUser(ctx context.Context,user *User)error
	UpdateUser(ctx context.Context,id int64,user *User)error
	DeleteUser(ctx context.Context,id int64)error
}

type UserUseCase struct {
	repo UserRepo
	logger log.Logger
}

func NewUserUseCase(repo UserRepo,logger log.Logger) *UserUseCase{
	return &UserUseCase{
		repo: repo,
		logger: logger,
	}
}

func (u *UserUseCase) Get(ctx context.Context,id int64)(us *User,err error){
	us,err=u.repo.GetUser(ctx,id)
	if err!=nil{
		return nil,err
	}
	return us,nil
}

func (u *UserUseCase) List(ctx context.Context)(us []*User,err error){
	us,err=u.repo.ListUser(ctx)
	if err!=nil{
		return nil, err
	}
	return us,nil
}

func(u *UserUseCase) Create(ctx context.Context,user *User)error{
	return u.repo.CreateUser(ctx,user)
}

func(u *UserUseCase) Update(ctx context.Context,id int64,user *User)error{
	return u.repo.UpdateUser(ctx,id,user)
}

func (u *UserUseCase) Delete(ctx context.Context,id int64)error{
	return u.repo.DeleteUser(ctx,id)
}
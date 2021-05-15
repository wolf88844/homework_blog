package data

import (
	"blog/internal/user/v1/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type userRepo struct {
	data *Data
	log *log.Helper
}
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log: log.NewHelper("user repo",logger),
	}
}

func (u *userRepo)GetUser(ctx context.Context,id int64) (*biz.User,error){
	user,err:=u.data.db.User.Get(ctx,id)
	if err!=nil{
		return nil,err
	}
	return &biz.User{
		Id: user.ID,
		Name: user.Name,
		Age: user.Age,
		Sex: user.Sex,
		CreatedAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	},nil
}
func (u *userRepo) ListUser(ctx context.Context) ([]*biz.User,error){
	users,err:=u.data.db.User.Query().All(ctx)
	if err!=nil{
		return nil,err
	}
	rv:=make([]*biz.User,0)
	for _,user:=range users{
		rv = append(rv,&biz.User{
			Id: user.ID,
			Name: user.Name,
			Age: user.Age,
			Sex: user.Sex,
			CreatedAt: user.CreateAt,
			UpdateAt: user.UpdateAt,
		})
	}
	return rv,nil
}
func (u *userRepo) CreateUser(ctx context.Context,user *biz.User)error{
	_,err:=u.data.db.User.Create().SetName(user.Name).SetAge(user.Age).SetSex(user.Sex).Save(ctx)
	return err
}
func (u *userRepo) UpdateUser(ctx context.Context,id int64,user *biz.User)error{
	ouser,err:=u.data.db.User.Get(ctx,id)
	if err!=nil{
		return err
	}
	_,err=ouser.Update().SetName(user.Name).SetAge(user.Age).SetSex(user.Sex).SetUpdateAt(time.Now()).Save(ctx)
	return err
}
func (u *userRepo) DeleteUser(ctx context.Context,id int64)error{
	return u.data.db.User.DeleteOneID(id).Exec(ctx)
}
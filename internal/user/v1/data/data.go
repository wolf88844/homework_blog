package data

import (
	"blog/internal/user/v1/conf"
	"blog/internal/user/v1/data/ent"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/go-sql-driver/mysql"
)

var ProviderSet = wire.NewSet(NewData,NewUserRepo)

type Data struct {
	db *ent.Client
}

func NewData(conf *conf.Data,logger log.Logger)(*Data,func(),error){
	log:=log.NewHelper("data",logger)
	client,err:=ent.Open(conf.Database.Driver,conf.Database.Source)
	if err!=nil{
		log.Errorf("failed opening connection to :%v",err)
		return nil,nil,err
	}
	if err:=client.Schema.Create(context.Background());err!=nil{
		log.Errorf("failed create schema resources: %v",err)
		return nil, nil, err
	}

	d:=&Data{
		db: client,
	}
	return d, func() {
		log.Info("message","closing the data resource")
		if err:=d.db.Close();err!=nil{
			log.Error(err)
		}
	},nil
}

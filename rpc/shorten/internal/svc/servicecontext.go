package svc

import (
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"shorturl/rpc/model"
	"shorturl/rpc/shorten/internal/config"
)

type ServiceContext struct {
	c config.Config
	Model *model.ShorturlModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:     c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
	}
}

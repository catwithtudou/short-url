package svc

import (
	"github.com/tal-tech/go-zero/rpcx"
	"shorturl/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Shortener rpcx.Client
	Expander rpcx.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Shortener: rpcx.MustNewClient(c.Shortener),
		Expander:  rpcx.MustNewClient(c.Expander),
	}
}

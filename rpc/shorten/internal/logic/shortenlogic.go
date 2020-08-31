package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/hash"
	"shorturl/rpc/model"

	"shorturl/rpc/shorten/internal/svc"
	shorten "shorturl/rpc/shorten/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

const keyLen = 6

type ShortenLogic struct {
	ctx context.Context
	logx.Logger
	model *model.ShorturlModel
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		model:  svcCtx.Model,
	}
}

func (l *ShortenLogic) Shorten(in *shorten.ShortenReq) (*shorten.ShortenResp, error) {
	key:=hash.Md5Hex([]byte(in.Url))[:keyLen]
	_,err:=l.model.Insert(model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	})
	if err!=nil{
		return nil,err
	}

	return &shorten.ShortenResp{
		Key: key,
	}, nil
}

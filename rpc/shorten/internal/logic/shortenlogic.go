package logic

import (
	"context"

	"shorturl/rpc/shorten/internal/svc"
	shorten "shorturl/rpc/shorten/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx context.Context
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *shorten.ShortenReq) (*shorten.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &shorten.ShortenResp{}, nil
}

package logic

import (
	"context"
	"github.com/tal-tech/go-zero/rpcx"
	"shorturl/rpc/shorten/shortener"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx context.Context
	logx.Logger
	shortener rpcx.Client
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShortenLogic {
	return ShortenLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		shortener: svcCtx.Shortener,
	}
}

func (l *ShortenLogic) Shorten(req types.ShortenReq) (*types.ShortenResp, error) {
	resp, err := shortener.NewShortener(l.shortener).Shorten(l.ctx, &shortener.ShortenReq{
		Url: req.Url,
	})
	if err != nil {
		return nil, err
	}

	return &types.ShortenResp{
		ShortUrl: resp.Key,
	}, nil
}

package logic

import (
	"context"
	"shorturl/rpc/model"

	"shorturl/rpc/expand/internal/svc"
	expand "shorturl/rpc/expand/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx context.Context
	logx.Logger
	model *model.ShorturlModel
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		Logger: logx.WithContext(ctx),
		model:  svcCtx.Model,
	}
}

func (l *ExpandLogic) Expand(in *expand.ExpandReq) (*expand.ExpandResp, error) {
	res,err:=l.model.FindOne(in.Key)
	if err!=nil{
		return nil,err
	}

	return &expand.ExpandResp{
		Url: res.Url,
	}, nil
}



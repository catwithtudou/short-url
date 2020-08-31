package logic

import (
	"context"
	"github.com/tal-tech/go-zero/rpcx"
	"shorturl/rpc/expand/expander"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx context.Context
	logx.Logger
	expander rpcx.Client
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandLogic {
	return ExpandLogic{
		ctx:      ctx,
		Logger:   logx.WithContext(ctx),
		expander: svcCtx.Expander,
	}
}

func (l *ExpandLogic) Expand(req types.ExpandReq) (*types.ExpandResp, error) {
	resp,err:=expander.NewExpander(l.expander).Expand(l.ctx,&expander.ExpandReq{
		Key: req.Key,
	})
	if err!=nil{
		return nil,err
	}

	return &types.ExpandResp{
		Url:resp.Url,
	},nil



}

package logic

import (
	"context"

	"bot/internal/svc"
	"bot/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BotLogic {
	return &BotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BotLogic) Bot(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

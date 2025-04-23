package normal

import (
	"context"

	"bot/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type SyncLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncLogic {
	return &SyncLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncLogic) Sync() error {
	// todo: add your logic here and delete this line

	return nil
}

package consumer

import (
	"bot/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MsgConsumer struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMsgConsumer(context context.Context, serverCtx *svc.ServiceContext) Consumer {
	return &MsgConsumer{
		ctx:    context,
		svcCtx: serverCtx,
	}
}

// Consume implements Consumer.
func (m *MsgConsumer) Consume() error {
	msg := <-m.svcCtx.MsgChannel
	m.Info(msg)
	return nil
}

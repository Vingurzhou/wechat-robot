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
		Logger: logx.WithContext(context),
		ctx:    context,
		svcCtx: serverCtx,
	}
}

// Consume implements Consumer.
func (m *MsgConsumer) Consume() error {
	for {
		select {
		case msg := <-m.svcCtx.MsgChannel:
			if msg != nil {
				m.Infof("%+v", msg)
			}
		case <-m.ctx.Done():
			return nil
		}
	}
}

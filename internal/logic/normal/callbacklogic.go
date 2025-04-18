package normal

import (
	"context"

	"bot/internal/dao/model"
	"bot/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CallbackLogic) Callback() error {
	msg := &model.Msg{
		MessageID:    "",
		CallbackURL:  "",
		Status:       "",
		Payload:      "",
		RetryCount:   0,
		ErrorMessage: "",
	}
	err := l.svcCtx.Query.WithContext(l.ctx).Msg.Create(msg)
	if err != nil {
		l.Error(err)
		return err
	}
	l.svcCtx.MsgChannel <- msg
	return nil
}

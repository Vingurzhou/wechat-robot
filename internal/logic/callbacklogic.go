package logic

import (
	"context"

	"bot/internal/dao/model"
	"bot/internal/svc"
	"bot/internal/types"

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

func (l *CallbackLogic) Callback(req *types.CallBackReq) error {
	// todo: add your logic here and delete this line
	msg := &model.Msg{
		MsgID:        0,
		FromUserName: req.Data.FromUserName.String,
		ToUserName:   req.Data.ToUserName.String,
		MsgType:      int32(req.Data.MsgType),
		Content:      req.Data.Content.String,
		Status:       int32(req.Data.Status),
		ImgStatus:    int32(req.Data.ImgStatus),
		ImgBufILen:   int32(req.Data.ImgBuf.ILen),
		CreateTime:   int32(req.Data.CreateTime),
		MsgSource:    req.Data.MsgSource,
		PushContent:  req.Data.PushContent,
		NewMsgID:     int32(req.Data.NewMsgID),
		MsgSeq:       int32(req.Data.MsgSeq),
		Wxid:         req.Wxid,
		Appid:        req.Appid,
		TypeName:     req.TypeName,
	}
	err := l.svcCtx.Query.Msg.WithContext(l.ctx).Create(msg)
	if err != nil {
		l.Error(err)
		return err
	}
	l.svcCtx.MsgChannel <- msg
	return nil
}

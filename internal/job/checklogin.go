package job

import (
	"bot/internal/svc"

	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckLoginJob struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}
type GetTokenIdResp struct {
	Ret  int64  `json:"ret"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
type GetLoginQrCodeResp struct {
	Ret  int64  `json:"ret"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	QrData      string `json:"qrData"`
	QrImgBase64 string `json:"qrImgBase64"`
	UUID        string `json:"uuid"`
	AppID       string `json:"appId"`
}

// Run implements Job.
func (j *CheckLoginJob) Run() {
}

func NewCheckLoginJob(ctx context.Context, svcCtx *svc.ServiceContext) Job {
	return &CheckLoginJob{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

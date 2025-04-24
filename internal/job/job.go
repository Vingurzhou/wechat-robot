package job

import (
	"bot/internal/svc"
	"fmt"
)

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
type Job interface {
	Run()
}

func RegisterJobs(svcCtx *svc.ServiceContext) {
	// ctx := context.Background()

	// NewLoginJob(ctx, svcCtx).Run()
	// svcCtx.Cron.AddFunc("@every 5s", NewCheckLoginJob(ctx, svcCtx).Run)

	fmt.Printf("Starting job at %s\n", svcCtx.Config.Host)
	svcCtx.Cron.Start()
}

package job

import (
	"bot/internal/svc"

	"context"
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginJob struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Run implements Job.
func (j *LoginJob) Run() {
	// curl --location --request POST 'http://127.0.0.1:2531/v2/api/tools/getTokenId'
	resp, err := j.svcCtx.HttpCli.Do("http://127.0.0.1:2531/v2/api/tools/getTokenId", "POST", "", nil)
	if err != nil {
		j.Error(err)
		return
	}
	var getTokenIdResp GetTokenIdResp
	err = json.Unmarshal(resp, &getTokenIdResp)
	if err != nil {
		j.Error(err)
		return
	}
	j.svcCtx.Config.GEWE.Token = getTokenIdResp.Data
	// curl --location --request POST 'http://127.0.0.1:2531/v2/api/login/getLoginQrCode' \
	// --header 'X-GEWE-TOKEN: 447d53fdf6354f1791531049fdba865c' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//	    "appId": ""
	// }'
	resp, err = j.svcCtx.HttpCli.Do("http://127.0.0.1:2531/v2/api/login/getLoginQrCode", "POST",
		`{"appId": ""}`,
		map[string]string{
			"X-GEWE-TOKEN": j.svcCtx.Config.GEWE.Token,
			"Content-Type": "application/json",
		})
	if err != nil {
		j.Error(err)
		return
	}
	var getLoginQrCodeResp GetLoginQrCodeResp
	err = json.Unmarshal(resp, &getLoginQrCodeResp)
	if err != nil {
		j.Error(err)
		return
	}
	list := strings.Split(getLoginQrCodeResp.Data.QrImgBase64, ",")
	if len(list) != 2 {
		j.Error("out of index,关闭代理")
		return
	}
	imgData, err := base64.StdEncoding.DecodeString(list[1])
	if err != nil {
		j.Error(err)
		return
	}
	err = os.WriteFile("qrImage.jpg", imgData, 0644)
	if err != nil {
		j.Error(err)
		return
	}
}

func NewLoginJob(ctx context.Context, svcCtx *svc.ServiceContext) Job {
	return &LoginJob{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

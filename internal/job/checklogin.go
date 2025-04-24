package job

import (
	"bot/internal/svc"
	"encoding/base64"
	"fmt"

	"context"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

// Run implements Job.
func (j *CheckLoginJob) Run() {
	// curl --location --request POST 'http://127.0.0.1:2531/v2/api/login/checkLogin' \
	// --header 'X-GEWE-TOKEN: 447d53fdf6354f1791531049fdba865c' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//     "appId": "{{appid}}",
	//     "uuid": "4dUAplBjL01QqBE6t4hK"
	// }'
	u := uuid.New()
	shortUUID := base64.RawURLEncoding.EncodeToString(u[:])
	resp, err := j.svcCtx.HttpCli.Do("http://127.0.0.1:2531/v2/api/login/checkLogin", "POST",
		fmt.Sprintf(`{"appId":"%s","uuid": "%s"}`, j.svcCtx.Config.GEWE.AppId, shortUUID),
		map[string]string{
			"X-GEWE-TOKEN": j.svcCtx.Config.GEWE.Token,
			"Content-Type": "application/json",
		})
	if err != nil {
		j.Error(err)
		return
	}
	j.Info(string(resp))
}

func NewCheckLoginJob(ctx context.Context, svcCtx *svc.ServiceContext) Job {
	return &CheckLoginJob{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

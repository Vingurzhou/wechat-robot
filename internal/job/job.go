package job

import (
	"bot/internal/svc"
	"context"
	"fmt"
)

type Job interface {
	Run()
}

func RegisterJobs(svcCtx *svc.ServiceContext) {
	ctx := context.Background()

	NewLoginJob(ctx, svcCtx).Run()
	svcCtx.Cron.AddFunc("@every 5s", NewCheckLoginJob(ctx, svcCtx).Run)

	fmt.Printf("Starting job at %s\n", svcCtx.Config.Host)
	svcCtx.Cron.Start()
}

package svc

import (
	"bot/internal/config"
	"bot/internal/dao/model"
	"bot/internal/dao/query"
	"bot/internal/middleware"

	"context"
	"fmt"

	"github.com/Vingurzhou/pkg/db"
	"github.com/Vingurzhou/pkg/httpz"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/rest"
	"google.golang.org/genai"
	"gorm.io/driver/sqlite"
)

type ServiceContext struct {
	Config             config.Config
	CallbackMiddleware rest.Middleware
	Query              *query.Query
	MsgChannel         chan *model.Msg
	HttpCli            *httpz.WrapperHttpCli
	Cron               *cron.Cron
	GenAICli           *genai.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// gentool -db sqlite -dsn "bot.db" -tables "" -outPath "./internal/dao/query"
	return &ServiceContext{
		Config:             c,
		CallbackMiddleware: middleware.NewCallbackMiddleware().Handle,
		Query:              query.Use(db.NewGormDB(sqlite.Open(c.DSN))),
		MsgChannel:         make(chan *model.Msg, 10),
		HttpCli:            httpz.NewHttpCli(),
		Cron: cron.New(cron.WithChain(
			cron.Recover(cron.DefaultLogger),
		)),
		GenAICli: NewGenAICLli(c.GenAI.ApiKey),
	}
}

func NewGenAICLli(apiKey string) *genai.Client {
	ctx := context.Background()
	cli, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		fmt.Println(err)
	}
	return cli
}

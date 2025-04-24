package svc

import (
	"bot/internal/config"
	"bot/internal/dao/model"
	"bot/internal/dao/query"
	"bot/internal/middleware"

	"github.com/Vingurzhou/pkg/db"
	"github.com/Vingurzhou/pkg/httpz"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/sqlite"
)

type ServiceContext struct {
	Config             config.Config
	CallbackMiddleware rest.Middleware
	Query              *query.Query
	MsgChannel         chan *model.Msg
	HttpCli            *httpz.WrapperHttpCli
	Cron               *cron.Cron
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
	}
}

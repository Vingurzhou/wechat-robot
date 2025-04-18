package consumer

import (
	"bot/internal/svc"
	"context"
)

type Consumer interface {
	Consume() error
}

func RegisterConsumers(serverCtx *svc.ServiceContext) {
	go NewMsgConsumer(context.Background(), serverCtx).Consume()
}

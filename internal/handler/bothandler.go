package handler

import (
	"net/http"

	"bot/internal/logic"
	"bot/internal/svc"
	"bot/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBotLogic(r.Context(), svcCtx)
		resp, err := l.Bot(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

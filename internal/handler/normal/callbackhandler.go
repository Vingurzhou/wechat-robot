package normal

import (
	"net/http"

	"bot/internal/logic/normal"
	"bot/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := normal.NewCallbackLogic(r.Context(), svcCtx)
		err := l.Callback()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

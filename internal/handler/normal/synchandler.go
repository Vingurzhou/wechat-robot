package normal

import (
	"net/http"

	"bot/internal/logic/normal"
	"bot/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := normal.NewSyncLogic(r.Context(), svcCtx)
		err := l.Sync()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}

package handler

import (
	"net/http"

	"business-platform/app/notify/internal/logic"
	"business-platform/app/notify/internal/svc"
	"business-platform/app/notify/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TxNotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewTxNotifyLogic(r.Context(), svcCtx)
		resp, err := l.TxNotify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

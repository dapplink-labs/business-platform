package withdraw

import (
	"net/http"

	"business-platform/app/exchange/internal/logic/withdraw"
	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 创建提现订单
func WithdrawHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := withdraw.NewWithdrawLogic(r.Context(), svcCtx)
		resp, err := l.Withdraw(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

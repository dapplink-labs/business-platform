package withdraw

import (
	"net/http"

	"business-platform/app/exchange/internal/logic/withdraw"
	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询提现订单详情
func GetWithdrawOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := withdraw.NewGetWithdrawOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetWithdrawOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

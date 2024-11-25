package withdraw

import (
	"net/http"

	"business-platform/app/exchange/internal/logic/withdraw"
	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 查询提现订单列表
func ListWithdrawOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WithdrawListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := withdraw.NewListWithdrawOrdersLogic(r.Context(), svcCtx)
		resp, err := l.ListWithdrawOrders(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

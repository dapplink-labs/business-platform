package withdraw

import (
	"context"

	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListWithdrawOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询提现订单列表
func NewListWithdrawOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListWithdrawOrdersLogic {
	return &ListWithdrawOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListWithdrawOrdersLogic) ListWithdrawOrders(req *types.WithdrawListRequest) (resp *types.WithdrawListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

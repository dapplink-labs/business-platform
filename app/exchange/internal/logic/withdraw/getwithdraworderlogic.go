package withdraw

import (
	"context"

	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWithdrawOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询提现订单详情
func NewGetWithdrawOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWithdrawOrderLogic {
	return &GetWithdrawOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWithdrawOrderLogic) GetWithdrawOrder(req *types.WithdrawOrderRequest) (resp *types.WithdrawOrderInfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

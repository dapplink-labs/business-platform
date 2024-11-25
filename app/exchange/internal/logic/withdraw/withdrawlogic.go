package withdraw

import (
	"context"

	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WithdrawLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建提现订单
func NewWithdrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WithdrawLogic {
	return &WithdrawLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WithdrawLogic) Withdraw(req *types.WithdrawRequest) (resp *types.WithdrawResponse, err error) {
	resp = &types.WithdrawResponse{
		OrderId: 123456789,
		TxHash:  "1111111",
	}
	return resp, nil
}

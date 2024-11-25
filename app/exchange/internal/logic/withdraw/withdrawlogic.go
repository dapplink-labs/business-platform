package withdraw

import (
	"business-platform/app/exchange/internal/svc"
	"business-platform/app/exchange/internal/types"
	"context"

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
	// 使用数据库
	//user, err := l.svcCtx.DB.FindOne(l.ctx, req.Uid)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// 使用 Redis
	//key := fmt.Sprintf("withdraw:%d", req.Uid)
	//exists, err := l.svcCtx.Redis.Exists(key)
	//if err != nil {
	//	return nil, err
	//}

	// 使用 Multichain 客户端
	//txResp, err := l.svcCtx.MultichainClient.CreateUnSignTransaction(l.ctx, &proto.UnSignWithdrawTransactionRequest{
	// ...
	//})

	resp = &types.WithdrawResponse{
		OrderId: 123456789,
		TxHash:  "1111111",
	}

	return resp, nil
}

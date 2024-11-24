package logic

import (
	"context"

	"business-platform/app/notify/internal/svc"
	"business-platform/app/notify/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TxNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTxNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TxNotifyLogic {
	return &TxNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TxNotifyLogic) TxNotify(req *types.NotifyRequest) (resp *types.NotifyResponse, err error) {
	// 打印完整的请求参数
	l.Logger.Infof("收到交易通知请求，交易数量: %d", len(req.Txn))

	// 详细打印每笔交易信息
	for i, tx := range req.Txn {
		l.Logger.Infof("交易 #%d 详情:", i+1)
		l.Logger.Infof("  区块哈希: %s", tx.BlockHash)
		l.Logger.Infof("  区块高度: %d", tx.BlockNumber)
		l.Logger.Infof("  交易哈希: %s", tx.Hash)
		l.Logger.Infof("  发送地址: %s", tx.FromAddress)
		l.Logger.Infof("  接收地址: %s", tx.ToAddress)
		l.Logger.Infof("  交易金额: %s", tx.Value)
		l.Logger.Infof("  交易费用: %s", tx.Fee)
		l.Logger.Infof("  交易类型: %s", getTxTypeName(tx.TxType))
		l.Logger.Infof("  确认数: %d", tx.Confirms)

		// 打印可选字段（如果存在）
		if tx.TokenAddress != "" {
			l.Logger.Infof("  代币地址: %s", tx.TokenAddress)
		}
		if tx.TokenId != "" {
			l.Logger.Infof("  代币ID: %s", tx.TokenId)
		}
		if tx.TokenMeta != "" {
			l.Logger.Infof("  代币元数据: %s", tx.TokenMeta)
		}
	}

	return &types.NotifyResponse{
		Success: true,
	}, nil
}

func getTxTypeName(txType string) string {
	switch txType {
	case "0":
		return "充值"
	case "1":
		return "提现"
	case "2":
		return "归集"
	case "3":
		return "热转冷"
	case "4":
		return "冷转热"
	default:
		return "未知类型"
	}
}

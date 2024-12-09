package multichain

import (
	"errors"
)

type TransactionType string

const (
	TxTypeUnKnow     TransactionType = "unknow"
	TxTypeDeposit    TransactionType = "deposit"
	TxTypeWithdraw   TransactionType = "withdraw"
	TxTypeCollection TransactionType = "collection"
	TxTypeHot2Cold   TransactionType = "hot2cold"
	TxTypeCold2Hot   TransactionType = "cold2hot"
)

func ParseTransactionType(s string) (TransactionType, error) {
	switch s {
	case string(TxTypeDeposit):
		return TxTypeDeposit, nil
	case string(TxTypeWithdraw):
		return TxTypeWithdraw, nil
	case string(TxTypeCollection):
		return TxTypeCollection, nil
	case string(TxTypeHot2Cold):
		return TxTypeHot2Cold, nil
	case string(TxTypeCold2Hot):
		return TxTypeCold2Hot, nil
	default:
		return TxTypeUnKnow, errors.New("unknown transaction type")
	}
}

type TokenType string

const (
	TokenTypeETH     TokenType = "ETH"
	TokenTypeERC20   TokenType = "ERC20"
	TokenTypeERC721  TokenType = "ERC721"
	TokenTypeERC1155 TokenType = "ERC1155"
)

type CreateUnSignTransactionRequest struct {
	ChainId   string
	Chain     string
	TxType    TransactionType
	TokenType TokenType
	TxETH     *UnSignTransactionRequestByETH
	TxERC     *UnSignTransactionRequestByERC
}

type UnSignTransactionRequestByETH struct {
	From  string
	To    string
	Value string
}

type UnSignTransactionRequestByERC struct {
	From            string
	To              string
	Value           string
	ContractAddress string
	TokenId         string
	TokenMeta       string
}

type CreateSignedTransactionRequest struct {
	Chain         string
	ChainId       string
	TransactionId string
	Signature     string
	TxType        TransactionType
}

package multichain

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"business-platform/app/common/grpcclient/multichain/proto"
)

const (
	multichainUrl    = "127.0.0.1:8987"
	notifyUrl        = "http://127.0.0.1:8001"
	CurrentRequestId = "1"
	CurrentChainId   = "17000"
	CurrentChain     = "ethereum"
)

func Test_GrpcClient_BusinessRegister(t *testing.T) {
	// 创建 gRPC 客户端
	client, err := NewGrpcClient(multichainUrl)
	if err != nil {
		t.Fatalf("创建 gRPC 客户端失败: %v", err)
	}
	defer func(client GrpcClient) {
		err := client.Close()
		if err != nil {
			t.Fatalf("gRPC conn.Close: %v", err)
		}
	}(client)

	// 创建上下文
	ctx := context.Background()

	// 调用 GetSupportChains 方法
	err = client.BusinessRegister(ctx, CurrentRequestId, notifyUrl)
	if err != nil {
		t.Fatalf("调用 GetSupportChains 失败: %v", err)
	}
	t.Logf("success")
}

func TestExportAddressesByPublicKeys(t *testing.T) {
	// 创建 gRPC 客户端
	client, err := NewGrpcClient(multichainUrl)
	if err != nil {
		t.Fatalf("创建 gRPC 客户端失败: %v", err)
	}
	defer func(client GrpcClient) {
		err := client.Close()
		if err != nil {
			t.Fatalf("gRPC conn.Close: %v", err)
		}
	}(client)

	// 创建上下文
	ctx := context.Background()

	// 准备请求参数
	publicKeyList := []*proto.PublicKey{
		{
			Type:      "eoa",
			PublicKey: "0422d39a1208b314bbbae7545c0b415167386d448ba9777b526e56d458db2f9f70d72f89373b7f53dfc9f0ff6aa55ae736fe2160d7ddd8be470250dd23fae9b0bc",
		},
		{
			Type:      "hot",
			PublicKey: "047b40b2707107640641c983919bfff36946849df442564a9bccc577680898c7449546e54eb4a2f63bfe8f061c9d7b7f6669a3154479746cc8e0d7c6ca2d490e6a",
		},
		{
			Type:      "cold",
			PublicKey: "04a84731792f6cdfb67d1c591d090844af1ecf4bb73193c7e389fedbdfc088564b3a1f9372781a0d92feb4251b3059f050873ada6ac2cb9b5b40f709900ce2a65d",
		},
	}

	// Call ExportAddressesByPublicKeys method
	addresses, err := client.ExportAddressesByPublicKeys(ctx, CurrentRequestId, publicKeyList)
	if err != nil {
		t.Fatalf("Failed to call ExportAddressesByPublicKeys: %v", err)
	}

	// Validate response
	if addresses == nil {
		t.Error("Received nil response")
		return
	}

	// Log the response
	respJson, err := json.Marshal(addresses)
	if err != nil {
		t.Fatalf("Failed to marshal response: %v", err)
	}
	t.Logf("Response: %s", respJson)
}

func Test_GrpcClient_CreateUnSignTransaction(t *testing.T) {
	// 创建 gRPC 客户端
	client, err := NewGrpcClient(multichainUrl)
	if err != nil {
		t.Fatalf("创建 gRPC 客户端失败: %v", err)
	}
	defer func(client GrpcClient) {
		err := client.Close()
		if err != nil {
			t.Fatalf("gRPC conn.Close: %v", err)
		}
	}(client)

	// 创建上下文
	ctx := context.Background()

	// 准备请求参数
	request := &CreateUnSignTransactionRequest{
		ChainId: CurrentChainId,
		Chain:   CurrentChain,
		TxType:  "collection",
		TxETH: &UnSignTransactionRequestByETH{
			From:  "0xD79053a14BC465d9C1434d4A4fAbdeA7b6a2A94b",
			To:    "0xDf894d39f6b33763bf55582Bb7A8b5515bccD982",
			Value: "10000000000000000", // 0.01 ETH
		},
	}

	// 调用方法
	txId, unSignTx, err := client.CreateUnSignTransaction(ctx, CurrentRequestId, request)
	if err != nil {
		t.Fatalf("CreateUnSignTransaction failed: %v", err)
	}

	// 验证响应
	assert.NotEmpty(t, txId, "Transaction ID should not be empty")
	assert.NotEmpty(t, unSignTx, "Unsigned transaction should not be empty")
	t.Logf("Transaction ID: %s, Unsigned Transaction: %s", txId, unSignTx)
}

func Test_GrpcClient_BuildSignedTransaction(t *testing.T) {
	// 创建 gRPC 客户端
	client, err := NewGrpcClient(multichainUrl)
	if err != nil {
		t.Fatalf("创建 gRPC 客户端失败: %v", err)
	}
	defer func(client GrpcClient) {
		err := client.Close()
		if err != nil {
			t.Fatalf("gRPC conn.Close: %v", err)
		}
	}(client)

	// 创建上下文
	ctx := context.Background()

	// 准备请求参数
	request := &CreateSignedTransactionRequest{
		Chain:         CurrentChain,
		ChainId:       CurrentChainId,
		TransactionId: "e7e656a5-3d37-4232-a8a7-4c79f6a864fc",
		Signature:     "aa8f64798957645c6e484716f856e7c87f0b5fdb1f7d2dd4367c472cf426849d3f9394bc473126b1d078ab9356245304c36f10474c38bef58f731d80ecbd532101",
		TxType:        "withdraw",
	}
	// 调用方法
	signedTx, err := client.BuildSignedTransaction(ctx, CurrentRequestId, request)
	if err != nil {
		t.Fatalf("BuildSignedTransaction failed: %v", err)
	}

	// 验证响应
	assert.NotEmpty(t, signedTx, "Signed transaction should not be empty")
	t.Logf("Signed Transaction: %s", signedTx)
}

func Test_GrpcClient_SetTokenAddress(t *testing.T) {
	// 创建 gRPC 客户端
	client, err := NewGrpcClient(multichainUrl)
	if err != nil {
		t.Fatalf("创建 gRPC 客户端失败: %v", err)
	}
	defer func(client GrpcClient) {
		err := client.Close()
		if err != nil {
			t.Fatalf("gRPC conn.Close: %v", err)
		}
	}(client)

	// 创建上下文
	ctx := context.Background()

	// 准备请求参数
	tokenList := []*proto.Token{
		{
			Decimals:      18,
			Address:       "0x789",
			TokenName:     "TEST",
			CollectAmount: "1000000000000000000",
			ColdAmount:    "500000000000000000",
		},
	}

	// 调用方法
	err = client.SetTokenAddress(ctx, CurrentRequestId, tokenList)
	if err != nil {
		t.Fatalf("SetTokenAddress failed: %v", err)
	}

	// 验证响应
	t.Logf("SetTokenAddress succeeded")
}

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
	notifyUrl        = "127.0.0.1:8001/dapplink/notify"
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

	// 准备请求参数
	req := &proto.BusinessRegisterRequest{
		RequestId: CurrentRequestId,
		NotifyUrl: notifyUrl,
	} // 根据实际需求设置请求参数

	// 调用 GetSupportChains 方法
	resp, err := client.BusinessRegister(ctx, req)
	if err != nil {
		t.Fatalf("调用 GetSupportChains 失败: %v", err)
	}

	// 验证响应
	if resp == nil {
		t.Error("获取到的响应为空")
		return
	}

	respJson, _ := json.Marshal(resp)
	t.Logf("响应: %s", respJson)
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
	req := &proto.ExportAddressesRequest{
		RequestId: CurrentRequestId,
		PublicKeys: []*proto.PublicKey{
			{
				Type:      1,
				PublicKey: "048318535b54105d4a7aae60c08fc45f9687181b4fdfc625bd1a753fa7397fed753547f11ca8696646f2f3acb08e31016afac23e630c5d11f59f61fef57b0d2aa5",
			},
		},
	}

	// 调用 GetSupportChains 方法
	resp, err := client.ExportAddressesByPublicKeys(ctx, req)
	if err != nil {
		t.Fatalf("调用 ExportAddressesByPublicKeys 失败: %v", err)
	}

	// 验证响应
	if resp == nil {
		t.Error("获取到的响应为空")
		return
	}
	// {"Code":1,"msg":"generate address success","addresses":[{"type":1}]}
	respJson, _ := json.Marshal(resp)
	t.Logf("响应: %s", respJson)
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
	req := &proto.UnSignWithdrawTransactionRequest{
		RequestId: CurrentRequestId,
		ChainId:   CurrentChainId,
		Chain:     CurrentChain,
		From:      "0xD79053a14BC465d9C1434d4A4fAbdeA7b6a2A94b",
		To:        "0xDf894d39f6b33763bf55582Bb7A8b5515bccD982",
		// 1 eth
		//Value: "1000000000000000000",
		// 0.01 eth
		Value:           "10000000000000000",
		ContractAddress: "0x00",
		TokenId:         "",
		TokenMeta:       "",
		TxType:          "withdraw",
	}

	// 调用方法
	resp, err := client.CreateUnSignTransaction(ctx, req)
	if err != nil {
		t.Fatalf("CreateUnSignTransaction failed: %v", err)
	}

	// 验证响应
	assert.NotNil(t, resp, "Response should not be nil")
	respJSON, err := json.Marshal(resp)
	assert.NoError(t, err, "Failed to marshal response to JSON")
	t.Logf("Response: %s", respJSON)
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
	req := &proto.SignedWithdrawTransactionRequest{
		ConsumerToken: "test-token",
		RequestId:     CurrentRequestId,
		Chain:         CurrentChain,
		ChainId:       CurrentChainId,
		TransactionId: "e7e656a5-3d37-4232-a8a7-4c79f6a864fc",
		Signature:     "aa8f64798957645c6e484716f856e7c87f0b5fdb1f7d2dd4367c472cf426849d3f9394bc473126b1d078ab9356245304c36f10474c38bef58f731d80ecbd532101",
		TxType:        "withdraw",
	}

	// 调用方法
	resp, err := client.BuildSignedTransaction(ctx, req)
	if err != nil {
		t.Fatalf("BuildSignedTransaction failed: %v", err)
	}

	// 验证响应
	assert.NotNil(t, resp, "Response should not be nil")
	respJSON, err := json.Marshal(resp)
	assert.NoError(t, err, "Failed to marshal response to JSON")
	t.Logf("Response: %s", respJSON)
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
	req := &proto.SetTokenAddressRequest{
		Code: proto.ReturnCode_SUCCESS,
		TokenList: []*proto.Token{
			{
				Decimals:      18,
				Address:       "0x789",
				TokenName:     "TEST",
				CollectAmount: "1000000000000000000",
				ColdAmount:    "500000000000000000",
			},
		},
	}

	// 调用方法
	resp, err := client.SetTokenAddress(ctx, req)
	if err != nil {
		t.Fatalf("SetTokenAddress failed: %v", err)
	}

	// 验证响应
	assert.NotNil(t, resp, "Response should not be nil")
	respJSON, err := json.Marshal(resp)
	assert.NoError(t, err, "Failed to marshal response to JSON")
	t.Logf("Response: %s", respJSON)
}

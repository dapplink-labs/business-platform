package account

import (
	"business-platform/app/common/grpcclient/account/proto"
	"context"
	"encoding/json"
	"testing"
)

func Test_GrpcClient_GetSupportChains(t *testing.T) {
	// 创建 gRPC 客户端
	endpoint := "127.0.0.1:8189" // 替换为你的实际 gRPC 服务器地址
	client, err := NewGrpcClient(endpoint)
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
	req := &proto.SupportChainsRequest{
		Chain: "Ethereum",
	} // 根据实际需求设置请求参数

	// 调用 GetSupportChains 方法
	resp, err := client.GetSupportChains(ctx, req)
	if err != nil {
		t.Fatalf("调用 GetSupportChains 失败: %v", err)
	}

	// 验证响应
	if resp == nil {
		t.Error("获取到的响应为空")
		return
	}

	respJson, _ := json.Marshal(resp)
	t.Logf("响应状态码: %s", respJson)
}

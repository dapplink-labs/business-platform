package sign

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"business-platform/app/common/grpcclient/sign/proto"
)

const (
	signUrl = "127.0.0.1:8983"
)

func Test_sign_ExportPublicKeyList_Integration(t *testing.T) {
	// 创建真实的 gRPC 客户端
	client, err := NewGrpcClient(signUrl)
	if err != nil {
		t.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 准备请求参数
	req := &proto.ExportPublicKeyRequest{
		Type:   "ecdsa",
		Number: 0,
	}

	// 调用方法
	resp, err := client.ExportPublicKeyList(ctx, req)
	if err != nil {
		t.Fatalf("ExportPublicKeyList failed: %v", err)
	}

	// 验证响应
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Code)
	assert.NotEmpty(t, resp.Msg)

	// 记录响应
	respJSON, err := json.Marshal(resp)
	assert.NoError(t, err)
	t.Logf("Integration test response: %s", respJSON)
}

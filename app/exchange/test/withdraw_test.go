package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"business-platform/app/exchange/internal/types"
)

func TestWithdraw(t *testing.T) {
	// 1. 准备请求数据
	reqData := map[string]interface{}{
		"amount":       "1000000000000000000", // 1 ETH
		"to_address":   "0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
		"chain_id":     1,
		"token_symbol": "ETH",
		"token_addr":   "",
		"uid":          12345,
		"memo":         "test withdraw",
	}

	// 2. 转换为JSON
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		t.Fatal("JSON序列化失败:", err)
	}

	// 3. 发送请求
	resp, err := http.Post(
		"http://localhost:8888/api/v1/withdraw/create",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		t.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	// 4. 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("读取响应失败:", err)
	}

	// 5. 打印响应
	fmt.Printf("1 响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("1 响应内容: %s\n", string(body))
	bodyJson, _ := json.Marshal(body)
	fmt.Printf("1 body json: %s\n", string(bodyJson))

	// 定义响应结构
	type Response struct {
		Code int                    `json:"code"`
		Msg  string                 `json:"msg"`
		Data types.WithdrawResponse `json:"data"`
	}
	// 解析响应
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal("2 解析响应失败:", err)
	}
	// 5. 打印响应
	fmt.Printf("2 响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("2 响应内容: %+v\n", result)

	// 6. 验证响应
	if result.Code != 200 {
		t.Errorf("预期状态码为200，实际获得：%d", result.Code)
	}

	fmt.Printf("3 订单ID: %d\n", result.Data.OrderId)
	fmt.Printf("3 交易哈希: %s\n", result.Data.TxHash)
}

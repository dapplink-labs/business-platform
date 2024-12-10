package sign

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"business-platform/app/common/grpcclient"
	"business-platform/app/common/grpcclient/sign/proto"
)

type GrpcClient interface {
	GetSupportSignWay(ctx context.Context, cryptoType CryptoType) error
	ExportPublicKeyList(ctx context.Context, cryptoType CryptoType, number uint64) (publicKeyList []*proto.PublicKey, err error)
	SignTxMessage(ctx context.Context, cryptoType CryptoType, pubKey, messageHash string) (signature string, err error)
	Close() error
}

type grpcClient struct {
	conn     *grpc.ClientConn
	client   proto.WalletServiceClient
	endpoint string
}

func NewGrpcClient(endpoint string) (GrpcClient, error) {
	conn, err := grpcclient.CreateGrpcConnection(endpoint)
	if err != nil {
		return nil, err
	}

	client := proto.NewWalletServiceClient(conn)

	return &grpcClient{
		conn:     conn,
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (c *grpcClient) GetSupportSignWay(ctx context.Context, cryptoType CryptoType) error {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	if cryptoType == "" {
		return fmt.Errorf("cryptoType cannot be empty")
	}
	in := &proto.SupportSignWayRequest{
		Type: string(cryptoType),
	}

	// Call the gRPC method
	resp, err := c.client.GetSupportSignWay(ctx, in)
	if err != nil {
		return fmt.Errorf("failed to get support sign way: %w", err)
	}
	// Check if the response code indicates success
	if proto.ReturnCode_SUCCESS != resp.Code {
		return fmt.Errorf("failed to export addresses msg: %s", resp.Msg)
	}
	// Return the support status
	return nil
}

func (c *grpcClient) ExportPublicKeyList(ctx context.Context, cryptoType CryptoType, number uint64) (publicKeyList []*proto.PublicKey, err error) {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	if cryptoType == "" {
		return nil, fmt.Errorf("cryptoType cannot be empty")
	}
	if number == 0 {
		return nil, fmt.Errorf("number must be greater than zero")
	}
	// Create the request object
	req := &proto.ExportPublicKeyRequest{
		Type:   string(cryptoType), // Assuming CryptoType can be converted to string
		Number: number,
	}
	// Call the gRPC method
	resp, err := c.client.ExportPublicKeyList(ctx, req)
	if err != nil {
		return nil, err
	}
	// Check if the response code indicates success
	if proto.ReturnCode_SUCCESS != resp.Code {
		return nil, fmt.Errorf("failed to export addresses msg: %s", resp.Msg)
	}
	// Return the public key list from the response
	return resp.PublicKey, nil
}

func (c *grpcClient) SignTxMessage(ctx context.Context, cryptoType CryptoType, pubKey, messageHash string) (signature string, err error) {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	if cryptoType == "" {
		return "", fmt.Errorf("cryptoType cannot be empty")
	}
	if pubKey == "" {
		return "", fmt.Errorf("public key cannot be empty")
	}
	if messageHash == "" {
		return "", fmt.Errorf("message hash cannot be empty")
	}
	// 创建请求对象
	req := &proto.SignTxMessageRequest{
		Type:        string(cryptoType), // Assuming CryptoType can be converted to string
		PublicKey:   pubKey,
		MessageHash: messageHash,
	}
	// 调用 gRPC 方法
	resp, err := c.client.SignTxMessage(ctx, req)
	if err != nil {
		return "", err
	}
	// Check if the response code indicates success
	if proto.ReturnCode_SUCCESS != resp.Code {
		return "", fmt.Errorf("failed to export addresses msg: %s", resp.Msg)
	}
	// 返回签名
	return resp.Signature, nil
}

func (c *grpcClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithTimeout(ctx, timeout)
}

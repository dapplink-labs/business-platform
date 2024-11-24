package sign

import (
	"context"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"business-platform/app/common/grpcclient/sign/proto"
)

type GrpcClient interface {
	GetSupportSignWay(ctx context.Context, in *proto.SupportSignWayRequest) (*proto.SupportSignWayResponse, error)
	ExportPublicKeyList(ctx context.Context, in *proto.ExportPublicKeyRequest) (*proto.ExportPublicKeyResponse, error)
	SignTxMessage(ctx context.Context, in *proto.SignTxMessageRequest) (*proto.SignTxMessageResponse, error)
	Close() error
}

type grpcClient struct {
	conn     *grpc.ClientConn
	client   proto.WalletServiceClient
	endpoint string
}

const (
	dialTimeout      = 5 * time.Second
	keepaliveTime    = 30 * time.Second
	keepaliveTimeout = 10 * time.Second
)

func NewGrpcClient(endpoint string) (GrpcClient, error) {
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "https://")

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff:           backoff.DefaultConfig,
			MinConnectTimeout: dialTimeout,
		}),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                keepaliveTime,
			Timeout:             keepaliveTimeout,
			PermitWithoutStream: true,
		}),
	}

	target := fmt.Sprintf("dns:///%s", endpoint)
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	conn.Connect()

	state := conn.GetState()
	for state != connectivity.Ready {
		if !conn.WaitForStateChange(ctx, state) {
			err := conn.Close()
			if err != nil {
				return nil, fmt.Errorf("grpc connection conn.Close: %w", err)
			}
			return nil, fmt.Errorf("grpc connection timeout")
		}
		state = conn.GetState()
	}

	client := proto.NewWalletServiceClient(conn)

	return &grpcClient{
		conn:     conn,
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (c *grpcClient) GetSupportSignWay(ctx context.Context, in *proto.SupportSignWayRequest) (*proto.SupportSignWayResponse, error) {
	return c.client.GetSupportSignWay(ctx, in)
}

func (c *grpcClient) ExportPublicKeyList(ctx context.Context, in *proto.ExportPublicKeyRequest) (*proto.ExportPublicKeyResponse, error) {
	return c.client.ExportPublicKeyList(ctx, in)
}

func (c *grpcClient) SignTxMessage(ctx context.Context, in *proto.SignTxMessageRequest) (*proto.SignTxMessageResponse, error) {
	return c.client.SignTxMessage(ctx, in)
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

package multichain

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

	"business-platform/app/common/grpcclient/multichain/proto"
)

type GrpcClient interface {
	BusinessRegister(ctx context.Context, in *proto.BusinessRegisterRequest) (*proto.BusinessRegisterResponse, error)
	ExportAddressesByPublicKeys(ctx context.Context, in *proto.ExportAddressesRequest) (*proto.ExportAddressesResponse, error)
	CreateUnSignTransaction(ctx context.Context, in *proto.UnSignWithdrawTransactionRequest) (*proto.UnSignWithdrawTransactionResponse, error)
	BuildSignedTransaction(ctx context.Context, in *proto.SignedWithdrawTransactionRequest) (*proto.SignedWithdrawTransactionResponse, error)
	// SetTokenAddress add token
	SetTokenAddress(ctx context.Context, in *proto.SetTokenAddressRequest) (*proto.SetTokenAddressResponse, error)
	Close() error
}

type grpcClient struct {
	conn     *grpc.ClientConn
	client   proto.BusinessMiddleWireServicesClient
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

	client := proto.NewBusinessMiddleWireServicesClient(conn)

	return &grpcClient{
		conn:     conn,
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (c *grpcClient) BusinessRegister(ctx context.Context, in *proto.BusinessRegisterRequest) (*proto.BusinessRegisterResponse, error) {
	return c.client.BusinessRegister(ctx, in)
}

func (c *grpcClient) ExportAddressesByPublicKeys(ctx context.Context, in *proto.ExportAddressesRequest) (*proto.ExportAddressesResponse, error) {
	return c.client.ExportAddressesByPublicKeys(ctx, in)
}

func (c *grpcClient) CreateUnSignTransaction(ctx context.Context, in *proto.UnSignWithdrawTransactionRequest) (*proto.UnSignWithdrawTransactionResponse, error) {
	return c.client.CreateUnSignTransaction(ctx, in)
}

func (c *grpcClient) BuildSignedTransaction(ctx context.Context, in *proto.SignedWithdrawTransactionRequest) (*proto.SignedWithdrawTransactionResponse, error) {
	return c.client.BuildSignedTransaction(ctx, in)
}

func (c *grpcClient) SetTokenAddress(ctx context.Context, in *proto.SetTokenAddressRequest) (*proto.SetTokenAddressResponse, error) {
	return c.client.SetTokenAddress(ctx, in)
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

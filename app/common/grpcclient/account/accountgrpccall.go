package account

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

	"business-platform/app/common/grpcclient/account/proto"
)

type GrpcClient interface {
	GetSupportChains(ctx context.Context, in *proto.SupportChainsRequest) (*proto.SupportChainsResponse, error)
	ConvertAddress(ctx context.Context, in *proto.ConvertAddressRequest) (*proto.ConvertAddressResponse, error)
	ValidAddress(ctx context.Context, in *proto.ValidAddressRequest) (*proto.ValidAddressResponse, error)
	GetBlockByNumber(ctx context.Context, in *proto.BlockNumberRequest) (*proto.BlockResponse, error)
	GetBlockByHash(ctx context.Context, in *proto.BlockHashRequest) (*proto.BlockResponse, error)
	GetBlockHeaderByHash(ctx context.Context, in *proto.BlockHeaderHashRequest) (*proto.BlockHeaderResponse, error)
	GetBlockHeaderByNumber(ctx context.Context, in *proto.BlockHeaderNumberRequest) (*proto.BlockHeaderResponse, error)
	GetBlockHeaderByRange(ctx context.Context, in *proto.BlockByRangeRequest) (*proto.BlockByRangeResponse, error)
	GetAccount(ctx context.Context, in *proto.AccountRequest) (*proto.AccountResponse, error)
	GetFee(ctx context.Context, in *proto.FeeRequest) (*proto.FeeResponse, error)
	SendTx(ctx context.Context, in *proto.SendTxRequest) (*proto.SendTxResponse, error)
	GetTxByAddress(ctx context.Context, in *proto.TxAddressRequest) (*proto.TxAddressResponse, error)
	GetTxByHash(ctx context.Context, in *proto.TxHashRequest) (*proto.TxHashResponse, error)
	CreateUnSignTransaction(ctx context.Context, in *proto.UnSignTransactionRequest) (*proto.UnSignTransactionResponse, error)
	BuildSignedTransaction(ctx context.Context, in *proto.SignedTransactionRequest) (*proto.SignedTransactionResponse, error)
	DecodeTransaction(ctx context.Context, in *proto.DecodeTransactionRequest) (*proto.DecodeTransactionResponse, error)
	VerifySignedTransaction(ctx context.Context, in *proto.VerifyTransactionRequest) (*proto.VerifyTransactionResponse, error)
	GetExtraData(ctx context.Context, in *proto.ExtraDataRequest) (*proto.ExtraDataResponse, error)
	Close() error
}

type grpcClient struct {
	conn     *grpc.ClientConn
	client   proto.WalletAccountServiceClient
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

	client := proto.NewWalletAccountServiceClient(conn)

	return &grpcClient{
		conn:     conn,
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (c *grpcClient) GetSupportChains(ctx context.Context, in *proto.SupportChainsRequest) (*proto.SupportChainsResponse, error) {
	return c.client.GetSupportChains(ctx, in)
}

func (c *grpcClient) ConvertAddress(ctx context.Context, in *proto.ConvertAddressRequest) (*proto.ConvertAddressResponse, error) {
	return c.client.ConvertAddress(ctx, in)
}

func (c *grpcClient) ValidAddress(ctx context.Context, in *proto.ValidAddressRequest) (*proto.ValidAddressResponse, error) {
	return c.client.ValidAddress(ctx, in)
}

func (c *grpcClient) GetBlockByNumber(ctx context.Context, in *proto.BlockNumberRequest) (*proto.BlockResponse, error) {
	return c.client.GetBlockByNumber(ctx, in)
}

func (c *grpcClient) GetBlockByHash(ctx context.Context, in *proto.BlockHashRequest) (*proto.BlockResponse, error) {
	return c.client.GetBlockByHash(ctx, in)
}

func (c *grpcClient) GetBlockHeaderByHash(ctx context.Context, in *proto.BlockHeaderHashRequest) (*proto.BlockHeaderResponse, error) {
	return c.client.GetBlockHeaderByHash(ctx, in)
}

func (c *grpcClient) GetBlockHeaderByNumber(ctx context.Context, in *proto.BlockHeaderNumberRequest) (*proto.BlockHeaderResponse, error) {
	return c.client.GetBlockHeaderByNumber(ctx, in)
}

func (c *grpcClient) GetBlockHeaderByRange(ctx context.Context, in *proto.BlockByRangeRequest) (*proto.BlockByRangeResponse, error) {
	return c.client.GetBlockHeaderByRange(ctx, in)
}

func (c *grpcClient) GetAccount(ctx context.Context, in *proto.AccountRequest) (*proto.AccountResponse, error) {
	return c.client.GetAccount(ctx, in)
}

func (c *grpcClient) GetFee(ctx context.Context, in *proto.FeeRequest) (*proto.FeeResponse, error) {
	return c.client.GetFee(ctx, in)
}

func (c *grpcClient) SendTx(ctx context.Context, in *proto.SendTxRequest) (*proto.SendTxResponse, error) {
	return c.client.SendTx(ctx, in)
}

func (c *grpcClient) GetTxByAddress(ctx context.Context, in *proto.TxAddressRequest) (*proto.TxAddressResponse, error) {
	return c.client.GetTxByAddress(ctx, in)
}

func (c *grpcClient) GetTxByHash(ctx context.Context, in *proto.TxHashRequest) (*proto.TxHashResponse, error) {
	return c.client.GetTxByHash(ctx, in)
}

func (c *grpcClient) CreateUnSignTransaction(ctx context.Context, in *proto.UnSignTransactionRequest) (*proto.UnSignTransactionResponse, error) {
	return c.client.CreateUnSignTransaction(ctx, in)
}

func (c *grpcClient) BuildSignedTransaction(ctx context.Context, in *proto.SignedTransactionRequest) (*proto.SignedTransactionResponse, error) {
	return c.client.BuildSignedTransaction(ctx, in)
}

func (c *grpcClient) DecodeTransaction(ctx context.Context, in *proto.DecodeTransactionRequest) (*proto.DecodeTransactionResponse, error) {
	return c.client.DecodeTransaction(ctx, in)
}

func (c *grpcClient) VerifySignedTransaction(ctx context.Context, in *proto.VerifyTransactionRequest) (*proto.VerifyTransactionResponse, error) {
	return c.client.VerifySignedTransaction(ctx, in)
}

func (c *grpcClient) GetExtraData(ctx context.Context, in *proto.ExtraDataRequest) (*proto.ExtraDataResponse, error) {
	return c.client.GetExtraData(ctx, in)
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

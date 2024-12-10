package multichain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"business-platform/app/common/grpcclient"
	"business-platform/app/common/grpcclient/multichain/proto"
)

type GrpcClient interface {
	BusinessRegister(ctx context.Context, businessId, notifyUrl string) error
	ExportAddressesByPublicKeys(ctx context.Context, businessId string, publicKeyList []*proto.PublicKey) ([]*proto.Address, error)
	CreateUnSignTransaction(ctx context.Context, businessId string, in *CreateUnSignTransactionRequest) (txId string, UnSignTx string, err error)
	BuildSignedTransaction(ctx context.Context, businessId string, in *CreateSignedTransactionRequest) (SignedTx string, err error)
	// SetTokenAddress add token
	SetTokenAddress(ctx context.Context, businessId string, tokenList []*proto.Token) error
	Close() error
}

type grpcClient struct {
	conn     *grpc.ClientConn
	client   proto.BusinessMiddleWireServicesClient
	endpoint string
}

func NewGrpcClient(endpoint string) (GrpcClient, error) {
	conn, err := grpcclient.CreateGrpcConnection(endpoint)
	if err != nil {
		return nil, err
	}

	client := proto.NewBusinessMiddleWireServicesClient(conn)

	return &grpcClient{
		conn:     conn,
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (c *grpcClient) BusinessRegister(ctx context.Context, businessId, notifyUrl string) error {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	// Check if the businessId is empty
	if businessId == "" {
		return fmt.Errorf("businessId cannot be empty")
	}

	// Check if the notifyUrl is empty
	if notifyUrl == "" {
		return fmt.Errorf("notifyUrl cannot be empty")
	}

	req := &proto.BusinessRegisterRequest{
		RequestId: businessId,
		NotifyUrl: notifyUrl,
	}
	resp, err := c.client.BusinessRegister(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to register business: %w", err)
	}
	if proto.ReturnCode_SUCCESS != resp.Code {
		return fmt.Errorf("failed to register business msg: %s", resp.Msg)
	}
	return nil
}

func (c *grpcClient) ExportAddressesByPublicKeys(ctx context.Context, businessId string, publicKeyList []*proto.PublicKey) ([]*proto.Address, error) {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	// Validate input parameters
	if businessId == "" {
		return nil, fmt.Errorf("businessId cannot be empty")
	}
	if len(publicKeyList) == 0 {
		return nil, fmt.Errorf("publicKeyList cannot be empty")
	}

	// Construct the request
	req := &proto.ExportAddressesRequest{
		RequestId:  businessId, // Assuming a function to generate a unique request ID
		PublicKeys: publicKeyList,
	}

	// Make the gRPC call
	resp, err := c.client.ExportAddressesByPublicKeys(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to export addresses: %w", err)
	}

	// Check if the response code indicates success
	if proto.ReturnCode_SUCCESS != resp.Code {
		return nil, fmt.Errorf("failed to export addresses msg: %s", resp.Msg)
	}

	return resp.Addresses, nil
}

func (c *grpcClient) CreateUnSignTransaction(ctx context.Context, businessId string, in *CreateUnSignTransactionRequest) (txId string, UnSignTx string, err error) {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	// Validate input parameters
	if businessId == "" {
		return "", "", errors.New("businessId cannot be empty")
	}
	if in == nil {
		return "", "", errors.New("request cannot be nil")
	}

	// Initialize the request
	req := &proto.UnSignTransactionRequest{
		ConsumerToken: businessId, // Assuming businessId maps to ConsumerToken
		RequestId:     in.ChainId, // Assuming you have a way to generate or map a request ID
		ChainId:       in.ChainId,
		Chain:         in.Chain,
		TxType:        string(in.TxType), // Convert TransactionType to string
	}

	// Distinguish based on TokenType
	switch in.TokenType {
	case TokenTypeETH:
		req.From = in.TxETH.From
		req.To = in.TxETH.To
		req.Value = in.TxETH.Value
	case TokenTypeERC20, TokenTypeERC721, TokenTypeERC1155:
		req.From = in.TxERC.From
		req.To = in.TxERC.To
		req.Value = in.TxERC.Value
		req.ContractAddress = in.TxERC.ContractAddress
		req.TokenId = in.TxERC.TokenId
		req.TokenMeta = in.TxERC.TokenMeta
	default:
		return "", "", fmt.Errorf("unsupported token type: %s", in.TokenType)
	}

	// Call the gRPC method
	resp, err := c.client.CreateUnSignTransaction(ctx, req)
	if err != nil {
		return "", "", fmt.Errorf("failed to create unsigned transaction: %w", err)
	}
	// Process the response
	// Assuming the response contains fields TxId and UnsignedTx
	txId = resp.TransactionId
	UnSignTx = resp.UnSignTx
	return txId, UnSignTx, nil
}

func (c *grpcClient) BuildSignedTransaction(ctx context.Context, businessId string, in *CreateSignedTransactionRequest) (SignedTx string, err error) {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()
	// Validate input parameters
	if in == nil {
		return "", errors.New("request cannot be nil")
	}

	// Map CreateSignedTransactionRequest to SignedTransactionRequest
	req := &proto.SignedTransactionRequest{
		ConsumerToken: "",         // Assuming you have a way to set this, possibly from context or configuration
		RequestId:     businessId, // Assuming TransactionId is used as RequestId
		Chain:         in.Chain,
		ChainId:       in.ChainId,
		TransactionId: in.TransactionId,
		Signature:     in.Signature,
		TxType:        string(in.TxType), // Convert TransactionType to string
	}

	// Call the gRPC method
	resp, err := c.client.BuildSignedTransaction(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to build signed transaction: %w", err)
	}

	// Check the response code
	if resp.Code != proto.ReturnCode_SUCCESS {
		return "", fmt.Errorf("failed to build signed transaction: %s", resp.Msg)
	}

	// Return the signed transaction
	return resp.SignedTx, nil
}

func (c *grpcClient) SetTokenAddress(ctx context.Context, businessId string, tokenList []*proto.Token) error {
	ctx, cancelFunc := WithTimeout(ctx, grpcclient.OnceRequestTimeout)
	defer cancelFunc()

	// Validate input parameters
	if businessId == "" {
		return errors.New("businessId cannot be empty")
	}
	if len(tokenList) == 0 {
		return errors.New("tokenList cannot be empty")
	}

	// Create the SetTokenAddressRequest
	req := &proto.SetTokenAddressRequest{
		RequestId: businessId, // Assuming businessId is used as RequestId
		TokenList: tokenList,
	}

	// Call the gRPC method
	resp, err := c.client.SetTokenAddress(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to set token address: %w", err)
	}

	// Check the response code
	if resp.Code != proto.ReturnCode_SUCCESS {
		return fmt.Errorf("failed to set token address: %s", resp.Msg)
	}

	return nil
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

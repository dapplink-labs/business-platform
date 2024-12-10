package grpcclient

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"strings"
	"time"
)

const (
	// ConnectTimeout is the maximum duration allowed for establishing a new gRPC connection.
	ConnectTimeout = 5 * time.Second
	// keepaliveTime is the interval at which the client sends "ping" messages to the server
	// to ensure the connection remains active, even when there are no active streams.
	keepaliveTime = 30 * time.Second
	// keepaliveTimeout is the maximum time the client waits for a "ping" response from the server.
	// If the server does not respond within this time, the connection is considered broken.
	keepaliveTimeout = 10 * time.Second

	OnceRequestTimeout = 5 * time.Second
)

func CreateGrpcConnection(endpoint string) (*grpc.ClientConn, error) {
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "https://")

	ctx, cancel := context.WithTimeout(context.Background(), ConnectTimeout)
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff:           backoff.DefaultConfig,
			MinConnectTimeout: ConnectTimeout,
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

	return conn, nil
}

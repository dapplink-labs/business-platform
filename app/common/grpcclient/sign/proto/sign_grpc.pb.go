// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: protobuf/wallet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WalletService_GetSupportSignWay_FullMethodName   = "/dapplink.wallet.WalletService/getSupportSignWay"
	WalletService_ExportPublicKeyList_FullMethodName = "/dapplink.wallet.WalletService/exportPublicKeyList"
	WalletService_SignTxMessage_FullMethodName       = "/dapplink.wallet.WalletService/signTxMessage"
)

// WalletServiceClient is the client API for WalletService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletServiceClient interface {
	GetSupportSignWay(ctx context.Context, in *SupportSignWayRequest, opts ...grpc.CallOption) (*SupportSignWayResponse, error)
	ExportPublicKeyList(ctx context.Context, in *ExportPublicKeyRequest, opts ...grpc.CallOption) (*ExportPublicKeyResponse, error)
	SignTxMessage(ctx context.Context, in *SignTxMessageRequest, opts ...grpc.CallOption) (*SignTxMessageResponse, error)
}

type walletServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletServiceClient(cc grpc.ClientConnInterface) WalletServiceClient {
	return &walletServiceClient{cc}
}

func (c *walletServiceClient) GetSupportSignWay(ctx context.Context, in *SupportSignWayRequest, opts ...grpc.CallOption) (*SupportSignWayResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SupportSignWayResponse)
	err := c.cc.Invoke(ctx, WalletService_GetSupportSignWay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) ExportPublicKeyList(ctx context.Context, in *ExportPublicKeyRequest, opts ...grpc.CallOption) (*ExportPublicKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportPublicKeyResponse)
	err := c.cc.Invoke(ctx, WalletService_ExportPublicKeyList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletServiceClient) SignTxMessage(ctx context.Context, in *SignTxMessageRequest, opts ...grpc.CallOption) (*SignTxMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignTxMessageResponse)
	err := c.cc.Invoke(ctx, WalletService_SignTxMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServiceServer is the server API for WalletService service.
// All implementations should embed UnimplementedWalletServiceServer
// for forward compatibility.
type WalletServiceServer interface {
	GetSupportSignWay(context.Context, *SupportSignWayRequest) (*SupportSignWayResponse, error)
	ExportPublicKeyList(context.Context, *ExportPublicKeyRequest) (*ExportPublicKeyResponse, error)
	SignTxMessage(context.Context, *SignTxMessageRequest) (*SignTxMessageResponse, error)
}

// UnimplementedWalletServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWalletServiceServer struct{}

func (UnimplementedWalletServiceServer) GetSupportSignWay(context.Context, *SupportSignWayRequest) (*SupportSignWayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportSignWay not implemented")
}
func (UnimplementedWalletServiceServer) ExportPublicKeyList(context.Context, *ExportPublicKeyRequest) (*ExportPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportPublicKeyList not implemented")
}
func (UnimplementedWalletServiceServer) SignTxMessage(context.Context, *SignTxMessageRequest) (*SignTxMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTxMessage not implemented")
}
func (UnimplementedWalletServiceServer) testEmbeddedByValue() {}

// UnsafeWalletServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletServiceServer will
// result in compilation errors.
type UnsafeWalletServiceServer interface {
	mustEmbedUnimplementedWalletServiceServer()
}

func RegisterWalletServiceServer(s grpc.ServiceRegistrar, srv WalletServiceServer) {
	// If the following call pancis, it indicates UnimplementedWalletServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WalletService_ServiceDesc, srv)
}

func _WalletService_GetSupportSignWay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportSignWayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).GetSupportSignWay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_GetSupportSignWay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).GetSupportSignWay(ctx, req.(*SupportSignWayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_ExportPublicKeyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).ExportPublicKeyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_ExportPublicKeyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).ExportPublicKeyList(ctx, req.(*ExportPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletService_SignTxMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTxMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServiceServer).SignTxMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletService_SignTxMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServiceServer).SignTxMessage(ctx, req.(*SignTxMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletService_ServiceDesc is the grpc.ServiceDesc for WalletService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapplink.wallet.WalletService",
	HandlerType: (*WalletServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSupportSignWay",
			Handler:    _WalletService_GetSupportSignWay_Handler,
		},
		{
			MethodName: "exportPublicKeyList",
			Handler:    _WalletService_ExportPublicKeyList_Handler,
		},
		{
			MethodName: "signTxMessage",
			Handler:    _WalletService_SignTxMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/wallet.proto",
}

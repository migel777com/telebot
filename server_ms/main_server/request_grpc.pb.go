// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TelegramBotServiceClient is the client API for TelegramBotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelegramBotServiceClient interface {
	//Unary
	CommandPack(ctx context.Context, in *CommandPackRequest, opts ...grpc.CallOption) (*CommandPackResponse, error)
}

type telegramBotServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegramBotServiceClient(cc grpc.ClientConnInterface) TelegramBotServiceClient {
	return &telegramBotServiceClient{cc}
}

func (c *telegramBotServiceClient) CommandPack(ctx context.Context, in *CommandPackRequest, opts ...grpc.CallOption) (*CommandPackResponse, error) {
	out := new(CommandPackResponse)
	err := c.cc.Invoke(ctx, "/commandpack.TelegramBotService/CommandPack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramBotServiceServer is the server API for TelegramBotService service.
// All implementations must embed UnimplementedTelegramBotServiceServer
// for forward compatibility
type TelegramBotServiceServer interface {
	//Unary
	CommandPack(context.Context, *CommandPackRequest) (*CommandPackResponse, error)
	mustEmbedUnimplementedTelegramBotServiceServer()
}

// UnimplementedTelegramBotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTelegramBotServiceServer struct {
}

func (UnimplementedTelegramBotServiceServer) CommandPack(context.Context, *CommandPackRequest) (*CommandPackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandPack not implemented")
}
func (UnimplementedTelegramBotServiceServer) mustEmbedUnimplementedTelegramBotServiceServer() {}

// UnsafeTelegramBotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegramBotServiceServer will
// result in compilation errors.
type UnsafeTelegramBotServiceServer interface {
	mustEmbedUnimplementedTelegramBotServiceServer()
}

func RegisterTelegramBotServiceServer(s grpc.ServiceRegistrar, srv TelegramBotServiceServer) {
	s.RegisterService(&TelegramBotService_ServiceDesc, srv)
}

func _TelegramBotService_CommandPack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandPackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramBotServiceServer).CommandPack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commandpack.TelegramBotService/CommandPack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramBotServiceServer).CommandPack(ctx, req.(*CommandPackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelegramBotService_ServiceDesc is the grpc.ServiceDesc for TelegramBotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelegramBotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commandpack.TelegramBotService",
	HandlerType: (*TelegramBotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommandPack",
			Handler:    _TelegramBotService_CommandPack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main_server/main_server.proto",
}

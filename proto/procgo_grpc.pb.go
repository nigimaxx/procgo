// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProcgoClient is the client API for Procgo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcgoClient interface {
	Start(ctx context.Context, in *Services, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Stop(ctx context.Context, in *Services, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Restart(ctx context.Context, in *Services, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Logs(ctx context.Context, in *AllOrServices, opts ...grpc.CallOption) (Procgo_LogsClient, error)
	KillAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type procgoClient struct {
	cc grpc.ClientConnInterface
}

func NewProcgoClient(cc grpc.ClientConnInterface) ProcgoClient {
	return &procgoClient{cc}
}

func (c *procgoClient) Start(ctx context.Context, in *Services, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Procgo/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procgoClient) Stop(ctx context.Context, in *Services, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Procgo/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procgoClient) Restart(ctx context.Context, in *Services, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Procgo/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procgoClient) Logs(ctx context.Context, in *AllOrServices, opts ...grpc.CallOption) (Procgo_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Procgo_serviceDesc.Streams[0], "/Procgo/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &procgoLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Procgo_LogsClient interface {
	Recv() (*wrapperspb.BytesValue, error)
	grpc.ClientStream
}

type procgoLogsClient struct {
	grpc.ClientStream
}

func (x *procgoLogsClient) Recv() (*wrapperspb.BytesValue, error) {
	m := new(wrapperspb.BytesValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *procgoClient) KillAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Procgo/KillAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procgoClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Procgo/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcgoServer is the server API for Procgo service.
// All implementations must embed UnimplementedProcgoServer
// for forward compatibility
type ProcgoServer interface {
	Start(context.Context, *Services) (*emptypb.Empty, error)
	Stop(context.Context, *Services) (*emptypb.Empty, error)
	Restart(context.Context, *Services) (*emptypb.Empty, error)
	Logs(*AllOrServices, Procgo_LogsServer) error
	KillAll(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedProcgoServer()
}

// UnimplementedProcgoServer must be embedded to have forward compatible implementations.
type UnimplementedProcgoServer struct {
}

func (UnimplementedProcgoServer) Start(context.Context, *Services) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedProcgoServer) Stop(context.Context, *Services) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedProcgoServer) Restart(context.Context, *Services) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedProcgoServer) Logs(*AllOrServices, Procgo_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedProcgoServer) KillAll(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KillAll not implemented")
}
func (UnimplementedProcgoServer) Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProcgoServer) mustEmbedUnimplementedProcgoServer() {}

// UnsafeProcgoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcgoServer will
// result in compilation errors.
type UnsafeProcgoServer interface {
	mustEmbedUnimplementedProcgoServer()
}

func RegisterProcgoServer(s grpc.ServiceRegistrar, srv ProcgoServer) {
	s.RegisterService(&_Procgo_serviceDesc, srv)
}

func _Procgo_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Services)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcgoServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Procgo/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcgoServer).Start(ctx, req.(*Services))
	}
	return interceptor(ctx, in, info, handler)
}

func _Procgo_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Services)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcgoServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Procgo/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcgoServer).Stop(ctx, req.(*Services))
	}
	return interceptor(ctx, in, info, handler)
}

func _Procgo_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Services)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcgoServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Procgo/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcgoServer).Restart(ctx, req.(*Services))
	}
	return interceptor(ctx, in, info, handler)
}

func _Procgo_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AllOrServices)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcgoServer).Logs(m, &procgoLogsServer{stream})
}

type Procgo_LogsServer interface {
	Send(*wrapperspb.BytesValue) error
	grpc.ServerStream
}

type procgoLogsServer struct {
	grpc.ServerStream
}

func (x *procgoLogsServer) Send(m *wrapperspb.BytesValue) error {
	return x.ServerStream.SendMsg(m)
}

func _Procgo_KillAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcgoServer).KillAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Procgo/KillAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcgoServer).KillAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Procgo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcgoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Procgo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcgoServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Procgo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Procgo",
	HandlerType: (*ProcgoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Procgo_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Procgo_Stop_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Procgo_Restart_Handler,
		},
		{
			MethodName: "KillAll",
			Handler:    _Procgo_KillAll_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Procgo_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _Procgo_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/procgo.proto",
}

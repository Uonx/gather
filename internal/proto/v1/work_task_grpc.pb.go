// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc1
// source: internal/proto/v1/work_task.proto

package v1

import (
	context "context"
	basic "github.com/Uonx/gather/internal/proto/basic"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SchedulerEvent_Registry_FullMethodName     = "/v1.SchedulerEvent/Registry"
	SchedulerEvent_HealthResult_FullMethodName = "/v1.SchedulerEvent/HealthResult"
	SchedulerEvent_WorkResult_FullMethodName   = "/v1.SchedulerEvent/WorkResult"
)

// SchedulerEventClient is the client API for SchedulerEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerEventClient interface {
	Registry(ctx context.Context, in *RegisterEvent, opts ...grpc.CallOption) (SchedulerEvent_RegistryClient, error)
	HealthResult(ctx context.Context, in *Health, opts ...grpc.CallOption) (SchedulerEvent_HealthResultClient, error)
	WorkResult(ctx context.Context, in *Result, opts ...grpc.CallOption) (*basic.Response, error)
}

type schedulerEventClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerEventClient(cc grpc.ClientConnInterface) SchedulerEventClient {
	return &schedulerEventClient{cc}
}

func (c *schedulerEventClient) Registry(ctx context.Context, in *RegisterEvent, opts ...grpc.CallOption) (SchedulerEvent_RegistryClient, error) {
	stream, err := c.cc.NewStream(ctx, &SchedulerEvent_ServiceDesc.Streams[0], SchedulerEvent_Registry_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerEventRegistryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SchedulerEvent_RegistryClient interface {
	Recv() (*Execution, error)
	grpc.ClientStream
}

type schedulerEventRegistryClient struct {
	grpc.ClientStream
}

func (x *schedulerEventRegistryClient) Recv() (*Execution, error) {
	m := new(Execution)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerEventClient) HealthResult(ctx context.Context, in *Health, opts ...grpc.CallOption) (SchedulerEvent_HealthResultClient, error) {
	stream, err := c.cc.NewStream(ctx, &SchedulerEvent_ServiceDesc.Streams[1], SchedulerEvent_HealthResult_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &schedulerEventHealthResultClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SchedulerEvent_HealthResultClient interface {
	Recv() (*basic.Response, error)
	grpc.ClientStream
}

type schedulerEventHealthResultClient struct {
	grpc.ClientStream
}

func (x *schedulerEventHealthResultClient) Recv() (*basic.Response, error) {
	m := new(basic.Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *schedulerEventClient) WorkResult(ctx context.Context, in *Result, opts ...grpc.CallOption) (*basic.Response, error) {
	out := new(basic.Response)
	err := c.cc.Invoke(ctx, SchedulerEvent_WorkResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerEventServer is the server API for SchedulerEvent service.
// All implementations must embed UnimplementedSchedulerEventServer
// for forward compatibility
type SchedulerEventServer interface {
	Registry(*RegisterEvent, SchedulerEvent_RegistryServer) error
	HealthResult(*Health, SchedulerEvent_HealthResultServer) error
	WorkResult(context.Context, *Result) (*basic.Response, error)
	mustEmbedUnimplementedSchedulerEventServer()
}

// UnimplementedSchedulerEventServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerEventServer struct {
}

func (UnimplementedSchedulerEventServer) Registry(*RegisterEvent, SchedulerEvent_RegistryServer) error {
	return status.Errorf(codes.Unimplemented, "method Registry not implemented")
}
func (UnimplementedSchedulerEventServer) HealthResult(*Health, SchedulerEvent_HealthResultServer) error {
	return status.Errorf(codes.Unimplemented, "method HealthResult not implemented")
}
func (UnimplementedSchedulerEventServer) WorkResult(context.Context, *Result) (*basic.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WorkResult not implemented")
}
func (UnimplementedSchedulerEventServer) mustEmbedUnimplementedSchedulerEventServer() {}

// UnsafeSchedulerEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerEventServer will
// result in compilation errors.
type UnsafeSchedulerEventServer interface {
	mustEmbedUnimplementedSchedulerEventServer()
}

func RegisterSchedulerEventServer(s grpc.ServiceRegistrar, srv SchedulerEventServer) {
	s.RegisterService(&SchedulerEvent_ServiceDesc, srv)
}

func _SchedulerEvent_Registry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RegisterEvent)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerEventServer).Registry(m, &schedulerEventRegistryServer{stream})
}

type SchedulerEvent_RegistryServer interface {
	Send(*Execution) error
	grpc.ServerStream
}

type schedulerEventRegistryServer struct {
	grpc.ServerStream
}

func (x *schedulerEventRegistryServer) Send(m *Execution) error {
	return x.ServerStream.SendMsg(m)
}

func _SchedulerEvent_HealthResult_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Health)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchedulerEventServer).HealthResult(m, &schedulerEventHealthResultServer{stream})
}

type SchedulerEvent_HealthResultServer interface {
	Send(*basic.Response) error
	grpc.ServerStream
}

type schedulerEventHealthResultServer struct {
	grpc.ServerStream
}

func (x *schedulerEventHealthResultServer) Send(m *basic.Response) error {
	return x.ServerStream.SendMsg(m)
}

func _SchedulerEvent_WorkResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Result)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerEventServer).WorkResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchedulerEvent_WorkResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerEventServer).WorkResult(ctx, req.(*Result))
	}
	return interceptor(ctx, in, info, handler)
}

// SchedulerEvent_ServiceDesc is the grpc.ServiceDesc for SchedulerEvent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchedulerEvent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.SchedulerEvent",
	HandlerType: (*SchedulerEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WorkResult",
			Handler:    _SchedulerEvent_WorkResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Registry",
			Handler:       _SchedulerEvent_Registry_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HealthResult",
			Handler:       _SchedulerEvent_HealthResult_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/proto/v1/work_task.proto",
}

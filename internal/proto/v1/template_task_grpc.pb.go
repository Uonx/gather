// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc1
// source: internal/proto/v1/template_task.proto

package v1

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

const (
	TemplateTaskEvent_TemplateResult_FullMethodName = "/v1.TemplateTaskEvent/TemplateResult"
)

// TemplateTaskEventClient is the client API for TemplateTaskEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateTaskEventClient interface {
	TemplateResult(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Result, error)
}

type templateTaskEventClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateTaskEventClient(cc grpc.ClientConnInterface) TemplateTaskEventClient {
	return &templateTaskEventClient{cc}
}

func (c *templateTaskEventClient) TemplateResult(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, TemplateTaskEvent_TemplateResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateTaskEventServer is the server API for TemplateTaskEvent service.
// All implementations must embed UnimplementedTemplateTaskEventServer
// for forward compatibility
type TemplateTaskEventServer interface {
	TemplateResult(context.Context, *Task) (*Result, error)
	mustEmbedUnimplementedTemplateTaskEventServer()
}

// UnimplementedTemplateTaskEventServer must be embedded to have forward compatible implementations.
type UnimplementedTemplateTaskEventServer struct {
}

func (UnimplementedTemplateTaskEventServer) TemplateResult(context.Context, *Task) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TemplateResult not implemented")
}
func (UnimplementedTemplateTaskEventServer) mustEmbedUnimplementedTemplateTaskEventServer() {}

// UnsafeTemplateTaskEventServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateTaskEventServer will
// result in compilation errors.
type UnsafeTemplateTaskEventServer interface {
	mustEmbedUnimplementedTemplateTaskEventServer()
}

func RegisterTemplateTaskEventServer(s grpc.ServiceRegistrar, srv TemplateTaskEventServer) {
	s.RegisterService(&TemplateTaskEvent_ServiceDesc, srv)
}

func _TemplateTaskEvent_TemplateResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateTaskEventServer).TemplateResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TemplateTaskEvent_TemplateResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateTaskEventServer).TemplateResult(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

// TemplateTaskEvent_ServiceDesc is the grpc.ServiceDesc for TemplateTaskEvent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemplateTaskEvent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TemplateTaskEvent",
	HandlerType: (*TemplateTaskEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TemplateResult",
			Handler:    _TemplateTaskEvent_TemplateResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/v1/template_task.proto",
}

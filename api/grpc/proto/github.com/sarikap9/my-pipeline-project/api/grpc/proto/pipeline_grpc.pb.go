// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pipeline.proto

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
	PipelineService_CreatePipeline_FullMethodName = "/pipeline.PipelineService/CreatePipeline"
	PipelineService_ListPipelines_FullMethodName  = "/pipeline.PipelineService/ListPipelines"
)

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineServiceClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error)
	ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error)
}

type pipelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineServiceClient(cc grpc.ClientConnInterface) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*CreatePipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_CreatePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) ListPipelines(ctx context.Context, in *ListPipelinesRequest, opts ...grpc.CallOption) (*ListPipelinesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPipelinesResponse)
	err := c.cc.Invoke(ctx, PipelineService_ListPipelines_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServiceServer is the server API for PipelineService service.
// All implementations must embed UnimplementedPipelineServiceServer
// for forward compatibility.
type PipelineServiceServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error)
	ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error)
	mustEmbedUnimplementedPipelineServiceServer()
}

// UnimplementedPipelineServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPipelineServiceServer struct{}

func (UnimplementedPipelineServiceServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*CreatePipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) ListPipelines(context.Context, *ListPipelinesRequest) (*ListPipelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelines not implemented")
}
func (UnimplementedPipelineServiceServer) mustEmbedUnimplementedPipelineServiceServer() {}
func (UnimplementedPipelineServiceServer) testEmbeddedByValue()                         {}

// UnsafePipelineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineServiceServer will
// result in compilation errors.
type UnsafePipelineServiceServer interface {
	mustEmbedUnimplementedPipelineServiceServer()
}

func RegisterPipelineServiceServer(s grpc.ServiceRegistrar, srv PipelineServiceServer) {
	// If the following call pancis, it indicates UnimplementedPipelineServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PipelineService_ServiceDesc, srv)
}

func _PipelineService_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_CreatePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_ListPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).ListPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_ListPipelines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).ListPipelines(ctx, req.(*ListPipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineService_ServiceDesc is the grpc.ServiceDesc for PipelineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipeline.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineService_CreatePipeline_Handler,
		},
		{
			MethodName: "ListPipelines",
			Handler:    _PipelineService_ListPipelines_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipeline.proto",
}

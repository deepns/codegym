// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: vector_service.proto

package proto

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
	VectorService_InsertVector_FullMethodName = "/vector_service.VectorService/InsertVector"
	VectorService_SearchVector_FullMethodName = "/vector_service.VectorService/SearchVector"
)

// VectorServiceClient is the client API for VectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VectorServiceClient interface {
	InsertVector(ctx context.Context, in *InsertVectorRequest, opts ...grpc.CallOption) (*InsertVectorResponse, error)
	SearchVector(ctx context.Context, in *SearchVectorRequest, opts ...grpc.CallOption) (*SearchVectorResponse, error)
}

type vectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVectorServiceClient(cc grpc.ClientConnInterface) VectorServiceClient {
	return &vectorServiceClient{cc}
}

func (c *vectorServiceClient) InsertVector(ctx context.Context, in *InsertVectorRequest, opts ...grpc.CallOption) (*InsertVectorResponse, error) {
	out := new(InsertVectorResponse)
	err := c.cc.Invoke(ctx, VectorService_InsertVector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vectorServiceClient) SearchVector(ctx context.Context, in *SearchVectorRequest, opts ...grpc.CallOption) (*SearchVectorResponse, error) {
	out := new(SearchVectorResponse)
	err := c.cc.Invoke(ctx, VectorService_SearchVector_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VectorServiceServer is the server API for VectorService service.
// All implementations must embed UnimplementedVectorServiceServer
// for forward compatibility
type VectorServiceServer interface {
	InsertVector(context.Context, *InsertVectorRequest) (*InsertVectorResponse, error)
	SearchVector(context.Context, *SearchVectorRequest) (*SearchVectorResponse, error)
	mustEmbedUnimplementedVectorServiceServer()
}

// UnimplementedVectorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVectorServiceServer struct {
}

func (UnimplementedVectorServiceServer) InsertVector(context.Context, *InsertVectorRequest) (*InsertVectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertVector not implemented")
}
func (UnimplementedVectorServiceServer) SearchVector(context.Context, *SearchVectorRequest) (*SearchVectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchVector not implemented")
}
func (UnimplementedVectorServiceServer) mustEmbedUnimplementedVectorServiceServer() {}

// UnsafeVectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VectorServiceServer will
// result in compilation errors.
type UnsafeVectorServiceServer interface {
	mustEmbedUnimplementedVectorServiceServer()
}

func RegisterVectorServiceServer(s grpc.ServiceRegistrar, srv VectorServiceServer) {
	s.RegisterService(&VectorService_ServiceDesc, srv)
}

func _VectorService_InsertVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertVectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VectorServiceServer).InsertVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VectorService_InsertVector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VectorServiceServer).InsertVector(ctx, req.(*InsertVectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VectorService_SearchVector_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchVectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VectorServiceServer).SearchVector(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VectorService_SearchVector_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VectorServiceServer).SearchVector(ctx, req.(*SearchVectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VectorService_ServiceDesc is the grpc.ServiceDesc for VectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vector_service.VectorService",
	HandlerType: (*VectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertVector",
			Handler:    _VectorService_InsertVector_Handler,
		},
		{
			MethodName: "SearchVector",
			Handler:    _VectorService_SearchVector_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vector_service.proto",
}

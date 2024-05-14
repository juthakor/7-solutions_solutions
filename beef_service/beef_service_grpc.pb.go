// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: beef_service.proto

package beef_service

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

// BeefServiceClient is the client API for BeefService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BeefServiceClient interface {
	GetBeefSummary(ctx context.Context, in *BeefRequest, opts ...grpc.CallOption) (*BeefResponse, error)
}

type beefServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBeefServiceClient(cc grpc.ClientConnInterface) BeefServiceClient {
	return &beefServiceClient{cc}
}

func (c *beefServiceClient) GetBeefSummary(ctx context.Context, in *BeefRequest, opts ...grpc.CallOption) (*BeefResponse, error) {
	out := new(BeefResponse)
	err := c.cc.Invoke(ctx, "/BeefService/GetBeefSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeefServiceServer is the server API for BeefService service.
// All implementations must embed UnimplementedBeefServiceServer
// for forward compatibility
type BeefServiceServer interface {
	GetBeefSummary(context.Context, *BeefRequest) (*BeefResponse, error)
	mustEmbedUnimplementedBeefServiceServer()
}

// UnimplementedBeefServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBeefServiceServer struct {
}

func (UnimplementedBeefServiceServer) GetBeefSummary(context.Context, *BeefRequest) (*BeefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeefSummary not implemented")
}
func (UnimplementedBeefServiceServer) mustEmbedUnimplementedBeefServiceServer() {}

// UnsafeBeefServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BeefServiceServer will
// result in compilation errors.
type UnsafeBeefServiceServer interface {
	mustEmbedUnimplementedBeefServiceServer()
}

func RegisterBeefServiceServer(s grpc.ServiceRegistrar, srv BeefServiceServer) {
	s.RegisterService(&BeefService_ServiceDesc, srv)
}

func _BeefService_GetBeefSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeefServiceServer).GetBeefSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BeefService/GetBeefSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeefServiceServer).GetBeefSummary(ctx, req.(*BeefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BeefService_ServiceDesc is the grpc.ServiceDesc for BeefService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BeefService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BeefService",
	HandlerType: (*BeefServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBeefSummary",
			Handler:    _BeefService_GetBeefSummary_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "beef_service.proto",
}

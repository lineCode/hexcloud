// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hexcloud

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

// HexagonServiceClient is the client API for HexagonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HexagonServiceClient interface {
	GetHexagonRing(ctx context.Context, in *HexagonRingRequest, opts ...grpc.CallOption) (*HexCubeResponse, error)
}

type hexagonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHexagonServiceClient(cc grpc.ClientConnInterface) HexagonServiceClient {
	return &hexagonServiceClient{cc}
}

func (c *hexagonServiceClient) GetHexagonRing(ctx context.Context, in *HexagonRingRequest, opts ...grpc.CallOption) (*HexCubeResponse, error) {
	out := new(HexCubeResponse)
	err := c.cc.Invoke(ctx, "/endpoints.hexworld.hexcloud.HexagonService/GetHexagonRing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HexagonServiceServer is the server API for HexagonService service.
// All implementations must embed UnimplementedHexagonServiceServer
// for forward compatibility
type HexagonServiceServer interface {
	GetHexagonRing(context.Context, *HexagonRingRequest) (*HexCubeResponse, error)
	mustEmbedUnimplementedHexagonServiceServer()
}

// UnimplementedHexagonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHexagonServiceServer struct {
}

func (UnimplementedHexagonServiceServer) GetHexagonRing(context.Context, *HexagonRingRequest) (*HexCubeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHexagonRing not implemented")
}
func (UnimplementedHexagonServiceServer) mustEmbedUnimplementedHexagonServiceServer() {}

// UnsafeHexagonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HexagonServiceServer will
// result in compilation errors.
type UnsafeHexagonServiceServer interface {
	mustEmbedUnimplementedHexagonServiceServer()
}

func RegisterHexagonServiceServer(s grpc.ServiceRegistrar, srv HexagonServiceServer) {
	s.RegisterService(&HexagonService_ServiceDesc, srv)
}

func _HexagonService_GetHexagonRing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexagonRingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetHexagonRing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/endpoints.hexworld.hexcloud.HexagonService/GetHexagonRing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetHexagonRing(ctx, req.(*HexagonRingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HexagonService_ServiceDesc is the grpc.ServiceDesc for HexagonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HexagonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoints.hexworld.hexcloud.HexagonService",
	HandlerType: (*HexagonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHexagonRing",
			Handler:    _HexagonService_GetHexagonRing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hexagon.proto",
}

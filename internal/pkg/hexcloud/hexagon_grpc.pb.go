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
	RepoAddHexagonInfo(ctx context.Context, in *HexInfoList, opts ...grpc.CallOption) (*Result, error)
	RepoDelHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*Result, error)
	RepoGetHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*HexInfoList, error)
	RepoGetAllHexagonInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HexInfoList, error)
	MapAdd(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error)
	MapGet(ctx context.Context, in *HexagonGetRequest, opts ...grpc.CallOption) (*HexLocationList, error)
	MapRemove(ctx context.Context, in *HexLocationList, opts ...grpc.CallOption) (*Result, error)
	GetStatusServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	GetStatusStorage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
	GetStatusClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error)
}

type hexagonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHexagonServiceClient(cc grpc.ClientConnInterface) HexagonServiceClient {
	return &hexagonServiceClient{cc}
}

func (c *hexagonServiceClient) RepoAddHexagonInfo(ctx context.Context, in *HexInfoList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/RepoAddHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoDelHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/RepoDelHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoGetHexagonInfo(ctx context.Context, in *HexIDList, opts ...grpc.CallOption) (*HexInfoList, error) {
	out := new(HexInfoList)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/RepoGetHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) RepoGetAllHexagonInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HexInfoList, error) {
	out := new(HexInfoList)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/RepoGetAllHexagonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapAdd(ctx context.Context, in *HexLocation, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/MapAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapGet(ctx context.Context, in *HexagonGetRequest, opts ...grpc.CallOption) (*HexLocationList, error) {
	out := new(HexLocationList)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/MapGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) MapRemove(ctx context.Context, in *HexLocationList, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/MapRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) GetStatusServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/GetStatusServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) GetStatusStorage(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/GetStatusStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hexagonServiceClient) GetStatusClients(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/hexcloud.HexagonService/GetStatusClients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HexagonServiceServer is the server API for HexagonService service.
// All implementations must embed UnimplementedHexagonServiceServer
// for forward compatibility
type HexagonServiceServer interface {
	RepoAddHexagonInfo(context.Context, *HexInfoList) (*Result, error)
	RepoDelHexagonInfo(context.Context, *HexIDList) (*Result, error)
	RepoGetHexagonInfo(context.Context, *HexIDList) (*HexInfoList, error)
	RepoGetAllHexagonInfo(context.Context, *Empty) (*HexInfoList, error)
	MapAdd(context.Context, *HexLocation) (*Result, error)
	MapGet(context.Context, *HexagonGetRequest) (*HexLocationList, error)
	MapRemove(context.Context, *HexLocationList) (*Result, error)
	GetStatusServer(context.Context, *Empty) (*Status, error)
	GetStatusStorage(context.Context, *Empty) (*Status, error)
	GetStatusClients(context.Context, *Empty) (*Status, error)
	mustEmbedUnimplementedHexagonServiceServer()
}

// UnimplementedHexagonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHexagonServiceServer struct {
}

func (UnimplementedHexagonServiceServer) RepoAddHexagonInfo(context.Context, *HexInfoList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoAddHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoDelHexagonInfo(context.Context, *HexIDList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoDelHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoGetHexagonInfo(context.Context, *HexIDList) (*HexInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoGetHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) RepoGetAllHexagonInfo(context.Context, *Empty) (*HexInfoList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepoGetAllHexagonInfo not implemented")
}
func (UnimplementedHexagonServiceServer) MapAdd(context.Context, *HexLocation) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapAdd not implemented")
}
func (UnimplementedHexagonServiceServer) MapGet(context.Context, *HexagonGetRequest) (*HexLocationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapGet not implemented")
}
func (UnimplementedHexagonServiceServer) MapRemove(context.Context, *HexLocationList) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapRemove not implemented")
}
func (UnimplementedHexagonServiceServer) GetStatusServer(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusServer not implemented")
}
func (UnimplementedHexagonServiceServer) GetStatusStorage(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusStorage not implemented")
}
func (UnimplementedHexagonServiceServer) GetStatusClients(context.Context, *Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatusClients not implemented")
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

func _HexagonService_RepoAddHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexInfoList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoAddHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/RepoAddHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoAddHexagonInfo(ctx, req.(*HexInfoList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoDelHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexIDList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoDelHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/RepoDelHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoDelHexagonInfo(ctx, req.(*HexIDList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoGetHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexIDList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoGetHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/RepoGetHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoGetHexagonInfo(ctx, req.(*HexIDList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_RepoGetAllHexagonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).RepoGetAllHexagonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/RepoGetAllHexagonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).RepoGetAllHexagonInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/MapAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapAdd(ctx, req.(*HexLocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexagonGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/MapGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapGet(ctx, req.(*HexagonGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_MapRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HexLocationList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).MapRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/MapRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).MapRemove(ctx, req.(*HexLocationList))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_GetStatusServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetStatusServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/GetStatusServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetStatusServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_GetStatusStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetStatusStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/GetStatusStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetStatusStorage(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HexagonService_GetStatusClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HexagonServiceServer).GetStatusClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hexcloud.HexagonService/GetStatusClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HexagonServiceServer).GetStatusClients(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HexagonService_ServiceDesc is the grpc.ServiceDesc for HexagonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HexagonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hexcloud.HexagonService",
	HandlerType: (*HexagonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RepoAddHexagonInfo",
			Handler:    _HexagonService_RepoAddHexagonInfo_Handler,
		},
		{
			MethodName: "RepoDelHexagonInfo",
			Handler:    _HexagonService_RepoDelHexagonInfo_Handler,
		},
		{
			MethodName: "RepoGetHexagonInfo",
			Handler:    _HexagonService_RepoGetHexagonInfo_Handler,
		},
		{
			MethodName: "RepoGetAllHexagonInfo",
			Handler:    _HexagonService_RepoGetAllHexagonInfo_Handler,
		},
		{
			MethodName: "MapAdd",
			Handler:    _HexagonService_MapAdd_Handler,
		},
		{
			MethodName: "MapGet",
			Handler:    _HexagonService_MapGet_Handler,
		},
		{
			MethodName: "MapRemove",
			Handler:    _HexagonService_MapRemove_Handler,
		},
		{
			MethodName: "GetStatusServer",
			Handler:    _HexagonService_GetStatusServer_Handler,
		},
		{
			MethodName: "GetStatusStorage",
			Handler:    _HexagonService_GetStatusStorage_Handler,
		},
		{
			MethodName: "GetStatusClients",
			Handler:    _HexagonService_GetStatusClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hexagon.proto",
}

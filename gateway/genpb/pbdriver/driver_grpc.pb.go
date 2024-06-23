// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbdriver

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

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverServiceClient interface {
	CreateDriver(ctx context.Context, in *DRequest, opts ...grpc.CallOption) (*DResponse, error)
	UpdateDriver(ctx context.Context, in *DRequest, opts ...grpc.CallOption) (*ResponseM, error)
	DeleteDriver(ctx context.Context, in *DResponse, opts ...grpc.CallOption) (*ResponseM, error)
	GetDriverById(ctx context.Context, in *DResponse, opts ...grpc.CallOption) (*DRequest, error)
	GetActiveDriverByLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) CreateDriver(ctx context.Context, in *DRequest, opts ...grpc.CallOption) (*DResponse, error) {
	out := new(DResponse)
	err := c.cc.Invoke(ctx, "/DriverService/CreateDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) UpdateDriver(ctx context.Context, in *DRequest, opts ...grpc.CallOption) (*ResponseM, error) {
	out := new(ResponseM)
	err := c.cc.Invoke(ctx, "/DriverService/UpdateDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) DeleteDriver(ctx context.Context, in *DResponse, opts ...grpc.CallOption) (*ResponseM, error) {
	out := new(ResponseM)
	err := c.cc.Invoke(ctx, "/DriverService/DeleteDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetDriverById(ctx context.Context, in *DResponse, opts ...grpc.CallOption) (*DRequest, error) {
	out := new(DRequest)
	err := c.cc.Invoke(ctx, "/DriverService/GetDriverById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) GetActiveDriverByLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*LocationResponse, error) {
	out := new(LocationResponse)
	err := c.cc.Invoke(ctx, "/DriverService/GetActiveDriverByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
// All implementations must embed UnimplementedDriverServiceServer
// for forward compatibility
type DriverServiceServer interface {
	CreateDriver(context.Context, *DRequest) (*DResponse, error)
	UpdateDriver(context.Context, *DRequest) (*ResponseM, error)
	DeleteDriver(context.Context, *DResponse) (*ResponseM, error)
	GetDriverById(context.Context, *DResponse) (*DRequest, error)
	GetActiveDriverByLocation(context.Context, *LocationRequest) (*LocationResponse, error)
	mustEmbedUnimplementedDriverServiceServer()
}

// UnimplementedDriverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDriverServiceServer struct {
}

func (UnimplementedDriverServiceServer) CreateDriver(context.Context, *DRequest) (*DResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDriver not implemented")
}
func (UnimplementedDriverServiceServer) UpdateDriver(context.Context, *DRequest) (*ResponseM, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDriver not implemented")
}
func (UnimplementedDriverServiceServer) DeleteDriver(context.Context, *DResponse) (*ResponseM, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDriver not implemented")
}
func (UnimplementedDriverServiceServer) GetDriverById(context.Context, *DResponse) (*DRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverById not implemented")
}
func (UnimplementedDriverServiceServer) GetActiveDriverByLocation(context.Context, *LocationRequest) (*LocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveDriverByLocation not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}

// UnsafeDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServiceServer will
// result in compilation errors.
type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv DriverServiceServer) {
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_CreateDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).CreateDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/CreateDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).CreateDriver(ctx, req.(*DRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_UpdateDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).UpdateDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/UpdateDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).UpdateDriver(ctx, req.(*DRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_DeleteDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).DeleteDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/DeleteDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).DeleteDriver(ctx, req.(*DResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetDriverById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetDriverById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/GetDriverById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetDriverById(ctx, req.(*DResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_GetActiveDriverByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).GetActiveDriverByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/GetActiveDriverByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).GetActiveDriverByLocation(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverService_ServiceDesc is the grpc.ServiceDesc for DriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDriver",
			Handler:    _DriverService_CreateDriver_Handler,
		},
		{
			MethodName: "UpdateDriver",
			Handler:    _DriverService_UpdateDriver_Handler,
		},
		{
			MethodName: "DeleteDriver",
			Handler:    _DriverService_DeleteDriver_Handler,
		},
		{
			MethodName: "GetDriverById",
			Handler:    _DriverService_GetDriverById_Handler,
		},
		{
			MethodName: "GetActiveDriverByLocation",
			Handler:    _DriverService_GetActiveDriverByLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/driverpb/driver.proto",
}

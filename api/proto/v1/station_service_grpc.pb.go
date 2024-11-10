// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: station_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StationService_ListStations_FullMethodName = "/api.v1.StationService/ListStations"
)

// StationServiceClient is the client API for StationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StationServiceClient interface {
	ListStations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStationsResponse, error)
}

type stationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStationServiceClient(cc grpc.ClientConnInterface) StationServiceClient {
	return &stationServiceClient{cc}
}

func (c *stationServiceClient) ListStations(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListStationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStationsResponse)
	err := c.cc.Invoke(ctx, StationService_ListStations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StationServiceServer is the server API for StationService service.
// All implementations must embed UnimplementedStationServiceServer
// for forward compatibility.
type StationServiceServer interface {
	ListStations(context.Context, *emptypb.Empty) (*ListStationsResponse, error)
	mustEmbedUnimplementedStationServiceServer()
}

// UnimplementedStationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStationServiceServer struct{}

func (UnimplementedStationServiceServer) ListStations(context.Context, *emptypb.Empty) (*ListStationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStations not implemented")
}
func (UnimplementedStationServiceServer) mustEmbedUnimplementedStationServiceServer() {}
func (UnimplementedStationServiceServer) testEmbeddedByValue()                        {}

// UnsafeStationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StationServiceServer will
// result in compilation errors.
type UnsafeStationServiceServer interface {
	mustEmbedUnimplementedStationServiceServer()
}

func RegisterStationServiceServer(s grpc.ServiceRegistrar, srv StationServiceServer) {
	// If the following call pancis, it indicates UnimplementedStationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StationService_ServiceDesc, srv)
}

func _StationService_ListStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StationServiceServer).ListStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StationService_ListStations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StationServiceServer).ListStations(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// StationService_ServiceDesc is the grpc.ServiceDesc for StationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.StationService",
	HandlerType: (*StationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListStations",
			Handler:    _StationService_ListStations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "station_service.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: discovery.proto

package pb

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
	Registration_RegisterDiscoveryHandler_FullMethodName = "/v0.Registration/RegisterDiscoveryHandler"
)

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Registration is the service advertised by the Akri Agent.
// Any `DiscoveryHandler` can register with the Akri Agent.
type RegistrationClient interface {
	RegisterDiscoveryHandler(ctx context.Context, in *RegisterDiscoveryHandlerRequest, opts ...grpc.CallOption) (*Empty, error)
}

type registrationClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationClient(cc grpc.ClientConnInterface) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) RegisterDiscoveryHandler(ctx context.Context, in *RegisterDiscoveryHandlerRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Registration_RegisterDiscoveryHandler_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
// All implementations must embed UnimplementedRegistrationServer
// for forward compatibility.
//
// Registration is the service advertised by the Akri Agent.
// Any `DiscoveryHandler` can register with the Akri Agent.
type RegistrationServer interface {
	RegisterDiscoveryHandler(context.Context, *RegisterDiscoveryHandlerRequest) (*Empty, error)
	mustEmbedUnimplementedRegistrationServer()
}

// UnimplementedRegistrationServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRegistrationServer struct{}

func (UnimplementedRegistrationServer) RegisterDiscoveryHandler(context.Context, *RegisterDiscoveryHandlerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDiscoveryHandler not implemented")
}
func (UnimplementedRegistrationServer) mustEmbedUnimplementedRegistrationServer() {}
func (UnimplementedRegistrationServer) testEmbeddedByValue()                      {}

// UnsafeRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServer will
// result in compilation errors.
type UnsafeRegistrationServer interface {
	mustEmbedUnimplementedRegistrationServer()
}

func RegisterRegistrationServer(s grpc.ServiceRegistrar, srv RegistrationServer) {
	// If the following call pancis, it indicates UnimplementedRegistrationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Registration_ServiceDesc, srv)
}

func _Registration_RegisterDiscoveryHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDiscoveryHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).RegisterDiscoveryHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Registration_RegisterDiscoveryHandler_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).RegisterDiscoveryHandler(ctx, req.(*RegisterDiscoveryHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Registration_ServiceDesc is the grpc.ServiceDesc for Registration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v0.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDiscoveryHandler",
			Handler:    _Registration_RegisterDiscoveryHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery.proto",
}

const (
	DiscoveryHandler_Discover_FullMethodName = "/v0.DiscoveryHandler/Discover"
)

// DiscoveryHandlerClient is the client API for DiscoveryHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoveryHandlerClient interface {
	Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DiscoverResponse], error)
}

type discoveryHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoveryHandlerClient(cc grpc.ClientConnInterface) DiscoveryHandlerClient {
	return &discoveryHandlerClient{cc}
}

func (c *discoveryHandlerClient) Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DiscoverResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DiscoveryHandler_ServiceDesc.Streams[0], DiscoveryHandler_Discover_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[DiscoverRequest, DiscoverResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DiscoveryHandler_DiscoverClient = grpc.ServerStreamingClient[DiscoverResponse]

// DiscoveryHandlerServer is the server API for DiscoveryHandler service.
// All implementations must embed UnimplementedDiscoveryHandlerServer
// for forward compatibility.
type DiscoveryHandlerServer interface {
	Discover(*DiscoverRequest, grpc.ServerStreamingServer[DiscoverResponse]) error
	mustEmbedUnimplementedDiscoveryHandlerServer()
}

// UnimplementedDiscoveryHandlerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDiscoveryHandlerServer struct{}

func (UnimplementedDiscoveryHandlerServer) Discover(*DiscoverRequest, grpc.ServerStreamingServer[DiscoverResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (UnimplementedDiscoveryHandlerServer) mustEmbedUnimplementedDiscoveryHandlerServer() {}
func (UnimplementedDiscoveryHandlerServer) testEmbeddedByValue()                          {}

// UnsafeDiscoveryHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoveryHandlerServer will
// result in compilation errors.
type UnsafeDiscoveryHandlerServer interface {
	mustEmbedUnimplementedDiscoveryHandlerServer()
}

func RegisterDiscoveryHandlerServer(s grpc.ServiceRegistrar, srv DiscoveryHandlerServer) {
	// If the following call pancis, it indicates UnimplementedDiscoveryHandlerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DiscoveryHandler_ServiceDesc, srv)
}

func _DiscoveryHandler_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiscoverRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiscoveryHandlerServer).Discover(m, &grpc.GenericServerStream[DiscoverRequest, DiscoverResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DiscoveryHandler_DiscoverServer = grpc.ServerStreamingServer[DiscoverResponse]

// DiscoveryHandler_ServiceDesc is the grpc.ServiceDesc for DiscoveryHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiscoveryHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v0.DiscoveryHandler",
	HandlerType: (*DiscoveryHandlerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _DiscoveryHandler_Discover_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "discovery.proto",
}

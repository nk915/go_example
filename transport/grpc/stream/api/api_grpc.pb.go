// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: api/api.proto

package api

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
	Stream_StartStreaming_FullMethodName = "/api.Stream/StartStreaming"
)

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClient interface {
	StartStreaming(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamMessage, StreamMessage], error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) StartStreaming(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamMessage, StreamMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Stream_ServiceDesc.Streams[0], Stream_StartStreaming_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamMessage, StreamMessage]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Stream_StartStreamingClient = grpc.BidiStreamingClient[StreamMessage, StreamMessage]

// StreamServer is the server API for Stream service.
// All implementations must embed UnimplementedStreamServer
// for forward compatibility.
type StreamServer interface {
	StartStreaming(grpc.BidiStreamingServer[StreamMessage, StreamMessage]) error
	mustEmbedUnimplementedStreamServer()
}

// UnimplementedStreamServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamServer struct{}

func (UnimplementedStreamServer) StartStreaming(grpc.BidiStreamingServer[StreamMessage, StreamMessage]) error {
	return status.Errorf(codes.Unimplemented, "method StartStreaming not implemented")
}
func (UnimplementedStreamServer) mustEmbedUnimplementedStreamServer() {}
func (UnimplementedStreamServer) testEmbeddedByValue()                {}

// UnsafeStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServer will
// result in compilation errors.
type UnsafeStreamServer interface {
	mustEmbedUnimplementedStreamServer()
}

func RegisterStreamServer(s grpc.ServiceRegistrar, srv StreamServer) {
	// If the following call pancis, it indicates UnimplementedStreamServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Stream_ServiceDesc, srv)
}

func _Stream_StartStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServer).StartStreaming(&grpc.GenericServerStream[StreamMessage, StreamMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Stream_StartStreamingServer = grpc.BidiStreamingServer[StreamMessage, StreamMessage]

// Stream_ServiceDesc is the grpc.ServiceDesc for Stream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartStreaming",
			Handler:       _Stream_StartStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api.proto",
}

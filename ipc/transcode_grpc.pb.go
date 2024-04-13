// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ipc

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

// TranscodeClient is the client API for Transcode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranscodeClient interface {
	ToJSON(ctx context.Context, in *JsonIn, opts ...grpc.CallOption) (*JsonOut, error)
}

type transcodeClient struct {
	cc grpc.ClientConnInterface
}

func NewTranscodeClient(cc grpc.ClientConnInterface) TranscodeClient {
	return &transcodeClient{cc}
}

func (c *transcodeClient) ToJSON(ctx context.Context, in *JsonIn, opts ...grpc.CallOption) (*JsonOut, error) {
	out := new(JsonOut)
	err := c.cc.Invoke(ctx, "/ipc.Transcode/ToJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranscodeServer is the server API for Transcode service.
// All implementations must embed UnimplementedTranscodeServer
// for forward compatibility
type TranscodeServer interface {
	ToJSON(context.Context, *JsonIn) (*JsonOut, error)
	mustEmbedUnimplementedTranscodeServer()
}

// UnimplementedTranscodeServer must be embedded to have forward compatible implementations.
type UnimplementedTranscodeServer struct {
}

func (UnimplementedTranscodeServer) ToJSON(context.Context, *JsonIn) (*JsonOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToJSON not implemented")
}
func (UnimplementedTranscodeServer) mustEmbedUnimplementedTranscodeServer() {}

// UnsafeTranscodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranscodeServer will
// result in compilation errors.
type UnsafeTranscodeServer interface {
	mustEmbedUnimplementedTranscodeServer()
}

func RegisterTranscodeServer(s grpc.ServiceRegistrar, srv TranscodeServer) {
	s.RegisterService(&Transcode_ServiceDesc, srv)
}

func _Transcode_ToJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JsonIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscodeServer).ToJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ipc.Transcode/ToJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscodeServer).ToJSON(ctx, req.(*JsonIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Transcode_ServiceDesc is the grpc.ServiceDesc for Transcode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transcode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ipc.Transcode",
	HandlerType: (*TranscodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ToJSON",
			Handler:    _Transcode_ToJSON_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipc/transcode.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package storage

import (
	context "context"
	common "github.com/zikwall/monit/src/protobuf/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	WriteHeatmap(ctx context.Context, in *HeatmapMessage, opts ...grpc.CallOption) (*common.EmptyResponse, error)
	WriteMetric(ctx context.Context, in *MetricMessage, opts ...grpc.CallOption) (*common.EmptyResponse, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) WriteHeatmap(ctx context.Context, in *HeatmapMessage, opts ...grpc.CallOption) (*common.EmptyResponse, error) {
	out := new(common.EmptyResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/WriteHeatmap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) WriteMetric(ctx context.Context, in *MetricMessage, opts ...grpc.CallOption) (*common.EmptyResponse, error) {
	out := new(common.EmptyResponse)
	err := c.cc.Invoke(ctx, "/storage.Storage/WriteMetric", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	WriteHeatmap(context.Context, *HeatmapMessage) (*common.EmptyResponse, error)
	WriteMetric(context.Context, *MetricMessage) (*common.EmptyResponse, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) WriteHeatmap(context.Context, *HeatmapMessage) (*common.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteHeatmap not implemented")
}
func (UnimplementedStorageServer) WriteMetric(context.Context, *MetricMessage) (*common.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMetric not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_WriteHeatmap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeatmapMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).WriteHeatmap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/WriteHeatmap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).WriteHeatmap(ctx, req.(*HeatmapMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_WriteMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).WriteMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.Storage/WriteMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).WriteMetric(ctx, req.(*MetricMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteHeatmap",
			Handler:    _Storage_WriteHeatmap_Handler,
		},
		{
			MethodName: "WriteMetric",
			Handler:    _Storage_WriteMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage.proto",
}

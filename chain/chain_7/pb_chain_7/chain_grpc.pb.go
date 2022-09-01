// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: pb_chain_7/chain.proto

package pb_chain_7

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

// Chain_7Client is the client API for Chain_7 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Chain_7Client interface {
	GetIds(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type chain_7Client struct {
	cc grpc.ClientConnInterface
}

func NewChain_7Client(cc grpc.ClientConnInterface) Chain_7Client {
	return &chain_7Client{cc}
}

func (c *chain_7Client) GetIds(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/chain_7.Chain_7/GetIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Chain_7Server is the server API for Chain_7 service.
// All implementations must embed UnimplementedChain_7Server
// for forward compatibility
type Chain_7Server interface {
	GetIds(context.Context, *Req) (*Res, error)
	mustEmbedUnimplementedChain_7Server()
}

// UnimplementedChain_7Server must be embedded to have forward compatible implementations.
type UnimplementedChain_7Server struct {
}

func (UnimplementedChain_7Server) GetIds(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIds not implemented")
}
func (UnimplementedChain_7Server) mustEmbedUnimplementedChain_7Server() {}

// UnsafeChain_7Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Chain_7Server will
// result in compilation errors.
type UnsafeChain_7Server interface {
	mustEmbedUnimplementedChain_7Server()
}

func RegisterChain_7Server(s grpc.ServiceRegistrar, srv Chain_7Server) {
	s.RegisterService(&Chain_7_ServiceDesc, srv)
}

func _Chain_7_GetIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Chain_7Server).GetIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chain_7.Chain_7/GetIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Chain_7Server).GetIds(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Chain_7_ServiceDesc is the grpc.ServiceDesc for Chain_7 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chain_7_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chain_7.Chain_7",
	HandlerType: (*Chain_7Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIds",
			Handler:    _Chain_7_GetIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_chain_7/chain.proto",
}

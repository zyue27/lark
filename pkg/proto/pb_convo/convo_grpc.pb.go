// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: pb_convo/convo.proto

package pb_convo

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

const (
	Convo_ConvoList_FullMethodName        = "/pb_convo.Convo/ConvoList"
	Convo_ConvoChatSeqList_FullMethodName = "/pb_convo.Convo/ConvoChatSeqList"
)

// ConvoClient is the client API for Convo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConvoClient interface {
	ConvoList(ctx context.Context, in *ConvoListReq, opts ...grpc.CallOption) (*ConvoListResp, error)
	ConvoChatSeqList(ctx context.Context, in *ConvoChatSeqListReq, opts ...grpc.CallOption) (*ConvoChatSeqListResp, error)
}

type convoClient struct {
	cc grpc.ClientConnInterface
}

func NewConvoClient(cc grpc.ClientConnInterface) ConvoClient {
	return &convoClient{cc}
}

func (c *convoClient) ConvoList(ctx context.Context, in *ConvoListReq, opts ...grpc.CallOption) (*ConvoListResp, error) {
	out := new(ConvoListResp)
	err := c.cc.Invoke(ctx, Convo_ConvoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *convoClient) ConvoChatSeqList(ctx context.Context, in *ConvoChatSeqListReq, opts ...grpc.CallOption) (*ConvoChatSeqListResp, error) {
	out := new(ConvoChatSeqListResp)
	err := c.cc.Invoke(ctx, Convo_ConvoChatSeqList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConvoServer is the server API for Convo service.
// All implementations must embed UnimplementedConvoServer
// for forward compatibility
type ConvoServer interface {
	ConvoList(context.Context, *ConvoListReq) (*ConvoListResp, error)
	ConvoChatSeqList(context.Context, *ConvoChatSeqListReq) (*ConvoChatSeqListResp, error)
	mustEmbedUnimplementedConvoServer()
}

// UnimplementedConvoServer must be embedded to have forward compatible implementations.
type UnimplementedConvoServer struct {
}

func (UnimplementedConvoServer) ConvoList(context.Context, *ConvoListReq) (*ConvoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvoList not implemented")
}
func (UnimplementedConvoServer) ConvoChatSeqList(context.Context, *ConvoChatSeqListReq) (*ConvoChatSeqListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvoChatSeqList not implemented")
}
func (UnimplementedConvoServer) mustEmbedUnimplementedConvoServer() {}

// UnsafeConvoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConvoServer will
// result in compilation errors.
type UnsafeConvoServer interface {
	mustEmbedUnimplementedConvoServer()
}

func RegisterConvoServer(s grpc.ServiceRegistrar, srv ConvoServer) {
	s.RegisterService(&Convo_ServiceDesc, srv)
}

func _Convo_ConvoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConvoServer).ConvoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Convo_ConvoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConvoServer).ConvoList(ctx, req.(*ConvoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Convo_ConvoChatSeqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvoChatSeqListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConvoServer).ConvoChatSeqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Convo_ConvoChatSeqList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConvoServer).ConvoChatSeqList(ctx, req.(*ConvoChatSeqListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Convo_ServiceDesc is the grpc.ServiceDesc for Convo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Convo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_convo.Convo",
	HandlerType: (*ConvoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvoList",
			Handler:    _Convo_ConvoList_Handler,
		},
		{
			MethodName: "ConvoChatSeqList",
			Handler:    _Convo_ConvoChatSeqList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_convo/convo.proto",
}

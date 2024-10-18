// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.1
// source: pb_user/user.proto

package pb_user

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
	User_EditUserInfo_FullMethodName         = "/pb_user.User/EditUserInfo"
	User_GetUserInfo_FullMethodName          = "/pb_user.User/GetUserInfo"
	User_GetBasicUserInfo_FullMethodName     = "/pb_user.User/GetBasicUserInfo"
	User_GetUserList_FullMethodName          = "/pb_user.User/GetUserList"
	User_SearchUser_FullMethodName           = "/pb_user.User/SearchUser"
	User_UploadAvatar_FullMethodName         = "/pb_user.User/UploadAvatar"
	User_GetBasicUserInfoList_FullMethodName = "/pb_user.User/GetBasicUserInfoList"
	User_GetServerIdList_FullMethodName      = "/pb_user.User/GetServerIdList"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	EditUserInfo(ctx context.Context, in *EditUserInfoReq, opts ...grpc.CallOption) (*EditUserInfoResp, error)
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	GetBasicUserInfo(ctx context.Context, in *GetBasicUserInfoReq, opts ...grpc.CallOption) (*GetBasicUserInfoResp, error)
	GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error)
	SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error)
	UploadAvatar(ctx context.Context, in *UploadAvatarReq, opts ...grpc.CallOption) (*UploadAvatarResp, error)
	GetBasicUserInfoList(ctx context.Context, in *GetBasicUserInfoListReq, opts ...grpc.CallOption) (*GetBasicUserInfoListResp, error)
	GetServerIdList(ctx context.Context, in *GetServerIdListReq, opts ...grpc.CallOption) (*GetServerIdListResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) EditUserInfo(ctx context.Context, in *EditUserInfoReq, opts ...grpc.CallOption) (*EditUserInfoResp, error) {
	out := new(EditUserInfoResp)
	err := c.cc.Invoke(ctx, User_EditUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, User_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetBasicUserInfo(ctx context.Context, in *GetBasicUserInfoReq, opts ...grpc.CallOption) (*GetBasicUserInfoResp, error) {
	out := new(GetBasicUserInfoResp)
	err := c.cc.Invoke(ctx, User_GetBasicUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error) {
	out := new(GetUserListResp)
	err := c.cc.Invoke(ctx, User_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error) {
	out := new(SearchUserResp)
	err := c.cc.Invoke(ctx, User_SearchUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UploadAvatar(ctx context.Context, in *UploadAvatarReq, opts ...grpc.CallOption) (*UploadAvatarResp, error) {
	out := new(UploadAvatarResp)
	err := c.cc.Invoke(ctx, User_UploadAvatar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetBasicUserInfoList(ctx context.Context, in *GetBasicUserInfoListReq, opts ...grpc.CallOption) (*GetBasicUserInfoListResp, error) {
	out := new(GetBasicUserInfoListResp)
	err := c.cc.Invoke(ctx, User_GetBasicUserInfoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetServerIdList(ctx context.Context, in *GetServerIdListReq, opts ...grpc.CallOption) (*GetServerIdListResp, error) {
	out := new(GetServerIdListResp)
	err := c.cc.Invoke(ctx, User_GetServerIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	EditUserInfo(context.Context, *EditUserInfoReq) (*EditUserInfoResp, error)
	GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	GetBasicUserInfo(context.Context, *GetBasicUserInfoReq) (*GetBasicUserInfoResp, error)
	GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error)
	SearchUser(context.Context, *SearchUserReq) (*SearchUserResp, error)
	UploadAvatar(context.Context, *UploadAvatarReq) (*UploadAvatarResp, error)
	GetBasicUserInfoList(context.Context, *GetBasicUserInfoListReq) (*GetBasicUserInfoListResp, error)
	GetServerIdList(context.Context, *GetServerIdListReq) (*GetServerIdListResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) EditUserInfo(context.Context, *EditUserInfoReq) (*EditUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserInfo not implemented")
}
func (UnimplementedUserServer) GetUserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServer) GetBasicUserInfo(context.Context, *GetBasicUserInfoReq) (*GetBasicUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasicUserInfo not implemented")
}
func (UnimplementedUserServer) GetUserList(context.Context, *GetUserListReq) (*GetUserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedUserServer) SearchUser(context.Context, *SearchUserReq) (*SearchUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (UnimplementedUserServer) UploadAvatar(context.Context, *UploadAvatarReq) (*UploadAvatarResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedUserServer) GetBasicUserInfoList(context.Context, *GetBasicUserInfoListReq) (*GetBasicUserInfoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasicUserInfoList not implemented")
}
func (UnimplementedUserServer) GetServerIdList(context.Context, *GetServerIdListReq) (*GetServerIdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerIdList not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_EditUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EditUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_EditUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EditUserInfo(ctx, req.(*EditUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetBasicUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasicUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetBasicUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetBasicUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetBasicUserInfo(ctx, req.(*GetBasicUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserList(ctx, req.(*GetUserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_SearchUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SearchUser(ctx, req.(*SearchUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAvatarReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UploadAvatar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UploadAvatar(ctx, req.(*UploadAvatarReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetBasicUserInfoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBasicUserInfoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetBasicUserInfoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetBasicUserInfoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetBasicUserInfoList(ctx, req.(*GetBasicUserInfoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetServerIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerIdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetServerIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetServerIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetServerIdList(ctx, req.(*GetServerIdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb_user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EditUserInfo",
			Handler:    _User_EditUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "GetBasicUserInfo",
			Handler:    _User_GetBasicUserInfo_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _User_GetUserList_Handler,
		},
		{
			MethodName: "SearchUser",
			Handler:    _User_SearchUser_Handler,
		},
		{
			MethodName: "UploadAvatar",
			Handler:    _User_UploadAvatar_Handler,
		},
		{
			MethodName: "GetBasicUserInfoList",
			Handler:    _User_GetBasicUserInfoList_Handler,
		},
		{
			MethodName: "GetServerIdList",
			Handler:    _User_GetServerIdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb_user/user.proto",
}

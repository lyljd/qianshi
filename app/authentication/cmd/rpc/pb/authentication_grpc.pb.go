// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: authentication.proto

package __

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
	Authentication_GenerateToken_FullMethodName        = "/service.Authentication/GenerateToken"
	Authentication_GenerateRefreshToken_FullMethodName = "/service.Authentication/GenerateRefreshToken"
	Authentication_VerifyToken_FullMethodName          = "/service.Authentication/VerifyToken"
	Authentication_VerifyRefreshToken_FullMethodName   = "/service.Authentication/VerifyRefreshToken"
)

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	GenerateRefreshToken(ctx context.Context, in *GenerateRefreshTokenReq, opts ...grpc.CallOption) (*GenerateRefreshTokenResp, error)
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenResp, error)
	VerifyRefreshToken(ctx context.Context, in *VerifyRefreshTokenReq, opts ...grpc.CallOption) (*VerifyRefreshTokenResp, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, Authentication_GenerateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GenerateRefreshToken(ctx context.Context, in *GenerateRefreshTokenReq, opts ...grpc.CallOption) (*GenerateRefreshTokenResp, error) {
	out := new(GenerateRefreshTokenResp)
	err := c.cc.Invoke(ctx, Authentication_GenerateRefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenResp, error) {
	out := new(VerifyTokenResp)
	err := c.cc.Invoke(ctx, Authentication_VerifyToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) VerifyRefreshToken(ctx context.Context, in *VerifyRefreshTokenReq, opts ...grpc.CallOption) (*VerifyRefreshTokenResp, error) {
	out := new(VerifyRefreshTokenResp)
	err := c.cc.Invoke(ctx, Authentication_VerifyRefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	GenerateRefreshToken(context.Context, *GenerateRefreshTokenReq) (*GenerateRefreshTokenResp, error)
	VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenResp, error)
	VerifyRefreshToken(context.Context, *VerifyRefreshTokenReq) (*VerifyRefreshTokenResp, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedAuthenticationServer) GenerateRefreshToken(context.Context, *GenerateRefreshTokenReq) (*GenerateRefreshTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateRefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthenticationServer) VerifyRefreshToken(context.Context, *VerifyRefreshTokenReq) (*VerifyRefreshTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyRefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_GenerateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GenerateRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GenerateRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_GenerateRefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GenerateRefreshToken(ctx, req.(*GenerateRefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_VerifyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).VerifyToken(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_VerifyRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).VerifyRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_VerifyRefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).VerifyRefreshToken(ctx, req.(*VerifyRefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateToken",
			Handler:    _Authentication_GenerateToken_Handler,
		},
		{
			MethodName: "GenerateRefreshToken",
			Handler:    _Authentication_GenerateRefreshToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _Authentication_VerifyToken_Handler,
		},
		{
			MethodName: "VerifyRefreshToken",
			Handler:    _Authentication_VerifyRefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

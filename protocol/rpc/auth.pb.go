// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AuthLoginReq struct {
	UID int64 `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
}

func (m *AuthLoginReq) Reset()                    { *m = AuthLoginReq{} }
func (m *AuthLoginReq) String() string            { return proto.CompactTextString(m) }
func (*AuthLoginReq) ProtoMessage()               {}
func (*AuthLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type AuthLoginRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *AuthLoginRes) Reset()                    { *m = AuthLoginRes{} }
func (m *AuthLoginRes) String() string            { return proto.CompactTextString(m) }
func (*AuthLoginRes) ProtoMessage()               {}
func (*AuthLoginRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type AuthAuthReq struct {
	UID   int64  `protobuf:"varint,1,opt,name=UID" json:"UID,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *AuthAuthReq) Reset()                    { *m = AuthAuthReq{} }
func (m *AuthAuthReq) String() string            { return proto.CompactTextString(m) }
func (*AuthAuthReq) ProtoMessage()               {}
func (*AuthAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type AuthAuthRes struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,2,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *AuthAuthRes) Reset()                    { *m = AuthAuthRes{} }
func (m *AuthAuthRes) String() string            { return proto.CompactTextString(m) }
func (*AuthAuthRes) ProtoMessage()               {}
func (*AuthAuthRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto.RegisterType((*AuthLoginReq)(nil), "rpc.AuthLoginReq")
	proto.RegisterType((*AuthLoginRes)(nil), "rpc.AuthLoginRes")
	proto.RegisterType((*AuthAuthReq)(nil), "rpc.AuthAuthReq")
	proto.RegisterType((*AuthAuthRes)(nil), "rpc.AuthAuthRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AuthServerRPC service

type AuthServerRPCClient interface {
	Login(ctx context.Context, in *AuthLoginReq, opts ...grpc.CallOption) (*AuthLoginRes, error)
	Auth(ctx context.Context, in *AuthAuthReq, opts ...grpc.CallOption) (*AuthAuthRes, error)
}

type authServerRPCClient struct {
	cc *grpc.ClientConn
}

func NewAuthServerRPCClient(cc *grpc.ClientConn) AuthServerRPCClient {
	return &authServerRPCClient{cc}
}

func (c *authServerRPCClient) Login(ctx context.Context, in *AuthLoginReq, opts ...grpc.CallOption) (*AuthLoginRes, error) {
	out := new(AuthLoginRes)
	err := grpc.Invoke(ctx, "/rpc.AuthServerRPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServerRPCClient) Auth(ctx context.Context, in *AuthAuthReq, opts ...grpc.CallOption) (*AuthAuthRes, error) {
	out := new(AuthAuthRes)
	err := grpc.Invoke(ctx, "/rpc.AuthServerRPC/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthServerRPC service

type AuthServerRPCServer interface {
	Login(context.Context, *AuthLoginReq) (*AuthLoginRes, error)
	Auth(context.Context, *AuthAuthReq) (*AuthAuthRes, error)
}

func RegisterAuthServerRPCServer(s *grpc.Server, srv AuthServerRPCServer) {
	s.RegisterService(&_AuthServerRPC_serviceDesc, srv)
}

func _AuthServerRPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServerRPCServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.AuthServerRPC/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServerRPCServer).Login(ctx, req.(*AuthLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthServerRPC_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServerRPCServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.AuthServerRPC/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServerRPCServer).Auth(ctx, req.(*AuthAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthServerRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.AuthServerRPC",
	HandlerType: (*AuthServerRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthServerRPC_Login_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _AuthServerRPC_Auth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x56, 0x52, 0xe0, 0xe2, 0x71,
	0x2c, 0x2d, 0xc9, 0xf0, 0xc9, 0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x14, 0x12, 0xe0, 0x62, 0x0e,
	0xf5, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x02, 0x31, 0x95, 0xc2, 0x50, 0x54, 0x14,
	0x0b, 0x49, 0x70, 0xb1, 0xa7, 0x16, 0x15, 0x39, 0xe7, 0xa7, 0xa4, 0x82, 0x55, 0xf1, 0x06, 0xc1,
	0xb8, 0x42, 0x62, 0x5c, 0x6c, 0xa9, 0x45, 0x45, 0xc1, 0x25, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x50, 0x9e, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x7e, 0x76, 0x6a, 0x9e, 0x04, 0x33, 0x58,
	0x18, 0xc2, 0x51, 0x32, 0xe5, 0xe2, 0x06, 0x99, 0x0b, 0xc2, 0x58, 0x2d, 0x46, 0x68, 0x63, 0x42,
	0xd6, 0x66, 0x8f, 0xac, 0x8d, 0x0c, 0xd7, 0x18, 0xe5, 0x71, 0xf1, 0x82, 0x34, 0x07, 0xa7, 0x16,
	0x95, 0xa5, 0x16, 0x05, 0x05, 0x38, 0x0b, 0xe9, 0x73, 0xb1, 0x82, 0x3d, 0x27, 0x24, 0xa8, 0x57,
	0x54, 0x90, 0xac, 0x87, 0x1c, 0x1c, 0x52, 0x18, 0x42, 0xc5, 0x4a, 0x0c, 0x42, 0x3a, 0x5c, 0x2c,
	0x20, 0x11, 0x21, 0x01, 0xb8, 0x24, 0xd4, 0x13, 0x52, 0xe8, 0x22, 0xc5, 0x4a, 0x0c, 0x49, 0x6c,
	0xe0, 0xd0, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x7e, 0x70, 0x92, 0x7b, 0x01, 0x00,
	0x00,
}
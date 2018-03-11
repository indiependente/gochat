// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

/*
Package gochat is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
*/
package gochat

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "gochat.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "gochat.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "gochat.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "gochat.LogoutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Chat service

type ChatClient interface {
	// Registers a user to the chat
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Unregisters a user to the chat
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/gochat.Chat/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/gochat.Chat/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Chat service

type ChatServer interface {
	// Registers a user to the chat
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Unregisters a user to the chat
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gochat.Chat/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gochat.Chat/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gochat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Chat_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Chat_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xcf, 0x07, 0xf1, 0x94, 0xec, 0xb8, 0x78,
	0x7c, 0xf2, 0xd3, 0x33, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58,
	0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x29, 0x2e,
	0x8e, 0x82, 0xc4, 0xe2, 0xe2, 0xf2, 0xfc, 0xa2, 0x14, 0x09, 0x26, 0xb0, 0x38, 0x9c, 0xaf, 0xa4,
	0xca, 0xc5, 0x0b, 0xd5, 0x5f, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc2, 0xc5, 0x5a, 0x92,
	0x9f, 0x9d, 0x9a, 0x07, 0x35, 0x01, 0xc2, 0x81, 0x2a, 0xcb, 0x2f, 0x2d, 0x81, 0xd9, 0x83, 0x5d,
	0x99, 0x00, 0x17, 0x1f, 0x4c, 0x19, 0xc4, 0x38, 0xa3, 0x4a, 0x2e, 0x16, 0xe7, 0x8c, 0xc4, 0x12,
	0x21, 0x33, 0x2e, 0x56, 0xb0, 0x3d, 0x42, 0x22, 0x7a, 0x10, 0x97, 0xeb, 0x21, 0x3b, 0x5b, 0x4a,
	0x14, 0x4d, 0x14, 0xa2, 0x5b, 0x89, 0x41, 0xc8, 0x92, 0x8b, 0x0d, 0x62, 0xa2, 0x10, 0xb2, 0x12,
	0x84, 0x43, 0xa4, 0xc4, 0xd0, 0x85, 0x61, 0x5a, 0x9d, 0x38, 0xa2, 0xa0, 0x81, 0x94, 0xc4, 0x06,
	0x0e, 0x33, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4f, 0xe0, 0x24, 0x41, 0x01, 0x00,
	0x00,
}

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
	StreamRequest
	StreamResponse
*/
package gochat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type StreamRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StreamRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*StreamResponse_ClientLogin
	//	*StreamResponse_ClientLogout
	//	*StreamResponse_ClientMessage
	//	*StreamResponse_ServerShutdown
	Event isStreamResponse_Event `protobuf_oneof:"event"`
}

func (m *StreamResponse) Reset()                    { *m = StreamResponse{} }
func (m *StreamResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()               {}
func (*StreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_ClientLogin struct {
	ClientLogin *StreamResponse_Login `protobuf:"bytes,2,opt,name=client_login,json=clientLogin,oneof"`
}
type StreamResponse_ClientLogout struct {
	ClientLogout *StreamResponse_Logout `protobuf:"bytes,3,opt,name=client_logout,json=clientLogout,oneof"`
}
type StreamResponse_ClientMessage struct {
	ClientMessage *StreamResponse_Message `protobuf:"bytes,4,opt,name=client_message,json=clientMessage,oneof"`
}
type StreamResponse_ServerShutdown struct {
	ServerShutdown *StreamResponse_Shutdown `protobuf:"bytes,5,opt,name=server_shutdown,json=serverShutdown,oneof"`
}

func (*StreamResponse_ClientLogin) isStreamResponse_Event()    {}
func (*StreamResponse_ClientLogout) isStreamResponse_Event()   {}
func (*StreamResponse_ClientMessage) isStreamResponse_Event()  {}
func (*StreamResponse_ServerShutdown) isStreamResponse_Event() {}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamResponse) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *StreamResponse) GetClientLogin() *StreamResponse_Login {
	if x, ok := m.GetEvent().(*StreamResponse_ClientLogin); ok {
		return x.ClientLogin
	}
	return nil
}

func (m *StreamResponse) GetClientLogout() *StreamResponse_Logout {
	if x, ok := m.GetEvent().(*StreamResponse_ClientLogout); ok {
		return x.ClientLogout
	}
	return nil
}

func (m *StreamResponse) GetClientMessage() *StreamResponse_Message {
	if x, ok := m.GetEvent().(*StreamResponse_ClientMessage); ok {
		return x.ClientMessage
	}
	return nil
}

func (m *StreamResponse) GetServerShutdown() *StreamResponse_Shutdown {
	if x, ok := m.GetEvent().(*StreamResponse_ServerShutdown); ok {
		return x.ServerShutdown
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamResponse_OneofMarshaler, _StreamResponse_OneofUnmarshaler, _StreamResponse_OneofSizer, []interface{}{
		(*StreamResponse_ClientLogin)(nil),
		(*StreamResponse_ClientLogout)(nil),
		(*StreamResponse_ClientMessage)(nil),
		(*StreamResponse_ServerShutdown)(nil),
	}
}

func _StreamResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_ClientLogin:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientLogin); err != nil {
			return err
		}
	case *StreamResponse_ClientLogout:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientLogout); err != nil {
			return err
		}
	case *StreamResponse_ClientMessage:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientMessage); err != nil {
			return err
		}
	case *StreamResponse_ServerShutdown:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ServerShutdown); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamResponse.Event has unexpected type %T", x)
	}
	return nil
}

func _StreamResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamResponse)
	switch tag {
	case 2: // event.client_login
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Login)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientLogin{msg}
		return true, err
	case 3: // event.client_logout
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Logout)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientLogout{msg}
		return true, err
	case 4: // event.client_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Message)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientMessage{msg}
		return true, err
	case 5: // event.server_shutdown
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Shutdown)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ServerShutdown{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_ClientLogin:
		s := proto.Size(x.ClientLogin)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientLogout:
		s := proto.Size(x.ClientLogout)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientMessage:
		s := proto.Size(x.ClientMessage)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ServerShutdown:
		s := proto.Size(x.ServerShutdown)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamResponse_Login struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StreamResponse_Login) Reset()                    { *m = StreamResponse_Login{} }
func (m *StreamResponse_Login) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Login) ProtoMessage()               {}
func (*StreamResponse_Login) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *StreamResponse_Login) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Logout struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StreamResponse_Logout) Reset()                    { *m = StreamResponse_Logout{} }
func (m *StreamResponse_Logout) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Logout) ProtoMessage()               {}
func (*StreamResponse_Logout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *StreamResponse_Logout) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Message struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *StreamResponse_Message) Reset()                    { *m = StreamResponse_Message{} }
func (m *StreamResponse_Message) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Message) ProtoMessage()               {}
func (*StreamResponse_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2} }

func (m *StreamResponse_Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamResponse_Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse_Shutdown struct {
}

func (m *StreamResponse_Shutdown) Reset()                    { *m = StreamResponse_Shutdown{} }
func (m *StreamResponse_Shutdown) String() string            { return proto.CompactTextString(m) }
func (*StreamResponse_Shutdown) ProtoMessage()               {}
func (*StreamResponse_Shutdown) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 3} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "gochat.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "gochat.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "gochat.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "gochat.LogoutResponse")
	proto.RegisterType((*StreamRequest)(nil), "gochat.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "gochat.StreamResponse")
	proto.RegisterType((*StreamResponse_Login)(nil), "gochat.StreamResponse.Login")
	proto.RegisterType((*StreamResponse_Logout)(nil), "gochat.StreamResponse.Logout")
	proto.RegisterType((*StreamResponse_Message)(nil), "gochat.StreamResponse.Message")
	proto.RegisterType((*StreamResponse_Shutdown)(nil), "gochat.StreamResponse.Shutdown")
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
	// Send/Receive messages
	Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error)
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

func (c *chatClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Chat_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Chat_serviceDesc.Streams[0], c.cc, "/gochat.Chat/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatStreamClient{stream}
	return x, nil
}

type Chat_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type chatStreamClient struct {
	grpc.ClientStream
}

func (x *chatStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Chat service

type ChatServer interface {
	// Registers a user to the chat
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Unregisters a user to the chat
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// Send/Receive messages
	Stream(Chat_StreamServer) error
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

func _Chat_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Stream(&chatStreamServer{stream})
}

type Chat_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type chatStreamServer struct {
	grpc.ServerStream
}

func (x *chatStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Chat_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0xce, 0xd2, 0x40,
	0x14, 0x6d, 0x3f, 0x68, 0x81, 0x0b, 0x54, 0x33, 0x41, 0xd3, 0x8c, 0x28, 0xa6, 0x89, 0x09, 0x6e,
	0x8a, 0xc1, 0xc4, 0x9f, 0x85, 0x26, 0xa2, 0x89, 0x8d, 0xd1, 0x4d, 0x71, 0xe5, 0x86, 0x14, 0x18,
	0x0b, 0x91, 0x76, 0x6a, 0x67, 0x0a, 0xaf, 0xe6, 0x6b, 0xf9, 0x06, 0x5f, 0x98, 0x9f, 0x96, 0x92,
	0xb2, 0xe3, 0x5e, 0xce, 0x39, 0xf7, 0xde, 0x33, 0xa7, 0x00, 0x9b, 0x5d, 0xc4, 0xfd, 0x2c, 0xa7,
	0x9c, 0x22, 0x3b, 0xa6, 0xe7, 0x0a, 0x4f, 0x62, 0x4a, 0xe3, 0x03, 0x99, 0x89, 0xee, 0xba, 0xf8,
	0x3d, 0xe3, 0xfb, 0x84, 0x30, 0x1e, 0x25, 0x99, 0x04, 0x7a, 0x1f, 0x61, 0xf0, 0x9d, 0xc6, 0xfb,
	0x34, 0x24, 0x7f, 0x0b, 0xc2, 0x38, 0x42, 0xd0, 0x4e, 0xa3, 0x84, 0xb8, 0xe6, 0x73, 0x73, 0xda,
	0x0b, 0xc5, 0x6f, 0x84, 0xa1, 0x9b, 0x45, 0x8c, 0x9d, 0x68, 0xbe, 0x75, 0xef, 0x44, 0xbf, 0xac,
	0xbd, 0x17, 0x30, 0x54, 0x7c, 0x96, 0xd1, 0x94, 0x11, 0x34, 0x02, 0x8b, 0xd3, 0x3f, 0x24, 0x55,
	0x0a, 0xb2, 0x50, 0x30, 0x5a, 0x70, 0x3d, 0xa7, 0x19, 0xf6, 0x10, 0x1c, 0x0d, 0x93, 0x72, 0xde,
	0x4b, 0x18, 0x2e, 0x79, 0x4e, 0xa2, 0x44, 0x13, 0x5d, 0xe8, 0x24, 0x84, 0xb1, 0x28, 0xd6, 0x3b,
	0xea, 0xd2, 0xfb, 0xdf, 0x02, 0x47, 0x63, 0xd5, 0x32, 0xef, 0xa0, 0x57, 0x1e, 0x2c, 0xe0, 0xfd,
	0x39, 0xf6, 0xa5, 0x25, 0xbe, 0xb6, 0xc4, 0xff, 0xa9, 0x11, 0x61, 0x05, 0x46, 0x9f, 0x60, 0xb0,
	0x39, 0xec, 0x49, 0xca, 0x57, 0x87, 0xf3, 0x79, 0xe2, 0xee, 0xfe, 0x7c, 0xec, 0x4b, 0x5f, 0xfd,
	0xfa, 0x1c, 0x5f, 0x58, 0x10, 0x18, 0x61, 0x5f, 0x72, 0x44, 0x89, 0xbe, 0xc0, 0xb0, 0x92, 0xa0,
	0x05, 0x77, 0x5b, 0x42, 0xe3, 0xe9, 0x6d, 0x0d, 0x5a, 0xf0, 0xc0, 0x08, 0x07, 0xa5, 0x08, 0x2d,
	0x38, 0xfa, 0x0a, 0x8e, 0x52, 0xd1, 0x67, 0xb7, 0x85, 0xcc, 0xb3, 0x1b, 0x32, 0x3f, 0x24, 0x2a,
	0x30, 0x42, 0x35, 0x5d, 0x35, 0xd0, 0x37, 0x78, 0xc0, 0x48, 0x7e, 0x24, 0xf9, 0x8a, 0xed, 0x0a,
	0xbe, 0xa5, 0xa7, 0xd4, 0xb5, 0x84, 0xd2, 0xe4, 0x86, 0xd2, 0x52, 0xc1, 0x02, 0x23, 0x74, 0x24,
	0x53, 0x77, 0xf0, 0x13, 0xb0, 0xe4, 0x8d, 0x0d, 0x71, 0xc1, 0x63, 0xb0, 0xd5, 0xee, 0x4d, 0xff,
	0xbe, 0x85, 0x8e, 0xde, 0xa8, 0x29, 0x6b, 0x17, 0xcf, 0x7b, 0x57, 0x7b, 0x5e, 0x0c, 0xd0, 0xd5,
	0xf3, 0x17, 0x1d, 0xb0, 0xc8, 0x91, 0xa4, 0x7c, 0xfe, 0xcf, 0x84, 0xf6, 0xe7, 0x5d, 0xc4, 0xd1,
	0x1b, 0xbd, 0xd1, 0x48, 0x5f, 0x73, 0x19, 0x6b, 0xfc, 0xe8, 0xaa, 0xab, 0xd2, 0x65, 0xa0, 0xf7,
	0xe5, 0xb2, 0x97, 0x90, 0x2a, 0xa8, 0xf8, 0xf1, 0x75, 0xbb, 0xa4, 0x7e, 0x00, 0x5b, 0x3a, 0x56,
	0x51, 0x6b, 0x51, 0xad, 0xa8, 0x75, 0x63, 0x3d, 0x63, 0x6a, 0xbe, 0x32, 0x17, 0xdd, 0x5f, 0xea,
	0x23, 0x5d, 0xdb, 0x22, 0x8a, 0xaf, 0xef, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x40, 0x8d, 0xfb,
	0xc1, 0x03, 0x00, 0x00,
}

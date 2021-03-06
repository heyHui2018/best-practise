// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twf.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UserInfoRequest struct {
	Greet                string   `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoRequest) Reset()         { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()    {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c724263dde6e8dba, []int{0}
}

func (m *UserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoRequest.Unmarshal(m, b)
}
func (m *UserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoRequest.Marshal(b, m, deterministic)
}
func (m *UserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoRequest.Merge(m, src)
}
func (m *UserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UserInfoRequest.Size(m)
}
func (m *UserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoRequest proto.InternalMessageInfo

func (m *UserInfoRequest) GetGreet() string {
	if m != nil {
		return m.Greet
	}
	return ""
}

type UserInfoResponse struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfoResponse) Reset()         { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()    {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c724263dde6e8dba, []int{1}
}

func (m *UserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfoResponse.Unmarshal(m, b)
}
func (m *UserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfoResponse.Marshal(b, m, deterministic)
}
func (m *UserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfoResponse.Merge(m, src)
}
func (m *UserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_UserInfoResponse.Size(m)
}
func (m *UserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfoResponse proto.InternalMessageInfo

func (m *UserInfoResponse) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

type Message struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c724263dde6e8dba, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfoRequest)(nil), "test.UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "test.UserInfoResponse")
	proto.RegisterType((*Message)(nil), "test.Message")
}

func init() { proto.RegisterFile("twf.proto", fileDescriptor_c724263dde6e8dba) }

var fileDescriptor_c724263dde6e8dba = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x29, 0x4f, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e, 0x51, 0x52, 0xe7, 0xe2, 0x0f,
	0x2d, 0x4e, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12,
	0xe1, 0x62, 0x4d, 0x2f, 0x4a, 0x4d, 0x2d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70,
	0x94, 0x34, 0xb8, 0x04, 0x10, 0x0a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0x2a, 0x8b, 0x52,
	0x0b, 0x72, 0x2a, 0x61, 0x2a, 0xc1, 0x1c, 0x25, 0x43, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4,
	0xf4, 0x54, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0xa8, 0x3c, 0x98, 0x0d, 0x12, 0x4b,
	0x49, 0x2c, 0x49, 0x94, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x8d, 0x2c, 0xb9, 0x58,
	0x40, 0x86, 0x0b, 0x19, 0x72, 0x71, 0x3b, 0xe7, 0xe7, 0xe6, 0x96, 0xe6, 0x65, 0x26, 0x27, 0x96,
	0xa4, 0x0a, 0xf1, 0xea, 0x81, 0xdc, 0xa8, 0x07, 0x35, 0x4d, 0x0a, 0x95, 0xab, 0xc4, 0xa0, 0xc1,
	0x68, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0x35, 0x69,
	0x9b, 0xda, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Communicate(ctx context.Context, opts ...grpc.CallOption) (User_CommunicateClient, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (User_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_User_serviceDesc.Streams[0], "/test.User/Communicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &userCommunicateClient{stream}
	return x, nil
}

type User_CommunicateClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type userCommunicateClient struct {
	grpc.ClientStream
}

func (x *userCommunicateClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userCommunicateClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Communicate(User_CommunicateServer) error
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServer).Communicate(&userCommunicateServer{stream})
}

type User_CommunicateServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type userCommunicateServer struct {
	grpc.ServerStream
}

func (x *userCommunicateServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userCommunicateServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.User",
	HandlerType: (*UserServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Communicate",
			Handler:       _User_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "twf.proto",
}

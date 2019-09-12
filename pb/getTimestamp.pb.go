// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getTimestamp.proto

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

// The request message containing the user's name.
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7efaed384311bfa, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

// The response message containing the greetings
type GetReply struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7efaed384311bfa, []int{1}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetReply) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *GetReply) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "timestamp.GetRequest")
	proto.RegisterType((*GetReply)(nil), "timestamp.GetReply")
}

func init() { proto.RegisterFile("getTimestamp.proto", fileDescriptor_f7efaed384311bfa) }

var fileDescriptor_f7efaed384311bfa = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4f, 0x2d, 0x09,
	0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x2c, 0x81, 0x09, 0x28, 0xf1, 0x70, 0x71, 0xb9, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x28, 0xe5, 0x71, 0x71, 0x80, 0x79, 0x05, 0x39, 0x95, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x25,
	0x89, 0x25, 0xa5, 0xc5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50, 0x9e, 0x90, 0x04, 0x17,
	0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c,
	0x2b, 0x24, 0xc3, 0x85, 0x30, 0x58, 0x82, 0x19, 0xac, 0x09, 0x21, 0x20, 0x24, 0xc4, 0xc5, 0x92,
	0x92, 0x58, 0x92, 0x28, 0xc1, 0x02, 0xd6, 0x04, 0x66, 0x1b, 0x39, 0x73, 0x31, 0xbb, 0xa7, 0x96,
	0x08, 0xd9, 0x70, 0xf1, 0xb8, 0x23, 0xb9, 0x52, 0x48, 0x54, 0x0f, 0xae, 0x4d, 0x0f, 0xe1, 0x3a,
	0x29, 0x61, 0x74, 0xe1, 0x82, 0x9c, 0x4a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xa7, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb1, 0x12, 0xe9, 0xab, 0xea, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetClient is the client API for Get service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetClient interface {
	GetTimestamp(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type getClient struct {
	cc *grpc.ClientConn
}

func NewGetClient(cc *grpc.ClientConn) GetClient {
	return &getClient{cc}
}

func (c *getClient) GetTimestamp(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/timestamp.Get/GetTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetServer is the server API for Get service.
type GetServer interface {
	GetTimestamp(context.Context, *GetRequest) (*GetReply, error)
}

func RegisterGetServer(s *grpc.Server, srv GetServer) {
	s.RegisterService(&_Get_serviceDesc, srv)
}

func _Get_GetTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetServer).GetTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timestamp.Get/GetTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetServer).GetTimestamp(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Get_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timestamp.Get",
	HandlerType: (*GetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimestamp",
			Handler:    _Get_GetTimestamp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getTimestamp.proto",
}
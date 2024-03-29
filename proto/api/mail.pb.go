// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendMailReq struct {
	Accountid            string    `protobuf:"bytes,1,opt,name=accountid,proto3" json:"accountid,omitempty"`
	Body                 *MailBody `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Record               int32     `protobuf:"varint,3,opt,name=Record,proto3" json:"Record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SendMailReq) Reset()         { *m = SendMailReq{} }
func (m *SendMailReq) String() string { return proto.CompactTextString(m) }
func (*SendMailReq) ProtoMessage()    {}
func (*SendMailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cda5f053e74676b, []int{0}
}

func (m *SendMailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMailReq.Unmarshal(m, b)
}
func (m *SendMailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMailReq.Marshal(b, m, deterministic)
}
func (m *SendMailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMailReq.Merge(m, src)
}
func (m *SendMailReq) XXX_Size() int {
	return xxx_messageInfo_SendMailReq.Size(m)
}
func (m *SendMailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMailReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMailReq proto.InternalMessageInfo

func (m *SendMailReq) GetAccountid() string {
	if m != nil {
		return m.Accountid
	}
	return ""
}

func (m *SendMailReq) GetBody() *MailBody {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *SendMailReq) GetRecord() int32 {
	if m != nil {
		return m.Record
	}
	return 0
}

type SendMailRes struct {
	Errcode              int32    `protobuf:"varint,1,opt,name=errcode,proto3" json:"errcode,omitempty"`
	Record               int32    `protobuf:"varint,2,opt,name=Record,proto3" json:"Record,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMailRes) Reset()         { *m = SendMailRes{} }
func (m *SendMailRes) String() string { return proto.CompactTextString(m) }
func (*SendMailRes) ProtoMessage()    {}
func (*SendMailRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cda5f053e74676b, []int{1}
}

func (m *SendMailRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMailRes.Unmarshal(m, b)
}
func (m *SendMailRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMailRes.Marshal(b, m, deterministic)
}
func (m *SendMailRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMailRes.Merge(m, src)
}
func (m *SendMailRes) XXX_Size() int {
	return xxx_messageInfo_SendMailRes.Size(m)
}
func (m *SendMailRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMailRes.DiscardUnknown(m)
}

var xxx_messageInfo_SendMailRes proto.InternalMessageInfo

func (m *SendMailRes) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *SendMailRes) GetRecord() int32 {
	if m != nil {
		return m.Record
	}
	return 0
}

type MailBody struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Txt                  string   `protobuf:"bytes,2,opt,name=txt,proto3" json:"txt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MailBody) Reset()         { *m = MailBody{} }
func (m *MailBody) String() string { return proto.CompactTextString(m) }
func (*MailBody) ProtoMessage()    {}
func (*MailBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cda5f053e74676b, []int{2}
}

func (m *MailBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailBody.Unmarshal(m, b)
}
func (m *MailBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailBody.Marshal(b, m, deterministic)
}
func (m *MailBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailBody.Merge(m, src)
}
func (m *MailBody) XXX_Size() int {
	return xxx_messageInfo_MailBody.Size(m)
}
func (m *MailBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MailBody.DiscardUnknown(m)
}

var xxx_messageInfo_MailBody proto.InternalMessageInfo

func (m *MailBody) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MailBody) GetTxt() string {
	if m != nil {
		return m.Txt
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMailReq)(nil), "api.sendMailReq")
	proto.RegisterType((*SendMailRes)(nil), "api.sendMailRes")
	proto.RegisterType((*MailBody)(nil), "api.MailBody")
}

func init() { proto.RegisterFile("mail.proto", fileDescriptor_7cda5f053e74676b) }

var fileDescriptor_7cda5f053e74676b = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x4d, 0xff, 0x69, 0xa7, 0x08, 0x65, 0x10, 0x09, 0xe2, 0xa1, 0xf6, 0xd4, 0x83, 0xf4,
	0x50, 0xbd, 0x0b, 0xde, 0xbd, 0xe4, 0x1b, 0xa4, 0x49, 0x84, 0x40, 0x6d, 0x6a, 0x9a, 0x85, 0xed,
	0xb7, 0x5f, 0x92, 0x6d, 0xd9, 0xee, 0xde, 0xf2, 0x26, 0xcc, 0xfb, 0xbd, 0x37, 0x00, 0x7f, 0x5c,
	0x0f, 0xed, 0x64, 0x8d, 0x33, 0x18, 0xf3, 0x49, 0xd7, 0xbf, 0x50, 0xcc, 0x6a, 0x94, 0x3f, 0x5c,
	0x0f, 0x4c, 0xfd, 0xe3, 0x2b, 0xe4, 0x5c, 0x08, 0x73, 0x18, 0x9d, 0x96, 0x94, 0x54, 0xa4, 0xc9,
	0xd9, 0x65, 0x80, 0x6f, 0x90, 0xf4, 0x46, 0x2e, 0x34, 0xaa, 0x48, 0x53, 0x74, 0x8f, 0x2d, 0x9f,
	0x74, 0xeb, 0x37, 0xbf, 0x8d, 0x5c, 0x58, 0xf8, 0xc2, 0x67, 0xc8, 0x98, 0x12, 0xc6, 0x4a, 0x1a,
	0x57, 0xa4, 0x49, 0xd9, 0xaa, 0xea, 0xaf, 0x3d, 0x67, 0x46, 0x0a, 0xf7, 0xca, 0x5a, 0x61, 0xa4,
	0x0a, 0x94, 0x94, 0x6d, 0x72, 0x67, 0x10, 0x5d, 0x19, 0x74, 0xf0, 0xb0, 0xa1, 0xf0, 0x09, 0x52,
	0xa7, 0xdd, 0xa0, 0xd6, 0x84, 0x67, 0x81, 0x25, 0xc4, 0xee, 0xe8, 0xc2, 0x5a, 0xce, 0xfc, 0xb3,
	0xfb, 0x84, 0xc4, 0xf7, 0xc5, 0x77, 0x48, 0x3c, 0x1c, 0xcb, 0x90, 0x78, 0xd7, 0xf7, 0xe5, 0x76,
	0x32, 0xd7, 0x77, 0x7d, 0x16, 0xce, 0xf3, 0x71, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x81, 0x7a,
	0xdc, 0x2c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MailClient is the client API for Mail service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MailClient interface {
	Send(ctx context.Context, in *SendMailReq, opts ...grpc.CallOption) (*SendMailRes, error)
}

type mailClient struct {
	cc *grpc.ClientConn
}

func NewMailClient(cc *grpc.ClientConn) MailClient {
	return &mailClient{cc}
}

func (c *mailClient) Send(ctx context.Context, in *SendMailReq, opts ...grpc.CallOption) (*SendMailRes, error) {
	out := new(SendMailRes)
	err := c.cc.Invoke(ctx, "/api.mail/send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServer is the server API for Mail service.
type MailServer interface {
	Send(context.Context, *SendMailReq) (*SendMailRes, error)
}

// UnimplementedMailServer can be embedded to have forward compatible implementations.
type UnimplementedMailServer struct {
}

func (*UnimplementedMailServer) Send(ctx context.Context, req *SendMailReq) (*SendMailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterMailServer(s *grpc.Server, srv MailServer) {
	s.RegisterService(&_Mail_serviceDesc, srv)
}

func _Mail_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.mail/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServer).Send(ctx, req.(*SendMailReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.mail",
	HandlerType: (*MailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "send",
			Handler:    _Mail_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mail.proto",
}

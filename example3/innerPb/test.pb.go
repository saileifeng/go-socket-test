// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package innerPb

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

//聊天消息状态
type ChatMessageStatus int32

const (
	ChatMessageStatus_Reached ChatMessageStatus = 0
	ChatMessageStatus_Looked  ChatMessageStatus = 1
)

var ChatMessageStatus_name = map[int32]string{
	0: "Reached",
	1: "Looked",
}

var ChatMessageStatus_value = map[string]int32{
	"Reached": 0,
	"Looked":  1,
}

func (x ChatMessageStatus) String() string {
	return proto.EnumName(ChatMessageStatus_name, int32(x))
}

func (ChatMessageStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

type Pid struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pid) Reset()         { *m = Pid{} }
func (m *Pid) String() string { return proto.CompactTextString(m) }
func (*Pid) ProtoMessage()    {}
func (*Pid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Pid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pid.Unmarshal(m, b)
}
func (m *Pid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pid.Marshal(b, m, deterministic)
}
func (m *Pid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pid.Merge(m, src)
}
func (m *Pid) XXX_Size() int {
	return xxx_messageInfo_Pid.Size(m)
}
func (m *Pid) XXX_DiscardUnknown() {
	xxx_messageInfo_Pid.DiscardUnknown(m)
}

var xxx_messageInfo_Pid proto.InternalMessageInfo

func (m *Pid) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Pid) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RPC_ChatToSomeOneMessageRequest struct {
	SenderPID            *Pid     `protobuf:"bytes,1,opt,name=SenderPID,proto3" json:"SenderPID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	MessageID            int64    `protobuf:"varint,3,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	SenderUID            string   `protobuf:"bytes,4,opt,name=SenderUID,proto3" json:"SenderUID,omitempty"`
	TargetUID            string   `protobuf:"bytes,5,opt,name=TargetUID,proto3" json:"TargetUID,omitempty"`
	TargetPID            *Pid     `protobuf:"bytes,6,opt,name=TargetPID,proto3" json:"TargetPID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPC_ChatToSomeOneMessageRequest) Reset()         { *m = RPC_ChatToSomeOneMessageRequest{} }
func (m *RPC_ChatToSomeOneMessageRequest) String() string { return proto.CompactTextString(m) }
func (*RPC_ChatToSomeOneMessageRequest) ProtoMessage()    {}
func (*RPC_ChatToSomeOneMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *RPC_ChatToSomeOneMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPC_ChatToSomeOneMessageRequest.Unmarshal(m, b)
}
func (m *RPC_ChatToSomeOneMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPC_ChatToSomeOneMessageRequest.Marshal(b, m, deterministic)
}
func (m *RPC_ChatToSomeOneMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC_ChatToSomeOneMessageRequest.Merge(m, src)
}
func (m *RPC_ChatToSomeOneMessageRequest) XXX_Size() int {
	return xxx_messageInfo_RPC_ChatToSomeOneMessageRequest.Size(m)
}
func (m *RPC_ChatToSomeOneMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC_ChatToSomeOneMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RPC_ChatToSomeOneMessageRequest proto.InternalMessageInfo

func (m *RPC_ChatToSomeOneMessageRequest) GetSenderPID() *Pid {
	if m != nil {
		return m.SenderPID
	}
	return nil
}

func (m *RPC_ChatToSomeOneMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RPC_ChatToSomeOneMessageRequest) GetMessageID() int64 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func (m *RPC_ChatToSomeOneMessageRequest) GetSenderUID() string {
	if m != nil {
		return m.SenderUID
	}
	return ""
}

func (m *RPC_ChatToSomeOneMessageRequest) GetTargetUID() string {
	if m != nil {
		return m.TargetUID
	}
	return ""
}

func (m *RPC_ChatToSomeOneMessageRequest) GetTargetPID() *Pid {
	if m != nil {
		return m.TargetPID
	}
	return nil
}

type RPC_ChatToSomeOneMessageResponse struct {
	SenderPID            *Pid              `protobuf:"bytes,1,opt,name=SenderPID,proto3" json:"SenderPID,omitempty"`
	ChatMessageStatus    ChatMessageStatus `protobuf:"varint,2,opt,name=ChatMessageStatus,proto3,enum=innerPb.ChatMessageStatus" json:"ChatMessageStatus,omitempty"`
	MessageID            int64             `protobuf:"varint,3,opt,name=MessageID,proto3" json:"MessageID,omitempty"`
	SenderUID            string            `protobuf:"bytes,4,opt,name=SenderUID,proto3" json:"SenderUID,omitempty"`
	TargetUID            string            `protobuf:"bytes,5,opt,name=TargetUID,proto3" json:"TargetUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RPC_ChatToSomeOneMessageResponse) Reset()         { *m = RPC_ChatToSomeOneMessageResponse{} }
func (m *RPC_ChatToSomeOneMessageResponse) String() string { return proto.CompactTextString(m) }
func (*RPC_ChatToSomeOneMessageResponse) ProtoMessage()    {}
func (*RPC_ChatToSomeOneMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *RPC_ChatToSomeOneMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPC_ChatToSomeOneMessageResponse.Unmarshal(m, b)
}
func (m *RPC_ChatToSomeOneMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPC_ChatToSomeOneMessageResponse.Marshal(b, m, deterministic)
}
func (m *RPC_ChatToSomeOneMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPC_ChatToSomeOneMessageResponse.Merge(m, src)
}
func (m *RPC_ChatToSomeOneMessageResponse) XXX_Size() int {
	return xxx_messageInfo_RPC_ChatToSomeOneMessageResponse.Size(m)
}
func (m *RPC_ChatToSomeOneMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RPC_ChatToSomeOneMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RPC_ChatToSomeOneMessageResponse proto.InternalMessageInfo

func (m *RPC_ChatToSomeOneMessageResponse) GetSenderPID() *Pid {
	if m != nil {
		return m.SenderPID
	}
	return nil
}

func (m *RPC_ChatToSomeOneMessageResponse) GetChatMessageStatus() ChatMessageStatus {
	if m != nil {
		return m.ChatMessageStatus
	}
	return ChatMessageStatus_Reached
}

func (m *RPC_ChatToSomeOneMessageResponse) GetMessageID() int64 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func (m *RPC_ChatToSomeOneMessageResponse) GetSenderUID() string {
	if m != nil {
		return m.SenderUID
	}
	return ""
}

func (m *RPC_ChatToSomeOneMessageResponse) GetTargetUID() string {
	if m != nil {
		return m.TargetUID
	}
	return ""
}

func init() {
	proto.RegisterEnum("innerPb.ChatMessageStatus", ChatMessageStatus_name, ChatMessageStatus_value)
	proto.RegisterType((*Pid)(nil), "innerPb.Pid")
	proto.RegisterType((*RPC_ChatToSomeOneMessageRequest)(nil), "innerPb.RPC_ChatToSomeOneMessageRequest")
	proto.RegisterType((*RPC_ChatToSomeOneMessageResponse)(nil), "innerPb.RPC_ChatToSomeOneMessageResponse")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x35, 0xdb, 0xdc, 0xd8, 0x57, 0x1d, 0x33, 0xa7, 0x32, 0x04, 0x4b, 0x4f, 0xdd, 0x90, 0x1e,
	0xea, 0x2f, 0x90, 0xf5, 0x60, 0x41, 0xb1, 0xa4, 0xf3, 0x2c, 0xdd, 0xf2, 0xb1, 0x15, 0xb1, 0x99,
	0x49, 0xe6, 0xbf, 0xf6, 0xe2, 0x2f, 0x90, 0xa4, 0xb5, 0x45, 0x36, 0x51, 0x0f, 0xde, 0x92, 0xf7,
	0xbe, 0xf7, 0xf2, 0xde, 0x47, 0x00, 0x34, 0x2a, 0x1d, 0x6e, 0xa5, 0xd0, 0x82, 0x0e, 0x8a, 0xb2,
	0x44, 0x99, 0x2e, 0xfd, 0x29, 0x74, 0xd3, 0x82, 0x53, 0x0a, 0xbd, 0x6b, 0xce, 0xa5, 0x4b, 0x3c,
	0x12, 0x0c, 0x99, 0x3d, 0xd3, 0x11, 0x74, 0x12, 0xee, 0x76, 0x2c, 0xd2, 0x49, 0xb8, 0xff, 0x46,
	0xe0, 0x82, 0xa5, 0xf3, 0xc7, 0xf9, 0x26, 0xd7, 0x0b, 0x91, 0x89, 0x67, 0xbc, 0x2f, 0xf1, 0x0e,
	0x95, 0xca, 0xd7, 0xc8, 0xf0, 0x65, 0x87, 0x4a, 0xd3, 0x19, 0x0c, 0x33, 0x2c, 0x39, 0xca, 0x34,
	0x89, 0xad, 0x99, 0x13, 0x9d, 0x84, 0xf5, 0x5b, 0x61, 0x5a, 0x70, 0xd6, 0xd2, 0xd4, 0x85, 0x41,
	0xad, 0xae, 0x1f, 0xf9, 0xbc, 0xd2, 0x73, 0x18, 0xd6, 0xc7, 0x24, 0x76, 0xbb, 0x1e, 0x09, 0xba,
	0xac, 0x05, 0x0c, 0x5b, 0x99, 0x3c, 0x24, 0xb1, 0xdb, 0xb3, 0xca, 0x16, 0x30, 0xec, 0x22, 0x97,
	0x6b, 0xd4, 0x86, 0x3d, 0xae, 0xd8, 0x06, 0x30, 0xf9, 0xaa, 0x8b, 0xc9, 0xd7, 0x3f, 0x94, 0xaf,
	0xa1, 0xfd, 0x77, 0x02, 0xde, 0xf7, 0x7d, 0xd5, 0x56, 0x94, 0x0a, 0xff, 0x54, 0xf8, 0x06, 0xce,
	0x8c, 0x57, 0x6d, 0x91, 0xe9, 0x5c, 0xef, 0x94, 0xad, 0x3e, 0x8a, 0x26, 0x8d, 0x66, 0x6f, 0x82,
	0xed, 0x8b, 0xfe, 0x6f, 0x41, 0xb3, 0xcb, 0x03, 0x19, 0xa9, 0x03, 0x03, 0x86, 0xf9, 0x6a, 0x83,
	0x7c, 0x7c, 0x44, 0x01, 0xfa, 0xb7, 0x42, 0x3c, 0x21, 0x1f, 0x93, 0x48, 0x81, 0x63, 0xa6, 0x33,
	0x94, 0xaf, 0xc5, 0x0a, 0x29, 0x87, 0xd3, 0x2f, 0xcb, 0xa2, 0x41, 0x53, 0xeb, 0x87, 0x8f, 0x33,
	0x99, 0xfe, 0x62, 0xb2, 0x5a, 0xf9, 0xb2, 0x6f, 0xbf, 0xf0, 0xd5, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x05, 0x3d, 0x69, 0x1c, 0xd0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	ChatToSomeOne(ctx context.Context, in *RPC_ChatToSomeOneMessageRequest, opts ...grpc.CallOption) (*RPC_ChatToSomeOneMessageResponse, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ChatToSomeOne(ctx context.Context, in *RPC_ChatToSomeOneMessageRequest, opts ...grpc.CallOption) (*RPC_ChatToSomeOneMessageResponse, error) {
	out := new(RPC_ChatToSomeOneMessageResponse)
	err := c.cc.Invoke(ctx, "/innerPb.ChatService/ChatToSomeOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	ChatToSomeOne(context.Context, *RPC_ChatToSomeOneMessageRequest) (*RPC_ChatToSomeOneMessageResponse, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) ChatToSomeOne(ctx context.Context, req *RPC_ChatToSomeOneMessageRequest) (*RPC_ChatToSomeOneMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChatToSomeOne not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_ChatToSomeOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPC_ChatToSomeOneMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ChatToSomeOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/innerPb.ChatService/ChatToSomeOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ChatToSomeOne(ctx, req.(*RPC_ChatToSomeOneMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "innerPb.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChatToSomeOne",
			Handler:    _ChatService_ChatToSomeOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
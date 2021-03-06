// Code generated by protoc-gen-go.
// source: interact.proto
// DO NOT EDIT!

/*
Package interact is a generated protocol buffer package.

It is generated from these files:
	interact.proto

It has these top-level messages:
	ServerStatus
	ClientStatus
	Token
	QUIZ
	Session
	Req
	Question
	Answer
	Response
	AnswerStatus
*/
package interact

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

type QUIZState int32

const (
	QUIZ_DEMO_NOT_TAKEN QUIZState = 0
	QUIZ_DEMO_STARTED   QUIZState = 1
	QUIZ_TEST_NOT_TAKEN QUIZState = 2
	QUIZ_TEST_STARTED   QUIZState = 3
	QUIZ_TEST_FINISHED  QUIZState = 4
)

var QUIZState_name = map[int32]string{
	0: "DEMO_NOT_TAKEN",
	1: "DEMO_STARTED",
	2: "TEST_NOT_TAKEN",
	3: "TEST_STARTED",
	4: "TEST_FINISHED",
}
var QUIZState_value = map[string]int32{
	"DEMO_NOT_TAKEN": 0,
	"DEMO_STARTED":   1,
	"TEST_NOT_TAKEN": 2,
	"TEST_STARTED":   3,
	"TEST_FINISHED":  4,
}

func (x QUIZState) String() string {
	return proto.EnumName(QUIZState_name, int32(x))
}
func (QUIZState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type ServerStatus struct {
	TimeLeft string `protobuf:"bytes,1,opt,name=timeLeft" json:"timeLeft,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *ServerStatus) Reset()                    { *m = ServerStatus{} }
func (m *ServerStatus) String() string            { return proto.CompactTextString(m) }
func (*ServerStatus) ProtoMessage()               {}
func (*ServerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClientStatus struct {
	CurQuestion string `protobuf:"bytes,1,opt,name=curQuestion" json:"curQuestion,omitempty"`
	Token       string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *ClientStatus) Reset()                    { *m = ClientStatus{} }
func (m *ClientStatus) String() string            { return proto.CompactTextString(m) }
func (*ClientStatus) ProtoMessage()               {}
func (*ClientStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Token struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type QUIZ struct {
}

func (m *QUIZ) Reset()                    { *m = QUIZ{} }
func (m *QUIZ) String() string            { return proto.CompactTextString(m) }
func (*QUIZ) ProtoMessage()               {}
func (*QUIZ) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Session struct {
	Id           string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	State        QUIZState `protobuf:"varint,2,opt,name=state,enum=interact.QUIZState" json:"state,omitempty"`
	TimeLeft     string    `protobuf:"bytes,3,opt,name=timeLeft" json:"timeLeft,omitempty"`
	TestDuration string    `protobuf:"bytes,4,opt,name=testDuration" json:"testDuration,omitempty"`
	DemoDuration string    `protobuf:"bytes,5,opt,name=demoDuration" json:"demoDuration,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Req struct {
	Repeat bool   `protobuf:"varint,1,opt,name=repeat" json:"repeat,omitempty"`
	Sid    string `protobuf:"bytes,2,opt,name=sid" json:"sid,omitempty"`
	Token  string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Question struct {
	Id         string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Str        string    `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
	Options    []*Answer `protobuf:"bytes,3,rep,name=options" json:"options,omitempty"`
	IsMultiple bool      `protobuf:"varint,4,opt,name=isMultiple" json:"isMultiple,omitempty"`
	Positive   float32   `protobuf:"fixed32,6,opt,name=positive" json:"positive,omitempty"`
	Negative   float32   `protobuf:"fixed32,7,opt,name=negative" json:"negative,omitempty"`
	Score      float32   `protobuf:"fixed32,8,opt,name=score" json:"score,omitempty"`
}

func (m *Question) Reset()                    { *m = Question{} }
func (m *Question) String() string            { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()               {}
func (*Question) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Question) GetOptions() []*Answer {
	if m != nil {
		return m.Options
	}
	return nil
}

type Answer struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Str string `protobuf:"bytes,2,opt,name=str" json:"str,omitempty"`
}

func (m *Answer) Reset()                    { *m = Answer{} }
func (m *Answer) String() string            { return proto.CompactTextString(m) }
func (*Answer) ProtoMessage()               {}
func (*Answer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Response struct {
	Qid   string   `protobuf:"bytes,1,opt,name=qid" json:"qid,omitempty"`
	Aid   []string `protobuf:"bytes,2,rep,name=aid" json:"aid,omitempty"`
	Sid   string   `protobuf:"bytes,3,opt,name=sid" json:"sid,omitempty"`
	Token string   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AnswerStatus struct {
	Status int64 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *AnswerStatus) Reset()                    { *m = AnswerStatus{} }
func (m *AnswerStatus) String() string            { return proto.CompactTextString(m) }
func (*AnswerStatus) ProtoMessage()               {}
func (*AnswerStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func init() {
	proto.RegisterType((*ServerStatus)(nil), "interact.ServerStatus")
	proto.RegisterType((*ClientStatus)(nil), "interact.ClientStatus")
	proto.RegisterType((*Token)(nil), "interact.Token")
	proto.RegisterType((*QUIZ)(nil), "interact.QUIZ")
	proto.RegisterType((*Session)(nil), "interact.Session")
	proto.RegisterType((*Req)(nil), "interact.Req")
	proto.RegisterType((*Question)(nil), "interact.Question")
	proto.RegisterType((*Answer)(nil), "interact.Answer")
	proto.RegisterType((*Response)(nil), "interact.Response")
	proto.RegisterType((*AnswerStatus)(nil), "interact.AnswerStatus")
	proto.RegisterEnum("interact.QUIZState", QUIZState_name, QUIZState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for GruQuiz service

type GruQuizClient interface {
	Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error)
	GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error)
	Status(ctx context.Context, in *Response, opts ...grpc.CallOption) (*AnswerStatus, error)
	Ping(ctx context.Context, in *ClientStatus, opts ...grpc.CallOption) (*ServerStatus, error)
}

type gruQuizClient struct {
	cc *grpc.ClientConn
}

func NewGruQuizClient(cc *grpc.ClientConn) GruQuizClient {
	return &gruQuizClient{cc}
}

func (c *gruQuizClient) Authenticate(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) GetQuestion(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Question, error) {
	out := new(Question)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/GetQuestion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) Status(ctx context.Context, in *Response, opts ...grpc.CallOption) (*AnswerStatus, error) {
	out := new(AnswerStatus)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gruQuizClient) Ping(ctx context.Context, in *ClientStatus, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := grpc.Invoke(ctx, "/interact.GruQuiz/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GruQuiz service

type GruQuizServer interface {
	Authenticate(context.Context, *Token) (*Session, error)
	GetQuestion(context.Context, *Req) (*Question, error)
	Status(context.Context, *Response) (*AnswerStatus, error)
	Ping(context.Context, *ClientStatus) (*ServerStatus, error)
}

func RegisterGruQuizServer(s *grpc.Server, srv GruQuizServer) {
	s.RegisterService(&_GruQuiz_serviceDesc, srv)
}

func _GruQuiz_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).Authenticate(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/GetQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).GetQuestion(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).Status(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

func _GruQuiz_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GruQuizServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/interact.GruQuiz/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GruQuizServer).Ping(ctx, req.(*ClientStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _GruQuiz_serviceDesc = grpc.ServiceDesc{
	ServiceName: "interact.GruQuiz",
	HandlerType: (*GruQuizServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _GruQuiz_Authenticate_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _GruQuiz_GetQuestion_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _GruQuiz_Status_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _GruQuiz_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("interact.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x6e, 0xe2, 0xfc, 0x98, 0x49, 0x1a, 0xd2, 0x55, 0x55, 0xa2, 0x1c, 0x50, 0xb5, 0x07, 0x54,
	0xf5, 0x90, 0x43, 0x40, 0x88, 0x6b, 0x20, 0x69, 0x89, 0xa0, 0x29, 0xb1, 0x0d, 0x07, 0x2e, 0x95,
	0x49, 0xa7, 0x65, 0x45, 0x6a, 0xbb, 0xbb, 0xeb, 0x22, 0xf1, 0x46, 0xbc, 0x08, 0x6f, 0xc2, 0x7b,
	0xb0, 0xbb, 0x5e, 0xdb, 0x4b, 0x84, 0xc4, 0x6d, 0xe7, 0x9b, 0x6f, 0xc7, 0xdf, 0x7e, 0x33, 0x63,
	0x18, 0xb0, 0x44, 0x22, 0x8f, 0x37, 0x72, 0x92, 0xf1, 0x54, 0xa6, 0xc4, 0x2f, 0x63, 0xfa, 0x1a,
	0xfa, 0x21, 0xf2, 0x07, 0xe4, 0xa1, 0x8c, 0x65, 0x2e, 0xc8, 0x18, 0x7c, 0xc9, 0xee, 0xf0, 0x3d,
	0xde, 0xc8, 0x51, 0xe3, 0xb8, 0x71, 0xf2, 0x28, 0xa8, 0x62, 0x72, 0x04, 0x1d, 0x61, 0x58, 0xa3,
	0xa6, 0xc9, 0xd8, 0x88, 0x9e, 0x41, 0xff, 0xcd, 0x96, 0x61, 0x22, 0x6d, 0x8d, 0x63, 0xe8, 0x6d,
	0x72, 0xbe, 0xce, 0x51, 0x48, 0x96, 0x26, 0xb6, 0x8c, 0x0b, 0x91, 0x43, 0x68, 0xcb, 0xf4, 0x1b,
	0x26, 0xb6, 0x50, 0x11, 0xd0, 0x27, 0xd0, 0x8e, 0xf4, 0x81, 0x0c, 0xa0, 0xc9, 0xae, 0xed, 0x3d,
	0x75, 0xa2, 0x09, 0xb4, 0xd6, 0x1f, 0x97, 0x9f, 0xe9, 0x0d, 0xb4, 0xf5, 0x27, 0x91, 0x10, 0x18,
	0xcc, 0x17, 0x17, 0x97, 0x57, 0xab, 0xcb, 0xe8, 0x2a, 0x9a, 0xbd, 0x5b, 0xac, 0x86, 0x7b, 0x64,
	0x08, 0x7d, 0x83, 0x85, 0xd1, 0x2c, 0x88, 0x16, 0xf3, 0x61, 0x43, 0xb3, 0xa2, 0x45, 0x18, 0x39,
	0xac, 0xa6, 0x66, 0x19, 0xac, 0x64, 0x79, 0xe4, 0x00, 0xf6, 0x0d, 0x72, 0xb6, 0x5c, 0x2d, 0xc3,
	0xb7, 0x0a, 0x6a, 0xd1, 0x9f, 0x0d, 0xe8, 0x86, 0x28, 0x84, 0x96, 0xba, 0xa3, 0x85, 0x9c, 0x5a,
	0x0d, 0x46, 0xfa, 0x60, 0x7a, 0x38, 0xa9, 0xac, 0xd5, 0x12, 0x27, 0x26, 0x17, 0x58, 0x99, 0xae,
	0x99, 0xde, 0x8e, 0x99, 0x14, 0xfa, 0x52, 0xb9, 0x31, 0xcf, 0x79, 0x6c, 0x5c, 0x6a, 0x99, 0xfc,
	0x5f, 0x98, 0xe6, 0x5c, 0xe3, 0x5d, 0x5a, 0x71, 0xda, 0x05, 0xc7, 0xc5, 0xe8, 0x02, 0xbc, 0x00,
	0xef, 0x75, 0x6f, 0x38, 0x66, 0x18, 0x17, 0x5d, 0xf3, 0x03, 0x1b, 0xa9, 0xf7, 0x7a, 0x42, 0xe9,
	0x2f, 0x7c, 0xd6, 0xc7, 0xda, 0x7b, 0xcf, 0xf5, 0xfe, 0x57, 0x03, 0xfc, 0xaa, 0x3d, 0xbb, 0x6f,
	0xd6, 0x45, 0x24, 0xaf, 0x8a, 0x48, 0xae, 0x5c, 0xe8, 0xa6, 0x99, 0xe6, 0x0a, 0x55, 0xc6, 0x3b,
	0xe9, 0x4d, 0x87, 0xb5, 0x0f, 0xb3, 0x44, 0x7c, 0x47, 0x1e, 0x94, 0x04, 0xf2, 0x14, 0x80, 0x89,
	0x8b, 0x7c, 0x2b, 0x59, 0xb6, 0x45, 0xf3, 0x4e, 0x3f, 0x70, 0x10, 0xed, 0x52, 0x96, 0x0a, 0x26,
	0xd9, 0x03, 0x8e, 0x3a, 0x2a, 0xdb, 0x0c, 0xaa, 0x58, 0xe7, 0x12, 0xbc, 0x8d, 0x4d, 0xae, 0x5b,
	0xe4, 0xca, 0x58, 0x3f, 0x44, 0x6c, 0x52, 0x8e, 0x23, 0xdf, 0x24, 0x8a, 0x80, 0x9e, 0x42, 0xa7,
	0x10, 0xf0, 0xff, 0x57, 0xd0, 0x4f, 0xe0, 0x07, 0x28, 0x32, 0x25, 0x12, 0x75, 0xf6, 0xbe, 0xa2,
	0xeb, 0xa3, 0x46, 0x62, 0x63, 0x9d, 0xa7, 0x91, 0xd8, 0x56, 0x50, 0x88, 0xf7, 0x0f, 0x33, 0x5b,
	0xae, 0x99, 0xcf, 0xa0, 0x5f, 0x68, 0xb0, 0x0b, 0x51, 0x2f, 0x8e, 0x2e, 0xef, 0x95, 0x8b, 0x33,
	0xfd, 0xad, 0xe6, 0xec, 0x9c, 0xe7, 0xeb, 0x9c, 0xfd, 0x20, 0x2f, 0xd4, 0x9d, 0x5c, 0x7e, 0x55,
	0x5b, 0xc4, 0x36, 0x7a, 0x76, 0x1e, 0xd7, 0x86, 0x9a, 0xa5, 0x18, 0x1f, 0xd4, 0x80, 0x9d, 0x4d,
	0xba, 0x47, 0xa6, 0xd0, 0x3b, 0x47, 0x59, 0x35, 0x6e, 0xbf, 0xe6, 0xa8, 0xa1, 0x18, 0x13, 0x67,
	0x38, 0x2d, 0x45, 0xdd, 0x79, 0x09, 0x1d, 0xab, 0x8b, 0xb8, 0xf4, 0xc2, 0x87, 0xf1, 0xd1, 0x6e,
	0x23, 0x0b, 0xae, 0xba, 0xf7, 0x0a, 0x5a, 0x1f, 0x58, 0x72, 0x4b, 0x1c, 0x86, 0xbb, 0xf6, 0xee,
	0x4d, 0xf7, 0x97, 0x42, 0xf7, 0xbe, 0x74, 0xcc, 0x5f, 0xe7, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x60, 0xdf, 0x7c, 0x61, 0x87, 0x04, 0x00, 0x00,
}

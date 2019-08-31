// Code generated by protoc-gen-go. DO NOT EDIT.
// source: timingsystem.proto

//go:generate protoc -I .\sysprotos\ --go_out=plugins=grpc:.\sysprotos .\sysprotos\timingsystem.proto

package timingsystem

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

type TimingSystemRequest_FinishTypes int32

const (
	TimingSystemRequest_FINISH_CORRIDOR TimingSystemRequest_FinishTypes = 0
	TimingSystemRequest_FINISH_LINE     TimingSystemRequest_FinishTypes = 1
)

var TimingSystemRequest_FinishTypes_name = map[int32]string{
	0: "FINISH_CORRIDOR",
	1: "FINISH_LINE",
}

var TimingSystemRequest_FinishTypes_value = map[string]int32{
	"FINISH_CORRIDOR": 0,
	"FINISH_LINE":     1,
}

func (x TimingSystemRequest_FinishTypes) String() string {
	return proto.EnumName(TimingSystemRequest_FinishTypes_name, int32(x))
}

func (TimingSystemRequest_FinishTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a6ffa42cb48888d, []int{0, 0}
}

type TimingSystemRequest struct {
	Id                   int32                           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 TimingSystemRequest_FinishTypes `protobuf:"varint,2,opt,name=type,proto3,enum=timingsystem.TimingSystemRequest_FinishTypes" json:"type,omitempty"`
	TimePoint            string                          `protobuf:"bytes,3,opt,name=timePoint,proto3" json:"timePoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *TimingSystemRequest) Reset()         { *m = TimingSystemRequest{} }
func (m *TimingSystemRequest) String() string { return proto.CompactTextString(m) }
func (*TimingSystemRequest) ProtoMessage()    {}
func (*TimingSystemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6ffa42cb48888d, []int{0}
}

func (m *TimingSystemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimingSystemRequest.Unmarshal(m, b)
}
func (m *TimingSystemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimingSystemRequest.Marshal(b, m, deterministic)
}
func (m *TimingSystemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimingSystemRequest.Merge(m, src)
}
func (m *TimingSystemRequest) XXX_Size() int {
	return xxx_messageInfo_TimingSystemRequest.Size(m)
}
func (m *TimingSystemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimingSystemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimingSystemRequest proto.InternalMessageInfo

func (m *TimingSystemRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TimingSystemRequest) GetType() TimingSystemRequest_FinishTypes {
	if m != nil {
		return m.Type
	}
	return TimingSystemRequest_FINISH_CORRIDOR
}

func (m *TimingSystemRequest) GetTimePoint() string {
	if m != nil {
		return m.TimePoint
	}
	return ""
}

type TimingSystemResponse struct {
	ResultStatus         bool     `protobuf:"varint,1,opt,name=resultStatus,proto3" json:"resultStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimingSystemResponse) Reset()         { *m = TimingSystemResponse{} }
func (m *TimingSystemResponse) String() string { return proto.CompactTextString(m) }
func (*TimingSystemResponse) ProtoMessage()    {}
func (*TimingSystemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6ffa42cb48888d, []int{1}
}

func (m *TimingSystemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimingSystemResponse.Unmarshal(m, b)
}
func (m *TimingSystemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimingSystemResponse.Marshal(b, m, deterministic)
}
func (m *TimingSystemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimingSystemResponse.Merge(m, src)
}
func (m *TimingSystemResponse) XXX_Size() int {
	return xxx_messageInfo_TimingSystemResponse.Size(m)
}
func (m *TimingSystemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TimingSystemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TimingSystemResponse proto.InternalMessageInfo

func (m *TimingSystemResponse) GetResultStatus() bool {
	if m != nil {
		return m.ResultStatus
	}
	return false
}

func init() {
	proto.RegisterEnum("timingsystem.TimingSystemRequest_FinishTypes", TimingSystemRequest_FinishTypes_name, TimingSystemRequest_FinishTypes_value)
	proto.RegisterType((*TimingSystemRequest)(nil), "timingsystem.TimingSystemRequest")
	proto.RegisterType((*TimingSystemResponse)(nil), "timingsystem.TimingSystemResponse")
}

func init() { proto.RegisterFile("timingsystem.proto", fileDescriptor_3a6ffa42cb48888d) }

var fileDescriptor_3a6ffa42cb48888d = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x06, 0xe0, 0x6e, 0xfc, 0xc0, 0x4e, 0x43, 0x5b, 0xa7, 0x1e, 0x82, 0x78, 0x88, 0x7b, 0xca,
	0xc5, 0x1c, 0xda, 0x9b, 0x37, 0x51, 0x8b, 0x01, 0x69, 0x65, 0xd3, 0xa3, 0x20, 0x6a, 0x06, 0x5d,
	0x68, 0x3e, 0xcc, 0x4c, 0x0e, 0xf9, 0x6d, 0xfe, 0x39, 0x61, 0x2b, 0x98, 0x80, 0xe8, 0xf5, 0x61,
	0xdf, 0xd9, 0x79, 0x07, 0x50, 0x6c, 0x6e, 0x8b, 0x37, 0x6e, 0x59, 0x28, 0x8f, 0xab, 0xba, 0x94,
	0x12, 0xfd, 0xae, 0xe9, 0x4f, 0x05, 0xb3, 0x8d, 0x83, 0xd4, 0x81, 0xa1, 0x8f, 0x86, 0x58, 0x70,
	0x0c, 0x9e, 0xcd, 0x02, 0x15, 0xaa, 0xe8, 0xc0, 0x78, 0x36, 0xc3, 0x2b, 0xd8, 0x97, 0xb6, 0xa2,
	0xc0, 0x0b, 0x55, 0x34, 0x9e, 0x5f, 0xc4, 0xbd, 0xc1, 0xbf, 0x0c, 0x88, 0x97, 0xb6, 0xb0, 0xfc,
	0xbe, 0x69, 0x2b, 0x62, 0xe3, 0xa2, 0x78, 0x06, 0x43, 0xb1, 0x39, 0x3d, 0x94, 0xb6, 0x90, 0x60,
	0x2f, 0x54, 0xd1, 0xd0, 0xfc, 0x80, 0x5e, 0xc0, 0xa8, 0x13, 0xc1, 0x19, 0x4c, 0x96, 0xc9, 0x2a,
	0x49, 0xef, 0x9e, 0xae, 0xd7, 0xc6, 0x24, 0x37, 0x6b, 0x33, 0x1d, 0xe0, 0x04, 0x46, 0xdf, 0x78,
	0x9f, 0xac, 0x6e, 0xa7, 0x4a, 0x5f, 0xc2, 0x49, 0xff, 0x6f, 0xae, 0xca, 0x82, 0x09, 0x35, 0xf8,
	0x35, 0x71, 0xb3, 0x95, 0x54, 0x9e, 0xa5, 0x61, 0xd7, 0xe3, 0xc8, 0xf4, 0x6c, 0xbe, 0x05, 0xbf,
	0x9b, 0xc5, 0x47, 0x38, 0x36, 0xf4, 0x5a, 0xd6, 0xd9, 0x4e, 0xdd, 0x56, 0x78, 0xfe, 0x6f, 0xd1,
	0x53, 0xfd, 0xd7, 0x93, 0xdd, 0x3e, 0x7a, 0xf0, 0x72, 0xe8, 0x8e, 0xbf, 0xf8, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xef, 0x01, 0xa7, 0xa8, 0x92, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimingSystemClient is the client API for TimingSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimingSystemClient interface {
	RecordTimingPoint(ctx context.Context, in *TimingSystemRequest, opts ...grpc.CallOption) (*TimingSystemResponse, error)
}

type timingSystemClient struct {
	cc *grpc.ClientConn
}

func NewTimingSystemClient(cc *grpc.ClientConn) TimingSystemClient {
	return &timingSystemClient{cc}
}

func (c *timingSystemClient) RecordTimingPoint(ctx context.Context, in *TimingSystemRequest, opts ...grpc.CallOption) (*TimingSystemResponse, error) {
	out := new(TimingSystemResponse)
	err := c.cc.Invoke(ctx, "/timingsystem.TimingSystem/RecordTimingPoint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimingSystemServer is the server API for TimingSystem service.
type TimingSystemServer interface {
	RecordTimingPoint(context.Context, *TimingSystemRequest) (*TimingSystemResponse, error)
}

// UnimplementedTimingSystemServer can be embedded to have forward compatible implementations.
type UnimplementedTimingSystemServer struct {
}

func (*UnimplementedTimingSystemServer) RecordTimingPoint(ctx context.Context, req *TimingSystemRequest) (*TimingSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordTimingPoint not implemented")
}

func RegisterTimingSystemServer(s *grpc.Server, srv TimingSystemServer) {
	s.RegisterService(&_TimingSystem_serviceDesc, srv)
}

func _TimingSystem_RecordTimingPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimingSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimingSystemServer).RecordTimingPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/timingsystem.TimingSystem/RecordTimingPoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimingSystemServer).RecordTimingPoint(ctx, req.(*TimingSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimingSystem_serviceDesc = grpc.ServiceDesc{
	ServiceName: "timingsystem.TimingSystem",
	HandlerType: (*TimingSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordTimingPoint",
			Handler:    _TimingSystem_RecordTimingPoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "timingsystem.proto",
}
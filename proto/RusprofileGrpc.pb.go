// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RusprofileGrpc.proto

package rusprofileparserservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type FirmByINNRequest struct {
	Inn                  string   `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirmByINNRequest) Reset()         { *m = FirmByINNRequest{} }
func (m *FirmByINNRequest) String() string { return proto.CompactTextString(m) }
func (*FirmByINNRequest) ProtoMessage()    {}
func (*FirmByINNRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00f9a7ad2bdce93, []int{0}
}

func (m *FirmByINNRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirmByINNRequest.Unmarshal(m, b)
}
func (m *FirmByINNRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirmByINNRequest.Marshal(b, m, deterministic)
}
func (m *FirmByINNRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirmByINNRequest.Merge(m, src)
}
func (m *FirmByINNRequest) XXX_Size() int {
	return xxx_messageInfo_FirmByINNRequest.Size(m)
}
func (m *FirmByINNRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FirmByINNRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FirmByINNRequest proto.InternalMessageInfo

func (m *FirmByINNRequest) GetInn() string {
	if m != nil {
		return m.Inn
	}
	return ""
}

type FirmInfoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kpp                  string   `protobuf:"bytes,2,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Inn                  string   `protobuf:"bytes,3,opt,name=inn,proto3" json:"inn,omitempty"`
	Boss                 string   `protobuf:"bytes,4,opt,name=boss,proto3" json:"boss,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FirmInfoResponse) Reset()         { *m = FirmInfoResponse{} }
func (m *FirmInfoResponse) String() string { return proto.CompactTextString(m) }
func (*FirmInfoResponse) ProtoMessage()    {}
func (*FirmInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00f9a7ad2bdce93, []int{1}
}

func (m *FirmInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FirmInfoResponse.Unmarshal(m, b)
}
func (m *FirmInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FirmInfoResponse.Marshal(b, m, deterministic)
}
func (m *FirmInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FirmInfoResponse.Merge(m, src)
}
func (m *FirmInfoResponse) XXX_Size() int {
	return xxx_messageInfo_FirmInfoResponse.Size(m)
}
func (m *FirmInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FirmInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FirmInfoResponse proto.InternalMessageInfo

func (m *FirmInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FirmInfoResponse) GetKpp() string {
	if m != nil {
		return m.Kpp
	}
	return ""
}

func (m *FirmInfoResponse) GetInn() string {
	if m != nil {
		return m.Inn
	}
	return ""
}

func (m *FirmInfoResponse) GetBoss() string {
	if m != nil {
		return m.Boss
	}
	return ""
}

func init() {
	proto.RegisterType((*FirmByINNRequest)(nil), "rusprofileparserservice.FirmByINNRequest")
	proto.RegisterType((*FirmInfoResponse)(nil), "rusprofileparserservice.FirmInfoResponse")
}

func init() { proto.RegisterFile("RusprofileGrpc.proto", fileDescriptor_d00f9a7ad2bdce93) }

var fileDescriptor_d00f9a7ad2bdce93 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x09, 0x2a, 0x2d, 0x2e,
	0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x75, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x2f, 0x82, 0x8b, 0x16, 0x24, 0x16, 0x15, 0xa7, 0x82, 0x50, 0x59, 0x66, 0x72, 0xaa,
	0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e,
	0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x9b, 0x92, 0x0a, 0x97, 0x80, 0x5b, 0x66,
	0x51, 0xae, 0x53, 0xa5, 0xa7, 0x9f, 0x5f, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x00,
	0x17, 0x73, 0x66, 0x5e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0xa9, 0x14, 0x07,
	0x51, 0xe5, 0x99, 0x97, 0x96, 0x1f, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc4,
	0xc5, 0x92, 0x97, 0x98, 0x9b, 0x0a, 0x55, 0x06, 0x66, 0x83, 0x74, 0x66, 0x17, 0x14, 0x48, 0x30,
	0x41, 0x74, 0x66, 0x17, 0x14, 0xc0, 0xcc, 0x62, 0x86, 0x9b, 0x05, 0xd2, 0x97, 0x94, 0x5f, 0x5c,
	0x2c, 0xc1, 0x02, 0xd1, 0x07, 0x62, 0x1b, 0xf5, 0x33, 0x72, 0x89, 0x23, 0x7c, 0x15, 0x00, 0x76,
	0x7f, 0x30, 0xc4, 0xfd, 0x42, 0x25, 0x5c, 0xdc, 0x30, 0xbb, 0xdd, 0x53, 0x4b, 0x84, 0x34, 0xf5,
	0x70, 0x78, 0x54, 0x0f, 0xdd, 0x1f, 0x52, 0xf8, 0x95, 0x22, 0x7b, 0x46, 0x89, 0xbf, 0xe9, 0xf2,
	0x93, 0xc9, 0x4c, 0x9c, 0x4a, 0x2c, 0xfa, 0x99, 0x79, 0x79, 0x56, 0x8c, 0x5a, 0x49, 0x6c, 0xe0,
	0xe0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x59, 0xdb, 0x56, 0x5a, 0x6d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RusprofileParserServiceClient is the client API for RusprofileParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RusprofileParserServiceClient interface {
	//rpc FirmInfoGet(FirmByINNRequest) returns (FirmInfoResponse) {}
	FirmInfoGet(ctx context.Context, in *FirmByINNRequest, opts ...grpc.CallOption) (*FirmInfoResponse, error)
}

type rusprofileParserServiceClient struct {
	cc *grpc.ClientConn
}

func NewRusprofileParserServiceClient(cc *grpc.ClientConn) RusprofileParserServiceClient {
	return &rusprofileParserServiceClient{cc}
}

func (c *rusprofileParserServiceClient) FirmInfoGet(ctx context.Context, in *FirmByINNRequest, opts ...grpc.CallOption) (*FirmInfoResponse, error) {
	out := new(FirmInfoResponse)
	err := c.cc.Invoke(ctx, "/rusprofileparserservice.RusprofileParserService/FirmInfoGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RusprofileParserServiceServer is the server API for RusprofileParserService service.
type RusprofileParserServiceServer interface {
	//rpc FirmInfoGet(FirmByINNRequest) returns (FirmInfoResponse) {}
	FirmInfoGet(context.Context, *FirmByINNRequest) (*FirmInfoResponse, error)
}

// UnimplementedRusprofileParserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRusprofileParserServiceServer struct {
}

func (*UnimplementedRusprofileParserServiceServer) FirmInfoGet(ctx context.Context, req *FirmByINNRequest) (*FirmInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FirmInfoGet not implemented")
}

func RegisterRusprofileParserServiceServer(s *grpc.Server, srv RusprofileParserServiceServer) {
	s.RegisterService(&_RusprofileParserService_serviceDesc, srv)
}

func _RusprofileParserService_FirmInfoGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FirmByINNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RusprofileParserServiceServer).FirmInfoGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rusprofileparserservice.RusprofileParserService/FirmInfoGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RusprofileParserServiceServer).FirmInfoGet(ctx, req.(*FirmByINNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RusprofileParserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rusprofileparserservice.RusprofileParserService",
	HandlerType: (*RusprofileParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FirmInfoGet",
			Handler:    _RusprofileParserService_FirmInfoGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RusprofileGrpc.proto",
}

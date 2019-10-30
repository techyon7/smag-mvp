// Code generated by protoc-gen-go. DO NOT EDIT.
// source: renewingAddress.proto

package proto

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

type RenewedElasticResult struct {
	IsRenewed            bool     `protobuf:"varint,1,opt,name=isRenewed,proto3" json:"isRenewed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewedElasticResult) Reset()         { *m = RenewedElasticResult{} }
func (m *RenewedElasticResult) String() string { return proto.CompactTextString(m) }
func (*RenewedElasticResult) ProtoMessage()    {}
func (*RenewedElasticResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_55c265b1376e8420, []int{0}
}

func (m *RenewedElasticResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewedElasticResult.Unmarshal(m, b)
}
func (m *RenewedElasticResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewedElasticResult.Marshal(b, m, deterministic)
}
func (m *RenewedElasticResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewedElasticResult.Merge(m, src)
}
func (m *RenewedElasticResult) XXX_Size() int {
	return xxx_messageInfo_RenewedElasticResult.Size(m)
}
func (m *RenewedElasticResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewedElasticResult.DiscardUnknown(m)
}

var xxx_messageInfo_RenewedElasticResult proto.InternalMessageInfo

func (m *RenewedElasticResult) GetIsRenewed() bool {
	if m != nil {
		return m.IsRenewed
	}
	return false
}

type RenewingElasticIp struct {
	InstanceId           string   `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Pod                  string   `protobuf:"bytes,3,opt,name=pod,proto3" json:"pod,omitempty"`
	PodIp                string   `protobuf:"bytes,4,opt,name=pod_ip,json=podIp,proto3" json:"pod_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenewingElasticIp) Reset()         { *m = RenewingElasticIp{} }
func (m *RenewingElasticIp) String() string { return proto.CompactTextString(m) }
func (*RenewingElasticIp) ProtoMessage()    {}
func (*RenewingElasticIp) Descriptor() ([]byte, []int) {
	return fileDescriptor_55c265b1376e8420, []int{1}
}

func (m *RenewingElasticIp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenewingElasticIp.Unmarshal(m, b)
}
func (m *RenewingElasticIp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenewingElasticIp.Marshal(b, m, deterministic)
}
func (m *RenewingElasticIp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenewingElasticIp.Merge(m, src)
}
func (m *RenewingElasticIp) XXX_Size() int {
	return xxx_messageInfo_RenewingElasticIp.Size(m)
}
func (m *RenewingElasticIp) XXX_DiscardUnknown() {
	xxx_messageInfo_RenewingElasticIp.DiscardUnknown(m)
}

var xxx_messageInfo_RenewingElasticIp proto.InternalMessageInfo

func (m *RenewingElasticIp) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *RenewingElasticIp) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *RenewingElasticIp) GetPod() string {
	if m != nil {
		return m.Pod
	}
	return ""
}

func (m *RenewingElasticIp) GetPodIp() string {
	if m != nil {
		return m.PodIp
	}
	return ""
}

func init() {
	proto.RegisterType((*RenewedElasticResult)(nil), "proto.RenewedElasticResult")
	proto.RegisterType((*RenewingElasticIp)(nil), "proto.RenewingElasticIp")
}

func init() { proto.RegisterFile("renewingAddress.proto", fileDescriptor_55c265b1376e8420) }

var fileDescriptor_55c265b1376e8420 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0xad, 0xfd, 0xc1, 0xce, 0x42, 0xea, 0x60, 0x21, 0xa8, 0x88, 0x74, 0xe5, 0xaa, 0x0b,
	0xf5, 0x05, 0x5c, 0xb8, 0xc8, 0x36, 0xae, 0x45, 0x6a, 0x33, 0x48, 0xa0, 0x24, 0x43, 0x12, 0xf5,
	0xf5, 0xa5, 0x69, 0xb9, 0x2d, 0xf7, 0xae, 0x32, 0x7c, 0x1f, 0x93, 0x39, 0x07, 0x5a, 0x4f, 0x96,
	0xfe, 0x8c, 0xfd, 0x7e, 0xd5, 0xda, 0x53, 0x08, 0x3d, 0x7b, 0x17, 0x1d, 0x96, 0xe9, 0xe9, 0x5e,
	0xe0, 0x5a, 0xcd, 0x9e, 0xf4, 0xdb, 0x34, 0x84, 0x68, 0x46, 0x45, 0xe1, 0x67, 0x8a, 0x78, 0x07,
	0xb5, 0x09, 0xab, 0x11, 0xd9, 0x43, 0xf6, 0x78, 0xa1, 0x36, 0xd0, 0x31, 0x5c, 0xa9, 0xf5, 0xd7,
	0x75, 0x4d, 0x32, 0xde, 0x03, 0x18, 0x1b, 0xe2, 0x60, 0x47, 0x92, 0xcb, 0x4e, 0xad, 0x76, 0x04,
	0x11, 0x0a, 0xeb, 0x34, 0x89, 0xf3, 0x64, 0xd2, 0x8c, 0x0d, 0xe4, 0xec, 0xb4, 0xc8, 0x13, 0x9a,
	0x47, 0x6c, 0xa1, 0x62, 0xa7, 0x3f, 0x0d, 0x8b, 0x22, 0xc1, 0x92, 0x9d, 0x96, 0xfc, 0xf4, 0x01,
	0xcd, 0xe1, 0xd2, 0x3b, 0xf9, 0x5f, 0x33, 0x12, 0x4a, 0xb8, 0x4c, 0xdd, 0xb6, 0x08, 0x62, 0x29,
	0xd7, 0x9f, 0x84, 0xbb, 0xb9, 0xdd, 0x9b, 0xa3, 0xb2, 0xdd, 0xd9, 0x57, 0x95, 0xec, 0xf3, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x22, 0x5f, 0xd7, 0xe4, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ElasticIpServiceClient is the client API for ElasticIpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElasticIpServiceClient interface {
	RenewElasticIp(ctx context.Context, in *RenewingElasticIp, opts ...grpc.CallOption) (*RenewedElasticResult, error)
}

type elasticIpServiceClient struct {
	cc *grpc.ClientConn
}

func NewElasticIpServiceClient(cc *grpc.ClientConn) ElasticIpServiceClient {
	return &elasticIpServiceClient{cc}
}

func (c *elasticIpServiceClient) RenewElasticIp(ctx context.Context, in *RenewingElasticIp, opts ...grpc.CallOption) (*RenewedElasticResult, error) {
	out := new(RenewedElasticResult)
	err := c.cc.Invoke(ctx, "/proto.ElasticIpService/renewElasticIp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticIpServiceServer is the server API for ElasticIpService service.
type ElasticIpServiceServer interface {
	RenewElasticIp(context.Context, *RenewingElasticIp) (*RenewedElasticResult, error)
}

// UnimplementedElasticIpServiceServer can be embedded to have forward compatible implementations.
type UnimplementedElasticIpServiceServer struct {
}

func (*UnimplementedElasticIpServiceServer) RenewElasticIp(ctx context.Context, req *RenewingElasticIp) (*RenewedElasticResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewElasticIp not implemented")
}

func RegisterElasticIpServiceServer(s *grpc.Server, srv ElasticIpServiceServer) {
	s.RegisterService(&_ElasticIpService_serviceDesc, srv)
}

func _ElasticIpService_RenewElasticIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewingElasticIp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticIpServiceServer).RenewElasticIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ElasticIpService/RenewElasticIp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticIpServiceServer).RenewElasticIp(ctx, req.(*RenewingElasticIp))
	}
	return interceptor(ctx, in, info, handler)
}

var _ElasticIpService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ElasticIpService",
	HandlerType: (*ElasticIpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "renewElasticIp",
			Handler:    _ElasticIpService_RenewElasticIp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "renewingAddress.proto",
}

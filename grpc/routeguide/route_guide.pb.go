/*
Package routeguide is a generated protocol buffer package.

It is generated from these files:
	route_guide.proto

It has these top-level messages:
	RandResponse
	RandRequest
*/
package routeguide

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

const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// request
type RandRequest struct {
	MaxNum int64 `protobuf:"varint,1,opt,name=maxnum" json:"maxnum,omitempty"`
}

func (m *RandRequest) Reset()                    { *m = RandRequest{} }
func (m *RandRequest) String() string            { return proto.CompactTextString(m) }
func (*RandRequest) ProtoMessage()               {}
func (*RandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// response
type RandResponse struct {
	RandNum int64 `protobuf:"varint,1,opt,name=randnum" json:"randnum,omitempty"`
}

func (m *RandResponse) Reset()                    { *m = RandResponse{} }
func (m *RandResponse) String() string            { return proto.CompactTextString(m) }
func (*RandResponse) ProtoMessage()               {}
func (*RandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*RandRequest)(nil), "HelloRequest")
	proto.RegisterType((*RandResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Jaguar service

type JaguarClient interface {
	// Sends random number request
	GetRand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandResponse, error)
}

type jaguarClient struct {
	cc *grpc.ClientConn
}

func NewJaguarClient(cc *grpc.ClientConn) JaguarClient {
	return &jaguarClient{cc}
}

func (c *jaguarClient) GetRand(ctx context.Context, in *RandRequest, opts ...grpc.CallOption) (*RandResponse, error) {
	out := new(RandResponse)
	err := grpc.Invoke(ctx, "/jaguar/GetRand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jaguar service
type JaguarServer interface {
	// Sends a random number response
	GetRand(context.Context, *RandRequest) (*RandResponse, error)
}

func RegisterJaguarServer(s *grpc.Server, srv JaguarServer) {
	s.RegisterService(&_Jaguar_serviceDesc, srv)
}

func _GetRand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JaguarServer).GetRand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaguar/GetRand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JaguarServer).GetRand(ctx, req.(*RandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Jaguar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jaguar",
	HandlerType: (*JaguarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRand",
			Handler:    _GetRand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0x6b, 0xe4, 0xc9, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a,
	0x24, 0x64, 0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x25, 0x24, 0xa1, 0x87, 0xe4, 0x02, 0x64,
	0xcb, 0xa4, 0xc4, 0xb0, 0xc8, 0x00, 0xad, 0x50, 0x62, 0x70, 0x32, 0xe0, 0x92, 0xce, 0xcc, 0xd7,
	0x4b, 0x2f, 0x2a, 0x48, 0xd6, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0x46, 0x52, 0xeb,
	0xc4, 0x0f, 0x56, 0x1c, 0x0e, 0x62, 0x07, 0x80, 0xbc, 0x14, 0xc0, 0x98, 0xc4, 0x06, 0xf6, 0x9b,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xb7, 0xcd, 0xf2, 0xef, 0x00, 0x00, 0x00,
}

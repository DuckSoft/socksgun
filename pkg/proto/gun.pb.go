// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gun.proto

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

type Packet struct {
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Padding              []byte   `protobuf:"bytes,3,opt,name=padding,proto3" json:"padding,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eb68c7936423302, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Packet) GetPadding() []byte {
	if m != nil {
		return m.Padding
	}
	return nil
}

func init() {
	proto.RegisterType((*Packet)(nil), "Packet")
}

func init() { proto.RegisterFile("gun.proto", fileDescriptor_5eb68c7936423302) }

var fileDescriptor_5eb68c7936423302 = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2f, 0xcd, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe1, 0x62, 0x0b, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x11,
	0x92, 0xe0, 0x62, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0,
	0x09, 0x82, 0x71, 0x41, 0x32, 0x05, 0x89, 0x29, 0x29, 0x99, 0x79, 0xe9, 0x12, 0xcc, 0x10, 0x19,
	0x28, 0xd7, 0x48, 0x8f, 0x8b, 0x3f, 0x38, 0x3f, 0x39, 0xbb, 0xd8, 0xbd, 0x34, 0x2f, 0x38, 0xb5,
	0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x9a, 0x8b, 0x39, 0xa4, 0x34, 0x4f, 0x88, 0x5d, 0x0f, 0x62,
	0xac, 0x14, 0x8c, 0xa1, 0xc1, 0x68, 0xc0, 0xe8, 0x24, 0x12, 0x25, 0x54, 0x0c, 0x52, 0x9f, 0x5e,
	0x9a, 0xa7, 0x5f, 0x90, 0x9d, 0xae, 0x0f, 0x76, 0x43, 0x12, 0x1b, 0x98, 0x32, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x76, 0xdc, 0x46, 0x66, 0x97, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SocksGunServiceClient is the client API for SocksGunService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SocksGunServiceClient interface {
	Tun(ctx context.Context, opts ...grpc.CallOption) (SocksGunService_TunClient, error)
}

type socksGunServiceClient struct {
	cc *grpc.ClientConn
}

func NewSocksGunServiceClient(cc *grpc.ClientConn) SocksGunServiceClient {
	return &socksGunServiceClient{cc}
}

func (c *socksGunServiceClient) Tun(ctx context.Context, opts ...grpc.CallOption) (SocksGunService_TunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SocksGunService_serviceDesc.Streams[0], "/SocksGunService/Tun", opts...)
	if err != nil {
		return nil, err
	}
	x := &socksGunServiceTunClient{stream}
	return x, nil
}

type SocksGunService_TunClient interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ClientStream
}

type socksGunServiceTunClient struct {
	grpc.ClientStream
}

func (x *socksGunServiceTunClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *socksGunServiceTunClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SocksGunServiceServer is the server API for SocksGunService service.
type SocksGunServiceServer interface {
	Tun(SocksGunService_TunServer) error
}

// UnimplementedSocksGunServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSocksGunServiceServer struct {
}

func (*UnimplementedSocksGunServiceServer) Tun(srv SocksGunService_TunServer) error {
	return status.Errorf(codes.Unimplemented, "method Tun not implemented")
}

func RegisterSocksGunServiceServer(s *grpc.Server, srv SocksGunServiceServer) {
	s.RegisterService(&_SocksGunService_serviceDesc, srv)
}

func _SocksGunService_Tun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SocksGunServiceServer).Tun(&socksGunServiceTunServer{stream})
}

type SocksGunService_TunServer interface {
	Send(*Packet) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type socksGunServiceTunServer struct {
	grpc.ServerStream
}

func (x *socksGunServiceTunServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *socksGunServiceTunServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SocksGunService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SocksGunService",
	HandlerType: (*SocksGunServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tun",
			Handler:       _SocksGunService_Tun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gun.proto",
}

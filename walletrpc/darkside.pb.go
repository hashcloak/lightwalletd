// Code generated by protoc-gen-go. DO NOT EDIT.
// source: darkside.proto

package walletrpc

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

type DarksideState struct {
	LatestHeight         uint64   `protobuf:"varint,1,opt,name=latestHeight,proto3" json:"latestHeight,omitempty"`
	ReorgHeight          uint64   `protobuf:"varint,2,opt,name=reorgHeight,proto3" json:"reorgHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DarksideState) Reset()         { *m = DarksideState{} }
func (m *DarksideState) String() string { return proto.CompactTextString(m) }
func (*DarksideState) ProtoMessage()    {}
func (*DarksideState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea18aa1b2b1f163, []int{0}
}

func (m *DarksideState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DarksideState.Unmarshal(m, b)
}
func (m *DarksideState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DarksideState.Marshal(b, m, deterministic)
}
func (m *DarksideState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DarksideState.Merge(m, src)
}
func (m *DarksideState) XXX_Size() int {
	return xxx_messageInfo_DarksideState.Size(m)
}
func (m *DarksideState) XXX_DiscardUnknown() {
	xxx_messageInfo_DarksideState.DiscardUnknown(m)
}

var xxx_messageInfo_DarksideState proto.InternalMessageInfo

func (m *DarksideState) GetLatestHeight() uint64 {
	if m != nil {
		return m.LatestHeight
	}
	return 0
}

func (m *DarksideState) GetReorgHeight() uint64 {
	if m != nil {
		return m.ReorgHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*DarksideState)(nil), "cash.z.wallet.sdk.rpc.DarksideState")
}

func init() {
	proto.RegisterFile("darkside.proto", fileDescriptor_5ea18aa1b2b1f163)
}

var fileDescriptor_5ea18aa1b2b1f163 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x5d, 0x4a, 0xc4, 0x30,
	0x10, 0x6e, 0x16, 0x11, 0x8c, 0xee, 0x22, 0x01, 0x41, 0x8a, 0xe0, 0x12, 0x14, 0x7c, 0x0a, 0xa2,
	0x37, 0x10, 0x45, 0x7d, 0x5d, 0x15, 0xc4, 0xb7, 0x31, 0x1d, 0xba, 0x71, 0xdb, 0x24, 0x4c, 0x06,
	0x17, 0x3d, 0x9a, 0x27, 0xf1, 0x38, 0x42, 0x5b, 0xb1, 0x05, 0xbb, 0xaf, 0xdf, 0xdf, 0x7c, 0xf3,
	0xc9, 0x59, 0x01, 0xb4, 0x4a, 0xae, 0x40, 0x13, 0x29, 0x70, 0x50, 0x07, 0x16, 0xd2, 0xd2, 0x7c,
	0x9a, 0x35, 0x54, 0x15, 0xb2, 0x49, 0xc5, 0xca, 0x50, 0xb4, 0xf9, 0x34, 0x21, 0xbd, 0x3b, 0xdb,
	0xa9, 0xf4, 0x93, 0x9c, 0x5e, 0x77, 0xbe, 0x07, 0x06, 0x46, 0xa5, 0xe5, 0x5e, 0x05, 0x8c, 0x89,
	0xef, 0xd0, 0x95, 0x4b, 0x3e, 0x14, 0x73, 0x71, 0xb6, 0xb5, 0x18, 0x60, 0x6a, 0x2e, 0x77, 0x09,
	0x03, 0x95, 0x9d, 0x64, 0xd2, 0x48, 0xfa, 0xd0, 0xc5, 0xb7, 0x90, 0xfb, 0x7f, 0xb9, 0x84, 0x50,
	0x23, 0xa9, 0x37, 0x79, 0xfc, 0x8b, 0xdd, 0x22, 0xdf, 0x7b, 0x1b, 0x6a, 0xe7, 0xcb, 0x47, 0x02,
	0x9f, 0xc0, 0xb2, 0x0b, 0x3e, 0xa9, 0x23, 0xf3, 0x6f, 0x6b, 0x73, 0x53, 0x47, 0xfe, 0xc8, 0x4f,
	0x47, 0xd8, 0x05, 0xac, 0x7b, 0x29, 0x3a, 0x3b, 0x17, 0xea, 0xb9, 0x77, 0x1f, 0xb9, 0x7d, 0xed,
	0x64, 0xc4, 0x3e, 0x18, 0x20, 0xdf, 0x58, 0x41, 0x67, 0x57, 0xb3, 0x97, 0x9d, 0x96, 0xa1, 0x68,
	0xbf, 0x26, 0xd9, 0xeb, 0x76, 0x33, 0xe4, 0xe5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x99,
	0x2c, 0x75, 0x80, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DarksideStreamerClient is the client API for DarksideStreamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DarksideStreamerClient interface {
	// Return the list of transactions that have been submitted (via SendTransaction).
	DarksideGetIncomingTransactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DarksideStreamer_DarksideGetIncomingTransactionsClient, error)
	// Set the information that GetLightdInfo returns, except that chainName specifies
	// a file of blocks within testdata/darkside that GetBlock will return.
	DarksideSetState(ctx context.Context, in *DarksideState, opts ...grpc.CallOption) (*Empty, error)
}

type darksideStreamerClient struct {
	cc grpc.ClientConnInterface
}

func NewDarksideStreamerClient(cc grpc.ClientConnInterface) DarksideStreamerClient {
	return &darksideStreamerClient{cc}
}

func (c *darksideStreamerClient) DarksideGetIncomingTransactions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (DarksideStreamer_DarksideGetIncomingTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DarksideStreamer_serviceDesc.Streams[0], "/cash.z.wallet.sdk.rpc.DarksideStreamer/DarksideGetIncomingTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &darksideStreamerDarksideGetIncomingTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DarksideStreamer_DarksideGetIncomingTransactionsClient interface {
	Recv() (*RawTransaction, error)
	grpc.ClientStream
}

type darksideStreamerDarksideGetIncomingTransactionsClient struct {
	grpc.ClientStream
}

func (x *darksideStreamerDarksideGetIncomingTransactionsClient) Recv() (*RawTransaction, error) {
	m := new(RawTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *darksideStreamerClient) DarksideSetState(ctx context.Context, in *DarksideState, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.DarksideStreamer/DarksideSetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DarksideStreamerServer is the server API for DarksideStreamer service.
type DarksideStreamerServer interface {
	// Return the list of transactions that have been submitted (via SendTransaction).
	DarksideGetIncomingTransactions(*Empty, DarksideStreamer_DarksideGetIncomingTransactionsServer) error
	// Set the information that GetLightdInfo returns, except that chainName specifies
	// a file of blocks within testdata/darkside that GetBlock will return.
	DarksideSetState(context.Context, *DarksideState) (*Empty, error)
}

// UnimplementedDarksideStreamerServer can be embedded to have forward compatible implementations.
type UnimplementedDarksideStreamerServer struct {
}

func (*UnimplementedDarksideStreamerServer) DarksideGetIncomingTransactions(req *Empty, srv DarksideStreamer_DarksideGetIncomingTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method DarksideGetIncomingTransactions not implemented")
}
func (*UnimplementedDarksideStreamerServer) DarksideSetState(ctx context.Context, req *DarksideState) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DarksideSetState not implemented")
}

func RegisterDarksideStreamerServer(s *grpc.Server, srv DarksideStreamerServer) {
	s.RegisterService(&_DarksideStreamer_serviceDesc, srv)
}

func _DarksideStreamer_DarksideGetIncomingTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DarksideStreamerServer).DarksideGetIncomingTransactions(m, &darksideStreamerDarksideGetIncomingTransactionsServer{stream})
}

type DarksideStreamer_DarksideGetIncomingTransactionsServer interface {
	Send(*RawTransaction) error
	grpc.ServerStream
}

type darksideStreamerDarksideGetIncomingTransactionsServer struct {
	grpc.ServerStream
}

func (x *darksideStreamerDarksideGetIncomingTransactionsServer) Send(m *RawTransaction) error {
	return x.ServerStream.SendMsg(m)
}

func _DarksideStreamer_DarksideSetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DarksideState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DarksideStreamerServer).DarksideSetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.DarksideStreamer/DarksideSetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DarksideStreamerServer).DarksideSetState(ctx, req.(*DarksideState))
	}
	return interceptor(ctx, in, info, handler)
}

var _DarksideStreamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cash.z.wallet.sdk.rpc.DarksideStreamer",
	HandlerType: (*DarksideStreamerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DarksideSetState",
			Handler:    _DarksideStreamer_DarksideSetState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DarksideGetIncomingTransactions",
			Handler:       _DarksideStreamer_DarksideGetIncomingTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "darkside.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

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

// A BlockID message contains identifiers to select a block: a height or a
// hash. Specification by hash is not implemented, but may be in the future.
type BlockID struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// BlockRange specifies a series of blocks from start to end inclusive.
// Both BlockIDs must be heights; specification by hash is not yet supported.
type BlockRange struct {
	Start                *BlockID `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *BlockID `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRange) Reset()         { *m = BlockRange{} }
func (m *BlockRange) String() string { return proto.CompactTextString(m) }
func (*BlockRange) ProtoMessage()    {}
func (*BlockRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *BlockRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRange.Unmarshal(m, b)
}
func (m *BlockRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRange.Marshal(b, m, deterministic)
}
func (m *BlockRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRange.Merge(m, src)
}
func (m *BlockRange) XXX_Size() int {
	return xxx_messageInfo_BlockRange.Size(m)
}
func (m *BlockRange) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRange.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRange proto.InternalMessageInfo

func (m *BlockRange) GetStart() *BlockID {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *BlockRange) GetEnd() *BlockID {
	if m != nil {
		return m.End
	}
	return nil
}

// A TxFilter contains the information needed to identify a particular
// transaction: either a block and an index, or a direct transaction hash.
// Currently, only specification by hash is supported.
type TxFilter struct {
	Block                *BlockID `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Index                uint64   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxFilter) Reset()         { *m = TxFilter{} }
func (m *TxFilter) String() string { return proto.CompactTextString(m) }
func (*TxFilter) ProtoMessage()    {}
func (*TxFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *TxFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxFilter.Unmarshal(m, b)
}
func (m *TxFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxFilter.Marshal(b, m, deterministic)
}
func (m *TxFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxFilter.Merge(m, src)
}
func (m *TxFilter) XXX_Size() int {
	return xxx_messageInfo_TxFilter.Size(m)
}
func (m *TxFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TxFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TxFilter proto.InternalMessageInfo

func (m *TxFilter) GetBlock() *BlockID {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *TxFilter) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxFilter) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// RawTransaction contains the complete transaction data. It also optionally includes
// the block height in which the transaction was included
type RawTransaction struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawTransaction) Reset()         { *m = RawTransaction{} }
func (m *RawTransaction) String() string { return proto.CompactTextString(m) }
func (*RawTransaction) ProtoMessage()    {}
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RawTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawTransaction.Unmarshal(m, b)
}
func (m *RawTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawTransaction.Marshal(b, m, deterministic)
}
func (m *RawTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTransaction.Merge(m, src)
}
func (m *RawTransaction) XXX_Size() int {
	return xxx_messageInfo_RawTransaction.Size(m)
}
func (m *RawTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_RawTransaction proto.InternalMessageInfo

func (m *RawTransaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RawTransaction) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// A SendResponse encodes an error code and a string. It is currently used
// only by SendTransaction(). If error code is zero, the operation was
// successful; if non-zero, it and the message specify the failure.
type SendResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *SendResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

// Chainspec is a placeholder to allow specification of a particular chain fork.
type ChainSpec struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainSpec) Reset()         { *m = ChainSpec{} }
func (m *ChainSpec) String() string { return proto.CompactTextString(m) }
func (*ChainSpec) ProtoMessage()    {}
func (*ChainSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ChainSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainSpec.Unmarshal(m, b)
}
func (m *ChainSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainSpec.Marshal(b, m, deterministic)
}
func (m *ChainSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainSpec.Merge(m, src)
}
func (m *ChainSpec) XXX_Size() int {
	return xxx_messageInfo_ChainSpec.Size(m)
}
func (m *ChainSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChainSpec proto.InternalMessageInfo

// Empty is for gRPCs that take no arguments, currently only GetLightdInfo.
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// LightdInfo returns various information about this lightwalletd instance
// and the state of the blockchain.
type LightdInfo struct {
	Version                 string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Vendor                  string   `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	TaddrSupport            bool     `protobuf:"varint,3,opt,name=taddrSupport,proto3" json:"taddrSupport,omitempty"`
	ChainName               string   `protobuf:"bytes,4,opt,name=chainName,proto3" json:"chainName,omitempty"`
	SaplingActivationHeight uint64   `protobuf:"varint,5,opt,name=saplingActivationHeight,proto3" json:"saplingActivationHeight,omitempty"`
	ConsensusBranchId       string   `protobuf:"bytes,6,opt,name=consensusBranchId,proto3" json:"consensusBranchId,omitempty"`
	BlockHeight             uint64   `protobuf:"varint,7,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *LightdInfo) Reset()         { *m = LightdInfo{} }
func (m *LightdInfo) String() string { return proto.CompactTextString(m) }
func (*LightdInfo) ProtoMessage()    {}
func (*LightdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *LightdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightdInfo.Unmarshal(m, b)
}
func (m *LightdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightdInfo.Marshal(b, m, deterministic)
}
func (m *LightdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightdInfo.Merge(m, src)
}
func (m *LightdInfo) XXX_Size() int {
	return xxx_messageInfo_LightdInfo.Size(m)
}
func (m *LightdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LightdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LightdInfo proto.InternalMessageInfo

func (m *LightdInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LightdInfo) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *LightdInfo) GetTaddrSupport() bool {
	if m != nil {
		return m.TaddrSupport
	}
	return false
}

func (m *LightdInfo) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *LightdInfo) GetSaplingActivationHeight() uint64 {
	if m != nil {
		return m.SaplingActivationHeight
	}
	return 0
}

func (m *LightdInfo) GetConsensusBranchId() string {
	if m != nil {
		return m.ConsensusBranchId
	}
	return ""
}

func (m *LightdInfo) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// TransparentAddressBlockFilter restricts the results to the given address
// or block range.
type TransparentAddressBlockFilter struct {
	Address              string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Range                *BlockRange `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransparentAddressBlockFilter) Reset()         { *m = TransparentAddressBlockFilter{} }
func (m *TransparentAddressBlockFilter) String() string { return proto.CompactTextString(m) }
func (*TransparentAddressBlockFilter) ProtoMessage()    {}
func (*TransparentAddressBlockFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *TransparentAddressBlockFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransparentAddressBlockFilter.Unmarshal(m, b)
}
func (m *TransparentAddressBlockFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransparentAddressBlockFilter.Marshal(b, m, deterministic)
}
func (m *TransparentAddressBlockFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransparentAddressBlockFilter.Merge(m, src)
}
func (m *TransparentAddressBlockFilter) XXX_Size() int {
	return xxx_messageInfo_TransparentAddressBlockFilter.Size(m)
}
func (m *TransparentAddressBlockFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TransparentAddressBlockFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TransparentAddressBlockFilter proto.InternalMessageInfo

func (m *TransparentAddressBlockFilter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TransparentAddressBlockFilter) GetRange() *BlockRange {
	if m != nil {
		return m.Range
	}
	return nil
}

// Duration is currently used only for testing, so that the Ping rpc
// can simulate a delay, to create many simultaneous connections. Units
// are microseconds.
type Duration struct {
	IntervalUs           int64    `protobuf:"varint,1,opt,name=intervalUs,proto3" json:"intervalUs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *Duration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duration.Unmarshal(m, b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return xxx_messageInfo_Duration.Size(m)
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

func (m *Duration) GetIntervalUs() int64 {
	if m != nil {
		return m.IntervalUs
	}
	return 0
}

// PingResponse is used to indicate concurrency, how many Ping rpcs
// are executing upon entry and upon exit (after the delay).
type PingResponse struct {
	Entry                int64    `protobuf:"varint,1,opt,name=entry,proto3" json:"entry,omitempty"`
	Exit                 int64    `protobuf:"varint,2,opt,name=exit,proto3" json:"exit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetEntry() int64 {
	if m != nil {
		return m.Entry
	}
	return 0
}

func (m *PingResponse) GetExit() int64 {
	if m != nil {
		return m.Exit
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockID)(nil), "cash.z.wallet.sdk.rpc.BlockID")
	proto.RegisterType((*BlockRange)(nil), "cash.z.wallet.sdk.rpc.BlockRange")
	proto.RegisterType((*TxFilter)(nil), "cash.z.wallet.sdk.rpc.TxFilter")
	proto.RegisterType((*RawTransaction)(nil), "cash.z.wallet.sdk.rpc.RawTransaction")
	proto.RegisterType((*SendResponse)(nil), "cash.z.wallet.sdk.rpc.SendResponse")
	proto.RegisterType((*ChainSpec)(nil), "cash.z.wallet.sdk.rpc.ChainSpec")
	proto.RegisterType((*Empty)(nil), "cash.z.wallet.sdk.rpc.Empty")
	proto.RegisterType((*LightdInfo)(nil), "cash.z.wallet.sdk.rpc.LightdInfo")
	proto.RegisterType((*TransparentAddressBlockFilter)(nil), "cash.z.wallet.sdk.rpc.TransparentAddressBlockFilter")
	proto.RegisterType((*Duration)(nil), "cash.z.wallet.sdk.rpc.Duration")
	proto.RegisterType((*PingResponse)(nil), "cash.z.wallet.sdk.rpc.PingResponse")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5f, 0x4f, 0x13, 0x41,
	0x10, 0xef, 0xbf, 0xa3, 0xed, 0x50, 0x20, 0x6c, 0x40, 0x9b, 0x06, 0xb1, 0xae, 0x31, 0x21, 0xc6,
	0x34, 0x04, 0x31, 0xf2, 0xe0, 0x0b, 0x7f, 0x14, 0x49, 0x90, 0xe0, 0xb6, 0xbe, 0xe0, 0x03, 0x59,
	0xee, 0x86, 0xde, 0x49, 0xbb, 0x77, 0xd9, 0x5d, 0x4a, 0xf1, 0x23, 0xf8, 0x91, 0xfc, 0x50, 0x7e,
	0x06, 0xb3, 0xbb, 0x57, 0x38, 0xa2, 0x47, 0xfb, 0x76, 0x33, 0x3b, 0xf3, 0x9b, 0xd9, 0xdf, 0xfe,
	0x66, 0x0e, 0x16, 0x14, 0xca, 0x51, 0xe4, 0x63, 0x27, 0x91, 0xb1, 0x8e, 0xc9, 0xaa, 0xcf, 0x55,
	0xd8, 0xf9, 0xd9, 0xb9, 0xe1, 0x83, 0x01, 0xea, 0x8e, 0x0a, 0xae, 0x3a, 0x32, 0xf1, 0x5b, 0xab,
	0x7e, 0x3c, 0x4c, 0xb8, 0xaf, 0xcf, 0x2f, 0x63, 0x39, 0xe4, 0x5a, 0xb9, 0x68, 0xfa, 0x0e, 0xaa,
	0x7b, 0x83, 0xd8, 0xbf, 0x3a, 0x3a, 0x20, 0x4f, 0x60, 0x2e, 0xc4, 0xa8, 0x1f, 0xea, 0x66, 0xb1,
	0x5d, 0xdc, 0xa8, 0xb0, 0xd4, 0x22, 0x04, 0x2a, 0x21, 0x57, 0x61, 0xb3, 0xd4, 0x2e, 0x6e, 0x34,
	0x98, 0xfd, 0xa6, 0x1a, 0xc0, 0xa6, 0x31, 0x2e, 0xfa, 0x48, 0xb6, 0xc1, 0x53, 0x9a, 0x4b, 0x97,
	0x38, 0xbf, 0xb5, 0xde, 0xf9, 0x6f, 0x0b, 0x9d, 0xb4, 0x10, 0x73, 0xc1, 0x64, 0x13, 0xca, 0x28,
	0x02, 0x0b, 0x3b, 0x3d, 0xc7, 0x84, 0xd2, 0x1f, 0x50, 0xeb, 0x8d, 0x3f, 0x45, 0x03, 0x8d, 0xd2,
	0xd4, 0xbc, 0x30, 0x67, 0xb3, 0xd6, 0xb4, 0xc1, 0x64, 0x05, 0xbc, 0x48, 0x04, 0x38, 0xb6, 0x55,
	0x2b, 0xcc, 0x19, 0x77, 0x37, 0x2c, 0x67, 0x6e, 0xf8, 0x01, 0x16, 0x19, 0xbf, 0xe9, 0x49, 0x2e,
	0x14, 0xf7, 0x75, 0x14, 0x0b, 0x13, 0x15, 0x70, 0xcd, 0x6d, 0xc1, 0x06, 0xb3, 0xdf, 0x19, 0xce,
	0x4a, 0x59, 0xce, 0xe8, 0x29, 0x34, 0xba, 0x28, 0x02, 0x86, 0x2a, 0x89, 0x85, 0x42, 0xb2, 0x06,
	0x75, 0x94, 0x32, 0x96, 0xfb, 0x71, 0x80, 0x16, 0xc0, 0x63, 0xf7, 0x0e, 0x42, 0xa1, 0x61, 0x8d,
	0x2f, 0xa8, 0x14, 0xef, 0xa3, 0xc5, 0xaa, 0xb3, 0x07, 0x3e, 0x3a, 0x0f, 0xf5, 0xfd, 0x90, 0x47,
	0xa2, 0x9b, 0xa0, 0x4f, 0xab, 0xe0, 0x7d, 0x1c, 0x26, 0xfa, 0x96, 0xfe, 0x2a, 0x01, 0x1c, 0x9b,
	0x8a, 0xc1, 0x91, 0xb8, 0x8c, 0x49, 0x13, 0xaa, 0x23, 0x94, 0x2a, 0x8a, 0x85, 0x2d, 0x52, 0x67,
	0x13, 0xd3, 0x34, 0x3a, 0x42, 0x11, 0xc4, 0x32, 0x05, 0x4f, 0x2d, 0x53, 0x5a, 0xf3, 0x20, 0x90,
	0xdd, 0xeb, 0x24, 0x89, 0xa5, 0xb6, 0x14, 0xd4, 0xd8, 0x03, 0x9f, 0x69, 0xde, 0x37, 0xa5, 0x4f,
	0xf8, 0x10, 0x9b, 0x15, 0x9b, 0x7e, 0xef, 0x20, 0x3b, 0xf0, 0x54, 0xf1, 0x64, 0x10, 0x89, 0xfe,
	0xae, 0xaf, 0xa3, 0x11, 0x37, 0x5c, 0x7d, 0x76, 0x9c, 0x78, 0x96, 0x93, 0xbc, 0x63, 0xf2, 0x06,
	0x96, 0x7d, 0xc3, 0x8e, 0x50, 0xd7, 0x6a, 0x4f, 0x72, 0xe1, 0x87, 0x47, 0x41, 0x73, 0xce, 0xe2,
	0xff, 0x7b, 0x40, 0xda, 0x30, 0x6f, 0xdf, 0x30, 0xc5, 0xae, 0x5a, 0xec, 0xac, 0x8b, 0x4a, 0x78,
	0x66, 0xdf, 0x2b, 0xe1, 0x12, 0x85, 0xde, 0x0d, 0x02, 0x89, 0x4a, 0x59, 0x01, 0xa4, 0x9a, 0x69,
	0x42, 0x95, 0x3b, 0xef, 0x84, 0x9e, 0xd4, 0x24, 0xef, 0xc1, 0x93, 0x46, 0xca, 0xa9, 0x1a, 0x5f,
	0x3c, 0xa6, 0x26, 0xab, 0x79, 0xe6, 0xe2, 0xe9, 0x6b, 0xa8, 0x1d, 0x5c, 0x4b, 0x7b, 0x2b, 0xb2,
	0x0e, 0x10, 0x09, 0x8d, 0x72, 0xc4, 0x07, 0xdf, 0x5c, 0x85, 0x32, 0xcb, 0x78, 0xe8, 0x0e, 0x34,
	0x4e, 0x23, 0xd1, 0xbf, 0x13, 0xc5, 0x0a, 0x78, 0x28, 0xb4, 0xbc, 0x4d, 0x43, 0x9d, 0x61, 0x64,
	0x86, 0xe3, 0xc8, 0x09, 0xaa, 0xcc, 0xec, 0xf7, 0xd6, 0x1f, 0x0f, 0x96, 0xf7, 0xdd, 0xfc, 0xf6,
	0xc6, 0x5d, 0x2d, 0x91, 0x0f, 0x51, 0x92, 0x1e, 0x2c, 0x1e, 0xa2, 0x3e, 0xe6, 0x1a, 0x95, 0xb6,
	0x9d, 0x91, 0x76, 0x4e, 0xdf, 0x77, 0xca, 0x69, 0x4d, 0x99, 0x13, 0x5a, 0x20, 0x5f, 0xa1, 0x76,
	0x88, 0x29, 0xde, 0x94, 0xe8, 0xd6, 0xcb, 0xbc, 0x7a, 0xae, 0x57, 0x1b, 0x46, 0x0b, 0xe4, 0x3b,
	0x2c, 0x4c, 0x20, 0xdd, 0xc2, 0x98, 0xce, 0xef, 0x8c, 0xd0, 0x9b, 0x45, 0x72, 0x66, 0x59, 0xc8,
	0x0e, 0xea, 0xf3, 0x9c, 0xd4, 0xc9, 0xee, 0x68, 0xbd, 0xca, 0x09, 0x78, 0x38, 0xf0, 0xb4, 0x40,
	0xce, 0x61, 0xc9, 0x8c, 0x71, 0x16, 0x7c, 0xb6, 0xdc, 0xdc, 0xf6, 0xb3, 0x5b, 0x81, 0x16, 0x88,
	0x84, 0xa5, 0x43, 0x9c, 0x48, 0xb5, 0x37, 0x8e, 0x02, 0x45, 0xb6, 0xf3, 0xba, 0x7f, 0x4c, 0xda,
	0x33, 0x5f, 0x69, 0xb3, 0x48, 0x98, 0x7d, 0x8d, 0xcc, 0xd6, 0x58, 0xcb, 0xc9, 0xb5, 0x2b, 0xa6,
	0x95, 0xf7, 0x56, 0xf7, 0x00, 0xb4, 0x40, 0x4e, 0xa0, 0x62, 0xa4, 0x9d, 0x4b, 0xfd, 0x64, 0x46,
	0x72, 0x79, 0xc9, 0x0e, 0x06, 0x2d, 0xec, 0x2d, 0x9e, 0xd5, 0x5d, 0x80, 0x4c, 0xfc, 0xdf, 0xa5,
	0xc2, 0xc5, 0x9c, 0xfd, 0x5b, 0xbd, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xd2, 0xc0, 0xf7,
	0xec, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompactTxStreamerClient is the client API for CompactTxStreamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompactTxStreamerClient interface {
	// Compact Blocks
	GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error)
	GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error)
	GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error)
	// Transactions
	GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error)
	SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error)
	// t-Address support
	GetAddressTxids(ctx context.Context, in *TransparentAddressBlockFilter, opts ...grpc.CallOption) (CompactTxStreamer_GetAddressTxidsClient, error)
	// Misc
	GetLightdInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LightdInfo, error)
	Ping(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*PingResponse, error)
}

type compactTxStreamerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompactTxStreamerClient(cc grpc.ClientConnInterface) CompactTxStreamerClient {
	return &compactTxStreamerClient{cc}
}

func (c *compactTxStreamerClient) GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error) {
	out := new(BlockID)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error) {
	out := new(CompactBlock)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompactTxStreamer_serviceDesc.Streams[0], "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlockRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetBlockRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetBlockRangeClient interface {
	Recv() (*CompactBlock, error)
	grpc.ClientStream
}

type compactTxStreamerGetBlockRangeClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetBlockRangeClient) Recv() (*CompactBlock, error) {
	m := new(CompactBlock)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error) {
	out := new(RawTransaction)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetAddressTxids(ctx context.Context, in *TransparentAddressBlockFilter, opts ...grpc.CallOption) (CompactTxStreamer_GetAddressTxidsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompactTxStreamer_serviceDesc.Streams[1], "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetAddressTxids", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetAddressTxidsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetAddressTxidsClient interface {
	Recv() (*RawTransaction, error)
	grpc.ClientStream
}

type compactTxStreamerGetAddressTxidsClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetAddressTxidsClient) Recv() (*RawTransaction, error) {
	m := new(RawTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetLightdInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LightdInfo, error) {
	out := new(LightdInfo)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLightdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) Ping(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompactTxStreamerServer is the server API for CompactTxStreamer service.
type CompactTxStreamerServer interface {
	// Compact Blocks
	GetLatestBlock(context.Context, *ChainSpec) (*BlockID, error)
	GetBlock(context.Context, *BlockID) (*CompactBlock, error)
	GetBlockRange(*BlockRange, CompactTxStreamer_GetBlockRangeServer) error
	// Transactions
	GetTransaction(context.Context, *TxFilter) (*RawTransaction, error)
	SendTransaction(context.Context, *RawTransaction) (*SendResponse, error)
	// t-Address support
	GetAddressTxids(*TransparentAddressBlockFilter, CompactTxStreamer_GetAddressTxidsServer) error
	// Misc
	GetLightdInfo(context.Context, *Empty) (*LightdInfo, error)
	Ping(context.Context, *Duration) (*PingResponse, error)
}

// UnimplementedCompactTxStreamerServer can be embedded to have forward compatible implementations.
type UnimplementedCompactTxStreamerServer struct {
}

func (*UnimplementedCompactTxStreamerServer) GetLatestBlock(ctx context.Context, req *ChainSpec) (*BlockID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetBlock(ctx context.Context, req *BlockID) (*CompactBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetBlockRange(req *BlockRange, srv CompactTxStreamer_GetBlockRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockRange not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetTransaction(ctx context.Context, req *TxFilter) (*RawTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (*UnimplementedCompactTxStreamerServer) SendTransaction(ctx context.Context, req *RawTransaction) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetAddressTxids(req *TransparentAddressBlockFilter, srv CompactTxStreamer_GetAddressTxidsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAddressTxids not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetLightdInfo(ctx context.Context, req *Empty) (*LightdInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLightdInfo not implemented")
}
func (*UnimplementedCompactTxStreamerServer) Ping(ctx context.Context, req *Duration) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterCompactTxStreamerServer(s *grpc.Server, srv CompactTxStreamerServer) {
	s.RegisterService(&_CompactTxStreamer_serviceDesc, srv)
}

func _CompactTxStreamer_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, req.(*ChainSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, req.(*BlockID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlockRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetBlockRange(m, &compactTxStreamerGetBlockRangeServer{stream})
}

type CompactTxStreamer_GetBlockRangeServer interface {
	Send(*CompactBlock) error
	grpc.ServerStream
}

type compactTxStreamerGetBlockRangeServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetBlockRangeServer) Send(m *CompactBlock) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, req.(*TxFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, req.(*RawTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetAddressTxids_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransparentAddressBlockFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetAddressTxids(m, &compactTxStreamerGetAddressTxidsServer{stream})
}

type CompactTxStreamer_GetAddressTxidsServer interface {
	Send(*RawTransaction) error
	grpc.ServerStream
}

type compactTxStreamerGetAddressTxidsServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetAddressTxidsServer) Send(m *RawTransaction) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetLightdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLightdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLightdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLightdInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Duration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).Ping(ctx, req.(*Duration))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompactTxStreamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cash.z.wallet.sdk.rpc.CompactTxStreamer",
	HandlerType: (*CompactTxStreamerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestBlock",
			Handler:    _CompactTxStreamer_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _CompactTxStreamer_GetBlock_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _CompactTxStreamer_GetTransaction_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _CompactTxStreamer_SendTransaction_Handler,
		},
		{
			MethodName: "GetLightdInfo",
			Handler:    _CompactTxStreamer_GetLightdInfo_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CompactTxStreamer_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlockRange",
			Handler:       _CompactTxStreamer_GetBlockRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAddressTxids",
			Handler:       _CompactTxStreamer_GetAddressTxids_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}

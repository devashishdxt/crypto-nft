// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cryptonft/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgNewClass struct {
	// creator defines creator of NFT classification
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// id defines the unique identifier of the NFT classification, similar to the contract address of ERC721
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// name defines the human-readable name of the NFT classification. Optional
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// symbol is an abbreviated name for nft classification. Optional
	Symbol string `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// description is a brief description of nft classification. Optional
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// uri for the class metadata stored off chain. It can define schema for Class and NFT `Data` attributes. Optional
	Uri string `protobuf:"bytes,6,opt,name=uri,proto3" json:"uri,omitempty"`
	// uri_hash is a hash of the document pointed by uri. Optional
	UriHash string `protobuf:"bytes,7,opt,name=uriHash,proto3" json:"uriHash,omitempty"`
	// mintRestricted specifies whether NFT classification is restricts minting NFTs by others
	MintRestricted bool `protobuf:"varint,8,opt,name=mintRestricted,proto3" json:"mintRestricted,omitempty"`
	// data is the app specific metadata of the NFT class. Optional
	Data *types.Any `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgNewClass) Reset()         { *m = MsgNewClass{} }
func (m *MsgNewClass) String() string { return proto.CompactTextString(m) }
func (*MsgNewClass) ProtoMessage()    {}
func (*MsgNewClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7f3988a15c0fee9, []int{0}
}
func (m *MsgNewClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewClass.Merge(m, src)
}
func (m *MsgNewClass) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewClass) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewClass.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewClass proto.InternalMessageInfo

func (m *MsgNewClass) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgNewClass) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgNewClass) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgNewClass) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgNewClass) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgNewClass) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *MsgNewClass) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

func (m *MsgNewClass) GetMintRestricted() bool {
	if m != nil {
		return m.MintRestricted
	}
	return false
}

func (m *MsgNewClass) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type MsgNewClassResponse struct {
}

func (m *MsgNewClassResponse) Reset()         { *m = MsgNewClassResponse{} }
func (m *MsgNewClassResponse) String() string { return proto.CompactTextString(m) }
func (*MsgNewClassResponse) ProtoMessage()    {}
func (*MsgNewClassResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7f3988a15c0fee9, []int{1}
}
func (m *MsgNewClassResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgNewClassResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgNewClassResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgNewClassResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgNewClassResponse.Merge(m, src)
}
func (m *MsgNewClassResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgNewClassResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgNewClassResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgNewClassResponse proto.InternalMessageInfo

type MsgMintNFT struct {
	// creator defines creator of NFT
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// class_id associated with the NFT, similar to the contract address of ERC721
	ClassId string `protobuf:"bytes,2,opt,name=classId,proto3" json:"classId,omitempty"`
	// id is a unique identifier of the NFT
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// uri for the NFT metadata stored off chain
	Uri string `protobuf:"bytes,4,opt,name=uri,proto3" json:"uri,omitempty"`
	// uri_hash is a hash of the document pointed by uri
	UriHash string `protobuf:"bytes,5,opt,name=uriHash,proto3" json:"uriHash,omitempty"`
	// receiver is the address of the receiver of the NFT
	Receiver string `protobuf:"bytes,6,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// data is an app specific data of the NFT. Optional
	Data *types.Any `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgMintNFT) Reset()         { *m = MsgMintNFT{} }
func (m *MsgMintNFT) String() string { return proto.CompactTextString(m) }
func (*MsgMintNFT) ProtoMessage()    {}
func (*MsgMintNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7f3988a15c0fee9, []int{2}
}
func (m *MsgMintNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintNFT.Merge(m, src)
}
func (m *MsgMintNFT) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintNFT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintNFT proto.InternalMessageInfo

func (m *MsgMintNFT) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgMintNFT) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *MsgMintNFT) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgMintNFT) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *MsgMintNFT) GetUriHash() string {
	if m != nil {
		return m.UriHash
	}
	return ""
}

func (m *MsgMintNFT) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgMintNFT) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type MsgMintNFTResponse struct {
}

func (m *MsgMintNFTResponse) Reset()         { *m = MsgMintNFTResponse{} }
func (m *MsgMintNFTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintNFTResponse) ProtoMessage()    {}
func (*MsgMintNFTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7f3988a15c0fee9, []int{3}
}
func (m *MsgMintNFTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintNFTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintNFTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintNFTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintNFTResponse.Merge(m, src)
}
func (m *MsgMintNFTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintNFTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintNFTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintNFTResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgNewClass)(nil), "devashishdxt.cryptonft.cryptonft.MsgNewClass")
	proto.RegisterType((*MsgNewClassResponse)(nil), "devashishdxt.cryptonft.cryptonft.MsgNewClassResponse")
	proto.RegisterType((*MsgMintNFT)(nil), "devashishdxt.cryptonft.cryptonft.MsgMintNFT")
	proto.RegisterType((*MsgMintNFTResponse)(nil), "devashishdxt.cryptonft.cryptonft.MsgMintNFTResponse")
}

func init() { proto.RegisterFile("cryptonft/tx.proto", fileDescriptor_f7f3988a15c0fee9) }

var fileDescriptor_f7f3988a15c0fee9 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xb1, 0x8e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0xb6, 0xd7, 0xe6, 0xbe, 0x4a, 0x27, 0x64, 0x0e, 0x64, 0x3a, 0x44, 0x51, 0x07,
	0xd4, 0x81, 0x4b, 0xa4, 0xe3, 0x78, 0x00, 0x40, 0x42, 0x30, 0xb4, 0x43, 0xc4, 0xc4, 0xe6, 0x26,
	0xbe, 0xd4, 0x52, 0x63, 0x47, 0xb6, 0x7b, 0x34, 0x33, 0x2f, 0xc0, 0x33, 0x31, 0x31, 0xde, 0xc8,
	0x88, 0x5a, 0x89, 0xe7, 0x40, 0x71, 0xe2, 0x5c, 0x40, 0x88, 0xbb, 0x6e, 0xfe, 0xfe, 0xdf, 0xdf,
	0x9f, 0xfd, 0xff, 0x25, 0x06, 0x9c, 0xa8, 0xb2, 0x30, 0x52, 0x5c, 0x9b, 0xc8, 0xec, 0xc2, 0x42,
	0x49, 0x23, 0x71, 0x90, 0xb2, 0x1b, 0xaa, 0xd7, 0x5c, 0xaf, 0xd3, 0x9d, 0x09, 0x5b, 0xc3, 0xdd,
	0x6a, 0xfa, 0x2c, 0x93, 0x32, 0xdb, 0xb0, 0xc8, 0xfa, 0x57, 0xdb, 0xeb, 0x88, 0x8a, 0xb2, 0xde,
	0x3c, 0xfb, 0xd2, 0x87, 0xc9, 0x42, 0x67, 0x4b, 0xf6, 0xf9, 0xed, 0x86, 0x6a, 0x8d, 0x09, 0x8c,
	0x13, 0xc5, 0xa8, 0x91, 0x8a, 0xa0, 0x00, 0xcd, 0x4f, 0x63, 0x57, 0xe2, 0x33, 0xe8, 0xf3, 0x94,
	0xf4, 0xad, 0xd8, 0xe7, 0x29, 0xc6, 0x30, 0x14, 0x34, 0x67, 0x64, 0x60, 0x15, 0xbb, 0xc6, 0x4f,
	0x61, 0xa4, 0xcb, 0x7c, 0x25, 0x37, 0x64, 0x68, 0xd5, 0xa6, 0xc2, 0x01, 0x4c, 0x52, 0xa6, 0x13,
	0xc5, 0x0b, 0xc3, 0xa5, 0x20, 0x27, 0xb6, 0xd9, 0x95, 0xf0, 0x23, 0x18, 0x6c, 0x15, 0x27, 0x23,
	0xdb, 0xa9, 0x96, 0xd5, 0x4d, 0xb6, 0x8a, 0xbf, 0xa7, 0x7a, 0x4d, 0xc6, 0xf5, 0x4d, 0x9a, 0x12,
	0x3f, 0x87, 0xb3, 0x9c, 0x0b, 0x13, 0x33, 0x6d, 0x14, 0x4f, 0x0c, 0x4b, 0x89, 0x17, 0xa0, 0xb9,
	0x17, 0xff, 0xa5, 0xe2, 0x39, 0x0c, 0x53, 0x6a, 0x28, 0x39, 0x0d, 0xd0, 0x7c, 0x72, 0x79, 0x1e,
	0xd6, 0x14, 0x42, 0x47, 0x21, 0x7c, 0x2d, 0xca, 0xd8, 0x3a, 0x66, 0x4f, 0xe0, 0x71, 0x07, 0x42,
	0xcc, 0x74, 0x21, 0x85, 0x66, 0xb3, 0x6f, 0x08, 0x60, 0xa1, 0xb3, 0x05, 0x17, 0x66, 0xf9, 0xee,
	0xe3, 0x7f, 0xd8, 0x54, 0x9d, 0x6a, 0xe7, 0x07, 0x07, 0xc8, 0x95, 0x0d, 0xb5, 0x41, 0x4b, 0xad,
	0xc9, 0x39, 0xfc, 0x67, 0xce, 0x93, 0x3f, 0x73, 0x4e, 0xc1, 0x53, 0x2c, 0x61, 0xfc, 0x86, 0xa9,
	0x06, 0x4c, 0x5b, 0xb7, 0xd9, 0xbc, 0x7b, 0xb3, 0x9d, 0x03, 0xbe, 0xcb, 0xe0, 0xa2, 0x5d, 0xfe,
	0x42, 0x30, 0x58, 0xe8, 0x0c, 0x17, 0xe0, 0xb5, 0xdf, 0xfe, 0x22, 0xbc, 0xef, 0x4f, 0x0a, 0x3b,
	0x94, 0xa6, 0xaf, 0x8e, 0xb2, 0xbb, 0x93, 0x71, 0x0e, 0x63, 0x07, 0xf4, 0xc5, 0x83, 0x26, 0x34,
	0xee, 0xe9, 0xd5, 0x31, 0x6e, 0x77, 0xdc, 0x9b, 0xe5, 0xf7, 0xbd, 0x8f, 0x6e, 0xf7, 0x3e, 0xfa,
	0xb9, 0xf7, 0xd1, 0xd7, 0x83, 0xdf, 0xbb, 0x3d, 0xf8, 0xbd, 0x1f, 0x07, 0xbf, 0xf7, 0xe9, 0x2a,
	0xe3, 0x66, 0xbd, 0x5d, 0x85, 0x89, 0xcc, 0xa3, 0xee, 0xe4, 0xa8, 0x9e, 0x77, 0x51, 0x3d, 0xb2,
	0x5d, 0xd4, 0x79, 0x70, 0x65, 0xc1, 0xf4, 0x6a, 0x64, 0x11, 0xbf, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0x09, 0x75, 0x4e, 0x8a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// NewClass creates a new NFT classification.
	NewClass(ctx context.Context, in *MsgNewClass, opts ...grpc.CallOption) (*MsgNewClassResponse, error)
	// MintNFT mints a new NFT.
	MintNFT(ctx context.Context, in *MsgMintNFT, opts ...grpc.CallOption) (*MsgMintNFTResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) NewClass(ctx context.Context, in *MsgNewClass, opts ...grpc.CallOption) (*MsgNewClassResponse, error) {
	out := new(MsgNewClassResponse)
	err := c.cc.Invoke(ctx, "/devashishdxt.cryptonft.cryptonft.Msg/NewClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintNFT(ctx context.Context, in *MsgMintNFT, opts ...grpc.CallOption) (*MsgMintNFTResponse, error) {
	out := new(MsgMintNFTResponse)
	err := c.cc.Invoke(ctx, "/devashishdxt.cryptonft.cryptonft.Msg/MintNFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// NewClass creates a new NFT classification.
	NewClass(context.Context, *MsgNewClass) (*MsgNewClassResponse, error)
	// MintNFT mints a new NFT.
	MintNFT(context.Context, *MsgMintNFT) (*MsgMintNFTResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) NewClass(ctx context.Context, req *MsgNewClass) (*MsgNewClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewClass not implemented")
}
func (*UnimplementedMsgServer) MintNFT(ctx context.Context, req *MsgMintNFT) (*MsgMintNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintNFT not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_NewClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewClass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devashishdxt.cryptonft.cryptonft.Msg/NewClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewClass(ctx, req.(*MsgNewClass))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintNFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintNFT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintNFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devashishdxt.cryptonft.cryptonft.Msg/MintNFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintNFT(ctx, req.(*MsgMintNFT))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "devashishdxt.cryptonft.cryptonft.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewClass",
			Handler:    _Msg_NewClass_Handler,
		},
		{
			MethodName: "MintNFT",
			Handler:    _Msg_MintNFT_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cryptonft/tx.proto",
}

func (m *MsgNewClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.MintRestricted {
		i--
		if m.MintRestricted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgNewClassResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgNewClassResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgNewClassResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgMintNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.UriHash) > 0 {
		i -= len(m.UriHash)
		copy(dAtA[i:], m.UriHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.UriHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintNFTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintNFTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintNFTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgNewClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.MintRestricted {
		n += 2
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgNewClassResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgMintNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.UriHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgMintNFTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgNewClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNewClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRestricted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.MintRestricted = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgNewClassResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgNewClassResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgNewClassResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMintNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMintNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UriHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UriHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMintNFTResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMintNFTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintNFTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

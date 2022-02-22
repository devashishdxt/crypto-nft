// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: compatnft/compatnft.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type DenomMetadata struct {
	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (m *DenomMetadata) Reset()         { *m = DenomMetadata{} }
func (m *DenomMetadata) String() string { return proto.CompactTextString(m) }
func (*DenomMetadata) ProtoMessage()    {}
func (*DenomMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_104575bc8f2e58e0, []int{0}
}
func (m *DenomMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomMetadata.Merge(m, src)
}
func (m *DenomMetadata) XXX_Size() int {
	return m.Size()
}
func (m *DenomMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DenomMetadata proto.InternalMessageInfo

func (m *DenomMetadata) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

type NFTMetadata struct {
	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *NFTMetadata) Reset()         { *m = NFTMetadata{} }
func (m *NFTMetadata) String() string { return proto.CompactTextString(m) }
func (*NFTMetadata) ProtoMessage()    {}
func (*NFTMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_104575bc8f2e58e0, []int{1}
}
func (m *NFTMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NFTMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NFTMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NFTMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NFTMetadata.Merge(m, src)
}
func (m *NFTMetadata) XXX_Size() int {
	return m.Size()
}
func (m *NFTMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_NFTMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_NFTMetadata proto.InternalMessageInfo

func (m *NFTMetadata) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*DenomMetadata)(nil), "devashishdxt.cryptonft.compatnft.DenomMetadata")
	proto.RegisterType((*NFTMetadata)(nil), "devashishdxt.cryptonft.compatnft.NFTMetadata")
}

func init() { proto.RegisterFile("compatnft/compatnft.proto", fileDescriptor_104575bc8f2e58e0) }

var fileDescriptor_104575bc8f2e58e0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xce, 0xcf, 0x2d,
	0x48, 0x2c, 0xc9, 0x4b, 0x2b, 0xd1, 0x87, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x14,
	0x52, 0x52, 0xcb, 0x12, 0x8b, 0x33, 0x32, 0x8b, 0x33, 0x52, 0x2a, 0x4a, 0xf4, 0x92, 0x8b, 0x2a,
	0x0b, 0x4a, 0xf2, 0x41, 0xb2, 0x70, 0x75, 0x4a, 0xea, 0x5c, 0xbc, 0x2e, 0xa9, 0x79, 0xf9, 0xb9,
	0xbe, 0xa9, 0x25, 0x89, 0x29, 0x89, 0x25, 0x89, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xc9, 0x19, 0xa9,
	0xb9, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x22, 0x17, 0xb7, 0x9f,
	0x5b, 0x08, 0x5c, 0x99, 0x10, 0x17, 0x0b, 0x88, 0x86, 0x2a, 0x02, 0xb3, 0x9d, 0xfc, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e,
	0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x09,
	0x64, 0xaf, 0x3e, 0xb2, 0x93, 0xf4, 0x21, 0x4e, 0xd2, 0x05, 0xb9, 0xbd, 0x02, 0xe1, 0x7a, 0xfd,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x27, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x80, 0x3b, 0xfb, 0x65, 0xe1, 0x00, 0x00, 0x00,
}

func (m *DenomMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintCompatnft(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NFTMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NFTMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NFTMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintCompatnft(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCompatnft(dAtA []byte, offset int, v uint64) int {
	offset -= sovCompatnft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovCompatnft(uint64(l))
	}
	return n
}

func (m *NFTMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovCompatnft(uint64(l))
	}
	return n
}

func sovCompatnft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCompatnft(x uint64) (n int) {
	return sovCompatnft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompatnft
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
			return fmt.Errorf("proto: DenomMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompatnft
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
				return ErrInvalidLengthCompatnft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompatnft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompatnft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompatnft
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
func (m *NFTMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCompatnft
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
			return fmt.Errorf("proto: NFTMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NFTMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCompatnft
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
				return ErrInvalidLengthCompatnft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCompatnft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCompatnft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCompatnft
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
func skipCompatnft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCompatnft
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
					return 0, ErrIntOverflowCompatnft
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
					return 0, ErrIntOverflowCompatnft
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
				return 0, ErrInvalidLengthCompatnft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCompatnft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCompatnft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCompatnft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCompatnft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCompatnft = fmt.Errorf("proto: unexpected end of group")
)

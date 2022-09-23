// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/accesscontrol.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type AccessOperation struct {
	AccessType         AccessType   `protobuf:"varint,1,opt,name=access_type,json=accessType,proto3,enum=cosmos.accesscontrol.v1beta1.AccessType" json:"access_type,omitempty"`
	ResourceType       ResourceType `protobuf:"varint,2,opt,name=resource_type,json=resourceType,proto3,enum=cosmos.accesscontrol.v1beta1.ResourceType" json:"resource_type,omitempty"`
	IdentifierTemplate string       `protobuf:"bytes,3,opt,name=identifier_template,json=identifierTemplate,proto3" json:"identifier_template,omitempty"`
}

func (m *AccessOperation) Reset()         { *m = AccessOperation{} }
func (m *AccessOperation) String() string { return proto.CompactTextString(m) }
func (*AccessOperation) ProtoMessage()    {}
func (*AccessOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{0}
}
func (m *AccessOperation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccessOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccessOperation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccessOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessOperation.Merge(m, src)
}
func (m *AccessOperation) XXX_Size() int {
	return m.Size()
}
func (m *AccessOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AccessOperation proto.InternalMessageInfo

func (m *AccessOperation) GetAccessType() AccessType {
	if m != nil {
		return m.AccessType
	}
	return AccessType_UNKNOWN
}

func (m *AccessOperation) GetResourceType() ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return ResourceType_ANY
}

func (m *AccessOperation) GetIdentifierTemplate() string {
	if m != nil {
		return m.IdentifierTemplate
	}
	return ""
}

type MessageDependencyMapping struct {
	MessageType MessageType       `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=cosmos.accesscontrol.v1beta1.MessageType" json:"message_type,omitempty"`
	ModuleName  string            `protobuf:"bytes,2,opt,name=module_name,json=moduleName,proto3" json:"module_name,omitempty"`
	AccessOps   []AccessOperation `protobuf:"bytes,3,rep,name=access_ops,json=accessOps,proto3" json:"access_ops"`
}

func (m *MessageDependencyMapping) Reset()         { *m = MessageDependencyMapping{} }
func (m *MessageDependencyMapping) String() string { return proto.CompactTextString(m) }
func (*MessageDependencyMapping) ProtoMessage()    {}
func (*MessageDependencyMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_d636a082612ba091, []int{1}
}
func (m *MessageDependencyMapping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MessageDependencyMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MessageDependencyMapping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MessageDependencyMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageDependencyMapping.Merge(m, src)
}
func (m *MessageDependencyMapping) XXX_Size() int {
	return m.Size()
}
func (m *MessageDependencyMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageDependencyMapping.DiscardUnknown(m)
}

var xxx_messageInfo_MessageDependencyMapping proto.InternalMessageInfo

func (m *MessageDependencyMapping) GetMessageType() MessageType {
	if m != nil {
		return m.MessageType
	}
	return MessageType_ANY_MESSAGE
}

func (m *MessageDependencyMapping) GetModuleName() string {
	if m != nil {
		return m.ModuleName
	}
	return ""
}

func (m *MessageDependencyMapping) GetAccessOps() []AccessOperation {
	if m != nil {
		return m.AccessOps
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessOperation)(nil), "cosmos.accesscontrol.v1beta1.AccessOperation")
	proto.RegisterType((*MessageDependencyMapping)(nil), "cosmos.accesscontrol.v1beta1.MessageDependencyMapping")
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/accesscontrol.proto", fileDescriptor_d636a082612ba091)
}

var fileDescriptor_d636a082612ba091 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0xb5, 0x75, 0x29, 0x78, 0xe5, 0xb6, 0xb0, 0xed, 0x41, 0x98, 0x22, 0x1b, 0xd3, 0x83,
	0x5a, 0xb0, 0x84, 0xdd, 0x27, 0xa8, 0xc9, 0x25, 0x60, 0xc7, 0x20, 0x7c, 0xca, 0x45, 0xac, 0x57,
	0x13, 0x45, 0xc4, 0xfb, 0x81, 0x76, 0x1d, 0xe2, 0xb7, 0xc8, 0x63, 0xf9, 0xe8, 0x63, 0x2e, 0x31,
	0xc1, 0x7e, 0x91, 0x60, 0xad, 0xf0, 0x47, 0x08, 0x26, 0x27, 0xed, 0xcc, 0xff, 0x3f, 0x3f, 0x66,
	0x46, 0x83, 0x03, 0x26, 0x35, 0x97, 0x3a, 0xa2, 0x8c, 0x81, 0xd6, 0x4c, 0x0a, 0x53, 0xc8, 0xd9,
	0x69, 0x14, 0xaa, 0x42, 0x1a, 0x49, 0x7e, 0x59, 0x67, 0x78, 0xaa, 0xdd, 0xf7, 0xa6, 0x60, 0x68,
	0xaf, 0xf9, 0x33, 0x93, 0x99, 0x2c, 0x8d, 0xd1, 0xee, 0x65, 0x6b, 0x9a, 0xbf, 0xdf, 0xa5, 0x33,
	0x29, 0xb4, 0xa1, 0xc2, 0x68, 0xeb, 0xea, 0xac, 0x11, 0xfe, 0xfe, 0xbf, 0x74, 0x8c, 0x15, 0x14,
	0xd4, 0xe4, 0x52, 0x90, 0x4b, 0xec, 0xda, 0xa2, 0xc4, 0x2c, 0x14, 0x78, 0xa8, 0x8d, 0x82, 0x6f,
	0xfd, 0x20, 0x3c, 0xd7, 0x43, 0x68, 0x19, 0x93, 0x85, 0x82, 0x18, 0xd3, 0xfd, 0x9b, 0x8c, 0xf1,
	0xd7, 0x02, 0xb4, 0x9c, 0x17, 0x0c, 0x2c, 0xec, 0x53, 0x09, 0xfb, 0x7b, 0x1e, 0x16, 0x57, 0x25,
	0x25, 0xae, 0x51, 0x1c, 0x45, 0x24, 0xc2, 0x3f, 0xf2, 0x14, 0x84, 0xc9, 0x6f, 0x72, 0x28, 0x12,
	0x03, 0x5c, 0xcd, 0xa8, 0x01, 0xaf, 0xd6, 0x46, 0x41, 0x3d, 0x26, 0x07, 0x69, 0x52, 0x29, 0x9d,
	0x67, 0x84, 0xbd, 0x11, 0x68, 0x4d, 0x33, 0xb8, 0x00, 0x05, 0x22, 0x05, 0xc1, 0x16, 0x23, 0xaa,
	0x54, 0x2e, 0x32, 0x32, 0xc4, 0x0d, 0x6e, 0xb5, 0xe3, 0x51, 0xff, 0x9c, 0xef, 0xae, 0xa2, 0x95,
	0xcd, 0xb9, 0xfc, 0x10, 0x90, 0x16, 0x76, 0xb9, 0x4c, 0xe7, 0x33, 0x48, 0x04, 0xe5, 0x76, 0xd4,
	0x7a, 0x8c, 0x6d, 0xea, 0x8a, 0x72, 0x20, 0x31, 0xae, 0x76, 0x93, 0x48, 0xa5, 0xbd, 0x5a, 0xbb,
	0x16, 0xb8, 0xfd, 0xee, 0x47, 0xf6, 0xba, 0xff, 0x37, 0x83, 0xcf, 0xcb, 0x75, 0xcb, 0x89, 0xeb,
	0xb4, 0x4a, 0xeb, 0xc1, 0x70, 0xb9, 0xf1, 0xd1, 0x6a, 0xe3, 0xa3, 0x97, 0x8d, 0x8f, 0x1e, 0xb7,
	0xbe, 0xb3, 0xda, 0xfa, 0xce, 0xd3, 0xd6, 0x77, 0xae, 0xfb, 0x59, 0x6e, 0x6e, 0xe7, 0xd3, 0x90,
	0x49, 0x1e, 0x55, 0xb7, 0x60, 0x3f, 0x5d, 0x9d, 0xde, 0x45, 0x0f, 0x6f, 0x0e, 0x63, 0x37, 0xbf,
	0x9e, 0x7e, 0x29, 0xaf, 0xe2, 0xdf, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x25, 0x6a, 0x8a,
	0x9b, 0x02, 0x00, 0x00,
}

func (m *AccessOperation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccessOperation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccessOperation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IdentifierTemplate) > 0 {
		i -= len(m.IdentifierTemplate)
		copy(dAtA[i:], m.IdentifierTemplate)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.IdentifierTemplate)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ResourceType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.ResourceType))
		i--
		dAtA[i] = 0x10
	}
	if m.AccessType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.AccessType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MessageDependencyMapping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MessageDependencyMapping) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MessageDependencyMapping) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccessOps) > 0 {
		for iNdEx := len(m.AccessOps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccessOps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccesscontrol(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ModuleName) > 0 {
		i -= len(m.ModuleName)
		copy(dAtA[i:], m.ModuleName)
		i = encodeVarintAccesscontrol(dAtA, i, uint64(len(m.ModuleName)))
		i--
		dAtA[i] = 0x12
	}
	if m.MessageType != 0 {
		i = encodeVarintAccesscontrol(dAtA, i, uint64(m.MessageType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccesscontrol(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccesscontrol(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccessOperation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccessType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.AccessType))
	}
	if m.ResourceType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.ResourceType))
	}
	l = len(m.IdentifierTemplate)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	return n
}

func (m *MessageDependencyMapping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MessageType != 0 {
		n += 1 + sovAccesscontrol(uint64(m.MessageType))
	}
	l = len(m.ModuleName)
	if l > 0 {
		n += 1 + l + sovAccesscontrol(uint64(l))
	}
	if len(m.AccessOps) > 0 {
		for _, e := range m.AccessOps {
			l = e.Size()
			n += 1 + l + sovAccesscontrol(uint64(l))
		}
	}
	return n
}

func sovAccesscontrol(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccesscontrol(x uint64) (n int) {
	return sovAccesscontrol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccessOperation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: AccessOperation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccessOperation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessType", wireType)
			}
			m.AccessType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccessType |= AccessType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceType", wireType)
			}
			m.ResourceType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResourceType |= ResourceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentifierTemplate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentifierTemplate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func (m *MessageDependencyMapping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccesscontrol
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
			return fmt.Errorf("proto: MessageDependencyMapping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageDependencyMapping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageType", wireType)
			}
			m.MessageType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageType |= MessageType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessOps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccesscontrol
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
				return ErrInvalidLengthAccesscontrol
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccesscontrol
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessOps = append(m.AccessOps, AccessOperation{})
			if err := m.AccessOps[len(m.AccessOps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccesscontrol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccesscontrol
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
func skipAccesscontrol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccesscontrol
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
					return 0, ErrIntOverflowAccesscontrol
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
					return 0, ErrIntOverflowAccesscontrol
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
				return 0, ErrInvalidLengthAccesscontrol
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccesscontrol
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccesscontrol
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccesscontrol        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccesscontrol          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccesscontrol = fmt.Errorf("proto: unexpected end of group")
)
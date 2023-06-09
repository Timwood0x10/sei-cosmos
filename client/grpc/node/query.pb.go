// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/node/v1beta1/query.proto

package node

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ConfigRequest defines the request structure for the Config gRPC query.
type ConfigRequest struct {
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8324226a07064341, []int{0}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

// ConfigResponse defines the response structure for the Config gRPC query.
type ConfigResponse struct {
	MinimumGasPrice string `protobuf:"bytes,1,opt,name=minimum_gas_price,json=minimumGasPrice,proto3" json:"minimum_gas_price,omitempty"`
	// pruning settings
	PruningKeepRecent string `protobuf:"bytes,2,opt,name=pruning_keep_recent,json=pruningKeepRecent,proto3" json:"pruning_keep_recent,omitempty"`
	PruningInterval   string `protobuf:"bytes,3,opt,name=pruning_interval,json=pruningInterval,proto3" json:"pruning_interval,omitempty"`
}

func (m *ConfigResponse) Reset()         { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()    {}
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8324226a07064341, []int{1}
}
func (m *ConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigResponse.Merge(m, src)
}
func (m *ConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigResponse proto.InternalMessageInfo

func (m *ConfigResponse) GetMinimumGasPrice() string {
	if m != nil {
		return m.MinimumGasPrice
	}
	return ""
}

func (m *ConfigResponse) GetPruningKeepRecent() string {
	if m != nil {
		return m.PruningKeepRecent
	}
	return ""
}

func (m *ConfigResponse) GetPruningInterval() string {
	if m != nil {
		return m.PruningInterval
	}
	return ""
}

// StateRequest defines the request structure for the status of a node.
type StatusRequest struct {
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8324226a07064341, []int{2}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return m.Size()
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

// StateResponse defines the response structure for the status of a node.
type StatusResponse struct {
	EarliestStoreHeight uint64     `protobuf:"varint,1,opt,name=earliest_store_height,json=earliestStoreHeight,proto3" json:"earliest_store_height,omitempty"`
	Height              uint64     `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp           *time.Time `protobuf:"bytes,3,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	AppHash             []byte     `protobuf:"bytes,4,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	ValidatorHash       []byte     `protobuf:"bytes,5,opt,name=validator_hash,json=validatorHash,proto3" json:"validator_hash,omitempty"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8324226a07064341, []int{3}
}
func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return m.Size()
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetEarliestStoreHeight() uint64 {
	if m != nil {
		return m.EarliestStoreHeight
	}
	return 0
}

func (m *StatusResponse) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *StatusResponse) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *StatusResponse) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *StatusResponse) GetValidatorHash() []byte {
	if m != nil {
		return m.ValidatorHash
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "cosmos.base.node.v1beta1.ConfigRequest")
	proto.RegisterType((*ConfigResponse)(nil), "cosmos.base.node.v1beta1.ConfigResponse")
	proto.RegisterType((*StatusRequest)(nil), "cosmos.base.node.v1beta1.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "cosmos.base.node.v1beta1.StatusResponse")
}

func init() {
	proto.RegisterFile("cosmos/base/node/v1beta1/query.proto", fileDescriptor_8324226a07064341)
}

var fileDescriptor_8324226a07064341 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6f, 0xd3, 0x3e,
	0x1c, 0xc5, 0x9b, 0xfe, 0xfa, 0xeb, 0x98, 0x61, 0x1d, 0xcb, 0x00, 0x95, 0x0a, 0x65, 0x55, 0x05,
	0xa2, 0x20, 0xcd, 0xd6, 0xca, 0x9d, 0xc3, 0x38, 0x6c, 0x88, 0x0b, 0x4a, 0x39, 0x71, 0x89, 0xdc,
	0xf4, 0xbb, 0xc4, 0x5a, 0x62, 0x7b, 0xb6, 0x53, 0x89, 0x2b, 0x12, 0xf7, 0x49, 0x1c, 0x10, 0xff,
	0x11, 0xc7, 0x49, 0x5c, 0x38, 0x01, 0x6a, 0xf9, 0x43, 0x50, 0x6c, 0x67, 0xa8, 0x87, 0x0d, 0x4e,
	0xb1, 0xdf, 0xfb, 0xd8, 0x79, 0x7a, 0x5f, 0xa3, 0x87, 0xa9, 0xd0, 0xa5, 0xd0, 0x64, 0x46, 0x35,
	0x10, 0x2e, 0xe6, 0x40, 0x16, 0x07, 0x33, 0x30, 0xf4, 0x80, 0x9c, 0x55, 0xa0, 0xde, 0x61, 0xa9,
	0x84, 0x11, 0x61, 0xdf, 0x51, 0xb8, 0xa6, 0x70, 0x4d, 0x61, 0x4f, 0x0d, 0x1e, 0x64, 0x42, 0x64,
	0x05, 0x10, 0x2a, 0x19, 0xa1, 0x9c, 0x0b, 0x43, 0x0d, 0x13, 0x5c, 0xbb, 0x73, 0x83, 0x3d, 0xef,
	0xda, 0xdd, 0xac, 0x3a, 0x21, 0x86, 0x95, 0xa0, 0x0d, 0x2d, 0xa5, 0x07, 0xee, 0x64, 0x22, 0x13,
	0x76, 0x49, 0xea, 0x95, 0x53, 0x47, 0xdb, 0x68, 0xeb, 0x85, 0xe0, 0x27, 0x2c, 0x8b, 0xe1, 0xac,
	0x02, 0x6d, 0x46, 0x9f, 0x02, 0xd4, 0x6b, 0x14, 0x2d, 0x05, 0xd7, 0x10, 0x3e, 0x45, 0x3b, 0x25,
	0xe3, 0xac, 0xac, 0xca, 0x24, 0xa3, 0x3a, 0x91, 0x8a, 0xa5, 0xd0, 0x0f, 0x86, 0xc1, 0x78, 0x33,
	0xde, 0xf6, 0xc6, 0x11, 0xd5, 0xaf, 0x6b, 0x39, 0xc4, 0x68, 0x57, 0xaa, 0x8a, 0x33, 0x9e, 0x25,
	0xa7, 0x00, 0x32, 0x51, 0x90, 0x02, 0x37, 0xfd, 0xb6, 0xa5, 0x77, 0xbc, 0xf5, 0x0a, 0x40, 0xc6,
	0xd6, 0x08, 0x9f, 0xa0, 0xdb, 0x0d, 0xcf, 0xb8, 0x01, 0xb5, 0xa0, 0x45, 0xff, 0x3f, 0x77, 0xb5,
	0xd7, 0x5f, 0x7a, 0xb9, 0x8e, 0x3a, 0x35, 0xd4, 0x54, 0xba, 0x89, 0xfa, 0x3d, 0x40, 0xbd, 0x46,
	0xf1, 0x51, 0x27, 0xe8, 0x2e, 0x50, 0x55, 0x30, 0xd0, 0x26, 0xd1, 0x46, 0x28, 0x48, 0x72, 0x60,
	0x59, 0x6e, 0x6c, 0xdc, 0x4e, 0xbc, 0xdb, 0x98, 0xd3, 0xda, 0x3b, 0xb6, 0x56, 0x78, 0x0f, 0x75,
	0x3d, 0xd4, 0xb6, 0x90, 0xdf, 0x85, 0xcf, 0xd1, 0xe6, 0x65, 0x87, 0x36, 0xd3, 0xcd, 0xc9, 0x00,
	0xbb, 0x96, 0x71, 0xd3, 0x32, 0x7e, 0xd3, 0x10, 0x87, 0x9d, 0xf3, 0x1f, 0x7b, 0x41, 0xfc, 0xe7,
	0x48, 0x78, 0x1f, 0xdd, 0xa0, 0x52, 0x26, 0x39, 0xd5, 0x79, 0xbf, 0x33, 0x0c, 0xc6, 0xb7, 0xe2,
	0x0d, 0x2a, 0xe5, 0x31, 0xd5, 0x79, 0xf8, 0x08, 0xf5, 0x16, 0xb4, 0x60, 0x73, 0x6a, 0x84, 0x72,
	0xc0, 0xff, 0x16, 0xd8, 0xba, 0x54, 0x6b, 0x6c, 0xf2, 0xb9, 0x8d, 0x36, 0xa6, 0xa0, 0x16, 0x75,
	0xb1, 0x1f, 0x02, 0xd4, 0x75, 0x73, 0x09, 0x1f, 0xe3, 0xab, 0xde, 0x08, 0x5e, 0x9b, 0xe5, 0x60,
	0xfc, 0x77, 0xd0, 0xf5, 0x36, 0x1a, 0xbf, 0xff, 0xfa, 0xeb, 0x63, 0x7b, 0x14, 0x0e, 0xc9, 0x95,
	0x8f, 0x34, 0x75, 0x3f, 0xaf, 0x73, 0xb8, 0xd2, 0xaf, 0xcb, 0xb1, 0x36, 0xa8, 0xeb, 0x72, 0xac,
	0xcf, 0xef, 0x5f, 0x72, 0x68, 0x7b, 0xe2, 0xf0, 0xe8, 0xcb, 0x32, 0x0a, 0x2e, 0x96, 0x51, 0xf0,
	0x73, 0x19, 0x05, 0xe7, 0xab, 0xa8, 0x75, 0xb1, 0x8a, 0x5a, 0xdf, 0x56, 0x51, 0xeb, 0xed, 0x7e,
	0xc6, 0x4c, 0x5e, 0xcd, 0x70, 0x2a, 0xca, 0xe6, 0x16, 0xf7, 0xd9, 0xd7, 0xf3, 0x53, 0x92, 0x16,
	0x0c, 0xb8, 0x21, 0x99, 0x92, 0xa9, 0xbd, 0x77, 0xd6, 0xb5, 0xb3, 0x7c, 0xf6, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x73, 0x23, 0x6e, 0x9f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// Config queries for the operator configuration.
	Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
	// Status queries for the node status.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Config(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.node.v1beta1.Service/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/cosmos.base.node.v1beta1.Service/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Config queries for the operator configuration.
	Config(context.Context, *ConfigRequest) (*ConfigResponse, error)
	// Status queries for the node status.
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) Config(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (*UnimplementedServiceServer) Status(ctx context.Context, req *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.node.v1beta1.Service/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Config(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.base.node.v1beta1.Service/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.base.node.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Config",
			Handler:    _Service_Config_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Service_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/base/node/v1beta1/query.proto",
}

func (m *ConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PruningInterval) > 0 {
		i -= len(m.PruningInterval)
		copy(dAtA[i:], m.PruningInterval)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PruningInterval)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PruningKeepRecent) > 0 {
		i -= len(m.PruningKeepRecent)
		copy(dAtA[i:], m.PruningKeepRecent)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PruningKeepRecent)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MinimumGasPrice) > 0 {
		i -= len(m.MinimumGasPrice)
		copy(dAtA[i:], m.MinimumGasPrice)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MinimumGasPrice)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *StatusResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorHash) > 0 {
		i -= len(m.ValidatorHash)
		copy(dAtA[i:], m.ValidatorHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ValidatorHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Timestamp != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Timestamp, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Timestamp):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintQuery(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.EarliestStoreHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.EarliestStoreHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MinimumGasPrice)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PruningKeepRecent)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PruningInterval)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *StatusRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *StatusResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EarliestStoreHeight != 0 {
		n += 1 + sovQuery(uint64(m.EarliestStoreHeight))
	}
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	if m.Timestamp != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Timestamp)
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ValidatorHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ConfigResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinimumGasPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PruningKeepRecent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PruningKeepRecent = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PruningInterval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PruningInterval = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *StatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: StatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *StatusResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: StatusResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarliestStoreHeight", wireType)
			}
			m.EarliestStoreHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EarliestStoreHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorHash = append(m.ValidatorHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorHash == nil {
				m.ValidatorHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

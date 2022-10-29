// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accesscontrol/constants.proto

package accesscontrol

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type AccessType int32

const (
	AccessType_UNKNOWN AccessType = 0
	AccessType_READ    AccessType = 1
	AccessType_WRITE   AccessType = 2
	AccessType_COMMIT  AccessType = 3
)

var AccessType_name = map[int32]string{
	0: "UNKNOWN",
	1: "READ",
	2: "WRITE",
	3: "COMMIT",
}

var AccessType_value = map[string]int32{
	"UNKNOWN": 0,
	"READ":    1,
	"WRITE":   2,
	"COMMIT":  3,
}

func (x AccessType) String() string {
	return proto.EnumName(AccessType_name, int32(x))
}

func (AccessType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{0}
}

type AccessOperationSelectorType int32

const (
	AccessOperationSelectorType_NONE AccessOperationSelectorType = 0
	AccessOperationSelectorType_JQ   AccessOperationSelectorType = 1
)

var AccessOperationSelectorType_name = map[int32]string{
	0: "NONE",
	1: "JQ",
}

var AccessOperationSelectorType_value = map[string]int32{
	"NONE": 0,
	"JQ":   1,
}

func (x AccessOperationSelectorType) String() string {
	return proto.EnumName(AccessOperationSelectorType_name, int32(x))
}

func (AccessOperationSelectorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{1}
}

type ResourceType int32

const (
	ResourceType_ANY                                 ResourceType = 0
	ResourceType_KV                                  ResourceType = 1
	ResourceType_Mem                                 ResourceType = 2
	ResourceType_DexMem                              ResourceType = 3
	ResourceType_KV_BANK                             ResourceType = 4
	ResourceType_KV_STAKING                          ResourceType = 5
	ResourceType_KV_WASM                             ResourceType = 6
	ResourceType_KV_ORACLE                           ResourceType = 7
	ResourceType_KV_DEX                              ResourceType = 8
	ResourceType_KV_EPOCH                            ResourceType = 9
	ResourceType_KV_TOKENFACTORY                     ResourceType = 10
	ResourceType_KV_ORACLE_VOTE_TARGETS              ResourceType = 11
	ResourceType_KV_ORACLE_AGGREGATE_VOTES           ResourceType = 12
	ResourceType_KV_ORACLE_FEEDERS                   ResourceType = 13
	ResourceType_KV_STAKING_DELEGATION               ResourceType = 14
	ResourceType_KV_STAKING_VALIDATOR                ResourceType = 15
	ResourceType_KV_AUTH                             ResourceType = 16
	ResourceType_KV_AUTH_ADDRESS_STORE               ResourceType = 17
	ResourceType_KV_BANK_SUPPLY                      ResourceType = 18
	ResourceType_KV_BANK_DENOM                       ResourceType = 19
	ResourceType_KV_BANK_BALANCES                    ResourceType = 20
	ResourceType_KV_TOKENFACTORY_DENOM               ResourceType = 21
	ResourceType_KV_TOKENFACTORY_METADATA            ResourceType = 22
	ResourceType_KV_TOKENFACTORY_ADMIN               ResourceType = 23
	ResourceType_KV_TOKENFACTORY_CREATOR             ResourceType = 24
	ResourceType_KV_ORACLE_EXCHANGE_RATE             ResourceType = 25
	ResourceType_KV_ORACLE_VOTE_PENALTY_COUNTER      ResourceType = 26
	ResourceType_KV_ORACLE_PRICE_SNAPSHOT            ResourceType = 27
	ResourceType_KV_STAKING_VALIDATION_POWER         ResourceType = 28
	ResourceType_KV_STAKING_TOTAL_POWER              ResourceType = 29
	ResourceType_KV_STAKING_VALIDATORS_CON_ADDR      ResourceType = 30
	ResourceType_KV_STAKING_UNBONDING_DELEGATION     ResourceType = 31
	ResourceType_KV_STAKING_UNBONDING_DELEGATION_VAL ResourceType = 32
	ResourceType_KV_STAKING_REDELEGATION             ResourceType = 33
	ResourceType_KV_STAKING_REDELEGATION_VAL_SRC     ResourceType = 34
	ResourceType_KV_STAKING_REDELEGATION_VAL_DST     ResourceType = 35
	ResourceType_KV_STAKING_REDELEGATION_QUEUE       ResourceType = 36
	ResourceType_KV_STAKING_VALIDATOR_QUEUE          ResourceType = 37
	ResourceType_KV_STAKING_HISTORICAL_INFO          ResourceType = 38
	ResourceType_KV_STAKING_UNBONDING                ResourceType = 39
)

var ResourceType_name = map[int32]string{
	0:  "ANY",
	1:  "KV",
	2:  "Mem",
	3:  "DexMem",
	4:  "KV_BANK",
	5:  "KV_STAKING",
	6:  "KV_WASM",
	7:  "KV_ORACLE",
	8:  "KV_DEX",
	9:  "KV_EPOCH",
	10: "KV_TOKENFACTORY",
	11: "KV_ORACLE_VOTE_TARGETS",
	12: "KV_ORACLE_AGGREGATE_VOTES",
	13: "KV_ORACLE_FEEDERS",
	14: "KV_STAKING_DELEGATION",
	15: "KV_STAKING_VALIDATOR",
	16: "KV_AUTH",
	17: "KV_AUTH_ADDRESS_STORE",
	18: "KV_BANK_SUPPLY",
	19: "KV_BANK_DENOM",
	20: "KV_BANK_BALANCES",
	21: "KV_TOKENFACTORY_DENOM",
	22: "KV_TOKENFACTORY_METADATA",
	23: "KV_TOKENFACTORY_ADMIN",
	24: "KV_TOKENFACTORY_CREATOR",
	25: "KV_ORACLE_EXCHANGE_RATE",
	26: "KV_ORACLE_VOTE_PENALTY_COUNTER",
	27: "KV_ORACLE_PRICE_SNAPSHOT",
	28: "KV_STAKING_VALIDATION_POWER",
	29: "KV_STAKING_TOTAL_POWER",
	30: "KV_STAKING_VALIDATORS_CON_ADDR",
	31: "KV_STAKING_UNBONDING_DELEGATION",
	32: "KV_STAKING_UNBONDING_DELEGATION_VAL",
	33: "KV_STAKING_REDELEGATION",
	34: "KV_STAKING_REDELEGATION_VAL_SRC",
	35: "KV_STAKING_REDELEGATION_VAL_DST",
	36: "KV_STAKING_REDELEGATION_QUEUE",
	37: "KV_STAKING_VALIDATOR_QUEUE",
	38: "KV_STAKING_HISTORICAL_INFO",
	39: "KV_STAKING_UNBONDING",
}

var ResourceType_value = map[string]int32{
	"ANY":                                 0,
	"KV":                                  1,
	"Mem":                                 2,
	"DexMem":                              3,
	"KV_BANK":                             4,
	"KV_STAKING":                          5,
	"KV_WASM":                             6,
	"KV_ORACLE":                           7,
	"KV_DEX":                              8,
	"KV_EPOCH":                            9,
	"KV_TOKENFACTORY":                     10,
	"KV_ORACLE_VOTE_TARGETS":              11,
	"KV_ORACLE_AGGREGATE_VOTES":           12,
	"KV_ORACLE_FEEDERS":                   13,
	"KV_STAKING_DELEGATION":               14,
	"KV_STAKING_VALIDATOR":                15,
	"KV_AUTH":                             16,
	"KV_AUTH_ADDRESS_STORE":               17,
	"KV_BANK_SUPPLY":                      18,
	"KV_BANK_DENOM":                       19,
	"KV_BANK_BALANCES":                    20,
	"KV_TOKENFACTORY_DENOM":               21,
	"KV_TOKENFACTORY_METADATA":            22,
	"KV_TOKENFACTORY_ADMIN":               23,
	"KV_TOKENFACTORY_CREATOR":             24,
	"KV_ORACLE_EXCHANGE_RATE":             25,
	"KV_ORACLE_VOTE_PENALTY_COUNTER":      26,
	"KV_ORACLE_PRICE_SNAPSHOT":            27,
	"KV_STAKING_VALIDATION_POWER":         28,
	"KV_STAKING_TOTAL_POWER":              29,
	"KV_STAKING_VALIDATORS_CON_ADDR":      30,
	"KV_STAKING_UNBONDING_DELEGATION":     31,
	"KV_STAKING_UNBONDING_DELEGATION_VAL": 32,
	"KV_STAKING_REDELEGATION":             33,
	"KV_STAKING_REDELEGATION_VAL_SRC":     34,
	"KV_STAKING_REDELEGATION_VAL_DST":     35,
	"KV_STAKING_REDELEGATION_QUEUE":       36,
	"KV_STAKING_VALIDATOR_QUEUE":          37,
	"KV_STAKING_HISTORICAL_INFO":          38,
	"KV_STAKING_UNBONDING":                39,
}

func (x ResourceType) String() string {
	return proto.EnumName(ResourceType_name, int32(x))
}

func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_36568f7561081112, []int{2}
}

func init() {
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessType", AccessType_name, AccessType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.AccessOperationSelectorType", AccessOperationSelectorType_name, AccessOperationSelectorType_value)
	proto.RegisterEnum("cosmos.accesscontrol.v1beta1.ResourceType", ResourceType_name, ResourceType_value)
}

func init() {
	proto.RegisterFile("cosmos/accesscontrol/constants.proto", fileDescriptor_36568f7561081112)
}

var fileDescriptor_36568f7561081112 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x41, 0x6f, 0xdb, 0x36,
	0x14, 0xc7, 0xed, 0x26, 0x8d, 0x9d, 0xd7, 0x24, 0x7d, 0x61, 0x93, 0x36, 0x89, 0x13, 0x65, 0x6d,
	0xba, 0x15, 0x08, 0x30, 0x7b, 0xc5, 0x6e, 0xbb, 0xd1, 0x12, 0x63, 0xab, 0x92, 0x49, 0x97, 0xa4,
	0xe5, 0x7a, 0x17, 0xc2, 0xd1, 0x84, 0xad, 0x58, 0x63, 0x19, 0x96, 0x3a, 0xac, 0xdf, 0x62, 0x1f,
	0x6b, 0xc7, 0x1e, 0x77, 0x1c, 0x92, 0xe3, 0xbe, 0xc4, 0x20, 0x9b, 0x6e, 0x3c, 0xb5, 0x43, 0x4e,
	0xb6, 0xf8, 0xff, 0xbd, 0x3f, 0xf9, 0xfe, 0x24, 0x1e, 0x3c, 0x8f, 0xd3, 0xec, 0x2a, 0xcd, 0x5a,
	0xe3, 0x38, 0x4e, 0xb2, 0x2c, 0x4e, 0x27, 0xf9, 0x2c, 0x7d, 0xd7, 0x8a, 0xd3, 0x49, 0x96, 0x8f,
	0x27, 0x79, 0xd6, 0x9c, 0xce, 0xd2, 0x3c, 0x25, 0xc7, 0x0b, 0xaa, 0xf9, 0x1f, 0xaa, 0xf9, 0xdb,
	0xcb, 0xcb, 0x24, 0x1f, 0xbf, 0x3c, 0xff, 0x01, 0x80, 0xce, 0x05, 0xfd, 0x61, 0x9a, 0x90, 0x07,
	0x50, 0x1b, 0xf0, 0x80, 0x8b, 0x21, 0xc7, 0x0a, 0xa9, 0xc3, 0xba, 0x64, 0xd4, 0xc3, 0x2a, 0xd9,
	0x84, 0xfb, 0x43, 0xe9, 0x6b, 0x86, 0xf7, 0x08, 0xc0, 0x86, 0x2b, 0x7a, 0x3d, 0x5f, 0xe3, 0xda,
	0x79, 0x0b, 0x1a, 0x8b, 0x5a, 0x31, 0x4d, 0x66, 0xe3, 0xfc, 0x6d, 0x3a, 0x51, 0xc9, 0xbb, 0x24,
	0xce, 0xd3, 0xd9, 0xdc, 0xac, 0x0e, 0xeb, 0x5c, 0x70, 0x86, 0x15, 0xb2, 0x01, 0xf7, 0x5e, 0xbd,
	0xc6, 0xea, 0xf9, 0x3f, 0x35, 0xd8, 0x92, 0x49, 0x96, 0xbe, 0x9f, 0xc5, 0xc9, 0x1c, 0xa9, 0xc1,
	0x1a, 0xe5, 0xa3, 0x05, 0x11, 0x44, 0x58, 0x2d, 0x16, 0x7a, 0xc9, 0xd5, 0x62, 0x1f, 0x2f, 0xf9,
	0xbd, 0xf8, 0xbf, 0x56, 0x9c, 0x2a, 0x88, 0x4c, 0x9b, 0xf2, 0x00, 0xd7, 0xc9, 0x0e, 0x40, 0x10,
	0x19, 0xa5, 0x69, 0xe0, 0xf3, 0x0e, 0xde, 0xb7, 0xe2, 0x90, 0xaa, 0x1e, 0x6e, 0x90, 0x6d, 0xd8,
	0x0c, 0x22, 0x23, 0x24, 0x75, 0x43, 0x86, 0xb5, 0xc2, 0x24, 0x88, 0x8c, 0xc7, 0xde, 0x60, 0x9d,
	0x6c, 0x41, 0x3d, 0x88, 0x0c, 0xeb, 0x0b, 0xb7, 0x8b, 0x9b, 0xe4, 0x11, 0x3c, 0x0c, 0x22, 0xa3,
	0x45, 0xc0, 0xf8, 0x05, 0x75, 0xb5, 0x90, 0x23, 0x04, 0x72, 0x04, 0x8f, 0x3f, 0x55, 0x9b, 0x48,
	0x68, 0x66, 0x34, 0x95, 0x1d, 0xa6, 0x15, 0x3e, 0x20, 0x27, 0x70, 0x78, 0xab, 0xd1, 0x4e, 0x47,
	0xb2, 0x0e, 0xd5, 0x0b, 0x4a, 0xe1, 0x16, 0xd9, 0x87, 0xdd, 0x5b, 0xf9, 0x82, 0x31, 0x8f, 0x49,
	0x85, 0xdb, 0xe4, 0x10, 0xf6, 0x6f, 0x0f, 0x6b, 0x3c, 0x16, 0x16, 0x55, 0xbe, 0xe0, 0xb8, 0x43,
	0x0e, 0x60, 0x6f, 0x45, 0x8a, 0x68, 0xe8, 0x7b, 0x54, 0x0b, 0x89, 0x0f, 0x6d, 0x47, 0x74, 0xa0,
	0xbb, 0x88, 0xd6, 0xa1, 0xf8, 0x30, 0xd4, 0xf3, 0x24, 0x53, 0xca, 0x28, 0x2d, 0x24, 0xc3, 0x5d,
	0x42, 0x60, 0xc7, 0xc6, 0x62, 0xd4, 0xa0, 0xdf, 0x0f, 0x47, 0x48, 0xc8, 0x2e, 0x6c, 0x2f, 0xd7,
	0x3c, 0xc6, 0x45, 0x0f, 0x1f, 0x91, 0x3d, 0xc0, 0xe5, 0x52, 0x9b, 0x86, 0x94, 0xbb, 0x4c, 0xe1,
	0x9e, 0xf5, 0x5d, 0x0d, 0xc0, 0x16, 0xec, 0x93, 0x63, 0x38, 0x28, 0x4b, 0x3d, 0xa6, 0xa9, 0x47,
	0x35, 0xc5, 0xc7, 0x5f, 0x2a, 0xa4, 0x5e, 0xcf, 0xe7, 0xf8, 0x84, 0x34, 0xe0, 0x49, 0x59, 0x72,
	0x25, 0x9b, 0x77, 0x75, 0x60, 0x45, 0x9b, 0x10, 0x7b, 0xe3, 0x76, 0x29, 0xef, 0x30, 0x23, 0xa9,
	0x66, 0x78, 0x48, 0x9e, 0x81, 0x53, 0x4a, 0xbe, 0xcf, 0x38, 0x0d, 0xf5, 0xc8, 0xb8, 0x62, 0xc0,
	0x35, 0x93, 0x78, 0x64, 0x8f, 0x65, 0x99, 0xbe, 0xf4, 0x5d, 0x66, 0x14, 0xa7, 0x7d, 0xd5, 0x15,
	0x1a, 0x1b, 0xe4, 0x14, 0x1a, 0x9f, 0xc7, 0xe9, 0x0b, 0x6e, 0xfa, 0x62, 0xc8, 0x24, 0x1e, 0xdb,
	0xcb, 0x5d, 0x02, 0x5a, 0x68, 0x1a, 0x5a, 0xed, 0xc4, 0x6e, 0xff, 0xd9, 0x5d, 0x28, 0xe3, 0x0a,
	0x3e, 0x8f, 0x1d, 0x1d, 0x72, 0x06, 0xa7, 0x2b, 0xcc, 0x80, 0xb7, 0x05, 0xf7, 0x4a, 0x97, 0x7a,
	0x4a, 0x5e, 0xc0, 0xd9, 0x1d, 0x50, 0xe1, 0x8e, 0x5f, 0xd9, 0x34, 0x96, 0xa0, 0x64, 0x2b, 0x2e,
	0x4f, 0x4b, 0x5b, 0xad, 0x8a, 0x45, 0xb5, 0x51, 0xd2, 0xc5, 0x67, 0x77, 0x41, 0x9e, 0xd2, 0x78,
	0x46, 0x9e, 0xc2, 0xc9, 0xff, 0x41, 0xaf, 0x07, 0x6c, 0xc0, 0xf0, 0x39, 0x71, 0xe0, 0xe8, 0x4b,
	0xbd, 0x5b, 0xfd, 0xeb, 0x92, 0xde, 0xf5, 0x8b, 0xd7, 0xe7, 0xbb, 0x34, 0x34, 0x3e, 0xbf, 0x10,
	0xf8, 0x4d, 0xe9, 0x1d, 0x7f, 0x6a, 0x19, 0x5f, 0xb4, 0x5f, 0xfd, 0x79, 0xed, 0x54, 0x3f, 0x5e,
	0x3b, 0xd5, 0xbf, 0xaf, 0x9d, 0xea, 0x1f, 0x37, 0x4e, 0xe5, 0xe3, 0x8d, 0x53, 0xf9, 0xeb, 0xc6,
	0xa9, 0xfc, 0xf8, 0xdd, 0xcf, 0x6f, 0xf3, 0x5f, 0xde, 0x5f, 0x36, 0xe3, 0xf4, 0xaa, 0x65, 0x67,
	0xd8, 0xe2, 0xe7, 0xdb, 0xec, 0xa7, 0x5f, 0x5b, 0xf9, 0x87, 0x69, 0x52, 0x1a, 0x6a, 0x97, 0x1b,
	0xf3, 0x59, 0xf6, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x8d, 0x21, 0xda, 0xf3, 0x04,
	0x00, 0x00,
}

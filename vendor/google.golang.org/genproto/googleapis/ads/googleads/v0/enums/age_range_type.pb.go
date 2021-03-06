// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/age_range_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of demographic age ranges (e.g. between 18 and 24 years old).
type AgeRangeTypeEnum_AgeRangeType int32

const (
	// Not specified.
	AgeRangeTypeEnum_UNSPECIFIED AgeRangeTypeEnum_AgeRangeType = 0
	// Used for return value only. Represents value unknown in this version.
	AgeRangeTypeEnum_UNKNOWN AgeRangeTypeEnum_AgeRangeType = 1
	// Between 18 and 24 years old.
	AgeRangeTypeEnum_AGE_RANGE_18_24 AgeRangeTypeEnum_AgeRangeType = 503001
	// Between 25 and 34 years old.
	AgeRangeTypeEnum_AGE_RANGE_25_34 AgeRangeTypeEnum_AgeRangeType = 503002
	// Between 35 and 44 years old.
	AgeRangeTypeEnum_AGE_RANGE_35_44 AgeRangeTypeEnum_AgeRangeType = 503003
	// Between 45 and 54 years old.
	AgeRangeTypeEnum_AGE_RANGE_45_54 AgeRangeTypeEnum_AgeRangeType = 503004
	// Between 55 and 64 years old.
	AgeRangeTypeEnum_AGE_RANGE_55_64 AgeRangeTypeEnum_AgeRangeType = 503005
	// 65 years old and beyond.
	AgeRangeTypeEnum_AGE_RANGE_65_UP AgeRangeTypeEnum_AgeRangeType = 503006
	// Undetermined age range.
	AgeRangeTypeEnum_AGE_RANGE_UNDETERMINED AgeRangeTypeEnum_AgeRangeType = 503999
)

var AgeRangeTypeEnum_AgeRangeType_name = map[int32]string{
	0:      "UNSPECIFIED",
	1:      "UNKNOWN",
	503001: "AGE_RANGE_18_24",
	503002: "AGE_RANGE_25_34",
	503003: "AGE_RANGE_35_44",
	503004: "AGE_RANGE_45_54",
	503005: "AGE_RANGE_55_64",
	503006: "AGE_RANGE_65_UP",
	503999: "AGE_RANGE_UNDETERMINED",
}
var AgeRangeTypeEnum_AgeRangeType_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"AGE_RANGE_18_24":        503001,
	"AGE_RANGE_25_34":        503002,
	"AGE_RANGE_35_44":        503003,
	"AGE_RANGE_45_54":        503004,
	"AGE_RANGE_55_64":        503005,
	"AGE_RANGE_65_UP":        503006,
	"AGE_RANGE_UNDETERMINED": 503999,
}

func (x AgeRangeTypeEnum_AgeRangeType) String() string {
	return proto.EnumName(AgeRangeTypeEnum_AgeRangeType_name, int32(x))
}
func (AgeRangeTypeEnum_AgeRangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_age_range_type_846fc75750f12f7f, []int{0, 0}
}

// Container for enum describing the type of demographic age ranges.
type AgeRangeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgeRangeTypeEnum) Reset()         { *m = AgeRangeTypeEnum{} }
func (m *AgeRangeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AgeRangeTypeEnum) ProtoMessage()    {}
func (*AgeRangeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_age_range_type_846fc75750f12f7f, []int{0}
}
func (m *AgeRangeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgeRangeTypeEnum.Unmarshal(m, b)
}
func (m *AgeRangeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgeRangeTypeEnum.Marshal(b, m, deterministic)
}
func (dst *AgeRangeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgeRangeTypeEnum.Merge(dst, src)
}
func (m *AgeRangeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AgeRangeTypeEnum.Size(m)
}
func (m *AgeRangeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AgeRangeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AgeRangeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AgeRangeTypeEnum)(nil), "google.ads.googleads.v0.enums.AgeRangeTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AgeRangeTypeEnum_AgeRangeType", AgeRangeTypeEnum_AgeRangeType_name, AgeRangeTypeEnum_AgeRangeType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/age_range_type.proto", fileDescriptor_age_range_type_846fc75750f12f7f)
}

var fileDescriptor_age_range_type_846fc75750f12f7f = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xc4, 0xf4, 0xd4, 0xf8, 0xa2, 0xc4, 0xbc, 0xf4, 0xd4, 0xf8, 0x92, 0xca,
	0x82, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x42, 0xbd, 0xc4, 0x94, 0x62,
	0x3d, 0xb8, 0x1e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x1e, 0xa5, 0x97, 0x8c, 0x5c, 0x02, 0x8e, 0xe9,
	0xa9, 0x41, 0x20, 0x6d, 0x21, 0x95, 0x05, 0xa9, 0xae, 0x79, 0xa5, 0xb9, 0x4a, 0x57, 0x18, 0xb9,
	0x78, 0x90, 0x05, 0x85, 0xf8, 0xb9, 0xb8, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d, 0x3d, 0xdd, 0x3c,
	0x5d, 0x5d, 0x04, 0x18, 0x84, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04,
	0x18, 0x85, 0x44, 0xb9, 0xf8, 0x1d, 0xdd, 0x5d, 0xe3, 0x83, 0x1c, 0xfd, 0xdc, 0x5d, 0xe3, 0x0d,
	0x2d, 0xe2, 0x8d, 0x4c, 0x04, 0x6e, 0xde, 0x94, 0x43, 0x15, 0x36, 0x32, 0x8d, 0x37, 0x36, 0x11,
	0xb8, 0x85, 0x2e, 0x6c, 0x6c, 0x1a, 0x6f, 0x62, 0x22, 0x70, 0x1b, 0x5d, 0xd8, 0xc4, 0x34, 0xde,
	0xd4, 0x44, 0xe0, 0x0e, 0xba, 0xb0, 0xa9, 0x69, 0xbc, 0x99, 0x89, 0xc0, 0x5d, 0x74, 0x61, 0x33,
	0xd3, 0xf8, 0xd0, 0x00, 0x81, 0x7b, 0x37, 0xe5, 0x84, 0x64, 0xb8, 0xc4, 0x10, 0xc2, 0xa1, 0x7e,
	0x2e, 0xae, 0x21, 0xae, 0x41, 0xbe, 0x9e, 0x7e, 0xae, 0x2e, 0x02, 0xfb, 0x1f, 0xca, 0x39, 0x1d,
	0x62, 0xe4, 0x52, 0x4c, 0xce, 0xcf, 0xd5, 0xc3, 0x1b, 0x22, 0x4e, 0x82, 0xc8, 0x3e, 0x0f, 0x00,
	0x85, 0x61, 0x00, 0x63, 0x94, 0x13, 0x54, 0x4f, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e,
	0x51, 0xba, 0x7e, 0x7a, 0x6a, 0x1e, 0x38, 0x84, 0x61, 0x31, 0x51, 0x90, 0x59, 0x8c, 0x23, 0x62,
	0xac, 0xc1, 0xe4, 0x22, 0x26, 0x66, 0x77, 0x47, 0xc7, 0x55, 0x4c, 0xb2, 0xee, 0x10, 0xa3, 0x1c,
	0x53, 0x8a, 0xf5, 0x20, 0x4c, 0x10, 0x2b, 0xcc, 0x40, 0x0f, 0x14, 0xf4, 0xc5, 0xa7, 0x60, 0xf2,
	0x31, 0x8e, 0x29, 0xc5, 0x31, 0x70, 0xf9, 0x98, 0x30, 0x83, 0x18, 0xb0, 0x7c, 0x12, 0x1b, 0xd8,
	0x52, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x72, 0xdc, 0x47, 0x0c, 0x02, 0x00, 0x00,
}

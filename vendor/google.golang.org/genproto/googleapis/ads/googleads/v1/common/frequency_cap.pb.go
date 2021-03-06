// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/frequency_cap.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v1/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A rule specifying the maximum number of times an ad (or some set of ads) can
// be shown to a user over a particular time period.
type FrequencyCapEntry struct {
	// The key of a particular frequency cap. There can be no more
	// than one frequency cap with the same key.
	Key *FrequencyCapKey `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Maximum number of events allowed during the time range by this cap.
	Cap                  *wrappers.Int32Value `protobuf:"bytes,2,opt,name=cap,proto3" json:"cap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapEntry) Reset()         { *m = FrequencyCapEntry{} }
func (m *FrequencyCapEntry) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapEntry) ProtoMessage()    {}
func (*FrequencyCapEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_31028d98fca9f9b0, []int{0}
}
func (m *FrequencyCapEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapEntry.Unmarshal(m, b)
}
func (m *FrequencyCapEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapEntry.Marshal(b, m, deterministic)
}
func (dst *FrequencyCapEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapEntry.Merge(dst, src)
}
func (m *FrequencyCapEntry) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapEntry.Size(m)
}
func (m *FrequencyCapEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapEntry proto.InternalMessageInfo

func (m *FrequencyCapEntry) GetKey() *FrequencyCapKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *FrequencyCapEntry) GetCap() *wrappers.Int32Value {
	if m != nil {
		return m.Cap
	}
	return nil
}

// A group of fields used as keys for a frequency cap.
// There can be no more than one frequency cap with the same key.
type FrequencyCapKey struct {
	// The level on which the cap is to be applied (e.g. ad group ad, ad group).
	// The cap is applied to all the entities of this level.
	Level enums.FrequencyCapLevelEnum_FrequencyCapLevel `protobuf:"varint,1,opt,name=level,proto3,enum=google.ads.googleads.v1.enums.FrequencyCapLevelEnum_FrequencyCapLevel" json:"level,omitempty"`
	// The type of event that the cap applies to (e.g. impression).
	EventType enums.FrequencyCapEventTypeEnum_FrequencyCapEventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=google.ads.googleads.v1.enums.FrequencyCapEventTypeEnum_FrequencyCapEventType" json:"event_type,omitempty"`
	// Unit of time the cap is defined at (e.g. day, week).
	TimeUnit enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit `protobuf:"varint,2,opt,name=time_unit,json=timeUnit,proto3,enum=google.ads.googleads.v1.enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit" json:"time_unit,omitempty"`
	// Number of time units the cap lasts.
	TimeLength           *wrappers.Int32Value `protobuf:"bytes,4,opt,name=time_length,json=timeLength,proto3" json:"time_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FrequencyCapKey) Reset()         { *m = FrequencyCapKey{} }
func (m *FrequencyCapKey) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapKey) ProtoMessage()    {}
func (*FrequencyCapKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_frequency_cap_31028d98fca9f9b0, []int{1}
}
func (m *FrequencyCapKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapKey.Unmarshal(m, b)
}
func (m *FrequencyCapKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapKey.Marshal(b, m, deterministic)
}
func (dst *FrequencyCapKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapKey.Merge(dst, src)
}
func (m *FrequencyCapKey) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapKey.Size(m)
}
func (m *FrequencyCapKey) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapKey.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapKey proto.InternalMessageInfo

func (m *FrequencyCapKey) GetLevel() enums.FrequencyCapLevelEnum_FrequencyCapLevel {
	if m != nil {
		return m.Level
	}
	return enums.FrequencyCapLevelEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetEventType() enums.FrequencyCapEventTypeEnum_FrequencyCapEventType {
	if m != nil {
		return m.EventType
	}
	return enums.FrequencyCapEventTypeEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeUnit() enums.FrequencyCapTimeUnitEnum_FrequencyCapTimeUnit {
	if m != nil {
		return m.TimeUnit
	}
	return enums.FrequencyCapTimeUnitEnum_UNSPECIFIED
}

func (m *FrequencyCapKey) GetTimeLength() *wrappers.Int32Value {
	if m != nil {
		return m.TimeLength
	}
	return nil
}

func init() {
	proto.RegisterType((*FrequencyCapEntry)(nil), "google.ads.googleads.v1.common.FrequencyCapEntry")
	proto.RegisterType((*FrequencyCapKey)(nil), "google.ads.googleads.v1.common.FrequencyCapKey")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/frequency_cap.proto", fileDescriptor_frequency_cap_31028d98fca9f9b0)
}

var fileDescriptor_frequency_cap_31028d98fca9f9b0 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6a, 0xd4, 0x40,
	0x18, 0x27, 0x1b, 0x15, 0x3b, 0x85, 0x8a, 0x39, 0x2d, 0x55, 0x8a, 0xec, 0xc9, 0x8b, 0x33, 0x24,
	0x3d, 0x08, 0x69, 0x2f, 0x69, 0xdd, 0x16, 0x71, 0x91, 0x12, 0xea, 0x1e, 0x24, 0xb0, 0x4c, 0x93,
	0xaf, 0x31, 0x98, 0xcc, 0x8c, 0xc9, 0x64, 0x25, 0x0f, 0x20, 0xbe, 0x87, 0x47, 0x1f, 0xc5, 0x47,
	0x11, 0x1f, 0x42, 0x66, 0x26, 0x13, 0x5d, 0x97, 0xb5, 0xe4, 0x94, 0x5f, 0xe6, 0xfb, 0xfd, 0xf9,
	0xf2, 0x7d, 0x13, 0x14, 0xe4, 0x9c, 0xe7, 0x25, 0x10, 0x9a, 0x35, 0xc4, 0x40, 0x85, 0xd6, 0x3e,
	0x49, 0x79, 0x55, 0x71, 0x46, 0x6e, 0x6b, 0xf8, 0xd4, 0x02, 0x4b, 0xbb, 0x55, 0x4a, 0x05, 0x16,
	0x35, 0x97, 0xdc, 0x3b, 0x32, 0x44, 0x4c, 0xb3, 0x06, 0x0f, 0x1a, 0xbc, 0xf6, 0xb1, 0xd1, 0x1c,
	0x9e, 0xee, 0xf2, 0x04, 0xd6, 0x56, 0xcd, 0xa6, 0xe5, 0x0a, 0xd6, 0xc0, 0xe4, 0x4a, 0x76, 0x02,
	0x8c, 0xfb, 0xe1, 0xcb, 0x31, 0xea, 0x12, 0xd6, 0x50, 0xf6, 0xc2, 0x93, 0x31, 0x42, 0x59, 0x54,
	0xb0, 0x6a, 0x59, 0x21, 0x7b, 0x71, 0xff, 0x4d, 0x44, 0xbf, 0xdd, 0xb4, 0xb7, 0xe4, 0x73, 0x4d,
	0x85, 0x80, 0xba, 0xe9, 0xeb, 0x4f, 0xad, 0xb9, 0x28, 0x08, 0x65, 0x8c, 0x4b, 0x2a, 0x0b, 0xce,
	0xfa, 0xea, 0xec, 0x8b, 0x83, 0x1e, 0x5f, 0x58, 0xff, 0x73, 0x2a, 0xe6, 0x4c, 0xd6, 0x9d, 0x17,
	0x21, 0xf7, 0x23, 0x74, 0x53, 0xe7, 0x99, 0xf3, 0x7c, 0x3f, 0x20, 0xf8, 0xff, 0x53, 0xc3, 0x7f,
	0xeb, 0xdf, 0x40, 0x17, 0x2b, 0xad, 0xf7, 0x02, 0xb9, 0x29, 0x15, 0xd3, 0x89, 0xb6, 0x78, 0x62,
	0x2d, 0x6c, 0x93, 0xf8, 0x35, 0x93, 0xc7, 0xc1, 0x92, 0x96, 0x2d, 0xc4, 0x8a, 0x37, 0xfb, 0xea,
	0xa2, 0x47, 0xff, 0xf8, 0x78, 0x09, 0xba, 0xaf, 0xa7, 0xa4, 0xfb, 0x38, 0x08, 0x2e, 0x76, 0xf6,
	0xa1, 0xc7, 0xb4, 0xd1, 0xc6, 0x42, 0xe9, 0xe6, 0xac, 0xad, 0xb6, 0x4f, 0x63, 0x63, 0xea, 0x55,
	0x08, 0xfd, 0xd9, 0xe0, 0xd4, 0xd5, 0x11, 0x6f, 0x47, 0x44, 0xcc, 0x95, 0xf8, 0xba, 0x13, 0xb0,
	0x15, 0x33, 0x54, 0xe2, 0x3d, 0xb0, 0xd0, 0x2b, 0xd0, 0xde, 0xb0, 0x39, 0x3d, 0x95, 0x83, 0x60,
	0x31, 0x22, 0xed, 0xba, 0xa8, 0xe0, 0x1d, 0x2b, 0xe4, 0x56, 0x98, 0x2d, 0xc4, 0x0f, 0x65, 0x8f,
	0xbc, 0x53, 0xb4, 0xaf, 0xa3, 0x4a, 0x60, 0xb9, 0xfc, 0x30, 0xbd, 0x77, 0xf7, 0x0a, 0x90, 0xe2,
	0x2f, 0x34, 0xfd, 0xec, 0x97, 0x83, 0x66, 0x29, 0xaf, 0xee, 0x58, 0xfa, 0xd9, 0xc6, 0xad, 0xb9,
	0x52, 0x9e, 0x57, 0xce, 0xfb, 0x57, 0xbd, 0x28, 0xe7, 0x25, 0x65, 0x39, 0xe6, 0x75, 0x4e, 0x72,
	0x60, 0x3a, 0xd1, 0x5e, 0x6c, 0x51, 0x34, 0xbb, 0x7e, 0xd9, 0x13, 0xf3, 0xf8, 0x36, 0x71, 0x2f,
	0xa3, 0xe8, 0xfb, 0xe4, 0xe8, 0xd2, 0x98, 0x45, 0x59, 0x83, 0x0d, 0x54, 0x68, 0xe9, 0xe3, 0x73,
	0x4d, 0xfb, 0x61, 0x09, 0x49, 0x94, 0x35, 0xc9, 0x40, 0x48, 0x96, 0x7e, 0x62, 0x08, 0x3f, 0x27,
	0x33, 0x73, 0x1a, 0x86, 0x51, 0xd6, 0x84, 0xe1, 0x40, 0x09, 0xc3, 0xa5, 0x1f, 0x86, 0x86, 0x74,
	0xf3, 0x40, 0x77, 0x77, 0xfc, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xa5, 0x50, 0x57, 0x4f, 0x04,
	0x00, 0x00,
}

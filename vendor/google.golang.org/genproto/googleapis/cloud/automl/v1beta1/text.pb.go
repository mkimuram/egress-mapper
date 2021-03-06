// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Dataset metadata for classification.
type TextClassificationDatasetMetadata struct {
	// Required.
	// Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TextClassificationDatasetMetadata) Reset()         { *m = TextClassificationDatasetMetadata{} }
func (m *TextClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationDatasetMetadata) ProtoMessage()    {}
func (*TextClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_11ac4f99f444a22c, []int{0}
}
func (m *TextClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *TextClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (dst *TextClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationDatasetMetadata.Merge(dst, src)
}
func (m *TextClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationDatasetMetadata.Size(m)
}
func (m *TextClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationDatasetMetadata proto.InternalMessageInfo

func (m *TextClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata that is specific to text classification.
type TextClassificationModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextClassificationModelMetadata) Reset()         { *m = TextClassificationModelMetadata{} }
func (m *TextClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*TextClassificationModelMetadata) ProtoMessage()    {}
func (*TextClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_text_11ac4f99f444a22c, []int{1}
}
func (m *TextClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextClassificationModelMetadata.Unmarshal(m, b)
}
func (m *TextClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (dst *TextClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextClassificationModelMetadata.Merge(dst, src)
}
func (m *TextClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_TextClassificationModelMetadata.Size(m)
}
func (m *TextClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TextClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TextClassificationModelMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TextClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationDatasetMetadata")
	proto.RegisterType((*TextClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.TextClassificationModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text.proto", fileDescriptor_text_11ac4f99f444a22c)
}

var fileDescriptor_text_11ac4f99f444a22c = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x39, 0x0b, 0xc1, 0x2b, 0x2c, 0xce, 0x46, 0x12, 0x21, 0x26, 0x85, 0x58, 0xed, 0x1a,
	0x2d, 0xad, 0x92, 0x08, 0x56, 0x07, 0x41, 0x82, 0x85, 0x1c, 0xe8, 0xe4, 0x6e, 0x5c, 0x16, 0x36,
	0x3b, 0x4b, 0x76, 0x4e, 0x92, 0x1f, 0x60, 0xed, 0xff, 0xf2, 0x57, 0xc9, 0xed, 0x9e, 0xc5, 0x91,
	0x70, 0xe5, 0xee, 0x7c, 0xef, 0xcd, 0x7b, 0x93, 0xde, 0x28, 0x22, 0x65, 0x50, 0x96, 0x86, 0xea,
	0x4a, 0x42, 0xcd, 0xb4, 0x31, 0xf2, 0x6b, 0xba, 0x46, 0x86, 0xa9, 0x64, 0xdc, 0xb1, 0x70, 0x5b,
	0x62, 0xca, 0x86, 0x91, 0x13, 0x81, 0x13, 0x91, 0x13, 0x2d, 0x37, 0xb8, 0x6a, 0x4d, 0xc0, 0x69,
	0x09, 0xd6, 0x12, 0x03, 0x6b, 0xb2, 0x3e, 0x4a, 0x07, 0x77, 0x7d, 0x2b, 0x4a, 0x03, 0xde, 0xeb,
	0x4f, 0x5d, 0x06, 0x49, 0x54, 0x4c, 0xbe, 0x93, 0x74, 0xbc, 0xc2, 0x1d, 0x2f, 0x3a, 0xc3, 0x27,
	0x60, 0xf0, 0xc8, 0x39, 0x32, 0x54, 0xc0, 0x90, 0x7d, 0xa4, 0x17, 0x5d, 0xf5, 0x3b, 0xef, 0x1d,
	0x5e, 0x26, 0xd7, 0xc9, 0xed, 0xf9, 0xbd, 0x14, 0x3d, 0x81, 0x45, 0xd7, 0x78, 0xb5, 0x77, 0xf8,
	0x92, 0x95, 0x07, 0x7f, 0x93, 0x71, 0x3a, 0x3a, 0x8c, 0x91, 0x53, 0x85, 0xe6, 0x3f, 0xc4, 0xfc,
	0x27, 0x49, 0x47, 0x25, 0x6d, 0xfa, 0xb6, 0xcd, 0xcf, 0x1a, 0x93, 0x65, 0xd3, 0x6c, 0x99, 0xbc,
	0xcd, 0x5a, 0x52, 0x91, 0x01, 0xab, 0x04, 0x6d, 0x95, 0x54, 0x68, 0x43, 0x6f, 0x19, 0x47, 0xe0,
	0xb4, 0x3f, 0x7a, 0xac, 0xc7, 0xf8, 0xfc, 0x3d, 0x19, 0x3e, 0x07, 0xb0, 0x58, 0x34, 0x50, 0x31,
	0xab, 0x99, 0x72, 0x53, 0xbc, 0x46, 0x68, 0x7d, 0x1a, 0xbc, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x24, 0xf6, 0x56, 0x97, 0xda, 0x01, 0x00, 0x00,
}

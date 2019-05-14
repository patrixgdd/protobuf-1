// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test3/test_import.proto

package test3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type ImportEnum int32

const (
	ImportEnum_IMPORT_ZERO ImportEnum = 0
)

// Deprecated: Use ImportEnum.Type.Values instead.
var ImportEnum_name = map[int32]string{
	0: "IMPORT_ZERO",
}

// Deprecated: Use ImportEnum.Type.Values instead.
var ImportEnum_value = map[string]int32{
	"IMPORT_ZERO": 0,
}

func (x ImportEnum) Enum() *ImportEnum {
	p := new(ImportEnum)
	*p = x
	return p
}

func (x ImportEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_test3_test_import_proto_enumTypes[0].Descriptor()
}

// Deprecated: Use Descriptor instead.
func (ImportEnum) Type() protoreflect.EnumType {
	return file_test3_test_import_proto_enumTypes[0]
}

func (x ImportEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportEnum.Type instead.
func (ImportEnum) EnumDescriptor() ([]byte, []int) {
	return file_test3_test_import_proto_rawDescGZIP(), []int{0}
}

type ImportMessage struct {
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *ImportMessage) Reset() {
	*x = ImportMessage{}
}

func (x *ImportMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportMessage) ProtoMessage() {}

func (x *ImportMessage) ProtoReflect() protoreflect.Message {
	return file_test3_test_import_proto_msgTypes[0].MessageOf(x)
}

func (m *ImportMessage) XXX_Methods() *protoiface.Methods {
	return file_test3_test_import_proto_msgTypes[0].Methods()
}

// Deprecated: Use ImportMessage.ProtoReflect.Type instead.
func (*ImportMessage) Descriptor() ([]byte, []int) {
	return file_test3_test_import_proto_rawDescGZIP(), []int{0}
}

var File_test3_test_import_proto protoreflect.FileDescriptor

var file_test3_test_import_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x33, 0x22, 0x0f,
	0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a,
	0x1d, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0f, 0x0a,
	0x0b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test3_test_import_proto_rawDescOnce sync.Once
	file_test3_test_import_proto_rawDescData = file_test3_test_import_proto_rawDesc
)

func file_test3_test_import_proto_rawDescGZIP() []byte {
	file_test3_test_import_proto_rawDescOnce.Do(func() {
		file_test3_test_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_test3_test_import_proto_rawDescData)
	})
	return file_test3_test_import_proto_rawDescData
}

var file_test3_test_import_proto_enumTypes = make([]protoreflect.EnumType, 1)
var file_test3_test_import_proto_msgTypes = make([]protoimpl.MessageType, 1)
var file_test3_test_import_proto_goTypes = []interface{}{
	(ImportEnum)(0),       // 0: goproto.proto.test3.ImportEnum
	(*ImportMessage)(nil), // 1: goproto.proto.test3.ImportMessage
}
var file_test3_test_import_proto_depIdxs = []int32{}

func init() { file_test3_test_import_proto_init() }
func file_test3_test_import_proto_init() {
	if File_test3_test_import_proto != nil {
		return
	}
	File_test3_test_import_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_test3_test_import_proto_rawDesc,
		GoTypes:            file_test3_test_import_proto_goTypes,
		DependencyIndexes:  file_test3_test_import_proto_depIdxs,
		EnumOutputTypes:    file_test3_test_import_proto_enumTypes,
		MessageOutputTypes: file_test3_test_import_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_test3_test_import_proto_rawDesc = nil
	file_test3_test_import_proto_goTypes = nil
	file_test3_test_import_proto_depIdxs = nil
}

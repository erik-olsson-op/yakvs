// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: model/entry.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Keyword int32

const (
	Keyword_SET    Keyword = 0
	Keyword_GET    Keyword = 1
	Keyword_DELETE Keyword = 2
)

// Enum value maps for Keyword.
var (
	Keyword_name = map[int32]string{
		0: "SET",
		1: "GET",
		2: "DELETE",
	}
	Keyword_value = map[string]int32{
		"SET":    0,
		"GET":    1,
		"DELETE": 2,
	}
)

func (x Keyword) Enum() *Keyword {
	p := new(Keyword)
	*p = x
	return p
}

func (x Keyword) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Keyword) Descriptor() protoreflect.EnumDescriptor {
	return file_model_entry_proto_enumTypes[0].Descriptor()
}

func (Keyword) Type() protoreflect.EnumType {
	return &file_model_entry_proto_enumTypes[0]
}

func (x Keyword) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Keyword.Descriptor instead.
func (Keyword) EnumDescriptor() ([]byte, []int) {
	return file_model_entry_proto_rawDescGZIP(), []int{0}
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword Keyword `protobuf:"varint,1,opt,name=keyword,proto3,enum=yakvs.Keyword" json:"keyword,omitempty"`
	Entry   *Entry  `protobuf:"bytes,2,opt,name=entry,proto3,oneof" json:"entry,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	mi := &file_model_entry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_model_entry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_model_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Query) GetKeyword() Keyword {
	if x != nil {
		return x.Keyword
	}
	return Keyword_SET
}

func (x *Query) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	mi := &file_model_entry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_model_entry_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_model_entry_proto_rawDescGZIP(), []int{1}
}

func (x *Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Entry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_model_entry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_model_entry_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_model_entry_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_model_entry_proto protoreflect.FileDescriptor

var file_model_entry_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x79, 0x61, 0x6b, 0x76, 0x73, 0x22, 0x64, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x79, 0x61, 0x6b, 0x76, 0x73, 0x2e, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x27, 0x0a,
	0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x79,
	0x61, 0x6b, 0x76, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x22, 0x2f, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x27, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x07,
	0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x32, 0x35, 0x0a, 0x09,
	0x59, 0x61, 0x6b, 0x76, 0x73, 0x47, 0x72, 0x70, 0x63, 0x12, 0x28, 0x0a, 0x07, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x79, 0x61, 0x6b, 0x76, 0x73, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x79, 0x61, 0x6b, 0x76, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x72, 0x69, 0x6b, 0x2d, 0x6f, 0x6c, 0x73, 0x73, 0x6f, 0x6e, 0x2d, 0x6f, 0x70,
	0x2f, 0x79, 0x61, 0x6b, 0x76, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_entry_proto_rawDescOnce sync.Once
	file_model_entry_proto_rawDescData = file_model_entry_proto_rawDesc
)

func file_model_entry_proto_rawDescGZIP() []byte {
	file_model_entry_proto_rawDescOnce.Do(func() {
		file_model_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_entry_proto_rawDescData)
	})
	return file_model_entry_proto_rawDescData
}

var file_model_entry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_entry_proto_goTypes = []any{
	(Keyword)(0),     // 0: yakvs.Keyword
	(*Query)(nil),    // 1: yakvs.Query
	(*Entry)(nil),    // 2: yakvs.Entry
	(*Response)(nil), // 3: yakvs.Response
}
var file_model_entry_proto_depIdxs = []int32{
	0, // 0: yakvs.Query.keyword:type_name -> yakvs.Keyword
	2, // 1: yakvs.Query.entry:type_name -> yakvs.Entry
	1, // 2: yakvs.YakvsGrpc.Execute:input_type -> yakvs.Query
	3, // 3: yakvs.YakvsGrpc.Execute:output_type -> yakvs.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_entry_proto_init() }
func file_model_entry_proto_init() {
	if File_model_entry_proto != nil {
		return
	}
	file_model_entry_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_entry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_model_entry_proto_goTypes,
		DependencyIndexes: file_model_entry_proto_depIdxs,
		EnumInfos:         file_model_entry_proto_enumTypes,
		MessageInfos:      file_model_entry_proto_msgTypes,
	}.Build()
	File_model_entry_proto = out.File
	file_model_entry_proto_rawDesc = nil
	file_model_entry_proto_goTypes = nil
	file_model_entry_proto_depIdxs = nil
}

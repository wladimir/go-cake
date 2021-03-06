// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cake/cake.proto

package cake

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProblemInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProblemInput) Reset() {
	*x = ProblemInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cake_cake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemInput) ProtoMessage() {}

func (x *ProblemInput) ProtoReflect() protoreflect.Message {
	mi := &file_cake_cake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemInput.ProtoReflect.Descriptor instead.
func (*ProblemInput) Descriptor() ([]byte, []int) {
	return file_cake_cake_proto_rawDescGZIP(), []int{0}
}

type ProblemDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProblemDefinition) Reset() {
	*x = ProblemDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cake_cake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemDefinition) ProtoMessage() {}

func (x *ProblemDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_cake_cake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemDefinition.ProtoReflect.Descriptor instead.
func (*ProblemDefinition) Descriptor() ([]byte, []int) {
	return file_cake_cake_proto_rawDescGZIP(), []int{1}
}

func (x *ProblemDefinition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProblemResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Output  string `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *ProblemResult) Reset() {
	*x = ProblemResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cake_cake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProblemResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemResult) ProtoMessage() {}

func (x *ProblemResult) ProtoReflect() protoreflect.Message {
	mi := &file_cake_cake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemResult.ProtoReflect.Descriptor instead.
func (*ProblemResult) Descriptor() ([]byte, []int) {
	return file_cake_cake_proto_rawDescGZIP(), []int{2}
}

func (x *ProblemResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProblemResult) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ProblemResult) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

var File_cake_cake_proto protoreflect.FileDescriptor

var file_cake_cake_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6b, 0x65, 0x2f, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x61, 0x6b, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x27, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x55, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0x85, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x6b, 0x65,
	0x12, 0x3f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73,
	0x12, 0x12, 0x2e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12,
	0x17, 0x2e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x44, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x2e, 0x63, 0x61, 0x6b, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6c,
	0x61, 0x64, 0x69, 0x6d, 0x69, 0x72, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x61, 0x6b, 0x65, 0x2f, 0x63,
	0x61, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cake_cake_proto_rawDescOnce sync.Once
	file_cake_cake_proto_rawDescData = file_cake_cake_proto_rawDesc
)

func file_cake_cake_proto_rawDescGZIP() []byte {
	file_cake_cake_proto_rawDescOnce.Do(func() {
		file_cake_cake_proto_rawDescData = protoimpl.X.CompressGZIP(file_cake_cake_proto_rawDescData)
	})
	return file_cake_cake_proto_rawDescData
}

var file_cake_cake_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cake_cake_proto_goTypes = []interface{}{
	(*ProblemInput)(nil),      // 0: cake.ProblemInput
	(*ProblemDefinition)(nil), // 1: cake.ProblemDefinition
	(*ProblemResult)(nil),     // 2: cake.ProblemResult
}
var file_cake_cake_proto_depIdxs = []int32{
	0, // 0: cake.Cake.ListProblems:input_type -> cake.ProblemInput
	1, // 1: cake.Cake.RunProblem:input_type -> cake.ProblemDefinition
	1, // 2: cake.Cake.ListProblems:output_type -> cake.ProblemDefinition
	2, // 3: cake.Cake.RunProblem:output_type -> cake.ProblemResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cake_cake_proto_init() }
func file_cake_cake_proto_init() {
	if File_cake_cake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cake_cake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cake_cake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemDefinition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cake_cake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProblemResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cake_cake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cake_cake_proto_goTypes,
		DependencyIndexes: file_cake_cake_proto_depIdxs,
		MessageInfos:      file_cake_cake_proto_msgTypes,
	}.Build()
	File_cake_cake_proto = out.File
	file_cake_cake_proto_rawDesc = nil
	file_cake_cake_proto_goTypes = nil
	file_cake_cake_proto_depIdxs = nil
}

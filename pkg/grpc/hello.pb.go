// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: api/v1/hello.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetHelloResponse) Reset() {
	*x = GetHelloResponse{}
	mi := &file_api_v1_hello_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHelloResponse) ProtoMessage() {}

func (x *GetHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_hello_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHelloResponse.ProtoReflect.Descriptor instead.
func (*GetHelloResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *GetHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_v1_hello_proto protoreflect.FileDescriptor

var file_api_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4c, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_hello_proto_rawDescOnce sync.Once
	file_api_v1_hello_proto_rawDescData = file_api_v1_hello_proto_rawDesc
)

func file_api_v1_hello_proto_rawDescGZIP() []byte {
	file_api_v1_hello_proto_rawDescOnce.Do(func() {
		file_api_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_hello_proto_rawDescData)
	})
	return file_api_v1_hello_proto_rawDescData
}

var file_api_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_v1_hello_proto_goTypes = []any{
	(*GetHelloResponse)(nil), // 0: api.v1.GetHelloResponse
	(*emptypb.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_api_v1_hello_proto_depIdxs = []int32{
	1, // 0: api.v1.HelloService.GetHello:input_type -> google.protobuf.Empty
	0, // 1: api.v1.HelloService.GetHello:output_type -> api.v1.GetHelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_hello_proto_init() }
func file_api_v1_hello_proto_init() {
	if File_api_v1_hello_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_hello_proto_goTypes,
		DependencyIndexes: file_api_v1_hello_proto_depIdxs,
		MessageInfos:      file_api_v1_hello_proto_msgTypes,
	}.Build()
	File_api_v1_hello_proto = out.File
	file_api_v1_hello_proto_rawDesc = nil
	file_api_v1_hello_proto_goTypes = nil
	file_api_v1_hello_proto_depIdxs = nil
}

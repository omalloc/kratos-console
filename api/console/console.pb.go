// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: console/console.proto

package console

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_console_console_proto protoreflect.FileDescriptor

var file_console_console_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x1a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x09, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x42, 0xc3, 0x02,
	0xba, 0x47, 0xf9, 0x01, 0x12, 0xc1, 0x01, 0x0a, 0x15, 0x47, 0x6f, 0x2d, 0x4b, 0x72, 0x61, 0x74,
	0x6f, 0x73, 0x20, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x20, 0x41, 0x50, 0x49, 0x12, 0x20,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x47, 0x6f, 0x2d, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x22, 0x36, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x79, 0x61, 0x12, 0x19, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x79, 0x61, 0x1a, 0x11, 0x79, 0x6c, 0x61, 0x64, 0x6d, 0x78, 0x61, 0x40, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x47, 0x0a, 0x0e, 0x41, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x62, 0x6c,
	0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53,
	0x45, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0x2a, 0x33, 0x3a, 0x31, 0x0a, 0x2f, 0x0a, 0x0a,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x1f, 0x0a, 0x06,
	0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x3b, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_console_console_proto_goTypes = []interface{}{}
var file_console_console_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_console_console_proto_init() }
func file_console_console_proto_init() {
	if File_console_console_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_console_console_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_console_console_proto_goTypes,
		DependencyIndexes: file_console_console_proto_depIdxs,
	}.Build()
	File_console_console_proto = out.File
	file_console_console_proto_rawDesc = nil
	file_console_console_proto_goTypes = nil
	file_console_console_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/cors.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_heyrenee_v1_cors_proto protoreflect.FileDescriptor

var file_heyrenee_v1_cors_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x5c, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x42, 0x0f, 0x0a,
	0x07, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x12, 0x04, 0x2f, 0x2a, 0x2f, 0x2a, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79,
	0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_heyrenee_v1_cors_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil), // 0: google.protobuf.Empty
}
var file_heyrenee_v1_cors_proto_depIdxs = []int32{
	0, // 0: heyrenee.v1.Cors.Preflight:input_type -> google.protobuf.Empty
	0, // 1: heyrenee.v1.Cors.Preflight:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_cors_proto_init() }
func file_heyrenee_v1_cors_proto_init() {
	if File_heyrenee_v1_cors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_heyrenee_v1_cors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_cors_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_cors_proto_depIdxs,
	}.Build()
	File_heyrenee_v1_cors_proto = out.File
	file_heyrenee_v1_cors_proto_rawDesc = nil
	file_heyrenee_v1_cors_proto_goTypes = nil
	file_heyrenee_v1_cors_proto_depIdxs = nil
}

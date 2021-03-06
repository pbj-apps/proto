// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/enums/specialty.proto

package enums

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

// Specialty specifies a healthcare speciality.
type Specialty int32

const (
	// The specialty is not specified.
	Specialty_SPECIALTY_UNSPECIFIED Specialty = 0
	// The specialty is Endocrinology.
	Specialty_SPECIALTY_ENDOCRINOLOGY Specialty = 1
)

// Enum value maps for Specialty.
var (
	Specialty_name = map[int32]string{
		0: "SPECIALTY_UNSPECIFIED",
		1: "SPECIALTY_ENDOCRINOLOGY",
	}
	Specialty_value = map[string]int32{
		"SPECIALTY_UNSPECIFIED":   0,
		"SPECIALTY_ENDOCRINOLOGY": 1,
	}
)

func (x Specialty) Enum() *Specialty {
	p := new(Specialty)
	*p = x
	return p
}

func (x Specialty) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Specialty) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_enums_specialty_proto_enumTypes[0].Descriptor()
}

func (Specialty) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_enums_specialty_proto_enumTypes[0]
}

func (x Specialty) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Specialty.Descriptor instead.
func (Specialty) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_enums_specialty_proto_rawDescGZIP(), []int{0}
}

var File_heyrenee_v1_enums_specialty_proto protoreflect.FileDescriptor

var file_heyrenee_v1_enums_specialty_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0x43, 0x0a, 0x09, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x54, 0x59,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b,
	0x0a, 0x17, 0x53, 0x50, 0x45, 0x43, 0x49, 0x41, 0x4c, 0x54, 0x59, 0x5f, 0x45, 0x4e, 0x44, 0x4f,
	0x43, 0x52, 0x49, 0x4e, 0x4f, 0x4c, 0x4f, 0x47, 0x59, 0x10, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e,
	0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_enums_specialty_proto_rawDescOnce sync.Once
	file_heyrenee_v1_enums_specialty_proto_rawDescData = file_heyrenee_v1_enums_specialty_proto_rawDesc
)

func file_heyrenee_v1_enums_specialty_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_enums_specialty_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_enums_specialty_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_enums_specialty_proto_rawDescData)
	})
	return file_heyrenee_v1_enums_specialty_proto_rawDescData
}

var file_heyrenee_v1_enums_specialty_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_heyrenee_v1_enums_specialty_proto_goTypes = []interface{}{
	(Specialty)(0), // 0: heyrenee.v1.enums.Specialty
}
var file_heyrenee_v1_enums_specialty_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_enums_specialty_proto_init() }
func file_heyrenee_v1_enums_specialty_proto_init() {
	if File_heyrenee_v1_enums_specialty_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_heyrenee_v1_enums_specialty_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_enums_specialty_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_enums_specialty_proto_depIdxs,
		EnumInfos:         file_heyrenee_v1_enums_specialty_proto_enumTypes,
	}.Build()
	File_heyrenee_v1_enums_specialty_proto = out.File
	file_heyrenee_v1_enums_specialty_proto_rawDesc = nil
	file_heyrenee_v1_enums_specialty_proto_goTypes = nil
	file_heyrenee_v1_enums_specialty_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/caregiver_service.proto

package v1

import (
	messages "github.com/HeyReneeInc/proto/golang/heyrenee/v1/messages"
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

type CreateCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caregiver *messages.Caregiver `protobuf:"bytes,1,opt,name=caregiver,proto3" json:"caregiver,omitempty"`
}

func (x *CreateCaregiverRequest) Reset() {
	*x = CreateCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_caregiver_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCaregiverRequest) ProtoMessage() {}

func (x *CreateCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_caregiver_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCaregiverRequest.ProtoReflect.Descriptor instead.
func (*CreateCaregiverRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_caregiver_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCaregiverRequest) GetCaregiver() *messages.Caregiver {
	if x != nil {
		return x.Caregiver
	}
	return nil
}

type GetCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaregiverId string `protobuf:"bytes,1,opt,name=caregiver_id,json=caregiverId,proto3" json:"caregiver_id,omitempty"`
}

func (x *GetCaregiverRequest) Reset() {
	*x = GetCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_caregiver_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaregiverRequest) ProtoMessage() {}

func (x *GetCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_caregiver_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCaregiverRequest.ProtoReflect.Descriptor instead.
func (*GetCaregiverRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_caregiver_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCaregiverRequest) GetCaregiverId() string {
	if x != nil {
		return x.CaregiverId
	}
	return ""
}

type UpdateCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caregiver *messages.Caregiver `protobuf:"bytes,1,opt,name=caregiver,proto3" json:"caregiver,omitempty"`
}

func (x *UpdateCaregiverRequest) Reset() {
	*x = UpdateCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_caregiver_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCaregiverRequest) ProtoMessage() {}

func (x *UpdateCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_caregiver_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCaregiverRequest.ProtoReflect.Descriptor instead.
func (*UpdateCaregiverRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_caregiver_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCaregiverRequest) GetCaregiver() *messages.Caregiver {
	if x != nil {
		return x.Caregiver
	}
	return nil
}

var File_heyrenee_v1_caregiver_service_proto protoreflect.FileDescriptor

var file_heyrenee_v1_caregiver_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x24, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72,
	0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67,
	0x69, 0x76, 0x65, 0x72, 0x32, 0x9d, 0x02, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x0f, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_caregiver_service_proto_rawDescOnce sync.Once
	file_heyrenee_v1_caregiver_service_proto_rawDescData = file_heyrenee_v1_caregiver_service_proto_rawDesc
)

func file_heyrenee_v1_caregiver_service_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_caregiver_service_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_caregiver_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_caregiver_service_proto_rawDescData)
	})
	return file_heyrenee_v1_caregiver_service_proto_rawDescData
}

var file_heyrenee_v1_caregiver_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_heyrenee_v1_caregiver_service_proto_goTypes = []interface{}{
	(*CreateCaregiverRequest)(nil), // 0: heyrenee.v1.CreateCaregiverRequest
	(*GetCaregiverRequest)(nil),    // 1: heyrenee.v1.GetCaregiverRequest
	(*UpdateCaregiverRequest)(nil), // 2: heyrenee.v1.UpdateCaregiverRequest
	(*messages.Caregiver)(nil),     // 3: heyrenee.v1.messages.Caregiver
}
var file_heyrenee_v1_caregiver_service_proto_depIdxs = []int32{
	3, // 0: heyrenee.v1.CreateCaregiverRequest.caregiver:type_name -> heyrenee.v1.messages.Caregiver
	3, // 1: heyrenee.v1.UpdateCaregiverRequest.caregiver:type_name -> heyrenee.v1.messages.Caregiver
	0, // 2: heyrenee.v1.CaregiverService.CreateCaregiver:input_type -> heyrenee.v1.CreateCaregiverRequest
	1, // 3: heyrenee.v1.CaregiverService.GetCaregiver:input_type -> heyrenee.v1.GetCaregiverRequest
	2, // 4: heyrenee.v1.CaregiverService.UpdateCaregiver:input_type -> heyrenee.v1.UpdateCaregiverRequest
	3, // 5: heyrenee.v1.CaregiverService.CreateCaregiver:output_type -> heyrenee.v1.messages.Caregiver
	3, // 6: heyrenee.v1.CaregiverService.GetCaregiver:output_type -> heyrenee.v1.messages.Caregiver
	3, // 7: heyrenee.v1.CaregiverService.UpdateCaregiver:output_type -> heyrenee.v1.messages.Caregiver
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_caregiver_service_proto_init() }
func file_heyrenee_v1_caregiver_service_proto_init() {
	if File_heyrenee_v1_caregiver_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_caregiver_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCaregiverRequest); i {
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
		file_heyrenee_v1_caregiver_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCaregiverRequest); i {
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
		file_heyrenee_v1_caregiver_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCaregiverRequest); i {
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
			RawDescriptor: file_heyrenee_v1_caregiver_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_caregiver_service_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_caregiver_service_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_caregiver_service_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_caregiver_service_proto = out.File
	file_heyrenee_v1_caregiver_service_proto_rawDesc = nil
	file_heyrenee_v1_caregiver_service_proto_goTypes = nil
	file_heyrenee_v1_caregiver_service_proto_depIdxs = nil
}

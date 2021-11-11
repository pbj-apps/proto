// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/appointment_service.proto

package v1

import (
	messages "github.com/HeyReneeInc/proto/golang/heyrenee/v1/messages"
	_ "github.com/HeyReneeInc/proto/golang/heyrenee/v1/options"
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

// Specifies the type of Appointment.
type AppointmentType int32

const (
	// No Appointment type specified.
	AppointmentType_APPOINTMENT_TYPE_UNSPECIFIED AppointmentType = 0
	// The Appointment is upcoming, that is it's scheduled datetime is >= the current datetime.
	AppointmentType_APPOINTMENT_UPCOMING AppointmentType = 1
)

// Enum value maps for AppointmentType.
var (
	AppointmentType_name = map[int32]string{
		0: "APPOINTMENT_TYPE_UNSPECIFIED",
		1: "APPOINTMENT_UPCOMING",
	}
	AppointmentType_value = map[string]int32{
		"APPOINTMENT_TYPE_UNSPECIFIED": 0,
		"APPOINTMENT_UPCOMING":         1,
	}
)

func (x AppointmentType) Enum() *AppointmentType {
	p := new(AppointmentType)
	*p = x
	return p
}

func (x AppointmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppointmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_appointment_service_proto_enumTypes[0].Descriptor()
}

func (AppointmentType) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_appointment_service_proto_enumTypes[0]
}

func (x AppointmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppointmentType.Descriptor instead.
func (AppointmentType) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_appointment_service_proto_rawDescGZIP(), []int{0}
}

// ListAppointmentsRequest is the request message for the ListAppointments method.
type ListAppointmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Patient to return Appointments for. Required.
	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	// The type of Appointments that should be returned.
	AppointmentType AppointmentType `protobuf:"varint,2,opt,name=appointment_type,json=appointmentType,proto3,enum=heyrenee.v1.AppointmentType" json:"appointment_type,omitempty"`
}

func (x *ListAppointmentsRequest) Reset() {
	*x = ListAppointmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_appointment_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppointmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppointmentsRequest) ProtoMessage() {}

func (x *ListAppointmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_appointment_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppointmentsRequest.ProtoReflect.Descriptor instead.
func (*ListAppointmentsRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_appointment_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListAppointmentsRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *ListAppointmentsRequest) GetAppointmentType() AppointmentType {
	if x != nil {
		return x.AppointmentType
	}
	return AppointmentType_APPOINTMENT_TYPE_UNSPECIFIED
}

// ListAppointmentsResponse is the response message for the ListAppointments method.
type ListAppointmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of Appointments returned.
	Appointments []*messages.Appointment `protobuf:"bytes,1,rep,name=appointments,proto3" json:"appointments,omitempty"`
}

func (x *ListAppointmentsResponse) Reset() {
	*x = ListAppointmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_appointment_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppointmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppointmentsResponse) ProtoMessage() {}

func (x *ListAppointmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_appointment_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppointmentsResponse.ProtoReflect.Descriptor instead.
func (*ListAppointmentsResponse) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_appointment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListAppointmentsResponse) GetAppointments() []*messages.Appointment {
	if x != nil {
		return x.Appointments
	}
	return nil
}

var File_heyrenee_v1_appointment_service_proto protoreflect.FileDescriptor

var file_heyrenee_v1_appointment_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x26, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x68, 0x65,
	0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52,
	0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x10, 0x61, 0x70,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x61, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x4d, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x50, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41,
	0x50, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x50, 0x43, 0x4f, 0x4d,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x32, 0xba, 0x01, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d,
	0x43, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x65, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x61, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_appointment_service_proto_rawDescOnce sync.Once
	file_heyrenee_v1_appointment_service_proto_rawDescData = file_heyrenee_v1_appointment_service_proto_rawDesc
)

func file_heyrenee_v1_appointment_service_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_appointment_service_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_appointment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_appointment_service_proto_rawDescData)
	})
	return file_heyrenee_v1_appointment_service_proto_rawDescData
}

var file_heyrenee_v1_appointment_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_heyrenee_v1_appointment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_heyrenee_v1_appointment_service_proto_goTypes = []interface{}{
	(AppointmentType)(0),             // 0: heyrenee.v1.AppointmentType
	(*ListAppointmentsRequest)(nil),  // 1: heyrenee.v1.ListAppointmentsRequest
	(*ListAppointmentsResponse)(nil), // 2: heyrenee.v1.ListAppointmentsResponse
	(*messages.Appointment)(nil),     // 3: heyrenee.v1.messages.Appointment
	(*emptypb.Empty)(nil),            // 4: google.protobuf.Empty
}
var file_heyrenee_v1_appointment_service_proto_depIdxs = []int32{
	0, // 0: heyrenee.v1.ListAppointmentsRequest.appointment_type:type_name -> heyrenee.v1.AppointmentType
	3, // 1: heyrenee.v1.ListAppointmentsResponse.appointments:type_name -> heyrenee.v1.messages.Appointment
	4, // 2: heyrenee.v1.AppointmentService.CorsPreflight:input_type -> google.protobuf.Empty
	1, // 3: heyrenee.v1.AppointmentService.ListAppointments:input_type -> heyrenee.v1.ListAppointmentsRequest
	4, // 4: heyrenee.v1.AppointmentService.CorsPreflight:output_type -> google.protobuf.Empty
	2, // 5: heyrenee.v1.AppointmentService.ListAppointments:output_type -> heyrenee.v1.ListAppointmentsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_appointment_service_proto_init() }
func file_heyrenee_v1_appointment_service_proto_init() {
	if File_heyrenee_v1_appointment_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_appointment_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppointmentsRequest); i {
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
		file_heyrenee_v1_appointment_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppointmentsResponse); i {
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
			RawDescriptor: file_heyrenee_v1_appointment_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_appointment_service_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_appointment_service_proto_depIdxs,
		EnumInfos:         file_heyrenee_v1_appointment_service_proto_enumTypes,
		MessageInfos:      file_heyrenee_v1_appointment_service_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_appointment_service_proto = out.File
	file_heyrenee_v1_appointment_service_proto_rawDesc = nil
	file_heyrenee_v1_appointment_service_proto_goTypes = nil
	file_heyrenee_v1_appointment_service_proto_depIdxs = nil
}

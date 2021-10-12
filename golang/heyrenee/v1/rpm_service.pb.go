// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/rpm_service.proto

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

type ListRpmSchedulesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId         string                     `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	RpmScheduleStatus messages.RpmScheduleStatus `protobuf:"varint,2,opt,name=rpm_schedule_status,json=rpmScheduleStatus,proto3,enum=heyrenee.v1.messages.RpmScheduleStatus" json:"rpm_schedule_status,omitempty"`
}

func (x *ListRpmSchedulesRequest) Reset() {
	*x = ListRpmSchedulesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpmSchedulesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpmSchedulesRequest) ProtoMessage() {}

func (x *ListRpmSchedulesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpmSchedulesRequest.ProtoReflect.Descriptor instead.
func (*ListRpmSchedulesRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_rpm_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListRpmSchedulesRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *ListRpmSchedulesRequest) GetRpmScheduleStatus() messages.RpmScheduleStatus {
	if x != nil {
		return x.RpmScheduleStatus
	}
	return messages.RpmScheduleStatus(0)
}

type ListRpmSchedulesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpmSchedules []*messages.RpmSchedule `protobuf:"bytes,1,rep,name=rpm_schedules,json=rpmSchedules,proto3" json:"rpm_schedules,omitempty"`
}

func (x *ListRpmSchedulesResponse) Reset() {
	*x = ListRpmSchedulesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpmSchedulesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpmSchedulesResponse) ProtoMessage() {}

func (x *ListRpmSchedulesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpmSchedulesResponse.ProtoReflect.Descriptor instead.
func (*ListRpmSchedulesResponse) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_rpm_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListRpmSchedulesResponse) GetRpmSchedules() []*messages.RpmSchedule {
	if x != nil {
		return x.RpmSchedules
	}
	return nil
}

type ListRpmMeasurementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpmScheduleId string `protobuf:"bytes,1,opt,name=rpm_schedule_id,json=rpmScheduleId,proto3" json:"rpm_schedule_id,omitempty"`
}

func (x *ListRpmMeasurementsRequest) Reset() {
	*x = ListRpmMeasurementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpmMeasurementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpmMeasurementsRequest) ProtoMessage() {}

func (x *ListRpmMeasurementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpmMeasurementsRequest.ProtoReflect.Descriptor instead.
func (*ListRpmMeasurementsRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_rpm_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListRpmMeasurementsRequest) GetRpmScheduleId() string {
	if x != nil {
		return x.RpmScheduleId
	}
	return ""
}

type ListRpmMeasurementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RpmMeasurements []*messages.RpmMeasurement `protobuf:"bytes,1,rep,name=rpm_measurements,json=rpmMeasurements,proto3" json:"rpm_measurements,omitempty"`
}

func (x *ListRpmMeasurementsResponse) Reset() {
	*x = ListRpmMeasurementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRpmMeasurementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRpmMeasurementsResponse) ProtoMessage() {}

func (x *ListRpmMeasurementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_rpm_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRpmMeasurementsResponse.ProtoReflect.Descriptor instead.
func (*ListRpmMeasurementsResponse) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_rpm_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListRpmMeasurementsResponse) GetRpmMeasurements() []*messages.RpmMeasurement {
	if x != nil {
		return x.RpmMeasurements
	}
	return nil
}

var File_heyrenee_v1_rpm_service_proto protoreflect.FileDescriptor

var file_heyrenee_v1_rpm_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x70,
	0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x68, 0x65,
	0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x72, 0x70, 0x6d, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x70, 0x6d, 0x5f,
	0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x91, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x57, 0x0a, 0x13,
	0x72, 0x70, 0x6d, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x52, 0x70, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x11, 0x72, 0x70, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x62, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x0d, 0x72, 0x70, 0x6d, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x52, 0x70, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0c, 0x72, 0x70, 0x6d,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x22, 0x44, 0x0a, 0x1a, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x70, 0x6d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x70, 0x6d, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x70, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x6e, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x10, 0x72, 0x70, 0x6d, 0x5f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x52, 0x70, 0x6d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0f,
	0x72, 0x70, 0x6d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32,
	0xdb, 0x01, 0x0a, 0x0a, 0x52, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x24, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x70, 0x6d, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52,
	0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_rpm_service_proto_rawDescOnce sync.Once
	file_heyrenee_v1_rpm_service_proto_rawDescData = file_heyrenee_v1_rpm_service_proto_rawDesc
)

func file_heyrenee_v1_rpm_service_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_rpm_service_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_rpm_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_rpm_service_proto_rawDescData)
	})
	return file_heyrenee_v1_rpm_service_proto_rawDescData
}

var file_heyrenee_v1_rpm_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_heyrenee_v1_rpm_service_proto_goTypes = []interface{}{
	(*ListRpmSchedulesRequest)(nil),     // 0: heyrenee.v1.ListRpmSchedulesRequest
	(*ListRpmSchedulesResponse)(nil),    // 1: heyrenee.v1.ListRpmSchedulesResponse
	(*ListRpmMeasurementsRequest)(nil),  // 2: heyrenee.v1.ListRpmMeasurementsRequest
	(*ListRpmMeasurementsResponse)(nil), // 3: heyrenee.v1.ListRpmMeasurementsResponse
	(messages.RpmScheduleStatus)(0),     // 4: heyrenee.v1.messages.RpmScheduleStatus
	(*messages.RpmSchedule)(nil),        // 5: heyrenee.v1.messages.RpmSchedule
	(*messages.RpmMeasurement)(nil),     // 6: heyrenee.v1.messages.RpmMeasurement
}
var file_heyrenee_v1_rpm_service_proto_depIdxs = []int32{
	4, // 0: heyrenee.v1.ListRpmSchedulesRequest.rpm_schedule_status:type_name -> heyrenee.v1.messages.RpmScheduleStatus
	5, // 1: heyrenee.v1.ListRpmSchedulesResponse.rpm_schedules:type_name -> heyrenee.v1.messages.RpmSchedule
	6, // 2: heyrenee.v1.ListRpmMeasurementsResponse.rpm_measurements:type_name -> heyrenee.v1.messages.RpmMeasurement
	0, // 3: heyrenee.v1.RpmService.ListRpmSchedules:input_type -> heyrenee.v1.ListRpmSchedulesRequest
	2, // 4: heyrenee.v1.RpmService.ListRpmMeasurements:input_type -> heyrenee.v1.ListRpmMeasurementsRequest
	1, // 5: heyrenee.v1.RpmService.ListRpmSchedules:output_type -> heyrenee.v1.ListRpmSchedulesResponse
	3, // 6: heyrenee.v1.RpmService.ListRpmMeasurements:output_type -> heyrenee.v1.ListRpmMeasurementsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_rpm_service_proto_init() }
func file_heyrenee_v1_rpm_service_proto_init() {
	if File_heyrenee_v1_rpm_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_rpm_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpmSchedulesRequest); i {
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
		file_heyrenee_v1_rpm_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpmSchedulesResponse); i {
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
		file_heyrenee_v1_rpm_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpmMeasurementsRequest); i {
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
		file_heyrenee_v1_rpm_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRpmMeasurementsResponse); i {
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
			RawDescriptor: file_heyrenee_v1_rpm_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_rpm_service_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_rpm_service_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_rpm_service_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_rpm_service_proto = out.File
	file_heyrenee_v1_rpm_service_proto_rawDesc = nil
	file_heyrenee_v1_rpm_service_proto_goTypes = nil
	file_heyrenee_v1_rpm_service_proto_depIdxs = nil
}

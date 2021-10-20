// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/messages/medication_dose.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A MedicationDose represents a single dose of Medication taken by a Patient as specified in a Prescription.
type MedicationDose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Patient who took the MedicationDose.
	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	// The ID of the Medication that was taken.
	MedicationId string `protobuf:"bytes,2,opt,name=medication_id,json=medicationId,proto3" json:"medication_id,omitempty"`
	// The ID of the Prescription that prescribed the Medication that was taken.
	PrescriptionId string `protobuf:"bytes,3,opt,name=prescription_id,json=prescriptionId,proto3" json:"prescription_id,omitempty"`
	// The ID of the MedicationDose.
	MedicationDoseId string `protobuf:"bytes,4,opt,name=medication_dose_id,json=medicationDoseId,proto3" json:"medication_dose_id,omitempty"`
	// The time at which the MedicationDose was taken.
	TimeTaken *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time_taken,json=timeTaken,proto3" json:"time_taken,omitempty"`
}

func (x *MedicationDose) Reset() {
	*x = MedicationDose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_messages_medication_dose_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MedicationDose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MedicationDose) ProtoMessage() {}

func (x *MedicationDose) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_messages_medication_dose_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MedicationDose.ProtoReflect.Descriptor instead.
func (*MedicationDose) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_medication_dose_proto_rawDescGZIP(), []int{0}
}

func (x *MedicationDose) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *MedicationDose) GetMedicationId() string {
	if x != nil {
		return x.MedicationId
	}
	return ""
}

func (x *MedicationDose) GetPrescriptionId() string {
	if x != nil {
		return x.PrescriptionId
	}
	return ""
}

func (x *MedicationDose) GetMedicationDoseId() string {
	if x != nil {
		return x.MedicationDoseId
	}
	return ""
}

func (x *MedicationDose) GetTimeTaken() *timestamppb.Timestamp {
	if x != nil {
		return x.TimeTaken
	}
	return nil
}

var File_heyrenee_v1_messages_medication_dose_proto protoreflect.FileDescriptor

var file_heyrenee_v1_messages_medication_dose_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x6f, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x68, 0x65,
	0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x6f, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x6f, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x73, 0x65, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65,
	0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_messages_medication_dose_proto_rawDescOnce sync.Once
	file_heyrenee_v1_messages_medication_dose_proto_rawDescData = file_heyrenee_v1_messages_medication_dose_proto_rawDesc
)

func file_heyrenee_v1_messages_medication_dose_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_messages_medication_dose_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_messages_medication_dose_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_messages_medication_dose_proto_rawDescData)
	})
	return file_heyrenee_v1_messages_medication_dose_proto_rawDescData
}

var file_heyrenee_v1_messages_medication_dose_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heyrenee_v1_messages_medication_dose_proto_goTypes = []interface{}{
	(*MedicationDose)(nil),        // 0: heyrenee.v1.messages.MedicationDose
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_heyrenee_v1_messages_medication_dose_proto_depIdxs = []int32{
	1, // 0: heyrenee.v1.messages.MedicationDose.time_taken:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_messages_medication_dose_proto_init() }
func file_heyrenee_v1_messages_medication_dose_proto_init() {
	if File_heyrenee_v1_messages_medication_dose_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_messages_medication_dose_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MedicationDose); i {
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
			RawDescriptor: file_heyrenee_v1_messages_medication_dose_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_messages_medication_dose_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_messages_medication_dose_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_messages_medication_dose_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_messages_medication_dose_proto = out.File
	file_heyrenee_v1_messages_medication_dose_proto_rawDesc = nil
	file_heyrenee_v1_messages_medication_dose_proto_goTypes = nil
	file_heyrenee_v1_messages_medication_dose_proto_depIdxs = nil
}

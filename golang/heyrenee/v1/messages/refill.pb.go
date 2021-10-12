// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/messages/refill.proto

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

type RefillType int32

const (
	RefillType_REFILL_TYPE_UNSPECIFIED RefillType = 0
	RefillType_REFILL_DELIVERY         RefillType = 1
	RefillType_REFILL_PICKUP           RefillType = 2
)

// Enum value maps for RefillType.
var (
	RefillType_name = map[int32]string{
		0: "REFILL_TYPE_UNSPECIFIED",
		1: "REFILL_DELIVERY",
		2: "REFILL_PICKUP",
	}
	RefillType_value = map[string]int32{
		"REFILL_TYPE_UNSPECIFIED": 0,
		"REFILL_DELIVERY":         1,
		"REFILL_PICKUP":           2,
	}
)

func (x RefillType) Enum() *RefillType {
	p := new(RefillType)
	*p = x
	return p
}

func (x RefillType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RefillType) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_refill_proto_enumTypes[0].Descriptor()
}

func (RefillType) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_refill_proto_enumTypes[0]
}

func (x RefillType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RefillType.Descriptor instead.
func (RefillType) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_refill_proto_rawDescGZIP(), []int{0}
}

type Refill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PatientId      string                 `protobuf:"bytes,2,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	MedicationId   string                 `protobuf:"bytes,3,opt,name=medication_id,json=medicationId,proto3" json:"medication_id,omitempty"`
	PrescriptionId string                 `protobuf:"bytes,4,opt,name=prescription_id,json=prescriptionId,proto3" json:"prescription_id,omitempty"`
	RefillId       string                 `protobuf:"bytes,5,opt,name=refill_id,json=refillId,proto3" json:"refill_id,omitempty"`
	DateWritten    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_written,json=dateWritten,proto3" json:"date_written,omitempty"`
	PharmacyId     string                 `protobuf:"bytes,9,opt,name=pharmacy_id,json=pharmacyId,proto3" json:"pharmacy_id,omitempty"`
	RefillType     RefillType             `protobuf:"varint,10,opt,name=refill_type,json=refillType,proto3,enum=heyrenee.v1.messages.RefillType" json:"refill_type,omitempty"`
}

func (x *Refill) Reset() {
	*x = Refill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_messages_refill_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Refill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Refill) ProtoMessage() {}

func (x *Refill) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_messages_refill_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Refill.ProtoReflect.Descriptor instead.
func (*Refill) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_refill_proto_rawDescGZIP(), []int{0}
}

func (x *Refill) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Refill) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *Refill) GetMedicationId() string {
	if x != nil {
		return x.MedicationId
	}
	return ""
}

func (x *Refill) GetPrescriptionId() string {
	if x != nil {
		return x.PrescriptionId
	}
	return ""
}

func (x *Refill) GetRefillId() string {
	if x != nil {
		return x.RefillId
	}
	return ""
}

func (x *Refill) GetDateWritten() *timestamppb.Timestamp {
	if x != nil {
		return x.DateWritten
	}
	return nil
}

func (x *Refill) GetPharmacyId() string {
	if x != nil {
		return x.PharmacyId
	}
	return ""
}

func (x *Refill) GetRefillType() RefillType {
	if x != nil {
		return x.RefillType
	}
	return RefillType_REFILL_TYPE_UNSPECIFIED
}

var File_heyrenee_v1_messages_refill_proto protoreflect.FileDescriptor

var file_heyrenee_v1_messages_refill_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x06, 0x52,
	0x65, 0x66, 0x69, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65,
	0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x68, 0x61, 0x72, 0x6d,
	0x61, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x68,
	0x61, 0x72, 0x6d, 0x61, 0x63, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x69,
	0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x51, 0x0a, 0x0a, 0x52,
	0x65, 0x66, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x46,
	0x49, 0x4c, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x46, 0x49, 0x4c, 0x4c,
	0x5f, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52,
	0x45, 0x46, 0x49, 0x4c, 0x4c, 0x5f, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x02, 0x42, 0x3a,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79,
	0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_heyrenee_v1_messages_refill_proto_rawDescOnce sync.Once
	file_heyrenee_v1_messages_refill_proto_rawDescData = file_heyrenee_v1_messages_refill_proto_rawDesc
)

func file_heyrenee_v1_messages_refill_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_messages_refill_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_messages_refill_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_messages_refill_proto_rawDescData)
	})
	return file_heyrenee_v1_messages_refill_proto_rawDescData
}

var file_heyrenee_v1_messages_refill_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_heyrenee_v1_messages_refill_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heyrenee_v1_messages_refill_proto_goTypes = []interface{}{
	(RefillType)(0),               // 0: heyrenee.v1.messages.RefillType
	(*Refill)(nil),                // 1: heyrenee.v1.messages.Refill
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_heyrenee_v1_messages_refill_proto_depIdxs = []int32{
	2, // 0: heyrenee.v1.messages.Refill.date_written:type_name -> google.protobuf.Timestamp
	0, // 1: heyrenee.v1.messages.Refill.refill_type:type_name -> heyrenee.v1.messages.RefillType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_messages_refill_proto_init() }
func file_heyrenee_v1_messages_refill_proto_init() {
	if File_heyrenee_v1_messages_refill_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_messages_refill_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Refill); i {
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
			RawDescriptor: file_heyrenee_v1_messages_refill_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_messages_refill_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_messages_refill_proto_depIdxs,
		EnumInfos:         file_heyrenee_v1_messages_refill_proto_enumTypes,
		MessageInfos:      file_heyrenee_v1_messages_refill_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_messages_refill_proto = out.File
	file_heyrenee_v1_messages_refill_proto_rawDesc = nil
	file_heyrenee_v1_messages_refill_proto_goTypes = nil
	file_heyrenee_v1_messages_refill_proto_depIdxs = nil
}

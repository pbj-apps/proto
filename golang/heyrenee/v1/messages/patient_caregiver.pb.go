// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/messages/patient_caregiver.proto

package messages

import (
	_ "github.com/HeyReneeInc/proto/golang/heyrenee/v1/options"
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

// Specifies a type of PatientCaregiver.
type PatientCaregiverType int32

const (
	// The type of PatientCaregiver is unspecified.
	PatientCaregiverType_PATIENT_CAREGIVER_TYPE_UNSPECIFIED PatientCaregiverType = 0
	// The Caregiver of the PatientCaregiver is the primary caregiver of the Patient.
	PatientCaregiverType_PATIENT_CAREGIVER_TYPE_PRIMARY PatientCaregiverType = 1
)

// Enum value maps for PatientCaregiverType.
var (
	PatientCaregiverType_name = map[int32]string{
		0: "PATIENT_CAREGIVER_TYPE_UNSPECIFIED",
		1: "PATIENT_CAREGIVER_TYPE_PRIMARY",
	}
	PatientCaregiverType_value = map[string]int32{
		"PATIENT_CAREGIVER_TYPE_UNSPECIFIED": 0,
		"PATIENT_CAREGIVER_TYPE_PRIMARY":     1,
	}
)

func (x PatientCaregiverType) Enum() *PatientCaregiverType {
	p := new(PatientCaregiverType)
	*p = x
	return p
}

func (x PatientCaregiverType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PatientCaregiverType) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes[0].Descriptor()
}

func (PatientCaregiverType) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes[0]
}

func (x PatientCaregiverType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PatientCaregiverType.Descriptor instead.
func (PatientCaregiverType) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_caregiver_proto_rawDescGZIP(), []int{0}
}

// Specifies the access that a Caregiver in a PatientCaregiver has to the Patient in a PatientCaregiver.
type PatientCaregiverAccess int32

const (
	// The access is not specified.
	PatientCaregiverAccess_PATIENT_CAREGIVER_ACCESS_UNSPECIFIED PatientCaregiverAccess = 0
	// The Caregiver has no access to the Patient.
	PatientCaregiverAccess_PATIENT_CAREGIVER_ACCESS_NONE PatientCaregiverAccess = 1
	// The Caregiver has access to view the Patient dashboard.
	PatientCaregiverAccess_PATIENT_CAREGIVER_ACCESS_DASHBOARD PatientCaregiverAccess = 2
)

// Enum value maps for PatientCaregiverAccess.
var (
	PatientCaregiverAccess_name = map[int32]string{
		0: "PATIENT_CAREGIVER_ACCESS_UNSPECIFIED",
		1: "PATIENT_CAREGIVER_ACCESS_NONE",
		2: "PATIENT_CAREGIVER_ACCESS_DASHBOARD",
	}
	PatientCaregiverAccess_value = map[string]int32{
		"PATIENT_CAREGIVER_ACCESS_UNSPECIFIED": 0,
		"PATIENT_CAREGIVER_ACCESS_NONE":        1,
		"PATIENT_CAREGIVER_ACCESS_DASHBOARD":   2,
	}
)

func (x PatientCaregiverAccess) Enum() *PatientCaregiverAccess {
	p := new(PatientCaregiverAccess)
	*p = x
	return p
}

func (x PatientCaregiverAccess) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PatientCaregiverAccess) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes[1].Descriptor()
}

func (PatientCaregiverAccess) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes[1]
}

func (x PatientCaregiverAccess) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PatientCaregiverAccess.Descriptor instead.
func (PatientCaregiverAccess) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_caregiver_proto_rawDescGZIP(), []int{1}
}

// Specifies the relationship of a Caregiver in a PatientCaregiver to the Patient in a PatientCaregiver.
type PatientCaregiverRelationship int32

const (
	// The relationship is not specified.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_UNSPECIFIED PatientCaregiverRelationship = 0
	// The Caregiver is the Patient's child.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_CHILD PatientCaregiverRelationship = 1
	// The Caregiver is the Patient's parent.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_PARENT PatientCaregiverRelationship = 2
	// The Caregiver is the Patient's grandparent.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_GRANDPARENT PatientCaregiverRelationship = 3
	// The Caregiver is the Patient's grandchild.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_GRANDCHILD PatientCaregiverRelationship = 4
	// The Caregiver is in the Patient's extended family.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_EXTENDED_FAMILY PatientCaregiverRelationship = 5
	// The Caregiver is the Patient's friend.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_FRIEND PatientCaregiverRelationship = 6
	// The Caregiver is the Patient's health aide.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_HEALTH_AIDE PatientCaregiverRelationship = 7
	// The Caregiver has another relationship to the Patient.
	PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_OTHER PatientCaregiverRelationship = 8
)

// Enum value maps for PatientCaregiverRelationship.
var (
	PatientCaregiverRelationship_name = map[int32]string{
		0: "PATIENT_CAREGIVER_RELATIONSHIP_UNSPECIFIED",
		1: "PATIENT_CAREGIVER_RELATIONSHIP_CHILD",
		2: "PATIENT_CAREGIVER_RELATIONSHIP_PARENT",
		3: "PATIENT_CAREGIVER_RELATIONSHIP_GRANDPARENT",
		4: "PATIENT_CAREGIVER_RELATIONSHIP_GRANDCHILD",
		5: "PATIENT_CAREGIVER_RELATIONSHIP_EXTENDED_FAMILY",
		6: "PATIENT_CAREGIVER_RELATIONSHIP_FRIEND",
		7: "PATIENT_CAREGIVER_RELATIONSHIP_HEALTH_AIDE",
		8: "PATIENT_CAREGIVER_RELATIONSHIP_OTHER",
	}
	PatientCaregiverRelationship_value = map[string]int32{
		"PATIENT_CAREGIVER_RELATIONSHIP_UNSPECIFIED":     0,
		"PATIENT_CAREGIVER_RELATIONSHIP_CHILD":           1,
		"PATIENT_CAREGIVER_RELATIONSHIP_PARENT":          2,
		"PATIENT_CAREGIVER_RELATIONSHIP_GRANDPARENT":     3,
		"PATIENT_CAREGIVER_RELATIONSHIP_GRANDCHILD":      4,
		"PATIENT_CAREGIVER_RELATIONSHIP_EXTENDED_FAMILY": 5,
		"PATIENT_CAREGIVER_RELATIONSHIP_FRIEND":          6,
		"PATIENT_CAREGIVER_RELATIONSHIP_HEALTH_AIDE":     7,
		"PATIENT_CAREGIVER_RELATIONSHIP_OTHER":           8,
	}
)

func (x PatientCaregiverRelationship) Enum() *PatientCaregiverRelationship {
	p := new(PatientCaregiverRelationship)
	*p = x
	return p
}

func (x PatientCaregiverRelationship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PatientCaregiverRelationship) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes[2].Descriptor()
}

func (PatientCaregiverRelationship) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes[2]
}

func (x PatientCaregiverRelationship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PatientCaregiverRelationship.Descriptor instead.
func (PatientCaregiverRelationship) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_caregiver_proto_rawDescGZIP(), []int{2}
}

// A PatientCaregiver represents a relationship between a Patient and a Caregiver. The relationship indicates that the
// Caregiver is responsible in some way for managing the healthcare of the Patient.
type PatientCaregiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Patient being cared for by the Caregiver. Required.
	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	// The Caregiver that is caring for the Patient. Required.
	//
	// Types that are assignable to Caregiver:
	//	*PatientCaregiver_CaregiverId
	//	*PatientCaregiver_CaregiverMessage
	Caregiver isPatientCaregiver_Caregiver `protobuf_oneof:"caregiver"`
	// The name that the Patient knows the Caregiver by.
	PreferredName string `protobuf:"bytes,4,opt,name=preferred_name,json=preferredName,proto3" json:"preferred_name,omitempty"`
	// The type of PatientCaregiver.
	PatientCaregiverType PatientCaregiverType `protobuf:"varint,5,opt,name=patient_caregiver_type,json=patientCaregiverType,proto3,enum=heyrenee.v1.messages.PatientCaregiverType" json:"patient_caregiver_type,omitempty"`
	// The access that the Caregiver has to the Patient. Required.
	PatientCaregiverAccess PatientCaregiverAccess `protobuf:"varint,6,opt,name=patient_caregiver_access,json=patientCaregiverAccess,proto3,enum=heyrenee.v1.messages.PatientCaregiverAccess" json:"patient_caregiver_access,omitempty"`
	// The relationship of the Caregiver to the Patient. Required.
	PatientCaregiverRelationship PatientCaregiverRelationship `protobuf:"varint,7,opt,name=patient_caregiver_relationship,json=patientCaregiverRelationship,proto3,enum=heyrenee.v1.messages.PatientCaregiverRelationship" json:"patient_caregiver_relationship,omitempty"`
}

func (x *PatientCaregiver) Reset() {
	*x = PatientCaregiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_messages_patient_caregiver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientCaregiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientCaregiver) ProtoMessage() {}

func (x *PatientCaregiver) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_messages_patient_caregiver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientCaregiver.ProtoReflect.Descriptor instead.
func (*PatientCaregiver) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_caregiver_proto_rawDescGZIP(), []int{0}
}

func (x *PatientCaregiver) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (m *PatientCaregiver) GetCaregiver() isPatientCaregiver_Caregiver {
	if m != nil {
		return m.Caregiver
	}
	return nil
}

func (x *PatientCaregiver) GetCaregiverId() string {
	if x, ok := x.GetCaregiver().(*PatientCaregiver_CaregiverId); ok {
		return x.CaregiverId
	}
	return ""
}

func (x *PatientCaregiver) GetCaregiverMessage() *Caregiver {
	if x, ok := x.GetCaregiver().(*PatientCaregiver_CaregiverMessage); ok {
		return x.CaregiverMessage
	}
	return nil
}

func (x *PatientCaregiver) GetPreferredName() string {
	if x != nil {
		return x.PreferredName
	}
	return ""
}

func (x *PatientCaregiver) GetPatientCaregiverType() PatientCaregiverType {
	if x != nil {
		return x.PatientCaregiverType
	}
	return PatientCaregiverType_PATIENT_CAREGIVER_TYPE_UNSPECIFIED
}

func (x *PatientCaregiver) GetPatientCaregiverAccess() PatientCaregiverAccess {
	if x != nil {
		return x.PatientCaregiverAccess
	}
	return PatientCaregiverAccess_PATIENT_CAREGIVER_ACCESS_UNSPECIFIED
}

func (x *PatientCaregiver) GetPatientCaregiverRelationship() PatientCaregiverRelationship {
	if x != nil {
		return x.PatientCaregiverRelationship
	}
	return PatientCaregiverRelationship_PATIENT_CAREGIVER_RELATIONSHIP_UNSPECIFIED
}

type isPatientCaregiver_Caregiver interface {
	isPatientCaregiver_Caregiver()
}

type PatientCaregiver_CaregiverId struct {
	// The ID of the Caregiver that is caring for the Patient. Required.
	CaregiverId string `protobuf:"bytes,2,opt,name=caregiver_id,json=caregiverId,proto3,oneof"`
}

type PatientCaregiver_CaregiverMessage struct {
	// The Caregiver that is caring for the Patient. Only returned in response, must not be set in requests.
	CaregiverMessage *Caregiver `protobuf:"bytes,3,opt,name=caregiver_message,json=caregiverMessage,proto3,oneof"`
}

func (*PatientCaregiver_CaregiverId) isPatientCaregiver_Caregiver() {}

func (*PatientCaregiver_CaregiverMessage) isPatientCaregiver_Caregiver() {}

var File_heyrenee_v1_messages_patient_caregiver_proto protoreflect.FileDescriptor

var file_heyrenee_v1_messages_patient_caregiver_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x1a, 0x24, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x67,
	0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x04, 0x0a, 0x10, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4e, 0x0a, 0x11, 0x63, 0x61, 0x72,
	0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x48, 0x00, 0x52, 0x10, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x60, 0x0a, 0x16, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2a, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x66, 0x0a, 0x18, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x16, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67,
	0x69, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x78, 0x0a, 0x1e, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x1c, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x2a, 0x62, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x41, 0x54,
	0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52,
	0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x4d,
	0x41, 0x52, 0x59, 0x10, 0x01, 0x2a, 0x8d, 0x01, 0x0a, 0x16, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x28, 0x0a, 0x24, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45,
	0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x41,
	0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x26, 0x0a,
	0x22, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56,
	0x45, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x41, 0x53, 0x48, 0x42, 0x4f,
	0x41, 0x52, 0x44, 0x10, 0x02, 0x2a, 0xbb, 0x03, 0x0a, 0x1c, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12, 0x2e, 0x0a, 0x2a, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e,
	0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e,
	0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x10, 0x01,
	0x12, 0x29, 0x0a, 0x25, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45,
	0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48,
	0x49, 0x50, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a, 0x50,
	0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x47, 0x52,
	0x41, 0x4e, 0x44, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x2d, 0x0a, 0x29, 0x50,
	0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x47, 0x52,
	0x41, 0x4e, 0x44, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x10, 0x04, 0x12, 0x32, 0x0a, 0x2e, 0x50, 0x41,
	0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f,
	0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x45, 0x58, 0x54,
	0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x10, 0x05, 0x12, 0x29,
	0x0a, 0x25, 0x50, 0x41, 0x54, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49,
	0x56, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50,
	0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x2e, 0x0a, 0x2a, 0x50, 0x41, 0x54,
	0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x48, 0x45, 0x41, 0x4c,
	0x54, 0x48, 0x5f, 0x41, 0x49, 0x44, 0x45, 0x10, 0x07, 0x12, 0x28, 0x0a, 0x24, 0x50, 0x41, 0x54,
	0x49, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x41, 0x52, 0x45, 0x47, 0x49, 0x56, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x4f, 0x54, 0x48, 0x45,
	0x52, 0x10, 0x08, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_messages_patient_caregiver_proto_rawDescOnce sync.Once
	file_heyrenee_v1_messages_patient_caregiver_proto_rawDescData = file_heyrenee_v1_messages_patient_caregiver_proto_rawDesc
)

func file_heyrenee_v1_messages_patient_caregiver_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_messages_patient_caregiver_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_messages_patient_caregiver_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_messages_patient_caregiver_proto_rawDescData)
	})
	return file_heyrenee_v1_messages_patient_caregiver_proto_rawDescData
}

var file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_heyrenee_v1_messages_patient_caregiver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heyrenee_v1_messages_patient_caregiver_proto_goTypes = []interface{}{
	(PatientCaregiverType)(0),         // 0: heyrenee.v1.messages.PatientCaregiverType
	(PatientCaregiverAccess)(0),       // 1: heyrenee.v1.messages.PatientCaregiverAccess
	(PatientCaregiverRelationship)(0), // 2: heyrenee.v1.messages.PatientCaregiverRelationship
	(*PatientCaregiver)(nil),          // 3: heyrenee.v1.messages.PatientCaregiver
	(*Caregiver)(nil),                 // 4: heyrenee.v1.messages.Caregiver
}
var file_heyrenee_v1_messages_patient_caregiver_proto_depIdxs = []int32{
	4, // 0: heyrenee.v1.messages.PatientCaregiver.caregiver_message:type_name -> heyrenee.v1.messages.Caregiver
	0, // 1: heyrenee.v1.messages.PatientCaregiver.patient_caregiver_type:type_name -> heyrenee.v1.messages.PatientCaregiverType
	1, // 2: heyrenee.v1.messages.PatientCaregiver.patient_caregiver_access:type_name -> heyrenee.v1.messages.PatientCaregiverAccess
	2, // 3: heyrenee.v1.messages.PatientCaregiver.patient_caregiver_relationship:type_name -> heyrenee.v1.messages.PatientCaregiverRelationship
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_messages_patient_caregiver_proto_init() }
func file_heyrenee_v1_messages_patient_caregiver_proto_init() {
	if File_heyrenee_v1_messages_patient_caregiver_proto != nil {
		return
	}
	file_heyrenee_v1_messages_caregiver_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_messages_patient_caregiver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientCaregiver); i {
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
	file_heyrenee_v1_messages_patient_caregiver_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PatientCaregiver_CaregiverId)(nil),
		(*PatientCaregiver_CaregiverMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_heyrenee_v1_messages_patient_caregiver_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_messages_patient_caregiver_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_messages_patient_caregiver_proto_depIdxs,
		EnumInfos:         file_heyrenee_v1_messages_patient_caregiver_proto_enumTypes,
		MessageInfos:      file_heyrenee_v1_messages_patient_caregiver_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_messages_patient_caregiver_proto = out.File
	file_heyrenee_v1_messages_patient_caregiver_proto_rawDesc = nil
	file_heyrenee_v1_messages_patient_caregiver_proto_goTypes = nil
	file_heyrenee_v1_messages_patient_caregiver_proto_depIdxs = nil
}

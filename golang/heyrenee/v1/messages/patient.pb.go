// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/messages/patient.proto

package messages

import (
	_ "github.com/HeyReneeInc/proto/golang/heyrenee/v1/options"
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

type Sex int32

const (
	Sex_SEX_UNSPECIFIED Sex = 0
	Sex_SEX_MALE        Sex = 1
	Sex_SEX_FEMALE      Sex = 2
)

// Enum value maps for Sex.
var (
	Sex_name = map[int32]string{
		0: "SEX_UNSPECIFIED",
		1: "SEX_MALE",
		2: "SEX_FEMALE",
	}
	Sex_value = map[string]int32{
		"SEX_UNSPECIFIED": 0,
		"SEX_MALE":        1,
		"SEX_FEMALE":      2,
	}
)

func (x Sex) Enum() *Sex {
	p := new(Sex)
	*p = x
	return p
}

func (x Sex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Sex) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_proto_enumTypes[0].Descriptor()
}

func (Sex) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_proto_enumTypes[0]
}

func (x Sex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Sex.Descriptor instead.
func (Sex) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_proto_rawDescGZIP(), []int{0}
}

type MaritalStatus int32

const (
	MaritalStatus_MARITAL_STATUS_UNSPECIFIED MaritalStatus = 0
	MaritalStatus_MARITAL_STATUS_SINGLE      MaritalStatus = 1
	MaritalStatus_MARITAL_STATUS_MARRIED     MaritalStatus = 2
)

// Enum value maps for MaritalStatus.
var (
	MaritalStatus_name = map[int32]string{
		0: "MARITAL_STATUS_UNSPECIFIED",
		1: "MARITAL_STATUS_SINGLE",
		2: "MARITAL_STATUS_MARRIED",
	}
	MaritalStatus_value = map[string]int32{
		"MARITAL_STATUS_UNSPECIFIED": 0,
		"MARITAL_STATUS_SINGLE":      1,
		"MARITAL_STATUS_MARRIED":     2,
	}
)

func (x MaritalStatus) Enum() *MaritalStatus {
	p := new(MaritalStatus)
	*p = x
	return p
}

func (x MaritalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MaritalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_proto_enumTypes[1].Descriptor()
}

func (MaritalStatus) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_proto_enumTypes[1]
}

func (x MaritalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MaritalStatus.Descriptor instead.
func (MaritalStatus) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_proto_rawDescGZIP(), []int{1}
}

type Language int32

const (
	Language_LANGUAGE_UNSPECIFIED Language = 0
	Language_LANGUAGE_ENGLISH     Language = 1
	Language_LANGUAGE_SPANISH     Language = 2
	Language_LANGUAGE_OTHER       Language = 3
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "LANGUAGE_UNSPECIFIED",
		1: "LANGUAGE_ENGLISH",
		2: "LANGUAGE_SPANISH",
		3: "LANGUAGE_OTHER",
	}
	Language_value = map[string]int32{
		"LANGUAGE_UNSPECIFIED": 0,
		"LANGUAGE_ENGLISH":     1,
		"LANGUAGE_SPANISH":     2,
		"LANGUAGE_OTHER":       3,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_proto_enumTypes[2].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_proto_enumTypes[2]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_proto_rawDescGZIP(), []int{2}
}

type Ethnicity int32

const (
	Ethnicity_ETHNICITY_UNSPECIFIED                         Ethnicity = 0
	Ethnicity_ETHNICITY_ASIAN_AMERICAN                      Ethnicity = 1
	Ethnicity_ETHNICITY_BLACK_OR_AFRICAN_AMERICAN           Ethnicity = 2
	Ethnicity_ETHNICITY_WHITE_OR_EUROPEAN_AMERICAN          Ethnicity = 3
	Ethnicity_ETHNICITY_AMERICAN_INDIAN_OR_ALASKA_NATIVE    Ethnicity = 4
	Ethnicity_ETHNICITY_NATIVE_HAWAIIAN_OR_PACIFIC_ISLANDER Ethnicity = 5
)

// Enum value maps for Ethnicity.
var (
	Ethnicity_name = map[int32]string{
		0: "ETHNICITY_UNSPECIFIED",
		1: "ETHNICITY_ASIAN_AMERICAN",
		2: "ETHNICITY_BLACK_OR_AFRICAN_AMERICAN",
		3: "ETHNICITY_WHITE_OR_EUROPEAN_AMERICAN",
		4: "ETHNICITY_AMERICAN_INDIAN_OR_ALASKA_NATIVE",
		5: "ETHNICITY_NATIVE_HAWAIIAN_OR_PACIFIC_ISLANDER",
	}
	Ethnicity_value = map[string]int32{
		"ETHNICITY_UNSPECIFIED":                         0,
		"ETHNICITY_ASIAN_AMERICAN":                      1,
		"ETHNICITY_BLACK_OR_AFRICAN_AMERICAN":           2,
		"ETHNICITY_WHITE_OR_EUROPEAN_AMERICAN":          3,
		"ETHNICITY_AMERICAN_INDIAN_OR_ALASKA_NATIVE":    4,
		"ETHNICITY_NATIVE_HAWAIIAN_OR_PACIFIC_ISLANDER": 5,
	}
)

func (x Ethnicity) Enum() *Ethnicity {
	p := new(Ethnicity)
	*p = x
	return p
}

func (x Ethnicity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ethnicity) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_patient_proto_enumTypes[3].Descriptor()
}

func (Ethnicity) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_patient_proto_enumTypes[3]
}

func (x Ethnicity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ethnicity.Descriptor instead.
func (Ethnicity) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_proto_rawDescGZIP(), []int{3}
}

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId         string                 `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	FirstName         string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName        string                 `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName          string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PrimaryPhone      string                 `protobuf:"bytes,5,opt,name=primary_phone,json=primaryPhone,proto3" json:"primary_phone,omitempty"`
	DateOfBirth       *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Sex               Sex                    `protobuf:"varint,7,opt,name=sex,proto3,enum=heyrenee.v1.messages.Sex" json:"sex,omitempty"`
	Email             string                 `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	AddressLines      []string               `protobuf:"bytes,9,rep,name=address_lines,json=addressLines,proto3" json:"address_lines,omitempty"`
	City              string                 `protobuf:"bytes,10,opt,name=city,proto3" json:"city,omitempty"`
	State             string                 `protobuf:"bytes,11,opt,name=state,proto3" json:"state,omitempty"`
	Zip               string                 `protobuf:"bytes,12,opt,name=zip,proto3" json:"zip,omitempty"`
	PreferredName     string                 `protobuf:"bytes,13,opt,name=preferred_name,json=preferredName,proto3" json:"preferred_name,omitempty"`
	MaritalStatus     MaritalStatus          `protobuf:"varint,14,opt,name=marital_status,json=maritalStatus,proto3,enum=heyrenee.v1.messages.MaritalStatus" json:"marital_status,omitempty"`
	PreferredLanguage Language               `protobuf:"varint,15,opt,name=preferred_language,json=preferredLanguage,proto3,enum=heyrenee.v1.messages.Language" json:"preferred_language,omitempty"`
	Ethnicity         Ethnicity              `protobuf:"varint,16,opt,name=ethnicity,proto3,enum=heyrenee.v1.messages.Ethnicity" json:"ethnicity,omitempty"`
	MobilePhone       string                 `protobuf:"bytes,17,opt,name=mobile_phone,json=mobilePhone,proto3" json:"mobile_phone,omitempty"`
	OtherPhone        string                 `protobuf:"bytes,18,opt,name=other_phone,json=otherPhone,proto3" json:"other_phone,omitempty"`
	Notes             string                 `protobuf:"bytes,19,opt,name=notes,proto3" json:"notes,omitempty"`
	NamePronunciation string                 `protobuf:"bytes,20,opt,name=name_pronunciation,json=namePronunciation,proto3" json:"name_pronunciation,omitempty"`
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_messages_patient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_messages_patient_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_patient_proto_rawDescGZIP(), []int{0}
}

func (x *Patient) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *Patient) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Patient) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *Patient) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Patient) GetPrimaryPhone() string {
	if x != nil {
		return x.PrimaryPhone
	}
	return ""
}

func (x *Patient) GetDateOfBirth() *timestamppb.Timestamp {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *Patient) GetSex() Sex {
	if x != nil {
		return x.Sex
	}
	return Sex_SEX_UNSPECIFIED
}

func (x *Patient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Patient) GetAddressLines() []string {
	if x != nil {
		return x.AddressLines
	}
	return nil
}

func (x *Patient) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Patient) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Patient) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

func (x *Patient) GetPreferredName() string {
	if x != nil {
		return x.PreferredName
	}
	return ""
}

func (x *Patient) GetMaritalStatus() MaritalStatus {
	if x != nil {
		return x.MaritalStatus
	}
	return MaritalStatus_MARITAL_STATUS_UNSPECIFIED
}

func (x *Patient) GetPreferredLanguage() Language {
	if x != nil {
		return x.PreferredLanguage
	}
	return Language_LANGUAGE_UNSPECIFIED
}

func (x *Patient) GetEthnicity() Ethnicity {
	if x != nil {
		return x.Ethnicity
	}
	return Ethnicity_ETHNICITY_UNSPECIFIED
}

func (x *Patient) GetMobilePhone() string {
	if x != nil {
		return x.MobilePhone
	}
	return ""
}

func (x *Patient) GetOtherPhone() string {
	if x != nil {
		return x.OtherPhone
	}
	return ""
}

func (x *Patient) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Patient) GetNamePronunciation() string {
	if x != nil {
		return x.NamePronunciation
	}
	return ""
}

var File_heyrenee_v1_messages_patient_proto protoreflect.FileDescriptor

var file_heyrenee_v1_messages_patient_proto_rawDesc = []byte{
	0x0a, 0x22, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x68, 0x65, 0x79,
	0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x06, 0x0a, 0x07,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18,
	0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x2b,
	0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x68, 0x65,
	0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x7a, 0x69, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a,
	0x69, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0e, 0x6d, 0x61, 0x72,
	0x69, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4d, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x64, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x65, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45,
	0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79, 0x52, 0x09, 0x65, 0x74, 0x68, 0x6e, 0x69, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x12, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6e, 0x75, 0x6e, 0x63, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6e, 0x61, 0x6d, 0x65, 0x50,
	0x72, 0x6f, 0x6e, 0x75, 0x6e, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x38, 0x0a, 0x03,
	0x53, 0x65, 0x78, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x58, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x58, 0x5f,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x58, 0x5f, 0x46, 0x45,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x2a, 0x66, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x52, 0x49, 0x54,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x41, 0x52, 0x49, 0x54,
	0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x52, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x41, 0x52, 0x52, 0x49, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x64,
	0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4c, 0x41,
	0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45,
	0x5f, 0x45, 0x4e, 0x47, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41,
	0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x10, 0x03, 0x2a, 0xfa, 0x01, 0x0a, 0x09, 0x45, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a,
	0x18, 0x45, 0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x53, 0x49, 0x41, 0x4e,
	0x5f, 0x41, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x41, 0x4e, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x45,
	0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x5f, 0x4f,
	0x52, 0x5f, 0x41, 0x46, 0x52, 0x49, 0x43, 0x41, 0x4e, 0x5f, 0x41, 0x4d, 0x45, 0x52, 0x49, 0x43,
	0x41, 0x4e, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x45, 0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54,
	0x59, 0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x55, 0x52, 0x4f, 0x50,
	0x45, 0x41, 0x4e, 0x5f, 0x41, 0x4d, 0x45, 0x52, 0x49, 0x43, 0x41, 0x4e, 0x10, 0x03, 0x12, 0x2e,
	0x0a, 0x2a, 0x45, 0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x41, 0x4d, 0x45, 0x52,
	0x49, 0x43, 0x41, 0x4e, 0x5f, 0x49, 0x4e, 0x44, 0x49, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x41,
	0x4c, 0x41, 0x53, 0x4b, 0x41, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x31,
	0x0a, 0x2d, 0x45, 0x54, 0x48, 0x4e, 0x49, 0x43, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x41, 0x54, 0x49,
	0x56, 0x45, 0x5f, 0x48, 0x41, 0x57, 0x41, 0x49, 0x49, 0x41, 0x4e, 0x5f, 0x4f, 0x52, 0x5f, 0x50,
	0x41, 0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x49, 0x53, 0x4c, 0x41, 0x4e, 0x44, 0x45, 0x52, 0x10,
	0x05, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_messages_patient_proto_rawDescOnce sync.Once
	file_heyrenee_v1_messages_patient_proto_rawDescData = file_heyrenee_v1_messages_patient_proto_rawDesc
)

func file_heyrenee_v1_messages_patient_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_messages_patient_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_messages_patient_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_messages_patient_proto_rawDescData)
	})
	return file_heyrenee_v1_messages_patient_proto_rawDescData
}

var file_heyrenee_v1_messages_patient_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_heyrenee_v1_messages_patient_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heyrenee_v1_messages_patient_proto_goTypes = []interface{}{
	(Sex)(0),                      // 0: heyrenee.v1.messages.Sex
	(MaritalStatus)(0),            // 1: heyrenee.v1.messages.MaritalStatus
	(Language)(0),                 // 2: heyrenee.v1.messages.Language
	(Ethnicity)(0),                // 3: heyrenee.v1.messages.Ethnicity
	(*Patient)(nil),               // 4: heyrenee.v1.messages.Patient
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_heyrenee_v1_messages_patient_proto_depIdxs = []int32{
	5, // 0: heyrenee.v1.messages.Patient.date_of_birth:type_name -> google.protobuf.Timestamp
	0, // 1: heyrenee.v1.messages.Patient.sex:type_name -> heyrenee.v1.messages.Sex
	1, // 2: heyrenee.v1.messages.Patient.marital_status:type_name -> heyrenee.v1.messages.MaritalStatus
	2, // 3: heyrenee.v1.messages.Patient.preferred_language:type_name -> heyrenee.v1.messages.Language
	3, // 4: heyrenee.v1.messages.Patient.ethnicity:type_name -> heyrenee.v1.messages.Ethnicity
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_messages_patient_proto_init() }
func file_heyrenee_v1_messages_patient_proto_init() {
	if File_heyrenee_v1_messages_patient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_messages_patient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
			RawDescriptor: file_heyrenee_v1_messages_patient_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_messages_patient_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_messages_patient_proto_depIdxs,
		EnumInfos:         file_heyrenee_v1_messages_patient_proto_enumTypes,
		MessageInfos:      file_heyrenee_v1_messages_patient_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_messages_patient_proto = out.File
	file_heyrenee_v1_messages_patient_proto_rawDesc = nil
	file_heyrenee_v1_messages_patient_proto_goTypes = nil
	file_heyrenee_v1_messages_patient_proto_depIdxs = nil
}

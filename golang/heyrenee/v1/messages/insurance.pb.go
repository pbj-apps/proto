// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/messages/insurance.proto

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

type InsuranceType int32

const (
	InsuranceType_INSURANCE_TYPE_UNSPECIFIED        InsuranceType = 0
	InsuranceType_INSURANCE_TYPE_MEDICAID           InsuranceType = 1
	InsuranceType_INSURANCE_TYPE_MEDICARE           InsuranceType = 2
	InsuranceType_INSURANCE_TYPE_MEDICARE_ADVANTAGE InsuranceType = 3
	InsuranceType_INSURANCE_TYPE_VETERANS_AFFAIRS   InsuranceType = 4
	InsuranceType_INSURANCE_TYPE_EMPLOYER_BASED     InsuranceType = 5
	InsuranceType_INSURANCE_TYPE_PRIVATE            InsuranceType = 6
)

// Enum value maps for InsuranceType.
var (
	InsuranceType_name = map[int32]string{
		0: "INSURANCE_TYPE_UNSPECIFIED",
		1: "INSURANCE_TYPE_MEDICAID",
		2: "INSURANCE_TYPE_MEDICARE",
		3: "INSURANCE_TYPE_MEDICARE_ADVANTAGE",
		4: "INSURANCE_TYPE_VETERANS_AFFAIRS",
		5: "INSURANCE_TYPE_EMPLOYER_BASED",
		6: "INSURANCE_TYPE_PRIVATE",
	}
	InsuranceType_value = map[string]int32{
		"INSURANCE_TYPE_UNSPECIFIED":        0,
		"INSURANCE_TYPE_MEDICAID":           1,
		"INSURANCE_TYPE_MEDICARE":           2,
		"INSURANCE_TYPE_MEDICARE_ADVANTAGE": 3,
		"INSURANCE_TYPE_VETERANS_AFFAIRS":   4,
		"INSURANCE_TYPE_EMPLOYER_BASED":     5,
		"INSURANCE_TYPE_PRIVATE":            6,
	}
)

func (x InsuranceType) Enum() *InsuranceType {
	p := new(InsuranceType)
	*p = x
	return p
}

func (x InsuranceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsuranceType) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_insurance_proto_enumTypes[0].Descriptor()
}

func (InsuranceType) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_insurance_proto_enumTypes[0]
}

func (x InsuranceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsuranceType.Descriptor instead.
func (InsuranceType) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_insurance_proto_rawDescGZIP(), []int{0}
}

type InsuranceStatus int32

const (
	InsuranceStatus_INSURANCE_STATUS_UNSPECIFIED InsuranceStatus = 0
	InsuranceStatus_INSURANCE_STATUS_ACTIVE      InsuranceStatus = 1
	InsuranceStatus_INSURANCE_STATUS_INACTIVE    InsuranceStatus = 2
)

// Enum value maps for InsuranceStatus.
var (
	InsuranceStatus_name = map[int32]string{
		0: "INSURANCE_STATUS_UNSPECIFIED",
		1: "INSURANCE_STATUS_ACTIVE",
		2: "INSURANCE_STATUS_INACTIVE",
	}
	InsuranceStatus_value = map[string]int32{
		"INSURANCE_STATUS_UNSPECIFIED": 0,
		"INSURANCE_STATUS_ACTIVE":      1,
		"INSURANCE_STATUS_INACTIVE":    2,
	}
)

func (x InsuranceStatus) Enum() *InsuranceStatus {
	p := new(InsuranceStatus)
	*p = x
	return p
}

func (x InsuranceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsuranceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_heyrenee_v1_messages_insurance_proto_enumTypes[1].Descriptor()
}

func (InsuranceStatus) Type() protoreflect.EnumType {
	return &file_heyrenee_v1_messages_insurance_proto_enumTypes[1]
}

func (x InsuranceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsuranceStatus.Descriptor instead.
func (InsuranceStatus) EnumDescriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_insurance_proto_rawDescGZIP(), []int{1}
}

type Insurance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId               string          `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	InsuranceId             string          `protobuf:"bytes,2,opt,name=insurance_id,json=insuranceId,proto3" json:"insurance_id,omitempty"`
	Insurer                 string          `protobuf:"bytes,3,opt,name=insurer,proto3" json:"insurer,omitempty"`
	InsuranceType           InsuranceType   `protobuf:"varint,4,opt,name=insurance_type,json=insuranceType,proto3,enum=heyrenee.v1.messages.InsuranceType" json:"insurance_type,omitempty"`
	InsuranceStatus         InsuranceStatus `protobuf:"varint,5,opt,name=insurance_status,json=insuranceStatus,proto3,enum=heyrenee.v1.messages.InsuranceStatus" json:"insurance_status,omitempty"`
	PolicyOwnerName         string          `protobuf:"bytes,6,opt,name=policy_owner_name,json=policyOwnerName,proto3" json:"policy_owner_name,omitempty"`
	PolicyOwnerAddressLines []string        `protobuf:"bytes,7,rep,name=policy_owner_address_lines,json=policyOwnerAddressLines,proto3" json:"policy_owner_address_lines,omitempty"`
	PolicyOwnerCity         string          `protobuf:"bytes,8,opt,name=policy_owner_city,json=policyOwnerCity,proto3" json:"policy_owner_city,omitempty"`
	PolicyOwnerState        string          `protobuf:"bytes,9,opt,name=policy_owner_state,json=policyOwnerState,proto3" json:"policy_owner_state,omitempty"`
	PolicyOwnerZip          string          `protobuf:"bytes,10,opt,name=policy_owner_zip,json=policyOwnerZip,proto3" json:"policy_owner_zip,omitempty"`
	PolicyOwnerPhone        string          `protobuf:"bytes,11,opt,name=policy_owner_phone,json=policyOwnerPhone,proto3" json:"policy_owner_phone,omitempty"`
	GroupNumber             string          `protobuf:"bytes,12,opt,name=group_number,json=groupNumber,proto3" json:"group_number,omitempty"`
	PolicyNumber            string          `protobuf:"bytes,13,opt,name=policy_number,json=policyNumber,proto3" json:"policy_number,omitempty"`
	RxPolicyNumber          string          `protobuf:"bytes,14,opt,name=rx_policy_number,json=rxPolicyNumber,proto3" json:"rx_policy_number,omitempty"`
	ClaimsAddressLines      []string        `protobuf:"bytes,15,rep,name=claims_address_lines,json=claimsAddressLines,proto3" json:"claims_address_lines,omitempty"`
	ClaimsCity              string          `protobuf:"bytes,16,opt,name=claims_city,json=claimsCity,proto3" json:"claims_city,omitempty"`
	ClaimsState             string          `protobuf:"bytes,17,opt,name=claims_state,json=claimsState,proto3" json:"claims_state,omitempty"`
	ClaimsZip               string          `protobuf:"bytes,18,opt,name=claims_zip,json=claimsZip,proto3" json:"claims_zip,omitempty"`
	ClaimsPhone             string          `protobuf:"bytes,19,opt,name=claims_phone,json=claimsPhone,proto3" json:"claims_phone,omitempty"`
	RxClaimsAddressLines    []string        `protobuf:"bytes,29,rep,name=rx_claims_address_lines,json=rxClaimsAddressLines,proto3" json:"rx_claims_address_lines,omitempty"`
	RxClaimsCity            string          `protobuf:"bytes,21,opt,name=rx_claims_city,json=rxClaimsCity,proto3" json:"rx_claims_city,omitempty"`
	RxClaimsState           string          `protobuf:"bytes,22,opt,name=rx_claims_state,json=rxClaimsState,proto3" json:"rx_claims_state,omitempty"`
	RxClaimsZip             string          `protobuf:"bytes,23,opt,name=rx_claims_zip,json=rxClaimsZip,proto3" json:"rx_claims_zip,omitempty"`
	RxClaimsPhone           string          `protobuf:"bytes,24,opt,name=rx_claims_phone,json=rxClaimsPhone,proto3" json:"rx_claims_phone,omitempty"`
}

func (x *Insurance) Reset() {
	*x = Insurance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_messages_insurance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Insurance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Insurance) ProtoMessage() {}

func (x *Insurance) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_messages_insurance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Insurance.ProtoReflect.Descriptor instead.
func (*Insurance) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_insurance_proto_rawDescGZIP(), []int{0}
}

func (x *Insurance) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *Insurance) GetInsuranceId() string {
	if x != nil {
		return x.InsuranceId
	}
	return ""
}

func (x *Insurance) GetInsurer() string {
	if x != nil {
		return x.Insurer
	}
	return ""
}

func (x *Insurance) GetInsuranceType() InsuranceType {
	if x != nil {
		return x.InsuranceType
	}
	return InsuranceType_INSURANCE_TYPE_UNSPECIFIED
}

func (x *Insurance) GetInsuranceStatus() InsuranceStatus {
	if x != nil {
		return x.InsuranceStatus
	}
	return InsuranceStatus_INSURANCE_STATUS_UNSPECIFIED
}

func (x *Insurance) GetPolicyOwnerName() string {
	if x != nil {
		return x.PolicyOwnerName
	}
	return ""
}

func (x *Insurance) GetPolicyOwnerAddressLines() []string {
	if x != nil {
		return x.PolicyOwnerAddressLines
	}
	return nil
}

func (x *Insurance) GetPolicyOwnerCity() string {
	if x != nil {
		return x.PolicyOwnerCity
	}
	return ""
}

func (x *Insurance) GetPolicyOwnerState() string {
	if x != nil {
		return x.PolicyOwnerState
	}
	return ""
}

func (x *Insurance) GetPolicyOwnerZip() string {
	if x != nil {
		return x.PolicyOwnerZip
	}
	return ""
}

func (x *Insurance) GetPolicyOwnerPhone() string {
	if x != nil {
		return x.PolicyOwnerPhone
	}
	return ""
}

func (x *Insurance) GetGroupNumber() string {
	if x != nil {
		return x.GroupNumber
	}
	return ""
}

func (x *Insurance) GetPolicyNumber() string {
	if x != nil {
		return x.PolicyNumber
	}
	return ""
}

func (x *Insurance) GetRxPolicyNumber() string {
	if x != nil {
		return x.RxPolicyNumber
	}
	return ""
}

func (x *Insurance) GetClaimsAddressLines() []string {
	if x != nil {
		return x.ClaimsAddressLines
	}
	return nil
}

func (x *Insurance) GetClaimsCity() string {
	if x != nil {
		return x.ClaimsCity
	}
	return ""
}

func (x *Insurance) GetClaimsState() string {
	if x != nil {
		return x.ClaimsState
	}
	return ""
}

func (x *Insurance) GetClaimsZip() string {
	if x != nil {
		return x.ClaimsZip
	}
	return ""
}

func (x *Insurance) GetClaimsPhone() string {
	if x != nil {
		return x.ClaimsPhone
	}
	return ""
}

func (x *Insurance) GetRxClaimsAddressLines() []string {
	if x != nil {
		return x.RxClaimsAddressLines
	}
	return nil
}

func (x *Insurance) GetRxClaimsCity() string {
	if x != nil {
		return x.RxClaimsCity
	}
	return ""
}

func (x *Insurance) GetRxClaimsState() string {
	if x != nil {
		return x.RxClaimsState
	}
	return ""
}

func (x *Insurance) GetRxClaimsZip() string {
	if x != nil {
		return x.RxClaimsZip
	}
	return ""
}

func (x *Insurance) GetRxClaimsPhone() string {
	if x != nil {
		return x.RxClaimsPhone
	}
	return ""
}

var File_heyrenee_v1_messages_insurance_proto protoreflect.FileDescriptor

var file_heyrenee_v1_messages_insurance_proto_rawDesc = []byte{
	0x0a, 0x24, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1e, 0x68, 0x65,
	0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x08, 0x0a,
	0x09, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x73, 0x75, 0x72, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0e,
	0x69, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x75,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x69, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x50, 0x0a, 0x10, 0x69, 0x6e, 0x73, 0x75,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x69, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x1a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x6e, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x7a, 0x69,
	0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x5a, 0x69, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a,
	0x10, 0x72, 0x78, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x78, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x43, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x7a, 0x69, 0x70, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5a, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x35, 0x0a, 0x17, 0x72, 0x78, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x14, 0x72, 0x78, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x78, 0x5f, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x5f, 0x63, 0x69, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x78, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x43, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0f,
	0x72, 0x78, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x78, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x78, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x5f, 0x7a, 0x69, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x78, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5a, 0x69, 0x70, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x78, 0x5f, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x78, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x2a, 0xf4, 0x01, 0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x43, 0x41, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x43, 0x41, 0x52, 0x45, 0x10, 0x02, 0x12, 0x25, 0x0a, 0x21,
	0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d,
	0x45, 0x44, 0x49, 0x43, 0x41, 0x52, 0x45, 0x5f, 0x41, 0x44, 0x56, 0x41, 0x4e, 0x54, 0x41, 0x47,
	0x45, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x56, 0x45, 0x54, 0x45, 0x52, 0x41, 0x4e, 0x53, 0x5f, 0x41,
	0x46, 0x46, 0x41, 0x49, 0x52, 0x53, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d, 0x49, 0x4e, 0x53, 0x55,
	0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d, 0x50, 0x4c, 0x4f,
	0x59, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x49,
	0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x06, 0x2a, 0x6f, 0x0a, 0x0f, 0x49, 0x6e, 0x73, 0x75, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x4e,
	0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17,
	0x49, 0x4e, 0x53, 0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x53,
	0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49,
	0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_messages_insurance_proto_rawDescOnce sync.Once
	file_heyrenee_v1_messages_insurance_proto_rawDescData = file_heyrenee_v1_messages_insurance_proto_rawDesc
)

func file_heyrenee_v1_messages_insurance_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_messages_insurance_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_messages_insurance_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_messages_insurance_proto_rawDescData)
	})
	return file_heyrenee_v1_messages_insurance_proto_rawDescData
}

var file_heyrenee_v1_messages_insurance_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_heyrenee_v1_messages_insurance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heyrenee_v1_messages_insurance_proto_goTypes = []interface{}{
	(InsuranceType)(0),   // 0: heyrenee.v1.messages.InsuranceType
	(InsuranceStatus)(0), // 1: heyrenee.v1.messages.InsuranceStatus
	(*Insurance)(nil),    // 2: heyrenee.v1.messages.Insurance
}
var file_heyrenee_v1_messages_insurance_proto_depIdxs = []int32{
	0, // 0: heyrenee.v1.messages.Insurance.insurance_type:type_name -> heyrenee.v1.messages.InsuranceType
	1, // 1: heyrenee.v1.messages.Insurance.insurance_status:type_name -> heyrenee.v1.messages.InsuranceStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_messages_insurance_proto_init() }
func file_heyrenee_v1_messages_insurance_proto_init() {
	if File_heyrenee_v1_messages_insurance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_messages_insurance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Insurance); i {
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
			RawDescriptor: file_heyrenee_v1_messages_insurance_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_messages_insurance_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_messages_insurance_proto_depIdxs,
		EnumInfos:         file_heyrenee_v1_messages_insurance_proto_enumTypes,
		MessageInfos:      file_heyrenee_v1_messages_insurance_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_messages_insurance_proto = out.File
	file_heyrenee_v1_messages_insurance_proto_rawDesc = nil
	file_heyrenee_v1_messages_insurance_proto_goTypes = nil
	file_heyrenee_v1_messages_insurance_proto_depIdxs = nil
}

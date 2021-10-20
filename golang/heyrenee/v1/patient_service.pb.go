// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/patient_service.proto

package v1

import (
	messages "github.com/HeyReneeInc/proto/golang/heyrenee/v1/messages"
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

// Request message for GetPatient.
type GetPatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Patient to retrieve. Required.
	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
}

func (x *GetPatientRequest) Reset() {
	*x = GetPatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientRequest) ProtoMessage() {}

func (x *GetPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPatientRequest.ProtoReflect.Descriptor instead.
func (*GetPatientRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPatientRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

// Request message for UpdatePatient.
type UpdatePatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Patient to update. Required.
	Patient *messages.Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *UpdatePatientRequest) Reset() {
	*x = UpdatePatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientRequest) ProtoMessage() {}

func (x *UpdatePatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatientRequest.ProtoReflect.Descriptor instead.
func (*UpdatePatientRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePatientRequest) GetPatient() *messages.Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

// Request message for CreatePatientProvider.
type CreatePatientProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The PatientProvider to create. Required.
	PatientProvider *messages.PatientProvider `protobuf:"bytes,1,opt,name=patient_provider,json=patientProvider,proto3" json:"patient_provider,omitempty"`
}

func (x *CreatePatientProviderRequest) Reset() {
	*x = CreatePatientProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientProviderRequest) ProtoMessage() {}

func (x *CreatePatientProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientProviderRequest.ProtoReflect.Descriptor instead.
func (*CreatePatientProviderRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePatientProviderRequest) GetPatientProvider() *messages.PatientProvider {
	if x != nil {
		return x.PatientProvider
	}
	return nil
}

// Request message for UpdatePatientProvider.
type UpdatePatientProviderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The PatientProvider to update. Required.
	PatientProvider *messages.PatientProvider `protobuf:"bytes,1,opt,name=patient_provider,json=patientProvider,proto3" json:"patient_provider,omitempty"`
}

func (x *UpdatePatientProviderRequest) Reset() {
	*x = UpdatePatientProviderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientProviderRequest) ProtoMessage() {}

func (x *UpdatePatientProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatientProviderRequest.ProtoReflect.Descriptor instead.
func (*UpdatePatientProviderRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePatientProviderRequest) GetPatientProvider() *messages.PatientProvider {
	if x != nil {
		return x.PatientProvider
	}
	return nil
}

// Request message for ListPatientProviders.
type ListPatientProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Patient to return PatientProviders for. Required.
	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
	// The PatientProvider status. If specified, only PatientProviders with this status will be returned.
	PatientProviderStatus messages.PatientProviderStatus `protobuf:"varint,2,opt,name=patient_provider_status,json=patientProviderStatus,proto3,enum=heyrenee.v1.messages.PatientProviderStatus" json:"patient_provider_status,omitempty"`
	// The PatientProvider type. If specified, only PatientProviders of this type will be returned.
	PatientProviderType messages.PatientProviderType `protobuf:"varint,3,opt,name=patient_provider_type,json=patientProviderType,proto3,enum=heyrenee.v1.messages.PatientProviderType" json:"patient_provider_type,omitempty"`
}

func (x *ListPatientProvidersRequest) Reset() {
	*x = ListPatientProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientProvidersRequest) ProtoMessage() {}

func (x *ListPatientProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPatientProvidersRequest.ProtoReflect.Descriptor instead.
func (*ListPatientProvidersRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListPatientProvidersRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

func (x *ListPatientProvidersRequest) GetPatientProviderStatus() messages.PatientProviderStatus {
	if x != nil {
		return x.PatientProviderStatus
	}
	return messages.PatientProviderStatus(0)
}

func (x *ListPatientProvidersRequest) GetPatientProviderType() messages.PatientProviderType {
	if x != nil {
		return x.PatientProviderType
	}
	return messages.PatientProviderType(0)
}

type ListPatientProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientProviders []*messages.PatientProvider `protobuf:"bytes,1,rep,name=patient_providers,json=patientProviders,proto3" json:"patient_providers,omitempty"`
}

func (x *ListPatientProvidersResponse) Reset() {
	*x = ListPatientProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientProvidersResponse) ProtoMessage() {}

func (x *ListPatientProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPatientProvidersResponse.ProtoReflect.Descriptor instead.
func (*ListPatientProvidersResponse) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListPatientProvidersResponse) GetPatientProviders() []*messages.PatientProvider {
	if x != nil {
		return x.PatientProviders
	}
	return nil
}

type CreatePatientCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientCaregiver *messages.PatientCaregiver `protobuf:"bytes,1,opt,name=patient_caregiver,json=patientCaregiver,proto3" json:"patient_caregiver,omitempty"`
}

func (x *CreatePatientCaregiverRequest) Reset() {
	*x = CreatePatientCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientCaregiverRequest) ProtoMessage() {}

func (x *CreatePatientCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientCaregiverRequest.ProtoReflect.Descriptor instead.
func (*CreatePatientCaregiverRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePatientCaregiverRequest) GetPatientCaregiver() *messages.PatientCaregiver {
	if x != nil {
		return x.PatientCaregiver
	}
	return nil
}

type UpdatePatientCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientCaregiver *messages.PatientCaregiver `protobuf:"bytes,1,opt,name=patient_caregiver,json=patientCaregiver,proto3" json:"patient_caregiver,omitempty"`
}

func (x *UpdatePatientCaregiverRequest) Reset() {
	*x = UpdatePatientCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientCaregiverRequest) ProtoMessage() {}

func (x *UpdatePatientCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePatientCaregiverRequest.ProtoReflect.Descriptor instead.
func (*UpdatePatientCaregiverRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePatientCaregiverRequest) GetPatientCaregiver() *messages.PatientCaregiver {
	if x != nil {
		return x.PatientCaregiver
	}
	return nil
}

type ListPatientCaregiversRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientId string `protobuf:"bytes,1,opt,name=patient_id,json=patientId,proto3" json:"patient_id,omitempty"`
}

func (x *ListPatientCaregiversRequest) Reset() {
	*x = ListPatientCaregiversRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientCaregiversRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientCaregiversRequest) ProtoMessage() {}

func (x *ListPatientCaregiversRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPatientCaregiversRequest.ProtoReflect.Descriptor instead.
func (*ListPatientCaregiversRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListPatientCaregiversRequest) GetPatientId() string {
	if x != nil {
		return x.PatientId
	}
	return ""
}

type ListPatientCaregiversResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PatientCaregivers []*messages.PatientCaregiver `protobuf:"bytes,1,rep,name=patient_caregivers,json=patientCaregivers,proto3" json:"patient_caregivers,omitempty"`
}

func (x *ListPatientCaregiversResponse) Reset() {
	*x = ListPatientCaregiversResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_patient_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientCaregiversResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientCaregiversResponse) ProtoMessage() {}

func (x *ListPatientCaregiversResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_patient_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPatientCaregiversResponse.ProtoReflect.Descriptor instead.
func (*ListPatientCaregiversResponse) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_patient_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListPatientCaregiversResponse) GetPatientCaregivers() []*messages.PatientCaregiver {
	if x != nil {
		return x.PatientCaregivers
	}
	return nil
}

var File_heyrenee_v1_patient_service_proto protoreflect.FileDescriptor

var file_heyrenee_v1_patient_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x22, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x38, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x1c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x10, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x1c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x10,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x86,
	0x02, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x63, 0x0a, 0x17, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x15, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5d, 0x0a, 0x15, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x13, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x72, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x10, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x74, 0x0a, 0x1d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x11,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52,
	0x10, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x22, 0x74, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x53, 0x0a, 0x11, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x10, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x11, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x73, 0x32, 0xef,
	0x06, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x1e, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x04,
	0x80, 0xb5, 0x18, 0x01, 0x12, 0x57, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x12, 0x6f, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x12, 0x6f,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x12,
	0x71, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x80, 0xb5,
	0x18, 0x01, 0x12, 0x72, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72,
	0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x12, 0x72, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72,
	0x12, 0x2a, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67,
	0x69, 0x76, 0x65, 0x72, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x12, 0x74, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x12, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72,
	0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48,
	0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_patient_service_proto_rawDescOnce sync.Once
	file_heyrenee_v1_patient_service_proto_rawDescData = file_heyrenee_v1_patient_service_proto_rawDesc
)

func file_heyrenee_v1_patient_service_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_patient_service_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_patient_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_patient_service_proto_rawDescData)
	})
	return file_heyrenee_v1_patient_service_proto_rawDescData
}

var file_heyrenee_v1_patient_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_heyrenee_v1_patient_service_proto_goTypes = []interface{}{
	(*GetPatientRequest)(nil),             // 0: heyrenee.v1.GetPatientRequest
	(*UpdatePatientRequest)(nil),          // 1: heyrenee.v1.UpdatePatientRequest
	(*CreatePatientProviderRequest)(nil),  // 2: heyrenee.v1.CreatePatientProviderRequest
	(*UpdatePatientProviderRequest)(nil),  // 3: heyrenee.v1.UpdatePatientProviderRequest
	(*ListPatientProvidersRequest)(nil),   // 4: heyrenee.v1.ListPatientProvidersRequest
	(*ListPatientProvidersResponse)(nil),  // 5: heyrenee.v1.ListPatientProvidersResponse
	(*CreatePatientCaregiverRequest)(nil), // 6: heyrenee.v1.CreatePatientCaregiverRequest
	(*UpdatePatientCaregiverRequest)(nil), // 7: heyrenee.v1.UpdatePatientCaregiverRequest
	(*ListPatientCaregiversRequest)(nil),  // 8: heyrenee.v1.ListPatientCaregiversRequest
	(*ListPatientCaregiversResponse)(nil), // 9: heyrenee.v1.ListPatientCaregiversResponse
	(*messages.Patient)(nil),              // 10: heyrenee.v1.messages.Patient
	(*messages.PatientProvider)(nil),      // 11: heyrenee.v1.messages.PatientProvider
	(messages.PatientProviderStatus)(0),   // 12: heyrenee.v1.messages.PatientProviderStatus
	(messages.PatientProviderType)(0),     // 13: heyrenee.v1.messages.PatientProviderType
	(*messages.PatientCaregiver)(nil),     // 14: heyrenee.v1.messages.PatientCaregiver
}
var file_heyrenee_v1_patient_service_proto_depIdxs = []int32{
	10, // 0: heyrenee.v1.UpdatePatientRequest.patient:type_name -> heyrenee.v1.messages.Patient
	11, // 1: heyrenee.v1.CreatePatientProviderRequest.patient_provider:type_name -> heyrenee.v1.messages.PatientProvider
	11, // 2: heyrenee.v1.UpdatePatientProviderRequest.patient_provider:type_name -> heyrenee.v1.messages.PatientProvider
	12, // 3: heyrenee.v1.ListPatientProvidersRequest.patient_provider_status:type_name -> heyrenee.v1.messages.PatientProviderStatus
	13, // 4: heyrenee.v1.ListPatientProvidersRequest.patient_provider_type:type_name -> heyrenee.v1.messages.PatientProviderType
	11, // 5: heyrenee.v1.ListPatientProvidersResponse.patient_providers:type_name -> heyrenee.v1.messages.PatientProvider
	14, // 6: heyrenee.v1.CreatePatientCaregiverRequest.patient_caregiver:type_name -> heyrenee.v1.messages.PatientCaregiver
	14, // 7: heyrenee.v1.UpdatePatientCaregiverRequest.patient_caregiver:type_name -> heyrenee.v1.messages.PatientCaregiver
	14, // 8: heyrenee.v1.ListPatientCaregiversResponse.patient_caregivers:type_name -> heyrenee.v1.messages.PatientCaregiver
	0,  // 9: heyrenee.v1.PatientService.GetPatient:input_type -> heyrenee.v1.GetPatientRequest
	1,  // 10: heyrenee.v1.PatientService.UpdatePatient:input_type -> heyrenee.v1.UpdatePatientRequest
	2,  // 11: heyrenee.v1.PatientService.CreatePatientProvider:input_type -> heyrenee.v1.CreatePatientProviderRequest
	3,  // 12: heyrenee.v1.PatientService.UpdatePatientProvider:input_type -> heyrenee.v1.UpdatePatientProviderRequest
	4,  // 13: heyrenee.v1.PatientService.ListPatientProviders:input_type -> heyrenee.v1.ListPatientProvidersRequest
	6,  // 14: heyrenee.v1.PatientService.CreatePatientCaregiver:input_type -> heyrenee.v1.CreatePatientCaregiverRequest
	7,  // 15: heyrenee.v1.PatientService.UpdatePatientCaregiver:input_type -> heyrenee.v1.UpdatePatientCaregiverRequest
	8,  // 16: heyrenee.v1.PatientService.ListPatientCaregivers:input_type -> heyrenee.v1.ListPatientCaregiversRequest
	10, // 17: heyrenee.v1.PatientService.GetPatient:output_type -> heyrenee.v1.messages.Patient
	10, // 18: heyrenee.v1.PatientService.UpdatePatient:output_type -> heyrenee.v1.messages.Patient
	11, // 19: heyrenee.v1.PatientService.CreatePatientProvider:output_type -> heyrenee.v1.messages.PatientProvider
	11, // 20: heyrenee.v1.PatientService.UpdatePatientProvider:output_type -> heyrenee.v1.messages.PatientProvider
	5,  // 21: heyrenee.v1.PatientService.ListPatientProviders:output_type -> heyrenee.v1.ListPatientProvidersResponse
	14, // 22: heyrenee.v1.PatientService.CreatePatientCaregiver:output_type -> heyrenee.v1.messages.PatientCaregiver
	14, // 23: heyrenee.v1.PatientService.UpdatePatientCaregiver:output_type -> heyrenee.v1.messages.PatientCaregiver
	9,  // 24: heyrenee.v1.PatientService.ListPatientCaregivers:output_type -> heyrenee.v1.ListPatientCaregiversResponse
	17, // [17:25] is the sub-list for method output_type
	9,  // [9:17] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_patient_service_proto_init() }
func file_heyrenee_v1_patient_service_proto_init() {
	if File_heyrenee_v1_patient_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_patient_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPatientRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePatientRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePatientProviderRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePatientProviderRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPatientProvidersRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPatientProvidersResponse); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePatientCaregiverRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePatientCaregiverRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPatientCaregiversRequest); i {
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
		file_heyrenee_v1_patient_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPatientCaregiversResponse); i {
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
			RawDescriptor: file_heyrenee_v1_patient_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_patient_service_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_patient_service_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_patient_service_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_patient_service_proto = out.File
	file_heyrenee_v1_patient_service_proto_rawDesc = nil
	file_heyrenee_v1_patient_service_proto_goTypes = nil
	file_heyrenee_v1_patient_service_proto_depIdxs = nil
}

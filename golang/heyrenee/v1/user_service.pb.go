// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/user_service.proto

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

// Request message for CreatePatient.
type CreatePatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Patient to create. Required.
	Patient *messages.Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
	// The Patient's password. If not provided, a random password will be generated.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreatePatientRequest) Reset() {
	*x = CreatePatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientRequest) ProtoMessage() {}

func (x *CreatePatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePatientRequest.ProtoReflect.Descriptor instead.
func (*CreatePatientRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePatientRequest) GetPatient() *messages.Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

func (x *CreatePatientRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

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
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPatientRequest) ProtoMessage() {}

func (x *GetPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[1]
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
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{1}
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
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientRequest) ProtoMessage() {}

func (x *UpdatePatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[2]
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
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePatientRequest) GetPatient() *messages.Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

// Request message for CreateCaregiver.
type CreateCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Caregiver to create. Required.
	Caregiver *messages.Caregiver `protobuf:"bytes,1,opt,name=caregiver,proto3" json:"caregiver,omitempty"`
	// Hash of the Caregiver's password. If not provided, a random password will be generated.
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *CreateCaregiverRequest) Reset() {
	*x = CreateCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCaregiverRequest) ProtoMessage() {}

func (x *CreateCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[3]
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
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCaregiverRequest) GetCaregiver() *messages.Caregiver {
	if x != nil {
		return x.Caregiver
	}
	return nil
}

func (x *CreateCaregiverRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// Request message for GetCaregiver.
type GetCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Caregiver to retrieve. Required.
	CaregiverId string `protobuf:"bytes,1,opt,name=caregiver_id,json=caregiverId,proto3" json:"caregiver_id,omitempty"`
}

func (x *GetCaregiverRequest) Reset() {
	*x = GetCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCaregiverRequest) ProtoMessage() {}

func (x *GetCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[4]
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
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCaregiverRequest) GetCaregiverId() string {
	if x != nil {
		return x.CaregiverId
	}
	return ""
}

// Request message for UpdateCaregiver.
type UpdateCaregiverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Caregiver to update. Required.
	Caregiver *messages.Caregiver `protobuf:"bytes,1,opt,name=caregiver,proto3" json:"caregiver,omitempty"`
}

func (x *UpdateCaregiverRequest) Reset() {
	*x = UpdateCaregiverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCaregiverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCaregiverRequest) ProtoMessage() {}

func (x *UpdateCaregiverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[5]
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
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCaregiverRequest) GetCaregiver() *messages.Caregiver {
	if x != nil {
		return x.Caregiver
	}
	return nil
}

// Request message for CreateConcierge.
type CreateConciergeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Concierge to create. Required.
	Concierge *messages.Concierge `protobuf:"bytes,1,opt,name=concierge,proto3" json:"concierge,omitempty"`
	// The Concierge's password. If not provided, a random password will be generated.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateConciergeRequest) Reset() {
	*x = CreateConciergeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConciergeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConciergeRequest) ProtoMessage() {}

func (x *CreateConciergeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConciergeRequest.ProtoReflect.Descriptor instead.
func (*CreateConciergeRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateConciergeRequest) GetConcierge() *messages.Concierge {
	if x != nil {
		return x.Concierge
	}
	return nil
}

func (x *CreateConciergeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_heyrenee_v1_user_service_proto protoreflect.FileDescriptor

var file_heyrenee_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x65, 0x72, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x38, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x6b, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3d, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69,
	0x76, 0x65, 0x72, 0x52, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x3e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0c, 0x63, 0x61, 0x72,
	0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0x80, 0xb5, 0x18, 0x02, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x57, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65,
	0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09,
	0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72,
	0x52, 0x09, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x22, 0x73, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x63, 0x69, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x65, 0x72,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x63, 0x69, 0x65, 0x72, 0x67, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x63, 0x69,
	0x65, 0x72, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x32, 0x80, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x12, 0x57, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x04, 0x80, 0xb5, 0x18,
	0x01, 0x12, 0x59, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x67,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x22,
	0x04, 0x80, 0xb5, 0x18, 0x01, 0x12, 0x5d, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x22, 0x04,
	0x80, 0xb5, 0x18, 0x01, 0x12, 0x5d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x63, 0x69, 0x65, 0x72, 0x67, 0x65, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x63,
	0x69, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68,
	0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x69, 0x65, 0x72, 0x67, 0x65, 0x22, 0x04, 0x80,
	0xb5, 0x18, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_user_service_proto_rawDescOnce sync.Once
	file_heyrenee_v1_user_service_proto_rawDescData = file_heyrenee_v1_user_service_proto_rawDesc
)

func file_heyrenee_v1_user_service_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_user_service_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_user_service_proto_rawDescData)
	})
	return file_heyrenee_v1_user_service_proto_rawDescData
}

var file_heyrenee_v1_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_heyrenee_v1_user_service_proto_goTypes = []interface{}{
	(*CreatePatientRequest)(nil),   // 0: heyrenee.v1.CreatePatientRequest
	(*GetPatientRequest)(nil),      // 1: heyrenee.v1.GetPatientRequest
	(*UpdatePatientRequest)(nil),   // 2: heyrenee.v1.UpdatePatientRequest
	(*CreateCaregiverRequest)(nil), // 3: heyrenee.v1.CreateCaregiverRequest
	(*GetCaregiverRequest)(nil),    // 4: heyrenee.v1.GetCaregiverRequest
	(*UpdateCaregiverRequest)(nil), // 5: heyrenee.v1.UpdateCaregiverRequest
	(*CreateConciergeRequest)(nil), // 6: heyrenee.v1.CreateConciergeRequest
	(*messages.Patient)(nil),       // 7: heyrenee.v1.messages.Patient
	(*messages.Caregiver)(nil),     // 8: heyrenee.v1.messages.Caregiver
	(*messages.Concierge)(nil),     // 9: heyrenee.v1.messages.Concierge
}
var file_heyrenee_v1_user_service_proto_depIdxs = []int32{
	7,  // 0: heyrenee.v1.CreatePatientRequest.patient:type_name -> heyrenee.v1.messages.Patient
	7,  // 1: heyrenee.v1.UpdatePatientRequest.patient:type_name -> heyrenee.v1.messages.Patient
	8,  // 2: heyrenee.v1.CreateCaregiverRequest.caregiver:type_name -> heyrenee.v1.messages.Caregiver
	8,  // 3: heyrenee.v1.UpdateCaregiverRequest.caregiver:type_name -> heyrenee.v1.messages.Caregiver
	9,  // 4: heyrenee.v1.CreateConciergeRequest.concierge:type_name -> heyrenee.v1.messages.Concierge
	0,  // 5: heyrenee.v1.UserService.CreatePatient:input_type -> heyrenee.v1.CreatePatientRequest
	1,  // 6: heyrenee.v1.UserService.GetPatient:input_type -> heyrenee.v1.GetPatientRequest
	2,  // 7: heyrenee.v1.UserService.UpdatePatient:input_type -> heyrenee.v1.UpdatePatientRequest
	3,  // 8: heyrenee.v1.UserService.CreateCaregiver:input_type -> heyrenee.v1.CreateCaregiverRequest
	4,  // 9: heyrenee.v1.UserService.GetCaregiver:input_type -> heyrenee.v1.GetCaregiverRequest
	5,  // 10: heyrenee.v1.UserService.UpdateCaregiver:input_type -> heyrenee.v1.UpdateCaregiverRequest
	6,  // 11: heyrenee.v1.UserService.CreateConcierge:input_type -> heyrenee.v1.CreateConciergeRequest
	7,  // 12: heyrenee.v1.UserService.CreatePatient:output_type -> heyrenee.v1.messages.Patient
	7,  // 13: heyrenee.v1.UserService.GetPatient:output_type -> heyrenee.v1.messages.Patient
	7,  // 14: heyrenee.v1.UserService.UpdatePatient:output_type -> heyrenee.v1.messages.Patient
	8,  // 15: heyrenee.v1.UserService.CreateCaregiver:output_type -> heyrenee.v1.messages.Caregiver
	8,  // 16: heyrenee.v1.UserService.GetCaregiver:output_type -> heyrenee.v1.messages.Caregiver
	8,  // 17: heyrenee.v1.UserService.UpdateCaregiver:output_type -> heyrenee.v1.messages.Caregiver
	9,  // 18: heyrenee.v1.UserService.CreateConcierge:output_type -> heyrenee.v1.messages.Concierge
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_user_service_proto_init() }
func file_heyrenee_v1_user_service_proto_init() {
	if File_heyrenee_v1_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePatientRequest); i {
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
		file_heyrenee_v1_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConciergeRequest); i {
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
			RawDescriptor: file_heyrenee_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_user_service_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_user_service_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_user_service_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_user_service_proto = out.File
	file_heyrenee_v1_user_service_proto_rawDesc = nil
	file_heyrenee_v1_user_service_proto_goTypes = nil
	file_heyrenee_v1_user_service_proto_depIdxs = nil
}

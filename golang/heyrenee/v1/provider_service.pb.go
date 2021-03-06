// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/provider_service.proto

package v1

import (
	enums "github.com/HeyReneeInc/proto/golang/heyrenee/v1/enums"
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

// Request message for SearchProviders.
type SearchProvidersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string          `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string          `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Specialty enums.Specialty `protobuf:"varint,3,opt,name=specialty,proto3,enum=heyrenee.v1.enums.Specialty" json:"specialty,omitempty"`
	City      string          `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	State     string          `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SearchProvidersRequest) Reset() {
	*x = SearchProvidersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_provider_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProvidersRequest) ProtoMessage() {}

func (x *SearchProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_provider_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProvidersRequest.ProtoReflect.Descriptor instead.
func (*SearchProvidersRequest) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_provider_service_proto_rawDescGZIP(), []int{0}
}

func (x *SearchProvidersRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SearchProvidersRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SearchProvidersRequest) GetSpecialty() enums.Specialty {
	if x != nil {
		return x.Specialty
	}
	return enums.Specialty(0)
}

func (x *SearchProvidersRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *SearchProvidersRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

// Response message for SearchProviders.
type SearchProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of Providers.
	Providers []*messages.Provider `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty"`
}

func (x *SearchProvidersResponse) Reset() {
	*x = SearchProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_provider_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProvidersResponse) ProtoMessage() {}

func (x *SearchProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_provider_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProvidersResponse.ProtoReflect.Descriptor instead.
func (*SearchProvidersResponse) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_provider_service_proto_rawDescGZIP(), []int{1}
}

func (x *SearchProvidersResponse) GetProviders() []*messages.Provider {
	if x != nil {
		return x.Providers
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
		mi := &file_heyrenee_v1_provider_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePatientProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePatientProviderRequest) ProtoMessage() {}

func (x *CreatePatientProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_provider_service_proto_msgTypes[2]
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
	return file_heyrenee_v1_provider_service_proto_rawDescGZIP(), []int{2}
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
		mi := &file_heyrenee_v1_provider_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePatientProviderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePatientProviderRequest) ProtoMessage() {}

func (x *UpdatePatientProviderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_provider_service_proto_msgTypes[3]
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
	return file_heyrenee_v1_provider_service_proto_rawDescGZIP(), []int{3}
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
		mi := &file_heyrenee_v1_provider_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientProvidersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientProvidersRequest) ProtoMessage() {}

func (x *ListPatientProvidersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_provider_service_proto_msgTypes[4]
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
	return file_heyrenee_v1_provider_service_proto_rawDescGZIP(), []int{4}
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

// Response message for ListPatientProviders.
type ListPatientProvidersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of PatientProviders.
	PatientProviders []*messages.PatientProvider `protobuf:"bytes,1,rep,name=patient_providers,json=patientProviders,proto3" json:"patient_providers,omitempty"`
}

func (x *ListPatientProvidersResponse) Reset() {
	*x = ListPatientProvidersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_provider_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPatientProvidersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPatientProvidersResponse) ProtoMessage() {}

func (x *ListPatientProvidersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_provider_service_proto_msgTypes[5]
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
	return file_heyrenee_v1_provider_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListPatientProvidersResponse) GetPatientProviders() []*messages.PatientProvider {
	if x != nil {
		return x.PatientProviders
	}
	return nil
}

var File_heyrenee_v1_provider_service_proto protoreflect.FileDescriptor

var file_heyrenee_v1_provider_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x2b, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x74, 0x79, 0x52, 0x09,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x57, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x22, 0x70, 0x0a, 0x1c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x10,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x70,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x70,
	0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50,
	0x0a, 0x10, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65,
	0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x0f, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x22, 0x86, 0x02, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x01, 0x52, 0x09, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x63, 0x0a, 0x17, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x15, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5d, 0x0a, 0x15, 0x70, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x72, 0x0a, 0x1c, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x11, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x10, 0x70, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x32, 0xba, 0x03,
	0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x68, 0x65, 0x79, 0x72,
	0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x6b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x68, 0x65, 0x79,
	0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x6b,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e,
	0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65, 0x79, 0x52, 0x65, 0x6e, 0x65,
	0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_provider_service_proto_rawDescOnce sync.Once
	file_heyrenee_v1_provider_service_proto_rawDescData = file_heyrenee_v1_provider_service_proto_rawDesc
)

func file_heyrenee_v1_provider_service_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_provider_service_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_provider_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_provider_service_proto_rawDescData)
	})
	return file_heyrenee_v1_provider_service_proto_rawDescData
}

var file_heyrenee_v1_provider_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_heyrenee_v1_provider_service_proto_goTypes = []interface{}{
	(*SearchProvidersRequest)(nil),       // 0: heyrenee.v1.SearchProvidersRequest
	(*SearchProvidersResponse)(nil),      // 1: heyrenee.v1.SearchProvidersResponse
	(*CreatePatientProviderRequest)(nil), // 2: heyrenee.v1.CreatePatientProviderRequest
	(*UpdatePatientProviderRequest)(nil), // 3: heyrenee.v1.UpdatePatientProviderRequest
	(*ListPatientProvidersRequest)(nil),  // 4: heyrenee.v1.ListPatientProvidersRequest
	(*ListPatientProvidersResponse)(nil), // 5: heyrenee.v1.ListPatientProvidersResponse
	(enums.Specialty)(0),                 // 6: heyrenee.v1.enums.Specialty
	(*messages.Provider)(nil),            // 7: heyrenee.v1.messages.Provider
	(*messages.PatientProvider)(nil),     // 8: heyrenee.v1.messages.PatientProvider
	(messages.PatientProviderStatus)(0),  // 9: heyrenee.v1.messages.PatientProviderStatus
	(messages.PatientProviderType)(0),    // 10: heyrenee.v1.messages.PatientProviderType
}
var file_heyrenee_v1_provider_service_proto_depIdxs = []int32{
	6,  // 0: heyrenee.v1.SearchProvidersRequest.specialty:type_name -> heyrenee.v1.enums.Specialty
	7,  // 1: heyrenee.v1.SearchProvidersResponse.providers:type_name -> heyrenee.v1.messages.Provider
	8,  // 2: heyrenee.v1.CreatePatientProviderRequest.patient_provider:type_name -> heyrenee.v1.messages.PatientProvider
	8,  // 3: heyrenee.v1.UpdatePatientProviderRequest.patient_provider:type_name -> heyrenee.v1.messages.PatientProvider
	9,  // 4: heyrenee.v1.ListPatientProvidersRequest.patient_provider_status:type_name -> heyrenee.v1.messages.PatientProviderStatus
	10, // 5: heyrenee.v1.ListPatientProvidersRequest.patient_provider_type:type_name -> heyrenee.v1.messages.PatientProviderType
	8,  // 6: heyrenee.v1.ListPatientProvidersResponse.patient_providers:type_name -> heyrenee.v1.messages.PatientProvider
	0,  // 7: heyrenee.v1.ProviderService.SearchProviders:input_type -> heyrenee.v1.SearchProvidersRequest
	2,  // 8: heyrenee.v1.ProviderService.CreatePatientProvider:input_type -> heyrenee.v1.CreatePatientProviderRequest
	3,  // 9: heyrenee.v1.ProviderService.UpdatePatientProvider:input_type -> heyrenee.v1.UpdatePatientProviderRequest
	4,  // 10: heyrenee.v1.ProviderService.ListPatientProviders:input_type -> heyrenee.v1.ListPatientProvidersRequest
	1,  // 11: heyrenee.v1.ProviderService.SearchProviders:output_type -> heyrenee.v1.SearchProvidersResponse
	8,  // 12: heyrenee.v1.ProviderService.CreatePatientProvider:output_type -> heyrenee.v1.messages.PatientProvider
	8,  // 13: heyrenee.v1.ProviderService.UpdatePatientProvider:output_type -> heyrenee.v1.messages.PatientProvider
	5,  // 14: heyrenee.v1.ProviderService.ListPatientProviders:output_type -> heyrenee.v1.ListPatientProvidersResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_provider_service_proto_init() }
func file_heyrenee_v1_provider_service_proto_init() {
	if File_heyrenee_v1_provider_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_provider_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProvidersRequest); i {
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
		file_heyrenee_v1_provider_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProvidersResponse); i {
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
		file_heyrenee_v1_provider_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_provider_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_provider_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_heyrenee_v1_provider_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_heyrenee_v1_provider_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heyrenee_v1_provider_service_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_provider_service_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_provider_service_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_provider_service_proto = out.File
	file_heyrenee_v1_provider_service_proto_rawDesc = nil
	file_heyrenee_v1_provider_service_proto_goTypes = nil
	file_heyrenee_v1_provider_service_proto_depIdxs = nil
}

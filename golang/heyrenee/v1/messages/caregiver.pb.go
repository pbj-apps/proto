// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: heyrenee/v1/messages/caregiver.proto

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

// A Caregiver represents a subuser that has access to the dashboard to monitor patients that they are caregivers for.
type Caregiver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the Caregiver. Required.
	CaregiverId string `protobuf:"bytes,1,opt,name=caregiver_id,json=caregiverId,proto3" json:"caregiver_id,omitempty"`
	// The legal first name of the Caregiver. Required.
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// The legal last name of the Caregiver. Required.
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The primary phone number for contacting the Caregiver. Required.
	PrimaryPhone string `protobuf:"bytes,4,opt,name=primary_phone,json=primaryPhone,proto3" json:"primary_phone,omitempty"`
	// The mobile phone number for contacting the Caregiver.
	MobilePhone string `protobuf:"bytes,5,opt,name=mobile_phone,json=mobilePhone,proto3" json:"mobile_phone,omitempty"`
	// The alternate phone number for contacting the Caregiver.
	OtherPhone string `protobuf:"bytes,6,opt,name=other_phone,json=otherPhone,proto3" json:"other_phone,omitempty"`
	// The email address for contacting the Caregiver. Required.
	Email string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	// The Caregiver's permanent residential address.
	AddressLines []string `protobuf:"bytes,8,rep,name=address_lines,json=addressLines,proto3" json:"address_lines,omitempty"`
	// The Caregiver's permanent residential city.
	City string `protobuf:"bytes,9,opt,name=city,proto3" json:"city,omitempty"`
	// The Caregiver's permanent residential state.
	State string `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	// The Caregiver's permanent residential zip code.
	Zip string `protobuf:"bytes,11,opt,name=zip,proto3" json:"zip,omitempty"`
}

func (x *Caregiver) Reset() {
	*x = Caregiver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heyrenee_v1_messages_caregiver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Caregiver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Caregiver) ProtoMessage() {}

func (x *Caregiver) ProtoReflect() protoreflect.Message {
	mi := &file_heyrenee_v1_messages_caregiver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Caregiver.ProtoReflect.Descriptor instead.
func (*Caregiver) Descriptor() ([]byte, []int) {
	return file_heyrenee_v1_messages_caregiver_proto_rawDescGZIP(), []int{0}
}

func (x *Caregiver) GetCaregiverId() string {
	if x != nil {
		return x.CaregiverId
	}
	return ""
}

func (x *Caregiver) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Caregiver) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Caregiver) GetPrimaryPhone() string {
	if x != nil {
		return x.PrimaryPhone
	}
	return ""
}

func (x *Caregiver) GetMobilePhone() string {
	if x != nil {
		return x.MobilePhone
	}
	return ""
}

func (x *Caregiver) GetOtherPhone() string {
	if x != nil {
		return x.OtherPhone
	}
	return ""
}

func (x *Caregiver) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Caregiver) GetAddressLines() []string {
	if x != nil {
		return x.AddressLines
	}
	return nil
}

func (x *Caregiver) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Caregiver) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Caregiver) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

var File_heyrenee_v1_messages_caregiver_proto protoreflect.FileDescriptor

var file_heyrenee_v1_messages_caregiver_proto_rawDesc = []byte{
	0x0a, 0x24, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1e, 0x68, 0x65,
	0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x02, 0x0a,
	0x09, 0x43, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0c, 0x63, 0x61,
	0x72, 0x65, 0x67, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0x80, 0xb5, 0x18, 0x02, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x65, 0x67, 0x69, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23,
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x6e, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x7a, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x7a, 0x69, 0x70, 0x42,
	0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x65,
	0x79, 0x52, 0x65, 0x6e, 0x65, 0x65, 0x49, 0x6e, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x79, 0x72, 0x65, 0x6e, 0x65, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_heyrenee_v1_messages_caregiver_proto_rawDescOnce sync.Once
	file_heyrenee_v1_messages_caregiver_proto_rawDescData = file_heyrenee_v1_messages_caregiver_proto_rawDesc
)

func file_heyrenee_v1_messages_caregiver_proto_rawDescGZIP() []byte {
	file_heyrenee_v1_messages_caregiver_proto_rawDescOnce.Do(func() {
		file_heyrenee_v1_messages_caregiver_proto_rawDescData = protoimpl.X.CompressGZIP(file_heyrenee_v1_messages_caregiver_proto_rawDescData)
	})
	return file_heyrenee_v1_messages_caregiver_proto_rawDescData
}

var file_heyrenee_v1_messages_caregiver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_heyrenee_v1_messages_caregiver_proto_goTypes = []interface{}{
	(*Caregiver)(nil), // 0: heyrenee.v1.messages.Caregiver
}
var file_heyrenee_v1_messages_caregiver_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heyrenee_v1_messages_caregiver_proto_init() }
func file_heyrenee_v1_messages_caregiver_proto_init() {
	if File_heyrenee_v1_messages_caregiver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heyrenee_v1_messages_caregiver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Caregiver); i {
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
			RawDescriptor: file_heyrenee_v1_messages_caregiver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_heyrenee_v1_messages_caregiver_proto_goTypes,
		DependencyIndexes: file_heyrenee_v1_messages_caregiver_proto_depIdxs,
		MessageInfos:      file_heyrenee_v1_messages_caregiver_proto_msgTypes,
	}.Build()
	File_heyrenee_v1_messages_caregiver_proto = out.File
	file_heyrenee_v1_messages_caregiver_proto_rawDesc = nil
	file_heyrenee_v1_messages_caregiver_proto_goTypes = nil
	file_heyrenee_v1_messages_caregiver_proto_depIdxs = nil
}

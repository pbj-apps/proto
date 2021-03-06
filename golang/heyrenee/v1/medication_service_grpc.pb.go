// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	messages "github.com/HeyReneeInc/proto/golang/heyrenee/v1/messages"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MedicationServiceClient is the client API for MedicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedicationServiceClient interface {
	// MedicationSuggest returns a list of MedicationSuggestions based on a partial search string. This should only be
	// used for autocomplete like features and not as a full fledged Medication search method.
	//
	// TODO(mitch): Use bi-directional streaming once client streaming supported by grpc-web
	// https://github.com/grpc/grpc-web/issues/24
	MedicationSuggest(ctx context.Context, in *MedicationSuggestRequest, opts ...grpc.CallOption) (*MedicationSuggestResponse, error)
	// CreatePrescription creates the provided Prescription resource.
	CreatePrescription(ctx context.Context, in *CreatePrescriptionRequest, opts ...grpc.CallOption) (*messages.Prescription, error)
	// UpdatePrescription updates the provided Prescription resource.
	UpdatePrescription(ctx context.Context, in *UpdatePrescriptionRequest, opts ...grpc.CallOption) (*messages.Prescription, error)
	// ListPrescriptions lists the Prescription resources for the specified Patient.
	ListPrescriptions(ctx context.Context, in *ListPrescriptionsRequest, opts ...grpc.CallOption) (*ListPrescriptionsResponse, error)
	// ListMedicationDoses lists the MedicationDoses for the specified Prescription.
	ListMedicationDoses(ctx context.Context, in *ListMedicationDosesRequest, opts ...grpc.CallOption) (*ListMedicationDosesResponse, error)
	// ListRefills lists the Refills for the specified Prescription.
	ListRefills(ctx context.Context, in *ListRefillsRequest, opts ...grpc.CallOption) (*ListRefillsResponse, error)
}

type medicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedicationServiceClient(cc grpc.ClientConnInterface) MedicationServiceClient {
	return &medicationServiceClient{cc}
}

func (c *medicationServiceClient) MedicationSuggest(ctx context.Context, in *MedicationSuggestRequest, opts ...grpc.CallOption) (*MedicationSuggestResponse, error) {
	out := new(MedicationSuggestResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.MedicationService/MedicationSuggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicationServiceClient) CreatePrescription(ctx context.Context, in *CreatePrescriptionRequest, opts ...grpc.CallOption) (*messages.Prescription, error) {
	out := new(messages.Prescription)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.MedicationService/CreatePrescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicationServiceClient) UpdatePrescription(ctx context.Context, in *UpdatePrescriptionRequest, opts ...grpc.CallOption) (*messages.Prescription, error) {
	out := new(messages.Prescription)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.MedicationService/UpdatePrescription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicationServiceClient) ListPrescriptions(ctx context.Context, in *ListPrescriptionsRequest, opts ...grpc.CallOption) (*ListPrescriptionsResponse, error) {
	out := new(ListPrescriptionsResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.MedicationService/ListPrescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicationServiceClient) ListMedicationDoses(ctx context.Context, in *ListMedicationDosesRequest, opts ...grpc.CallOption) (*ListMedicationDosesResponse, error) {
	out := new(ListMedicationDosesResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.MedicationService/ListMedicationDoses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *medicationServiceClient) ListRefills(ctx context.Context, in *ListRefillsRequest, opts ...grpc.CallOption) (*ListRefillsResponse, error) {
	out := new(ListRefillsResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.MedicationService/ListRefills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedicationServiceServer is the server API for MedicationService service.
// All implementations must embed UnimplementedMedicationServiceServer
// for forward compatibility
type MedicationServiceServer interface {
	// MedicationSuggest returns a list of MedicationSuggestions based on a partial search string. This should only be
	// used for autocomplete like features and not as a full fledged Medication search method.
	//
	// TODO(mitch): Use bi-directional streaming once client streaming supported by grpc-web
	// https://github.com/grpc/grpc-web/issues/24
	MedicationSuggest(context.Context, *MedicationSuggestRequest) (*MedicationSuggestResponse, error)
	// CreatePrescription creates the provided Prescription resource.
	CreatePrescription(context.Context, *CreatePrescriptionRequest) (*messages.Prescription, error)
	// UpdatePrescription updates the provided Prescription resource.
	UpdatePrescription(context.Context, *UpdatePrescriptionRequest) (*messages.Prescription, error)
	// ListPrescriptions lists the Prescription resources for the specified Patient.
	ListPrescriptions(context.Context, *ListPrescriptionsRequest) (*ListPrescriptionsResponse, error)
	// ListMedicationDoses lists the MedicationDoses for the specified Prescription.
	ListMedicationDoses(context.Context, *ListMedicationDosesRequest) (*ListMedicationDosesResponse, error)
	// ListRefills lists the Refills for the specified Prescription.
	ListRefills(context.Context, *ListRefillsRequest) (*ListRefillsResponse, error)
	mustEmbedUnimplementedMedicationServiceServer()
}

// UnimplementedMedicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMedicationServiceServer struct {
}

func (UnimplementedMedicationServiceServer) MedicationSuggest(context.Context, *MedicationSuggestRequest) (*MedicationSuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MedicationSuggest not implemented")
}
func (UnimplementedMedicationServiceServer) CreatePrescription(context.Context, *CreatePrescriptionRequest) (*messages.Prescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePrescription not implemented")
}
func (UnimplementedMedicationServiceServer) UpdatePrescription(context.Context, *UpdatePrescriptionRequest) (*messages.Prescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrescription not implemented")
}
func (UnimplementedMedicationServiceServer) ListPrescriptions(context.Context, *ListPrescriptionsRequest) (*ListPrescriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPrescriptions not implemented")
}
func (UnimplementedMedicationServiceServer) ListMedicationDoses(context.Context, *ListMedicationDosesRequest) (*ListMedicationDosesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMedicationDoses not implemented")
}
func (UnimplementedMedicationServiceServer) ListRefills(context.Context, *ListRefillsRequest) (*ListRefillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRefills not implemented")
}
func (UnimplementedMedicationServiceServer) mustEmbedUnimplementedMedicationServiceServer() {}

// UnsafeMedicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedicationServiceServer will
// result in compilation errors.
type UnsafeMedicationServiceServer interface {
	mustEmbedUnimplementedMedicationServiceServer()
}

func RegisterMedicationServiceServer(s grpc.ServiceRegistrar, srv MedicationServiceServer) {
	s.RegisterService(&MedicationService_ServiceDesc, srv)
}

func _MedicationService_MedicationSuggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MedicationSuggestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicationServiceServer).MedicationSuggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.MedicationService/MedicationSuggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicationServiceServer).MedicationSuggest(ctx, req.(*MedicationSuggestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicationService_CreatePrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePrescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicationServiceServer).CreatePrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.MedicationService/CreatePrescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicationServiceServer).CreatePrescription(ctx, req.(*CreatePrescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicationService_UpdatePrescription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePrescriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicationServiceServer).UpdatePrescription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.MedicationService/UpdatePrescription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicationServiceServer).UpdatePrescription(ctx, req.(*UpdatePrescriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicationService_ListPrescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPrescriptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicationServiceServer).ListPrescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.MedicationService/ListPrescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicationServiceServer).ListPrescriptions(ctx, req.(*ListPrescriptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicationService_ListMedicationDoses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMedicationDosesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicationServiceServer).ListMedicationDoses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.MedicationService/ListMedicationDoses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicationServiceServer).ListMedicationDoses(ctx, req.(*ListMedicationDosesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MedicationService_ListRefills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRefillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedicationServiceServer).ListRefills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.MedicationService/ListRefills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedicationServiceServer).ListRefills(ctx, req.(*ListRefillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MedicationService_ServiceDesc is the grpc.ServiceDesc for MedicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heyrenee.v1.MedicationService",
	HandlerType: (*MedicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MedicationSuggest",
			Handler:    _MedicationService_MedicationSuggest_Handler,
		},
		{
			MethodName: "CreatePrescription",
			Handler:    _MedicationService_CreatePrescription_Handler,
		},
		{
			MethodName: "UpdatePrescription",
			Handler:    _MedicationService_UpdatePrescription_Handler,
		},
		{
			MethodName: "ListPrescriptions",
			Handler:    _MedicationService_ListPrescriptions_Handler,
		},
		{
			MethodName: "ListMedicationDoses",
			Handler:    _MedicationService_ListMedicationDoses_Handler,
		},
		{
			MethodName: "ListRefills",
			Handler:    _MedicationService_ListRefills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heyrenee/v1/medication_service.proto",
}

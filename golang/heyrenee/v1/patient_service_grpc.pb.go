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

// PatientServiceClient is the client API for PatientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientServiceClient interface {
	GetPatient(ctx context.Context, in *GetPatientRequest, opts ...grpc.CallOption) (*messages.Patient, error)
	UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*messages.Patient, error)
	CreatePatientProvider(ctx context.Context, in *CreatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error)
	UpdatePatientProvider(ctx context.Context, in *UpdatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error)
	ListPatientProviders(ctx context.Context, in *ListPatientProvidersRequest, opts ...grpc.CallOption) (*ListPatientProvidersResponse, error)
	CreatePatientCaregiver(ctx context.Context, in *CreatePatientCaregiverRequest, opts ...grpc.CallOption) (*messages.PatientCaregiver, error)
	UpdatePatientCaregiver(ctx context.Context, in *UpdatePatientCaregiverRequest, opts ...grpc.CallOption) (*messages.PatientCaregiver, error)
	ListPatientCaregivers(ctx context.Context, in *ListPatientCaregiversRequest, opts ...grpc.CallOption) (*ListPatientCaregiversResponse, error)
}

type patientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientServiceClient(cc grpc.ClientConnInterface) PatientServiceClient {
	return &patientServiceClient{cc}
}

func (c *patientServiceClient) GetPatient(ctx context.Context, in *GetPatientRequest, opts ...grpc.CallOption) (*messages.Patient, error) {
	out := new(messages.Patient)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/GetPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*messages.Patient, error) {
	out := new(messages.Patient)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/UpdatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) CreatePatientProvider(ctx context.Context, in *CreatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error) {
	out := new(messages.PatientProvider)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/CreatePatientProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) UpdatePatientProvider(ctx context.Context, in *UpdatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error) {
	out := new(messages.PatientProvider)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/UpdatePatientProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) ListPatientProviders(ctx context.Context, in *ListPatientProvidersRequest, opts ...grpc.CallOption) (*ListPatientProvidersResponse, error) {
	out := new(ListPatientProvidersResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/ListPatientProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) CreatePatientCaregiver(ctx context.Context, in *CreatePatientCaregiverRequest, opts ...grpc.CallOption) (*messages.PatientCaregiver, error) {
	out := new(messages.PatientCaregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/CreatePatientCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) UpdatePatientCaregiver(ctx context.Context, in *UpdatePatientCaregiverRequest, opts ...grpc.CallOption) (*messages.PatientCaregiver, error) {
	out := new(messages.PatientCaregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/UpdatePatientCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) ListPatientCaregivers(ctx context.Context, in *ListPatientCaregiversRequest, opts ...grpc.CallOption) (*ListPatientCaregiversResponse, error) {
	out := new(ListPatientCaregiversResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/ListPatientCaregivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServiceServer is the server API for PatientService service.
// All implementations must embed UnimplementedPatientServiceServer
// for forward compatibility
type PatientServiceServer interface {
	GetPatient(context.Context, *GetPatientRequest) (*messages.Patient, error)
	UpdatePatient(context.Context, *UpdatePatientRequest) (*messages.Patient, error)
	CreatePatientProvider(context.Context, *CreatePatientProviderRequest) (*messages.PatientProvider, error)
	UpdatePatientProvider(context.Context, *UpdatePatientProviderRequest) (*messages.PatientProvider, error)
	ListPatientProviders(context.Context, *ListPatientProvidersRequest) (*ListPatientProvidersResponse, error)
	CreatePatientCaregiver(context.Context, *CreatePatientCaregiverRequest) (*messages.PatientCaregiver, error)
	UpdatePatientCaregiver(context.Context, *UpdatePatientCaregiverRequest) (*messages.PatientCaregiver, error)
	ListPatientCaregivers(context.Context, *ListPatientCaregiversRequest) (*ListPatientCaregiversResponse, error)
	mustEmbedUnimplementedPatientServiceServer()
}

// UnimplementedPatientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServiceServer struct {
}

func (UnimplementedPatientServiceServer) GetPatient(context.Context, *GetPatientRequest) (*messages.Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatient not implemented")
}
func (UnimplementedPatientServiceServer) UpdatePatient(context.Context, *UpdatePatientRequest) (*messages.Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatient not implemented")
}
func (UnimplementedPatientServiceServer) CreatePatientProvider(context.Context, *CreatePatientProviderRequest) (*messages.PatientProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientProvider not implemented")
}
func (UnimplementedPatientServiceServer) UpdatePatientProvider(context.Context, *UpdatePatientProviderRequest) (*messages.PatientProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatientProvider not implemented")
}
func (UnimplementedPatientServiceServer) ListPatientProviders(context.Context, *ListPatientProvidersRequest) (*ListPatientProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientProviders not implemented")
}
func (UnimplementedPatientServiceServer) CreatePatientCaregiver(context.Context, *CreatePatientCaregiverRequest) (*messages.PatientCaregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientCaregiver not implemented")
}
func (UnimplementedPatientServiceServer) UpdatePatientCaregiver(context.Context, *UpdatePatientCaregiverRequest) (*messages.PatientCaregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatientCaregiver not implemented")
}
func (UnimplementedPatientServiceServer) ListPatientCaregivers(context.Context, *ListPatientCaregiversRequest) (*ListPatientCaregiversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientCaregivers not implemented")
}
func (UnimplementedPatientServiceServer) mustEmbedUnimplementedPatientServiceServer() {}

// UnsafePatientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientServiceServer will
// result in compilation errors.
type UnsafePatientServiceServer interface {
	mustEmbedUnimplementedPatientServiceServer()
}

func RegisterPatientServiceServer(s grpc.ServiceRegistrar, srv PatientServiceServer) {
	s.RegisterService(&PatientService_ServiceDesc, srv)
}

func _PatientService_GetPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).GetPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/GetPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).GetPatient(ctx, req.(*GetPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_UpdatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).UpdatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/UpdatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).UpdatePatient(ctx, req.(*UpdatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_CreatePatientProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).CreatePatientProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/CreatePatientProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).CreatePatientProvider(ctx, req.(*CreatePatientProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_UpdatePatientProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).UpdatePatientProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/UpdatePatientProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).UpdatePatientProvider(ctx, req.(*UpdatePatientProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_ListPatientProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).ListPatientProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/ListPatientProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).ListPatientProviders(ctx, req.(*ListPatientProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_CreatePatientCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).CreatePatientCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/CreatePatientCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).CreatePatientCaregiver(ctx, req.(*CreatePatientCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_UpdatePatientCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).UpdatePatientCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/UpdatePatientCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).UpdatePatientCaregiver(ctx, req.(*UpdatePatientCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_ListPatientCaregivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientCaregiversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).ListPatientCaregivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/ListPatientCaregivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).ListPatientCaregivers(ctx, req.(*ListPatientCaregiversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientService_ServiceDesc is the grpc.ServiceDesc for PatientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heyrenee.v1.PatientService",
	HandlerType: (*PatientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPatient",
			Handler:    _PatientService_GetPatient_Handler,
		},
		{
			MethodName: "UpdatePatient",
			Handler:    _PatientService_UpdatePatient_Handler,
		},
		{
			MethodName: "CreatePatientProvider",
			Handler:    _PatientService_CreatePatientProvider_Handler,
		},
		{
			MethodName: "UpdatePatientProvider",
			Handler:    _PatientService_UpdatePatientProvider_Handler,
		},
		{
			MethodName: "ListPatientProviders",
			Handler:    _PatientService_ListPatientProviders_Handler,
		},
		{
			MethodName: "CreatePatientCaregiver",
			Handler:    _PatientService_CreatePatientCaregiver_Handler,
		},
		{
			MethodName: "UpdatePatientCaregiver",
			Handler:    _PatientService_UpdatePatientCaregiver_Handler,
		},
		{
			MethodName: "ListPatientCaregivers",
			Handler:    _PatientService_ListPatientCaregivers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heyrenee/v1/patient_service.proto",
}

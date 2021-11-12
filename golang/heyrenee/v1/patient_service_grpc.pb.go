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
	// CreatePatientCaregiver creates a specified PatientCaregiver.
	CreatePatientCaregiver(ctx context.Context, in *CreatePatientCaregiverRequest, opts ...grpc.CallOption) (*messages.PatientCaregiver, error)
	// UpdatePatientCaregiver updates a specified PatientCaregiver.
	UpdatePatientCaregiver(ctx context.Context, in *UpdatePatientCaregiverRequest, opts ...grpc.CallOption) (*messages.PatientCaregiver, error)
	// ListPatientCaregivers lists PatientCaregivers for a specific Patient.
	ListPatientCaregivers(ctx context.Context, in *ListPatientCaregiversRequest, opts ...grpc.CallOption) (*ListPatientCaregiversResponse, error)
	// CreatePatientAssessment creates a specified PatientAssessment.
	CreatePatientAssessment(ctx context.Context, in *CreatePatientAssessmentRequest, opts ...grpc.CallOption) (*messages.PatientAssessment, error)
	// ListPatientAssessments lists PatientAssessments for a specific Patient.
	ListPatientAssessments(ctx context.Context, in *ListPatientAssessmentsRequest, opts ...grpc.CallOption) (*ListPatientAssessmentsResponse, error)
	// CreatePatientSatisfactionQuestionnaire creates a specified PatientSatisfactionQuestionnaire.
	CreatePatientSatisfactionQuestionnaire(ctx context.Context, in *CreatePatientSatisfactionQuestionnaireRequest, opts ...grpc.CallOption) (*messages.PatientSatisfactionQuestionnaire, error)
	// ListPatientSatisfactionQuestionnaires lists PatientSatisfactionQuestionnaires for a specific Patient.
	ListPatientSatisfactionQuestionnaires(ctx context.Context, in *ListPatientSatisfactionQuestionnairesRequest, opts ...grpc.CallOption) (*ListPatientSatisfactionQuestionnairesResponse, error)
	// CreatePatientHealthQuestionnaire creates a specified PatientHealthQuestionnaire.
	CreatePatientHealthQuestionnaire(ctx context.Context, in *CreatePatientHealthQuestionnaireRequest, opts ...grpc.CallOption) (*messages.PatientHealthQuestionnaire, error)
	// ListPatientHealthQuestionnaires lists PatientHealthQuestionnaires for a specific Patient.
	ListPatientHealthQuestionnaires(ctx context.Context, in *ListPatientHealthQuestionnairesRequest, opts ...grpc.CallOption) (*ListPatientHealthQuestionnairesResponse, error)
	// CreatePatientSdohQuestionnaire creates a specified PatientSdohQuestionnaire.
	CreatePatientSdohQuestionnaire(ctx context.Context, in *CreatePatientSdohQuestionnaireRequest, opts ...grpc.CallOption) (*messages.PatientSdohQuestionnaire, error)
	// ListPatientSdohQuestionnaires lists PatientSdohQuestionnaires for a specific Patient.
	ListPatientSdohQuestionnaires(ctx context.Context, in *ListPatientSdohQuestionnairesRequest, opts ...grpc.CallOption) (*ListPatientSdohQuestionnairesResponse, error)
}

type patientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientServiceClient(cc grpc.ClientConnInterface) PatientServiceClient {
	return &patientServiceClient{cc}
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

func (c *patientServiceClient) CreatePatientAssessment(ctx context.Context, in *CreatePatientAssessmentRequest, opts ...grpc.CallOption) (*messages.PatientAssessment, error) {
	out := new(messages.PatientAssessment)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/CreatePatientAssessment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) ListPatientAssessments(ctx context.Context, in *ListPatientAssessmentsRequest, opts ...grpc.CallOption) (*ListPatientAssessmentsResponse, error) {
	out := new(ListPatientAssessmentsResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/ListPatientAssessments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) CreatePatientSatisfactionQuestionnaire(ctx context.Context, in *CreatePatientSatisfactionQuestionnaireRequest, opts ...grpc.CallOption) (*messages.PatientSatisfactionQuestionnaire, error) {
	out := new(messages.PatientSatisfactionQuestionnaire)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/CreatePatientSatisfactionQuestionnaire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) ListPatientSatisfactionQuestionnaires(ctx context.Context, in *ListPatientSatisfactionQuestionnairesRequest, opts ...grpc.CallOption) (*ListPatientSatisfactionQuestionnairesResponse, error) {
	out := new(ListPatientSatisfactionQuestionnairesResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/ListPatientSatisfactionQuestionnaires", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) CreatePatientHealthQuestionnaire(ctx context.Context, in *CreatePatientHealthQuestionnaireRequest, opts ...grpc.CallOption) (*messages.PatientHealthQuestionnaire, error) {
	out := new(messages.PatientHealthQuestionnaire)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/CreatePatientHealthQuestionnaire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) ListPatientHealthQuestionnaires(ctx context.Context, in *ListPatientHealthQuestionnairesRequest, opts ...grpc.CallOption) (*ListPatientHealthQuestionnairesResponse, error) {
	out := new(ListPatientHealthQuestionnairesResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/ListPatientHealthQuestionnaires", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) CreatePatientSdohQuestionnaire(ctx context.Context, in *CreatePatientSdohQuestionnaireRequest, opts ...grpc.CallOption) (*messages.PatientSdohQuestionnaire, error) {
	out := new(messages.PatientSdohQuestionnaire)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/CreatePatientSdohQuestionnaire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *patientServiceClient) ListPatientSdohQuestionnaires(ctx context.Context, in *ListPatientSdohQuestionnairesRequest, opts ...grpc.CallOption) (*ListPatientSdohQuestionnairesResponse, error) {
	out := new(ListPatientSdohQuestionnairesResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.PatientService/ListPatientSdohQuestionnaires", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServiceServer is the server API for PatientService service.
// All implementations must embed UnimplementedPatientServiceServer
// for forward compatibility
type PatientServiceServer interface {
	// CreatePatientCaregiver creates a specified PatientCaregiver.
	CreatePatientCaregiver(context.Context, *CreatePatientCaregiverRequest) (*messages.PatientCaregiver, error)
	// UpdatePatientCaregiver updates a specified PatientCaregiver.
	UpdatePatientCaregiver(context.Context, *UpdatePatientCaregiverRequest) (*messages.PatientCaregiver, error)
	// ListPatientCaregivers lists PatientCaregivers for a specific Patient.
	ListPatientCaregivers(context.Context, *ListPatientCaregiversRequest) (*ListPatientCaregiversResponse, error)
	// CreatePatientAssessment creates a specified PatientAssessment.
	CreatePatientAssessment(context.Context, *CreatePatientAssessmentRequest) (*messages.PatientAssessment, error)
	// ListPatientAssessments lists PatientAssessments for a specific Patient.
	ListPatientAssessments(context.Context, *ListPatientAssessmentsRequest) (*ListPatientAssessmentsResponse, error)
	// CreatePatientSatisfactionQuestionnaire creates a specified PatientSatisfactionQuestionnaire.
	CreatePatientSatisfactionQuestionnaire(context.Context, *CreatePatientSatisfactionQuestionnaireRequest) (*messages.PatientSatisfactionQuestionnaire, error)
	// ListPatientSatisfactionQuestionnaires lists PatientSatisfactionQuestionnaires for a specific Patient.
	ListPatientSatisfactionQuestionnaires(context.Context, *ListPatientSatisfactionQuestionnairesRequest) (*ListPatientSatisfactionQuestionnairesResponse, error)
	// CreatePatientHealthQuestionnaire creates a specified PatientHealthQuestionnaire.
	CreatePatientHealthQuestionnaire(context.Context, *CreatePatientHealthQuestionnaireRequest) (*messages.PatientHealthQuestionnaire, error)
	// ListPatientHealthQuestionnaires lists PatientHealthQuestionnaires for a specific Patient.
	ListPatientHealthQuestionnaires(context.Context, *ListPatientHealthQuestionnairesRequest) (*ListPatientHealthQuestionnairesResponse, error)
	// CreatePatientSdohQuestionnaire creates a specified PatientSdohQuestionnaire.
	CreatePatientSdohQuestionnaire(context.Context, *CreatePatientSdohQuestionnaireRequest) (*messages.PatientSdohQuestionnaire, error)
	// ListPatientSdohQuestionnaires lists PatientSdohQuestionnaires for a specific Patient.
	ListPatientSdohQuestionnaires(context.Context, *ListPatientSdohQuestionnairesRequest) (*ListPatientSdohQuestionnairesResponse, error)
	mustEmbedUnimplementedPatientServiceServer()
}

// UnimplementedPatientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServiceServer struct {
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
func (UnimplementedPatientServiceServer) CreatePatientAssessment(context.Context, *CreatePatientAssessmentRequest) (*messages.PatientAssessment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientAssessment not implemented")
}
func (UnimplementedPatientServiceServer) ListPatientAssessments(context.Context, *ListPatientAssessmentsRequest) (*ListPatientAssessmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientAssessments not implemented")
}
func (UnimplementedPatientServiceServer) CreatePatientSatisfactionQuestionnaire(context.Context, *CreatePatientSatisfactionQuestionnaireRequest) (*messages.PatientSatisfactionQuestionnaire, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientSatisfactionQuestionnaire not implemented")
}
func (UnimplementedPatientServiceServer) ListPatientSatisfactionQuestionnaires(context.Context, *ListPatientSatisfactionQuestionnairesRequest) (*ListPatientSatisfactionQuestionnairesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientSatisfactionQuestionnaires not implemented")
}
func (UnimplementedPatientServiceServer) CreatePatientHealthQuestionnaire(context.Context, *CreatePatientHealthQuestionnaireRequest) (*messages.PatientHealthQuestionnaire, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientHealthQuestionnaire not implemented")
}
func (UnimplementedPatientServiceServer) ListPatientHealthQuestionnaires(context.Context, *ListPatientHealthQuestionnairesRequest) (*ListPatientHealthQuestionnairesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientHealthQuestionnaires not implemented")
}
func (UnimplementedPatientServiceServer) CreatePatientSdohQuestionnaire(context.Context, *CreatePatientSdohQuestionnaireRequest) (*messages.PatientSdohQuestionnaire, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientSdohQuestionnaire not implemented")
}
func (UnimplementedPatientServiceServer) ListPatientSdohQuestionnaires(context.Context, *ListPatientSdohQuestionnairesRequest) (*ListPatientSdohQuestionnairesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientSdohQuestionnaires not implemented")
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

func _PatientService_CreatePatientAssessment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientAssessmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).CreatePatientAssessment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/CreatePatientAssessment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).CreatePatientAssessment(ctx, req.(*CreatePatientAssessmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_ListPatientAssessments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientAssessmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).ListPatientAssessments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/ListPatientAssessments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).ListPatientAssessments(ctx, req.(*ListPatientAssessmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_CreatePatientSatisfactionQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientSatisfactionQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).CreatePatientSatisfactionQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/CreatePatientSatisfactionQuestionnaire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).CreatePatientSatisfactionQuestionnaire(ctx, req.(*CreatePatientSatisfactionQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_ListPatientSatisfactionQuestionnaires_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientSatisfactionQuestionnairesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).ListPatientSatisfactionQuestionnaires(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/ListPatientSatisfactionQuestionnaires",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).ListPatientSatisfactionQuestionnaires(ctx, req.(*ListPatientSatisfactionQuestionnairesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_CreatePatientHealthQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientHealthQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).CreatePatientHealthQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/CreatePatientHealthQuestionnaire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).CreatePatientHealthQuestionnaire(ctx, req.(*CreatePatientHealthQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_ListPatientHealthQuestionnaires_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientHealthQuestionnairesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).ListPatientHealthQuestionnaires(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/ListPatientHealthQuestionnaires",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).ListPatientHealthQuestionnaires(ctx, req.(*ListPatientHealthQuestionnairesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_CreatePatientSdohQuestionnaire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientSdohQuestionnaireRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).CreatePatientSdohQuestionnaire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/CreatePatientSdohQuestionnaire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).CreatePatientSdohQuestionnaire(ctx, req.(*CreatePatientSdohQuestionnaireRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PatientService_ListPatientSdohQuestionnaires_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientSdohQuestionnairesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).ListPatientSdohQuestionnaires(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.PatientService/ListPatientSdohQuestionnaires",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).ListPatientSdohQuestionnaires(ctx, req.(*ListPatientSdohQuestionnairesRequest))
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
		{
			MethodName: "CreatePatientAssessment",
			Handler:    _PatientService_CreatePatientAssessment_Handler,
		},
		{
			MethodName: "ListPatientAssessments",
			Handler:    _PatientService_ListPatientAssessments_Handler,
		},
		{
			MethodName: "CreatePatientSatisfactionQuestionnaire",
			Handler:    _PatientService_CreatePatientSatisfactionQuestionnaire_Handler,
		},
		{
			MethodName: "ListPatientSatisfactionQuestionnaires",
			Handler:    _PatientService_ListPatientSatisfactionQuestionnaires_Handler,
		},
		{
			MethodName: "CreatePatientHealthQuestionnaire",
			Handler:    _PatientService_CreatePatientHealthQuestionnaire_Handler,
		},
		{
			MethodName: "ListPatientHealthQuestionnaires",
			Handler:    _PatientService_ListPatientHealthQuestionnaires_Handler,
		},
		{
			MethodName: "CreatePatientSdohQuestionnaire",
			Handler:    _PatientService_CreatePatientSdohQuestionnaire_Handler,
		},
		{
			MethodName: "ListPatientSdohQuestionnaires",
			Handler:    _PatientService_ListPatientSdohQuestionnaires_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heyrenee/v1/patient_service.proto",
}

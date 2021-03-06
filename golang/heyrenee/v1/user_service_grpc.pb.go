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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// CreatePatient creates a new User with a Patient subuser. If the request is unauthenticated, a standalone user with
	// a Patient subuser will be created. If the request is authenticated by a Caregiver or Concierge subuser then the new
	// Patient subuser will be created with a PatientCaregiver or a PatientConcierge for the authenticated subuser.
	CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*messages.Patient, error)
	// GetPatient retrieves a specified Patient subuser.
	GetPatient(ctx context.Context, in *GetPatientRequest, opts ...grpc.CallOption) (*messages.Patient, error)
	// UpdatePatient updates a specified Patient subuser.
	UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*messages.Patient, error)
	// CreateCaregiver creates a new User with a Caregiver subuser.
	CreateCaregiver(ctx context.Context, in *CreateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error)
	// GetCaregiver retrieves a specified Caregiver subuser.
	GetCaregiver(ctx context.Context, in *GetCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error)
	// UpdateCaregiver updates a specified Caregiver subuser.
	UpdateCaregiver(ctx context.Context, in *UpdateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreatePatient(ctx context.Context, in *CreatePatientRequest, opts ...grpc.CallOption) (*messages.Patient, error) {
	out := new(messages.Patient)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.UserService/CreatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPatient(ctx context.Context, in *GetPatientRequest, opts ...grpc.CallOption) (*messages.Patient, error) {
	out := new(messages.Patient)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.UserService/GetPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePatient(ctx context.Context, in *UpdatePatientRequest, opts ...grpc.CallOption) (*messages.Patient, error) {
	out := new(messages.Patient)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.UserService/UpdatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateCaregiver(ctx context.Context, in *CreateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error) {
	out := new(messages.Caregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.UserService/CreateCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetCaregiver(ctx context.Context, in *GetCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error) {
	out := new(messages.Caregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.UserService/GetCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateCaregiver(ctx context.Context, in *UpdateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error) {
	out := new(messages.Caregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.UserService/UpdateCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	// CreatePatient creates a new User with a Patient subuser. If the request is unauthenticated, a standalone user with
	// a Patient subuser will be created. If the request is authenticated by a Caregiver or Concierge subuser then the new
	// Patient subuser will be created with a PatientCaregiver or a PatientConcierge for the authenticated subuser.
	CreatePatient(context.Context, *CreatePatientRequest) (*messages.Patient, error)
	// GetPatient retrieves a specified Patient subuser.
	GetPatient(context.Context, *GetPatientRequest) (*messages.Patient, error)
	// UpdatePatient updates a specified Patient subuser.
	UpdatePatient(context.Context, *UpdatePatientRequest) (*messages.Patient, error)
	// CreateCaregiver creates a new User with a Caregiver subuser.
	CreateCaregiver(context.Context, *CreateCaregiverRequest) (*messages.Caregiver, error)
	// GetCaregiver retrieves a specified Caregiver subuser.
	GetCaregiver(context.Context, *GetCaregiverRequest) (*messages.Caregiver, error)
	// UpdateCaregiver updates a specified Caregiver subuser.
	UpdateCaregiver(context.Context, *UpdateCaregiverRequest) (*messages.Caregiver, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreatePatient(context.Context, *CreatePatientRequest) (*messages.Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}
func (UnimplementedUserServiceServer) GetPatient(context.Context, *GetPatientRequest) (*messages.Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatient not implemented")
}
func (UnimplementedUserServiceServer) UpdatePatient(context.Context, *UpdatePatientRequest) (*messages.Patient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatient not implemented")
}
func (UnimplementedUserServiceServer) CreateCaregiver(context.Context, *CreateCaregiverRequest) (*messages.Caregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCaregiver not implemented")
}
func (UnimplementedUserServiceServer) GetCaregiver(context.Context, *GetCaregiverRequest) (*messages.Caregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaregiver not implemented")
}
func (UnimplementedUserServiceServer) UpdateCaregiver(context.Context, *UpdateCaregiverRequest) (*messages.Caregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCaregiver not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.UserService/CreatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreatePatient(ctx, req.(*CreatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.UserService/GetPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPatient(ctx, req.(*GetPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.UserService/UpdatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePatient(ctx, req.(*UpdatePatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.UserService/CreateCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateCaregiver(ctx, req.(*CreateCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.UserService/GetCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetCaregiver(ctx, req.(*GetCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.UserService/UpdateCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateCaregiver(ctx, req.(*UpdateCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heyrenee.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePatient",
			Handler:    _UserService_CreatePatient_Handler,
		},
		{
			MethodName: "GetPatient",
			Handler:    _UserService_GetPatient_Handler,
		},
		{
			MethodName: "UpdatePatient",
			Handler:    _UserService_UpdatePatient_Handler,
		},
		{
			MethodName: "CreateCaregiver",
			Handler:    _UserService_CreateCaregiver_Handler,
		},
		{
			MethodName: "GetCaregiver",
			Handler:    _UserService_GetCaregiver_Handler,
		},
		{
			MethodName: "UpdateCaregiver",
			Handler:    _UserService_UpdateCaregiver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heyrenee/v1/user_service.proto",
}

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

// CaregiverServiceClient is the client API for CaregiverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CaregiverServiceClient interface {
	CreateCaregiver(ctx context.Context, in *CreateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error)
	GetCaregiver(ctx context.Context, in *GetCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error)
	UpdateCaregiver(ctx context.Context, in *UpdateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error)
}

type caregiverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaregiverServiceClient(cc grpc.ClientConnInterface) CaregiverServiceClient {
	return &caregiverServiceClient{cc}
}

func (c *caregiverServiceClient) CreateCaregiver(ctx context.Context, in *CreateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error) {
	out := new(messages.Caregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.CaregiverService/CreateCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caregiverServiceClient) GetCaregiver(ctx context.Context, in *GetCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error) {
	out := new(messages.Caregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.CaregiverService/GetCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *caregiverServiceClient) UpdateCaregiver(ctx context.Context, in *UpdateCaregiverRequest, opts ...grpc.CallOption) (*messages.Caregiver, error) {
	out := new(messages.Caregiver)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.CaregiverService/UpdateCaregiver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaregiverServiceServer is the server API for CaregiverService service.
// All implementations must embed UnimplementedCaregiverServiceServer
// for forward compatibility
type CaregiverServiceServer interface {
	CreateCaregiver(context.Context, *CreateCaregiverRequest) (*messages.Caregiver, error)
	GetCaregiver(context.Context, *GetCaregiverRequest) (*messages.Caregiver, error)
	UpdateCaregiver(context.Context, *UpdateCaregiverRequest) (*messages.Caregiver, error)
	mustEmbedUnimplementedCaregiverServiceServer()
}

// UnimplementedCaregiverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCaregiverServiceServer struct {
}

func (UnimplementedCaregiverServiceServer) CreateCaregiver(context.Context, *CreateCaregiverRequest) (*messages.Caregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCaregiver not implemented")
}
func (UnimplementedCaregiverServiceServer) GetCaregiver(context.Context, *GetCaregiverRequest) (*messages.Caregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaregiver not implemented")
}
func (UnimplementedCaregiverServiceServer) UpdateCaregiver(context.Context, *UpdateCaregiverRequest) (*messages.Caregiver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCaregiver not implemented")
}
func (UnimplementedCaregiverServiceServer) mustEmbedUnimplementedCaregiverServiceServer() {}

// UnsafeCaregiverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaregiverServiceServer will
// result in compilation errors.
type UnsafeCaregiverServiceServer interface {
	mustEmbedUnimplementedCaregiverServiceServer()
}

func RegisterCaregiverServiceServer(s grpc.ServiceRegistrar, srv CaregiverServiceServer) {
	s.RegisterService(&CaregiverService_ServiceDesc, srv)
}

func _CaregiverService_CreateCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaregiverServiceServer).CreateCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.CaregiverService/CreateCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaregiverServiceServer).CreateCaregiver(ctx, req.(*CreateCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaregiverService_GetCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaregiverServiceServer).GetCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.CaregiverService/GetCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaregiverServiceServer).GetCaregiver(ctx, req.(*GetCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaregiverService_UpdateCaregiver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaregiverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaregiverServiceServer).UpdateCaregiver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.CaregiverService/UpdateCaregiver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaregiverServiceServer).UpdateCaregiver(ctx, req.(*UpdateCaregiverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaregiverService_ServiceDesc is the grpc.ServiceDesc for CaregiverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaregiverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heyrenee.v1.CaregiverService",
	HandlerType: (*CaregiverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCaregiver",
			Handler:    _CaregiverService_CreateCaregiver_Handler,
		},
		{
			MethodName: "GetCaregiver",
			Handler:    _CaregiverService_GetCaregiver_Handler,
		},
		{
			MethodName: "UpdateCaregiver",
			Handler:    _CaregiverService_UpdateCaregiver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heyrenee/v1/caregiver_service.proto",
}

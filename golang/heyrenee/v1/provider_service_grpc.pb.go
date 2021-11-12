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

// ProviderServiceClient is the client API for ProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServiceClient interface {
	// ProviderSuggest returns a list of ProviderSuggestions based on a partial search string. This should only be
	// used for autocomplete like features and not as a full fledged Provider search method.
	//
	// TODO(mitch): Use bi-directional streaming once client streaming supported by grpc-web
	// https://github.com/grpc/grpc-web/issues/24
	ProviderSuggest(ctx context.Context, in *ProviderSuggestRequest, opts ...grpc.CallOption) (*ProviderSuggestResponse, error)
	// CreatePatientProvider creates a specified PatientProvider.
	CreatePatientProvider(ctx context.Context, in *CreatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error)
	// UpdatePatientProvider updates a specified PatientProvider.
	UpdatePatientProvider(ctx context.Context, in *UpdatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error)
	// ListPatientProviders lists PatientProviders for a specified Patient.
	ListPatientProviders(ctx context.Context, in *ListPatientProvidersRequest, opts ...grpc.CallOption) (*ListPatientProvidersResponse, error)
}

type providerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServiceClient(cc grpc.ClientConnInterface) ProviderServiceClient {
	return &providerServiceClient{cc}
}

func (c *providerServiceClient) ProviderSuggest(ctx context.Context, in *ProviderSuggestRequest, opts ...grpc.CallOption) (*ProviderSuggestResponse, error) {
	out := new(ProviderSuggestResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.ProviderService/ProviderSuggest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) CreatePatientProvider(ctx context.Context, in *CreatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error) {
	out := new(messages.PatientProvider)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.ProviderService/CreatePatientProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) UpdatePatientProvider(ctx context.Context, in *UpdatePatientProviderRequest, opts ...grpc.CallOption) (*messages.PatientProvider, error) {
	out := new(messages.PatientProvider)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.ProviderService/UpdatePatientProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) ListPatientProviders(ctx context.Context, in *ListPatientProvidersRequest, opts ...grpc.CallOption) (*ListPatientProvidersResponse, error) {
	out := new(ListPatientProvidersResponse)
	err := c.cc.Invoke(ctx, "/heyrenee.v1.ProviderService/ListPatientProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServiceServer is the server API for ProviderService service.
// All implementations must embed UnimplementedProviderServiceServer
// for forward compatibility
type ProviderServiceServer interface {
	// ProviderSuggest returns a list of ProviderSuggestions based on a partial search string. This should only be
	// used for autocomplete like features and not as a full fledged Provider search method.
	//
	// TODO(mitch): Use bi-directional streaming once client streaming supported by grpc-web
	// https://github.com/grpc/grpc-web/issues/24
	ProviderSuggest(context.Context, *ProviderSuggestRequest) (*ProviderSuggestResponse, error)
	// CreatePatientProvider creates a specified PatientProvider.
	CreatePatientProvider(context.Context, *CreatePatientProviderRequest) (*messages.PatientProvider, error)
	// UpdatePatientProvider updates a specified PatientProvider.
	UpdatePatientProvider(context.Context, *UpdatePatientProviderRequest) (*messages.PatientProvider, error)
	// ListPatientProviders lists PatientProviders for a specified Patient.
	ListPatientProviders(context.Context, *ListPatientProvidersRequest) (*ListPatientProvidersResponse, error)
	mustEmbedUnimplementedProviderServiceServer()
}

// UnimplementedProviderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServiceServer struct {
}

func (UnimplementedProviderServiceServer) ProviderSuggest(context.Context, *ProviderSuggestRequest) (*ProviderSuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProviderSuggest not implemented")
}
func (UnimplementedProviderServiceServer) CreatePatientProvider(context.Context, *CreatePatientProviderRequest) (*messages.PatientProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatientProvider not implemented")
}
func (UnimplementedProviderServiceServer) UpdatePatientProvider(context.Context, *UpdatePatientProviderRequest) (*messages.PatientProvider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatientProvider not implemented")
}
func (UnimplementedProviderServiceServer) ListPatientProviders(context.Context, *ListPatientProvidersRequest) (*ListPatientProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPatientProviders not implemented")
}
func (UnimplementedProviderServiceServer) mustEmbedUnimplementedProviderServiceServer() {}

// UnsafeProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServiceServer will
// result in compilation errors.
type UnsafeProviderServiceServer interface {
	mustEmbedUnimplementedProviderServiceServer()
}

func RegisterProviderServiceServer(s grpc.ServiceRegistrar, srv ProviderServiceServer) {
	s.RegisterService(&ProviderService_ServiceDesc, srv)
}

func _ProviderService_ProviderSuggest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderSuggestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).ProviderSuggest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.ProviderService/ProviderSuggest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).ProviderSuggest(ctx, req.(*ProviderSuggestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_CreatePatientProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePatientProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).CreatePatientProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.ProviderService/CreatePatientProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).CreatePatientProvider(ctx, req.(*CreatePatientProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_UpdatePatientProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatientProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).UpdatePatientProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.ProviderService/UpdatePatientProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).UpdatePatientProvider(ctx, req.(*UpdatePatientProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_ListPatientProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPatientProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).ListPatientProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heyrenee.v1.ProviderService/ListPatientProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).ListPatientProviders(ctx, req.(*ListPatientProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderService_ServiceDesc is the grpc.ServiceDesc for ProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "heyrenee.v1.ProviderService",
	HandlerType: (*ProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProviderSuggest",
			Handler:    _ProviderService_ProviderSuggest_Handler,
		},
		{
			MethodName: "CreatePatientProvider",
			Handler:    _ProviderService_CreatePatientProvider_Handler,
		},
		{
			MethodName: "UpdatePatientProvider",
			Handler:    _ProviderService_UpdatePatientProvider_Handler,
		},
		{
			MethodName: "ListPatientProviders",
			Handler:    _ProviderService_ListPatientProviders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heyrenee/v1/provider_service.proto",
}

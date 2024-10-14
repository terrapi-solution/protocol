// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: deployment/deployment.proto

package deployment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeploymentService_Get_FullMethodName = "/terrapi.v1.DeploymentService/Get"
)

// DeploymentServiceClient is the client API for DeploymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploymentServiceClient interface {
	Get(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Deployment, error)
}

type deploymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentServiceClient(cc grpc.ClientConnInterface) DeploymentServiceClient {
	return &deploymentServiceClient{cc}
}

func (c *deploymentServiceClient) Get(ctx context.Context, in *RetrieveRequest, opts ...grpc.CallOption) (*Deployment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Deployment)
	err := c.cc.Invoke(ctx, DeploymentService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServiceServer is the server API for DeploymentService service.
// All implementations must embed UnimplementedDeploymentServiceServer
// for forward compatibility.
type DeploymentServiceServer interface {
	Get(context.Context, *RetrieveRequest) (*Deployment, error)
	mustEmbedUnimplementedDeploymentServiceServer()
}

// UnimplementedDeploymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeploymentServiceServer struct{}

func (UnimplementedDeploymentServiceServer) Get(context.Context, *RetrieveRequest) (*Deployment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDeploymentServiceServer) mustEmbedUnimplementedDeploymentServiceServer() {}
func (UnimplementedDeploymentServiceServer) testEmbeddedByValue()                           {}

// UnsafeDeploymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploymentServiceServer will
// result in compilation errors.
type UnsafeDeploymentServiceServer interface {
	mustEmbedUnimplementedDeploymentServiceServer()
}

func RegisterDeploymentServiceServer(s grpc.ServiceRegistrar, srv DeploymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeploymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeploymentService_ServiceDesc, srv)
}

func _DeploymentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeploymentService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServiceServer).Get(ctx, req.(*RetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploymentService_ServiceDesc is the grpc.ServiceDesc for DeploymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "terrapi.v1.DeploymentService",
	HandlerType: (*DeploymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DeploymentService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployment/deployment.proto",
}

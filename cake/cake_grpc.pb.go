// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cake

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CakeClient is the client API for Cake service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CakeClient interface {
	ListProblems(ctx context.Context, in *ProblemInput, opts ...grpc.CallOption) (Cake_ListProblemsClient, error)
	RunProblem(ctx context.Context, in *ProblemDefinition, opts ...grpc.CallOption) (*ProblemResult, error)
}

type cakeClient struct {
	cc grpc.ClientConnInterface
}

func NewCakeClient(cc grpc.ClientConnInterface) CakeClient {
	return &cakeClient{cc}
}

func (c *cakeClient) ListProblems(ctx context.Context, in *ProblemInput, opts ...grpc.CallOption) (Cake_ListProblemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Cake_ServiceDesc.Streams[0], "/cake.Cake/ListProblems", opts...)
	if err != nil {
		return nil, err
	}
	x := &cakeListProblemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cake_ListProblemsClient interface {
	Recv() (*ProblemDefinition, error)
	grpc.ClientStream
}

type cakeListProblemsClient struct {
	grpc.ClientStream
}

func (x *cakeListProblemsClient) Recv() (*ProblemDefinition, error) {
	m := new(ProblemDefinition)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cakeClient) RunProblem(ctx context.Context, in *ProblemDefinition, opts ...grpc.CallOption) (*ProblemResult, error) {
	out := new(ProblemResult)
	err := c.cc.Invoke(ctx, "/cake.Cake/RunProblem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CakeServer is the server API for Cake service.
// All implementations must embed UnimplementedCakeServer
// for forward compatibility
type CakeServer interface {
	ListProblems(*ProblemInput, Cake_ListProblemsServer) error
	RunProblem(context.Context, *ProblemDefinition) (*ProblemResult, error)
	mustEmbedUnimplementedCakeServer()
}

// UnimplementedCakeServer must be embedded to have forward compatible implementations.
type UnimplementedCakeServer struct {
}

func (UnimplementedCakeServer) ListProblems(*ProblemInput, Cake_ListProblemsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProblems not implemented")
}
func (UnimplementedCakeServer) RunProblem(context.Context, *ProblemDefinition) (*ProblemResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunProblem not implemented")
}
func (UnimplementedCakeServer) mustEmbedUnimplementedCakeServer() {}

// UnsafeCakeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CakeServer will
// result in compilation errors.
type UnsafeCakeServer interface {
	mustEmbedUnimplementedCakeServer()
}

func RegisterCakeServer(s grpc.ServiceRegistrar, srv CakeServer) {
	s.RegisterService(&Cake_ServiceDesc, srv)
}

func _Cake_ListProblems_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProblemInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CakeServer).ListProblems(m, &cakeListProblemsServer{stream})
}

type Cake_ListProblemsServer interface {
	Send(*ProblemDefinition) error
	grpc.ServerStream
}

type cakeListProblemsServer struct {
	grpc.ServerStream
}

func (x *cakeListProblemsServer) Send(m *ProblemDefinition) error {
	return x.ServerStream.SendMsg(m)
}

func _Cake_RunProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProblemDefinition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CakeServer).RunProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cake.Cake/RunProblem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CakeServer).RunProblem(ctx, req.(*ProblemDefinition))
	}
	return interceptor(ctx, in, info, handler)
}

// Cake_ServiceDesc is the grpc.ServiceDesc for Cake service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cake_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cake.Cake",
	HandlerType: (*CakeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunProblem",
			Handler:    _Cake_RunProblem_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProblems",
			Handler:       _Cake_ListProblems_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cake/cake.proto",
}

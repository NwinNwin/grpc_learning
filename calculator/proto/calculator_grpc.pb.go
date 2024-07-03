// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: calculator.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CalculatorService_Sum_FullMethodName     = "/calculator.CalculatorService/Sum"
	CalculatorService_Prime_FullMethodName   = "/calculator.CalculatorService/Prime"
	CalculatorService_Average_FullMethodName = "/calculator.CalculatorService/Average"
	CalculatorService_Max_FullMethodName     = "/calculator.CalculatorService/Max"
)

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, CalculatorService_Sum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Prime(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], CalculatorService_Prime_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], CalculatorService_Average_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{ClientStream: stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Max(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_MaxClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], CalculatorService_Max_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceMaxClient{ClientStream: stream}
	return x, nil
}

type CalculatorService_MaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxResponse, error)
	grpc.ClientStream
}

type calculatorServiceMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceMaxClient) Recv() (*MaxResponse, error) {
	m := new(MaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Prime(*PrimeRequest, CalculatorService_PrimeServer) error
	Average(CalculatorService_AverageServer) error
	Max(CalculatorService_MaxServer) error
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) Prime(*PrimeRequest, CalculatorService_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}
func (UnimplementedCalculatorServiceServer) Average(CalculatorService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedCalculatorServiceServer) Max(CalculatorService_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CalculatorService_Sum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Prime(m, &calculatorServicePrimeServer{ServerStream: stream})
}

type CalculatorService_PrimeServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{ServerStream: stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Max(&calculatorServiceMaxServer{ServerStream: stream})
}

type CalculatorService_MaxServer interface {
	Send(*MaxResponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceMaxServer) Send(m *MaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Prime",
			Handler:       _CalculatorService_Prime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _CalculatorService_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}

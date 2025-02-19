// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/data-integration.proto

package proto

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
	DataIngestionService_StreamData_FullMethodName          = "/DataIngestionService/StreamData"
	DataIngestionService_ConfirmDataReceived_FullMethodName = "/DataIngestionService/ConfirmDataReceived"
	DataIngestionService_GetDataStatus_FullMethodName       = "/DataIngestionService/GetDataStatus"
	DataIngestionService_GetDataMetadata_FullMethodName     = "/DataIngestionService/GetDataMetadata"
	DataIngestionService_ReprocessData_FullMethodName       = "/DataIngestionService/ReprocessData"
	DataIngestionService_ReportError_FullMethodName         = "/DataIngestionService/ReportError"
)

// DataIngestionServiceClient is the client API for DataIngestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataIngestionServiceClient interface {
	// Получить данные из источника
	StreamData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error)
	// Запрос на подтверждение получения данных
	ConfirmDataReceived(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
	// Статус получения данных
	GetDataStatus(ctx context.Context, in *DataStatusRequest, opts ...grpc.CallOption) (*DataStatusResponse, error)
	// Получить метаданные о данных
	GetDataMetadata(ctx context.Context, in *DataMetadataRequest, opts ...grpc.CallOption) (*DataMetadataResponse, error)
	// Запуск повторной обработки данных
	ReprocessData(ctx context.Context, in *ReprocessRequest, opts ...grpc.CallOption) (*ReprocessResponse, error)
	// Обработка ошибок
	ReportError(ctx context.Context, in *ErrorReportRequest, opts ...grpc.CallOption) (*ErrorReportResponse, error)
}

type dataIngestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataIngestionServiceClient(cc grpc.ClientConnInterface) DataIngestionServiceClient {
	return &dataIngestionServiceClient{cc}
}

func (c *dataIngestionServiceClient) StreamData(ctx context.Context, in *DataRequest, opts ...grpc.CallOption) (*DataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataResponse)
	err := c.cc.Invoke(ctx, DataIngestionService_StreamData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataIngestionServiceClient) ConfirmDataReceived(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, DataIngestionService_ConfirmDataReceived_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataIngestionServiceClient) GetDataStatus(ctx context.Context, in *DataStatusRequest, opts ...grpc.CallOption) (*DataStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataStatusResponse)
	err := c.cc.Invoke(ctx, DataIngestionService_GetDataStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataIngestionServiceClient) GetDataMetadata(ctx context.Context, in *DataMetadataRequest, opts ...grpc.CallOption) (*DataMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DataMetadataResponse)
	err := c.cc.Invoke(ctx, DataIngestionService_GetDataMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataIngestionServiceClient) ReprocessData(ctx context.Context, in *ReprocessRequest, opts ...grpc.CallOption) (*ReprocessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReprocessResponse)
	err := c.cc.Invoke(ctx, DataIngestionService_ReprocessData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataIngestionServiceClient) ReportError(ctx context.Context, in *ErrorReportRequest, opts ...grpc.CallOption) (*ErrorReportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ErrorReportResponse)
	err := c.cc.Invoke(ctx, DataIngestionService_ReportError_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataIngestionServiceServer is the server API for DataIngestionService service.
// All implementations must embed UnimplementedDataIngestionServiceServer
// for forward compatibility.
type DataIngestionServiceServer interface {
	// Получить данные из источника
	StreamData(context.Context, *DataRequest) (*DataResponse, error)
	// Запрос на подтверждение получения данных
	ConfirmDataReceived(context.Context, *ConfirmRequest) (*ConfirmResponse, error)
	// Статус получения данных
	GetDataStatus(context.Context, *DataStatusRequest) (*DataStatusResponse, error)
	// Получить метаданные о данных
	GetDataMetadata(context.Context, *DataMetadataRequest) (*DataMetadataResponse, error)
	// Запуск повторной обработки данных
	ReprocessData(context.Context, *ReprocessRequest) (*ReprocessResponse, error)
	// Обработка ошибок
	ReportError(context.Context, *ErrorReportRequest) (*ErrorReportResponse, error)
	mustEmbedUnimplementedDataIngestionServiceServer()
}

// UnimplementedDataIngestionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataIngestionServiceServer struct{}

func (UnimplementedDataIngestionServiceServer) StreamData(context.Context, *DataRequest) (*DataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}
func (UnimplementedDataIngestionServiceServer) ConfirmDataReceived(context.Context, *ConfirmRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDataReceived not implemented")
}
func (UnimplementedDataIngestionServiceServer) GetDataStatus(context.Context, *DataStatusRequest) (*DataStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataStatus not implemented")
}
func (UnimplementedDataIngestionServiceServer) GetDataMetadata(context.Context, *DataMetadataRequest) (*DataMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataMetadata not implemented")
}
func (UnimplementedDataIngestionServiceServer) ReprocessData(context.Context, *ReprocessRequest) (*ReprocessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReprocessData not implemented")
}
func (UnimplementedDataIngestionServiceServer) ReportError(context.Context, *ErrorReportRequest) (*ErrorReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportError not implemented")
}
func (UnimplementedDataIngestionServiceServer) mustEmbedUnimplementedDataIngestionServiceServer() {}
func (UnimplementedDataIngestionServiceServer) testEmbeddedByValue()                              {}

// UnsafeDataIngestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataIngestionServiceServer will
// result in compilation errors.
type UnsafeDataIngestionServiceServer interface {
	mustEmbedUnimplementedDataIngestionServiceServer()
}

func RegisterDataIngestionServiceServer(s grpc.ServiceRegistrar, srv DataIngestionServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataIngestionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataIngestionService_ServiceDesc, srv)
}

func _DataIngestionService_StreamData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataIngestionServiceServer).StreamData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataIngestionService_StreamData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataIngestionServiceServer).StreamData(ctx, req.(*DataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataIngestionService_ConfirmDataReceived_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataIngestionServiceServer).ConfirmDataReceived(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataIngestionService_ConfirmDataReceived_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataIngestionServiceServer).ConfirmDataReceived(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataIngestionService_GetDataStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataIngestionServiceServer).GetDataStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataIngestionService_GetDataStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataIngestionServiceServer).GetDataStatus(ctx, req.(*DataStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataIngestionService_GetDataMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataIngestionServiceServer).GetDataMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataIngestionService_GetDataMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataIngestionServiceServer).GetDataMetadata(ctx, req.(*DataMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataIngestionService_ReprocessData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReprocessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataIngestionServiceServer).ReprocessData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataIngestionService_ReprocessData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataIngestionServiceServer).ReprocessData(ctx, req.(*ReprocessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataIngestionService_ReportError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ErrorReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataIngestionServiceServer).ReportError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataIngestionService_ReportError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataIngestionServiceServer).ReportError(ctx, req.(*ErrorReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataIngestionService_ServiceDesc is the grpc.ServiceDesc for DataIngestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataIngestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataIngestionService",
	HandlerType: (*DataIngestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StreamData",
			Handler:    _DataIngestionService_StreamData_Handler,
		},
		{
			MethodName: "ConfirmDataReceived",
			Handler:    _DataIngestionService_ConfirmDataReceived_Handler,
		},
		{
			MethodName: "GetDataStatus",
			Handler:    _DataIngestionService_GetDataStatus_Handler,
		},
		{
			MethodName: "GetDataMetadata",
			Handler:    _DataIngestionService_GetDataMetadata_Handler,
		},
		{
			MethodName: "ReprocessData",
			Handler:    _DataIngestionService_ReprocessData_Handler,
		},
		{
			MethodName: "ReportError",
			Handler:    _DataIngestionService_ReportError_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/data-integration.proto",
}

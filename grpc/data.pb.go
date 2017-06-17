// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	data.proto

It has these top-level messages:
	EmptyRequest
	LivenessResponse
	ExchangeRateResponse
	ExchangeRatesResponse
	WeatherRequest
	WeatherResponse
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// https://github.com/google/protobuf/blob/master/src/google/protobuf/empty.proto
// https://stackoverflow.com/questions/31768665/can-i-define-a-grpc-call-with-a-null-request-or-response
// https://stackoverflow.com/questions/29687243/protobuf-rpc-service-method-without-parameters
// note: The reason this is required is so that if you later need to add an optional parameter,
// you can do so without breaking existing code.
type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LivenessResponse struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *LivenessResponse) Reset()                    { *m = LivenessResponse{} }
func (m *LivenessResponse) String() string            { return proto.CompactTextString(m) }
func (*LivenessResponse) ProtoMessage()               {}
func (*LivenessResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LivenessResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// https://developers.google.com/protocol-buffers/docs/proto3#scalar
type ExchangeRateResponse struct {
	Title       string  `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	PubDate     string  `protobuf:"bytes,2,opt,name=pubDate" json:"pubDate,omitempty"`
	Description float32 `protobuf:"fixed32,3,opt,name=description" json:"description,omitempty"`
	Quant       int32   `protobuf:"varint,4,opt,name=quant" json:"quant,omitempty"`
	Index       string  `protobuf:"bytes,5,opt,name=index" json:"index,omitempty"`
	Change      float32 `protobuf:"fixed32,6,opt,name=change" json:"change,omitempty"`
}

func (m *ExchangeRateResponse) Reset()                    { *m = ExchangeRateResponse{} }
func (m *ExchangeRateResponse) String() string            { return proto.CompactTextString(m) }
func (*ExchangeRateResponse) ProtoMessage()               {}
func (*ExchangeRateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExchangeRateResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ExchangeRateResponse) GetPubDate() string {
	if m != nil {
		return m.PubDate
	}
	return ""
}

func (m *ExchangeRateResponse) GetDescription() float32 {
	if m != nil {
		return m.Description
	}
	return 0
}

func (m *ExchangeRateResponse) GetQuant() int32 {
	if m != nil {
		return m.Quant
	}
	return 0
}

func (m *ExchangeRateResponse) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ExchangeRateResponse) GetChange() float32 {
	if m != nil {
		return m.Change
	}
	return 0
}

type ExchangeRatesResponse struct {
	Data []*ExchangeRateResponse `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *ExchangeRatesResponse) Reset()                    { *m = ExchangeRatesResponse{} }
func (m *ExchangeRatesResponse) String() string            { return proto.CompactTextString(m) }
func (*ExchangeRatesResponse) ProtoMessage()               {}
func (*ExchangeRatesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExchangeRatesResponse) GetData() []*ExchangeRateResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

type WeatherRequest struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *WeatherRequest) Reset()                    { *m = WeatherRequest{} }
func (m *WeatherRequest) String() string            { return proto.CompactTextString(m) }
func (*WeatherRequest) ProtoMessage()               {}
func (*WeatherRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WeatherRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *WeatherRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type WeatherResponse struct {
	Temp     float64 `protobuf:"fixed64,1,opt,name=temp" json:"temp,omitempty"`
	Pressure float64 `protobuf:"fixed64,2,opt,name=pressure" json:"pressure,omitempty"`
	Humidity float64 `protobuf:"fixed64,3,opt,name=humidity" json:"humidity,omitempty"`
}

func (m *WeatherResponse) Reset()                    { *m = WeatherResponse{} }
func (m *WeatherResponse) String() string            { return proto.CompactTextString(m) }
func (*WeatherResponse) ProtoMessage()               {}
func (*WeatherResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WeatherResponse) GetTemp() float64 {
	if m != nil {
		return m.Temp
	}
	return 0
}

func (m *WeatherResponse) GetPressure() float64 {
	if m != nil {
		return m.Pressure
	}
	return 0
}

func (m *WeatherResponse) GetHumidity() float64 {
	if m != nil {
		return m.Humidity
	}
	return 0
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "grpc.EmptyRequest")
	proto.RegisterType((*LivenessResponse)(nil), "grpc.LivenessResponse")
	proto.RegisterType((*ExchangeRateResponse)(nil), "grpc.ExchangeRateResponse")
	proto.RegisterType((*ExchangeRatesResponse)(nil), "grpc.ExchangeRatesResponse")
	proto.RegisterType((*WeatherRequest)(nil), "grpc.WeatherRequest")
	proto.RegisterType((*WeatherResponse)(nil), "grpc.WeatherResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for GrpcService service

type GrpcServiceClient interface {
	// Check grpc server liveness.
	CheckGrpcServerLiveness(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (GrpcService_CheckGrpcServerLivenessClient, error)
	// Get all exchange rates - A server-to-client streaming RPC.
	GetExchangeRates(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (GrpcService_GetExchangeRatesClient, error)
	// Get weather info - A server-to-client streaming RPC.
	GetWeatherInfo(ctx context.Context, in *WeatherRequest, opts ...grpc1.CallOption) (GrpcService_GetWeatherInfoClient, error)
}

type grpcServiceClient struct {
	cc *grpc1.ClientConn
}

func NewGrpcServiceClient(cc *grpc1.ClientConn) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) CheckGrpcServerLiveness(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (GrpcService_CheckGrpcServerLivenessClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_GrpcService_serviceDesc.Streams[0], c.cc, "/grpc.GrpcService/CheckGrpcServerLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcServiceCheckGrpcServerLivenessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcService_CheckGrpcServerLivenessClient interface {
	Recv() (*LivenessResponse, error)
	grpc1.ClientStream
}

type grpcServiceCheckGrpcServerLivenessClient struct {
	grpc1.ClientStream
}

func (x *grpcServiceCheckGrpcServerLivenessClient) Recv() (*LivenessResponse, error) {
	m := new(LivenessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcServiceClient) GetExchangeRates(ctx context.Context, in *EmptyRequest, opts ...grpc1.CallOption) (GrpcService_GetExchangeRatesClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_GrpcService_serviceDesc.Streams[1], c.cc, "/grpc.GrpcService/GetExchangeRates", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcServiceGetExchangeRatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcService_GetExchangeRatesClient interface {
	Recv() (*ExchangeRatesResponse, error)
	grpc1.ClientStream
}

type grpcServiceGetExchangeRatesClient struct {
	grpc1.ClientStream
}

func (x *grpcServiceGetExchangeRatesClient) Recv() (*ExchangeRatesResponse, error) {
	m := new(ExchangeRatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcServiceClient) GetWeatherInfo(ctx context.Context, in *WeatherRequest, opts ...grpc1.CallOption) (GrpcService_GetWeatherInfoClient, error) {
	stream, err := grpc1.NewClientStream(ctx, &_GrpcService_serviceDesc.Streams[2], c.cc, "/grpc.GrpcService/GetWeatherInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcServiceGetWeatherInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcService_GetWeatherInfoClient interface {
	Recv() (*WeatherResponse, error)
	grpc1.ClientStream
}

type grpcServiceGetWeatherInfoClient struct {
	grpc1.ClientStream
}

func (x *grpcServiceGetWeatherInfoClient) Recv() (*WeatherResponse, error) {
	m := new(WeatherResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GrpcService service

type GrpcServiceServer interface {
	// Check grpc server liveness.
	CheckGrpcServerLiveness(*EmptyRequest, GrpcService_CheckGrpcServerLivenessServer) error
	// Get all exchange rates - A server-to-client streaming RPC.
	GetExchangeRates(*EmptyRequest, GrpcService_GetExchangeRatesServer) error
	// Get weather info - A server-to-client streaming RPC.
	GetWeatherInfo(*WeatherRequest, GrpcService_GetWeatherInfoServer) error
}

func RegisterGrpcServiceServer(s *grpc1.Server, srv GrpcServiceServer) {
	s.RegisterService(&_GrpcService_serviceDesc, srv)
}

func _GrpcService_CheckGrpcServerLiveness_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcServiceServer).CheckGrpcServerLiveness(m, &grpcServiceCheckGrpcServerLivenessServer{stream})
}

type GrpcService_CheckGrpcServerLivenessServer interface {
	Send(*LivenessResponse) error
	grpc1.ServerStream
}

type grpcServiceCheckGrpcServerLivenessServer struct {
	grpc1.ServerStream
}

func (x *grpcServiceCheckGrpcServerLivenessServer) Send(m *LivenessResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcService_GetExchangeRates_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(EmptyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcServiceServer).GetExchangeRates(m, &grpcServiceGetExchangeRatesServer{stream})
}

type GrpcService_GetExchangeRatesServer interface {
	Send(*ExchangeRatesResponse) error
	grpc1.ServerStream
}

type grpcServiceGetExchangeRatesServer struct {
	grpc1.ServerStream
}

func (x *grpcServiceGetExchangeRatesServer) Send(m *ExchangeRatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcService_GetWeatherInfo_Handler(srv interface{}, stream grpc1.ServerStream) error {
	m := new(WeatherRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcServiceServer).GetWeatherInfo(m, &grpcServiceGetWeatherInfoServer{stream})
}

type GrpcService_GetWeatherInfoServer interface {
	Send(*WeatherResponse) error
	grpc1.ServerStream
}

type grpcServiceGetWeatherInfoServer struct {
	grpc1.ServerStream
}

func (x *grpcServiceGetWeatherInfoServer) Send(m *WeatherResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GrpcService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.GrpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods:     []grpc1.MethodDesc{},
	Streams: []grpc1.StreamDesc{
		{
			StreamName:    "CheckGrpcServerLiveness",
			Handler:       _GrpcService_CheckGrpcServerLiveness_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetExchangeRates",
			Handler:       _GrpcService_GetExchangeRates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetWeatherInfo",
			Handler:       _GrpcService_GetWeatherInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "data.proto",
}

func init() { proto.RegisterFile("data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x65, 0xf3, 0x05, 0x9d, 0xa0, 0x10, 0xad, 0xd2, 0xb2, 0x0a, 0x1c, 0x2c, 0x8b, 0x43, 0x4e,
	0x11, 0x2a, 0xbf, 0x00, 0x41, 0x65, 0x15, 0x71, 0x5a, 0x0e, 0x9c, 0x38, 0x6c, 0xed, 0x21, 0x5e,
	0x11, 0xaf, 0xb7, 0xbb, 0xe3, 0xaa, 0xfd, 0x4f, 0xfc, 0x2c, 0x7e, 0x08, 0xf2, 0xae, 0xed, 0x24,
	0x6d, 0x6f, 0xfb, 0xde, 0x8c, 0xdf, 0xcc, 0x7b, 0x1e, 0x80, 0x42, 0x91, 0xda, 0x5a, 0x57, 0x53,
	0xcd, 0x27, 0x3b, 0x67, 0xf3, 0x74, 0x01, 0xaf, 0xaf, 0x2a, 0x4b, 0x0f, 0x12, 0x6f, 0x1b, 0xf4,
	0x94, 0x7e, 0x80, 0xe5, 0x77, 0x7d, 0x87, 0x06, 0xbd, 0x97, 0xe8, 0x6d, 0x6d, 0x3c, 0xf2, 0x25,
	0x8c, 0x2b, 0xbf, 0x13, 0x2c, 0x61, 0x9b, 0x33, 0xd9, 0x3e, 0xd3, 0xbf, 0x0c, 0x56, 0x57, 0xf7,
	0x79, 0xa9, 0xcc, 0x0e, 0xa5, 0x22, 0x1c, 0x5a, 0x57, 0x30, 0x25, 0x4d, 0x7b, 0xec, 0x9a, 0x23,
	0xe0, 0x02, 0x5e, 0xda, 0xe6, 0xe6, 0xab, 0x22, 0x14, 0xa3, 0xc0, 0xf7, 0x90, 0x27, 0x30, 0x2f,
	0xd0, 0xe7, 0x4e, 0x5b, 0xd2, 0xb5, 0x11, 0xe3, 0x84, 0x6d, 0x46, 0xf2, 0x98, 0x6a, 0x15, 0x6f,
	0x1b, 0x65, 0x48, 0x4c, 0x12, 0xb6, 0x99, 0xca, 0x08, 0x5a, 0x56, 0x9b, 0x02, 0xef, 0xc5, 0x34,
	0xce, 0x09, 0x80, 0x5f, 0xc0, 0x2c, 0xee, 0x24, 0x66, 0x41, 0xa8, 0x43, 0x69, 0x06, 0xe7, 0xc7,
	0xdb, 0x1e, 0x9c, 0x6d, 0x61, 0xd2, 0x26, 0x22, 0x58, 0x32, 0xde, 0xcc, 0x2f, 0xd7, 0xdb, 0x36,
	0x92, 0xed, 0x73, 0xc6, 0x64, 0xe8, 0x4b, 0xbf, 0xc1, 0xe2, 0x27, 0x2a, 0x2a, 0xd1, 0x75, 0x79,
	0xf1, 0x35, 0xbc, 0xda, 0x2b, 0xd2, 0xd4, 0x14, 0xd1, 0x33, 0x93, 0x03, 0xe6, 0xef, 0xe1, 0x6c,
	0x5f, 0x9b, 0x5d, 0x2c, 0x8e, 0x42, 0xf1, 0x40, 0xa4, 0xbf, 0xe0, 0xcd, 0xa0, 0xd5, 0xad, 0xc3,
	0x61, 0x42, 0x58, 0xd9, 0x4e, 0x28, 0xbc, 0xdb, 0x01, 0xd6, 0xa1, 0xf7, 0x8d, 0xeb, 0x35, 0x06,
	0xdc, 0xd6, 0xca, 0xa6, 0xd2, 0x85, 0xa6, 0x87, 0x10, 0x1d, 0x93, 0x03, 0xbe, 0xfc, 0xc7, 0x60,
	0x9e, 0x39, 0x9b, 0xff, 0x40, 0x77, 0xa7, 0x73, 0xe4, 0xd7, 0xf0, 0xf6, 0x4b, 0x89, 0xf9, 0x9f,
	0x9e, 0x43, 0xd7, 0xff, 0x67, 0xce, 0x3b, 0xdf, 0x47, 0x77, 0xb0, 0xbe, 0x88, 0xdc, 0xe3, 0x5b,
	0x48, 0x5f, 0x7c, 0x64, 0x3c, 0x83, 0x65, 0x86, 0x74, 0x92, 0xe8, 0xb3, 0x1a, 0xef, 0x9e, 0xe6,
	0x79, 0x2a, 0xf4, 0x19, 0x16, 0x19, 0x52, 0x97, 0xc2, 0xb5, 0xf9, 0x5d, 0xf3, 0x55, 0xfc, 0xe4,
	0x34, 0xe4, 0xf5, 0xf9, 0x23, 0xf6, 0x20, 0x71, 0x33, 0x0b, 0xc7, 0xfc, 0xe9, 0x7f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf7, 0xfe, 0x52, 0xa5, 0xda, 0x02, 0x00, 0x00,
}

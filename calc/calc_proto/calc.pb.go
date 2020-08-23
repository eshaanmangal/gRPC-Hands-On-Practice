// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        (unknown)
// source: calc/calc_proto/calc.proto

package calc_proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  float32 `protobuf:"fixed32,1,opt,name=firstNumber,proto3" json:"firstNumber,omitempty"`
	SecondNumber float32 `protobuf:"fixed32,2,opt,name=secondNumber,proto3" json:"secondNumber,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_calc_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetFirstNumber() float32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *SumRequest) GetSecondNumber() float32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_calc_proto_rawDescGZIP(), []int{1}
}

func (x *SumResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type TwoNumbers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  float32 `protobuf:"fixed32,1,opt,name=firstNumber,proto3" json:"firstNumber,omitempty"`
	SecondNumber float32 `protobuf:"fixed32,2,opt,name=secondNumber,proto3" json:"secondNumber,omitempty"`
}

func (x *TwoNumbers) Reset() {
	*x = TwoNumbers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwoNumbers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwoNumbers) ProtoMessage() {}

func (x *TwoNumbers) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwoNumbers.ProtoReflect.Descriptor instead.
func (*TwoNumbers) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_calc_proto_rawDescGZIP(), []int{2}
}

func (x *TwoNumbers) GetFirstNumber() float32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *TwoNumbers) GetSecondNumber() float32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type AirthmeticOperationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AirthmeticOperationsResponse) Reset() {
	*x = AirthmeticOperationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calc_calc_proto_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirthmeticOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirthmeticOperationsResponse) ProtoMessage() {}

func (x *AirthmeticOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calc_calc_proto_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirthmeticOperationsResponse.ProtoReflect.Descriptor instead.
func (*AirthmeticOperationsResponse) Descriptor() ([]byte, []int) {
	return file_calc_calc_proto_calc_proto_rawDescGZIP(), []int{3}
}

func (x *AirthmeticOperationsResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calc_calc_proto_calc_proto protoreflect.FileDescriptor

var file_calc_calc_proto_calc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61,
	0x6c, 0x63, 0x22, 0x52, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x52, 0x0a,
	0x0a, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x36, 0x0a, 0x1c, 0x41, 0x69, 0x72, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x98, 0x01, 0x0a, 0x0a, 0x53, 0x75,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x4f,
	0x66, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x50, 0x0a, 0x14, 0x41, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x2e, 0x54, 0x77, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x1a, 0x22, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x41, 0x69, 0x72, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x63, 0x61, 0x6c, 0x63, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calc_calc_proto_calc_proto_rawDescOnce sync.Once
	file_calc_calc_proto_calc_proto_rawDescData = file_calc_calc_proto_calc_proto_rawDesc
)

func file_calc_calc_proto_calc_proto_rawDescGZIP() []byte {
	file_calc_calc_proto_calc_proto_rawDescOnce.Do(func() {
		file_calc_calc_proto_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calc_calc_proto_calc_proto_rawDescData)
	})
	return file_calc_calc_proto_calc_proto_rawDescData
}

var file_calc_calc_proto_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calc_calc_proto_calc_proto_goTypes = []interface{}{
	(*SumRequest)(nil),                   // 0: calc.SumRequest
	(*SumResponse)(nil),                  // 1: calc.SumResponse
	(*TwoNumbers)(nil),                   // 2: calc.TwoNumbers
	(*AirthmeticOperationsResponse)(nil), // 3: calc.AirthmeticOperationsResponse
}
var file_calc_calc_proto_calc_proto_depIdxs = []int32{
	0, // 0: calc.SumService.SumOfTwoNumbers:input_type -> calc.SumRequest
	2, // 1: calc.SumService.ArithmeticOperations:input_type -> calc.TwoNumbers
	1, // 2: calc.SumService.SumOfTwoNumbers:output_type -> calc.SumResponse
	3, // 3: calc.SumService.ArithmeticOperations:output_type -> calc.AirthmeticOperationsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calc_calc_proto_calc_proto_init() }
func file_calc_calc_proto_calc_proto_init() {
	if File_calc_calc_proto_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calc_calc_proto_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calc_calc_proto_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calc_calc_proto_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwoNumbers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calc_calc_proto_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirthmeticOperationsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calc_calc_proto_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calc_calc_proto_calc_proto_goTypes,
		DependencyIndexes: file_calc_calc_proto_calc_proto_depIdxs,
		MessageInfos:      file_calc_calc_proto_calc_proto_msgTypes,
	}.Build()
	File_calc_calc_proto_calc_proto = out.File
	file_calc_calc_proto_calc_proto_rawDesc = nil
	file_calc_calc_proto_calc_proto_goTypes = nil
	file_calc_calc_proto_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	SumOfTwoNumbers(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	ArithmeticOperations(ctx context.Context, in *TwoNumbers, opts ...grpc.CallOption) (SumService_ArithmeticOperationsClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) SumOfTwoNumbers(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calc.SumService/SumOfTwoNumbers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) ArithmeticOperations(ctx context.Context, in *TwoNumbers, opts ...grpc.CallOption) (SumService_ArithmeticOperationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/calc.SumService/ArithmeticOperations", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceArithmeticOperationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_ArithmeticOperationsClient interface {
	Recv() (*AirthmeticOperationsResponse, error)
	grpc.ClientStream
}

type sumServiceArithmeticOperationsClient struct {
	grpc.ClientStream
}

func (x *sumServiceArithmeticOperationsClient) Recv() (*AirthmeticOperationsResponse, error) {
	m := new(AirthmeticOperationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	SumOfTwoNumbers(context.Context, *SumRequest) (*SumResponse, error)
	ArithmeticOperations(*TwoNumbers, SumService_ArithmeticOperationsServer) error
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) SumOfTwoNumbers(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumOfTwoNumbers not implemented")
}
func (*UnimplementedSumServiceServer) ArithmeticOperations(*TwoNumbers, SumService_ArithmeticOperationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ArithmeticOperations not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_SumOfTwoNumbers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).SumOfTwoNumbers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.SumService/SumOfTwoNumbers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).SumOfTwoNumbers(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_ArithmeticOperations_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TwoNumbers)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).ArithmeticOperations(m, &sumServiceArithmeticOperationsServer{stream})
}

type SumService_ArithmeticOperationsServer interface {
	Send(*AirthmeticOperationsResponse) error
	grpc.ServerStream
}

type sumServiceArithmeticOperationsServer struct {
	grpc.ServerStream
}

func (x *sumServiceArithmeticOperationsServer) Send(m *AirthmeticOperationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SumOfTwoNumbers",
			Handler:    _SumService_SumOfTwoNumbers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ArithmeticOperations",
			Handler:       _SumService_ArithmeticOperations_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calc/calc_proto/calc.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MathRequest struct {
	Operand1             int64    `protobuf:"varint,1,opt,name=operand1,proto3" json:"operand1,omitempty"`
	Operand2             int64    `protobuf:"varint,2,opt,name=operand2,proto3" json:"operand2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathRequest) Reset()         { *m = MathRequest{} }
func (m *MathRequest) String() string { return proto.CompactTextString(m) }
func (*MathRequest) ProtoMessage()    {}
func (*MathRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *MathRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathRequest.Unmarshal(m, b)
}
func (m *MathRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathRequest.Marshal(b, m, deterministic)
}
func (m *MathRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathRequest.Merge(m, src)
}
func (m *MathRequest) XXX_Size() int {
	return xxx_messageInfo_MathRequest.Size(m)
}
func (m *MathRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MathRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MathRequest proto.InternalMessageInfo

func (m *MathRequest) GetOperand1() int64 {
	if m != nil {
		return m.Operand1
	}
	return 0
}

func (m *MathRequest) GetOperand2() int64 {
	if m != nil {
		return m.Operand2
	}
	return 0
}

type MathResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MathResponse) Reset()         { *m = MathResponse{} }
func (m *MathResponse) String() string { return proto.CompactTextString(m) }
func (*MathResponse) ProtoMessage()    {}
func (*MathResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *MathResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MathResponse.Unmarshal(m, b)
}
func (m *MathResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MathResponse.Marshal(b, m, deterministic)
}
func (m *MathResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MathResponse.Merge(m, src)
}
func (m *MathResponse) XXX_Size() int {
	return xxx_messageInfo_MathResponse.Size(m)
}
func (m *MathResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MathResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MathResponse proto.InternalMessageInfo

func (m *MathResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type RandomRequest struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Bounded              bool     `protobuf:"varint,2,opt,name=bounded,proto3" json:"bounded,omitempty"`
	MinValue             int64    `protobuf:"varint,3,opt,name=minValue,proto3" json:"minValue,omitempty"`
	MaxValue             int64    `protobuf:"varint,4,opt,name=maxValue,proto3" json:"maxValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomRequest) Reset()         { *m = RandomRequest{} }
func (m *RandomRequest) String() string { return proto.CompactTextString(m) }
func (*RandomRequest) ProtoMessage()    {}
func (*RandomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *RandomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomRequest.Unmarshal(m, b)
}
func (m *RandomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomRequest.Marshal(b, m, deterministic)
}
func (m *RandomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomRequest.Merge(m, src)
}
func (m *RandomRequest) XXX_Size() int {
	return xxx_messageInfo_RandomRequest.Size(m)
}
func (m *RandomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RandomRequest proto.InternalMessageInfo

func (m *RandomRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *RandomRequest) GetBounded() bool {
	if m != nil {
		return m.Bounded
	}
	return false
}

func (m *RandomRequest) GetMinValue() int64 {
	if m != nil {
		return m.MinValue
	}
	return 0
}

func (m *RandomRequest) GetMaxValue() int64 {
	if m != nil {
		return m.MaxValue
	}
	return 0
}

type RandomResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RandomResponse) Reset()         { *m = RandomResponse{} }
func (m *RandomResponse) String() string { return proto.CompactTextString(m) }
func (*RandomResponse) ProtoMessage()    {}
func (*RandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *RandomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RandomResponse.Unmarshal(m, b)
}
func (m *RandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RandomResponse.Marshal(b, m, deterministic)
}
func (m *RandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RandomResponse.Merge(m, src)
}
func (m *RandomResponse) XXX_Size() int {
	return xxx_messageInfo_RandomResponse.Size(m)
}
func (m *RandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RandomResponse proto.InternalMessageInfo

func (m *RandomResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*MathRequest)(nil), "model.MathRequest")
	proto.RegisterType((*MathResponse)(nil), "model.MathResponse")
	proto.RegisterType((*RandomRequest)(nil), "model.RandomRequest")
	proto.RegisterType((*RandomResponse)(nil), "model.RandomResponse")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x4f, 0x84, 0x40,
	0x10, 0xc5, 0x83, 0x08, 0x5e, 0xe6, 0xbc, 0x2b, 0x56, 0x34, 0x84, 0xca, 0x50, 0x5c, 0xac, 0x88,
	0x62, 0x61, 0x6d, 0x72, 0xda, 0xd1, 0x70, 0x89, 0xfd, 0x72, 0x3b, 0x51, 0x12, 0xd8, 0x45, 0x96,
	0x45, 0xfd, 0xf4, 0x9a, 0xdb, 0x3f, 0xe6, 0xb0, 0x92, 0xf2, 0xb7, 0xef, 0xe5, 0xcd, 0x9b, 0x01,
	0x58, 0xb7, 0x28, 0x25, 0x7d, 0x45, 0x99, 0x75, 0xbd, 0x18, 0x04, 0x09, 0x5a, 0xc1, 0xb0, 0x49,
	0x9f, 0x60, 0x59, 0xd0, 0xe1, 0xad, 0xc4, 0x77, 0x85, 0x72, 0x20, 0x09, 0x2c, 0x44, 0x87, 0x3d,
	0xe5, 0xec, 0x2e, 0xf6, 0xae, 0xbd, 0x1b, 0xbf, 0xfc, 0xe5, 0x23, 0x2d, 0x8f, 0x4f, 0x26, 0x5a,
	0x9e, 0x6e, 0xe0, 0xdc, 0xc4, 0xc8, 0x4e, 0x70, 0x89, 0xe4, 0x0a, 0xc2, 0x1e, 0xa5, 0x6a, 0x06,
	0x9b, 0x62, 0x29, 0xfd, 0x80, 0x55, 0x49, 0x39, 0x13, 0xad, 0x1b, 0x18, 0x41, 0xb0, 0x17, 0x8a,
	0x1b, 0x5f, 0x50, 0x1a, 0x20, 0x31, 0x9c, 0x55, 0x42, 0x71, 0x86, 0x4c, 0x4f, 0x5a, 0x94, 0x0e,
	0x0f, 0x25, 0xda, 0x9a, 0xbf, 0xd0, 0x46, 0x61, 0xec, 0x9b, 0x12, 0x8e, 0xb5, 0x46, 0x3f, 0x8d,
	0x76, 0x6a, 0x35, 0xcb, 0xe9, 0x06, 0xd6, 0x6e, 0xb0, 0xad, 0x18, 0x41, 0x30, 0x6a, 0xab, 0x69,
	0x68, 0x20, 0xff, 0xf6, 0x60, 0x55, 0x7c, 0x1d, 0x76, 0xd9, 0x61, 0x3f, 0xd6, 0x7b, 0x24, 0x19,
	0xf8, 0x8f, 0x8c, 0x11, 0x92, 0xe9, 0x83, 0x65, 0x47, 0xd7, 0x4a, 0x2e, 0x26, 0x6f, 0x36, 0x37,
	0x03, 0x7f, 0xa7, 0xaa, 0x59, 0xfe, 0x42, 0x35, 0xb3, 0xfc, 0xdb, 0x7a, 0x9c, 0x97, 0x2f, 0xfe,
	0xdf, 0x3f, 0x7f, 0x86, 0xe5, 0x96, 0x0e, 0xd4, 0xad, 0xff, 0x00, 0xa1, 0x39, 0x1c, 0x89, 0xac,
	0x7b, 0xf2, 0x01, 0x93, 0xcb, 0x3f, 0xaf, 0x26, 0xe5, 0xd6, 0xab, 0x42, 0xfd, 0x9f, 0xdd, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x27, 0x7c, 0xeb, 0x83, 0x79, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MyMathServiceClient is the client API for MyMathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyMathServiceClient interface {
	Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Sub(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Mul(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Div(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Mod(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
}

type myMathServiceClient struct {
	cc *grpc.ClientConn
}

func NewMyMathServiceClient(cc *grpc.ClientConn) MyMathServiceClient {
	return &myMathServiceClient{cc}
}

func (c *myMathServiceClient) Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MyMathService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMathServiceClient) Sub(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MyMathService/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMathServiceClient) Mul(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MyMathService/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMathServiceClient) Div(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MyMathService/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myMathServiceClient) Mod(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/model.MyMathService/Mod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyMathServiceServer is the server API for MyMathService service.
type MyMathServiceServer interface {
	Add(context.Context, *MathRequest) (*MathResponse, error)
	Sub(context.Context, *MathRequest) (*MathResponse, error)
	Mul(context.Context, *MathRequest) (*MathResponse, error)
	Div(context.Context, *MathRequest) (*MathResponse, error)
	Mod(context.Context, *MathRequest) (*MathResponse, error)
}

// UnimplementedMyMathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMyMathServiceServer struct {
}

func (*UnimplementedMyMathServiceServer) Add(ctx context.Context, req *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedMyMathServiceServer) Sub(ctx context.Context, req *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (*UnimplementedMyMathServiceServer) Mul(ctx context.Context, req *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (*UnimplementedMyMathServiceServer) Div(ctx context.Context, req *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}
func (*UnimplementedMyMathServiceServer) Mod(ctx context.Context, req *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mod not implemented")
}

func RegisterMyMathServiceServer(s *grpc.Server, srv MyMathServiceServer) {
	s.RegisterService(&_MyMathService_serviceDesc, srv)
}

func _MyMathService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMathServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MyMathService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMathServiceServer).Add(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMathService_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMathServiceServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MyMathService/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMathServiceServer).Sub(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMathService_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMathServiceServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MyMathService/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMathServiceServer).Mul(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMathService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMathServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MyMathService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMathServiceServer).Div(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyMathService_Mod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyMathServiceServer).Mod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.MyMathService/Mod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyMathServiceServer).Mod(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyMathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.MyMathService",
	HandlerType: (*MyMathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MyMathService_Add_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _MyMathService_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _MyMathService_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _MyMathService_Div_Handler,
		},
		{
			MethodName: "Mod",
			Handler:    _MyMathService_Mod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataServiceClient interface {
	Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (DataService_RandomClient, error)
}

type dataServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataServiceClient(cc *grpc.ClientConn) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (DataService_RandomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataService_serviceDesc.Streams[0], "/model.DataService/Random", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceRandomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_RandomClient interface {
	Recv() (*RandomResponse, error)
	grpc.ClientStream
}

type dataServiceRandomClient struct {
	grpc.ClientStream
}

func (x *dataServiceRandomClient) Recv() (*RandomResponse, error) {
	m := new(RandomResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
type DataServiceServer interface {
	Random(*RandomRequest, DataService_RandomServer) error
}

// UnimplementedDataServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (*UnimplementedDataServiceServer) Random(req *RandomRequest, srv DataService_RandomServer) error {
	return status.Errorf(codes.Unimplemented, "method Random not implemented")
}

func RegisterDataServiceServer(s *grpc.Server, srv DataServiceServer) {
	s.RegisterService(&_DataService_serviceDesc, srv)
}

func _DataService_Random_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RandomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).Random(m, &dataServiceRandomServer{stream})
}

type DataService_RandomServer interface {
	Send(*RandomResponse) error
	grpc.ServerStream
}

type dataServiceRandomServer struct {
	grpc.ServerStream
}

func (x *dataServiceRandomServer) Send(m *RandomResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _DataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Random",
			Handler:       _DataService_Random_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messages.proto",
}

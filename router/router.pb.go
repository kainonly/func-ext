// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package funcext

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

type ExportToExcelParameter struct {
	Sheets               []*Sheet `protobuf:"bytes,2,rep,name=sheets,proto3" json:"sheets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToExcelParameter) Reset()         { *m = ExportToExcelParameter{} }
func (m *ExportToExcelParameter) String() string { return proto.CompactTextString(m) }
func (*ExportToExcelParameter) ProtoMessage()    {}
func (*ExportToExcelParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *ExportToExcelParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToExcelParameter.Unmarshal(m, b)
}
func (m *ExportToExcelParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToExcelParameter.Marshal(b, m, deterministic)
}
func (m *ExportToExcelParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToExcelParameter.Merge(m, src)
}
func (m *ExportToExcelParameter) XXX_Size() int {
	return xxx_messageInfo_ExportToExcelParameter.Size(m)
}
func (m *ExportToExcelParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToExcelParameter.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToExcelParameter proto.InternalMessageInfo

func (m *ExportToExcelParameter) GetSheets() []*Sheet {
	if m != nil {
		return m.Sheets
	}
	return nil
}

type Sheet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rows                 []*Row   `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sheet) Reset()         { *m = Sheet{} }
func (m *Sheet) String() string { return proto.CompactTextString(m) }
func (*Sheet) ProtoMessage()    {}
func (*Sheet) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *Sheet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sheet.Unmarshal(m, b)
}
func (m *Sheet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sheet.Marshal(b, m, deterministic)
}
func (m *Sheet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sheet.Merge(m, src)
}
func (m *Sheet) XXX_Size() int {
	return xxx_messageInfo_Sheet.Size(m)
}
func (m *Sheet) XXX_DiscardUnknown() {
	xxx_messageInfo_Sheet.DiscardUnknown(m)
}

var xxx_messageInfo_Sheet proto.InternalMessageInfo

func (m *Sheet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sheet) GetRows() []*Row {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Row struct {
	Axis                 string   `protobuf:"bytes,1,opt,name=axis,proto3" json:"axis,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Row) Reset()         { *m = Row{} }
func (m *Row) String() string { return proto.CompactTextString(m) }
func (*Row) ProtoMessage()    {}
func (*Row) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *Row) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Row.Unmarshal(m, b)
}
func (m *Row) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Row.Marshal(b, m, deterministic)
}
func (m *Row) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Row.Merge(m, src)
}
func (m *Row) XXX_Size() int {
	return xxx_messageInfo_Row.Size(m)
}
func (m *Row) XXX_DiscardUnknown() {
	xxx_messageInfo_Row.DiscardUnknown(m)
}

var xxx_messageInfo_Row proto.InternalMessageInfo

func (m *Row) GetAxis() string {
	if m != nil {
		return m.Axis
	}
	return ""
}

func (m *Row) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ExportToExcelResponse struct {
	Error                uint32   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToExcelResponse) Reset()         { *m = ExportToExcelResponse{} }
func (m *ExportToExcelResponse) String() string { return proto.CompactTextString(m) }
func (*ExportToExcelResponse) ProtoMessage()    {}
func (*ExportToExcelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *ExportToExcelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToExcelResponse.Unmarshal(m, b)
}
func (m *ExportToExcelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToExcelResponse.Marshal(b, m, deterministic)
}
func (m *ExportToExcelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToExcelResponse.Merge(m, src)
}
func (m *ExportToExcelResponse) XXX_Size() int {
	return xxx_messageInfo_ExportToExcelResponse.Size(m)
}
func (m *ExportToExcelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToExcelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToExcelResponse proto.InternalMessageInfo

func (m *ExportToExcelResponse) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *ExportToExcelResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ExportToExcelResponseData struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportToExcelResponseData) Reset()         { *m = ExportToExcelResponseData{} }
func (m *ExportToExcelResponseData) String() string { return proto.CompactTextString(m) }
func (*ExportToExcelResponseData) ProtoMessage()    {}
func (*ExportToExcelResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3, 0}
}

func (m *ExportToExcelResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportToExcelResponseData.Unmarshal(m, b)
}
func (m *ExportToExcelResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportToExcelResponseData.Marshal(b, m, deterministic)
}
func (m *ExportToExcelResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportToExcelResponseData.Merge(m, src)
}
func (m *ExportToExcelResponseData) XXX_Size() int {
	return xxx_messageInfo_ExportToExcelResponseData.Size(m)
}
func (m *ExportToExcelResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportToExcelResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_ExportToExcelResponseData proto.InternalMessageInfo

func (m *ExportToExcelResponseData) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*ExportToExcelParameter)(nil), "funcext.ExportToExcelParameter")
	proto.RegisterType((*Sheet)(nil), "funcext.Sheet")
	proto.RegisterType((*Row)(nil), "funcext.Row")
	proto.RegisterType((*ExportToExcelResponse)(nil), "funcext.ExportToExcelResponse")
	proto.RegisterType((*ExportToExcelResponseData)(nil), "funcext.ExportToExcelResponse.data")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xb5, 0x4d, 0x1a, 0x71, 0x6c, 0x45, 0x16, 0x95, 0xd0, 0x83, 0x86, 0x3d, 0x48, 0x4f, 0x11,
	0xea, 0x59, 0xf0, 0xd2, 0xbb, 0x8c, 0x5e, 0x04, 0x2f, 0x6b, 0x1d, 0x3f, 0x20, 0xc9, 0x84, 0xd9,
	0x8d, 0xc9, 0xcf, 0x2f, 0xbb, 0x4d, 0x0a, 0x85, 0xdc, 0xde, 0xcc, 0x9b, 0xf7, 0xe6, 0xf1, 0x60,
	0x2e, 0xdc, 0x38, 0x92, 0xbc, 0x16, 0x76, 0xac, 0x4e, 0xbf, 0x9b, 0x6a, 0x4b, 0x9d, 0xd3, 0xcf,
	0x70, 0xb3, 0xe9, 0x6a, 0x16, 0xf7, 0xc6, 0x9b, 0x6e, 0x4b, 0xc5, 0x8b, 0x11, 0x53, 0x92, 0x23,
	0x51, 0xf7, 0x90, 0xd8, 0x5f, 0x22, 0x67, 0xd3, 0x69, 0x16, 0xad, 0xce, 0xd7, 0x17, 0x79, 0xaf,
	0xc9, 0x5f, 0xfd, 0x1a, 0x7b, 0x56, 0x3f, 0xc1, 0x2c, 0x2c, 0x94, 0x82, 0xb8, 0x32, 0x25, 0xa5,
	0x93, 0x6c, 0xb2, 0x3a, 0xc3, 0x80, 0x55, 0x06, 0xb1, 0x70, 0x3b, 0x58, 0xcc, 0x0f, 0x16, 0xc8,
	0x2d, 0x06, 0x46, 0x3f, 0x40, 0x84, 0xdc, 0x7a, 0xb1, 0xe9, 0xfe, 0xec, 0x20, 0xf6, 0x58, 0x5d,
	0xc1, 0xec, 0xdf, 0x14, 0x0d, 0xa5, 0xd3, 0xb0, 0xdc, 0x0f, 0xfa, 0x1d, 0xae, 0x8f, 0x12, 0x23,
	0xd9, 0x9a, 0x2b, 0x4b, 0xfe, 0x9c, 0x44, 0x58, 0x82, 0xc7, 0x02, 0xf7, 0x83, 0xba, 0x84, 0xa8,
	0xb4, 0x3f, 0xbd, 0x85, 0x87, 0xcb, 0x14, 0xe2, 0x2f, 0xe3, 0x8c, 0x67, 0x1a, 0x29, 0xfa, 0x8f,
	0x1e, 0xae, 0x3f, 0x20, 0xc1, 0xd0, 0x92, 0x42, 0x58, 0x1c, 0x3d, 0x51, 0x77, 0x87, 0xe8, 0xe3,
	0x75, 0x2d, 0x6f, 0xc7, 0x0f, 0x86, 0x74, 0xfa, 0xe4, 0x33, 0x09, 0xd5, 0x3f, 0xee, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xb0, 0x81, 0x85, 0x8a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	ExportToExcel(ctx context.Context, in *ExportToExcelParameter, opts ...grpc.CallOption) (*ExportToExcelResponse, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) ExportToExcel(ctx context.Context, in *ExportToExcelParameter, opts ...grpc.CallOption) (*ExportToExcelResponse, error) {
	out := new(ExportToExcelResponse)
	err := c.cc.Invoke(ctx, "/funcext.Router/ExportToExcel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	ExportToExcel(context.Context, *ExportToExcelParameter) (*ExportToExcelResponse, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) ExportToExcel(ctx context.Context, req *ExportToExcelParameter) (*ExportToExcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportToExcel not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_ExportToExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportToExcelParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).ExportToExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funcext.Router/ExportToExcel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).ExportToExcel(ctx, req.(*ExportToExcelParameter))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "funcext.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExportToExcel",
			Handler:    _Router_ExportToExcel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}

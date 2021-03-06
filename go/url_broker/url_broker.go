// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/url_broker.proto

package url_broker

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

//URL parameter (string format)
type URL struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URL) Reset()         { *m = URL{} }
func (m *URL) String() string { return proto.CompactTextString(m) }
func (*URL) ProtoMessage()    {}
func (*URL) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fe0cfe562ebbd3, []int{0}
}

func (m *URL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URL.Unmarshal(m, b)
}
func (m *URL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URL.Marshal(b, m, deterministic)
}
func (m *URL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URL.Merge(m, src)
}
func (m *URL) XXX_Size() int {
	return xxx_messageInfo_URL.Size(m)
}
func (m *URL) XXX_DiscardUnknown() {
	xxx_messageInfo_URL.DiscardUnknown(m)
}

var xxx_messageInfo_URL proto.InternalMessageInfo

func (m *URL) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

//Broker response code to PushURL request
type Response struct {
	ResponseCode         int32    `protobuf:"varint,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_64fe0cfe562ebbd3, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return 0
}

func init() {
	proto.RegisterType((*URL)(nil), "URL")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("protos/url_broker.proto", fileDescriptor_64fe0cfe562ebbd3)
}

var fileDescriptor_64fe0cfe562ebbd3 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x2d, 0xca, 0x89, 0x4f, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x03, 0x8b,
	0x28, 0x89, 0x73, 0x31, 0x87, 0x06, 0xf9, 0x08, 0x09, 0x70, 0x31, 0x97, 0x16, 0xe5, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0xfa, 0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9,
	0x79, 0xc5, 0xa9, 0x42, 0xca, 0x5c, 0xbc, 0x45, 0x50, 0x76, 0x7c, 0x72, 0x7e, 0x4a, 0x2a, 0x58,
	0x1d, 0x6b, 0x10, 0x0f, 0x4c, 0xd0, 0x39, 0x3f, 0x25, 0xd5, 0xc8, 0x80, 0x4b, 0x20, 0x34, 0xc8,
	0xc7, 0x09, 0x6c, 0x78, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x0c, 0x17, 0x7b, 0x40,
	0x69, 0x71, 0x06, 0xc8, 0x06, 0x16, 0xbd, 0xd0, 0x20, 0x1f, 0x29, 0x4e, 0x3d, 0x98, 0xa1, 0x4a,
	0x0c, 0x49, 0x6c, 0x60, 0x27, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x09, 0x33, 0x4a, 0x16,
	0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// URLBrokerServiceClient is the client API for URLBrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type URLBrokerServiceClient interface {
	//Pushes an URL to the server to open it on the default browser
	PushURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Response, error)
}

type uRLBrokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewURLBrokerServiceClient(cc grpc.ClientConnInterface) URLBrokerServiceClient {
	return &uRLBrokerServiceClient{cc}
}

func (c *uRLBrokerServiceClient) PushURL(ctx context.Context, in *URL, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/URLBrokerService/PushURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLBrokerServiceServer is the server API for URLBrokerService service.
type URLBrokerServiceServer interface {
	//Pushes an URL to the server to open it on the default browser
	PushURL(context.Context, *URL) (*Response, error)
}

// UnimplementedURLBrokerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedURLBrokerServiceServer struct {
}

func (*UnimplementedURLBrokerServiceServer) PushURL(ctx context.Context, req *URL) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushURL not implemented")
}

func RegisterURLBrokerServiceServer(s *grpc.Server, srv URLBrokerServiceServer) {
	s.RegisterService(&_URLBrokerService_serviceDesc, srv)
}

func _URLBrokerService_PushURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLBrokerServiceServer).PushURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/URLBrokerService/PushURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLBrokerServiceServer).PushURL(ctx, req.(*URL))
	}
	return interceptor(ctx, in, info, handler)
}

var _URLBrokerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "URLBrokerService",
	HandlerType: (*URLBrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushURL",
			Handler:    _URLBrokerService_PushURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/url_broker.proto",
}

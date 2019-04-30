// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/feature_demo/demo_multi_file_service.proto

package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/TvdBrink/protoc-gen-gorm/options"
import infoblox_api "github.com/infobloxopen/atlas-app-toolkit/query"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReadAccountRequest struct {
	// For a read request, the id field is the only to be specified
	Id     uint64                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Fields *infoblox_api.FieldSelection `protobuf:"bytes,2,opt,name=fields" json:"fields,omitempty"`
}

func (m *ReadAccountRequest) Reset()                    { *m = ReadAccountRequest{} }
func (m *ReadAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadAccountRequest) ProtoMessage()               {}
func (*ReadAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ReadAccountRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadAccountRequest) GetFields() *infoblox_api.FieldSelection {
	if m != nil {
		return m.Fields
	}
	return nil
}

type ReadBlogPostsResponse struct {
	Posts []*BlogPost `protobuf:"bytes,1,rep,name=posts" json:"posts,omitempty"`
}

func (m *ReadBlogPostsResponse) Reset()                    { *m = ReadBlogPostsResponse{} }
func (m *ReadBlogPostsResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadBlogPostsResponse) ProtoMessage()               {}
func (*ReadBlogPostsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ReadBlogPostsResponse) GetPosts() []*BlogPost {
	if m != nil {
		return m.Posts
	}
	return nil
}

func init() {
	proto.RegisterType((*ReadAccountRequest)(nil), "example.ReadAccountRequest")
	proto.RegisterType((*ReadBlogPostsResponse)(nil), "example.ReadBlogPostsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlogPostService service

type BlogPostServiceClient interface {
	Read(ctx context.Context, in *ReadAccountRequest, opts ...grpc.CallOption) (*ReadBlogPostsResponse, error)
}

type blogPostServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogPostServiceClient(cc *grpc.ClientConn) BlogPostServiceClient {
	return &blogPostServiceClient{cc}
}

func (c *blogPostServiceClient) Read(ctx context.Context, in *ReadAccountRequest, opts ...grpc.CallOption) (*ReadBlogPostsResponse, error) {
	out := new(ReadBlogPostsResponse)
	err := grpc.Invoke(ctx, "/example.BlogPostService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlogPostService service

type BlogPostServiceServer interface {
	Read(context.Context, *ReadAccountRequest) (*ReadBlogPostsResponse, error)
}

func RegisterBlogPostServiceServer(s *grpc.Server, srv BlogPostServiceServer) {
	s.RegisterService(&_BlogPostService_serviceDesc, srv)
}

func _BlogPostService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogPostServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.BlogPostService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogPostServiceServer).Read(ctx, req.(*ReadAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogPostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.BlogPostService",
	HandlerType: (*BlogPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _BlogPostService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/feature_demo/demo_multi_file_service.proto",
}

func init() { proto.RegisterFile("example/feature_demo/demo_multi_file_service.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x9d, 0x53, 0x32, 0x50, 0x2c, 0x08, 0xb3, 0x8a, 0x8c, 0x5d, 0x1c, 0x42, 0x1b,
	0xa8, 0x9e, 0xf4, 0xa2, 0x53, 0x76, 0xf0, 0x24, 0xdd, 0x6d, 0x07, 0x4b, 0xd6, 0xbe, 0xd6, 0x60,
	0xda, 0x97, 0x25, 0xa9, 0xcc, 0x3f, 0x4d, 0xff, 0x3a, 0xc9, 0xda, 0x09, 0x32, 0x85, 0x5d, 0x02,
	0x79, 0xdf, 0x1f, 0x7c, 0xf2, 0x42, 0x22, 0x58, 0xb2, 0x52, 0x0a, 0xa0, 0x39, 0x30, 0x53, 0x2b,
	0x48, 0x32, 0x28, 0x91, 0xda, 0x23, 0x29, 0x6b, 0x61, 0x78, 0x92, 0x73, 0x01, 0x89, 0x06, 0xf5,
	0xce, 0x53, 0x08, 0xa5, 0x42, 0x83, 0xde, 0x5e, 0x9b, 0xf1, 0x6f, 0x0a, 0x6e, 0x5e, 0xeb, 0x79,
	0x98, 0x62, 0x49, 0x79, 0x95, 0xe3, 0x5c, 0xe0, 0x12, 0x25, 0x54, 0x74, 0xe5, 0x4b, 0x83, 0x02,
	0xaa, 0xa0, 0x40, 0x55, 0x52, 0x94, 0x86, 0x63, 0xa5, 0xa9, 0xbd, 0x34, 0x25, 0xfe, 0xd3, 0x7f,
	0x59, 0x66, 0x04, 0xd3, 0x01, 0x93, 0x32, 0x30, 0x88, 0xe2, 0x8d, 0x1b, 0xba, 0xa8, 0x41, 0x7d,
	0xd0, 0x14, 0x85, 0x80, 0xd4, 0xf6, 0x24, 0x28, 0x41, 0x31, 0x83, 0x4a, 0xb7, 0x5d, 0x97, 0xdb,
	0x3c, 0xa2, 0xf1, 0x0e, 0x67, 0xc4, 0x8b, 0x81, 0x65, 0xf7, 0x69, 0x8a, 0x75, 0x65, 0x62, 0x58,
	0xd4, 0xa0, 0x8d, 0x77, 0x40, 0x5c, 0x9e, 0xf5, 0x9d, 0x81, 0x33, 0xea, 0xc4, 0x2e, 0xcf, 0xbc,
	0x6b, 0xd2, 0xcd, 0x39, 0x88, 0x4c, 0xf7, 0xdd, 0x81, 0x33, 0xea, 0x45, 0x67, 0xe1, 0x9a, 0x31,
	0x64, 0x92, 0x87, 0x13, 0xab, 0x4d, 0xa1, 0xe5, 0x89, 0x5b, 0xef, 0xf0, 0x8e, 0x1c, 0xdb, 0xee,
	0xb1, 0xc0, 0xe2, 0x19, 0xb5, 0xd1, 0x31, 0x68, 0x89, 0x95, 0x06, 0xef, 0x82, 0xec, 0x4a, 0x3b,
	0xe8, 0x3b, 0x83, 0x9d, 0x51, 0x2f, 0x3a, 0x0a, 0x5b, 0xe0, 0x70, 0x6d, 0x8d, 0x1b, 0x3d, 0x7a,
	0x21, 0x87, 0xeb, 0xd1, 0xb4, 0xd9, 0xb9, 0xf7, 0x40, 0x3a, 0xb6, 0xd4, 0x3b, 0xfd, 0x09, 0x6d,
	0xf2, 0xfb, 0xe7, 0xbf, 0xc4, 0x0d, 0x00, 0xbf, 0xfb, 0xf5, 0x79, 0xe2, 0xee, 0x3b, 0xe3, 0xc9,
	0xec, 0x71, 0xdb, 0x3f, 0xfb, 0x6b, 0xa7, 0xb7, 0xed, 0x70, 0xde, 0x5d, 0xb9, 0xaf, 0xbe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xb9, 0x74, 0x51, 0x1f, 0x3f, 0x02, 0x00, 0x00,
}

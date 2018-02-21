// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fetcher.proto

/*
Package fetcher is a generated protocol buffer package.

It is generated from these files:
	fetcher.proto

It has these top-level messages:
	FetchReq
	FetchReply
	Issue
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type FetchReq struct {
	RepoId string `protobuf:"bytes,1,opt,name=repoId" json:"repoId,omitempty"`
}

func (m *FetchReq) Reset()                    { *m = FetchReq{} }
func (m *FetchReq) String() string            { return proto.CompactTextString(m) }
func (*FetchReq) ProtoMessage()               {}
func (*FetchReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FetchReq) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

type FetchReply struct {
	Issues []*Issue `protobuf:"bytes,1,rep,name=issues" json:"issues,omitempty"`
}

func (m *FetchReply) Reset()                    { *m = FetchReply{} }
func (m *FetchReply) String() string            { return proto.CompactTextString(m) }
func (*FetchReply) ProtoMessage()               {}
func (*FetchReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchReply) GetIssues() []*Issue {
	if m != nil {
		return m.Issues
	}
	return nil
}

type Issue struct {
	RepoId    string `protobuf:"bytes,1,opt,name=repoId" json:"repoId,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *Issue) Reset()                    { *m = Issue{} }
func (m *Issue) String() string            { return proto.CompactTextString(m) }
func (*Issue) ProtoMessage()               {}
func (*Issue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Issue) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *Issue) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Issue) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Issue) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*FetchReq)(nil), "FetchReq")
	proto.RegisterType((*FetchReply)(nil), "FetchReply")
	proto.RegisterType((*Issue)(nil), "Issue")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IssueFetcher service

type IssueFetcherClient interface {
	FetchIssues(ctx context.Context, in *FetchReq, opts ...grpc.CallOption) (*FetchReply, error)
}

type issueFetcherClient struct {
	cc *grpc.ClientConn
}

func NewIssueFetcherClient(cc *grpc.ClientConn) IssueFetcherClient {
	return &issueFetcherClient{cc}
}

func (c *issueFetcherClient) FetchIssues(ctx context.Context, in *FetchReq, opts ...grpc.CallOption) (*FetchReply, error) {
	out := new(FetchReply)
	err := grpc.Invoke(ctx, "/IssueFetcher/FetchIssues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IssueFetcher service

type IssueFetcherServer interface {
	FetchIssues(context.Context, *FetchReq) (*FetchReply, error)
}

func RegisterIssueFetcherServer(s *grpc.Server, srv IssueFetcherServer) {
	s.RegisterService(&_IssueFetcher_serviceDesc, srv)
}

func _IssueFetcher_FetchIssues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssueFetcherServer).FetchIssues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IssueFetcher/FetchIssues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssueFetcherServer).FetchIssues(ctx, req.(*FetchReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _IssueFetcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "IssueFetcher",
	HandlerType: (*IssueFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchIssues",
			Handler:    _IssueFetcher_FetchIssues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fetcher.proto",
}

func init() { proto.RegisterFile("fetcher.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x41, 0x0b, 0x82, 0x30,
	0x18, 0x86, 0x33, 0x73, 0xe4, 0x67, 0x5d, 0x46, 0xc4, 0x88, 0x08, 0xd9, 0x25, 0x0f, 0xb1, 0x83,
	0x1d, 0xfa, 0x07, 0x82, 0x57, 0xff, 0x81, 0xd5, 0x22, 0x49, 0x71, 0xb9, 0xcf, 0x43, 0xff, 0x3e,
	0xfc, 0x66, 0x78, 0xea, 0xb6, 0xe7, 0xd9, 0x78, 0xc7, 0x03, 0xeb, 0x87, 0xc6, 0xdb, 0x53, 0x77,
	0xca, 0x74, 0x2d, 0xb6, 0x52, 0xc2, 0x32, 0x1b, 0x44, 0xa1, 0xdf, 0x7c, 0x0b, 0xac, 0xd3, 0xa6,
	0xcd, 0xef, 0xc2, 0x8b, 0xbd, 0x24, 0x2c, 0x46, 0x92, 0x27, 0x80, 0xf1, 0x8d, 0xa9, 0x3f, 0xfc,
	0x00, 0xac, 0xb2, 0xb6, 0xd7, 0x56, 0x78, 0xb1, 0x9f, 0x44, 0x29, 0x53, 0xf9, 0x80, 0xc5, 0x68,
	0xe5, 0x0b, 0x02, 0x12, 0xff, 0xe6, 0xf8, 0x06, 0x02, 0xac, 0xb0, 0xd6, 0x62, 0x4e, 0xda, 0x01,
	0xdf, 0x43, 0x88, 0x55, 0xa3, 0x2d, 0x96, 0x8d, 0x11, 0x3e, 0xdd, 0x4c, 0x62, 0xd8, 0xb2, 0x58,
	0x62, 0x6f, 0xc5, 0xc2, 0x6d, 0x39, 0x4a, 0x2f, 0xb0, 0xa2, 0xcf, 0x32, 0x17, 0xc5, 0x8f, 0x10,
	0xd1, 0x91, 0xa4, 0xe5, 0xa1, 0xfa, 0xc5, 0xed, 0x22, 0x35, 0x35, 0xc8, 0xd9, 0x95, 0x51, 0xfe,
	0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x74, 0xdf, 0x9f, 0x41, 0x0f, 0x01, 0x00, 0x00,
}

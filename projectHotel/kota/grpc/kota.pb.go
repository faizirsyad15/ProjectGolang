// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kota.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	kota.proto

It has these top-level messages:
	AddKotaReq
	ReadKotaByNamaReq
	ReadKotaByNamaResp
	ReadKotaResp
	UpdateKotaReq
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type AddKotaReq struct {
	IdKota   string `protobuf:"bytes,1,opt,name=idKota" json:"idKota,omitempty"`
	NamaKota string `protobuf:"bytes,2,opt,name=namaKota" json:"namaKota,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *AddKotaReq) Reset()                    { *m = AddKotaReq{} }
func (m *AddKotaReq) String() string            { return proto.CompactTextString(m) }
func (*AddKotaReq) ProtoMessage()               {}
func (*AddKotaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddKotaReq) GetIdKota() string {
	if m != nil {
		return m.IdKota
	}
	return ""
}

func (m *AddKotaReq) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

func (m *AddKotaReq) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ReadKotaByNamaReq struct {
	NamaKota string `protobuf:"bytes,1,opt,name=namaKota" json:"namaKota,omitempty"`
}

func (m *ReadKotaByNamaReq) Reset()                    { *m = ReadKotaByNamaReq{} }
func (m *ReadKotaByNamaReq) String() string            { return proto.CompactTextString(m) }
func (*ReadKotaByNamaReq) ProtoMessage()               {}
func (*ReadKotaByNamaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ReadKotaByNamaReq) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

type ReadKotaByNamaResp struct {
	IdKota   string `protobuf:"bytes,1,opt,name=idKota" json:"idKota,omitempty"`
	NamaKota string `protobuf:"bytes,2,opt,name=namaKota" json:"namaKota,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *ReadKotaByNamaResp) Reset()                    { *m = ReadKotaByNamaResp{} }
func (m *ReadKotaByNamaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKotaByNamaResp) ProtoMessage()               {}
func (*ReadKotaByNamaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadKotaByNamaResp) GetIdKota() string {
	if m != nil {
		return m.IdKota
	}
	return ""
}

func (m *ReadKotaByNamaResp) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

func (m *ReadKotaByNamaResp) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ReadKotaResp struct {
	AllKota []*ReadKotaByNamaResp `protobuf:"bytes,1,rep,name=allKota" json:"allKota,omitempty"`
}

func (m *ReadKotaResp) Reset()                    { *m = ReadKotaResp{} }
func (m *ReadKotaResp) String() string            { return proto.CompactTextString(m) }
func (*ReadKotaResp) ProtoMessage()               {}
func (*ReadKotaResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadKotaResp) GetAllKota() []*ReadKotaByNamaResp {
	if m != nil {
		return m.AllKota
	}
	return nil
}

type UpdateKotaReq struct {
	IdKota   string `protobuf:"bytes,1,opt,name=idKota" json:"idKota,omitempty"`
	NamaKota string `protobuf:"bytes,2,opt,name=namaKota" json:"namaKota,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *UpdateKotaReq) Reset()                    { *m = UpdateKotaReq{} }
func (m *UpdateKotaReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateKotaReq) ProtoMessage()               {}
func (*UpdateKotaReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateKotaReq) GetIdKota() string {
	if m != nil {
		return m.IdKota
	}
	return ""
}

func (m *UpdateKotaReq) GetNamaKota() string {
	if m != nil {
		return m.NamaKota
	}
	return ""
}

func (m *UpdateKotaReq) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*AddKotaReq)(nil), "grpc.AddKotaReq")
	proto.RegisterType((*ReadKotaByNamaReq)(nil), "grpc.ReadKotaByNamaReq")
	proto.RegisterType((*ReadKotaByNamaResp)(nil), "grpc.ReadKotaByNamaResp")
	proto.RegisterType((*ReadKotaResp)(nil), "grpc.ReadKotaResp")
	proto.RegisterType((*UpdateKotaReq)(nil), "grpc.UpdateKotaReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for KotaService service

type KotaServiceClient interface {
	AddKota(ctx context.Context, in *AddKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadKota(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKotaResp, error)
	UpdateKota(ctx context.Context, in *UpdateKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error)
	ReadKotaByNama(ctx context.Context, in *ReadKotaByNamaReq, opts ...grpc1.CallOption) (*ReadKotaByNamaResp, error)
}

type kotaServiceClient struct {
	cc *grpc1.ClientConn
}

func NewKotaServiceClient(cc *grpc1.ClientConn) KotaServiceClient {
	return &kotaServiceClient{cc}
}

func (c *kotaServiceClient) AddKota(ctx context.Context, in *AddKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/AddKota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotaServiceClient) ReadKota(ctx context.Context, in *google_protobuf.Empty, opts ...grpc1.CallOption) (*ReadKotaResp, error) {
	out := new(ReadKotaResp)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/ReadKota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotaServiceClient) UpdateKota(ctx context.Context, in *UpdateKotaReq, opts ...grpc1.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/UpdateKota", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kotaServiceClient) ReadKotaByNama(ctx context.Context, in *ReadKotaByNamaReq, opts ...grpc1.CallOption) (*ReadKotaByNamaResp, error) {
	out := new(ReadKotaByNamaResp)
	err := grpc1.Invoke(ctx, "/grpc.KotaService/ReadKotaByNama", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KotaService service

type KotaServiceServer interface {
	AddKota(context.Context, *AddKotaReq) (*google_protobuf.Empty, error)
	ReadKota(context.Context, *google_protobuf.Empty) (*ReadKotaResp, error)
	UpdateKota(context.Context, *UpdateKotaReq) (*google_protobuf.Empty, error)
	ReadKotaByNama(context.Context, *ReadKotaByNamaReq) (*ReadKotaByNamaResp, error)
}

func RegisterKotaServiceServer(s *grpc1.Server, srv KotaServiceServer) {
	s.RegisterService(&_KotaService_serviceDesc, srv)
}

func _KotaService_AddKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKotaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).AddKota(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/AddKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).AddKota(ctx, req.(*AddKotaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotaService_ReadKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).ReadKota(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/ReadKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).ReadKota(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotaService_UpdateKota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKotaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).UpdateKota(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/UpdateKota",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).UpdateKota(ctx, req.(*UpdateKotaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KotaService_ReadKotaByNama_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadKotaByNamaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KotaServiceServer).ReadKotaByNama(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.KotaService/ReadKotaByNama",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KotaServiceServer).ReadKotaByNama(ctx, req.(*ReadKotaByNamaReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _KotaService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.KotaService",
	HandlerType: (*KotaServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "AddKota",
			Handler:    _KotaService_AddKota_Handler,
		},
		{
			MethodName: "ReadKota",
			Handler:    _KotaService_ReadKota_Handler,
		},
		{
			MethodName: "UpdateKota",
			Handler:    _KotaService_UpdateKota_Handler,
		},
		{
			MethodName: "ReadKotaByNama",
			Handler:    _KotaService_ReadKotaByNama_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "kota.proto",
}

func init() { proto.RegisterFile("kota.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x6d, 0x37, 0xd9, 0xe6, 0x37, 0x15, 0xfd, 0x84, 0x59, 0xe2, 0x45, 0x72, 0xda, 0x29, 0x85,
	0x8a, 0x20, 0x78, 0x72, 0xb0, 0x93, 0xe0, 0xa1, 0x22, 0x08, 0x5e, 0xcc, 0xd6, 0x58, 0x86, 0xad,
	0xc9, 0xda, 0x4c, 0xd8, 0xd5, 0x5f, 0x2e, 0x49, 0x5a, 0x67, 0xd5, 0x7a, 0x72, 0xc7, 0x2f, 0x79,
	0xef, 0x7d, 0xdf, 0x7b, 0x0f, 0xe0, 0x45, 0x6a, 0xce, 0x54, 0x21, 0xb5, 0xc4, 0x9d, 0xb4, 0x50,
	0x73, 0x72, 0x9a, 0x4a, 0x99, 0x66, 0x22, 0xb4, 0x6f, 0xb3, 0xd5, 0x73, 0x28, 0x72, 0xa5, 0xd7,
	0x0e, 0x42, 0x1f, 0x00, 0xae, 0x93, 0xe4, 0x46, 0x6a, 0x1e, 0x8b, 0x25, 0x8e, 0xa0, 0xb7, 0xb0,
	0x43, 0xe0, 0x9f, 0xf9, 0xe3, 0xdd, 0xb8, 0x9a, 0x90, 0xc0, 0xe0, 0x95, 0xe7, 0xdc, 0xfe, 0x74,
	0xec, 0xcf, 0xe7, 0x6c, 0x38, 0xa5, 0xe6, 0x7a, 0x55, 0x06, 0x5d, 0xc7, 0x71, 0x13, 0x0d, 0xe1,
	0x28, 0x16, 0xdc, 0xf2, 0x27, 0xeb, 0x5b, 0x9e, 0xdb, 0x05, 0x5f, 0x85, 0xfc, 0xa6, 0x10, 0x7d,
	0x02, 0xfc, 0x4e, 0x28, 0xd5, 0xbf, 0x9e, 0x34, 0x81, 0xbd, 0x7a, 0x83, 0xd5, 0x8e, 0xa0, 0xcf,
	0xb3, 0xac, 0x12, 0xef, 0x8e, 0x87, 0x51, 0xc0, 0x4c, 0x62, 0xec, 0xe7, 0x19, 0x71, 0x0d, 0xa4,
	0x8f, 0xb0, 0x7f, 0xaf, 0x12, 0xae, 0xc5, 0x16, 0x32, 0x8b, 0xde, 0x3b, 0x30, 0x34, 0x80, 0x3b,
	0x51, 0xbc, 0x2d, 0xe6, 0x02, 0x2f, 0xa0, 0x5f, 0xb5, 0x83, 0x87, 0xee, 0xb4, 0x4d, 0x59, 0x64,
	0xc4, 0x5c, 0xb1, 0xac, 0x2e, 0x96, 0x4d, 0x4d, 0xb1, 0xd4, 0xc3, 0x4b, 0x18, 0xd4, 0x16, 0xb0,
	0x05, 0x45, 0xb0, 0x69, 0xd5, 0x98, 0xa4, 0x1e, 0x5e, 0x01, 0x6c, 0xdc, 0xe1, 0xb1, 0xc3, 0x34,
	0xfc, 0xfe, 0xb1, 0x76, 0x0a, 0x07, 0xcd, 0xe4, 0xf0, 0xe4, 0xf7, 0x3c, 0x97, 0xa4, 0x35, 0x68,
	0xea, 0xcd, 0x7a, 0x56, 0xf8, 0xfc, 0x23, 0x00, 0x00, 0xff, 0xff, 0x52, 0xe9, 0x82, 0xf6, 0xca,
	0x02, 0x00, 0x00,
}

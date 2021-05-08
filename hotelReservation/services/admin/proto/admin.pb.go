// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin.proto

/*
Package admin is a generated protocol buffer package.

It is generated from these files:
	admin.proto

It has these top-level messages:
	CheckRequest
	CheckReply
	UpdateRequest
	UpdateReply
	RegisterRequest
	RegisterReply
	LoginRequest
	LoginReply
*/
package admin

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

type CheckRequest struct {
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CheckRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CheckRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CheckReply struct {
	Correct bool `protobuf:"varint,1,opt,name=correct" json:"correct,omitempty"`
}

func (m *CheckReply) Reset()                    { *m = CheckReply{} }
func (m *CheckReply) String() string            { return proto.CompactTextString(m) }
func (*CheckReply) ProtoMessage()               {}
func (*CheckReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckReply) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

type UpdateRequest struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Target  string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *UpdateRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdateReply struct {
	Correct bool `protobuf:"varint,1,opt,name=correct" json:"correct,omitempty"`
}

func (m *UpdateReply) Reset()                    { *m = UpdateReply{} }
func (m *UpdateReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateReply) ProtoMessage()               {}
func (*UpdateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateReply) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

type RegisterRequest struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string   `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Hotels   []string `protobuf:"bytes,4,rep,name=hotels" json:"hotels,omitempty"`
	Id       string   `protobuf:"bytes,5,opt,name=id" json:"id,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetHotels() []string {
	if m != nil {
		return m.Hotels
	}
	return nil
}

func (m *RegisterRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RegisterReply struct {
	Correct bool `protobuf:"varint,1,opt,name=correct" json:"correct,omitempty"`
}

func (m *RegisterReply) Reset()                    { *m = RegisterReply{} }
func (m *RegisterReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()               {}
func (*RegisterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegisterReply) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

type LoginRequest struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	Correct bool `protobuf:"varint,1,opt,name=correct" json:"correct,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LoginReply) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "admin.CheckRequest")
	proto.RegisterType((*CheckReply)(nil), "admin.CheckReply")
	proto.RegisterType((*UpdateRequest)(nil), "admin.UpdateRequest")
	proto.RegisterType((*UpdateReply)(nil), "admin.UpdateReply")
	proto.RegisterType((*RegisterRequest)(nil), "admin.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "admin.RegisterReply")
	proto.RegisterType((*LoginRequest)(nil), "admin.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "admin.LoginReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Admin service

type AdminClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	CheckHotel(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckReply, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/admin.Admin/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := grpc.Invoke(ctx, "/admin.Admin/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := grpc.Invoke(ctx, "/admin.Admin/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) CheckHotel(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckReply, error) {
	out := new(CheckReply)
	err := grpc.Invoke(ctx, "/admin.Admin/CheckHotel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Admin service

type AdminServer interface {
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	CheckHotel(context.Context, *CheckRequest) (*CheckReply, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_CheckHotel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CheckHotel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.Admin/CheckHotel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CheckHotel(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Admin_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Admin_Register_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Admin_Update_Handler,
		},
		{
			MethodName: "CheckHotel",
			Handler:    _Admin_CheckHotel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x4f, 0xbb, 0x40,
	0x10, 0x0d, 0xb4, 0xf4, 0xd7, 0xdf, 0xb4, 0xd5, 0xb8, 0x36, 0x0d, 0xe1, 0xd4, 0x70, 0xf0, 0xcf,
	0xa5, 0x26, 0xb5, 0x07, 0x8f, 0x1a, 0x2f, 0x1e, 0xbc, 0x48, 0xe2, 0x07, 0x58, 0x61, 0x42, 0x37,
	0x02, 0x8b, 0xcb, 0x1a, 0xe3, 0xc9, 0x2f, 0xea, 0x87, 0x31, 0xbb, 0xec, 0xd2, 0x42, 0x53, 0x6e,
	0xbc, 0x59, 0xe6, 0xbd, 0x37, 0x33, 0x0f, 0x26, 0x34, 0xc9, 0x59, 0xb1, 0x2a, 0x05, 0x97, 0x9c,
	0x78, 0x1a, 0x84, 0x1b, 0x98, 0x3e, 0x6e, 0x31, 0x7e, 0x8f, 0xf0, 0xe3, 0x13, 0x2b, 0x49, 0xe6,
	0xe0, 0x61, 0x4e, 0x59, 0xe6, 0xbb, 0x4b, 0xe7, 0xea, 0x7f, 0x54, 0x03, 0x72, 0x02, 0x2e, 0x4b,
	0x7c, 0x47, 0x97, 0x5c, 0x96, 0x84, 0x17, 0x00, 0xa6, 0xab, 0xcc, 0xbe, 0x89, 0x0f, 0xff, 0x62,
	0x2e, 0x04, 0xc6, 0x52, 0xff, 0x32, 0x8e, 0x2c, 0x0c, 0x5f, 0x60, 0xf6, 0x5a, 0x26, 0x54, 0xa2,
	0xa5, 0xef, 0x10, 0x91, 0x05, 0x8c, 0x24, 0x15, 0x29, 0x4a, 0xa3, 0x67, 0x50, 0x4d, 0x59, 0x48,
	0x2c, 0xa4, 0x3f, 0xd0, 0x0f, 0x16, 0x86, 0x97, 0x30, 0xb1, 0x94, 0xfd, 0xda, 0x3f, 0x70, 0x1a,
	0x61, 0xca, 0x2a, 0x89, 0xc2, 0xaa, 0x13, 0x18, 0x16, 0x34, 0x47, 0xa3, 0xaf, 0xbf, 0x8f, 0x0c,
	0x1c, 0xc0, 0xb8, 0xa4, 0x55, 0xf5, 0xc5, 0x45, 0x62, 0x0c, 0x34, 0x58, 0x79, 0xde, 0x72, 0x89,
	0x59, 0xe5, 0x0f, 0x97, 0x03, 0xe5, 0xb9, 0x46, 0x66, 0x36, 0xaf, 0x59, 0xd2, 0x35, 0xcc, 0x76,
	0x06, 0xfa, 0xbd, 0xde, 0xc3, 0xf4, 0x99, 0xa7, 0xac, 0x38, 0xb8, 0x82, 0x73, 0xcc, 0x94, 0xdb,
	0x36, 0xa5, 0x2e, 0x62, 0x18, 0x7a, 0x95, 0xd6, 0xbf, 0x0e, 0x78, 0x0f, 0xea, 0xf2, 0xe4, 0x06,
	0x3c, 0xdd, 0x41, 0xce, 0x57, 0x75, 0x2e, 0xf6, 0x1d, 0x04, 0x67, 0xed, 0xa2, 0x22, 0xbd, 0x83,
	0xb1, 0x9d, 0x87, 0x2c, 0xcc, 0x73, 0x67, 0xc3, 0xc1, 0xfc, 0xa0, 0xae, 0x3a, 0xd7, 0x30, 0xaa,
	0x6f, 0x46, 0xec, 0x7b, 0x2b, 0x15, 0x01, 0xe9, 0x54, 0x55, 0xcf, 0xc6, 0x44, 0xec, 0x49, 0x2d,
	0xb7, 0xf1, 0xb8, 0x9f, 0xd5, 0xc6, 0xe3, 0x2e, 0x8a, 0x6f, 0x23, 0x1d, 0xee, 0xdb, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x70, 0xd5, 0xa4, 0xeb, 0x02, 0x00, 0x00,
}
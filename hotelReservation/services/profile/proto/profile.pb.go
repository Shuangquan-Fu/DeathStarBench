// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Request
	Result
	Hotel
	Address
	Image
	ScoreRequest
	ScoreResult
*/
package profile

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

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	Locale   string   `protobuf:"bytes,2,opt,name=locale" json:"locale,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type Result struct {
	Hotels []*Hotel `protobuf:"bytes,1,rep,name=hotels" json:"hotels,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" json:"images,omitempty"`
	Price       float32  `protobuf:"fixed32,7,opt,name=price" json:"price,omitempty"`
	Score       float32  `protobuf:"fixed32,8,opt,name=score" json:"score,omitempty"`
	ScoreTimes  int32    `protobuf:"varint,9,opt,name=scoreTimes" json:"scoreTimes,omitempty"`
}

func (m *Hotel) Reset()                    { *m = Hotel{} }
func (m *Hotel) String() string            { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()               {}
func (*Hotel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Hotel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hotel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hotel) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Hotel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *Hotel) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Hotel) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Hotel) GetScoreTimes() int32 {
	if m != nil {
		return m.ScoreTimes
	}
	return 0
}

type Address struct {
	StreetNumber string  `protobuf:"bytes,1,opt,name=streetNumber" json:"streetNumber,omitempty"`
	StreetName   string  `protobuf:"bytes,2,opt,name=streetName" json:"streetName,omitempty"`
	City         string  `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State        string  `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country      string  `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	PostalCode   string  `protobuf:"bytes,6,opt,name=postalCode" json:"postalCode,omitempty"`
	Lat          float32 `protobuf:"fixed32,7,opt,name=lat" json:"lat,omitempty"`
	Lon          float32 `protobuf:"fixed32,8,opt,name=lon" json:"lon,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Address) GetStreetNumber() string {
	if m != nil {
		return m.StreetNumber
	}
	return ""
}

func (m *Address) GetStreetName() string {
	if m != nil {
		return m.StreetName
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Address) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Address) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Image struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" json:"default,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

type ScoreRequest struct {
	HotelId string  `protobuf:"bytes,1,opt,name=hotelId" json:"hotelId,omitempty"`
	Score   float32 `protobuf:"fixed32,2,opt,name=score" json:"score,omitempty"`
}

func (m *ScoreRequest) Reset()                    { *m = ScoreRequest{} }
func (m *ScoreRequest) String() string            { return proto.CompactTextString(m) }
func (*ScoreRequest) ProtoMessage()               {}
func (*ScoreRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ScoreRequest) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *ScoreRequest) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type ScoreResult struct {
	Correct bool `protobuf:"varint,1,opt,name=correct" json:"correct,omitempty"`
}

func (m *ScoreResult) Reset()                    { *m = ScoreResult{} }
func (m *ScoreResult) String() string            { return proto.CompactTextString(m) }
func (*ScoreResult) ProtoMessage()               {}
func (*ScoreResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ScoreResult) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
	proto.RegisterType((*ScoreRequest)(nil), "profile.ScoreRequest")
	proto.RegisterType((*ScoreResult)(nil), "profile.ScoreResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Profile service

type ProfileClient interface {
	GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	UpdateScore(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*ScoreResult, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/profile.Profile/GetProfiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateScore(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*ScoreResult, error) {
	out := new(ScoreResult)
	err := grpc.Invoke(ctx, "/profile.Profile/UpdateScore", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	GetProfiles(context.Context, *Request) (*Result, error)
	UpdateScore(context.Context, *ScoreRequest) (*ScoreResult, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfiles(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/UpdateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateScore(ctx, req.(*ScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profile_GetProfiles_Handler,
		},
		{
			MethodName: "UpdateScore",
			Handler:    _Profile_UpdateScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0x6c, 0x93, 0xec, 0x4e, 0x4a, 0xa9, 0xac, 0x82, 0xac, 0x3d, 0xa0, 0x28, 0x07,
	0x88, 0x38, 0x54, 0xd5, 0xf6, 0x86, 0x04, 0x12, 0xe2, 0x00, 0xbd, 0x20, 0x64, 0xe0, 0x01, 0xd2,
	0x64, 0x4a, 0x2d, 0x79, 0xe3, 0x60, 0x3b, 0x42, 0x7d, 0x0a, 0x9e, 0x8e, 0xf7, 0x41, 0x1e, 0x3b,
	0xbb, 0x69, 0x6f, 0xf3, 0x7f, 0xe3, 0xf1, 0xcc, 0xfc, 0x71, 0xe0, 0xd9, 0x68, 0xf4, 0x9d, 0x54,
	0x78, 0x39, 0x1a, 0xed, 0x34, 0x2b, 0xa2, 0xac, 0xdf, 0x43, 0x21, 0xf0, 0xf7, 0x84, 0xd6, 0xb1,
	0x2d, 0xac, 0xef, 0xb5, 0x43, 0x75, 0xd3, 0x5b, 0x9e, 0x54, 0xab, 0x66, 0x23, 0x0e, 0x9a, 0xbd,
	0x84, 0x5c, 0xe9, 0xae, 0x55, 0xc8, 0xd3, 0x2a, 0x69, 0x36, 0x22, 0xaa, 0xfa, 0x0a, 0x72, 0x81,
	0x76, 0x52, 0x8e, 0xbd, 0x86, 0x9c, 0x4e, 0x87, 0xda, 0x72, 0x77, 0x76, 0x39, 0x77, 0xfc, 0xe2,
	0xb1, 0x88, 0xd9, 0xfa, 0x6f, 0x0a, 0x19, 0x11, 0x76, 0x06, 0xa9, 0xec, 0x79, 0x42, 0xf7, 0xa5,
	0xb2, 0x67, 0x0c, 0x4e, 0x86, 0x76, 0x3f, 0x77, 0xa0, 0x98, 0x55, 0x50, 0x8e, 0xf7, 0x7a, 0xc0,
	0xaf, 0xd3, 0xfe, 0x16, 0x0d, 0x5f, 0x51, 0x6a, 0x89, 0xfc, 0x89, 0x1e, 0x6d, 0x67, 0xe4, 0xe8,
	0xa4, 0x1e, 0xf8, 0x49, 0x38, 0xb1, 0x40, 0xec, 0x2d, 0x14, 0x6d, 0xdf, 0x1b, 0xb4, 0x96, 0x67,
	0x55, 0xd2, 0x94, 0xbb, 0xf3, 0xc3, 0x68, 0x1f, 0x03, 0x17, 0xf3, 0x01, 0xbf, 0x85, 0xdc, 0xb7,
	0xbf, 0xd0, 0xf2, 0xfc, 0xc9, 0x16, 0x37, 0x1e, 0x8b, 0x98, 0x65, 0x17, 0x90, 0x8d, 0x46, 0x76,
	0xc8, 0x8b, 0x2a, 0x69, 0x52, 0x11, 0x84, 0xa7, 0xb6, 0xd3, 0x06, 0xf9, 0x3a, 0x50, 0x12, 0xec,
	0x15, 0x00, 0x05, 0x3f, 0xe4, 0x1e, 0x2d, 0xdf, 0x54, 0x49, 0x93, 0x89, 0x05, 0xa9, 0xff, 0x25,
	0x50, 0xc4, 0x41, 0x58, 0x0d, 0xa7, 0xd6, 0x19, 0x44, 0x17, 0x17, 0x0e, 0xee, 0x3c, 0x62, 0x74,
	0x5f, 0xd0, 0x47, 0xb7, 0x16, 0xc4, 0xfb, 0xd8, 0x49, 0xf7, 0x10, 0xcd, 0xa2, 0x98, 0x26, 0x73,
	0xad, 0xc3, 0xe8, 0x4f, 0x10, 0x8c, 0x43, 0xd1, 0xe9, 0x69, 0x70, 0xe6, 0x81, 0x9c, 0xd9, 0x88,
	0x59, 0xfa, 0x1e, 0xa3, 0xb6, 0xae, 0x55, 0x9f, 0x74, 0x8f, 0x3c, 0x0f, 0x3d, 0x8e, 0x84, 0x9d,
	0xc3, 0x4a, 0xb5, 0x2e, 0x6e, 0xef, 0x43, 0x22, 0x7a, 0x88, 0x9b, 0xfb, 0xb0, 0xbe, 0x86, 0x8c,
	0x4c, 0xf3, 0xa9, 0xc9, 0xa8, 0xb8, 0x8b, 0x0f, 0x7d, 0xe3, 0x1e, 0xef, 0xda, 0x49, 0x39, 0x9a,
	0x7f, 0x2d, 0x66, 0x59, 0x7f, 0x80, 0xd3, 0xef, 0xde, 0x9a, 0xf9, 0x51, 0x72, 0x28, 0xe2, 0x23,
	0x8c, 0xf5, 0xb3, 0x3c, 0x9a, 0x9d, 0x2e, 0xcc, 0xae, 0xdf, 0x40, 0x19, 0xeb, 0xe9, 0x55, 0xd2,
	0x86, 0xc6, 0x60, 0xe7, 0xa8, 0x7c, 0x2d, 0x66, 0xb9, 0xfb, 0x03, 0xc5, 0xb7, 0xf0, 0x69, 0xd9,
	0x15, 0x94, 0x9f, 0xd1, 0x45, 0x65, 0xd9, 0xf1, 0x79, 0xc4, 0x21, 0xb6, 0xcf, 0x17, 0x84, 0xae,
	0x7d, 0x07, 0xe5, 0xcf, 0xb1, 0x6f, 0x1d, 0x52, 0x2f, 0xf6, 0xe2, 0x90, 0x5f, 0xce, 0xbe, 0xbd,
	0x78, 0x8a, 0x7d, 0xed, 0x6d, 0x4e, 0x7f, 0xe0, 0xf5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0x3d, 0xf4, 0x1b, 0x92, 0x03, 0x00, 0x00,
}

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

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
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

// Server API for Profile service

type ProfileServer interface {
	GetProfiles(context.Context, *Request) (*Result, error)
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

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profile_GetProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x6c, 0x93, 0xec, 0xce, 0x42, 0xa9, 0x46, 0x08, 0x59, 0x3d, 0xa0, 0x28, 0x07,
	0x14, 0x71, 0xa8, 0xaa, 0xed, 0x11, 0x71, 0x40, 0x1c, 0xa0, 0x17, 0x84, 0x2c, 0x5e, 0xc0, 0x8d,
	0xa7, 0xd4, 0x92, 0x37, 0x0e, 0xb6, 0x73, 0xe8, 0x53, 0xf0, 0x74, 0xbc, 0x0f, 0xf2, 0x9f, 0xec,
	0x06, 0x6e, 0xf3, 0xfd, 0x66, 0xe2, 0x99, 0xf9, 0x32, 0xf0, 0x72, 0xb2, 0xe6, 0x51, 0x69, 0xba,
	0x99, 0xac, 0xf1, 0x06, 0x9b, 0x2c, 0xbb, 0x8f, 0xd0, 0x70, 0xfa, 0x35, 0x93, 0xf3, 0x78, 0x0d,
	0xdb, 0x27, 0xe3, 0x49, 0xdf, 0x4b, 0xc7, 0x8a, 0x76, 0xd3, 0xef, 0xf8, 0x49, 0xe3, 0x1b, 0xa8,
	0xb5, 0x19, 0x84, 0x26, 0x56, 0xb6, 0x45, 0xbf, 0xe3, 0x59, 0x75, 0xb7, 0x50, 0x73, 0x72, 0xb3,
	0xf6, 0xf8, 0x0e, 0xea, 0x58, 0x9d, 0xbe, 0xdd, 0x1f, 0x2e, 0x6f, 0x96, 0x8e, 0x5f, 0x03, 0xe6,
	0x39, 0xdb, 0xfd, 0x2e, 0xa1, 0x8a, 0x04, 0x2f, 0xa1, 0x54, 0x92, 0x15, 0xf1, 0xbd, 0x52, 0x49,
	0x44, 0xb8, 0x18, 0xc5, 0x71, 0xe9, 0x10, 0x63, 0x6c, 0x61, 0x3f, 0x3d, 0x99, 0x91, 0xbe, 0xcd,
	0xc7, 0x07, 0xb2, 0x6c, 0x13, 0x53, 0x6b, 0x14, 0x2a, 0x24, 0xb9, 0xc1, 0xaa, 0xc9, 0x2b, 0x33,
	0xb2, 0x8b, 0x54, 0xb1, 0x42, 0xf8, 0x1e, 0x1a, 0x21, 0xa5, 0x25, 0xe7, 0x58, 0xd5, 0x16, 0xfd,
	0xfe, 0x70, 0x75, 0x1a, 0xed, 0x53, 0xe2, 0x7c, 0x29, 0x08, 0x5b, 0xa8, 0xa3, 0xf8, 0x49, 0x8e,
	0xd5, 0xff, 0x6d, 0x71, 0x1f, 0x30, 0xcf, 0x59, 0x7c, 0x0d, 0xd5, 0x64, 0xd5, 0x40, 0xac, 0x69,
	0x8b, 0xbe, 0xe4, 0x49, 0x04, 0xea, 0x06, 0x63, 0x89, 0x6d, 0x13, 0x8d, 0x02, 0xdf, 0x02, 0xc4,
	0xe0, 0x87, 0x3a, 0x92, 0x63, 0xbb, 0xb6, 0xe8, 0x2b, 0xbe, 0x22, 0xdd, 0x9f, 0x02, 0x9a, 0x3c,
	0x08, 0x76, 0xf0, 0xc2, 0x79, 0x4b, 0xe4, 0xf3, 0xc2, 0xc9, 0x9d, 0x7f, 0x58, 0x7c, 0x2f, 0xe9,
	0xb3, 0x5b, 0x2b, 0x12, 0x7c, 0x1c, 0x94, 0x7f, 0xce, 0x66, 0xc5, 0x38, 0x4e, 0xe6, 0x85, 0xa7,
	0xec, 0x4f, 0x12, 0xc8, 0xa0, 0x19, 0xcc, 0x3c, 0x7a, 0xfb, 0x1c, 0x9d, 0xd9, 0xf1, 0x45, 0x86,
	0x1e, 0x93, 0x71, 0x5e, 0xe8, 0xcf, 0x46, 0x12, 0xab, 0x53, 0x8f, 0x33, 0xc1, 0x2b, 0xd8, 0x68,
	0xe1, 0xf3, 0xf6, 0x21, 0x8c, 0xc4, 0x8c, 0x79, 0xf3, 0x10, 0x76, 0x77, 0x50, 0x45, 0xd3, 0x42,
	0x6a, 0xb6, 0x3a, 0xef, 0x12, 0xc2, 0xd0, 0x58, 0xd2, 0xa3, 0x98, 0xb5, 0x8f, 0xf3, 0x6f, 0xf9,
	0x22, 0x0f, 0x1f, 0xa0, 0xf9, 0x9e, 0x1c, 0xc7, 0x5b, 0xd8, 0x7f, 0x21, 0x9f, 0x95, 0xc3, 0xf3,
	0x5f, 0xcb, 0x07, 0x7b, 0xfd, 0x6a, 0x45, 0xc2, 0x0d, 0x3e, 0xd4, 0xf1, 0xb8, 0xef, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x73, 0x9f, 0xbb, 0xd6, 0xed, 0x02, 0x00, 0x00,
}

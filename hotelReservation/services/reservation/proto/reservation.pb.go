// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reservation.proto

/*
Package reservation is a generated protocol buffer package.

It is generated from these files:
	reservation.proto

It has these top-level messages:
	Request
	Result
*/
package reservation

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
	CustomerName string   `protobuf:"bytes,1,opt,name=customerName" json:"customerName,omitempty"`
	HotelId      []string `protobuf:"bytes,2,rep,name=hotelId" json:"hotelId,omitempty"`
	InDate       string   `protobuf:"bytes,3,opt,name=inDate" json:"inDate,omitempty"`
	OutDate      string   `protobuf:"bytes,4,opt,name=outDate" json:"outDate,omitempty"`
	RoomNumber   int32    `protobuf:"varint,5,opt,name=roomNumber" json:"roomNumber,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetCustomerName() string {
	if m != nil {
		return m.CustomerName
	}
	return ""
}

func (m *Request) GetHotelId() []string {
	if m != nil {
		return m.HotelId
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *Request) GetRoomNumber() int32 {
	if m != nil {
		return m.RoomNumber
	}
	return 0
}

type Result struct {
	HotelId []string `protobuf:"bytes,1,rep,name=hotelId" json:"hotelId,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotelId() []string {
	if m != nil {
		return m.HotelId
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "reservation.Request")
	proto.RegisterType((*Result)(nil), "reservation.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Reservation service

type ReservationClient interface {
	// MakeReservation makes a reservation based on given information
	MakeReservation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	// CancelReservation Cancels a reservation based on given information
	CancelReservation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
	// CheckAvailability checks if given information is available
	CheckAvailability(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type reservationClient struct {
	cc *grpc.ClientConn
}

func NewReservationClient(cc *grpc.ClientConn) ReservationClient {
	return &reservationClient{cc}
}

func (c *reservationClient) MakeReservation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/reservation.Reservation/MakeReservation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) CancelReservation(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/reservation.Reservation/CancelReservation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) CheckAvailability(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/reservation.Reservation/CheckAvailability", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reservation service

type ReservationServer interface {
	// MakeReservation makes a reservation based on given information
	MakeReservation(context.Context, *Request) (*Result, error)
	// CancelReservation Cancels a reservation based on given information
	CancelReservation(context.Context, *Request) (*Result, error)
	// CheckAvailability checks if given information is available
	CheckAvailability(context.Context, *Request) (*Result, error)
}

func RegisterReservationServer(s *grpc.Server, srv ReservationServer) {
	s.RegisterService(&_Reservation_serviceDesc, srv)
}

func _Reservation_MakeReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).MakeReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/MakeReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).MakeReservation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/CancelReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).CancelReservation(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reservation.Reservation/CheckAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).CheckAvailability(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reservation_serviceDesc = grpc.ServiceDesc{
	ServiceName: "reservation.Reservation",
	HandlerType: (*ReservationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeReservation",
			Handler:    _Reservation_MakeReservation_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _Reservation_CancelReservation_Handler,
		},
		{
			MethodName: "CheckAvailability",
			Handler:    _Reservation_CheckAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation.proto",
}

func init() { proto.RegisterFile("reservation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xb5, 0x5b, 0x3a, 0x15, 0xa4, 0x51, 0x24, 0x78, 0x90, 0x65, 0x4f, 0x7b, 0xea,
	0x41, 0xaf, 0x22, 0x88, 0x5e, 0x3c, 0xd8, 0x43, 0xde, 0x20, 0xbb, 0x0e, 0x34, 0x34, 0xd9, 0xd1,
	0x64, 0x52, 0xf0, 0x51, 0x7c, 0x26, 0x5f, 0x4a, 0x0c, 0x2d, 0x66, 0xbd, 0x89, 0xc7, 0xff, 0x1b,
	0x3e, 0x7e, 0xfe, 0x04, 0x96, 0x01, 0x23, 0x86, 0x9d, 0x61, 0x4b, 0xc3, 0xea, 0x35, 0x10, 0x93,
	0x5c, 0x14, 0xa8, 0xf9, 0x10, 0x30, 0xd3, 0xf8, 0x96, 0x30, 0xb2, 0x6c, 0xe0, 0xa4, 0x4f, 0x91,
	0xc9, 0x63, 0x58, 0x1b, 0x8f, 0x4a, 0xd4, 0xa2, 0x9d, 0xeb, 0x11, 0x93, 0x0a, 0x66, 0x1b, 0x62,
	0x74, 0x4f, 0x2f, 0xea, 0xa8, 0x9e, 0xb4, 0x73, 0x7d, 0x88, 0xf2, 0x02, 0x2a, 0x3b, 0x3c, 0x1a,
	0x46, 0x35, 0xc9, 0xde, 0x3e, 0x7d, 0x1b, 0x94, 0x38, 0x1f, 0x8e, 0xf3, 0xe1, 0x10, 0xe5, 0x15,
	0x40, 0x20, 0xf2, 0xeb, 0xe4, 0x3b, 0x0c, 0x6a, 0x5a, 0x8b, 0x76, 0xaa, 0x0b, 0xd2, 0x34, 0x50,
	0x69, 0x8c, 0xc9, 0x71, 0xd9, 0x2a, 0x46, 0xad, 0xd7, 0x9f, 0x02, 0x16, 0xfa, 0x67, 0x8f, 0xbc,
	0x85, 0xd3, 0x67, 0xb3, 0xc5, 0x12, 0x9d, 0xaf, 0xca, 0x37, 0xd8, 0x8f, 0xbd, 0x3c, 0xfb, 0x45,
	0x73, 0xcf, 0x1d, 0x2c, 0x1f, 0xcc, 0xd0, 0xa3, 0xfb, 0x87, 0xbf, 0xc1, 0x7e, 0x7b, 0xbf, 0x33,
	0xd6, 0x99, 0xce, 0x3a, 0xcb, 0xef, 0x7f, 0xf0, 0xbb, 0x2a, 0xff, 0xd0, 0xcd, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0x75, 0x5e, 0xc5, 0xb6, 0x01, 0x00, 0x00,
}

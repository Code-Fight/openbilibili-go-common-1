// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: app/service/openplatform/ticket-item/api/grpc/v1/guest.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request message containing the guest info to update/insert
type GuestInfoRequest struct {
	ID          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id" validate:"min=0"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" validate:"max=32"`
	GuestImg    string `protobuf:"bytes,3,opt,name=guest_img,json=guestImg,proto3" json:"guest_img" validate:"min=0"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description" validate:"max=128"`
	GuestID     int64  `protobuf:"varint,5,opt,name=guest_id,json=guestId,proto3" json:"guest_id"`
}

func (m *GuestInfoRequest) Reset()                    { *m = GuestInfoRequest{} }
func (m *GuestInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GuestInfoRequest) ProtoMessage()               {}
func (*GuestInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorGuest, []int{0} }

// The request message containing the required info to change guest's status
type GuestStatusRequest struct {
	ID     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" validate:"required,min=1"`
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status" validate:"max=1"`
}

func (m *GuestStatusRequest) Reset()                    { *m = GuestStatusRequest{} }
func (m *GuestStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*GuestStatusRequest) ProtoMessage()               {}
func (*GuestStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptorGuest, []int{1} }

// The general response message contaning the result after updating/inserting the guest info
type GuestInfoReply struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success"`
}

func (m *GuestInfoReply) Reset()                    { *m = GuestInfoReply{} }
func (m *GuestInfoReply) String() string            { return proto.CompactTextString(m) }
func (*GuestInfoReply) ProtoMessage()               {}
func (*GuestInfoReply) Descriptor() ([]byte, []int) { return fileDescriptorGuest, []int{2} }

func init() {
	proto.RegisterType((*GuestInfoRequest)(nil), "ticket.service.item.v1.GuestInfoRequest")
	proto.RegisterType((*GuestStatusRequest)(nil), "ticket.service.item.v1.GuestStatusRequest")
	proto.RegisterType((*GuestInfoReply)(nil), "ticket.service.item.v1.GuestInfoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Guest service

type GuestClient interface {
	GuestInfo(ctx context.Context, in *GuestInfoRequest, opts ...grpc.CallOption) (*GuestInfoReply, error)
	GuestStatus(ctx context.Context, in *GuestStatusRequest, opts ...grpc.CallOption) (*GuestInfoReply, error)
}

type guestClient struct {
	cc *grpc.ClientConn
}

func NewGuestClient(cc *grpc.ClientConn) GuestClient {
	return &guestClient{cc}
}

func (c *guestClient) GuestInfo(ctx context.Context, in *GuestInfoRequest, opts ...grpc.CallOption) (*GuestInfoReply, error) {
	out := new(GuestInfoReply)
	err := grpc.Invoke(ctx, "/ticket.service.item.v1.Guest/GuestInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestClient) GuestStatus(ctx context.Context, in *GuestStatusRequest, opts ...grpc.CallOption) (*GuestInfoReply, error) {
	out := new(GuestInfoReply)
	err := grpc.Invoke(ctx, "/ticket.service.item.v1.Guest/GuestStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Guest service

type GuestServer interface {
	GuestInfo(context.Context, *GuestInfoRequest) (*GuestInfoReply, error)
	GuestStatus(context.Context, *GuestStatusRequest) (*GuestInfoReply, error)
}

func RegisterGuestServer(s *grpc.Server, srv GuestServer) {
	s.RegisterService(&_Guest_serviceDesc, srv)
}

func _Guest_GuestInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServer).GuestInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.service.item.v1.Guest/GuestInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServer).GuestInfo(ctx, req.(*GuestInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Guest_GuestStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServer).GuestStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket.service.item.v1.Guest/GuestStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServer).GuestStatus(ctx, req.(*GuestStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Guest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ticket.service.item.v1.Guest",
	HandlerType: (*GuestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GuestInfo",
			Handler:    _Guest_GuestInfo_Handler,
		},
		{
			MethodName: "GuestStatus",
			Handler:    _Guest_GuestStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/service/openplatform/ticket-item/api/grpc/v1/guest.proto",
}

func (m *GuestInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuestInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGuest(dAtA, i, uint64(m.ID))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGuest(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.GuestImg) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintGuest(dAtA, i, uint64(len(m.GuestImg)))
		i += copy(dAtA[i:], m.GuestImg)
	}
	if len(m.Description) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGuest(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	}
	if m.GuestID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintGuest(dAtA, i, uint64(m.GuestID))
	}
	return i, nil
}

func (m *GuestStatusRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuestStatusRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGuest(dAtA, i, uint64(m.ID))
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGuest(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *GuestInfoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GuestInfoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintGuest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GuestInfoRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovGuest(uint64(m.ID))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGuest(uint64(l))
	}
	l = len(m.GuestImg)
	if l > 0 {
		n += 1 + l + sovGuest(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGuest(uint64(l))
	}
	if m.GuestID != 0 {
		n += 1 + sovGuest(uint64(m.GuestID))
	}
	return n
}

func (m *GuestStatusRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovGuest(uint64(m.ID))
	}
	if m.Status != 0 {
		n += 1 + sovGuest(uint64(m.Status))
	}
	return n
}

func (m *GuestInfoReply) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovGuest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGuest(x uint64) (n int) {
	return sovGuest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GuestInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GuestInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuestInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuestImg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GuestImg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGuest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GuestID", wireType)
			}
			m.GuestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GuestID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GuestStatusRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GuestStatusRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuestStatusRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGuest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GuestInfoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGuest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GuestInfoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GuestInfoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGuest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGuest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGuest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGuest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGuest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGuest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGuest
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGuest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGuest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGuest   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("app/service/openplatform/ticket-item/api/grpc/v1/guest.proto", fileDescriptorGuest)
}

var fileDescriptorGuest = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x8e, 0x9d, 0xa6, 0x49, 0x36, 0x02, 0x95, 0x3d, 0x80, 0xa9, 0x84, 0x37, 0x58, 0x2a, 0x0a,
	0x88, 0x7a, 0xeb, 0x54, 0xe2, 0x4f, 0x94, 0x83, 0x85, 0x54, 0xe5, 0x6a, 0x6e, 0x5c, 0x2a, 0xc7,
	0xde, 0x98, 0x15, 0x71, 0x76, 0x6b, 0xaf, 0x2d, 0xfa, 0x04, 0xbc, 0x10, 0x0f, 0xd1, 0x0b, 0x52,
	0x9f, 0x60, 0x45, 0x7d, 0xf4, 0x31, 0x4f, 0x80, 0xbc, 0x4e, 0x52, 0x37, 0xaa, 0xaa, 0xde, 0x76,
	0x66, 0xbe, 0xf9, 0x66, 0xbe, 0x6f, 0x07, 0x7c, 0xf6, 0x39, 0xc7, 0x29, 0x49, 0x72, 0x1a, 0x10,
	0xcc, 0x38, 0x59, 0xf0, 0xb9, 0x2f, 0x66, 0x2c, 0x89, 0xb1, 0xa0, 0xc1, 0x4f, 0x22, 0x0e, 0xa9,
	0x20, 0x31, 0xf6, 0x39, 0xc5, 0x51, 0xc2, 0x03, 0x9c, 0x3b, 0x38, 0xca, 0x48, 0x2a, 0x6c, 0x9e,
	0x30, 0xc1, 0xe0, 0xd3, 0x1a, 0x64, 0xaf, 0x08, 0xec, 0x0a, 0x6c, 0xe7, 0xce, 0xfe, 0x61, 0x44,
	0xc5, 0x8f, 0x6c, 0x6a, 0x07, 0x2c, 0xc6, 0x11, 0x8b, 0x18, 0x56, 0xf0, 0x69, 0x36, 0x53, 0x91,
	0x0a, 0xd4, 0xab, 0xa6, 0xb1, 0xfe, 0xe8, 0x60, 0xef, 0xb4, 0xa2, 0x9d, 0x2c, 0x66, 0xcc, 0x23,
	0xe7, 0xd5, 0x13, 0x1e, 0x01, 0x9d, 0x86, 0x86, 0x36, 0xd4, 0x46, 0x8f, 0xdc, 0x61, 0x21, 0x91,
	0x3e, 0xf9, 0x5a, 0x4a, 0xa4, 0xd3, 0x70, 0x29, 0xd1, 0x5e, 0xee, 0xcf, 0x69, 0xe8, 0x0b, 0xf2,
	0xc9, 0x8a, 0xe9, 0xe2, 0xe4, 0xc8, 0xf2, 0x74, 0x1a, 0x42, 0x07, 0xec, 0x2c, 0xfc, 0x98, 0x18,
	0xfa, 0x50, 0x1b, 0xf5, 0xdd, 0x17, 0xa5, 0x44, 0x2a, 0x5e, 0x4a, 0xf4, 0xa4, 0x81, 0xf7, 0x7f,
	0x9d, 0x1c, 0x8f, 0x2d, 0x4f, 0x95, 0xe0, 0x17, 0xd0, 0x57, 0x7a, 0xce, 0x68, 0x1c, 0x19, 0x6d,
	0xd5, 0xf7, 0xb2, 0x94, 0xe8, 0x26, 0x79, 0xe7, 0xb0, 0x9e, 0x2a, 0x4f, 0xe2, 0x08, 0x9e, 0x82,
	0x41, 0x48, 0xd2, 0x20, 0xa1, 0x5c, 0x50, 0xb6, 0x30, 0x76, 0x14, 0xc3, 0x41, 0x29, 0x51, 0x33,
	0xbd, 0x94, 0x08, 0xde, 0x5e, 0xc0, 0x19, 0x7f, 0xb0, 0xbc, 0x26, 0x04, 0x8e, 0x41, 0x6f, 0x35,
	0x33, 0x34, 0x3a, 0x43, 0x6d, 0xd4, 0x76, 0x9f, 0x15, 0x12, 0x75, 0x6b, 0x57, 0x2a, 0xe1, 0x9b,
	0xb2, 0xd7, 0xad, 0xa7, 0x87, 0xd6, 0x6f, 0x0d, 0x40, 0x05, 0xf8, 0x26, 0x7c, 0x91, 0xa5, 0x6b,
	0xe3, 0x3e, 0x6e, 0x8c, 0x6b, 0xbb, 0xaf, 0xb7, 0x8c, 0x7b, 0x7e, 0xb3, 0x47, 0x42, 0xce, 0x33,
	0x9a, 0x90, 0xf0, 0x6d, 0x25, 0xca, 0xa9, 0x1d, 0x7c, 0x07, 0x76, 0x53, 0xc5, 0xa5, 0x3c, 0xec,
	0xb8, 0x66, 0x29, 0xd1, 0x2a, 0xb3, 0x65, 0x44, 0x25, 0xc2, 0xf2, 0x56, 0x35, 0xeb, 0x3d, 0x78,
	0xdc, 0xf8, 0x3f, 0x3e, 0xbf, 0x80, 0x07, 0xa0, 0x9b, 0x66, 0x41, 0x40, 0xd2, 0x54, 0x6d, 0xd2,
	0x73, 0x07, 0xa5, 0x44, 0xeb, 0x94, 0xb7, 0x7e, 0x8c, 0xff, 0x6a, 0xa0, 0xa3, 0x3a, 0xe1, 0x19,
	0xe8, 0x6f, 0x28, 0xe0, 0xc8, 0xbe, 0xfb, 0xb0, 0xec, 0xed, 0x2b, 0xd9, 0x7f, 0xf5, 0x00, 0x24,
	0x9f, 0x5f, 0x58, 0x2d, 0x18, 0x80, 0x41, 0xc3, 0x2c, 0xf8, 0xe6, 0xde, 0xc6, 0x5b, 0x8e, 0x3e,
	0x7c, 0x88, 0x6b, 0x5c, 0x5e, 0x9b, 0xad, 0xab, 0x6b, 0xb3, 0x75, 0x59, 0x98, 0xda, 0x55, 0x61,
	0x6a, 0xff, 0x0a, 0x53, 0xfb, 0xae, 0xe7, 0xce, 0x74, 0x57, 0x9d, 0xfa, 0xf1, 0xff, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x44, 0x78, 0x81, 0x56, 0x71, 0x03, 0x00, 0x00,
}

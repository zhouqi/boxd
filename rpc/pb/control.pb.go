// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: control.proto

package rpcpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing debug level.
type DebugLevelRequest struct {
	Level string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (m *DebugLevelRequest) Reset()         { *m = DebugLevelRequest{} }
func (m *DebugLevelRequest) String() string { return proto.CompactTextString(m) }
func (*DebugLevelRequest) ProtoMessage()    {}
func (*DebugLevelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_7e77d2d8349cbed8, []int{0}
}
func (m *DebugLevelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DebugLevelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DebugLevelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *DebugLevelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugLevelRequest.Merge(dst, src)
}
func (m *DebugLevelRequest) XXX_Size() int {
	return m.Size()
}
func (m *DebugLevelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugLevelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DebugLevelRequest proto.InternalMessageInfo

func (m *DebugLevelRequest) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

// The response message containing the result message
type Reply struct {
	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_control_7e77d2d8349cbed8, []int{1}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return m.Size()
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*DebugLevelRequest)(nil), "rpcpb.DebugLevelRequest")
	proto.RegisterType((*Reply)(nil), "rpcpb.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContorlCommandClient is the client API for ContorlCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContorlCommandClient interface {
	// set boxd debug level
	SetDebugLevel(ctx context.Context, in *DebugLevelRequest, opts ...grpc.CallOption) (*Reply, error)
}

type contorlCommandClient struct {
	cc *grpc.ClientConn
}

func NewContorlCommandClient(cc *grpc.ClientConn) ContorlCommandClient {
	return &contorlCommandClient{cc}
}

func (c *contorlCommandClient) SetDebugLevel(ctx context.Context, in *DebugLevelRequest, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/rpcpb.ContorlCommand/SetDebugLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContorlCommandServer is the server API for ContorlCommand service.
type ContorlCommandServer interface {
	// set boxd debug level
	SetDebugLevel(context.Context, *DebugLevelRequest) (*Reply, error)
}

func RegisterContorlCommandServer(s *grpc.Server, srv ContorlCommandServer) {
	s.RegisterService(&_ContorlCommand_serviceDesc, srv)
}

func _ContorlCommand_SetDebugLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebugLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContorlCommandServer).SetDebugLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.ContorlCommand/SetDebugLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContorlCommandServer).SetDebugLevel(ctx, req.(*DebugLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContorlCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.ContorlCommand",
	HandlerType: (*ContorlCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetDebugLevel",
			Handler:    _ContorlCommand_SetDebugLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}

func (m *DebugLevelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DebugLevelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Level) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Level)))
		i += copy(dAtA[i:], m.Level)
	}
	return i, nil
}

func (m *Reply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.Code))
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	return i, nil
}

func encodeVarintControl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DebugLevelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Level)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}

func (m *Reply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovControl(uint64(m.Code))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	return n
}

func sovControl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozControl(x uint64) (n int) {
	return sovControl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DebugLevelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
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
			return fmt.Errorf("proto: DebugLevelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DebugLevelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Level = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
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
func (m *Reply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
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
			return fmt.Errorf("proto: Reply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
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
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
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
func skipControl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
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
				return 0, ErrInvalidLengthControl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowControl
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
				next, err := skipControl(dAtA[start:])
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
	ErrInvalidLengthControl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControl   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("control.proto", fileDescriptor_control_7e77d2d8349cbed8) }

var fileDescriptor_control_7e77d2d8349cbed8 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2a, 0x48, 0x2e, 0x48,
	0x92, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x52, 0xd2, 0xe4, 0x12, 0x74, 0x49,
	0x4d, 0x2a, 0x4d, 0xf7, 0x49, 0x2d, 0x4b, 0xcd, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe1, 0x62, 0xcd, 0x01, 0xf1, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25,
	0x53, 0x2e, 0xd6, 0xa0, 0xd4, 0x82, 0x9c, 0x4a, 0x21, 0x21, 0x2e, 0x96, 0xe4, 0xfc, 0x94, 0x54,
	0xb0, 0x2c, 0x6b, 0x10, 0x98, 0x2d, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e,
	0x2a, 0xc1, 0x04, 0xd6, 0x04, 0xe3, 0x1a, 0x65, 0x70, 0xf1, 0x39, 0xe7, 0xe7, 0x95, 0xe4, 0x17,
	0xe5, 0x38, 0xe7, 0xe7, 0xe6, 0x26, 0xe6, 0xa5, 0x08, 0x85, 0x71, 0xf1, 0x06, 0xa7, 0x96, 0x20,
	0xac, 0x15, 0x92, 0xd0, 0x03, 0x3b, 0x55, 0x0f, 0xc3, 0x25, 0x52, 0x3c, 0x50, 0x19, 0xb0, 0xc5,
	0x4a, 0xb2, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12, 0x57, 0x12, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x2e,
	0xc9, 0xd1, 0x4f, 0x01, 0x69, 0x00, 0xbb, 0xce, 0x8a, 0x51, 0xcb, 0x49, 0xe2, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x9e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x0e, 0xe0, 0x9a, 0x2c, 0x22, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/kata-containers/kata-containers/src/libs/protocols/protos/portforward.proto

package grpc

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PortForwardRequest struct {
	ContainerId          string   `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	VsockPort            uint32   `protobuf:"varint,3,opt,name=vsock_port,json=vsockPort,proto3" json:"vsock_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortForwardRequest) Reset()      { *m = PortForwardRequest{} }
func (*PortForwardRequest) ProtoMessage() {}
func (*PortForwardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d3f7de0f8f0b0a9, []int{0}
}
func (m *PortForwardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PortForwardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PortForwardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PortForwardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortForwardRequest.Merge(m, src)
}
func (m *PortForwardRequest) XXX_Size() int {
	return m.Size()
}
func (m *PortForwardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PortForwardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PortForwardRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PortForwardRequest)(nil), "grpc.PortForwardRequest")
}

func init() {
	proto.RegisterFile("github.com/kata-containers/kata-containers/src/libs/protocols/protos/portforward.proto", fileDescriptor_9d3f7de0f8f0b0a9)
}

var fileDescriptor_9d3f7de0f8f0b0a9 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0xad, 0xbf, 0xaf, 0x42, 0xaa, 0x0b, 0x8b, 0x07, 0x14, 0x15, 0x61, 0x15, 0xa6, 0x2e, 0xd8,
	0x12, 0x3c, 0x00, 0x12, 0x12, 0x48, 0x6c, 0x28, 0x08, 0x06, 0x06, 0xaa, 0xc4, 0x71, 0x8d, 0xc9,
	0x8f, 0x83, 0x7d, 0x13, 0x60, 0xe3, 0xf1, 0x3a, 0x32, 0x32, 0xd2, 0x3c, 0x09, 0x8a, 0x03, 0x55,
	0x50, 0x27, 0xb6, 0x7b, 0xcf, 0xb9, 0x57, 0xe7, 0x07, 0xdf, 0x2a, 0x0d, 0x0f, 0x55, 0xcc, 0x84,
	0xc9, 0x79, 0x1a, 0x41, 0x74, 0x24, 0x4c, 0x01, 0x91, 0x2e, 0xa4, 0x75, 0x1b, 0xbb, 0xb3, 0x82,
	0x67, 0x3a, 0x76, 0xbc, 0xb4, 0x06, 0x8c, 0x30, 0xd9, 0xf7, 0xe4, 0x78, 0x69, 0x2c, 0x2c, 0x8c,
	0x7d, 0x8e, 0x6c, 0xc2, 0x3c, 0x44, 0x86, 0xca, 0x96, 0x62, 0xb2, 0xa7, 0x8c, 0x51, 0x99, 0xec,
	0xce, 0xe2, 0x6a, 0xc1, 0x65, 0x5e, 0xc2, 0x6b, 0x77, 0x72, 0xf8, 0x88, 0xc9, 0x95, 0xb1, 0x70,
	0xd1, 0xfd, 0x85, 0xf2, 0xa9, 0x92, 0x0e, 0xc8, 0x01, 0xde, 0x5e, 0x4b, 0xce, 0x75, 0x12, 0xa0,
	0x29, 0x9a, 0x8d, 0xc2, 0xf1, 0x1a, 0xbb, 0x4c, 0x08, 0xc1, 0xc3, 0x56, 0x30, 0xf8, 0x37, 0x45,
	0xb3, 0x9d, 0xd0, 0xcf, 0x64, 0x1f, 0xe3, 0xda, 0x19, 0x91, 0xce, 0x3d, 0xf3, 0xdf, 0x33, 0x23,
	0x8f, 0xb4, 0x1a, 0xc7, 0x37, 0xbf, 0xb4, 0xae, 0xa5, 0xad, 0xb5, 0x90, 0xe4, 0x14, 0x8f, 0x7b,
	0x28, 0x09, 0x58, 0x6b, 0x9a, 0x6d, 0x9a, 0x9a, 0xec, 0xb2, 0x2e, 0x08, 0xfb, 0x09, 0xc2, 0xce,
	0xdb, 0x20, 0x67, 0x2f, 0xcb, 0x15, 0x1d, 0x7c, 0xac, 0xe8, 0xe0, 0xad, 0xa1, 0x68, 0xd9, 0x50,
	0xf4, 0xde, 0x50, 0xf4, 0xd9, 0x50, 0x74, 0x77, 0xff, 0xc7, 0x5e, 0x6d, 0x55, 0x80, 0xce, 0x25,
	0xaf, 0xb5, 0x85, 0x1e, 0x55, 0xa6, 0x8a, 0x47, 0x4a, 0x16, 0xd0, 0xeb, 0xbc, 0x35, 0x19, 0x6f,
	0xf9, 0xfd, 0xe4, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x29, 0x47, 0xcb, 0x0b, 0xc0, 0x01, 0x00, 0x00,
}

func (m *PortForwardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PortForwardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PortForwardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VsockPort != 0 {
		i = encodeVarintPortforward(dAtA, i, uint64(m.VsockPort))
		i--
		dAtA[i] = 0x18
	}
	if m.Port != 0 {
		i = encodeVarintPortforward(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContainerId) > 0 {
		i -= len(m.ContainerId)
		copy(dAtA[i:], m.ContainerId)
		i = encodeVarintPortforward(dAtA, i, uint64(len(m.ContainerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPortforward(dAtA []byte, offset int, v uint64) int {
	offset -= sovPortforward(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PortForwardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContainerId)
	if l > 0 {
		n += 1 + l + sovPortforward(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovPortforward(uint64(m.Port))
	}
	if m.VsockPort != 0 {
		n += 1 + sovPortforward(uint64(m.VsockPort))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPortforward(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPortforward(x uint64) (n int) {
	return sovPortforward(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PortForwardRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PortForwardRequest{`,
		`ContainerId:` + fmt.Sprintf("%v", this.ContainerId) + `,`,
		`Port:` + fmt.Sprintf("%v", this.Port) + `,`,
		`VsockPort:` + fmt.Sprintf("%v", this.VsockPort) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPortforward(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}

type PortForwardServiceService interface {
	PortForward(ctx context.Context, req *PortForwardRequest) (*types.Empty, error)
}

func RegisterPortForwardServiceService(srv *github_com_containerd_ttrpc.Server, svc PortForwardServiceService) {
	srv.Register("grpc.PortForwardService", map[string]github_com_containerd_ttrpc.Method{
		"PortForward": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req PortForwardRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.PortForward(ctx, &req)
		},
	})
}

type portForwardServiceClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewPortForwardServiceClient(client *github_com_containerd_ttrpc.Client) PortForwardServiceService {
	return &portForwardServiceClient{
		client: client,
	}
}

func (c *portForwardServiceClient) PortForward(ctx context.Context, req *PortForwardRequest) (*types.Empty, error) {
	var resp types.Empty
	if err := c.client.Call(ctx, "grpc.PortForwardService", "PortForward", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
func (m *PortForwardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPortforward
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PortForwardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PortForwardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortforward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPortforward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPortforward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortforward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VsockPort", wireType)
			}
			m.VsockPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPortforward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VsockPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPortforward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPortforward
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPortforward(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPortforward
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
					return 0, ErrIntOverflowPortforward
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPortforward
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
			if length < 0 {
				return 0, ErrInvalidLengthPortforward
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPortforward
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPortforward
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPortforward        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPortforward          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPortforward = fmt.Errorf("proto: unexpected end of group")
)

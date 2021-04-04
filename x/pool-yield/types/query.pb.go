// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/pool-yield/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryFarmIdsRequest struct {
	PoolId uint64 `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty" yaml:"pool_id"`
}

func (m *QueryFarmIdsRequest) Reset()         { *m = QueryFarmIdsRequest{} }
func (m *QueryFarmIdsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFarmIdsRequest) ProtoMessage()    {}
func (*QueryFarmIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e128144e22b6aa3, []int{0}
}
func (m *QueryFarmIdsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFarmIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFarmIdsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFarmIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFarmIdsRequest.Merge(m, src)
}
func (m *QueryFarmIdsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFarmIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFarmIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFarmIdsRequest proto.InternalMessageInfo

func (m *QueryFarmIdsRequest) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

type QueryFarmIdsResponse struct {
	FarmIdsWithDuration []*QueryFarmIdsResponse_FarmIdWithDuration `protobuf:"bytes,1,rep,name=farmIdsWithDuration,proto3" json:"farmIdsWithDuration,omitempty"`
}

func (m *QueryFarmIdsResponse) Reset()         { *m = QueryFarmIdsResponse{} }
func (m *QueryFarmIdsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFarmIdsResponse) ProtoMessage()    {}
func (*QueryFarmIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e128144e22b6aa3, []int{1}
}
func (m *QueryFarmIdsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFarmIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFarmIdsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFarmIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFarmIdsResponse.Merge(m, src)
}
func (m *QueryFarmIdsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFarmIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFarmIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFarmIdsResponse proto.InternalMessageInfo

func (m *QueryFarmIdsResponse) GetFarmIdsWithDuration() []*QueryFarmIdsResponse_FarmIdWithDuration {
	if m != nil {
		return m.FarmIdsWithDuration
	}
	return nil
}

type QueryFarmIdsResponse_FarmIdWithDuration struct {
	FarmId   uint64        `protobuf:"varint,1,opt,name=farmId,proto3" json:"farmId,omitempty" yaml:"farm_id"`
	Duration time.Duration `protobuf:"bytes,2,opt,name=duration,proto3,stdduration" json:"duration"`
}

func (m *QueryFarmIdsResponse_FarmIdWithDuration) Reset() {
	*m = QueryFarmIdsResponse_FarmIdWithDuration{}
}
func (m *QueryFarmIdsResponse_FarmIdWithDuration) String() string { return proto.CompactTextString(m) }
func (*QueryFarmIdsResponse_FarmIdWithDuration) ProtoMessage()    {}
func (*QueryFarmIdsResponse_FarmIdWithDuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e128144e22b6aa3, []int{1, 0}
}
func (m *QueryFarmIdsResponse_FarmIdWithDuration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFarmIdsResponse_FarmIdWithDuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFarmIdsResponse_FarmIdWithDuration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFarmIdsResponse_FarmIdWithDuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFarmIdsResponse_FarmIdWithDuration.Merge(m, src)
}
func (m *QueryFarmIdsResponse_FarmIdWithDuration) XXX_Size() int {
	return m.Size()
}
func (m *QueryFarmIdsResponse_FarmIdWithDuration) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFarmIdsResponse_FarmIdWithDuration.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFarmIdsResponse_FarmIdWithDuration proto.InternalMessageInfo

func (m *QueryFarmIdsResponse_FarmIdWithDuration) GetFarmId() uint64 {
	if m != nil {
		return m.FarmId
	}
	return 0
}

func (m *QueryFarmIdsResponse_FarmIdWithDuration) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryFarmIdsRequest)(nil), "osmosis.poolyield.v1beta1.QueryFarmIdsRequest")
	proto.RegisterType((*QueryFarmIdsResponse)(nil), "osmosis.poolyield.v1beta1.QueryFarmIdsResponse")
	proto.RegisterType((*QueryFarmIdsResponse_FarmIdWithDuration)(nil), "osmosis.poolyield.v1beta1.QueryFarmIdsResponse.FarmIdWithDuration")
}

func init() {
	proto.RegisterFile("osmosis/pool-yield/v1beta1/query.proto", fileDescriptor_4e128144e22b6aa3)
}

var fileDescriptor_4e128144e22b6aa3 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0xeb, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0x55, 0x6b, 0xb9, 0x82, 0xc3, 0xb5, 0x43, 0x1b, 0x24, 0x2d, 0x19, 0xa4, 0x08,
	0xb9, 0xa3, 0x15, 0x17, 0x17, 0x31, 0x88, 0xd0, 0xd1, 0x2c, 0x82, 0x8b, 0x24, 0xcd, 0x35, 0x3d,
	0x48, 0x72, 0x69, 0xee, 0x22, 0x06, 0x71, 0xd0, 0x4f, 0x50, 0x70, 0x71, 0x75, 0xf6, 0x8b, 0x74,
	0x2c, 0xb8, 0x38, 0x55, 0x69, 0xfd, 0x04, 0x7e, 0x02, 0xb9, 0x5c, 0x22, 0x95, 0x56, 0xf9, 0xff,
	0xa7, 0xe4, 0xde, 0xf7, 0x77, 0xcf, 0xcb, 0xf3, 0xbc, 0x07, 0xef, 0x71, 0x91, 0x70, 0xc1, 0x04,
	0xc9, 0x38, 0x8f, 0x9d, 0x92, 0xd1, 0x38, 0x24, 0xaf, 0xa7, 0x01, 0x95, 0xfe, 0x94, 0xac, 0x0b,
	0x9a, 0x97, 0x38, 0xcb, 0xb9, 0xe4, 0x68, 0x58, 0x73, 0x58, 0x71, 0x15, 0x86, 0x6b, 0xcc, 0xec,
	0x47, 0x3c, 0xe2, 0x15, 0x45, 0xd4, 0x9f, 0xbe, 0x60, 0xde, 0x8d, 0x38, 0x8f, 0x62, 0x4a, 0xfc,
	0x8c, 0x11, 0x3f, 0x4d, 0xb9, 0xf4, 0x25, 0xe3, 0xa9, 0xa8, 0xbb, 0x56, 0xdd, 0xad, 0x4e, 0x41,
	0xb1, 0x24, 0x61, 0x91, 0x57, 0x80, 0xee, 0xdb, 0x4f, 0x60, 0xef, 0xb9, 0x9a, 0xfe, 0xcc, 0xcf,
	0x93, 0x79, 0x28, 0x3c, 0xba, 0x2e, 0xa8, 0x90, 0xe8, 0x3e, 0x6c, 0xab, 0xf9, 0xf3, 0x70, 0x00,
	0xc6, 0x60, 0x72, 0xd3, 0x45, 0xbf, 0xf6, 0xa3, 0x3b, 0xa5, 0x9f, 0xc4, 0x8f, 0x6c, 0x55, 0x7f,
	0xc5, 0x42, 0xdb, 0xab, 0x09, 0x7b, 0xd3, 0x82, 0xfd, 0xbf, 0x35, 0x44, 0xc6, 0x53, 0x41, 0x91,
	0x84, 0xbd, 0xa5, 0x2e, 0xbd, 0x60, 0x72, 0xf5, 0xb4, 0x1e, 0x3c, 0x00, 0xe3, 0x1b, 0x93, 0xee,
	0xcc, 0xc5, 0xff, 0x34, 0x8a, 0x2f, 0xa9, 0x61, 0x7d, 0x3e, 0x55, 0xf2, 0x2e, 0xc9, 0x9b, 0xef,
	0x01, 0x44, 0xe7, 0xac, 0x72, 0xa4, 0xe9, 0x73, 0x47, 0xaa, 0xae, 0x1d, 0x69, 0x02, 0x3d, 0x86,
	0x9d, 0x26, 0xa6, 0x41, 0x6b, 0x0c, 0x26, 0xdd, 0xd9, 0x10, 0xeb, 0x1c, 0x71, 0x93, 0x23, 0x6e,
	0x84, 0xdd, 0xce, 0x76, 0x3f, 0x32, 0x3e, 0x7d, 0x1f, 0x01, 0xef, 0xcf, 0xa5, 0xd9, 0x17, 0x00,
	0x6f, 0x55, 0x26, 0xd0, 0x67, 0x00, 0x6f, 0xd7, 0x4e, 0x10, 0xbe, 0xb2, 0xe5, 0x6a, 0x09, 0x26,
	0xb9, 0x66, 0x44, 0xf6, 0xc3, 0x0f, 0x5f, 0x7f, 0x7e, 0x6c, 0x11, 0xe4, 0x90, 0xff, 0x3c, 0x36,
	0xe5, 0xd1, 0x61, 0xa1, 0x20, 0x6f, 0xf5, 0xfe, 0xde, 0xb9, 0xf3, 0xed, 0xc1, 0x02, 0xbb, 0x83,
	0x05, 0x7e, 0x1c, 0x2c, 0xb0, 0x39, 0x5a, 0xc6, 0xee, 0x68, 0x19, 0xdf, 0x8e, 0x96, 0xf1, 0x92,
	0x44, 0x4c, 0xae, 0x8a, 0x00, 0x2f, 0x78, 0x42, 0x16, 0x4e, 0x23, 0xda, 0x7c, 0xdf, 0x9c, 0xca,
	0xcb, 0x32, 0xa3, 0x22, 0x68, 0x57, 0xf9, 0x3c, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xed,
	0x7c, 0xb4, 0xee, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// FarmIds takes the pool id and returns the matching farm ids and durations
	FarmIds(ctx context.Context, in *QueryFarmIdsRequest, opts ...grpc.CallOption) (*QueryFarmIdsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) FarmIds(ctx context.Context, in *QueryFarmIdsRequest, opts ...grpc.CallOption) (*QueryFarmIdsResponse, error) {
	out := new(QueryFarmIdsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.poolyield.v1beta1.Query/FarmIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// FarmIds takes the pool id and returns the matching farm ids and durations
	FarmIds(context.Context, *QueryFarmIdsRequest) (*QueryFarmIdsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) FarmIds(ctx context.Context, req *QueryFarmIdsRequest) (*QueryFarmIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FarmIds not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_FarmIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFarmIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FarmIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.poolyield.v1beta1.Query/FarmIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FarmIds(ctx, req.(*QueryFarmIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.poolyield.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FarmIds",
			Handler:    _Query_FarmIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/pool-yield/v1beta1/query.proto",
}

func (m *QueryFarmIdsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFarmIdsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFarmIdsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryFarmIdsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFarmIdsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFarmIdsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FarmIdsWithDuration) > 0 {
		for iNdEx := len(m.FarmIdsWithDuration) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FarmIdsWithDuration[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryFarmIdsResponse_FarmIdWithDuration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFarmIdsResponse_FarmIdWithDuration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFarmIdsResponse_FarmIdWithDuration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintQuery(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.FarmId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.FarmId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryFarmIdsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovQuery(uint64(m.PoolId))
	}
	return n
}

func (m *QueryFarmIdsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FarmIdsWithDuration) > 0 {
		for _, e := range m.FarmIdsWithDuration {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryFarmIdsResponse_FarmIdWithDuration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FarmId != 0 {
		n += 1 + sovQuery(uint64(m.FarmId))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryFarmIdsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFarmIdsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFarmIdsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFarmIdsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryFarmIdsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFarmIdsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmIdsWithDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FarmIdsWithDuration = append(m.FarmIdsWithDuration, &QueryFarmIdsResponse_FarmIdWithDuration{})
			if err := m.FarmIdsWithDuration[len(m.FarmIdsWithDuration)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryFarmIdsResponse_FarmIdWithDuration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: FarmIdWithDuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FarmIdWithDuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmId", wireType)
			}
			m.FarmId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FarmId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
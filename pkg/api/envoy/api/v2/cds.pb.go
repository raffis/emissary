// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/cds.proto

package envoy_api_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/datawire/ambassador/pkg/api/envoy/annotations"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
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

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221 and protoxform to upgrade the file.
type CdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdsDummy) Reset()         { *m = CdsDummy{} }
func (m *CdsDummy) String() string { return proto.CompactTextString(m) }
func (*CdsDummy) ProtoMessage()    {}
func (*CdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e73f50fbb1daa302, []int{0}
}
func (m *CdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdsDummy.Merge(m, src)
}
func (m *CdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *CdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_CdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_CdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CdsDummy)(nil), "envoy.api.v2.CdsDummy")
}

func init() { proto.RegisterFile("envoy/api/v2/cds.proto", fileDescriptor_e73f50fbb1daa302) }

var fileDescriptor_e73f50fbb1daa302 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0x7b, 0x95, 0x56, 0x39, 0xac, 0x48, 0x90, 0x5a, 0x8e, 0x1a, 0xa5, 0x0a, 0x8a, 0x43,
	0x22, 0xed, 0xd6, 0xb1, 0x2d, 0xce, 0xc5, 0x0e, 0x8e, 0x7a, 0x26, 0x2f, 0x35, 0xd0, 0xe4, 0xe2,
	0xdd, 0x25, 0x98, 0x4d, 0x9c, 0xc4, 0xc5, 0x41, 0x10, 0x3f, 0x80, 0x9f, 0xc2, 0x4f, 0xd0, 0x51,
	0xf1, 0x0b, 0x48, 0x71, 0x75, 0xd1, 0x5d, 0xa4, 0xf9, 0x53, 0x3d, 0x8b, 0x4e, 0xae, 0xf7, 0x7b,
	0xde, 0xe7, 0x9e, 0xf7, 0x7d, 0x70, 0x19, 0xbc, 0x90, 0x45, 0x26, 0xf5, 0x1d, 0x33, 0xac, 0x9b,
	0x96, 0x2d, 0x0c, 0x9f, 0x33, 0xc9, 0xb4, 0xf9, 0xf8, 0xdd, 0xa0, 0xbe, 0x63, 0x84, 0x75, 0x52,
	0x55, 0x54, 0xb6, 0x23, 0x2c, 0x16, 0x02, 0x8f, 0x12, 0x2d, 0xa9, 0xf6, 0x19, 0xeb, 0x0f, 0x20,
	0xc6, 0xd4, 0xf3, 0x98, 0xa4, 0xd2, 0x61, 0x5e, 0xea, 0x44, 0xd6, 0xd2, 0xd9, 0x2f, 0x60, 0x72,
	0x10, 0x2c, 0xe0, 0x16, 0xa4, 0x0a, 0x3d, 0xb0, 0x7d, 0xaa, 0x08, 0x5c, 0xa7, 0xcf, 0xa9, 0xcc,
	0xf8, 0xca, 0x14, 0x17, 0x92, 0xca, 0x20, 0xfb, 0x80, 0xa8, 0x2b, 0x0c, 0x02, 0x21, 0x81, 0x27,
	0xac, 0x86, 0xf1, 0x5c, 0xdb, 0x16, 0x9d, 0xc0, 0x75, 0xa3, 0xfa, 0x6b, 0x1e, 0x2f, 0xb7, 0x13,
	0xda, 0xc9, 0x36, 0xe8, 0x01, 0x0f, 0x1d, 0x0b, 0xb4, 0x7d, 0xbc, 0xd0, 0x93, 0x1c, 0xa8, 0x9b,
	0x0a, 0x84, 0xa6, 0x1b, 0xdf, 0x2f, 0x60, 0x4c, 0x26, 0xf6, 0xe0, 0x24, 0x00, 0x21, 0xc9, 0xea,
	0xaf, 0x5c, 0xf8, 0xcc, 0x13, 0x50, 0xcb, 0x6d, 0xa1, 0x1d, 0xa4, 0x1d, 0xe2, 0x52, 0x07, 0x06,
	0x92, 0x4e, 0x7c, 0xd7, 0x7f, 0xcc, 0x8d, 0xe1, 0x94, 0xf9, 0xc6, 0xdf, 0x22, 0xe5, 0x87, 0x08,
	0x97, 0x76, 0x41, 0x5a, 0xc7, 0xff, 0x97, 0x7c, 0xf3, 0xfc, 0xe9, 0xe5, 0x3a, 0x5f, 0xa9, 0x95,
	0x95, 0xb2, 0x9b, 0xe9, 0x65, 0x45, 0x4c, 0x67, 0x9a, 0x68, 0x9b, 0x54, 0x2f, 0xef, 0x6e, 0xde,
	0x67, 0xcb, 0x78, 0x49, 0x31, 0x4c, 0x83, 0xb4, 0x0e, 0x86, 0x23, 0x1d, 0x3d, 0x8c, 0x74, 0xf4,
	0x3c, 0xd2, 0xd1, 0xdb, 0xed, 0xc7, 0x55, 0x81, 0x68, 0x95, 0x44, 0x29, 0x92, 0xb3, 0x1b, 0x59,
	0x59, 0x61, 0xe3, 0xfe, 0x6c, 0xf8, 0x58, 0xcc, 0x2f, 0xe6, 0x30, 0x71, 0x58, 0x92, 0xcf, 0xe7,
	0xec, 0x34, 0x52, 0xa2, 0xb6, 0xc6, 0x5d, 0x76, 0xc7, 0xbd, 0x76, 0xd1, 0x05, 0x42, 0xdd, 0xc2,
	0x51, 0x31, 0x6e, 0xb9, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x12, 0x9b, 0x45, 0xc6, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterDiscoveryServiceClient is the client API for ClusterDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterDiscoveryServiceClient interface {
	StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error)
	DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error)
	FetchClusters(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type clusterDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterDiscoveryServiceClient(cc *grpc.ClientConn) ClusterDiscoveryServiceClient {
	return &clusterDiscoveryServiceClient{cc}
}

func (c *clusterDiscoveryServiceClient) StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v2.ClusterDiscoveryService/StreamClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceStreamClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_StreamClustersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceStreamClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceStreamClustersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v2.ClusterDiscoveryService/DeltaClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceDeltaClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_DeltaClustersClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceDeltaClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) FetchClusters(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v2.ClusterDiscoveryService/FetchClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterDiscoveryServiceServer is the server API for ClusterDiscoveryService service.
type ClusterDiscoveryServiceServer interface {
	StreamClusters(ClusterDiscoveryService_StreamClustersServer) error
	DeltaClusters(ClusterDiscoveryService_DeltaClustersServer) error
	FetchClusters(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedClusterDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterDiscoveryServiceServer struct {
}

func (*UnimplementedClusterDiscoveryServiceServer) StreamClusters(srv ClusterDiscoveryService_StreamClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) DeltaClusters(srv ClusterDiscoveryService_DeltaClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) FetchClusters(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClusters not implemented")
}

func RegisterClusterDiscoveryServiceServer(s *grpc.Server, srv ClusterDiscoveryServiceServer) {
	s.RegisterService(&_ClusterDiscoveryService_serviceDesc, srv)
}

func _ClusterDiscoveryService_StreamClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).StreamClusters(&clusterDiscoveryServiceStreamClustersServer{stream})
}

type ClusterDiscoveryService_StreamClustersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceStreamClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceStreamClustersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_DeltaClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).DeltaClusters(&clusterDiscoveryServiceDeltaClustersServer{stream})
}

type ClusterDiscoveryService_DeltaClustersServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceDeltaClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_FetchClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ClusterDiscoveryService/FetchClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ClusterDiscoveryService",
	HandlerType: (*ClusterDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClusters",
			Handler:    _ClusterDiscoveryService_FetchClusters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClusters",
			Handler:       _ClusterDiscoveryService_StreamClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaClusters",
			Handler:       _ClusterDiscoveryService_DeltaClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v2/cds.proto",
}

func (m *CdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CdsDummy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CdsDummy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintCds(dAtA []byte, offset int, v uint64) int {
	offset -= sovCds(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CdsDummy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCds(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCds(x uint64) (n int) {
	return sovCds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCds
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
			return fmt.Errorf("proto: CdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCds
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCds
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
func skipCds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCds
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
					return 0, ErrIntOverflowCds
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
					return 0, ErrIntOverflowCds
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
				return 0, ErrInvalidLengthCds
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCds
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCds
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCds        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCds          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCds = fmt.Errorf("proto: unexpected end of group")
)

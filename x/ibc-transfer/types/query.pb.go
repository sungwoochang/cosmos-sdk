// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/transfer/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
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

// QueryDenomTraceRequest is the request type for the Query/DenomTrace RPC method
type QueryDenomTraceRequest struct {
	// hex hash of the denomination trace info
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *QueryDenomTraceRequest) Reset()         { *m = QueryDenomTraceRequest{} }
func (m *QueryDenomTraceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTraceRequest) ProtoMessage()    {}
func (*QueryDenomTraceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b3e8b4e9dff1c1, []int{0}
}
func (m *QueryDenomTraceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTraceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTraceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTraceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTraceRequest.Merge(m, src)
}
func (m *QueryDenomTraceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTraceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTraceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTraceRequest proto.InternalMessageInfo

func (m *QueryDenomTraceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// QueryDenomTraceResponse is the response type for the Query/DenomTrace RPC method.
type QueryDenomTraceResponse struct {
	// denomination trace associated with the request hash
	DenomTrace *DenomTrace `protobuf:"bytes,1,opt,name=denom_trace,json=denomTrace,proto3" json:"denom_trace,omitempty"`
}

func (m *QueryDenomTraceResponse) Reset()         { *m = QueryDenomTraceResponse{} }
func (m *QueryDenomTraceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTraceResponse) ProtoMessage()    {}
func (*QueryDenomTraceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b3e8b4e9dff1c1, []int{1}
}
func (m *QueryDenomTraceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTraceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTraceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTraceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTraceResponse.Merge(m, src)
}
func (m *QueryDenomTraceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTraceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTraceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTraceResponse proto.InternalMessageInfo

func (m *QueryDenomTraceResponse) GetDenomTrace() *DenomTrace {
	if m != nil {
		return m.DenomTrace
	}
	return nil
}

// QueryConnectionsRequest is the request type for the Query/Connections RPC method
type QueryDenomTracesRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDenomTracesRequest) Reset()         { *m = QueryDenomTracesRequest{} }
func (m *QueryDenomTracesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTracesRequest) ProtoMessage()    {}
func (*QueryDenomTracesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b3e8b4e9dff1c1, []int{2}
}
func (m *QueryDenomTracesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTracesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTracesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTracesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTracesRequest.Merge(m, src)
}
func (m *QueryDenomTracesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTracesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTracesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTracesRequest proto.InternalMessageInfo

func (m *QueryDenomTracesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryConnectionsResponse is the response type for the Query/Connections RPC method.
type QueryDenomTracesResponse struct {
	// list of stored connections of the chain.
	DenomTraces []*IdentifiedDenomTrace `protobuf:"bytes,1,rep,name=denom_traces,json=denomTraces,proto3" json:"denom_traces,omitempty"`
	// pagination response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDenomTracesResponse) Reset()         { *m = QueryDenomTracesResponse{} }
func (m *QueryDenomTracesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomTracesResponse) ProtoMessage()    {}
func (*QueryDenomTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b3e8b4e9dff1c1, []int{3}
}
func (m *QueryDenomTracesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomTracesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomTracesResponse.Merge(m, src)
}
func (m *QueryDenomTracesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomTracesResponse proto.InternalMessageInfo

func (m *QueryDenomTracesResponse) GetDenomTraces() []*IdentifiedDenomTrace {
	if m != nil {
		return m.DenomTraces
	}
	return nil
}

func (m *QueryDenomTracesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryDenomTraceRequest)(nil), "ibc.transfer.QueryDenomTraceRequest")
	proto.RegisterType((*QueryDenomTraceResponse)(nil), "ibc.transfer.QueryDenomTraceResponse")
	proto.RegisterType((*QueryDenomTracesRequest)(nil), "ibc.transfer.QueryDenomTracesRequest")
	proto.RegisterType((*QueryDenomTracesResponse)(nil), "ibc.transfer.QueryDenomTracesResponse")
}

func init() { proto.RegisterFile("ibc/transfer/query.proto", fileDescriptor_26b3e8b4e9dff1c1) }

var fileDescriptor_26b3e8b4e9dff1c1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0xfe, 0x24, 0x6e, 0x3a, 0x79, 0x80, 0x10, 0x44, 0x54, 0x45, 0x80, 0x3a, 0x50,
	0x47, 0x94, 0xa9, 0x8c, 0x08, 0x06, 0xc4, 0x02, 0x55, 0x27, 0x90, 0x40, 0xf9, 0x71, 0x5b, 0x0b,
	0x35, 0x4e, 0x63, 0x57, 0xa2, 0x6f, 0xc1, 0x0b, 0xf0, 0x36, 0x0c, 0x8c, 0x1d, 0x19, 0x51, 0xfb,
	0x22, 0xa8, 0x71, 0x43, 0x1c, 0x35, 0xa2, 0x53, 0xae, 0xec, 0x73, 0xce, 0xfd, 0x7c, 0x73, 0xc1,
	0x62, 0x41, 0xe8, 0xc9, 0xd4, 0x8f, 0x45, 0x8f, 0xa6, 0xde, 0x68, 0x4c, 0xd3, 0x09, 0x49, 0x52,
	0x2e, 0x39, 0xae, 0xb1, 0x20, 0x24, 0xf9, 0x8d, 0x7d, 0x14, 0x72, 0x31, 0xe4, 0x42, 0x29, 0xbc,
	0xc4, 0xef, 0xb3, 0xd8, 0x97, 0x8c, 0xc7, 0x4a, 0x6c, 0x1f, 0x96, 0x62, 0xf2, 0x42, 0x5d, 0xba,
	0x67, 0xb0, 0xf7, 0xb0, 0xb0, 0x5d, 0xd3, 0x98, 0x0f, 0xbb, 0xa9, 0x1f, 0xd2, 0x0e, 0x1d, 0x8d,
	0xa9, 0x90, 0x18, 0xc3, 0xd6, 0xc0, 0x17, 0x03, 0x0b, 0xd5, 0x51, 0x63, 0xb7, 0x93, 0xd5, 0x6e,
	0x17, 0xf6, 0x57, 0xd4, 0x22, 0xe1, 0xb1, 0xa0, 0xb8, 0x0d, 0x66, 0xb4, 0x38, 0x7d, 0x91, 0x8b,
	0xe3, 0xcc, 0x65, 0xb6, 0x2c, 0xa2, 0x83, 0x12, 0xcd, 0x06, 0xd1, 0x5f, 0x5d, 0x91, 0x2a, 0x72,
	0x88, 0x36, 0x40, 0xf1, 0x9e, 0x65, 0xe8, 0x01, 0x51, 0xef, 0x25, 0x6a, 0x22, 0xf7, 0x7e, 0x3f,
	0x67, 0xee, 0x68, 0x62, 0xf7, 0x03, 0x81, 0xb5, 0x1a, 0xbb, 0xa4, 0xbd, 0x81, 0x9a, 0x46, 0x2b,
	0x2c, 0x54, 0xdf, 0x6c, 0x98, 0x2d, 0xb7, 0x8c, 0x7b, 0x1b, 0xd1, 0x58, 0xb2, 0x1e, 0xa3, 0x91,
	0x06, 0x6e, 0x16, 0xe0, 0x02, 0x5f, 0x96, 0xf0, 0x36, 0x32, 0x3c, 0xbb, 0x0a, 0x4f, 0xb5, 0xd5,
	0xf9, 0x5a, 0x9f, 0x08, 0xb6, 0x33, 0x3e, 0xfc, 0x04, 0x50, 0x34, 0xc0, 0xc7, 0x65, 0x88, 0xea,
	0xbf, 0x63, 0x9f, 0xac, 0x51, 0xa9, 0x86, 0xae, 0x81, 0x9f, 0xc1, 0xd4, 0x06, 0x80, 0xff, 0xf7,
	0xe5, 0x73, 0xb7, 0x4f, 0xd7, 0xc9, 0xf2, 0xfc, 0xab, 0xbb, 0xaf, 0x99, 0x83, 0xa6, 0x33, 0x07,
	0xfd, 0xcc, 0x1c, 0xf4, 0x3e, 0x77, 0x8c, 0xe9, 0xdc, 0x31, 0xbe, 0xe7, 0x8e, 0xf1, 0x78, 0xde,
	0x67, 0x72, 0x30, 0x0e, 0x48, 0xc8, 0x87, 0xde, 0x72, 0x43, 0xd5, 0xa7, 0x29, 0xa2, 0x57, 0xef,
	0xcd, 0x63, 0x41, 0xd8, 0x2c, 0xd6, 0x72, 0x92, 0x50, 0x11, 0xec, 0x64, 0x4b, 0x79, 0xf1, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0xb2, 0x05, 0x8d, 0xfa, 0x02, 0x00, 0x00,
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
	// DenomTrace queries an IBC connection end.
	DenomTrace(ctx context.Context, in *QueryDenomTraceRequest, opts ...grpc.CallOption) (*QueryDenomTraceResponse, error)
	// DenomTraces queries all the IBC connections of a chain.
	DenomTraces(ctx context.Context, in *QueryDenomTracesRequest, opts ...grpc.CallOption) (*QueryDenomTracesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DenomTrace(ctx context.Context, in *QueryDenomTraceRequest, opts ...grpc.CallOption) (*QueryDenomTraceResponse, error) {
	out := new(QueryDenomTraceResponse)
	err := c.cc.Invoke(ctx, "/ibc.transfer.Query/DenomTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomTraces(ctx context.Context, in *QueryDenomTracesRequest, opts ...grpc.CallOption) (*QueryDenomTracesResponse, error) {
	out := new(QueryDenomTracesResponse)
	err := c.cc.Invoke(ctx, "/ibc.transfer.Query/DenomTraces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DenomTrace queries an IBC connection end.
	DenomTrace(context.Context, *QueryDenomTraceRequest) (*QueryDenomTraceResponse, error)
	// DenomTraces queries all the IBC connections of a chain.
	DenomTraces(context.Context, *QueryDenomTracesRequest) (*QueryDenomTracesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DenomTrace(ctx context.Context, req *QueryDenomTraceRequest) (*QueryDenomTraceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTrace not implemented")
}
func (*UnimplementedQueryServer) DenomTraces(ctx context.Context, req *QueryDenomTracesRequest) (*QueryDenomTracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomTraces not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DenomTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomTraceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.transfer.Query/DenomTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTrace(ctx, req.(*QueryDenomTraceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomTraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomTracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomTraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.transfer.Query/DenomTraces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomTraces(ctx, req.(*QueryDenomTracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.transfer.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DenomTrace",
			Handler:    _Query_DenomTrace_Handler,
		},
		{
			MethodName: "DenomTraces",
			Handler:    _Query_DenomTraces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/transfer/query.proto",
}

func (m *QueryDenomTraceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTraceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTraceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomTraceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTraceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTraceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomTrace != nil {
		{
			size, err := m.DenomTrace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomTracesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTracesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTracesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomTracesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomTracesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomTracesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DenomTraces) > 0 {
		for iNdEx := len(m.DenomTraces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomTraces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryDenomTraceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomTraceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DenomTrace != nil {
		l = m.DenomTrace.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomTracesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomTracesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomTraces) > 0 {
		for _, e := range m.DenomTraces {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDenomTraceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomTraceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTraceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDenomTraceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomTraceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTraceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTrace", wireType)
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
			if m.DenomTrace == nil {
				m.DenomTrace = &DenomTrace{}
			}
			if err := m.DenomTrace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDenomTracesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomTracesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTracesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDenomTracesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDenomTracesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomTracesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTraces", wireType)
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
			m.DenomTraces = append(m.DenomTraces, &IdentifiedDenomTrace{})
			if err := m.DenomTraces[len(m.DenomTraces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

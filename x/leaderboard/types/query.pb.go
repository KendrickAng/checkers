// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: leaderboard/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3340ab8601aaf86, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3340ab8601aaf86, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetPlayerInfoRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetPlayerInfoRequest) Reset()         { *m = QueryGetPlayerInfoRequest{} }
func (m *QueryGetPlayerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetPlayerInfoRequest) ProtoMessage()    {}
func (*QueryGetPlayerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3340ab8601aaf86, []int{2}
}
func (m *QueryGetPlayerInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPlayerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPlayerInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPlayerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPlayerInfoRequest.Merge(m, src)
}
func (m *QueryGetPlayerInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPlayerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPlayerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPlayerInfoRequest proto.InternalMessageInfo

func (m *QueryGetPlayerInfoRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetPlayerInfoResponse struct {
	PlayerInfo PlayerInfo `protobuf:"bytes,1,opt,name=playerInfo,proto3" json:"playerInfo"`
}

func (m *QueryGetPlayerInfoResponse) Reset()         { *m = QueryGetPlayerInfoResponse{} }
func (m *QueryGetPlayerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetPlayerInfoResponse) ProtoMessage()    {}
func (*QueryGetPlayerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3340ab8601aaf86, []int{3}
}
func (m *QueryGetPlayerInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetPlayerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetPlayerInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetPlayerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetPlayerInfoResponse.Merge(m, src)
}
func (m *QueryGetPlayerInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetPlayerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetPlayerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetPlayerInfoResponse proto.InternalMessageInfo

func (m *QueryGetPlayerInfoResponse) GetPlayerInfo() PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return PlayerInfo{}
}

type QueryAllPlayerInfoRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPlayerInfoRequest) Reset()         { *m = QueryAllPlayerInfoRequest{} }
func (m *QueryAllPlayerInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPlayerInfoRequest) ProtoMessage()    {}
func (*QueryAllPlayerInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3340ab8601aaf86, []int{4}
}
func (m *QueryAllPlayerInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPlayerInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPlayerInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPlayerInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPlayerInfoRequest.Merge(m, src)
}
func (m *QueryAllPlayerInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPlayerInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPlayerInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPlayerInfoRequest proto.InternalMessageInfo

func (m *QueryAllPlayerInfoRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllPlayerInfoResponse struct {
	PlayerInfo []PlayerInfo        `protobuf:"bytes,1,rep,name=playerInfo,proto3" json:"playerInfo"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllPlayerInfoResponse) Reset()         { *m = QueryAllPlayerInfoResponse{} }
func (m *QueryAllPlayerInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPlayerInfoResponse) ProtoMessage()    {}
func (*QueryAllPlayerInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3340ab8601aaf86, []int{5}
}
func (m *QueryAllPlayerInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPlayerInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPlayerInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPlayerInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPlayerInfoResponse.Merge(m, src)
}
func (m *QueryAllPlayerInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPlayerInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPlayerInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPlayerInfoResponse proto.InternalMessageInfo

func (m *QueryAllPlayerInfoResponse) GetPlayerInfo() []PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

func (m *QueryAllPlayerInfoResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "alice.checkers.leaderboard.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "alice.checkers.leaderboard.QueryParamsResponse")
	proto.RegisterType((*QueryGetPlayerInfoRequest)(nil), "alice.checkers.leaderboard.QueryGetPlayerInfoRequest")
	proto.RegisterType((*QueryGetPlayerInfoResponse)(nil), "alice.checkers.leaderboard.QueryGetPlayerInfoResponse")
	proto.RegisterType((*QueryAllPlayerInfoRequest)(nil), "alice.checkers.leaderboard.QueryAllPlayerInfoRequest")
	proto.RegisterType((*QueryAllPlayerInfoResponse)(nil), "alice.checkers.leaderboard.QueryAllPlayerInfoResponse")
}

func init() { proto.RegisterFile("leaderboard/query.proto", fileDescriptor_b3340ab8601aaf86) }

var fileDescriptor_b3340ab8601aaf86 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xb5, 0x0d, 0xf8, 0xc4, 0xcb, 0x18, 0xb0, 0x2e, 0xba, 0xca, 0x20, 0x56, 0x0b,
	0xce, 0x90, 0x8a, 0x7a, 0xb5, 0x3d, 0x18, 0x04, 0x0f, 0x31, 0x17, 0xc1, 0x8b, 0xcc, 0x6e, 0x5e,
	0xb7, 0xab, 0x9b, 0x9d, 0xed, 0xce, 0x46, 0x1a, 0xc4, 0x8b, 0x27, 0x8f, 0x82, 0xf8, 0x39, 0x3c,
	0xe8, 0x87, 0xe8, 0xb1, 0xe0, 0xc5, 0x93, 0x48, 0xe2, 0x07, 0x91, 0xcc, 0x4c, 0x93, 0x4d, 0x93,
	0xa6, 0x2b, 0xbd, 0x25, 0x33, 0xef, 0xff, 0x7f, 0xbf, 0x37, 0xef, 0x9f, 0xc0, 0xd5, 0x04, 0x65,
	0x17, 0xf3, 0x40, 0xc9, 0xbc, 0x2b, 0xf6, 0xfb, 0x98, 0x0f, 0x78, 0x96, 0xab, 0x42, 0x51, 0x4f,
	0x26, 0x71, 0x88, 0x3c, 0xdc, 0xc3, 0xf0, 0x2d, 0xe6, 0x9a, 0x97, 0xea, 0xbc, 0x46, 0xa4, 0x22,
	0x65, 0xca, 0xc4, 0xf8, 0x93, 0x55, 0x78, 0xd7, 0x23, 0xa5, 0xa2, 0x04, 0x85, 0xcc, 0x62, 0x21,
	0xd3, 0x54, 0x15, 0xb2, 0x88, 0x55, 0xaa, 0xdd, 0xed, 0x66, 0xa8, 0x74, 0x4f, 0x69, 0x11, 0x48,
	0x8d, 0xb6, 0x91, 0x78, 0xd7, 0x0c, 0xb0, 0x90, 0x4d, 0x91, 0xc9, 0x28, 0x4e, 0x4d, 0xb1, 0xab,
	0x5d, 0x2f, 0x43, 0x65, 0x32, 0x97, 0xbd, 0x63, 0x97, 0x1b, 0x33, 0x37, 0x89, 0x1c, 0x60, 0xfe,
	0x3a, 0x4e, 0x77, 0x1d, 0x02, 0x6b, 0x00, 0x7d, 0x31, 0xb6, 0x6e, 0x1b, 0x4d, 0x07, 0xf7, 0xfb,
	0xa8, 0x0b, 0xf6, 0x12, 0xae, 0xcc, 0x9c, 0xea, 0x4c, 0xa5, 0x1a, 0xe9, 0x13, 0xa8, 0x5b, 0xef,
	0x75, 0x72, 0x8b, 0xdc, 0xbd, 0xb4, 0xc5, 0xf8, 0xe9, 0x23, 0x73, 0xab, 0xdd, 0x59, 0x3d, 0xfc,
	0x7d, 0xb3, 0xd6, 0x71, 0x3a, 0xd6, 0x84, 0x6b, 0xc6, 0xb8, 0x85, 0x45, 0xdb, 0xb0, 0x3c, 0x4b,
	0x77, 0x95, 0xeb, 0x4a, 0x1b, 0xb0, 0x16, 0xa7, 0x5d, 0x3c, 0x30, 0xee, 0x17, 0x3b, 0xf6, 0x0b,
	0x7b, 0x03, 0xde, 0x22, 0x89, 0x43, 0x7a, 0x0e, 0x90, 0x4d, 0x4e, 0x1d, 0xd6, 0x9d, 0xa5, 0x58,
	0x93, 0x6a, 0x87, 0x56, 0xd2, 0xb3, 0xd0, 0xe1, 0x6d, 0x27, 0xc9, 0x3c, 0xde, 0x53, 0x80, 0xe9,
	0xbb, 0x4f, 0x5a, 0xd9, 0x25, 0xf1, 0xf1, 0x92, 0xb8, 0x4d, 0x83, 0x5b, 0x12, 0x6f, 0xcb, 0x08,
	0x9d, 0xb6, 0x53, 0x52, 0xb2, 0xef, 0xc4, 0x4d, 0x74, 0xa2, 0xcb, 0x29, 0x13, 0x5d, 0x38, 0xcf,
	0x44, 0xb4, 0x35, 0x03, 0xbd, 0x62, 0xa0, 0x37, 0xce, 0x84, 0xb6, 0x28, 0x65, 0xea, 0xad, 0x4f,
	0xab, 0xb0, 0x66, 0xa8, 0xe9, 0x57, 0x02, 0x75, 0xbb, 0x5c, 0xca, 0x97, 0x71, 0xcd, 0xe7, 0xca,
	0x13, 0x95, 0xeb, 0x2d, 0x01, 0xdb, 0xfc, 0xf8, 0xf3, 0xef, 0x97, 0x95, 0xdb, 0x94, 0x09, 0x23,
	0x14, 0xc7, 0x42, 0x31, 0x9f, 0x77, 0xfa, 0x83, 0x00, 0x4c, 0xdf, 0x82, 0x3e, 0x3c, 0xb3, 0xd7,
	0xa2, 0x10, 0x7a, 0x8f, 0xfe, 0x57, 0xe6, 0x48, 0x1f, 0x1b, 0xd2, 0x26, 0x15, 0x4b, 0x49, 0xa7,
	0xbf, 0x3f, 0xf1, 0xde, 0xc4, 0xfb, 0x03, 0xfd, 0x46, 0xe0, 0xf2, 0xd4, 0x6f, 0x3b, 0x49, 0x2a,
	0x90, 0x2f, 0xca, 0x67, 0x05, 0xf2, 0x85, 0x81, 0x63, 0xc2, 0x90, 0xdf, 0xa3, 0x1b, 0x15, 0xc9,
	0x77, 0x5a, 0x87, 0x43, 0x9f, 0x1c, 0x0d, 0x7d, 0xf2, 0x67, 0xe8, 0x93, 0xcf, 0x23, 0xbf, 0x76,
	0x34, 0xf2, 0x6b, 0xbf, 0x46, 0x7e, 0xed, 0xd5, 0xfd, 0x28, 0x2e, 0xf6, 0xfa, 0x01, 0x0f, 0x55,
	0xef, 0xa4, 0xd9, 0xc1, 0x8c, 0x5d, 0x31, 0xc8, 0x50, 0x07, 0x75, 0xf3, 0x1f, 0xf4, 0xe0, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x4f, 0x03, 0x6e, 0x53, 0x05, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a PlayerInfo by index.
	PlayerInfo(ctx context.Context, in *QueryGetPlayerInfoRequest, opts ...grpc.CallOption) (*QueryGetPlayerInfoResponse, error)
	// Queries a list of PlayerInfo items.
	PlayerInfoAll(ctx context.Context, in *QueryAllPlayerInfoRequest, opts ...grpc.CallOption) (*QueryAllPlayerInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/alice.checkers.leaderboard.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PlayerInfo(ctx context.Context, in *QueryGetPlayerInfoRequest, opts ...grpc.CallOption) (*QueryGetPlayerInfoResponse, error) {
	out := new(QueryGetPlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/alice.checkers.leaderboard.Query/PlayerInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PlayerInfoAll(ctx context.Context, in *QueryAllPlayerInfoRequest, opts ...grpc.CallOption) (*QueryAllPlayerInfoResponse, error) {
	out := new(QueryAllPlayerInfoResponse)
	err := c.cc.Invoke(ctx, "/alice.checkers.leaderboard.Query/PlayerInfoAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a PlayerInfo by index.
	PlayerInfo(context.Context, *QueryGetPlayerInfoRequest) (*QueryGetPlayerInfoResponse, error)
	// Queries a list of PlayerInfo items.
	PlayerInfoAll(context.Context, *QueryAllPlayerInfoRequest) (*QueryAllPlayerInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) PlayerInfo(ctx context.Context, req *QueryGetPlayerInfoRequest) (*QueryGetPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerInfo not implemented")
}
func (*UnimplementedQueryServer) PlayerInfoAll(ctx context.Context, req *QueryAllPlayerInfoRequest) (*QueryAllPlayerInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerInfoAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.checkers.leaderboard.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.checkers.leaderboard.Query/PlayerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PlayerInfo(ctx, req.(*QueryGetPlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PlayerInfoAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPlayerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PlayerInfoAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alice.checkers.leaderboard.Query/PlayerInfoAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PlayerInfoAll(ctx, req.(*QueryAllPlayerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alice.checkers.leaderboard.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "PlayerInfo",
			Handler:    _Query_PlayerInfo_Handler,
		},
		{
			MethodName: "PlayerInfoAll",
			Handler:    _Query_PlayerInfoAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "leaderboard/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetPlayerInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPlayerInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPlayerInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetPlayerInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetPlayerInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetPlayerInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.PlayerInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllPlayerInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPlayerInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPlayerInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllPlayerInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPlayerInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPlayerInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.PlayerInfo) > 0 {
		for iNdEx := len(m.PlayerInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PlayerInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetPlayerInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetPlayerInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PlayerInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllPlayerInfoRequest) Size() (n int) {
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

func (m *QueryAllPlayerInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PlayerInfo) > 0 {
		for _, e := range m.PlayerInfo {
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryGetPlayerInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetPlayerInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPlayerInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryGetPlayerInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetPlayerInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetPlayerInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerInfo", wireType)
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
			if err := m.PlayerInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryAllPlayerInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPlayerInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPlayerInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryAllPlayerInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPlayerInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPlayerInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerInfo", wireType)
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
			m.PlayerInfo = append(m.PlayerInfo, PlayerInfo{})
			if err := m.PlayerInfo[len(m.PlayerInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
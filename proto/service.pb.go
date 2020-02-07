// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Key struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Value struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PrefixKey struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrefixKey) Reset()         { *m = PrefixKey{} }
func (m *PrefixKey) String() string { return proto.CompactTextString(m) }
func (*PrefixKey) ProtoMessage()    {}
func (*PrefixKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *PrefixKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrefixKey.Unmarshal(m, b)
}
func (m *PrefixKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrefixKey.Marshal(b, m, deterministic)
}
func (m *PrefixKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrefixKey.Merge(m, src)
}
func (m *PrefixKey) XXX_Size() int {
	return xxx_messageInfo_PrefixKey.Size(m)
}
func (m *PrefixKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PrefixKey.DiscardUnknown(m)
}

var xxx_messageInfo_PrefixKey proto.InternalMessageInfo

func (m *PrefixKey) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*Key)(nil), "proto.Key")
	proto.RegisterType((*Value)(nil), "proto.Value")
	proto.RegisterType((*KeyValuePair)(nil), "proto.KeyValuePair")
	proto.RegisterType((*SetResponse)(nil), "proto.SetResponse")
	proto.RegisterType((*PrefixKey)(nil), "proto.PrefixKey")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x17, 0x43, 0x37, 0x7b, 0x9d, 0x30, 0xae, 0xa2, 0x63, 0x20, 0x68, 0x7c, 0xd0, 0xa7,
	0x39, 0x14, 0xfc, 0x0b, 0x7b, 0xc8, 0xcb, 0x68, 0x61, 0xef, 0xb3, 0x1c, 0xa1, 0x28, 0xa6, 0x24,
	0xd9, 0x30, 0xbf, 0xc2, 0xbf, 0x2c, 0x4d, 0xe3, 0x56, 0xdc, 0x53, 0xcf, 0xb9, 0x3d, 0x27, 0xf7,
	0xbb, 0x74, 0xee, 0x60, 0x77, 0x75, 0x85, 0x79, 0x63, 0x8d, 0x37, 0x9c, 0xc5, 0x8f, 0xba, 0x26,
	0xa9, 0x11, 0x78, 0x42, 0xf2, 0x03, 0x61, 0x2a, 0x6e, 0xc5, 0x63, 0x5e, 0xb4, 0x52, 0xdd, 0x50,
	0xb6, 0xde, 0x7c, 0x6e, 0xc1, 0x97, 0x94, 0xed, 0x5a, 0x91, 0x7e, 0x76, 0x46, 0xbd, 0xd2, 0x58,
	0x23, 0xc4, 0xc4, 0x6a, 0x53, 0xdb, 0xe3, 0x07, 0x0e, 0xbd, 0x93, 0x7e, 0xef, 0x81, 0xce, 0x4a,
	0xf8, 0x02, 0xae, 0x31, 0x5f, 0x0e, 0x3c, 0xa5, 0x91, 0xdb, 0x56, 0x15, 0x9c, 0x8b, 0xd5, 0xd3,
	0xe2, 0xcf, 0xaa, 0x7b, 0xca, 0x57, 0x16, 0xef, 0xf5, 0x77, 0x8b, 0x77, 0x45, 0xc3, 0x26, 0x9a,
	0xb4, 0x20, 0xb9, 0xe7, 0x1f, 0x41, 0x23, 0xbd, 0x2e, 0xbd, 0xb1, 0xe0, 0x3b, 0x92, 0x4b, 0x78,
	0xa6, 0xee, 0xbe, 0xb9, 0x46, 0x98, 0x8d, 0x93, 0x8e, 0x98, 0x6a, 0xc0, 0x0b, 0x92, 0x25, 0x3c,
	0x5f, 0x1c, 0x22, 0xfb, 0x03, 0x66, 0x9c, 0x86, 0x3d, 0x3a, 0x35, 0xe0, 0x27, 0xca, 0x97, 0xf0,
	0x1d, 0x08, 0x4f, 0x52, 0x64, 0xcf, 0xf5, 0x7f, 0xc1, 0x42, 0xbc, 0x0d, 0xe3, 0xe0, 0xe5, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x25, 0x01, 0x3f, 0x83, 0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVStoreClient is the client API for KVStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KVStoreClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Set(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*SetResponse, error)
	GetPrefix(ctx context.Context, in *PrefixKey, opts ...grpc.CallOption) (KVStore_GetPrefixClient, error)
}

type kVStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKVStoreClient(cc grpc.ClientConnInterface) KVStoreClient {
	return &kVStoreClient{cc}
}

func (c *kVStoreClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/proto.KVStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) Set(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/proto.KVStore/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) GetPrefix(ctx context.Context, in *PrefixKey, opts ...grpc.CallOption) (KVStore_GetPrefixClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KVStore_serviceDesc.Streams[0], "/proto.KVStore/GetPrefix", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVStoreGetPrefixClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KVStore_GetPrefixClient interface {
	Recv() (*Value, error)
	grpc.ClientStream
}

type kVStoreGetPrefixClient struct {
	grpc.ClientStream
}

func (x *kVStoreGetPrefixClient) Recv() (*Value, error) {
	m := new(Value)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVStoreServer is the server API for KVStore service.
type KVStoreServer interface {
	Get(context.Context, *Key) (*Value, error)
	Set(context.Context, *KeyValuePair) (*SetResponse, error)
	GetPrefix(*PrefixKey, KVStore_GetPrefixServer) error
}

// UnimplementedKVStoreServer can be embedded to have forward compatible implementations.
type UnimplementedKVStoreServer struct {
}

func (*UnimplementedKVStoreServer) Get(ctx context.Context, req *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedKVStoreServer) Set(ctx context.Context, req *KeyValuePair) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedKVStoreServer) GetPrefix(req *PrefixKey, srv KVStore_GetPrefixServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPrefix not implemented")
}

func RegisterKVStoreServer(s *grpc.Server, srv KVStoreServer) {
	s.RegisterService(&_KVStore_serviceDesc, srv)
}

func _KVStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KVStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.KVStore/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Set(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_GetPrefix_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrefixKey)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVStoreServer).GetPrefix(m, &kVStoreGetPrefixServer{stream})
}

type KVStore_GetPrefixServer interface {
	Send(*Value) error
	grpc.ServerStream
}

type kVStoreGetPrefixServer struct {
	grpc.ServerStream
}

func (x *kVStoreGetPrefixServer) Send(m *Value) error {
	return x.ServerStream.SendMsg(m)
}

var _KVStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.KVStore",
	HandlerType: (*KVStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KVStore_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _KVStore_Set_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPrefix",
			Handler:       _KVStore_GetPrefix_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}

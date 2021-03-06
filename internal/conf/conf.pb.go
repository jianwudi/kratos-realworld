// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/conf/conf.proto

package conf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Bootstrap struct {
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data                 *Data    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Jwt                  *JWT     `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bootstrap) Reset()         { *m = Bootstrap{} }
func (m *Bootstrap) String() string { return proto.CompactTextString(m) }
func (*Bootstrap) ProtoMessage()    {}
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{0}
}

func (m *Bootstrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap.Unmarshal(m, b)
}
func (m *Bootstrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap.Marshal(b, m, deterministic)
}
func (m *Bootstrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap.Merge(m, src)
}
func (m *Bootstrap) XXX_Size() int {
	return xxx_messageInfo_Bootstrap.Size(m)
}
func (m *Bootstrap) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap proto.InternalMessageInfo

func (m *Bootstrap) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *Bootstrap) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Bootstrap) GetJwt() *JWT {
	if m != nil {
		return m.Jwt
	}
	return nil
}

type Server struct {
	Http                 *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc                 *Server_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{1}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetHttp() *Server_HTTP {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Server) GetGrpc() *Server_GRPC {
	if m != nil {
		return m.Grpc
	}
	return nil
}

type Server_HTTP struct {
	Network              string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr                 string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout              *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Server_HTTP) Reset()         { *m = Server_HTTP{} }
func (m *Server_HTTP) String() string { return proto.CompactTextString(m) }
func (*Server_HTTP) ProtoMessage()    {}
func (*Server_HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{1, 0}
}

func (m *Server_HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server_HTTP.Unmarshal(m, b)
}
func (m *Server_HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server_HTTP.Marshal(b, m, deterministic)
}
func (m *Server_HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server_HTTP.Merge(m, src)
}
func (m *Server_HTTP) XXX_Size() int {
	return xxx_messageInfo_Server_HTTP.Size(m)
}
func (m *Server_HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_Server_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_Server_HTTP proto.InternalMessageInfo

func (m *Server_HTTP) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Server_HTTP) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Server_HTTP) GetTimeout() *durationpb.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type Server_GRPC struct {
	Network              string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr                 string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout              *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Server_GRPC) Reset()         { *m = Server_GRPC{} }
func (m *Server_GRPC) String() string { return proto.CompactTextString(m) }
func (*Server_GRPC) ProtoMessage()    {}
func (*Server_GRPC) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{1, 1}
}

func (m *Server_GRPC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server_GRPC.Unmarshal(m, b)
}
func (m *Server_GRPC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server_GRPC.Marshal(b, m, deterministic)
}
func (m *Server_GRPC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server_GRPC.Merge(m, src)
}
func (m *Server_GRPC) XXX_Size() int {
	return xxx_messageInfo_Server_GRPC.Size(m)
}
func (m *Server_GRPC) XXX_DiscardUnknown() {
	xxx_messageInfo_Server_GRPC.DiscardUnknown(m)
}

var xxx_messageInfo_Server_GRPC proto.InternalMessageInfo

func (m *Server_GRPC) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Server_GRPC) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Server_GRPC) GetTimeout() *durationpb.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type Data struct {
	Database             *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{2}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetDatabase() *Data_Database {
	if m != nil {
		return m.Database
	}
	return nil
}

type Data_Database struct {
	Dsn                  string   `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data_Database) Reset()         { *m = Data_Database{} }
func (m *Data_Database) String() string { return proto.CompactTextString(m) }
func (*Data_Database) ProtoMessage()    {}
func (*Data_Database) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{2, 0}
}

func (m *Data_Database) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data_Database.Unmarshal(m, b)
}
func (m *Data_Database) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data_Database.Marshal(b, m, deterministic)
}
func (m *Data_Database) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data_Database.Merge(m, src)
}
func (m *Data_Database) XXX_Size() int {
	return xxx_messageInfo_Data_Database.Size(m)
}
func (m *Data_Database) XXX_DiscardUnknown() {
	xxx_messageInfo_Data_Database.DiscardUnknown(m)
}

var xxx_messageInfo_Data_Database proto.InternalMessageInfo

func (m *Data_Database) GetDsn() string {
	if m != nil {
		return m.Dsn
	}
	return ""
}

type JWT struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWT) Reset()         { *m = JWT{} }
func (m *JWT) String() string { return proto.CompactTextString(m) }
func (*JWT) ProtoMessage()    {}
func (*JWT) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c69a7f648509b54, []int{3}
}

func (m *JWT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JWT.Unmarshal(m, b)
}
func (m *JWT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JWT.Marshal(b, m, deterministic)
}
func (m *JWT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWT.Merge(m, src)
}
func (m *JWT) XXX_Size() int {
	return xxx_messageInfo_JWT.Size(m)
}
func (m *JWT) XXX_DiscardUnknown() {
	xxx_messageInfo_JWT.DiscardUnknown(m)
}

var xxx_messageInfo_JWT proto.InternalMessageInfo

func (m *JWT) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func init() {
	proto.RegisterType((*Bootstrap)(nil), "kratos.api.Bootstrap")
	proto.RegisterType((*Server)(nil), "kratos.api.Server")
	proto.RegisterType((*Server_HTTP)(nil), "kratos.api.Server.HTTP")
	proto.RegisterType((*Server_GRPC)(nil), "kratos.api.Server.GRPC")
	proto.RegisterType((*Data)(nil), "kratos.api.Data")
	proto.RegisterType((*Data_Database)(nil), "kratos.api.Data.Database")
	proto.RegisterType((*JWT)(nil), "kratos.api.JWT")
}

func init() { proto.RegisterFile("internal/conf/conf.proto", fileDescriptor_9c69a7f648509b54) }

var fileDescriptor_9c69a7f648509b54 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x4d, 0x4b, 0xfb, 0x40,
	0x10, 0xc6, 0x69, 0x13, 0xfa, 0x32, 0xff, 0xc3, 0xbf, 0xec, 0x41, 0xd3, 0xa2, 0xa2, 0x51, 0x41,
	0x14, 0x37, 0x60, 0xf1, 0xe4, 0xad, 0x16, 0x94, 0x9e, 0xca, 0x1a, 0x28, 0xe8, 0x69, 0xdb, 0x6c,
	0x6b, 0x6c, 0xdc, 0x0d, 0x9b, 0xa9, 0xbd, 0xfa, 0xb1, 0xfc, 0x78, 0xb2, 0x9b, 0x8d, 0xd6, 0x97,
	0xb3, 0x97, 0x61, 0x37, 0xcf, 0x2f, 0xf3, 0x4c, 0x9e, 0x0c, 0x04, 0xa9, 0x44, 0xa1, 0x25, 0xcf,
	0xa2, 0x99, 0x92, 0x73, 0x5b, 0x68, 0xae, 0x15, 0x2a, 0x02, 0x4b, 0xcd, 0x51, 0x15, 0x94, 0xe7,
	0x69, 0x6f, 0x6f, 0xa1, 0xd4, 0x22, 0x13, 0x91, 0x55, 0xa6, 0xab, 0x79, 0x94, 0xac, 0x34, 0xc7,
	0x54, 0xc9, 0x92, 0x0d, 0x5f, 0x6b, 0xd0, 0x1e, 0x28, 0x85, 0x05, 0x6a, 0x9e, 0x93, 0x53, 0x68,
	0x14, 0x42, 0xbf, 0x08, 0x1d, 0xd4, 0xf6, 0x6b, 0x27, 0xff, 0x2e, 0x08, 0xfd, 0x6c, 0x45, 0xef,
	0xac, 0xc2, 0x1c, 0x41, 0x8e, 0xc0, 0x4f, 0x38, 0xf2, 0xa0, 0x6e, 0xc9, 0xce, 0x26, 0x39, 0xe4,
	0xc8, 0x99, 0x55, 0xc9, 0x01, 0x78, 0x4f, 0x6b, 0x0c, 0x3c, 0x0b, 0xfd, 0xdf, 0x84, 0x46, 0x93,
	0x98, 0x19, 0x2d, 0x7c, 0xab, 0x43, 0xa3, 0xec, 0x4d, 0xce, 0xc0, 0x7f, 0x44, 0xcc, 0x9d, 0xfb,
	0xf6, 0x4f, 0x77, 0x7a, 0x1b, 0xc7, 0x63, 0x66, 0x21, 0x03, 0x2f, 0x74, 0x3e, 0x73, 0x03, 0xfc,
	0x06, 0xdf, 0xb0, 0xf1, 0x35, 0xb3, 0x50, 0x2f, 0x05, 0xdf, 0xbc, 0x4a, 0x02, 0x68, 0x4a, 0x81,
	0x6b, 0xa5, 0x97, 0xd6, 0xa4, 0xcd, 0xaa, 0x2b, 0x21, 0xe0, 0xf3, 0x24, 0xd1, 0xb6, 0x5d, 0x9b,
	0xd9, 0x33, 0xe9, 0x43, 0x13, 0xd3, 0x67, 0xa1, 0x56, 0xd5, 0x17, 0x74, 0x69, 0x99, 0x27, 0xad,
	0xf2, 0xa4, 0x43, 0x97, 0x27, 0xab, 0x48, 0x63, 0x65, 0x8c, 0xff, 0xc0, 0x2a, 0x7c, 0x00, 0xdf,
	0x64, 0x4d, 0x2e, 0xa1, 0x65, 0xd2, 0x9e, 0xf2, 0x42, 0xb8, 0xec, 0xba, 0xdf, 0xff, 0x87, 0x2d,
	0x06, 0x60, 0x1f, 0x68, 0x6f, 0x07, 0x5a, 0xd5, 0x53, 0xd2, 0x01, 0x2f, 0x29, 0xa4, 0x9b, 0xd4,
	0x1c, 0xc3, 0x5d, 0xf0, 0x46, 0x93, 0x98, 0x6c, 0x99, 0x9d, 0x98, 0x69, 0x81, 0x4e, 0x73, 0xb7,
	0xc1, 0xf1, 0xfd, 0x61, 0x69, 0x71, 0xae, 0x05, 0xcf, 0xd6, 0x4a, 0x67, 0x49, 0xf4, 0x65, 0x25,
	0xaf, 0x4c, 0x99, 0x36, 0xec, 0xf8, 0xfd, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xdc, 0x61,
	0x76, 0xaf, 0x02, 0x00, 0x00,
}

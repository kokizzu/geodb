// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type StreamRequest struct {
	Regex                string   `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

type StreamResponse struct {
	Data                 *Data    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpsertRequest struct {
	Data                 map[string]*Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UpsertRequest) Reset()         { *m = UpsertRequest{} }
func (m *UpsertRequest) String() string { return proto.CompactTextString(m) }
func (*UpsertRequest) ProtoMessage()    {}
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *UpsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertRequest.Unmarshal(m, b)
}
func (m *UpsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertRequest.Marshal(b, m, deterministic)
}
func (m *UpsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertRequest.Merge(m, src)
}
func (m *UpsertRequest) XXX_Size() int {
	return xxx_messageInfo_UpsertRequest.Size(m)
}
func (m *UpsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertRequest proto.InternalMessageInfo

func (m *UpsertRequest) GetData() map[string]*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpsertResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpsertResponse) Reset()         { *m = UpsertResponse{} }
func (m *UpsertResponse) String() string { return proto.CompactTextString(m) }
func (*UpsertResponse) ProtoMessage()    {}
func (*UpsertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *UpsertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpsertResponse.Unmarshal(m, b)
}
func (m *UpsertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpsertResponse.Marshal(b, m, deterministic)
}
func (m *UpsertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpsertResponse.Merge(m, src)
}
func (m *UpsertResponse) XXX_Size() int {
	return xxx_messageInfo_UpsertResponse.Size(m)
}
func (m *UpsertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpsertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpsertResponse proto.InternalMessageInfo

type GetRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type GetResponse struct {
	Data                 map[string]*Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetData() map[string]*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteRequest struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type Point struct {
	Lat                  float64  `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon                  float64  `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Point) GetLon() float64 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Data struct {
	Key                  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Point                *Point            `protobuf:"bytes,2,opt,name=point,proto3" json:"point,omitempty"`
	Radius               int64             `protobuf:"varint,3,opt,name=radius,proto3" json:"radius,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	UpdatedUnix          int64             `protobuf:"varint,5,opt,name=updated_unix,json=updatedUnix,proto3" json:"updated_unix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
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

func (m *Data) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Data) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *Data) GetRadius() int64 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *Data) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Data) GetUpdatedUnix() int64 {
	if m != nil {
		return m.UpdatedUnix
	}
	return 0
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "api.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "api.StreamResponse")
	proto.RegisterType((*UpsertRequest)(nil), "api.UpsertRequest")
	proto.RegisterMapType((map[string]*Data)(nil), "api.UpsertRequest.DataEntry")
	proto.RegisterType((*UpsertResponse)(nil), "api.UpsertResponse")
	proto.RegisterType((*GetRequest)(nil), "api.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "api.GetResponse")
	proto.RegisterMapType((map[string]*Data)(nil), "api.GetResponse.DataEntry")
	proto.RegisterType((*DeleteRequest)(nil), "api.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "api.DeleteResponse")
	proto.RegisterType((*PingRequest)(nil), "api.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "api.PingResponse")
	proto.RegisterType((*Point)(nil), "api.Point")
	proto.RegisterType((*Data)(nil), "api.Data")
	proto.RegisterMapType((map[string]string)(nil), "api.Data.MetadataEntry")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x6d, 0xbe, 0xaa, 0xcd, 0x64, 0x53, 0x82, 0x59, 0x41, 0x15, 0xf1, 0x51, 0x8c, 0x90, 0x2a,
	0x50, 0xd3, 0x55, 0x2b, 0x10, 0x82, 0x5b, 0x55, 0xd4, 0x13, 0xd2, 0x2a, 0x68, 0x2f, 0x5c, 0x90,
	0x97, 0x58, 0xc5, 0x4a, 0x1b, 0x87, 0xc4, 0xd9, 0x6d, 0x8f, 0x08, 0xf1, 0x5b, 0xf9, 0x1b, 0x28,
	0xb6, 0x9b, 0x6d, 0x4a, 0xc5, 0x6d, 0x6f, 0xf6, 0xcc, 0x7b, 0x33, 0x6f, 0x5e, 0xc6, 0x01, 0x97,
	0xe4, 0x2c, 0xca, 0x0b, 0x2e, 0x38, 0xb2, 0x48, 0xce, 0xc2, 0xb7, 0x4b, 0x26, 0xbe, 0x57, 0x57,
	0xd1, 0x37, 0xbe, 0x1e, 0xaf, 0x6f, 0x98, 0x48, 0xf9, 0xcd, 0x78, 0xc9, 0x47, 0x12, 0x31, 0xba,
	0x26, 0x2b, 0x96, 0x10, 0xc1, 0x8b, 0x72, 0xdc, 0x1c, 0x15, 0x19, 0xbf, 0x04, 0xff, 0xb3, 0x28,
	0x28, 0x59, 0xc7, 0xf4, 0x47, 0x45, 0x4b, 0x81, 0xce, 0xc0, 0x29, 0xe8, 0x92, 0x6e, 0xfa, 0xc6,
	0xc0, 0x18, 0xba, 0xb1, 0xba, 0xe0, 0x31, 0xf4, 0x76, 0xb0, 0x32, 0xe7, 0x59, 0x49, 0xd1, 0x13,
	0xb0, 0x13, 0x22, 0x88, 0x84, 0x79, 0x13, 0x37, 0xaa, 0xf5, 0xcc, 0x89, 0x20, 0xb1, 0x0c, 0xe3,
	0xdf, 0x06, 0xf8, 0x97, 0x79, 0x49, 0x0b, 0xb1, 0x2b, 0x7c, 0xde, 0x10, 0xac, 0xa1, 0x37, 0x79,
	0x2c, 0x09, 0x2d, 0x84, 0xa4, 0x7f, 0xcc, 0x44, 0xb1, 0x55, 0x35, 0xc2, 0x19, 0xb8, 0x4d, 0x08,
	0x05, 0x60, 0xa5, 0x74, 0xab, 0x55, 0xd5, 0x47, 0xf4, 0x0c, 0x9c, 0x6b, 0xb2, 0xaa, 0x68, 0xdf,
	0x3c, 0x94, 0xa0, 0xe2, 0xef, 0xcd, 0x77, 0x06, 0x0e, 0xa0, 0xb7, 0x6b, 0xa2, 0x84, 0xe3, 0x01,
	0xc0, 0x82, 0x36, 0xaa, 0x10, 0xd8, 0x29, 0xdd, 0x96, 0x52, 0x95, 0x1b, 0xcb, 0x33, 0xfe, 0x69,
	0x80, 0x27, 0x21, 0x7a, 0xd4, 0xa8, 0xa5, 0x3c, 0x94, 0x7d, 0xf6, 0xf2, 0x77, 0xa2, 0xfb, 0x05,
	0xf8, 0x73, 0xba, 0xa2, 0x82, 0xfe, 0x4f, 0x68, 0x00, 0xbd, 0x1d, 0x48, 0x0f, 0xe7, 0x83, 0x77,
	0xc1, 0xb2, 0xa5, 0x26, 0xe1, 0xa7, 0x70, 0xaa, 0xae, 0x7a, 0x92, 0x1e, 0x98, 0x3c, 0x95, 0x5a,
	0x4e, 0x62, 0x93, 0xa7, 0xf8, 0x35, 0x38, 0x17, 0x9c, 0x65, 0xa2, 0x56, 0xb9, 0x22, 0x42, 0x66,
	0x8c, 0xb8, 0x3e, 0xca, 0x08, 0xcf, 0xa4, 0xc6, 0x3a, 0xc2, 0x33, 0xfc, 0xc7, 0x00, 0xbb, 0x96,
	0x79, 0x64, 0xa4, 0x01, 0x38, 0x79, 0x5d, 0x47, 0x8f, 0x04, 0x72, 0x24, 0x59, 0x39, 0x56, 0x09,
	0xf4, 0x10, 0xba, 0x05, 0x49, 0x58, 0x55, 0xf6, 0xad, 0x81, 0x31, 0xb4, 0x62, 0x7d, 0x43, 0x53,
	0x38, 0x59, 0x53, 0x41, 0xa4, 0xbf, 0xb6, 0xf4, 0xf7, 0x51, 0xe3, 0x47, 0xf4, 0x49, 0x67, 0x94,
	0xb9, 0x0d, 0x10, 0x3d, 0x87, 0xd3, 0x2a, 0x4f, 0x88, 0xa0, 0xc9, 0xd7, 0x2a, 0x63, 0x9b, 0xbe,
	0x23, 0x4b, 0x7a, 0x3a, 0x76, 0x99, 0xb1, 0x4d, 0xf8, 0x01, 0xfc, 0x16, 0xfb, 0x88, 0xe8, 0xb3,
	0xfd, 0xef, 0xe0, 0xee, 0x99, 0x3f, 0xf9, 0x65, 0x82, 0xb3, 0xa0, 0x7c, 0x3e, 0x43, 0x23, 0xb0,
	0x6b, 0x03, 0x51, 0xa0, 0x26, 0xba, 0xb5, 0x36, 0xbc, 0xbf, 0x17, 0xd1, 0xe6, 0x77, 0xd0, 0x14,
	0xba, 0x6a, 0xdb, 0x10, 0xfa, 0x77, 0xbf, 0xc3, 0x07, 0xad, 0x58, 0x43, 0x7a, 0x05, 0xd6, 0x82,
	0x0a, 0x74, 0xef, 0x76, 0xaf, 0x14, 0x3c, 0x38, 0x5c, 0x34, 0xd5, 0x40, 0x7d, 0x71, 0xdd, 0xa0,
	0xb5, 0x23, 0xba, 0xc1, 0xc1, 0x4a, 0x74, 0xd0, 0x1b, 0xe8, 0xaa, 0xc7, 0xab, 0x49, 0xad, 0x07,
	0xaf, 0x49, 0xed, 0xd7, 0x8d, 0x3b, 0xe7, 0xc6, 0xcc, 0xf9, 0x52, 0xff, 0x59, 0xae, 0xba, 0xf2,
	0x47, 0x31, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x2d, 0xcd, 0xce, 0x72, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GeoDBClient is the client API for GeoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeoDBClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (GeoDB_StreamClient, error)
}

type geoDBClient struct {
	cc *grpc.ClientConn
}

func NewGeoDBClient(cc *grpc.ClientConn) GeoDBClient {
	return &geoDBClient{cc}
}

func (c *geoDBClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertResponse, error) {
	out := new(UpsertResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/api.GeoDB/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoDBClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (GeoDB_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GeoDB_serviceDesc.Streams[0], "/api.GeoDB/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &geoDBStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GeoDB_StreamClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type geoDBStreamClient struct {
	grpc.ClientStream
}

func (x *geoDBStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GeoDBServer is the server API for GeoDB service.
type GeoDBServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Upsert(context.Context, *UpsertRequest) (*UpsertResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Stream(*StreamRequest, GeoDB_StreamServer) error
}

// UnimplementedGeoDBServer can be embedded to have forward compatible implementations.
type UnimplementedGeoDBServer struct {
}

func (*UnimplementedGeoDBServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedGeoDBServer) Upsert(ctx context.Context, req *UpsertRequest) (*UpsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (*UnimplementedGeoDBServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedGeoDBServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedGeoDBServer) Stream(req *StreamRequest, srv GeoDB_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterGeoDBServer(s *grpc.Server, srv GeoDBServer) {
	s.RegisterService(&_GeoDB_serviceDesc, srv)
}

func _GeoDB_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GeoDB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoDBServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoDB_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GeoDBServer).Stream(m, &geoDBStreamServer{stream})
}

type GeoDB_StreamServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type geoDBStreamServer struct {
	grpc.ServerStream
}

func (x *geoDBStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GeoDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GeoDB",
	HandlerType: (*GeoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GeoDB_Ping_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _GeoDB_Upsert_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GeoDB_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GeoDB_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _GeoDB_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

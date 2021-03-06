// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package go_micro_srv_consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
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

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=Consignments,proto3" json:"Consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4e, 0xf3, 0x30,
	0x10, 0xc5, 0xbf, 0xf4, 0x7f, 0x26, 0xd5, 0x57, 0xe1, 0x05, 0x58, 0xb0, 0x20, 0x4a, 0x41, 0xea,
	0x2a, 0x48, 0xe5, 0x08, 0x59, 0x54, 0xd9, 0xba, 0x6b, 0x04, 0x25, 0x1e, 0xa5, 0x96, 0x88, 0x1d,
	0x6c, 0xb7, 0x5c, 0x0d, 0xae, 0xc1, 0x89, 0x50, 0x9d, 0x86, 0x1a, 0xa1, 0xa2, 0xee, 0xfc, 0xe6,
	0xf9, 0x4d, 0x7e, 0x79, 0x32, 0x4c, 0x6b, 0xad, 0xac, 0xba, 0x2b, 0x94, 0x34, 0xa2, 0x94, 0x15,
	0x4a, 0xeb, 0x9f, 0x53, 0xe7, 0x12, 0x5a, 0xaa, 0xb4, 0x12, 0x85, 0x56, 0xa9, 0xd1, 0xdb, 0xd4,
	0xf3, 0x93, 0x8f, 0x00, 0xa2, 0xec, 0xa0, 0xc9, 0x7f, 0xe8, 0x08, 0x4e, 0x83, 0x38, 0x98, 0x85,
	0xac, 0x23, 0x38, 0x89, 0x21, 0xe2, 0x68, 0x0a, 0x2d, 0x6a, 0x2b, 0x94, 0xa4, 0x1d, 0x67, 0xf8,
	0x23, 0x72, 0x0e, 0x83, 0x37, 0x14, 0xe5, 0xda, 0xd2, 0x6e, 0x1c, 0xcc, 0xfa, 0x6c, 0xaf, 0x48,
	0x06, 0x50, 0x28, 0x69, 0x57, 0x42, 0xa2, 0x36, 0xb4, 0x17, 0x77, 0x67, 0xd1, 0x7c, 0x9a, 0x1e,
	0x03, 0x49, 0xb3, 0xf6, 0x2e, 0xf3, 0x62, 0xe4, 0x0a, 0xc2, 0x2d, 0x1a, 0x83, 0x2f, 0x8f, 0x82,
	0xd3, 0xbe, 0xfb, 0xf8, 0xa8, 0x19, 0xe4, 0x3c, 0xa9, 0x20, 0xfc, 0x4e, 0xfd, 0x02, 0xbf, 0x86,
	0xa8, 0xd8, 0x18, 0xab, 0x2a, 0xd4, 0xbb, 0x6c, 0x03, 0x0e, 0xed, 0x28, 0xe7, 0x3b, 0x6e, 0xa5,
	0x45, 0x29, 0xa4, 0xe3, 0x0e, 0xd9, 0x5e, 0x91, 0x0b, 0x18, 0x6e, 0x4c, 0x13, 0xea, 0x35, 0xc6,
	0x4e, 0xe6, 0x3c, 0x19, 0x03, 0x2c, 0xd0, 0x32, 0x7c, 0xdd, 0xa0, 0xb1, 0xc9, 0x7b, 0x00, 0x23,
	0x86, 0xa6, 0x56, 0xd2, 0x20, 0xa1, 0x30, 0x2c, 0x34, 0xae, 0x2c, 0x36, 0x04, 0x23, 0xd6, 0x4a,
	0xb2, 0x80, 0xc8, 0xfb, 0x4b, 0x87, 0x11, 0xcd, 0x6f, 0xff, 0xac, 0xa1, 0x3d, 0x33, 0x3f, 0x49,
	0x72, 0x18, 0x7b, 0x9e, 0xa1, 0x5d, 0x57, 0xe8, 0x89, 0x9b, 0x7e, 0x44, 0xe7, 0x9f, 0x01, 0x4c,
	0x96, 0x6b, 0x51, 0xd7, 0x42, 0x96, 0x4b, 0xd4, 0x5b, 0x51, 0x20, 0x79, 0x82, 0xb3, 0xcc, 0x21,
	0xfb, 0x8f, 0xe1, 0xb4, 0xed, 0x97, 0xc9, 0xf1, 0x6b, 0x6d, 0x43, 0xc9, 0x3f, 0xf2, 0x00, 0x93,
	0x05, 0x5a, 0x1f, 0x84, 0xdc, 0x1c, 0x0f, 0x1e, 0x9a, 0x3e, 0x6d, 0xfd, 0xf3, 0xc0, 0xbd, 0xf4,
	0xfb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x27, 0x35, 0xbd, 0x10, 0x03, 0x00, 0x00,
}

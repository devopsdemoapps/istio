// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: distrib/manifest.proto

/*
Package distrib is a generated protocol buffer package.

It is generated from these files:
	distrib/manifest.proto

It has these top-level messages:
	Manifest
*/
package distrib

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Manifest struct {
	// Unique id of this manifest
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The id of component that this manifest is intended for
	ComponentId string `protobuf:"bytes,2,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
	// The type of component that this manifest is intended for
	ComponentType string `protobuf:"bytes,3,opt,name=component_type,json=componentType,proto3" json:"component_type,omitempty"`
	// The ids of config fragments.
	FragmentIds []string `protobuf:"bytes,4,rep,name=fragment_ids,json=fragmentIds" json:"fragment_ids,omitempty"`
}

func (m *Manifest) Reset()                    { *m = Manifest{} }
func (m *Manifest) String() string            { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()               {}
func (*Manifest) Descriptor() ([]byte, []int) { return fileDescriptorManifest, []int{0} }

func (m *Manifest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Manifest) GetComponentId() string {
	if m != nil {
		return m.ComponentId
	}
	return ""
}

func (m *Manifest) GetComponentType() string {
	if m != nil {
		return m.ComponentType
	}
	return ""
}

func (m *Manifest) GetFragmentIds() []string {
	if m != nil {
		return m.FragmentIds
	}
	return nil
}

func init() {
	proto.RegisterType((*Manifest)(nil), "istio.component.config.Manifest")
}

func init() { proto.RegisterFile("distrib/manifest.proto", fileDescriptorManifest) }

var fileDescriptorManifest = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0xd2, 0xcf, 0x4d, 0xcc, 0xcb, 0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xcb, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x4b, 0xce, 0xcf, 0x2d, 0xc8, 0xcf,
	0x4b, 0xcd, 0x2b, 0xd1, 0x4b, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x57, 0x6a, 0x67, 0xe4, 0xe2, 0xf0,
	0x85, 0x2a, 0x15, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0xca, 0x4c, 0x11, 0x52, 0xe4, 0xe2, 0x81, 0x6b, 0x88, 0xcf, 0x4c, 0x91, 0x60, 0x02, 0xcb, 0x70,
	0xc3, 0xc5, 0x3c, 0x53, 0x84, 0x54, 0xb9, 0xf8, 0x10, 0x4a, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x98,
	0xc1, 0x8a, 0x78, 0xe1, 0xa2, 0x21, 0x95, 0x05, 0xa9, 0x20, 0x93, 0xd2, 0x8a, 0x12, 0xd3, 0x73,
	0x21, 0x06, 0x15, 0x4b, 0xb0, 0x28, 0x30, 0x83, 0x4c, 0x82, 0x89, 0x79, 0xa6, 0x14, 0x3b, 0xa9,
	0x47, 0xa9, 0x42, 0xdc, 0x98, 0x99, 0xaf, 0x0f, 0x66, 0xe8, 0xa7, 0x27, 0xe6, 0xe4, 0xa4, 0x56,
	0xea, 0x17, 0x64, 0xa7, 0xeb, 0x27, 0x16, 0x64, 0xea, 0x43, 0x7d, 0x96, 0xc4, 0x06, 0xf6, 0x91,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x33, 0xf7, 0xda, 0x5a, 0xeb, 0x00, 0x00, 0x00,
}
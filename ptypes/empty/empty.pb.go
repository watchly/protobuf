// Code generated by protoc-gen-govalidate.
// source: github.com/watchly/protobuf/ptypes/empty/empty.proto
// DO NOT EDIT!

/*
Package empty is a generated protocol buffer package.

It is generated from these files:
	github.com/watchly/protobuf/ptypes/empty/empty.proto

It has these top-level messages:
	Empty
*/
package empty

import proto "github.com/watchly/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (*Empty) XXX_WellKnownType() string   { return "Empty" }

func init() {
	proto.RegisterType((*Empty)(nil), "google.protobuf.Empty")
}

func init() {
	proto.RegisterFile("github.com/watchly/protobuf/ptypes/empty/empty.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0xcd,
	0x2d, 0x28, 0xa9, 0x84, 0x90, 0x7a, 0x60, 0x39, 0x21, 0xfe, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54,
	0x3d, 0x98, 0x4a, 0x25, 0x76, 0x2e, 0x56, 0x57, 0x90, 0xbc, 0x53, 0x25, 0x97, 0x70, 0x72, 0x7e,
	0xae, 0x1e, 0x9a, 0xbc, 0x13, 0x17, 0x58, 0x36, 0x00, 0xc4, 0x0d, 0x60, 0x8c, 0x52, 0x27, 0xd2,
	0xce, 0x05, 0x8c, 0x8c, 0x3f, 0x18, 0x19, 0x17, 0x31, 0x31, 0xbb, 0x07, 0x38, 0xad, 0x62, 0x92,
	0x73, 0x87, 0x18, 0x1a, 0x00, 0x55, 0xaa, 0x17, 0x9e, 0x9a, 0x93, 0xe3, 0x9d, 0x97, 0x5f, 0x9e,
	0x17, 0x02, 0xd2, 0x92, 0xc4, 0x06, 0x36, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xbb,
	0xf4, 0x0e, 0xd2, 0x00, 0x00, 0x00,
}

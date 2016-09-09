// Code generated by protoc-gen-govalidate.
// source: github.com/watchly/protobuf/ptypes/duration/duration.proto
// DO NOT EDIT!

/*
Package duration is a generated protocol buffer package.

It is generated from these files:
	github.com/watchly/protobuf/ptypes/duration/duration.proto

It has these top-level messages:
	Duration
*/
package duration

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

// A Duration represents a signed, fixed-length span of time represented
// as a count of seconds and fractions of seconds at nanosecond
// resolution. It is independent of any calendar and concepts like "day"
// or "month". It is related to Timestamp in that the difference between
// two Timestamp values is a Duration and it can be added or subtracted
// from a Timestamp. Range is approximately +-10,000 years.
//
// Example 1: Compute Duration from two Timestamps in pseudo code.
//
//     Timestamp start = ...;
//     Timestamp end = ...;
//     Duration duration = ...;
//
//     duration.seconds = end.seconds - start.seconds;
//     duration.nanos = end.nanos - start.nanos;
//
//     if (duration.seconds < 0 && duration.nanos > 0) {
//       duration.seconds += 1;
//       duration.nanos -= 1000000000;
//     } else if (durations.seconds > 0 && duration.nanos < 0) {
//       duration.seconds -= 1;
//       duration.nanos += 1000000000;
//     }
//
// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
//
//     Timestamp start = ...;
//     Duration duration = ...;
//     Timestamp end = ...;
//
//     end.seconds = start.seconds + duration.seconds;
//     end.nanos = start.nanos + duration.nanos;
//
//     if (end.nanos < 0) {
//       end.seconds -= 1;
//       end.nanos += 1000000000;
//     } else if (end.nanos >= 1000000000) {
//       end.seconds += 1;
//       end.nanos -= 1000000000;
//     }
//
//
type Duration struct {
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos int32 `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
}

func (m *Duration) Reset()                    { *m = Duration{} }
func (m *Duration) String() string            { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()               {}
func (*Duration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (*Duration) XXX_WellKnownType() string   { return "Duration" }

func init() {
	proto.RegisterType((*Duration)(nil), "google.protobuf.Duration")
}

func init() {
	proto.RegisterFile("github.com/watchly/protobuf/ptypes/duration/duration.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x28,
	0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x29,
	0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0x33, 0xf4, 0xc0, 0x2a, 0x84, 0xf8, 0xd3, 0xf3, 0xf3,
	0xd3, 0x73, 0x52, 0xf5, 0x60, 0xea, 0x95, 0xac, 0xb8, 0x38, 0x5c, 0xa0, 0x4a, 0x84, 0x24, 0xb8,
	0xd8, 0x8b, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x8a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x60,
	0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xbc, 0xc4, 0xbc, 0xfc, 0x62, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x08, 0xc7, 0xa9, 0x86, 0x4b, 0x38, 0x39, 0x3f, 0x57, 0x0f, 0xcd, 0x48, 0x27, 0x5e, 0x98,
	0x81, 0x01, 0x20, 0x91, 0x00, 0xc6, 0x28, 0x2d, 0xe2, 0xdd, 0xbb, 0x80, 0x91, 0x71, 0x11, 0x13,
	0xb3, 0x7b, 0x80, 0xd3, 0x2a, 0x26, 0x39, 0x77, 0x88, 0xb9, 0x01, 0x50, 0xa5, 0x7a, 0xe1, 0xa9,
	0x39, 0x39, 0xde, 0x79, 0xf9, 0xe5, 0x79, 0x21, 0x20, 0x2d, 0x49, 0x6c, 0x60, 0x33, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0xfb, 0xb1, 0x51, 0x0e, 0x01, 0x00, 0x00,
}

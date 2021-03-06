// Code generated by protoc-gen-go.
// source: gopilosa_pbuf/public.proto
// DO NOT EDIT!

/*
Package gopilosa_pbuf is a generated protocol buffer package.

It is generated from these files:
	gopilosa_pbuf/public.proto

It has these top-level messages:
	Bitmap
	Pair
	SumCount
	Bit
	ColumnAttrSet
	Attr
	AttrMap
	QueryRequest
	QueryResponse
	QueryResult
	ImportRequest
	ImportValueRequest
*/
package gopilosa_pbuf

import "github.com/golang/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Bitmap struct {
	Bits  []uint64 `protobuf:"varint,1,rep,packed,name=Bits" json:"Bits,omitempty"`
	Attrs []*Attr  `protobuf:"bytes,2,rep,name=Attrs" json:"Attrs,omitempty"`
}

func (m *Bitmap) Reset()                    { *m = Bitmap{} }
func (m *Bitmap) String() string            { return proto.CompactTextString(m) }
func (*Bitmap) ProtoMessage()               {}
func (*Bitmap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bitmap) GetBits() []uint64 {
	if m != nil {
		return m.Bits
	}
	return nil
}

func (m *Bitmap) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Pair struct {
	Key   uint64 `protobuf:"varint,1,opt,name=Key" json:"Key,omitempty"`
	Count uint64 `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
}

func (m *Pair) Reset()                    { *m = Pair{} }
func (m *Pair) String() string            { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()               {}
func (*Pair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Pair) GetKey() uint64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Pair) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SumCount struct {
	Sum   int64 `protobuf:"varint,1,opt,name=Sum" json:"Sum,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=Count" json:"Count,omitempty"`
}

func (m *SumCount) Reset()                    { *m = SumCount{} }
func (m *SumCount) String() string            { return proto.CompactTextString(m) }
func (*SumCount) ProtoMessage()               {}
func (*SumCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SumCount) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *SumCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Bit struct {
	RowID     uint64 `protobuf:"varint,1,opt,name=RowID" json:"RowID,omitempty"`
	ColumnID  uint64 `protobuf:"varint,2,opt,name=ColumnID" json:"ColumnID,omitempty"`
	Timestamp int64  `protobuf:"varint,3,opt,name=Timestamp" json:"Timestamp,omitempty"`
}

func (m *Bit) Reset()                    { *m = Bit{} }
func (m *Bit) String() string            { return proto.CompactTextString(m) }
func (*Bit) ProtoMessage()               {}
func (*Bit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Bit) GetRowID() uint64 {
	if m != nil {
		return m.RowID
	}
	return 0
}

func (m *Bit) GetColumnID() uint64 {
	if m != nil {
		return m.ColumnID
	}
	return 0
}

func (m *Bit) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ColumnAttrSet struct {
	ID    uint64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Attrs []*Attr `protobuf:"bytes,2,rep,name=Attrs" json:"Attrs,omitempty"`
}

func (m *ColumnAttrSet) Reset()                    { *m = ColumnAttrSet{} }
func (m *ColumnAttrSet) String() string            { return proto.CompactTextString(m) }
func (*ColumnAttrSet) ProtoMessage()               {}
func (*ColumnAttrSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ColumnAttrSet) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ColumnAttrSet) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Attr struct {
	Key         string  `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Type        uint64  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	StringValue string  `protobuf:"bytes,3,opt,name=StringValue" json:"StringValue,omitempty"`
	IntValue    int64   `protobuf:"varint,4,opt,name=IntValue" json:"IntValue,omitempty"`
	BoolValue   bool    `protobuf:"varint,5,opt,name=BoolValue" json:"BoolValue,omitempty"`
	FloatValue  float64 `protobuf:"fixed64,6,opt,name=FloatValue" json:"FloatValue,omitempty"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Attr) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Attr) GetType() uint64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Attr) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Attr) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *Attr) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *Attr) GetFloatValue() float64 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

type AttrMap struct {
	Attrs []*Attr `protobuf:"bytes,1,rep,name=Attrs" json:"Attrs,omitempty"`
}

func (m *AttrMap) Reset()                    { *m = AttrMap{} }
func (m *AttrMap) String() string            { return proto.CompactTextString(m) }
func (*AttrMap) ProtoMessage()               {}
func (*AttrMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AttrMap) GetAttrs() []*Attr {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type QueryRequest struct {
	Query        string   `protobuf:"bytes,1,opt,name=Query" json:"Query,omitempty"`
	Slices       []uint64 `protobuf:"varint,2,rep,packed,name=Slices" json:"Slices,omitempty"`
	ColumnAttrs  bool     `protobuf:"varint,3,opt,name=ColumnAttrs" json:"ColumnAttrs,omitempty"`
	Remote       bool     `protobuf:"varint,5,opt,name=Remote" json:"Remote,omitempty"`
	ExcludeAttrs bool     `protobuf:"varint,6,opt,name=ExcludeAttrs" json:"ExcludeAttrs,omitempty"`
	ExcludeBits  bool     `protobuf:"varint,7,opt,name=ExcludeBits" json:"ExcludeBits,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *QueryRequest) GetSlices() []uint64 {
	if m != nil {
		return m.Slices
	}
	return nil
}

func (m *QueryRequest) GetColumnAttrs() bool {
	if m != nil {
		return m.ColumnAttrs
	}
	return false
}

func (m *QueryRequest) GetRemote() bool {
	if m != nil {
		return m.Remote
	}
	return false
}

func (m *QueryRequest) GetExcludeAttrs() bool {
	if m != nil {
		return m.ExcludeAttrs
	}
	return false
}

func (m *QueryRequest) GetExcludeBits() bool {
	if m != nil {
		return m.ExcludeBits
	}
	return false
}

type QueryResponse struct {
	Err            string           `protobuf:"bytes,1,opt,name=Err" json:"Err,omitempty"`
	Results        []*QueryResult   `protobuf:"bytes,2,rep,name=Results" json:"Results,omitempty"`
	ColumnAttrSets []*ColumnAttrSet `protobuf:"bytes,3,rep,name=ColumnAttrSets" json:"ColumnAttrSets,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *QueryResponse) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetColumnAttrSets() []*ColumnAttrSet {
	if m != nil {
		return m.ColumnAttrSets
	}
	return nil
}

type QueryResult struct {
	Bitmap   *Bitmap   `protobuf:"bytes,1,opt,name=Bitmap" json:"Bitmap,omitempty"`
	N        uint64    `protobuf:"varint,2,opt,name=N" json:"N,omitempty"`
	Pairs    []*Pair   `protobuf:"bytes,3,rep,name=Pairs" json:"Pairs,omitempty"`
	SumCount *SumCount `protobuf:"bytes,5,opt,name=SumCount" json:"SumCount,omitempty"`
	Changed  bool      `protobuf:"varint,4,opt,name=Changed" json:"Changed,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueryResult) GetBitmap() *Bitmap {
	if m != nil {
		return m.Bitmap
	}
	return nil
}

func (m *QueryResult) GetN() uint64 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *QueryResult) GetPairs() []*Pair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

func (m *QueryResult) GetSumCount() *SumCount {
	if m != nil {
		return m.SumCount
	}
	return nil
}

func (m *QueryResult) GetChanged() bool {
	if m != nil {
		return m.Changed
	}
	return false
}

type ImportRequest struct {
	Index      string   `protobuf:"bytes,1,opt,name=Index" json:"Index,omitempty"`
	Frame      string   `protobuf:"bytes,2,opt,name=Frame" json:"Frame,omitempty"`
	Slice      uint64   `protobuf:"varint,3,opt,name=Slice" json:"Slice,omitempty"`
	RowIDs     []uint64 `protobuf:"varint,4,rep,packed,name=RowIDs" json:"RowIDs,omitempty"`
	ColumnIDs  []uint64 `protobuf:"varint,5,rep,packed,name=ColumnIDs" json:"ColumnIDs,omitempty"`
	Timestamps []int64  `protobuf:"varint,6,rep,packed,name=Timestamps" json:"Timestamps,omitempty"`
}

func (m *ImportRequest) Reset()                    { *m = ImportRequest{} }
func (m *ImportRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportRequest) ProtoMessage()               {}
func (*ImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ImportRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ImportRequest) GetFrame() string {
	if m != nil {
		return m.Frame
	}
	return ""
}

func (m *ImportRequest) GetSlice() uint64 {
	if m != nil {
		return m.Slice
	}
	return 0
}

func (m *ImportRequest) GetRowIDs() []uint64 {
	if m != nil {
		return m.RowIDs
	}
	return nil
}

func (m *ImportRequest) GetColumnIDs() []uint64 {
	if m != nil {
		return m.ColumnIDs
	}
	return nil
}

func (m *ImportRequest) GetTimestamps() []int64 {
	if m != nil {
		return m.Timestamps
	}
	return nil
}

type ImportValueRequest struct {
	Index     string   `protobuf:"bytes,1,opt,name=Index" json:"Index,omitempty"`
	Frame     string   `protobuf:"bytes,2,opt,name=Frame" json:"Frame,omitempty"`
	Slice     uint64   `protobuf:"varint,3,opt,name=Slice" json:"Slice,omitempty"`
	Field     string   `protobuf:"bytes,4,opt,name=Field" json:"Field,omitempty"`
	ColumnIDs []uint64 `protobuf:"varint,5,rep,packed,name=ColumnIDs" json:"ColumnIDs,omitempty"`
	Values    []uint64 `protobuf:"varint,6,rep,packed,name=Values" json:"Values,omitempty"`
}

func (m *ImportValueRequest) Reset()                    { *m = ImportValueRequest{} }
func (m *ImportValueRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportValueRequest) ProtoMessage()               {}
func (*ImportValueRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ImportValueRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ImportValueRequest) GetFrame() string {
	if m != nil {
		return m.Frame
	}
	return ""
}

func (m *ImportValueRequest) GetSlice() uint64 {
	if m != nil {
		return m.Slice
	}
	return 0
}

func (m *ImportValueRequest) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ImportValueRequest) GetColumnIDs() []uint64 {
	if m != nil {
		return m.ColumnIDs
	}
	return nil
}

func (m *ImportValueRequest) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Bitmap)(nil), "gopilosa_pbuf.Bitmap")
	proto.RegisterType((*Pair)(nil), "gopilosa_pbuf.Pair")
	proto.RegisterType((*SumCount)(nil), "gopilosa_pbuf.SumCount")
	proto.RegisterType((*Bit)(nil), "gopilosa_pbuf.Bit")
	proto.RegisterType((*ColumnAttrSet)(nil), "gopilosa_pbuf.ColumnAttrSet")
	proto.RegisterType((*Attr)(nil), "gopilosa_pbuf.Attr")
	proto.RegisterType((*AttrMap)(nil), "gopilosa_pbuf.AttrMap")
	proto.RegisterType((*QueryRequest)(nil), "gopilosa_pbuf.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "gopilosa_pbuf.QueryResponse")
	proto.RegisterType((*QueryResult)(nil), "gopilosa_pbuf.QueryResult")
	proto.RegisterType((*ImportRequest)(nil), "gopilosa_pbuf.ImportRequest")
	proto.RegisterType((*ImportValueRequest)(nil), "gopilosa_pbuf.ImportValueRequest")
}

func init() { proto.RegisterFile("gopilosa_pbuf/public.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xd5, 0xc4, 0xce, 0xeb, 0xa6, 0xa9, 0xd0, 0x00, 0xc5, 0xaa, 0x2a, 0x14, 0x79, 0x15, 0x16,
	0x04, 0x29, 0xed, 0x0f, 0x90, 0xa6, 0x45, 0x01, 0x51, 0xc1, 0xa4, 0xb0, 0x45, 0x6e, 0x3b, 0x14,
	0x4b, 0xb6, 0xc7, 0xd8, 0x33, 0xa2, 0xf9, 0x18, 0x36, 0x6c, 0xe0, 0x23, 0x58, 0xf2, 0x61, 0xe8,
	0xde, 0xf1, 0xc4, 0x8e, 0x37, 0x15, 0x12, 0xbb, 0x39, 0xf7, 0xe5, 0x73, 0x1f, 0xc7, 0x70, 0x78,
	0xab, 0xf2, 0x38, 0x51, 0x65, 0xf4, 0x29, 0xbf, 0x32, 0x9f, 0x5f, 0xe4, 0xe6, 0x2a, 0x89, 0xaf,
	0x67, 0x79, 0xa1, 0xb4, 0xe2, 0xe3, 0x1d, 0x5f, 0xf8, 0x0a, 0x7a, 0x8b, 0x58, 0xa7, 0x51, 0xce,
	0x39, 0xf8, 0x8b, 0x58, 0x97, 0x01, 0x9b, 0x78, 0x53, 0x5f, 0xd0, 0x9b, 0x3f, 0x83, 0xee, 0x4b,
	0xad, 0x8b, 0x32, 0xe8, 0x4c, 0xbc, 0xe9, 0x68, 0xfe, 0x70, 0xb6, 0x93, 0x3c, 0x43, 0x9f, 0xb0,
	0x11, 0xe1, 0x0c, 0xfc, 0x77, 0x51, 0x5c, 0xf0, 0x07, 0xe0, 0xbd, 0x91, 0x9b, 0x80, 0x4d, 0xd8,
	0xd4, 0x17, 0xf8, 0xe4, 0x8f, 0xa0, 0x7b, 0xaa, 0x4c, 0xa6, 0x83, 0x0e, 0xd9, 0x2c, 0x08, 0xe7,
	0x30, 0x58, 0x9b, 0x94, 0xde, 0x98, 0xb3, 0x36, 0x29, 0xe5, 0x78, 0x02, 0x9f, 0xbb, 0x39, 0x9e,
	0xcb, 0xf9, 0x00, 0xde, 0x22, 0xd6, 0xe8, 0x14, 0xea, 0xdb, 0x6a, 0x59, 0x7d, 0xc4, 0x02, 0x7e,
	0x08, 0x83, 0x53, 0x95, 0x98, 0x34, 0x5b, 0x2d, 0xab, 0x2f, 0x6d, 0x31, 0x3f, 0x82, 0xe1, 0x65,
	0x9c, 0xca, 0x52, 0x47, 0x69, 0x1e, 0x78, 0x54, 0xb2, 0x36, 0x84, 0xaf, 0x61, 0x6c, 0x23, 0xb1,
	0x93, 0xb5, 0xd4, 0x7c, 0x1f, 0x3a, 0xdb, 0xea, 0x9d, 0xd5, 0xf2, 0x5f, 0xc6, 0xf0, 0x8b, 0x81,
	0x8f, 0xaf, 0xe6, 0x1c, 0x86, 0x76, 0x0e, 0x1c, 0xfc, 0xcb, 0x4d, 0x2e, 0x2b, 0x72, 0xf4, 0xe6,
	0x13, 0x18, 0xad, 0x75, 0x11, 0x67, 0xb7, 0x1f, 0xa3, 0xc4, 0x48, 0xa2, 0x36, 0x14, 0x4d, 0x13,
	0xb6, 0xb5, 0xca, 0xb4, 0x75, 0xfb, 0xc4, 0x7c, 0x8b, 0xb1, 0xad, 0x85, 0x52, 0x89, 0x75, 0x76,
	0x27, 0x6c, 0x3a, 0x10, 0xb5, 0x81, 0x3f, 0x05, 0x38, 0x4f, 0x54, 0x54, 0xe5, 0xf6, 0x26, 0x6c,
	0xca, 0x44, 0xc3, 0x12, 0x9e, 0x40, 0x1f, 0x99, 0xbe, 0x8d, 0xf2, 0xba, 0x41, 0x76, 0x6f, 0x83,
	0xbf, 0x19, 0xec, 0xbd, 0x37, 0xb2, 0xd8, 0x08, 0xf9, 0xd5, 0xc8, 0x92, 0xb6, 0x41, 0xb8, 0x6a,
	0xd5, 0x02, 0x7e, 0x00, 0xbd, 0x75, 0x12, 0x5f, 0x4b, 0x3b, 0x33, 0x5f, 0x54, 0x08, 0x1b, 0xae,
	0x67, 0x5d, 0x52, 0xc3, 0x03, 0xd1, 0x34, 0x61, 0xa6, 0x90, 0xa9, 0xd2, 0xae, 0xa3, 0x0a, 0xf1,
	0x10, 0xf6, 0xce, 0xee, 0xae, 0x13, 0x73, 0x23, 0x6d, 0x6a, 0x8f, 0xbc, 0x3b, 0x36, 0xac, 0x5e,
	0x61, 0x3a, 0xe5, 0xbe, 0xad, 0xde, 0x30, 0x85, 0xdf, 0x19, 0x8c, 0x2b, 0xfa, 0x65, 0xae, 0xb2,
	0x52, 0xe2, 0xa2, 0xce, 0x8a, 0xc2, 0x2d, 0xea, 0xac, 0x28, 0xf8, 0x09, 0xf4, 0x85, 0x2c, 0x4d,
	0xa2, 0xdd, 0xc2, 0x0f, 0x5b, 0xf3, 0x70, 0x05, 0x4c, 0xa2, 0x85, 0x0b, 0xe5, 0x4b, 0xd8, 0xdf,
	0xb9, 0x22, 0x6c, 0x0e, 0x93, 0x8f, 0x5a, 0xc9, 0x3b, 0x41, 0xa2, 0x95, 0x13, 0xfe, 0x61, 0x30,
	0x6a, 0x94, 0xe7, 0xcf, 0x9d, 0x3e, 0x89, 0xe0, 0x68, 0xfe, 0xb8, 0x55, 0xcd, 0x3a, 0x85, 0x13,
	0xf1, 0x1e, 0xb0, 0x8b, 0xea, 0xc0, 0xd8, 0x05, 0xae, 0x15, 0x35, 0xe9, 0x98, 0xb4, 0xd7, 0x8a,
	0x3e, 0x61, 0x23, 0xf8, 0x71, 0x2d, 0x47, 0x9a, 0xfb, 0x68, 0xfe, 0xa4, 0x15, 0xed, 0xdc, 0xa2,
	0xd6, 0x6d, 0x00, 0xfd, 0xd3, 0x2f, 0x51, 0x76, 0x2b, 0x6f, 0xe8, 0x34, 0x07, 0xc2, 0xc1, 0xf0,
	0x27, 0x83, 0xf1, 0x2a, 0xcd, 0x55, 0xa1, 0x1b, 0x67, 0xb2, 0xca, 0x6e, 0xe4, 0x9d, 0x3b, 0x13,
	0x02, 0x68, 0x3d, 0x2f, 0xa2, 0xd4, 0x8a, 0x62, 0x28, 0x2c, 0x40, 0x2b, 0x9d, 0x0b, 0x9d, 0x87,
	0x2f, 0x2c, 0xa0, 0xc3, 0x40, 0xa5, 0x97, 0x81, 0x6f, 0x4f, 0xca, 0x22, 0x54, 0x81, 0x13, 0x7a,
	0x19, 0x74, 0xc9, 0x55, 0x1b, 0x50, 0x05, 0x5b, 0xa5, 0xe3, 0xd1, 0x78, 0x53, 0x4f, 0x34, 0x2c,
	0xe1, 0x0f, 0x06, 0xdc, 0x32, 0x25, 0x55, 0xfc, 0x3f, 0xba, 0x18, 0x1b, 0xcb, 0xc4, 0x8e, 0x06,
	0x63, 0x11, 0xdc, 0x43, 0xf6, 0x00, 0x7a, 0xc4, 0xc2, 0x12, 0xf5, 0x45, 0x85, 0xae, 0x7a, 0xf4,
	0xef, 0x3e, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x1b, 0x7f, 0x47, 0xd9, 0x05, 0x00, 0x00,
}

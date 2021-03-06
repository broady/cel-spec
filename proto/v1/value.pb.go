// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/value.proto

package v1 // import "github.com/google/cel-spec/proto/v1"

/*
Contains representations for CEL runtime values.
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Builtin type designators for CEL primitive types.
type TypeValue_BasicType int32

const (
	// Unspecified type.
	TypeValue_BASIC_TYPE_UNSPECIFIED TypeValue_BasicType = 0
	// The null type.
	TypeValue_NULL_TYPE TypeValue_BasicType = 1
	// The boolean type.
	TypeValue_BOOL_TYPE TypeValue_BasicType = 2
	// The signed integer type.
	//
	// Proto-based integer values are widened to int64.
	TypeValue_INT_TYPE TypeValue_BasicType = 3
	// The unsigned integer type.
	//
	// Proto-based unsigned integer values are widened to uint64.
	TypeValue_UINT_TYPE TypeValue_BasicType = 4
	// The double type.
	//
	// Proto-based float values are widened to double values.
	TypeValue_DOUBLE_TYPE TypeValue_BasicType = 5
	// The string type.
	TypeValue_STRING_TYPE TypeValue_BasicType = 6
	// The bytes type.
	TypeValue_BYTES_TYPE TypeValue_BasicType = 7
	// The type type.
	TypeValue_TYPE_TYPE TypeValue_BasicType = 8
	// The map type.
	TypeValue_MAP_TYPE TypeValue_BasicType = 9
	// The list type.
	TypeValue_LIST_TYPE TypeValue_BasicType = 10
)

var TypeValue_BasicType_name = map[int32]string{
	0:  "BASIC_TYPE_UNSPECIFIED",
	1:  "NULL_TYPE",
	2:  "BOOL_TYPE",
	3:  "INT_TYPE",
	4:  "UINT_TYPE",
	5:  "DOUBLE_TYPE",
	6:  "STRING_TYPE",
	7:  "BYTES_TYPE",
	8:  "TYPE_TYPE",
	9:  "MAP_TYPE",
	10: "LIST_TYPE",
}
var TypeValue_BasicType_value = map[string]int32{
	"BASIC_TYPE_UNSPECIFIED": 0,
	"NULL_TYPE":              1,
	"BOOL_TYPE":              2,
	"INT_TYPE":               3,
	"UINT_TYPE":              4,
	"DOUBLE_TYPE":            5,
	"STRING_TYPE":            6,
	"BYTES_TYPE":             7,
	"TYPE_TYPE":              8,
	"MAP_TYPE":               9,
	"LIST_TYPE":              10,
}

func (x TypeValue_BasicType) String() string {
	return proto.EnumName(TypeValue_BasicType_name, int32(x))
}
func (TypeValue_BasicType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_value_37fb3afa46b8f581, []int{3, 0}
}

// Represents a CEL value.
//
// This is similar to `google.protobuf.Value`, but can represent CEL's full
// range of values.
type Value struct {
	// Required. The valid kinds of values.
	//
	// Types that are valid to be assigned to Kind:
	//	*Value_NullValue
	//	*Value_BoolValue
	//	*Value_Int64Value
	//	*Value_Uint64Value
	//	*Value_DoubleValue
	//	*Value_StringValue
	//	*Value_BytesValue
	//	*Value_ObjectValue
	//	*Value_MapValue
	//	*Value_ListValue
	//	*Value_TypeValue
	Kind                 isValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_37fb3afa46b8f581, []int{0}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	NullValue _struct.NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=google.protobuf.NullValue,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Value_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,4,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,6,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_ObjectValue struct {
	ObjectValue *any.Any `protobuf:"bytes,10,opt,name=object_value,json=objectValue,proto3,oneof"`
}

type Value_MapValue struct {
	MapValue *MapValue `protobuf:"bytes,11,opt,name=map_value,json=mapValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,12,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_TypeValue struct {
	TypeValue *TypeValue `protobuf:"bytes,15,opt,name=type_value,json=typeValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_Int64Value) isValue_Kind() {}

func (*Value_Uint64Value) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BytesValue) isValue_Kind() {}

func (*Value_ObjectValue) isValue_Kind() {}

func (*Value_MapValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_TypeValue) isValue_Kind() {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetNullValue() _struct.NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return _struct.NullValue_NULL_VALUE
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetInt64Value() int64 {
	if x, ok := m.GetKind().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Value) GetUint64Value() uint64 {
	if x, ok := m.GetKind().(*Value_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if x, ok := m.GetKind().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBytesValue() []byte {
	if x, ok := m.GetKind().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (m *Value) GetObjectValue() *any.Any {
	if x, ok := m.GetKind().(*Value_ObjectValue); ok {
		return x.ObjectValue
	}
	return nil
}

func (m *Value) GetMapValue() *MapValue {
	if x, ok := m.GetKind().(*Value_MapValue); ok {
		return x.MapValue
	}
	return nil
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (m *Value) GetTypeValue() *TypeValue {
	if x, ok := m.GetKind().(*Value_TypeValue); ok {
		return x.TypeValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_NullValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_Int64Value)(nil),
		(*Value_Uint64Value)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_ObjectValue)(nil),
		(*Value_MapValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_TypeValue)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.NullValue))
	case *Value_BoolValue:
		t := uint64(0)
		if x.BoolValue {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_Int64Value:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Value))
	case *Value_Uint64Value:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Uint64Value))
	case *Value_DoubleValue:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleValue))
	case *Value_StringValue:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringValue)
	case *Value_BytesValue:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesValue)
	case *Value_ObjectValue:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ObjectValue); err != nil {
			return err
		}
	case *Value_MapValue:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MapValue); err != nil {
			return err
		}
	case *Value_ListValue:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ListValue); err != nil {
			return err
		}
	case *Value_TypeValue:
		b.EncodeVarint(15<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypeValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Kind has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // kind.null_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_NullValue{_struct.NullValue(x)}
		return true, err
	case 2: // kind.bool_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_BoolValue{x != 0}
		return true, err
	case 3: // kind.int64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_Int64Value{int64(x)}
		return true, err
	case 4: // kind.uint64_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Kind = &Value_Uint64Value{x}
		return true, err
	case 5: // kind.double_value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Kind = &Value_DoubleValue{math.Float64frombits(x)}
		return true, err
	case 6: // kind.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Kind = &Value_StringValue{x}
		return true, err
	case 7: // kind.bytes_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Kind = &Value_BytesValue{x}
		return true, err
	case 10: // kind.object_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(any.Any)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_ObjectValue{msg}
		return true, err
	case 11: // kind.map_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MapValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_MapValue{msg}
		return true, err
	case 12: // kind.list_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ListValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_ListValue{msg}
		return true, err
	case 15: // kind.type_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TypeValue)
		err := b.DecodeMessage(msg)
		m.Kind = &Value_TypeValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// kind
	switch x := m.Kind.(type) {
	case *Value_NullValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.NullValue))
	case *Value_BoolValue:
		n += 1 // tag and wire
		n += 1
	case *Value_Int64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64Value))
	case *Value_Uint64Value:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Uint64Value))
	case *Value_DoubleValue:
		n += 1 // tag and wire
		n += 8
	case *Value_StringValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringValue)))
		n += len(x.StringValue)
	case *Value_BytesValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BytesValue)))
		n += len(x.BytesValue)
	case *Value_ObjectValue:
		s := proto.Size(x.ObjectValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_MapValue:
		s := proto.Size(x.MapValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_ListValue:
		s := proto.Size(x.ListValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Value_TypeValue:
		s := proto.Size(x.TypeValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A list.
//
// Wrapped in a message so 'not set' and empty can be differentiated, which is
// required for use in a 'oneof'.
type ListValue struct {
	// The ordered values in the list.
	Values               []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_37fb3afa46b8f581, []int{1}
}
func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListValue.Unmarshal(m, b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
}
func (dst *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(dst, src)
}
func (m *ListValue) XXX_Size() int {
	return xxx_messageInfo_ListValue.Size(m)
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

// A map.
//
// Wrapped in a message so 'not set' and empty can be differentiated, which is
// required for use in a 'oneof'.
type MapValue struct {
	// The set of map entries.
	//
	// CEL has fewer restrictions on keys, so a protobuf map representation
	// cannot be used.
	Entries              []*MapValue_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_37fb3afa46b8f581, []int{2}
}
func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (dst *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(dst, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetEntries() []*MapValue_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type MapValue_Entry struct {
	// The key.
	//
	// Must be unique with in the map.
	// Currently only boolean, int, uint, and string values can be keys.
	Key *Value `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value.
	Value                *Value   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapValue_Entry) Reset()         { *m = MapValue_Entry{} }
func (m *MapValue_Entry) String() string { return proto.CompactTextString(m) }
func (*MapValue_Entry) ProtoMessage()    {}
func (*MapValue_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_37fb3afa46b8f581, []int{2, 0}
}
func (m *MapValue_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue_Entry.Unmarshal(m, b)
}
func (m *MapValue_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue_Entry.Marshal(b, m, deterministic)
}
func (dst *MapValue_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue_Entry.Merge(dst, src)
}
func (m *MapValue_Entry) XXX_Size() int {
	return xxx_messageInfo_MapValue_Entry.Size(m)
}
func (m *MapValue_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue_Entry proto.InternalMessageInfo

func (m *MapValue_Entry) GetKey() *Value {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MapValue_Entry) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// A type value.
//
// TypeValue is trivially derivable from a Value.
//
type TypeValue struct {
	// The type designator.
	//
	// Type designators are split into two categories due to representational
	// and ease of use considerations. This split has no semantic implications.
	//
	// Types that are valid to be assigned to DesignatorKind:
	//	*TypeValue_BasicType_
	//	*TypeValue_ObjectType
	DesignatorKind       isTypeValue_DesignatorKind `protobuf_oneof:"designator_kind"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TypeValue) Reset()         { *m = TypeValue{} }
func (m *TypeValue) String() string { return proto.CompactTextString(m) }
func (*TypeValue) ProtoMessage()    {}
func (*TypeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_value_37fb3afa46b8f581, []int{3}
}
func (m *TypeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeValue.Unmarshal(m, b)
}
func (m *TypeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeValue.Marshal(b, m, deterministic)
}
func (dst *TypeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeValue.Merge(dst, src)
}
func (m *TypeValue) XXX_Size() int {
	return xxx_messageInfo_TypeValue.Size(m)
}
func (m *TypeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeValue.DiscardUnknown(m)
}

var xxx_messageInfo_TypeValue proto.InternalMessageInfo

type isTypeValue_DesignatorKind interface {
	isTypeValue_DesignatorKind()
}

type TypeValue_BasicType_ struct {
	BasicType TypeValue_BasicType `protobuf:"varint,1,opt,name=basic_type,json=basicType,proto3,enum=google.api.expr.v1.TypeValue_BasicType,oneof"`
}

type TypeValue_ObjectType struct {
	ObjectType string `protobuf:"bytes,6,opt,name=object_type,json=objectType,proto3,oneof"`
}

func (*TypeValue_BasicType_) isTypeValue_DesignatorKind() {}

func (*TypeValue_ObjectType) isTypeValue_DesignatorKind() {}

func (m *TypeValue) GetDesignatorKind() isTypeValue_DesignatorKind {
	if m != nil {
		return m.DesignatorKind
	}
	return nil
}

func (m *TypeValue) GetBasicType() TypeValue_BasicType {
	if x, ok := m.GetDesignatorKind().(*TypeValue_BasicType_); ok {
		return x.BasicType
	}
	return TypeValue_BASIC_TYPE_UNSPECIFIED
}

func (m *TypeValue) GetObjectType() string {
	if x, ok := m.GetDesignatorKind().(*TypeValue_ObjectType); ok {
		return x.ObjectType
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TypeValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TypeValue_OneofMarshaler, _TypeValue_OneofUnmarshaler, _TypeValue_OneofSizer, []interface{}{
		(*TypeValue_BasicType_)(nil),
		(*TypeValue_ObjectType)(nil),
	}
}

func _TypeValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TypeValue)
	// designator_kind
	switch x := m.DesignatorKind.(type) {
	case *TypeValue_BasicType_:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.BasicType))
	case *TypeValue_ObjectType:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ObjectType)
	case nil:
	default:
		return fmt.Errorf("TypeValue.DesignatorKind has unexpected type %T", x)
	}
	return nil
}

func _TypeValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TypeValue)
	switch tag {
	case 1: // designator_kind.basic_type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.DesignatorKind = &TypeValue_BasicType_{TypeValue_BasicType(x)}
		return true, err
	case 6: // designator_kind.object_type
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.DesignatorKind = &TypeValue_ObjectType{x}
		return true, err
	default:
		return false, nil
	}
}

func _TypeValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TypeValue)
	// designator_kind
	switch x := m.DesignatorKind.(type) {
	case *TypeValue_BasicType_:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.BasicType))
	case *TypeValue_ObjectType:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ObjectType)))
		n += len(x.ObjectType)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Value)(nil), "google.api.expr.v1.Value")
	proto.RegisterType((*ListValue)(nil), "google.api.expr.v1.ListValue")
	proto.RegisterType((*MapValue)(nil), "google.api.expr.v1.MapValue")
	proto.RegisterType((*MapValue_Entry)(nil), "google.api.expr.v1.MapValue.Entry")
	proto.RegisterType((*TypeValue)(nil), "google.api.expr.v1.TypeValue")
	proto.RegisterEnum("google.api.expr.v1.TypeValue_BasicType", TypeValue_BasicType_name, TypeValue_BasicType_value)
}

func init() { proto.RegisterFile("proto/v1/value.proto", fileDescriptor_value_37fb3afa46b8f581) }

var fileDescriptor_value_37fb3afa46b8f581 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdb, 0x6e, 0xd3, 0x4c,
	0x14, 0x85, 0x33, 0xcd, 0xa1, 0xf1, 0x4e, 0xfe, 0xb6, 0xff, 0xa8, 0xaa, 0xda, 0xa8, 0x08, 0x93,
	0x5e, 0x60, 0x09, 0x61, 0xab, 0x05, 0x21, 0x21, 0x50, 0xa5, 0xba, 0x0d, 0x24, 0x52, 0x9a, 0x46,
	0x4e, 0x82, 0x54, 0x6e, 0x22, 0xdb, 0x1d, 0x82, 0xa9, 0x63, 0x5b, 0xf6, 0xb8, 0xc2, 0xaf, 0xc2,
	0x4b, 0xf0, 0x0a, 0x3c, 0x01, 0xcf, 0xc3, 0x25, 0x9a, 0x93, 0x39, 0x34, 0xe4, 0x2e, 0x6b, 0xed,
	0x6f, 0xed, 0x99, 0x6c, 0x6f, 0x1b, 0x76, 0x93, 0x34, 0xa6, 0xb1, 0x75, 0x77, 0x6c, 0xdd, 0xb9,
	0x61, 0x4e, 0x4c, 0x2e, 0x31, 0x5e, 0xc4, 0xf1, 0x22, 0x24, 0xa6, 0x9b, 0x04, 0x26, 0xf9, 0x9c,
	0xa4, 0xe6, 0xdd, 0x71, 0xe7, 0x40, 0x78, 0x16, 0x27, 0xbc, 0xfc, 0x83, 0xe5, 0x46, 0x85, 0xc0,
	0x3b, 0x87, 0x7f, 0x97, 0x32, 0x9a, 0xe6, 0x3e, 0x15, 0xd5, 0xee, 0x97, 0x1a, 0xd4, 0xdf, 0xb1,
	0xe6, 0xf8, 0x15, 0x40, 0x94, 0x87, 0xe1, 0x9c, 0x1f, 0xb5, 0x8f, 0x74, 0x64, 0x6c, 0x9d, 0x74,
	0x4c, 0x79, 0x96, 0x0a, 0x9b, 0xa3, 0x3c, 0x0c, 0x39, 0xdf, 0xaf, 0x38, 0x5a, 0xa4, 0x04, 0x7e,
	0x08, 0xe0, 0xc5, 0xb1, 0x0a, 0x6f, 0xe8, 0xc8, 0x68, 0x32, 0x80, 0x79, 0x02, 0x78, 0x04, 0xad,
	0x20, 0xa2, 0x2f, 0x9e, 0x4b, 0xa2, 0xaa, 0x23, 0xa3, 0xda, 0xaf, 0x38, 0xc0, 0x4d, 0x81, 0x1c,
	0x41, 0x3b, 0xff, 0x9d, 0xa9, 0xe9, 0xc8, 0xa8, 0xf5, 0x2b, 0x4e, 0x2b, 0xff, 0x13, 0xba, 0x89,
	0x73, 0x2f, 0x24, 0x12, 0xaa, 0xeb, 0xc8, 0x40, 0x0c, 0x12, 0x6e, 0x09, 0x65, 0x34, 0x0d, 0xa2,
	0x85, 0x84, 0x1a, 0x3a, 0x32, 0x34, 0x06, 0x09, 0xb7, 0xbc, 0x91, 0x57, 0x50, 0x92, 0x49, 0x66,
	0x53, 0x47, 0x46, 0x9b, 0xdd, 0x88, 0x9b, 0x02, 0x79, 0x09, 0xed, 0xd8, 0xfb, 0x44, 0x7c, 0x2a,
	0x19, 0xd0, 0x91, 0xd1, 0x3a, 0xd9, 0xbd, 0x37, 0x94, 0xb3, 0xa8, 0x60, 0xdd, 0x05, 0xab, 0xa6,
	0xa9, 0x2d, 0xdd, 0x44, 0xe6, 0x5a, 0x3c, 0x77, 0x68, 0xde, 0x7f, 0x70, 0xe6, 0xa5, 0x9b, 0xa8,
	0x71, 0x36, 0x97, 0xf2, 0x37, 0x3e, 0x05, 0x08, 0x83, 0x4c, 0x9d, 0xda, 0xe6, 0xe9, 0x07, 0xab,
	0xd2, 0xc3, 0x20, 0xa3, 0xe5, 0xd3, 0x08, 0x95, 0x60, 0x79, 0x5a, 0x24, 0x6a, 0x44, 0xdb, 0xff,
	0xce, 0x4f, 0x8b, 0x84, 0x94, 0x79, 0xaa, 0x84, 0xdd, 0x80, 0xda, 0x6d, 0x10, 0xdd, 0x74, 0x4f,
	0x41, 0x2b, 0x4f, 0xc0, 0xc7, 0xd0, 0xe0, 0xfd, 0xb2, 0x7d, 0xa4, 0x57, 0x8d, 0xd6, 0xc9, 0xc1,
	0xaa, 0x86, 0x1c, 0x75, 0x24, 0xd8, 0xfd, 0x8a, 0xa0, 0xa9, 0xfe, 0x20, 0x7e, 0x0d, 0x9b, 0x24,
	0xa2, 0x69, 0x50, 0x36, 0xe8, 0xae, 0x9b, 0x87, 0xd9, 0x8b, 0x68, 0x5a, 0x38, 0x2a, 0xd2, 0x21,
	0x50, 0xe7, 0x0e, 0x7e, 0x02, 0xd5, 0x5b, 0x52, 0xf0, 0xfd, 0x5c, 0x7b, 0x07, 0x46, 0x61, 0x0b,
	0xea, 0xbf, 0x36, 0x72, 0x2d, 0x2e, 0xb8, 0xee, 0xf7, 0x0d, 0xd0, 0xca, 0xa1, 0xe0, 0x3e, 0x80,
	0xe7, 0x66, 0x81, 0x3f, 0x67, 0xa3, 0x91, 0xaf, 0xc4, 0xe3, 0xb5, 0x73, 0x34, 0x6d, 0xc6, 0x33,
	0xc9, 0xd7, 0x5f, 0x09, 0xb6, 0x6c, 0x72, 0x93, 0x78, 0x2b, 0xb5, 0x90, 0x20, 0x4c, 0x86, 0x74,
	0xbf, 0x21, 0xd0, 0xca, 0x34, 0xee, 0xc0, 0x9e, 0x7d, 0x36, 0x19, 0x9c, 0xcf, 0xa7, 0xd7, 0xe3,
	0xde, 0x7c, 0x36, 0x9a, 0x8c, 0x7b, 0xe7, 0x83, 0x37, 0x83, 0xde, 0xc5, 0x4e, 0x05, 0xff, 0x07,
	0xda, 0x68, 0x36, 0x1c, 0xf2, 0xd2, 0x0e, 0x62, 0xd2, 0xbe, 0xba, 0x92, 0x72, 0x03, 0xb7, 0xa1,
	0x39, 0x18, 0x4d, 0x85, 0xaa, 0xb2, 0xe2, 0xac, 0x94, 0x35, 0xbc, 0x0d, 0xad, 0x8b, 0xab, 0x99,
	0x3d, 0xec, 0x09, 0xa3, 0xce, 0x8c, 0xc9, 0xd4, 0x19, 0x8c, 0xde, 0x0a, 0xa3, 0x81, 0xb7, 0x00,
	0xec, 0xeb, 0x69, 0x6f, 0x22, 0xf4, 0x26, 0x6b, 0xc0, 0xaf, 0xc0, 0x65, 0x93, 0x75, 0xbf, 0x3c,
	0x1b, 0x0b, 0xa5, 0xb1, 0xe2, 0x70, 0x30, 0x91, 0xdd, 0xc1, 0xfe, 0x1f, 0xb6, 0x6f, 0x48, 0x16,
	0x2c, 0x22, 0x97, 0xc6, 0xe9, 0x9c, 0xad, 0x90, 0x3d, 0x82, 0x3d, 0x3f, 0x5e, 0xae, 0x98, 0x99,
	0x0d, 0x7c, 0x60, 0x63, 0xf6, 0x0e, 0x8d, 0xd1, 0xfb, 0xa3, 0x45, 0x40, 0x3f, 0xe6, 0x9e, 0xe9,
	0xc7, 0x4b, 0x4b, 0x7e, 0xb0, 0x7c, 0x12, 0x3e, 0xcd, 0x12, 0xe2, 0x5b, 0xea, 0x2b, 0xf8, 0x03,
	0x21, 0xaf, 0xc1, 0xc5, 0xb3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xaf, 0x26, 0xb6, 0x1b,
	0x05, 0x00, 0x00,
}

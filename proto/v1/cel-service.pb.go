// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/cel-service.proto

package v1 // import "github.com/google/cel-spec/proto/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import status "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Severities of issues.
type StatusDetails_Severity int32

const (
	// An unspecified severity.
	StatusDetails_SEVERITY_UNSPECIFIED StatusDetails_Severity = 0
	// Deprecation issue for statements and method that may no longer be
	// supported or maintained.
	StatusDetails_DEPRECATION StatusDetails_Severity = 1
	// Warnings such as: unused variables.
	StatusDetails_WARNING StatusDetails_Severity = 2
	// Errors such as: unmatched curly braces or variable redefinition.
	StatusDetails_ERROR StatusDetails_Severity = 3
)

var StatusDetails_Severity_name = map[int32]string{
	0: "SEVERITY_UNSPECIFIED",
	1: "DEPRECATION",
	2: "WARNING",
	3: "ERROR",
}
var StatusDetails_Severity_value = map[string]int32{
	"SEVERITY_UNSPECIFIED": 0,
	"DEPRECATION":          1,
	"WARNING":              2,
	"ERROR":                3,
}

func (x StatusDetails_Severity) String() string {
	return proto.EnumName(StatusDetails_Severity_name, int32(x))
}
func (StatusDetails_Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{0, 0}
}

// Warnings or errors in service execution are represented by
// [google.rpc.Status][] messages, with the following message
// in the details field.
type StatusDetails struct {
	// The severity of the issue.
	Severity StatusDetails_Severity `protobuf:"varint,1,opt,name=severity,proto3,enum=google.api.expr.v1.StatusDetails_Severity" json:"severity,omitempty"`
	// The 1-based index of the starting line in the source text
	// where the issue occurs, or 0 if unknown.
	Line int32 `protobuf:"varint,2,opt,name=line,proto3" json:"line,omitempty"`
	// The 0-based index of the starting position within the line of source text
	// where the issue occurs.  Only meaningful if line is nonzer.
	Column int32 `protobuf:"varint,3,opt,name=column,proto3" json:"column,omitempty"`
	// Expression ID from [Expr][], 0 if unknown.
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusDetails) Reset()         { *m = StatusDetails{} }
func (m *StatusDetails) String() string { return proto.CompactTextString(m) }
func (*StatusDetails) ProtoMessage()    {}
func (*StatusDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{0}
}
func (m *StatusDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails.Unmarshal(m, b)
}
func (m *StatusDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails.Marshal(b, m, deterministic)
}
func (dst *StatusDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails.Merge(dst, src)
}
func (m *StatusDetails) XXX_Size() int {
	return xxx_messageInfo_StatusDetails.Size(m)
}
func (m *StatusDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails proto.InternalMessageInfo

func (m *StatusDetails) GetSeverity() StatusDetails_Severity {
	if m != nil {
		return m.Severity
	}
	return StatusDetails_SEVERITY_UNSPECIFIED
}

func (m *StatusDetails) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *StatusDetails) GetColumn() int32 {
	if m != nil {
		return m.Column
	}
	return 0
}

func (m *StatusDetails) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request message for the Parse method.
type ParseRequest struct {
	// Required. Source text in CEL syntax.
	CelSource string `protobuf:"bytes,1,opt,name=cel_source,json=celSource,proto3" json:"cel_source,omitempty"`
	// Tag for version of CEL syntax, for future use.
	SyntaxVersion string `protobuf:"bytes,2,opt,name=syntax_version,json=syntaxVersion,proto3" json:"syntax_version,omitempty"`
	// File or resource for source text, used in [SourceInfo][].
	SourceLocation string `protobuf:"bytes,3,opt,name=source_location,json=sourceLocation,proto3" json:"source_location,omitempty"`
	// Prevent macro expansion.  See "Macros" in Language Defiinition.
	DisableMacros        bool     `protobuf:"varint,4,opt,name=disable_macros,json=disableMacros,proto3" json:"disable_macros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseRequest) Reset()         { *m = ParseRequest{} }
func (m *ParseRequest) String() string { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()    {}
func (*ParseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{1}
}
func (m *ParseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseRequest.Unmarshal(m, b)
}
func (m *ParseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseRequest.Marshal(b, m, deterministic)
}
func (dst *ParseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseRequest.Merge(dst, src)
}
func (m *ParseRequest) XXX_Size() int {
	return xxx_messageInfo_ParseRequest.Size(m)
}
func (m *ParseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseRequest proto.InternalMessageInfo

func (m *ParseRequest) GetCelSource() string {
	if m != nil {
		return m.CelSource
	}
	return ""
}

func (m *ParseRequest) GetSyntaxVersion() string {
	if m != nil {
		return m.SyntaxVersion
	}
	return ""
}

func (m *ParseRequest) GetSourceLocation() string {
	if m != nil {
		return m.SourceLocation
	}
	return ""
}

func (m *ParseRequest) GetDisableMacros() bool {
	if m != nil {
		return m.DisableMacros
	}
	return false
}

// Response message for the Parse method.
type ParseResponse struct {
	// The parsed representation, or unset if parsing failed.
	ParsedExpr *ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
	// Any number of issues with [StatusDetails][] as the details.
	Issues               []*status.Status `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ParseResponse) Reset()         { *m = ParseResponse{} }
func (m *ParseResponse) String() string { return proto.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()    {}
func (*ParseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{2}
}
func (m *ParseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResponse.Unmarshal(m, b)
}
func (m *ParseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResponse.Marshal(b, m, deterministic)
}
func (dst *ParseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResponse.Merge(dst, src)
}
func (m *ParseResponse) XXX_Size() int {
	return xxx_messageInfo_ParseResponse.Size(m)
}
func (m *ParseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResponse proto.InternalMessageInfo

func (m *ParseResponse) GetParsedExpr() *ParsedExpr {
	if m != nil {
		return m.ParsedExpr
	}
	return nil
}

func (m *ParseResponse) GetIssues() []*status.Status {
	if m != nil {
		return m.Issues
	}
	return nil
}

// Request message for the Check method.
type CheckRequest struct {
	// Required. The parsed representation of the CEL program.
	ParsedExpr *ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3" json:"parsed_expr,omitempty"`
	// Declarations of types for external variables and functions.
	// Required if program uses external variables or functions
	// not in the default environment.
	TypeEnv []*Decl `protobuf:"bytes,2,rep,name=type_env,json=typeEnv,proto3" json:"type_env,omitempty"`
	// The protocol buffer context.  See "Name Resolution" in the
	// Language Definition.
	Container string `protobuf:"bytes,3,opt,name=container,proto3" json:"container,omitempty"`
	// If true, use only the declarations in [type_env][].  If false (default), add
	// declarations for the standard definitions to the type environment.  See
	// "Standard Definitions" in the Language Definition.
	NoStdEnv             bool     `protobuf:"varint,4,opt,name=no_std_env,json=noStdEnv,proto3" json:"no_std_env,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{3}
}
func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (dst *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(dst, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetParsedExpr() *ParsedExpr {
	if m != nil {
		return m.ParsedExpr
	}
	return nil
}

func (m *CheckRequest) GetTypeEnv() []*Decl {
	if m != nil {
		return m.TypeEnv
	}
	return nil
}

func (m *CheckRequest) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *CheckRequest) GetNoStdEnv() bool {
	if m != nil {
		return m.NoStdEnv
	}
	return false
}

// Response message for the Check method.
type CheckResponse struct {
	// The annotated representation, or unset if checking failed.
	CheckedExpr *CheckedExpr `protobuf:"bytes,1,opt,name=checked_expr,json=checkedExpr,proto3" json:"checked_expr,omitempty"`
	// Any number of issues with [StatusDetails][] as the details.
	Issues               []*status.Status `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{4}
}
func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (dst *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(dst, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetCheckedExpr() *CheckedExpr {
	if m != nil {
		return m.CheckedExpr
	}
	return nil
}

func (m *CheckResponse) GetIssues() []*status.Status {
	if m != nil {
		return m.Issues
	}
	return nil
}

// Request message for the Eval method.
type EvalRequest struct {
	// Required. Either the parsed or annotated representation of the CEL program.
	//
	// Types that are valid to be assigned to ExprKind:
	//	*EvalRequest_ParsedExpr
	//	*EvalRequest_CheckedExpr
	ExprKind isEvalRequest_ExprKind `protobuf_oneof:"expr_kind"`
	// Bindings for the external variables.  The types SHOULD be compatible
	// with the type environment in [CheckRequest][], if checked.
	Bindings map[string]*ExprValue `protobuf:"bytes,3,rep,name=bindings,proto3" json:"bindings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// SHOULD be the same container as used in [CheckRequest][], if checked.
	Container            string   `protobuf:"bytes,4,opt,name=container,proto3" json:"container,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvalRequest) Reset()         { *m = EvalRequest{} }
func (m *EvalRequest) String() string { return proto.CompactTextString(m) }
func (*EvalRequest) ProtoMessage()    {}
func (*EvalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{5}
}
func (m *EvalRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalRequest.Unmarshal(m, b)
}
func (m *EvalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalRequest.Marshal(b, m, deterministic)
}
func (dst *EvalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalRequest.Merge(dst, src)
}
func (m *EvalRequest) XXX_Size() int {
	return xxx_messageInfo_EvalRequest.Size(m)
}
func (m *EvalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvalRequest proto.InternalMessageInfo

type isEvalRequest_ExprKind interface {
	isEvalRequest_ExprKind()
}

type EvalRequest_ParsedExpr struct {
	ParsedExpr *ParsedExpr `protobuf:"bytes,1,opt,name=parsed_expr,json=parsedExpr,proto3,oneof"`
}

type EvalRequest_CheckedExpr struct {
	CheckedExpr *CheckedExpr `protobuf:"bytes,2,opt,name=checked_expr,json=checkedExpr,proto3,oneof"`
}

func (*EvalRequest_ParsedExpr) isEvalRequest_ExprKind() {}

func (*EvalRequest_CheckedExpr) isEvalRequest_ExprKind() {}

func (m *EvalRequest) GetExprKind() isEvalRequest_ExprKind {
	if m != nil {
		return m.ExprKind
	}
	return nil
}

func (m *EvalRequest) GetParsedExpr() *ParsedExpr {
	if x, ok := m.GetExprKind().(*EvalRequest_ParsedExpr); ok {
		return x.ParsedExpr
	}
	return nil
}

func (m *EvalRequest) GetCheckedExpr() *CheckedExpr {
	if x, ok := m.GetExprKind().(*EvalRequest_CheckedExpr); ok {
		return x.CheckedExpr
	}
	return nil
}

func (m *EvalRequest) GetBindings() map[string]*ExprValue {
	if m != nil {
		return m.Bindings
	}
	return nil
}

func (m *EvalRequest) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EvalRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EvalRequest_OneofMarshaler, _EvalRequest_OneofUnmarshaler, _EvalRequest_OneofSizer, []interface{}{
		(*EvalRequest_ParsedExpr)(nil),
		(*EvalRequest_CheckedExpr)(nil),
	}
}

func _EvalRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EvalRequest)
	// expr_kind
	switch x := m.ExprKind.(type) {
	case *EvalRequest_ParsedExpr:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ParsedExpr); err != nil {
			return err
		}
	case *EvalRequest_CheckedExpr:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CheckedExpr); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EvalRequest.ExprKind has unexpected type %T", x)
	}
	return nil
}

func _EvalRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EvalRequest)
	switch tag {
	case 1: // expr_kind.parsed_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ParsedExpr)
		err := b.DecodeMessage(msg)
		m.ExprKind = &EvalRequest_ParsedExpr{msg}
		return true, err
	case 2: // expr_kind.checked_expr
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CheckedExpr)
		err := b.DecodeMessage(msg)
		m.ExprKind = &EvalRequest_CheckedExpr{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EvalRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EvalRequest)
	// expr_kind
	switch x := m.ExprKind.(type) {
	case *EvalRequest_ParsedExpr:
		s := proto.Size(x.ParsedExpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EvalRequest_CheckedExpr:
		s := proto.Size(x.CheckedExpr)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Response message for the Eval method.
type EvalResponse struct {
	// The execution result, or unset if execution couldn't start.
	Result *ExprValue `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	// Any number of issues with [StatusDetails][] as the details.
	// Note that CEL execution errors are reified into [ExprValue][].
	// Nevertheless, we'll allow out-of-band issues to be raised,
	// which also makes the replies more regular.
	Issues               []*status.Status `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EvalResponse) Reset()         { *m = EvalResponse{} }
func (m *EvalResponse) String() string { return proto.CompactTextString(m) }
func (*EvalResponse) ProtoMessage()    {}
func (*EvalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cel_service_96c249772e35deed, []int{6}
}
func (m *EvalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvalResponse.Unmarshal(m, b)
}
func (m *EvalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvalResponse.Marshal(b, m, deterministic)
}
func (dst *EvalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvalResponse.Merge(dst, src)
}
func (m *EvalResponse) XXX_Size() int {
	return xxx_messageInfo_EvalResponse.Size(m)
}
func (m *EvalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvalResponse proto.InternalMessageInfo

func (m *EvalResponse) GetResult() *ExprValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *EvalResponse) GetIssues() []*status.Status {
	if m != nil {
		return m.Issues
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusDetails)(nil), "google.api.expr.v1.StatusDetails")
	proto.RegisterType((*ParseRequest)(nil), "google.api.expr.v1.ParseRequest")
	proto.RegisterType((*ParseResponse)(nil), "google.api.expr.v1.ParseResponse")
	proto.RegisterType((*CheckRequest)(nil), "google.api.expr.v1.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "google.api.expr.v1.CheckResponse")
	proto.RegisterType((*EvalRequest)(nil), "google.api.expr.v1.EvalRequest")
	proto.RegisterMapType((map[string]*ExprValue)(nil), "google.api.expr.v1.EvalRequest.BindingsEntry")
	proto.RegisterType((*EvalResponse)(nil), "google.api.expr.v1.EvalResponse")
	proto.RegisterEnum("google.api.expr.v1.StatusDetails_Severity", StatusDetails_Severity_name, StatusDetails_Severity_value)
}

func init() {
	proto.RegisterFile("proto/v1/cel-service.proto", fileDescriptor_cel_service_96c249772e35deed)
}

var fileDescriptor_cel_service_96c249772e35deed = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x6f, 0xe3, 0x44,
	0x14, 0xad, 0x9d, 0x8f, 0x4d, 0xae, 0x93, 0x34, 0x1a, 0xa0, 0x58, 0xd1, 0x2e, 0x1b, 0x82, 0x10,
	0xd1, 0x4a, 0xeb, 0x68, 0x53, 0x21, 0x21, 0x5e, 0x50, 0x93, 0x78, 0xd9, 0x48, 0xdd, 0x6c, 0x34,
	0x5e, 0x8a, 0xe8, 0x8b, 0xe5, 0x8c, 0x47, 0xa9, 0x55, 0x67, 0xec, 0x7a, 0x6c, 0x2b, 0x91, 0x90,
	0xf8, 0x33, 0xfc, 0x10, 0x7e, 0x0b, 0x4f, 0xbc, 0xf3, 0xc2, 0x23, 0xf2, 0x78, 0x92, 0x34, 0xc5,
	0xa5, 0x94, 0x7d, 0x9b, 0x39, 0x39, 0xf7, 0xf8, 0x9e, 0x73, 0x67, 0x32, 0xd0, 0x09, 0xa3, 0x20,
	0x0e, 0x06, 0xe9, 0xab, 0x01, 0xa1, 0xfe, 0x4b, 0x4e, 0xa3, 0xd4, 0x23, 0xd4, 0x10, 0x20, 0x42,
	0xcb, 0x20, 0x58, 0xfa, 0xd4, 0x70, 0x42, 0xcf, 0xa0, 0xeb, 0x30, 0x32, 0xd2, 0x57, 0x9d, 0x4f,
	0x73, 0x6c, 0x10, 0x85, 0x64, 0xc0, 0x63, 0x27, 0x4e, 0x78, 0x4e, 0xee, 0x9c, 0xec, 0x85, 0xae,
	0x28, 0xb9, 0xa6, 0xae, 0xc4, 0x3f, 0xda, 0xe1, 0x34, 0x75, 0x7c, 0x09, 0x7e, 0xb2, 0x03, 0xf9,
	0x86, 0xc5, 0xce, 0x3a, 0x87, 0x7b, 0xbf, 0x2b, 0xd0, 0xb4, 0x84, 0xe8, 0x84, 0xc6, 0x8e, 0xe7,
	0x73, 0xf4, 0x1a, 0x6a, 0x9c, 0xa6, 0x34, 0xf2, 0xe2, 0x8d, 0xae, 0x74, 0x95, 0x7e, 0x6b, 0xf8,
	0xc2, 0xf8, 0x67, 0x57, 0xc6, 0x41, 0x91, 0x61, 0xc9, 0x0a, 0xbc, 0xab, 0x45, 0x08, 0xca, 0xbe,
	0xc7, 0xa8, 0xae, 0x76, 0x95, 0x7e, 0x05, 0x8b, 0x35, 0x3a, 0x81, 0x2a, 0x09, 0xfc, 0x64, 0xc5,
	0xf4, 0x92, 0x40, 0xe5, 0x0e, 0xb5, 0x40, 0xf5, 0x5c, 0xbd, 0xdc, 0x55, 0xfa, 0x25, 0xac, 0x7a,
	0x6e, 0xef, 0x2d, 0xd4, 0xb6, 0x8a, 0x48, 0x87, 0x8f, 0x2d, 0xf3, 0xc2, 0xc4, 0xd3, 0xf7, 0x3f,
	0xd9, 0x3f, 0xcc, 0xac, 0xb9, 0x39, 0x9e, 0xbe, 0x9e, 0x9a, 0x93, 0xf6, 0x11, 0x3a, 0x06, 0x6d,
	0x62, 0xce, 0xb1, 0x39, 0x3e, 0x7b, 0x3f, 0x7d, 0x37, 0x6b, 0x2b, 0x48, 0x83, 0x27, 0x3f, 0x9e,
	0xe1, 0xd9, 0x74, 0xf6, 0x7d, 0x5b, 0x45, 0x75, 0xa8, 0x98, 0x18, 0xbf, 0xc3, 0xed, 0x52, 0xef,
	0x57, 0x05, 0x1a, 0x73, 0x27, 0xe2, 0x14, 0xd3, 0x9b, 0x84, 0xf2, 0x18, 0x3d, 0x03, 0x20, 0xd4,
	0xb7, 0x79, 0x90, 0x44, 0x84, 0x0a, 0x97, 0x75, 0x5c, 0x27, 0xd4, 0xb7, 0x04, 0x80, 0xbe, 0x84,
	0x56, 0x1e, 0x92, 0x9d, 0xd2, 0x88, 0x7b, 0x01, 0x13, 0x26, 0xea, 0xb8, 0x99, 0xa3, 0x17, 0x39,
	0x88, 0xbe, 0x82, 0xe3, 0x5c, 0xc1, 0xf6, 0x03, 0xe2, 0xc4, 0x19, 0xaf, 0x24, 0x78, 0xad, 0x1c,
	0x3e, 0x97, 0x68, 0xa6, 0xe7, 0x7a, 0xdc, 0x59, 0xf8, 0xd4, 0x5e, 0x39, 0x24, 0x0a, 0xb8, 0xb0,
	0x5a, 0xc3, 0x4d, 0x89, 0xbe, 0x15, 0x60, 0xef, 0x67, 0x68, 0xca, 0x2e, 0x79, 0x18, 0x30, 0x4e,
	0xd1, 0x77, 0xa0, 0x85, 0x19, 0xe0, 0xda, 0x59, 0xea, 0xa2, 0x4f, 0x6d, 0xf8, 0x59, 0xd1, 0x34,
	0x44, 0x9d, 0x6b, 0xae, 0xc3, 0x08, 0x43, 0xb8, 0x5b, 0xa3, 0x17, 0x50, 0xf5, 0x38, 0x4f, 0x28,
	0xd7, 0xd5, 0x6e, 0xa9, 0xaf, 0x0d, 0xd1, 0xb6, 0x36, 0x0a, 0x89, 0x9c, 0x20, 0x96, 0x8c, 0xde,
	0x6f, 0x0a, 0x34, 0xc6, 0xd9, 0x39, 0xda, 0x86, 0xf4, 0xc1, 0x5f, 0x3f, 0x85, 0x5a, 0xbc, 0x09,
	0xa9, 0x4d, 0x59, 0x2a, 0xbf, 0xaf, 0x17, 0x55, 0x4f, 0x28, 0xf1, 0xf1, 0x93, 0x8c, 0x69, 0xb2,
	0x14, 0x3d, 0x85, 0x3a, 0x09, 0x58, 0xec, 0x78, 0x8c, 0x46, 0x32, 0xce, 0x3d, 0x80, 0x9e, 0x02,
	0xb0, 0xc0, 0xe6, 0xb1, 0x2b, 0x44, 0xf3, 0x14, 0x6b, 0x2c, 0xb0, 0x62, 0xd7, 0x64, 0x69, 0xef,
	0x17, 0x68, 0x4a, 0x07, 0x32, 0xc0, 0x11, 0x34, 0xe4, 0xd5, 0xb8, 0xed, 0xe1, 0x79, 0x51, 0x17,
	0xe3, 0x9c, 0x27, 0x4c, 0x68, 0x64, 0xbf, 0x79, 0x54, 0x86, 0x7f, 0xa8, 0xa0, 0x99, 0xa9, 0xe3,
	0x6f, 0x23, 0x3c, 0xfb, 0x1f, 0x11, 0xbe, 0x39, 0x3a, 0x08, 0x71, 0x72, 0xc7, 0x82, 0xfa, 0x9f,
	0x2c, 0xbc, 0x39, 0x3a, 0x34, 0x31, 0x85, 0xda, 0xc2, 0x63, 0xae, 0xc7, 0x96, 0x5c, 0x2f, 0x09,
	0x1b, 0x2f, 0x8b, 0x14, 0x6e, 0xf5, 0x6e, 0x8c, 0x24, 0xdf, 0x64, 0x71, 0xb4, 0xc1, 0xbb, 0xf2,
	0xc3, 0x01, 0x95, 0xef, 0x0c, 0xa8, 0x73, 0x09, 0xcd, 0x83, 0x42, 0xd4, 0x86, 0xd2, 0x35, 0xdd,
	0xc8, 0x3b, 0x96, 0x2d, 0xd1, 0x29, 0x54, 0x52, 0xc7, 0x4f, 0xa8, 0xb4, 0xf2, 0xac, 0xb0, 0x91,
	0x75, 0x18, 0x5d, 0x64, 0x24, 0x9c, 0x73, 0xbf, 0x55, 0xbf, 0x51, 0x46, 0x1a, 0xd4, 0xb3, 0xdf,
	0xed, 0x6b, 0x8f, 0xb9, 0xbd, 0x1b, 0x68, 0xe4, 0xdd, 0xca, 0x51, 0x7f, 0x0d, 0xd5, 0x88, 0xf2,
	0xc4, 0x8f, 0x65, 0xca, 0x0f, 0xc8, 0x4a, 0xf2, 0x63, 0xa6, 0x3b, 0xfc, 0x53, 0x01, 0x18, 0x53,
	0xdf, 0xca, 0xff, 0xb1, 0xd1, 0x39, 0x54, 0xc4, 0xd4, 0x50, 0xf7, 0xde, 0x81, 0xca, 0x2c, 0x3b,
	0x9f, 0xff, 0x0b, 0x43, 0xf6, 0x7f, 0x0e, 0x15, 0x31, 0xbf, 0x62, 0xb5, 0xdb, 0x17, 0xb3, 0x58,
	0xed, 0xf0, 0xe0, 0x4f, 0xa1, 0x9c, 0xa5, 0x83, 0x9e, 0x3f, 0x30, 0xe5, 0x4e, 0xf7, 0x7e, 0x42,
	0x2e, 0x35, 0xb2, 0xe0, 0x84, 0x04, 0xab, 0x02, 0xda, 0xe8, 0x78, 0x1f, 0xc6, 0x3c, 0x7b, 0x4c,
	0xe6, 0xca, 0xe5, 0x17, 0x4b, 0x2f, 0xbe, 0x4a, 0x16, 0x06, 0x09, 0x56, 0x03, 0xf9, 0x6c, 0x89,
	0x47, 0x2e, 0xa4, 0x64, 0xb0, 0x7d, 0x80, 0xfe, 0x52, 0x94, 0x45, 0x55, 0x6c, 0x4e, 0xff, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xbc, 0x72, 0x9a, 0x70, 0x0c, 0x07, 0x00, 0x00,
}
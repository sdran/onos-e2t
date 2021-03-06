// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/e2ap/v1beta1/e2ap_commondatatypes.proto

package e2ap_commondatatypes

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// enumerated from e2ap-v01.00.00.asn1:1202
type Criticality int32

const (
	Criticality_CRITICALITY_REJECT Criticality = 0
	Criticality_CRITICALITY_IGNORE Criticality = 1
	Criticality_CRITICALITY_NOTIFY Criticality = 2
)

var Criticality_name = map[int32]string{
	0: "CRITICALITY_REJECT",
	1: "CRITICALITY_IGNORE",
	2: "CRITICALITY_NOTIFY",
}

var Criticality_value = map[string]int32{
	"CRITICALITY_REJECT": 0,
	"CRITICALITY_IGNORE": 1,
	"CRITICALITY_NOTIFY": 2,
}

func (x Criticality) String() string {
	return proto.EnumName(Criticality_name, int32(x))
}

func (Criticality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{0}
}

// enumerated from e2ap-v01.00.00.asn:1294
type Presence int32

const (
	Presence_PRESENCE_OPTIONAL    Presence = 0
	Presence_PRESENCE_CONDITIONAL Presence = 1
	Presence_PRESENCE_MANDATORY   Presence = 2
)

var Presence_name = map[int32]string{
	0: "PRESENCE_OPTIONAL",
	1: "PRESENCE_CONDITIONAL",
	2: "PRESENCE_MANDATORY",
}

var Presence_value = map[string]int32{
	"PRESENCE_OPTIONAL":    0,
	"PRESENCE_CONDITIONAL": 1,
	"PRESENCE_MANDATORY":   2,
}

func (x Presence) String() string {
	return proto.EnumName(Presence_name, int32(x))
}

func (Presence) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{1}
}

// enumerated from e2ap-v01.00.00.asn:1300
type TriggeringMessage int32

const (
	TriggeringMessage_TRIGGERING_MESSAGE_INITIATING_MESSAGE    TriggeringMessage = 0
	TriggeringMessage_TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME    TriggeringMessage = 1
	TriggeringMessage_TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME TriggeringMessage = 2
)

var TriggeringMessage_name = map[int32]string{
	0: "TRIGGERING_MESSAGE_INITIATING_MESSAGE",
	1: "TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME",
	2: "TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME",
}

var TriggeringMessage_value = map[string]int32{
	"TRIGGERING_MESSAGE_INITIATING_MESSAGE":    0,
	"TRIGGERING_MESSAGE_SUCCESSFUL_OUTCOME":    1,
	"TRIGGERING_MESSAGE_UNSUCCESSFULL_OUTCOME": 2,
}

func (x TriggeringMessage) String() string {
	return proto.EnumName(TriggeringMessage_name, int32(x))
}

func (TriggeringMessage) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{2}
}

type CriticalityReject struct {
	Criticality          Criticality `protobuf:"varint,1,opt,name=criticality,proto3,enum=e2ap.v1beta1.Criticality" json:"criticality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CriticalityReject) Reset()         { *m = CriticalityReject{} }
func (m *CriticalityReject) String() string { return proto.CompactTextString(m) }
func (*CriticalityReject) ProtoMessage()    {}
func (*CriticalityReject) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{0}
}
func (m *CriticalityReject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalityReject.Unmarshal(m, b)
}
func (m *CriticalityReject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalityReject.Marshal(b, m, deterministic)
}
func (m *CriticalityReject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalityReject.Merge(m, src)
}
func (m *CriticalityReject) XXX_Size() int {
	return xxx_messageInfo_CriticalityReject.Size(m)
}
func (m *CriticalityReject) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalityReject.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalityReject proto.InternalMessageInfo

func (m *CriticalityReject) GetCriticality() Criticality {
	if m != nil {
		return m.Criticality
	}
	return Criticality_CRITICALITY_REJECT
}

type CriticalityIgnore struct {
	Criticality          Criticality `protobuf:"varint,1,opt,name=criticality,proto3,enum=e2ap.v1beta1.Criticality" json:"criticality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CriticalityIgnore) Reset()         { *m = CriticalityIgnore{} }
func (m *CriticalityIgnore) String() string { return proto.CompactTextString(m) }
func (*CriticalityIgnore) ProtoMessage()    {}
func (*CriticalityIgnore) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{1}
}
func (m *CriticalityIgnore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalityIgnore.Unmarshal(m, b)
}
func (m *CriticalityIgnore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalityIgnore.Marshal(b, m, deterministic)
}
func (m *CriticalityIgnore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalityIgnore.Merge(m, src)
}
func (m *CriticalityIgnore) XXX_Size() int {
	return xxx_messageInfo_CriticalityIgnore.Size(m)
}
func (m *CriticalityIgnore) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalityIgnore.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalityIgnore proto.InternalMessageInfo

func (m *CriticalityIgnore) GetCriticality() Criticality {
	if m != nil {
		return m.Criticality
	}
	return Criticality_CRITICALITY_REJECT
}

type CriticalityNotify struct {
	Criticality          Criticality `protobuf:"varint,1,opt,name=criticality,proto3,enum=e2ap.v1beta1.Criticality" json:"criticality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CriticalityNotify) Reset()         { *m = CriticalityNotify{} }
func (m *CriticalityNotify) String() string { return proto.CompactTextString(m) }
func (*CriticalityNotify) ProtoMessage()    {}
func (*CriticalityNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{2}
}
func (m *CriticalityNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalityNotify.Unmarshal(m, b)
}
func (m *CriticalityNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalityNotify.Marshal(b, m, deterministic)
}
func (m *CriticalityNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalityNotify.Merge(m, src)
}
func (m *CriticalityNotify) XXX_Size() int {
	return xxx_messageInfo_CriticalityNotify.Size(m)
}
func (m *CriticalityNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalityNotify.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalityNotify proto.InternalMessageInfo

func (m *CriticalityNotify) GetCriticality() Criticality {
	if m != nil {
		return m.Criticality
	}
	return Criticality_CRITICALITY_REJECT
}

type PresenceOptional struct {
	Presence             Presence `protobuf:"varint,1,opt,name=presence,proto3,enum=e2ap.v1beta1.Presence" json:"presence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PresenceOptional) Reset()         { *m = PresenceOptional{} }
func (m *PresenceOptional) String() string { return proto.CompactTextString(m) }
func (*PresenceOptional) ProtoMessage()    {}
func (*PresenceOptional) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{3}
}
func (m *PresenceOptional) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceOptional.Unmarshal(m, b)
}
func (m *PresenceOptional) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceOptional.Marshal(b, m, deterministic)
}
func (m *PresenceOptional) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceOptional.Merge(m, src)
}
func (m *PresenceOptional) XXX_Size() int {
	return xxx_messageInfo_PresenceOptional.Size(m)
}
func (m *PresenceOptional) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceOptional.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceOptional proto.InternalMessageInfo

func (m *PresenceOptional) GetPresence() Presence {
	if m != nil {
		return m.Presence
	}
	return Presence_PRESENCE_OPTIONAL
}

type PresenceConditional struct {
	Presence             Presence `protobuf:"varint,1,opt,name=presence,proto3,enum=e2ap.v1beta1.Presence" json:"presence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PresenceConditional) Reset()         { *m = PresenceConditional{} }
func (m *PresenceConditional) String() string { return proto.CompactTextString(m) }
func (*PresenceConditional) ProtoMessage()    {}
func (*PresenceConditional) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{4}
}
func (m *PresenceConditional) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceConditional.Unmarshal(m, b)
}
func (m *PresenceConditional) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceConditional.Marshal(b, m, deterministic)
}
func (m *PresenceConditional) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceConditional.Merge(m, src)
}
func (m *PresenceConditional) XXX_Size() int {
	return xxx_messageInfo_PresenceConditional.Size(m)
}
func (m *PresenceConditional) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceConditional.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceConditional proto.InternalMessageInfo

func (m *PresenceConditional) GetPresence() Presence {
	if m != nil {
		return m.Presence
	}
	return Presence_PRESENCE_OPTIONAL
}

type PresenceMandatory struct {
	Presence             Presence `protobuf:"varint,1,opt,name=presence,proto3,enum=e2ap.v1beta1.Presence" json:"presence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PresenceMandatory) Reset()         { *m = PresenceMandatory{} }
func (m *PresenceMandatory) String() string { return proto.CompactTextString(m) }
func (*PresenceMandatory) ProtoMessage()    {}
func (*PresenceMandatory) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{5}
}
func (m *PresenceMandatory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PresenceMandatory.Unmarshal(m, b)
}
func (m *PresenceMandatory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PresenceMandatory.Marshal(b, m, deterministic)
}
func (m *PresenceMandatory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PresenceMandatory.Merge(m, src)
}
func (m *PresenceMandatory) XXX_Size() int {
	return xxx_messageInfo_PresenceMandatory.Size(m)
}
func (m *PresenceMandatory) XXX_DiscardUnknown() {
	xxx_messageInfo_PresenceMandatory.DiscardUnknown(m)
}

var xxx_messageInfo_PresenceMandatory proto.InternalMessageInfo

func (m *PresenceMandatory) GetPresence() Presence {
	if m != nil {
		return m.Presence
	}
	return Presence_PRESENCE_OPTIONAL
}

// range of Integer from e2ap-v01.00.00.asn1:1206
type ProcedureCode struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcedureCode) Reset()         { *m = ProcedureCode{} }
func (m *ProcedureCode) String() string { return proto.CompactTextString(m) }
func (*ProcedureCode) ProtoMessage()    {}
func (*ProcedureCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{6}
}
func (m *ProcedureCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcedureCode.Unmarshal(m, b)
}
func (m *ProcedureCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcedureCode.Marshal(b, m, deterministic)
}
func (m *ProcedureCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcedureCode.Merge(m, src)
}
func (m *ProcedureCode) XXX_Size() int {
	return xxx_messageInfo_ProcedureCode.Size(m)
}
func (m *ProcedureCode) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcedureCode.DiscardUnknown(m)
}

var xxx_messageInfo_ProcedureCode proto.InternalMessageInfo

func (m *ProcedureCode) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// range of Integer from e2ap-v01.00.00.asn1:1208
type ProtocolIeId struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtocolIeId) Reset()         { *m = ProtocolIeId{} }
func (m *ProtocolIeId) String() string { return proto.CompactTextString(m) }
func (*ProtocolIeId) ProtoMessage()    {}
func (*ProtocolIeId) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{7}
}
func (m *ProtocolIeId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolIeId.Unmarshal(m, b)
}
func (m *ProtocolIeId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolIeId.Marshal(b, m, deterministic)
}
func (m *ProtocolIeId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolIeId.Merge(m, src)
}
func (m *ProtocolIeId) XXX_Size() int {
	return xxx_messageInfo_ProtocolIeId.Size(m)
}
func (m *ProtocolIeId) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolIeId.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolIeId proto.InternalMessageInfo

func (m *ProtocolIeId) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type RiccallProcessId struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiccallProcessId) Reset()         { *m = RiccallProcessId{} }
func (m *RiccallProcessId) String() string { return proto.CompactTextString(m) }
func (*RiccallProcessId) ProtoMessage()    {}
func (*RiccallProcessId) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{8}
}
func (m *RiccallProcessId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiccallProcessId.Unmarshal(m, b)
}
func (m *RiccallProcessId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiccallProcessId.Marshal(b, m, deterministic)
}
func (m *RiccallProcessId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiccallProcessId.Merge(m, src)
}
func (m *RiccallProcessId) XXX_Size() int {
	return xxx_messageInfo_RiccallProcessId.Size(m)
}
func (m *RiccallProcessId) XXX_DiscardUnknown() {
	xxx_messageInfo_RiccallProcessId.DiscardUnknown(m)
}

var xxx_messageInfo_RiccallProcessId proto.InternalMessageInfo

func (m *RiccallProcessId) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RiccontrolHeader struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiccontrolHeader) Reset()         { *m = RiccontrolHeader{} }
func (m *RiccontrolHeader) String() string { return proto.CompactTextString(m) }
func (*RiccontrolHeader) ProtoMessage()    {}
func (*RiccontrolHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{9}
}
func (m *RiccontrolHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiccontrolHeader.Unmarshal(m, b)
}
func (m *RiccontrolHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiccontrolHeader.Marshal(b, m, deterministic)
}
func (m *RiccontrolHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiccontrolHeader.Merge(m, src)
}
func (m *RiccontrolHeader) XXX_Size() int {
	return xxx_messageInfo_RiccontrolHeader.Size(m)
}
func (m *RiccontrolHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RiccontrolHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RiccontrolHeader proto.InternalMessageInfo

func (m *RiccontrolHeader) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RiccontrolMessage struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiccontrolMessage) Reset()         { *m = RiccontrolMessage{} }
func (m *RiccontrolMessage) String() string { return proto.CompactTextString(m) }
func (*RiccontrolMessage) ProtoMessage()    {}
func (*RiccontrolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{10}
}
func (m *RiccontrolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiccontrolMessage.Unmarshal(m, b)
}
func (m *RiccontrolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiccontrolMessage.Marshal(b, m, deterministic)
}
func (m *RiccontrolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiccontrolMessage.Merge(m, src)
}
func (m *RiccontrolMessage) XXX_Size() int {
	return xxx_messageInfo_RiccontrolMessage.Size(m)
}
func (m *RiccontrolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RiccontrolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RiccontrolMessage proto.InternalMessageInfo

func (m *RiccontrolMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RiccontrolOutcome struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiccontrolOutcome) Reset()         { *m = RiccontrolOutcome{} }
func (m *RiccontrolOutcome) String() string { return proto.CompactTextString(m) }
func (*RiccontrolOutcome) ProtoMessage()    {}
func (*RiccontrolOutcome) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{11}
}
func (m *RiccontrolOutcome) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiccontrolOutcome.Unmarshal(m, b)
}
func (m *RiccontrolOutcome) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiccontrolOutcome.Marshal(b, m, deterministic)
}
func (m *RiccontrolOutcome) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiccontrolOutcome.Merge(m, src)
}
func (m *RiccontrolOutcome) XXX_Size() int {
	return xxx_messageInfo_RiccontrolOutcome.Size(m)
}
func (m *RiccontrolOutcome) XXX_DiscardUnknown() {
	xxx_messageInfo_RiccontrolOutcome.DiscardUnknown(m)
}

var xxx_messageInfo_RiccontrolOutcome proto.InternalMessageInfo

func (m *RiccontrolOutcome) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RicindicationHeader struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RicindicationHeader) Reset()         { *m = RicindicationHeader{} }
func (m *RicindicationHeader) String() string { return proto.CompactTextString(m) }
func (*RicindicationHeader) ProtoMessage()    {}
func (*RicindicationHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{12}
}
func (m *RicindicationHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RicindicationHeader.Unmarshal(m, b)
}
func (m *RicindicationHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RicindicationHeader.Marshal(b, m, deterministic)
}
func (m *RicindicationHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RicindicationHeader.Merge(m, src)
}
func (m *RicindicationHeader) XXX_Size() int {
	return xxx_messageInfo_RicindicationHeader.Size(m)
}
func (m *RicindicationHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_RicindicationHeader.DiscardUnknown(m)
}

var xxx_messageInfo_RicindicationHeader proto.InternalMessageInfo

func (m *RicindicationHeader) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RicindicationMessage struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RicindicationMessage) Reset()         { *m = RicindicationMessage{} }
func (m *RicindicationMessage) String() string { return proto.CompactTextString(m) }
func (*RicindicationMessage) ProtoMessage()    {}
func (*RicindicationMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{13}
}
func (m *RicindicationMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RicindicationMessage.Unmarshal(m, b)
}
func (m *RicindicationMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RicindicationMessage.Marshal(b, m, deterministic)
}
func (m *RicindicationMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RicindicationMessage.Merge(m, src)
}
func (m *RicindicationMessage) XXX_Size() int {
	return xxx_messageInfo_RicindicationMessage.Size(m)
}
func (m *RicindicationMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RicindicationMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RicindicationMessage proto.InternalMessageInfo

func (m *RicindicationMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RanfunctionDefinition struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RanfunctionDefinition) Reset()         { *m = RanfunctionDefinition{} }
func (m *RanfunctionDefinition) String() string { return proto.CompactTextString(m) }
func (*RanfunctionDefinition) ProtoMessage()    {}
func (*RanfunctionDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{14}
}
func (m *RanfunctionDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RanfunctionDefinition.Unmarshal(m, b)
}
func (m *RanfunctionDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RanfunctionDefinition.Marshal(b, m, deterministic)
}
func (m *RanfunctionDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RanfunctionDefinition.Merge(m, src)
}
func (m *RanfunctionDefinition) XXX_Size() int {
	return xxx_messageInfo_RanfunctionDefinition.Size(m)
}
func (m *RanfunctionDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_RanfunctionDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_RanfunctionDefinition proto.InternalMessageInfo

func (m *RanfunctionDefinition) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PlmnIdentity struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlmnIdentity) Reset()         { *m = PlmnIdentity{} }
func (m *PlmnIdentity) String() string { return proto.CompactTextString(m) }
func (*PlmnIdentity) ProtoMessage()    {}
func (*PlmnIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{15}
}
func (m *PlmnIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmnIdentity.Unmarshal(m, b)
}
func (m *PlmnIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmnIdentity.Marshal(b, m, deterministic)
}
func (m *PlmnIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmnIdentity.Merge(m, src)
}
func (m *PlmnIdentity) XXX_Size() int {
	return xxx_messageInfo_PlmnIdentity.Size(m)
}
func (m *PlmnIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmnIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_PlmnIdentity proto.InternalMessageInfo

func (m *PlmnIdentity) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RiceventTriggerDefinition struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RiceventTriggerDefinition) Reset()         { *m = RiceventTriggerDefinition{} }
func (m *RiceventTriggerDefinition) String() string { return proto.CompactTextString(m) }
func (*RiceventTriggerDefinition) ProtoMessage()    {}
func (*RiceventTriggerDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{16}
}
func (m *RiceventTriggerDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RiceventTriggerDefinition.Unmarshal(m, b)
}
func (m *RiceventTriggerDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RiceventTriggerDefinition.Marshal(b, m, deterministic)
}
func (m *RiceventTriggerDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiceventTriggerDefinition.Merge(m, src)
}
func (m *RiceventTriggerDefinition) XXX_Size() int {
	return xxx_messageInfo_RiceventTriggerDefinition.Size(m)
}
func (m *RiceventTriggerDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_RiceventTriggerDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_RiceventTriggerDefinition proto.InternalMessageInfo

func (m *RiceventTriggerDefinition) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RicactionDefinition struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RicactionDefinition) Reset()         { *m = RicactionDefinition{} }
func (m *RicactionDefinition) String() string { return proto.CompactTextString(m) }
func (*RicactionDefinition) ProtoMessage()    {}
func (*RicactionDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{17}
}
func (m *RicactionDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RicactionDefinition.Unmarshal(m, b)
}
func (m *RicactionDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RicactionDefinition.Marshal(b, m, deterministic)
}
func (m *RicactionDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RicactionDefinition.Merge(m, src)
}
func (m *RicactionDefinition) XXX_Size() int {
	return xxx_messageInfo_RicactionDefinition.Size(m)
}
func (m *RicactionDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_RicactionDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_RicactionDefinition proto.InternalMessageInfo

func (m *RicactionDefinition) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BitString struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Len                  uint32   `protobuf:"varint,2,opt,name=len,proto3" json:"len,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BitString) Reset()         { *m = BitString{} }
func (m *BitString) String() string { return proto.CompactTextString(m) }
func (*BitString) ProtoMessage()    {}
func (*BitString) Descriptor() ([]byte, []int) {
	return fileDescriptor_863e491415873d9a, []int{18}
}
func (m *BitString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BitString.Unmarshal(m, b)
}
func (m *BitString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BitString.Marshal(b, m, deterministic)
}
func (m *BitString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitString.Merge(m, src)
}
func (m *BitString) XXX_Size() int {
	return xxx_messageInfo_BitString.Size(m)
}
func (m *BitString) XXX_DiscardUnknown() {
	xxx_messageInfo_BitString.DiscardUnknown(m)
}

var xxx_messageInfo_BitString proto.InternalMessageInfo

func (m *BitString) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *BitString) GetLen() uint32 {
	if m != nil {
		return m.Len
	}
	return 0
}

func init() {
	proto.RegisterEnum("e2ap.v1beta1.Criticality", Criticality_name, Criticality_value)
	proto.RegisterEnum("e2ap.v1beta1.Presence", Presence_name, Presence_value)
	proto.RegisterEnum("e2ap.v1beta1.TriggeringMessage", TriggeringMessage_name, TriggeringMessage_value)
	proto.RegisterType((*CriticalityReject)(nil), "e2ap.v1beta1.CriticalityReject")
	proto.RegisterType((*CriticalityIgnore)(nil), "e2ap.v1beta1.CriticalityIgnore")
	proto.RegisterType((*CriticalityNotify)(nil), "e2ap.v1beta1.CriticalityNotify")
	proto.RegisterType((*PresenceOptional)(nil), "e2ap.v1beta1.PresenceOptional")
	proto.RegisterType((*PresenceConditional)(nil), "e2ap.v1beta1.PresenceConditional")
	proto.RegisterType((*PresenceMandatory)(nil), "e2ap.v1beta1.PresenceMandatory")
	proto.RegisterType((*ProcedureCode)(nil), "e2ap.v1beta1.ProcedureCode")
	proto.RegisterType((*ProtocolIeId)(nil), "e2ap.v1beta1.ProtocolIeId")
	proto.RegisterType((*RiccallProcessId)(nil), "e2ap.v1beta1.RiccallProcessId")
	proto.RegisterType((*RiccontrolHeader)(nil), "e2ap.v1beta1.RiccontrolHeader")
	proto.RegisterType((*RiccontrolMessage)(nil), "e2ap.v1beta1.RiccontrolMessage")
	proto.RegisterType((*RiccontrolOutcome)(nil), "e2ap.v1beta1.RiccontrolOutcome")
	proto.RegisterType((*RicindicationHeader)(nil), "e2ap.v1beta1.RicindicationHeader")
	proto.RegisterType((*RicindicationMessage)(nil), "e2ap.v1beta1.RicindicationMessage")
	proto.RegisterType((*RanfunctionDefinition)(nil), "e2ap.v1beta1.RanfunctionDefinition")
	proto.RegisterType((*PlmnIdentity)(nil), "e2ap.v1beta1.PlmnIdentity")
	proto.RegisterType((*RiceventTriggerDefinition)(nil), "e2ap.v1beta1.RiceventTriggerDefinition")
	proto.RegisterType((*RicactionDefinition)(nil), "e2ap.v1beta1.RicactionDefinition")
	proto.RegisterType((*BitString)(nil), "e2ap.v1beta1.BitString")
}

func init() {
	proto.RegisterFile("api/e2ap/v1beta1/e2ap_commondatatypes.proto", fileDescriptor_863e491415873d9a)
}

var fileDescriptor_863e491415873d9a = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcd, 0x6e, 0xe2, 0x3a,
	0x14, 0x80, 0x49, 0x5a, 0x5a, 0xae, 0x4b, 0xaf, 0x20, 0xfd, 0x11, 0x45, 0xba, 0x52, 0x6f, 0xa5,
	0x91, 0xe8, 0x1f, 0x88, 0xce, 0xb6, 0x8b, 0x21, 0x69, 0xca, 0x78, 0x54, 0x12, 0xc6, 0x09, 0x8b,
	0x76, 0x83, 0x5c, 0xc7, 0xa5, 0x1e, 0x05, 0x3b, 0x4a, 0x0c, 0x12, 0xb3, 0x9c, 0x97, 0x98, 0x77,
	0x65, 0xd3, 0x51, 0x02, 0x94, 0x80, 0x4a, 0xa7, 0x62, 0x17, 0xce, 0xf9, 0xce, 0x97, 0x63, 0xfb,
	0x10, 0x83, 0x73, 0x1c, 0xb0, 0x1a, 0xbd, 0xc2, 0x41, 0x6d, 0x58, 0x7f, 0xa4, 0x12, 0xd7, 0x93,
	0x1f, 0x5d, 0x22, 0xfa, 0x7d, 0xc1, 0x3d, 0x2c, 0xb1, 0x1c, 0x05, 0x34, 0xaa, 0x06, 0xa1, 0x90,
	0x42, 0xcb, 0xc7, 0xb9, 0xea, 0x14, 0x2c, 0x97, 0x87, 0xd8, 0x67, 0x1e, 0x96, 0xb4, 0x36, 0xac,
	0xd7, 0x66, 0xcf, 0x13, 0xf2, 0xe4, 0x01, 0x14, 0x8d, 0x90, 0x49, 0x46, 0xb0, 0xcf, 0xe4, 0x08,
	0xd1, 0x1f, 0x94, 0x48, 0xcd, 0x04, 0x3b, 0x64, 0x1e, 0x2c, 0x29, 0xc7, 0x4a, 0xe5, 0xdf, 0xab,
	0xa3, 0x6a, 0x5a, 0x5a, 0x4d, 0x55, 0xe9, 0xb9, 0xb1, 0x9e, 0xfd, 0xa5, 0xa8, 0xb9, 0x0c, 0x4a,
	0xd7, 0x2d, 0xb9, 0x61, 0x8f, 0x8b, 0x90, 0xae, 0xed, 0x56, 0xde, 0x73, 0x5b, 0x42, 0xb2, 0xa7,
	0xd1, 0xda, 0x6e, 0x75, 0xd1, 0xdd, 0x06, 0x85, 0x76, 0x48, 0x23, 0xca, 0x09, 0xb5, 0x03, 0xc9,
	0x04, 0xc7, 0xbe, 0x76, 0x0d, 0x72, 0xc1, 0x34, 0x36, 0xf5, 0x1e, 0x2e, 0x7a, 0x67, 0x15, 0xa9,
	0xcd, 0x78, 0xad, 0x38, 0x71, 0xc0, 0xde, 0x2c, 0x6f, 0x08, 0xee, 0xb1, 0xb5, 0xa5, 0x4a, 0x4a,
	0xfa, 0x1d, 0x14, 0x67, 0xf9, 0x16, 0x8e, 0x47, 0x40, 0x84, 0xa3, 0x35, 0x94, 0x6a, 0x4a, 0x59,
	0x07, 0xbb, 0xed, 0x50, 0x10, 0xea, 0x0d, 0x42, 0x6a, 0x08, 0x8f, 0x6a, 0xc7, 0x20, 0x3b, 0xc4,
	0xfe, 0x60, 0xe2, 0xca, 0xea, 0x60, 0xac, 0x6f, 0x97, 0xb3, 0xa5, 0x17, 0xa5, 0x92, 0x41, 0x93,
	0xc4, 0x49, 0x1d, 0xe4, 0xdb, 0xf1, 0x24, 0x11, 0xe1, 0x43, 0x0a, 0x3d, 0xed, 0xff, 0xc5, 0x8a,
	0x9d, 0xb1, 0x9e, 0x2b, 0x6f, 0x95, 0x5e, 0x5e, 0x36, 0xe6, 0x25, 0x15, 0x50, 0x40, 0x8c, 0x10,
	0xec, 0xfb, 0xc9, 0xcb, 0xa2, 0x08, 0x7a, 0xda, 0x7e, 0xba, 0x2c, 0xbf, 0x44, 0x0a, 0x2e, 0x43,
	0xe1, 0x7f, 0xa5, 0xd8, 0xa3, 0xe1, 0x0a, 0xf2, 0x14, 0x14, 0xe7, 0x64, 0x8b, 0x46, 0x11, 0xee,
	0xd1, 0x8f, 0xa0, 0xf6, 0x40, 0x12, 0xd1, 0x5f, 0x85, 0x9e, 0x83, 0x3d, 0xc4, 0x08, 0xe3, 0x1e,
	0x23, 0x38, 0x3e, 0xb3, 0x77, 0x5b, 0xb8, 0x00, 0xfb, 0x0b, 0xf0, 0xfb, 0x5d, 0x5c, 0x82, 0x03,
	0x84, 0xf9, 0xd3, 0x80, 0x93, 0x98, 0xbd, 0xa1, 0x4f, 0x8c, 0x27, 0x63, 0xb1, 0x12, 0xcf, 0xb7,
	0xfd, 0x3e, 0x87, 0x1e, 0xe5, 0x92, 0xc9, 0x91, 0xf6, 0xdf, 0x02, 0xa5, 0x6f, 0x8f, 0xf5, 0xcd,
	0x9f, 0xea, 0xf3, 0xc6, 0xfc, 0x54, 0x8e, 0x10, 0x23, 0x74, 0x48, 0xb9, 0x74, 0x43, 0xd6, 0xeb,
	0xd1, 0xf0, 0xaf, 0x6f, 0x98, 0xac, 0x15, 0x7f, 0xac, 0x9d, 0x6b, 0xf0, 0x8f, 0xce, 0xa4, 0x23,
	0x43, 0xc6, 0x7b, 0x8b, 0xc8, 0xe6, 0x14, 0xd1, 0x8e, 0xc0, 0x86, 0x4f, 0x79, 0x49, 0x3d, 0x56,
	0x2a, 0xbb, 0x49, 0x7f, 0x67, 0x6a, 0xe9, 0x0b, 0x8a, 0x63, 0x67, 0x1d, 0xb0, 0x93, 0xfa, 0x1b,
	0x6a, 0x87, 0x40, 0x33, 0x10, 0x74, 0xa1, 0xd1, 0xb8, 0x83, 0xee, 0x7d, 0x17, 0x99, 0xdf, 0x4c,
	0xc3, 0x2d, 0x64, 0x96, 0xe3, 0xb0, 0x69, 0xd9, 0xc8, 0x2c, 0x28, 0xcb, 0x71, 0xcb, 0x76, 0xe1,
	0xed, 0x7d, 0x41, 0x3d, 0x73, 0x40, 0x6e, 0x36, 0xdd, 0xda, 0x01, 0x28, 0xb6, 0x91, 0xe9, 0x98,
	0x96, 0x61, 0x76, 0xed, 0xb6, 0x0b, 0x6d, 0xab, 0x71, 0x57, 0xc8, 0x68, 0x25, 0xb0, 0xff, 0x1a,
	0x36, 0x6c, 0xeb, 0x06, 0x4e, 0x33, 0x89, 0xf4, 0x35, 0xd3, 0x6a, 0x58, 0x37, 0x0d, 0xd7, 0x46,
	0xb1, 0xf4, 0xb7, 0x02, 0x8a, 0xd3, 0x2d, 0x64, 0xbc, 0x37, 0x3b, 0xd3, 0x53, 0xf0, 0xc9, 0x45,
	0xb0, 0xd9, 0x34, 0x11, 0xb4, 0x9a, 0xdd, 0x96, 0xe9, 0x38, 0x8d, 0xa6, 0xd9, 0x85, 0x16, 0x74,
	0x61, 0xc3, 0x4d, 0x85, 0x0a, 0x99, 0x15, 0xa8, 0xd3, 0x31, 0x0c, 0xd3, 0x71, 0x6e, 0x3b, 0x77,
	0x5d, 0xbb, 0xe3, 0x1a, 0x76, 0x2b, 0x5e, 0xd8, 0x05, 0xa8, 0xbc, 0x81, 0x76, 0xac, 0x39, 0x3c,
	0xa7, 0x55, 0xbd, 0xf9, 0x60, 0xf6, 0x98, 0x7c, 0x1e, 0x3c, 0x56, 0x89, 0xe8, 0xd7, 0x04, 0x17,
	0x51, 0x10, 0x8a, 0xf8, 0xfb, 0x9d, 0x3c, 0x5f, 0xd2, 0x2b, 0x59, 0x7b, 0xf3, 0xce, 0xb8, 0x5c,
	0xba, 0x33, 0x1e, 0xb7, 0x92, 0xab, 0xe0, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0x4a,
	0x89, 0x61, 0x63, 0x06, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/eventexecution.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type EventExecution_Status int32

const (
	EventExecution_IN_PROGRESS EventExecution_Status = 0
	EventExecution_COMPLETED   EventExecution_Status = 1
	EventExecution_FAILED      EventExecution_Status = 2
	EventExecution_SKIPPED     EventExecution_Status = 3
)

var EventExecution_Status_name = map[int32]string{
	0: "IN_PROGRESS",
	1: "COMPLETED",
	2: "FAILED",
	3: "SKIPPED",
}
var EventExecution_Status_value = map[string]int32{
	"IN_PROGRESS": 0,
	"COMPLETED":   1,
	"FAILED":      2,
	"SKIPPED":     3,
}

func (x EventExecution_Status) String() string {
	return proto.EnumName(EventExecution_Status_name, int32(x))
}
func (EventExecution_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eventexecution_680c67ac3fada8e2, []int{0, 0}
}

type EventExecution struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MessageId            string                    `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Name                 string                    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Event                string                    `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Created              int64                     `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Status               EventExecution_Status     `protobuf:"varint,6,opt,name=status,proto3,enum=conductor.proto.EventExecution_Status" json:"status,omitempty"`
	Action               EventHandler_Action_Type  `protobuf:"varint,7,opt,name=action,proto3,enum=conductor.proto.EventHandler_Action_Type" json:"action,omitempty"`
	Output               map[string]*_struct.Value `protobuf:"bytes,8,rep,name=output,proto3" json:"output,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EventExecution) Reset()         { *m = EventExecution{} }
func (m *EventExecution) String() string { return proto.CompactTextString(m) }
func (*EventExecution) ProtoMessage()    {}
func (*EventExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventexecution_680c67ac3fada8e2, []int{0}
}
func (m *EventExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventExecution.Unmarshal(m, b)
}
func (m *EventExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventExecution.Marshal(b, m, deterministic)
}
func (dst *EventExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExecution.Merge(dst, src)
}
func (m *EventExecution) XXX_Size() int {
	return xxx_messageInfo_EventExecution.Size(m)
}
func (m *EventExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExecution.DiscardUnknown(m)
}

var xxx_messageInfo_EventExecution proto.InternalMessageInfo

func (m *EventExecution) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventExecution) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *EventExecution) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventExecution) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventExecution) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *EventExecution) GetStatus() EventExecution_Status {
	if m != nil {
		return m.Status
	}
	return EventExecution_IN_PROGRESS
}

func (m *EventExecution) GetAction() EventHandler_Action_Type {
	if m != nil {
		return m.Action
	}
	return EventHandler_Action_START_WORKFLOW
}

func (m *EventExecution) GetOutput() map[string]*_struct.Value {
	if m != nil {
		return m.Output
	}
	return nil
}

func init() {
	proto.RegisterType((*EventExecution)(nil), "conductor.proto.EventExecution")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.EventExecution.OutputEntry")
	proto.RegisterEnum("conductor.proto.EventExecution_Status", EventExecution_Status_name, EventExecution_Status_value)
}

func init() {
	proto.RegisterFile("model/eventexecution.proto", fileDescriptor_eventexecution_680c67ac3fada8e2)
}

var fileDescriptor_eventexecution_680c67ac3fada8e2 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0xb2, 0x4d, 0xed, 0x0b, 0x76, 0xc3, 0x20, 0x32, 0x54, 0x85, 0xb2, 0x07, 0xa9,
	0x28, 0x13, 0xa8, 0x17, 0xd9, 0x83, 0xd0, 0xdd, 0x46, 0x2d, 0xae, 0x36, 0xa6, 0x8b, 0x07, 0x2f,
	0x4b, 0x3a, 0x79, 0x9b, 0x0d, 0x9b, 0xcc, 0x94, 0x64, 0x66, 0xd9, 0xfe, 0xb9, 0xfe, 0x27, 0xd2,
	0x49, 0x22, 0xdd, 0x22, 0xec, 0x6d, 0xde, 0xf7, 0x7d, 0xbf, 0xe4, 0xbd, 0x37, 0x03, 0xa3, 0x52,
	0xa6, 0x58, 0x04, 0x78, 0x87, 0x42, 0xe1, 0x3d, 0x72, 0xad, 0x72, 0x29, 0xd8, 0xa6, 0x92, 0x4a,
	0x92, 0x63, 0x2e, 0x45, 0xaa, 0xb9, 0x92, 0x55, 0x23, 0x8c, 0xe8, 0x5e, 0xf8, 0x26, 0x11, 0x69,
	0x81, 0x9d, 0xf3, 0x2a, 0x93, 0x32, 0x2b, 0x30, 0x30, 0xd5, 0x5a, 0x5f, 0x07, 0xb5, 0xaa, 0x34,
	0x57, 0x8d, 0x7b, 0xf2, 0xc7, 0x81, 0x61, 0xb8, 0x83, 0xc2, 0xee, 0x0f, 0x64, 0x08, 0x76, 0x9e,
	0x52, 0x6b, 0x6c, 0x4d, 0x06, 0xb1, 0x9d, 0xa7, 0xe4, 0x35, 0x40, 0x89, 0x75, 0x9d, 0x64, 0x78,
	0x95, 0xa7, 0xd4, 0x36, 0xfa, 0xa0, 0x55, 0x16, 0x29, 0x21, 0x70, 0x24, 0x92, 0x12, 0xa9, 0x63,
	0x0c, 0x73, 0x26, 0xcf, 0xa1, 0x67, 0x3a, 0xa1, 0x47, 0x46, 0x6c, 0x0a, 0x42, 0xa1, 0xcf, 0x2b,
	0x4c, 0x14, 0xa6, 0xb4, 0x37, 0xb6, 0x26, 0x4e, 0xdc, 0x95, 0xe4, 0x13, 0xb8, 0xb5, 0x4a, 0x94,
	0xae, 0xa9, 0x3b, 0xb6, 0x26, 0xc3, 0xe9, 0x1b, 0x76, 0x30, 0x1f, 0x7b, 0xd8, 0x23, 0x5b, 0x99,
	0x74, 0xdc, 0x52, 0x64, 0x06, 0x6e, 0xc2, 0x77, 0x06, 0xed, 0x1b, 0xfe, 0xed, 0xff, 0xf9, 0xaf,
	0xed, 0x62, 0x66, 0x26, 0xcb, 0x2e, 0xb7, 0x1b, 0x8c, 0x5b, 0x90, 0x9c, 0x83, 0x2b, 0xb5, 0xda,
	0x68, 0x45, 0x9f, 0x8e, 0x9d, 0x89, 0x37, 0x7d, 0xf7, 0x58, 0x0b, 0x4b, 0x93, 0x0e, 0x85, 0xaa,
	0xb6, 0x71, 0x8b, 0x8e, 0x7e, 0x82, 0xb7, 0x27, 0x13, 0x1f, 0x9c, 0x5b, 0xdc, 0xb6, 0xab, 0xdc,
	0x1d, 0xc9, 0x7b, 0xe8, 0xdd, 0x25, 0x85, 0x46, 0xb3, 0x46, 0x6f, 0xfa, 0x82, 0x35, 0x97, 0xc3,
	0xba, 0xcb, 0x61, 0xbf, 0x76, 0x6e, 0xdc, 0x84, 0x4e, 0xed, 0x8f, 0xd6, 0xc9, 0x0c, 0xdc, 0x66,
	0x58, 0x72, 0x0c, 0xde, 0xe2, 0xc7, 0x55, 0x14, 0x2f, 0xbf, 0xc4, 0xe1, 0x6a, 0xe5, 0x3f, 0x21,
	0xcf, 0x60, 0x70, 0xbe, 0xfc, 0x1e, 0x5d, 0x84, 0x97, 0xe1, 0xdc, 0xb7, 0x08, 0x80, 0xfb, 0x79,
	0xb6, 0xb8, 0x08, 0xe7, 0xbe, 0x4d, 0x3c, 0xe8, 0xaf, 0xbe, 0x2d, 0xa2, 0x28, 0x9c, 0xfb, 0xce,
	0xd9, 0x2d, 0xbc, 0xe4, 0xb2, 0x64, 0x02, 0xd5, 0x75, 0x91, 0xdf, 0x1f, 0xce, 0x75, 0xe6, 0x3f,
	0x1c, 0x2c, 0x5a, 0xff, 0x3e, 0xcd, 0x72, 0x75, 0xa3, 0xd7, 0x8c, 0xcb, 0x32, 0x68, 0xa9, 0xe0,
	0x1f, 0x15, 0xf0, 0x22, 0x47, 0xa1, 0x82, 0x4c, 0x66, 0xd5, 0x86, 0xef, 0xe9, 0xe6, 0x05, 0xae,
	0x5d, 0xf3, 0xd1, 0x0f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x30, 0x90, 0x3d, 0xc6, 0xbe, 0x02,
	0x00, 0x00,
}

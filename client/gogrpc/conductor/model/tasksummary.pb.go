// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/tasksummary.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
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

type TaskSummary struct {
	WorkflowId            string      `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	WorkflowType          string      `protobuf:"bytes,2,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type,omitempty"`
	CorrelationId         string      `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ScheduledTime         string      `protobuf:"bytes,4,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
	StartTime             string      `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	UpdateTime            string      `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	EndTime               string      `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Status                Task_Status `protobuf:"varint,8,opt,name=status,proto3,enum=conductor.proto.Task_Status" json:"status,omitempty"`
	ReasonForIncompletion string      `protobuf:"bytes,9,opt,name=reason_for_incompletion,json=reasonForIncompletion,proto3" json:"reason_for_incompletion,omitempty"`
	ExecutionTime         int64       `protobuf:"varint,10,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	QueueWaitTime         int64       `protobuf:"varint,11,opt,name=queue_wait_time,json=queueWaitTime,proto3" json:"queue_wait_time,omitempty"`
	TaskDefName           string      `protobuf:"bytes,12,opt,name=task_def_name,json=taskDefName,proto3" json:"task_def_name,omitempty"`
	TaskType              string      `protobuf:"bytes,13,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	Input                 string      `protobuf:"bytes,14,opt,name=input,proto3" json:"input,omitempty"`
	Output                string      `protobuf:"bytes,15,opt,name=output,proto3" json:"output,omitempty"`
	TaskId                string      `protobuf:"bytes,16,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}    `json:"-"`
	XXX_unrecognized      []byte      `json:"-"`
	XXX_sizecache         int32       `json:"-"`
}

func (m *TaskSummary) Reset()         { *m = TaskSummary{} }
func (m *TaskSummary) String() string { return proto.CompactTextString(m) }
func (*TaskSummary) ProtoMessage()    {}
func (*TaskSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_tasksummary_ab439d130c50da04, []int{0}
}
func (m *TaskSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskSummary.Unmarshal(m, b)
}
func (m *TaskSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskSummary.Marshal(b, m, deterministic)
}
func (dst *TaskSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskSummary.Merge(dst, src)
}
func (m *TaskSummary) XXX_Size() int {
	return xxx_messageInfo_TaskSummary.Size(m)
}
func (m *TaskSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskSummary.DiscardUnknown(m)
}

var xxx_messageInfo_TaskSummary proto.InternalMessageInfo

func (m *TaskSummary) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *TaskSummary) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *TaskSummary) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *TaskSummary) GetScheduledTime() string {
	if m != nil {
		return m.ScheduledTime
	}
	return ""
}

func (m *TaskSummary) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *TaskSummary) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *TaskSummary) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *TaskSummary) GetStatus() Task_Status {
	if m != nil {
		return m.Status
	}
	return Task_IN_PROGRESS
}

func (m *TaskSummary) GetReasonForIncompletion() string {
	if m != nil {
		return m.ReasonForIncompletion
	}
	return ""
}

func (m *TaskSummary) GetExecutionTime() int64 {
	if m != nil {
		return m.ExecutionTime
	}
	return 0
}

func (m *TaskSummary) GetQueueWaitTime() int64 {
	if m != nil {
		return m.QueueWaitTime
	}
	return 0
}

func (m *TaskSummary) GetTaskDefName() string {
	if m != nil {
		return m.TaskDefName
	}
	return ""
}

func (m *TaskSummary) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *TaskSummary) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *TaskSummary) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *TaskSummary) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskSummary)(nil), "conductor.proto.TaskSummary")
}

func init() {
	proto.RegisterFile("model/tasksummary.proto", fileDescriptor_tasksummary_ab439d130c50da04)
}

var fileDescriptor_tasksummary_ab439d130c50da04 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0xa9, 0xbb, 0xdb, 0x1f, 0xaf, 0x3b, 0xed, 0x32, 0xa8, 0x1d, 0x5d, 0x65, 0xcb, 0x8a,
	0xd2, 0xd3, 0x14, 0x54, 0x3c, 0x78, 0x5c, 0x44, 0xe8, 0x45, 0xa4, 0x5b, 0x10, 0xbc, 0x0c, 0x69,
	0xf2, 0xa6, 0x0d, 0x9d, 0x24, 0x63, 0x26, 0xa1, 0xdb, 0x3f, 0xcf, 0xff, 0x4c, 0xf2, 0x32, 0x5b,
	0xcb, 0x1e, 0xf3, 0xf9, 0x7e, 0xf2, 0x92, 0xf7, 0x12, 0x98, 0x28, 0x23, 0xb0, 0x9a, 0x3b, 0xd6,
	0xec, 0x1a, 0xaf, 0x14, 0xb3, 0x87, 0xbc, 0xb6, 0xc6, 0x99, 0x74, 0xcc, 0x8d, 0x16, 0x9e, 0x3b,
	0x63, 0x23, 0x78, 0x7d, 0xf5, 0xdf, 0x8c, 0xe4, 0xf6, 0xef, 0x39, 0x0c, 0x57, 0xac, 0xd9, 0xdd,
	0xc7, 0x8d, 0xe9, 0x0d, 0x0c, 0xf7, 0xc6, 0xee, 0xca, 0xca, 0xec, 0x0b, 0x29, 0xb2, 0xce, 0xb4,
	0x33, 0x1b, 0x2c, 0xe1, 0x11, 0x2d, 0x44, 0xfa, 0x0e, 0x92, 0xa3, 0xe0, 0x0e, 0x35, 0x66, 0xcf,
	0x48, 0xb9, 0x7c, 0x84, 0xab, 0x43, 0x8d, 0xe9, 0x7b, 0x18, 0x71, 0x63, 0x2d, 0x56, 0xcc, 0x49,
	0xa3, 0x43, 0xa1, 0x33, 0xb2, 0x92, 0x13, 0xba, 0x10, 0x41, 0x6b, 0xf8, 0x16, 0x85, 0xaf, 0x50,
	0x14, 0x4e, 0x2a, 0xcc, 0xce, 0xa3, 0x76, 0xa4, 0x2b, 0xa9, 0x30, 0x7d, 0x0b, 0xd0, 0x38, 0x66,
	0x5d, 0x54, 0x2e, 0x48, 0x19, 0x10, 0xa1, 0xf8, 0x06, 0x86, 0xbe, 0x16, 0xcc, 0x61, 0xcc, 0xbb,
	0xf1, 0xca, 0x11, 0x91, 0xf0, 0x0a, 0xfa, 0xa8, 0xdb, 0x03, 0x7a, 0x94, 0xf6, 0x50, 0xc7, 0xd2,
	0x9f, 0xa1, 0xdb, 0x38, 0xe6, 0x7c, 0x93, 0xf5, 0xa7, 0x9d, 0xd9, 0xe8, 0xe3, 0x9b, 0xfc, 0xc9,
	0xc8, 0xf2, 0x30, 0x9c, 0xfc, 0x9e, 0x9c, 0x65, 0xeb, 0xa6, 0x5f, 0x60, 0x62, 0x91, 0x35, 0x46,
	0x17, 0xa5, 0xb1, 0x85, 0xd4, 0xdc, 0xa8, 0xba, 0xc2, 0xd0, 0x54, 0x36, 0xa0, 0xfa, 0x2f, 0x62,
	0xfc, 0xdd, 0xd8, 0xc5, 0x49, 0x18, 0xfa, 0xc5, 0x07, 0xe4, 0x9e, 0x86, 0x42, 0xd7, 0x81, 0x69,
	0x67, 0x76, 0xb6, 0x4c, 0x8e, 0x94, 0x2e, 0xf5, 0x01, 0xc6, 0x7f, 0x3c, 0x7a, 0x2c, 0xf6, 0x4c,
	0xb6, 0x4d, 0x0f, 0xa3, 0x47, 0xf8, 0x17, 0x93, 0xb1, 0xf1, 0x5b, 0x48, 0xc2, 0x4b, 0x16, 0x02,
	0xcb, 0x42, 0x33, 0x85, 0xd9, 0x25, 0x1d, 0x3e, 0x0c, 0xf0, 0x1b, 0x96, 0x3f, 0x98, 0xc2, 0xf4,
	0x1a, 0x06, 0xe4, 0xd0, 0x53, 0x25, 0x94, 0xf7, 0x03, 0xa0, 0x67, 0x7a, 0x0e, 0x17, 0x52, 0xd7,
	0xde, 0x65, 0x23, 0x0a, 0xe2, 0x22, 0x7d, 0x09, 0x5d, 0xe3, 0x5d, 0xc0, 0x63, 0xc2, 0xed, 0x2a,
	0x9d, 0x40, 0x8f, 0x4a, 0x49, 0x91, 0x5d, 0xc5, 0x20, 0x2c, 0x17, 0xe2, 0x6e, 0x0b, 0xd7, 0xdc,
	0xa8, 0x5c, 0xa3, 0x2b, 0x2b, 0xf9, 0xf0, 0x74, 0x82, 0x77, 0xc9, 0xc9, 0xff, 0xfa, 0xb9, 0xfe,
	0xfd, 0x75, 0x23, 0xdd, 0xd6, 0xaf, 0x73, 0x6e, 0xd4, 0xbc, 0xdd, 0x32, 0x3f, 0x6e, 0x99, 0xf3,
	0x4a, 0xa2, 0x76, 0xf3, 0x8d, 0xd9, 0xd8, 0x9a, 0x9f, 0x70, 0xfa, 0xb8, 0xeb, 0x2e, 0x55, 0xfc,
	0xf4, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x32, 0xdb, 0x34, 0x28, 0xf2, 0x02, 0x00, 0x00,
}

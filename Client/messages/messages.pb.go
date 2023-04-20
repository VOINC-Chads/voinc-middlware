// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package messages

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

type Role int32

const (
	Role_ROLE_UNKNOWN   Role = 0
	Role_ROLE_VOLUNTEER Role = 1
)

var Role_name = map[int32]string{
	0: "ROLE_UNKNOWN",
	1: "ROLE_VOLUNTEER",
}

var Role_value = map[string]int32{
	"ROLE_UNKNOWN":   0,
	"ROLE_VOLUNTEER": 1,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

// an enumeration for the status of the message request
type Status int32

const (
	Status_STATUS_UNKNOWN     Status = 0
	Status_STATUS_SUCCESS     Status = 1
	Status_STATUS_FAILURE     Status = 2
	Status_STATUS_CHECK_AGAIN Status = 3
)

var Status_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "STATUS_SUCCESS",
	2: "STATUS_FAILURE",
	3: "STATUS_CHECK_AGAIN",
}

var Status_value = map[string]int32{
	"STATUS_UNKNOWN":     0,
	"STATUS_SUCCESS":     1,
	"STATUS_FAILURE":     2,
	"STATUS_CHECK_AGAIN": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

type MsgTypes int32

const (
	MsgTypes_TYPE_UNKNOWN   MsgTypes = 0
	MsgTypes_TYPE_REGISTER  MsgTypes = 1
	MsgTypes_TYPE_JOB       MsgTypes = 2
	MsgTypes_TYPE_CODE      MsgTypes = 3
	MsgTypes_TYPE_HEARTBEAT MsgTypes = 4
)

var MsgTypes_name = map[int32]string{
	0: "TYPE_UNKNOWN",
	1: "TYPE_REGISTER",
	2: "TYPE_JOB",
	3: "TYPE_CODE",
	4: "TYPE_HEARTBEAT",
}

var MsgTypes_value = map[string]int32{
	"TYPE_UNKNOWN":   0,
	"TYPE_REGISTER":  1,
	"TYPE_JOB":       2,
	"TYPE_CODE":      3,
	"TYPE_HEARTBEAT": 4,
}

func (x MsgTypes) String() string {
	return proto.EnumName(MsgTypes_name, int32(x))
}

func (MsgTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

type Heartbeat struct {
	PendingJobs          uint32   `protobuf:"varint,1,opt,name=pending_jobs,json=pendingJobs,proto3" json:"pending_jobs,omitempty"`
	TotalWorkers         uint32   `protobuf:"varint,2,opt,name=total_workers,json=totalWorkers,proto3" json:"total_workers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

func (m *Heartbeat) GetPendingJobs() uint32 {
	if m != nil {
		return m.PendingJobs
	}
	return 0
}

func (m *Heartbeat) GetTotalWorkers() uint32 {
	if m != nil {
		return m.TotalWorkers
	}
	return 0
}

type RegistrantInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Addr                 string   `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	Capacity             uint32   `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistrantInfo) Reset()         { *m = RegistrantInfo{} }
func (m *RegistrantInfo) String() string { return proto.CompactTextString(m) }
func (*RegistrantInfo) ProtoMessage()    {}
func (*RegistrantInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *RegistrantInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistrantInfo.Unmarshal(m, b)
}
func (m *RegistrantInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistrantInfo.Marshal(b, m, deterministic)
}
func (m *RegistrantInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrantInfo.Merge(m, src)
}
func (m *RegistrantInfo) XXX_Size() int {
	return xxx_messageInfo_RegistrantInfo.Size(m)
}
func (m *RegistrantInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrantInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrantInfo proto.InternalMessageInfo

func (m *RegistrantInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegistrantInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *RegistrantInfo) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RegistrantInfo) GetCapacity() uint32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type CodeMsg struct {
	Requirements         string   `protobuf:"bytes,1,opt,name=requirements,proto3" json:"requirements,omitempty"`
	ProcessCode          string   `protobuf:"bytes,2,opt,name=process_code,json=processCode,proto3" json:"process_code,omitempty"`
	ExecuteCode          string   `protobuf:"bytes,3,opt,name=execute_code,json=executeCode,proto3" json:"execute_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeMsg) Reset()         { *m = CodeMsg{} }
func (m *CodeMsg) String() string { return proto.CompactTextString(m) }
func (*CodeMsg) ProtoMessage()    {}
func (*CodeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *CodeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeMsg.Unmarshal(m, b)
}
func (m *CodeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeMsg.Marshal(b, m, deterministic)
}
func (m *CodeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeMsg.Merge(m, src)
}
func (m *CodeMsg) XXX_Size() int {
	return xxx_messageInfo_CodeMsg.Size(m)
}
func (m *CodeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CodeMsg proto.InternalMessageInfo

func (m *CodeMsg) GetRequirements() string {
	if m != nil {
		return m.Requirements
	}
	return ""
}

func (m *CodeMsg) GetProcessCode() string {
	if m != nil {
		return m.ProcessCode
	}
	return ""
}

func (m *CodeMsg) GetExecuteCode() string {
	if m != nil {
		return m.ExecuteCode
	}
	return ""
}

type RegisterReq struct {
	Role                 Role            `protobuf:"varint,1,opt,name=role,proto3,enum=Role" json:"role,omitempty"`
	Info                 *RegistrantInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_ROLE_UNKNOWN
}

func (m *RegisterReq) GetInfo() *RegistrantInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type JobResult struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Result               string   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobResult) Reset()         { *m = JobResult{} }
func (m *JobResult) String() string { return proto.CompactTextString(m) }
func (*JobResult) ProtoMessage()    {}
func (*JobResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *JobResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobResult.Unmarshal(m, b)
}
func (m *JobResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobResult.Marshal(b, m, deterministic)
}
func (m *JobResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobResult.Merge(m, src)
}
func (m *JobResult) XXX_Size() int {
	return xxx_messageInfo_JobResult.Size(m)
}
func (m *JobResult) XXX_DiscardUnknown() {
	xxx_messageInfo_JobResult.DiscardUnknown(m)
}

var xxx_messageInfo_JobResult proto.InternalMessageInfo

func (m *JobResult) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *JobResult) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type JobResp struct {
	Status               Status       `protobuf:"varint,1,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
	Results              []*JobResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *JobResp) Reset()         { *m = JobResp{} }
func (m *JobResp) String() string { return proto.CompactTextString(m) }
func (*JobResp) ProtoMessage()    {}
func (*JobResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *JobResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobResp.Unmarshal(m, b)
}
func (m *JobResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobResp.Marshal(b, m, deterministic)
}
func (m *JobResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobResp.Merge(m, src)
}
func (m *JobResp) XXX_Size() int {
	return xxx_messageInfo_JobResp.Size(m)
}
func (m *JobResp) XXX_DiscardUnknown() {
	xxx_messageInfo_JobResp.DiscardUnknown(m)
}

var xxx_messageInfo_JobResp proto.InternalMessageInfo

func (m *JobResp) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNKNOWN
}

func (m *JobResp) GetResults() []*JobResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type JobMsg struct {
	Jobs                 []string `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobMsg) Reset()         { *m = JobMsg{} }
func (m *JobMsg) String() string { return proto.CompactTextString(m) }
func (*JobMsg) ProtoMessage()    {}
func (*JobMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{6}
}

func (m *JobMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobMsg.Unmarshal(m, b)
}
func (m *JobMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobMsg.Marshal(b, m, deterministic)
}
func (m *JobMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobMsg.Merge(m, src)
}
func (m *JobMsg) XXX_Size() int {
	return xxx_messageInfo_JobMsg.Size(m)
}
func (m *JobMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_JobMsg.DiscardUnknown(m)
}

var xxx_messageInfo_JobMsg proto.InternalMessageInfo

func (m *JobMsg) GetJobs() []string {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type MainReq struct {
	MsgType MsgTypes `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=MsgTypes" json:"msg_type,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*MainReq_JobMsg
	//	*MainReq_CodeMsg
	//	*MainReq_Heartbeat
	//	*MainReq_RegisterReq
	Content              isMainReq_Content `protobuf_oneof:"Content"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MainReq) Reset()         { *m = MainReq{} }
func (m *MainReq) String() string { return proto.CompactTextString(m) }
func (*MainReq) ProtoMessage()    {}
func (*MainReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{7}
}

func (m *MainReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MainReq.Unmarshal(m, b)
}
func (m *MainReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MainReq.Marshal(b, m, deterministic)
}
func (m *MainReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainReq.Merge(m, src)
}
func (m *MainReq) XXX_Size() int {
	return xxx_messageInfo_MainReq.Size(m)
}
func (m *MainReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MainReq.DiscardUnknown(m)
}

var xxx_messageInfo_MainReq proto.InternalMessageInfo

func (m *MainReq) GetMsgType() MsgTypes {
	if m != nil {
		return m.MsgType
	}
	return MsgTypes_TYPE_UNKNOWN
}

type isMainReq_Content interface {
	isMainReq_Content()
}

type MainReq_JobMsg struct {
	JobMsg *JobMsg `protobuf:"bytes,2,opt,name=job_msg,json=jobMsg,proto3,oneof"`
}

type MainReq_CodeMsg struct {
	CodeMsg *CodeMsg `protobuf:"bytes,3,opt,name=code_msg,json=codeMsg,proto3,oneof"`
}

type MainReq_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,4,opt,name=heartbeat,proto3,oneof"`
}

type MainReq_RegisterReq struct {
	RegisterReq *RegisterReq `protobuf:"bytes,5,opt,name=register_req,json=registerReq,proto3,oneof"`
}

func (*MainReq_JobMsg) isMainReq_Content() {}

func (*MainReq_CodeMsg) isMainReq_Content() {}

func (*MainReq_Heartbeat) isMainReq_Content() {}

func (*MainReq_RegisterReq) isMainReq_Content() {}

func (m *MainReq) GetContent() isMainReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MainReq) GetJobMsg() *JobMsg {
	if x, ok := m.GetContent().(*MainReq_JobMsg); ok {
		return x.JobMsg
	}
	return nil
}

func (m *MainReq) GetCodeMsg() *CodeMsg {
	if x, ok := m.GetContent().(*MainReq_CodeMsg); ok {
		return x.CodeMsg
	}
	return nil
}

func (m *MainReq) GetHeartbeat() *Heartbeat {
	if x, ok := m.GetContent().(*MainReq_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

func (m *MainReq) GetRegisterReq() *RegisterReq {
	if x, ok := m.GetContent().(*MainReq_RegisterReq); ok {
		return x.RegisterReq
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MainReq) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MainReq_JobMsg)(nil),
		(*MainReq_CodeMsg)(nil),
		(*MainReq_Heartbeat)(nil),
		(*MainReq_RegisterReq)(nil),
	}
}

// Response to discovery req will be similar oneof of the responses.
type MainResp struct {
	MsgType MsgTypes `protobuf:"varint,1,opt,name=msg_type,json=msgType,proto3,enum=MsgTypes" json:"msg_type,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*MainResp_JobResp
	//	*MainResp_Heartbeat
	Content              isMainResp_Content `protobuf_oneof:"Content"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MainResp) Reset()         { *m = MainResp{} }
func (m *MainResp) String() string { return proto.CompactTextString(m) }
func (*MainResp) ProtoMessage()    {}
func (*MainResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{8}
}

func (m *MainResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MainResp.Unmarshal(m, b)
}
func (m *MainResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MainResp.Marshal(b, m, deterministic)
}
func (m *MainResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MainResp.Merge(m, src)
}
func (m *MainResp) XXX_Size() int {
	return xxx_messageInfo_MainResp.Size(m)
}
func (m *MainResp) XXX_DiscardUnknown() {
	xxx_messageInfo_MainResp.DiscardUnknown(m)
}

var xxx_messageInfo_MainResp proto.InternalMessageInfo

func (m *MainResp) GetMsgType() MsgTypes {
	if m != nil {
		return m.MsgType
	}
	return MsgTypes_TYPE_UNKNOWN
}

type isMainResp_Content interface {
	isMainResp_Content()
}

type MainResp_JobResp struct {
	JobResp *JobResp `protobuf:"bytes,2,opt,name=job_resp,json=jobResp,proto3,oneof"`
}

type MainResp_Heartbeat struct {
	Heartbeat *Heartbeat `protobuf:"bytes,3,opt,name=heartbeat,proto3,oneof"`
}

func (*MainResp_JobResp) isMainResp_Content() {}

func (*MainResp_Heartbeat) isMainResp_Content() {}

func (m *MainResp) GetContent() isMainResp_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MainResp) GetJobResp() *JobResp {
	if x, ok := m.GetContent().(*MainResp_JobResp); ok {
		return x.JobResp
	}
	return nil
}

func (m *MainResp) GetHeartbeat() *Heartbeat {
	if x, ok := m.GetContent().(*MainResp_Heartbeat); ok {
		return x.Heartbeat
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MainResp) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MainResp_JobResp)(nil),
		(*MainResp_Heartbeat)(nil),
	}
}

func init() {
	proto.RegisterEnum("Role", Role_name, Role_value)
	proto.RegisterEnum("Status", Status_name, Status_value)
	proto.RegisterEnum("MsgTypes", MsgTypes_name, MsgTypes_value)
	proto.RegisterType((*Heartbeat)(nil), "Heartbeat")
	proto.RegisterType((*RegistrantInfo)(nil), "RegistrantInfo")
	proto.RegisterType((*CodeMsg)(nil), "CodeMsg")
	proto.RegisterType((*RegisterReq)(nil), "RegisterReq")
	proto.RegisterType((*JobResult)(nil), "JobResult")
	proto.RegisterType((*JobResp)(nil), "JobResp")
	proto.RegisterType((*JobMsg)(nil), "JobMsg")
	proto.RegisterType((*MainReq)(nil), "MainReq")
	proto.RegisterType((*MainResp)(nil), "MainResp")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x6f, 0x6f, 0xd3, 0x3e,
	0x10, 0xc7, 0xdb, 0xa6, 0x6b, 0x92, 0xeb, 0x9f, 0x5f, 0x7e, 0x16, 0x9a, 0x0a, 0x42, 0x02, 0xb2,
	0x4d, 0x9a, 0x2a, 0x14, 0x89, 0xf1, 0x88, 0x87, 0x5d, 0x09, 0x6b, 0xbb, 0xb5, 0x9d, 0x9c, 0x94,
	0x09, 0x09, 0x29, 0xa4, 0x8d, 0x57, 0x52, 0xda, 0x38, 0xb3, 0x5d, 0x60, 0x6f, 0x82, 0x57, 0xc9,
	0x0b, 0x41, 0x76, 0xdc, 0xad, 0x7d, 0x82, 0x78, 0xe6, 0xfb, 0xdc, 0xd9, 0x77, 0xdf, 0xbb, 0x93,
	0xa1, 0xb5, 0x26, 0x9c, 0xc7, 0x0b, 0xc2, 0xbd, 0x9c, 0x51, 0x41, 0xdd, 0x00, 0xec, 0x3e, 0x89,
	0x99, 0x98, 0x91, 0x58, 0xa0, 0x57, 0xd0, 0xc8, 0x49, 0x96, 0xa4, 0xd9, 0x22, 0x5a, 0xd2, 0x19,
	0x6f, 0x97, 0x5f, 0x96, 0x4f, 0x9b, 0xb8, 0xae, 0xd9, 0x90, 0xce, 0x38, 0x3a, 0x82, 0xa6, 0xa0,
	0x22, 0x5e, 0x45, 0x3f, 0x28, 0xfb, 0x46, 0x18, 0x6f, 0x57, 0x54, 0x4c, 0x43, 0xc1, 0x9b, 0x82,
	0xb9, 0x09, 0xb4, 0x30, 0x59, 0xa4, 0x5c, 0xb0, 0x38, 0x13, 0x83, 0xec, 0x96, 0xa2, 0x16, 0x54,
	0xd2, 0x44, 0xbd, 0x67, 0xe3, 0x4a, 0x9a, 0x20, 0x04, 0xd5, 0x38, 0x49, 0x98, 0xba, 0x6d, 0x63,
	0x75, 0x96, 0x2c, 0xa7, 0x4c, 0xb4, 0x0d, 0xf5, 0xa2, 0x3a, 0xa3, 0x67, 0x60, 0xcd, 0xe3, 0x3c,
	0x9e, 0xa7, 0xe2, 0xbe, 0x5d, 0x55, 0xfc, 0xc1, 0x76, 0x39, 0x98, 0x3d, 0x9a, 0x90, 0x11, 0x5f,
	0x20, 0x17, 0x1a, 0x8c, 0xdc, 0x6d, 0x52, 0x46, 0xd6, 0x24, 0x13, 0x5c, 0x27, 0xda, 0x63, 0x4a,
	0x1c, 0xa3, 0x73, 0xc2, 0x79, 0x34, 0xa7, 0x09, 0xd1, 0xa9, 0xeb, 0x9a, 0xc9, 0x97, 0x64, 0x08,
	0xf9, 0x49, 0xe6, 0x1b, 0x41, 0x8a, 0x10, 0xa3, 0x08, 0xd1, 0x4c, 0x86, 0xb8, 0x23, 0xa8, 0x17,
	0xd2, 0x08, 0xc3, 0xe4, 0x0e, 0x3d, 0x85, 0x2a, 0xa3, 0x2b, 0xa2, 0x12, 0xb6, 0xce, 0x0e, 0x3c,
	0x4c, 0x57, 0x04, 0x2b, 0x84, 0x8e, 0xa0, 0x9a, 0x66, 0xb7, 0x54, 0xe5, 0xa9, 0x9f, 0xfd, 0xe7,
	0xed, 0x77, 0x04, 0x2b, 0xa7, 0xfb, 0x0e, 0xec, 0x21, 0x9d, 0x61, 0xc2, 0x37, 0x2b, 0x81, 0x9e,
	0xc0, 0xc1, 0xf7, 0x78, 0xb5, 0x21, 0xba, 0xfc, 0xc2, 0x40, 0x87, 0x50, 0x63, 0xca, 0xaf, 0x2b,
	0xd6, 0x96, 0x7b, 0x0d, 0x66, 0x71, 0x35, 0x47, 0x2f, 0xa0, 0xc6, 0x45, 0x2c, 0x36, 0x5c, 0xd7,
	0x61, 0x7a, 0x81, 0x32, 0xb1, 0xc6, 0xe8, 0x18, 0xcc, 0xe2, 0x96, 0x9c, 0x97, 0x71, 0x5a, 0x3f,
	0x03, 0xef, 0x21, 0x2d, 0xde, 0xba, 0xdc, 0xe7, 0x50, 0x1b, 0xd2, 0x99, 0xec, 0x27, 0x82, 0xaa,
	0x5e, 0x00, 0x43, 0x8e, 0x47, 0x9e, 0xdd, 0xdf, 0x65, 0x30, 0x47, 0x71, 0x9a, 0x49, 0xd9, 0xc7,
	0x60, 0xad, 0xf9, 0x22, 0x12, 0xf7, 0xf9, 0x56, 0xba, 0xed, 0x8d, 0xf8, 0x22, 0xbc, 0xcf, 0x09,
	0xc7, 0xe6, 0xba, 0x38, 0x21, 0x17, 0xcc, 0x25, 0x9d, 0x45, 0x6b, 0xbe, 0xd0, 0x4d, 0x30, 0xbd,
	0xe2, 0xfd, 0x7e, 0x09, 0xd7, 0x96, 0x45, 0xa6, 0x13, 0xb0, 0x64, 0xab, 0x55, 0x90, 0xa1, 0x82,
	0x2c, 0x4f, 0x4f, 0xb5, 0x5f, 0xc2, 0xe6, 0x5c, 0x0f, 0xb8, 0x03, 0xf6, 0xd7, 0xed, 0x9a, 0xaa,
	0x45, 0x90, 0x12, 0x1e, 0x16, 0xb7, 0x5f, 0xc2, 0x8f, 0x6e, 0xf4, 0x46, 0x2e, 0x43, 0x31, 0xa2,
	0x88, 0x91, 0xbb, 0xf6, 0x81, 0x0a, 0x6f, 0x78, 0x3b, 0x73, 0xeb, 0x97, 0x70, 0x9d, 0x3d, 0x9a,
	0xe7, 0xb6, 0x5c, 0xa5, 0x4c, 0x90, 0x4c, 0xb8, 0xbf, 0xca, 0x60, 0x15, 0x32, 0x79, 0xfe, 0x8f,
	0x3a, 0x4f, 0xc0, 0x92, 0x3a, 0x19, 0xe1, 0xb9, 0x16, 0x6a, 0xe9, 0xf6, 0xe6, 0x52, 0xc3, 0x52,
	0x4f, 0x69, 0x4f, 0x83, 0xf1, 0x57, 0x0d, 0x3b, 0x05, 0x75, 0x5e, 0x43, 0x55, 0x6e, 0x15, 0x72,
	0xa0, 0x81, 0x27, 0x57, 0x7e, 0x34, 0x1d, 0x5f, 0x8e, 0x27, 0x37, 0x63, 0xa7, 0x84, 0x10, 0xb4,
	0x14, 0xf9, 0x38, 0xb9, 0x9a, 0x8e, 0x43, 0xdf, 0xc7, 0x4e, 0xb9, 0xf3, 0x19, 0x6a, 0xc5, 0xec,
	0xa5, 0x37, 0x08, 0xbb, 0xe1, 0x34, 0xd8, 0xbf, 0xa1, 0x59, 0x30, 0xed, 0xf5, 0xfc, 0x20, 0x70,
	0xca, 0x3b, 0xec, 0x43, 0x77, 0x70, 0x35, 0xc5, 0xbe, 0x53, 0x41, 0x87, 0x80, 0x34, 0xeb, 0xf5,
	0xfd, 0xde, 0x65, 0xd4, 0xbd, 0xe8, 0x0e, 0xc6, 0x8e, 0xd1, 0xf9, 0x02, 0xd6, 0x56, 0xbe, 0xac,
	0x27, 0xfc, 0x74, 0xbd, 0x5b, 0xcf, 0xff, 0xd0, 0x54, 0x04, 0xfb, 0x17, 0x83, 0x20, 0x94, 0xe5,
	0xa0, 0x06, 0x58, 0x0a, 0x0d, 0x27, 0xe7, 0x4e, 0x05, 0x35, 0xc1, 0x56, 0x56, 0x6f, 0xf2, 0xde,
	0x77, 0x0c, 0x99, 0x59, 0x99, 0x7d, 0xbf, 0x8b, 0xc3, 0x73, 0xbf, 0x1b, 0x3a, 0xd5, 0x59, 0x4d,
	0x7d, 0x4b, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xb9, 0xb4, 0x5f, 0xa8, 0x04, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: deikstra.proto

package protofiles

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskSubmStatus int32

const (
	TaskSubmStatus_TSS_IQS TaskSubmStatus = 0  // In Queue State
	TaskSubmStatus_TSS_ICS TaskSubmStatus = 1  // In Compilation State
	TaskSubmStatus_TSS_ITS TaskSubmStatus = 2  // In Testing State
	TaskSubmStatus_TSS_CE  TaskSubmStatus = 3  // Compilation Error
	TaskSubmStatus_TSS_TLE TaskSubmStatus = 4  // Time Limit Exceeded
	TaskSubmStatus_TSS_MLE TaskSubmStatus = 5  // Memory Limit Exceed
	TaskSubmStatus_TSS_OK  TaskSubmStatus = 6  // Accepted
	TaskSubmStatus_TSS_PT  TaskSubmStatus = 7  // Partial solution
	TaskSubmStatus_TSS_WA  TaskSubmStatus = 8  // Wrong Answer
	TaskSubmStatus_TSS_RE  TaskSubmStatus = 9  // Runtime Error
	TaskSubmStatus_TSS_PE  TaskSubmStatus = 10 // Presentation Error
	TaskSubmStatus_TSS_ILE TaskSubmStatus = 11 // Idleness Limit Exceeded
	TaskSubmStatus_TSS_CF  TaskSubmStatus = 12 // Check Failed
	TaskSubmStatus_TSS_SV  TaskSubmStatus = 13 // Security Violation
	TaskSubmStatus_TSS_RJ  TaskSubmStatus = 14 // Rejected
	TaskSubmStatus_TSS_DQ  TaskSubmStatus = 15 // Disqualified
)

// Enum value maps for TaskSubmStatus.
var (
	TaskSubmStatus_name = map[int32]string{
		0:  "TSS_IQS",
		1:  "TSS_ICS",
		2:  "TSS_ITS",
		3:  "TSS_CE",
		4:  "TSS_TLE",
		5:  "TSS_MLE",
		6:  "TSS_OK",
		7:  "TSS_PT",
		8:  "TSS_WA",
		9:  "TSS_RE",
		10: "TSS_PE",
		11: "TSS_ILE",
		12: "TSS_CF",
		13: "TSS_SV",
		14: "TSS_RJ",
		15: "TSS_DQ",
	}
	TaskSubmStatus_value = map[string]int32{
		"TSS_IQS": 0,
		"TSS_ICS": 1,
		"TSS_ITS": 2,
		"TSS_CE":  3,
		"TSS_TLE": 4,
		"TSS_MLE": 5,
		"TSS_OK":  6,
		"TSS_PT":  7,
		"TSS_WA":  8,
		"TSS_RE":  9,
		"TSS_PE":  10,
		"TSS_ILE": 11,
		"TSS_CF":  12,
		"TSS_SV":  13,
		"TSS_RJ":  14,
		"TSS_DQ":  15,
	}
)

func (x TaskSubmStatus) Enum() *TaskSubmStatus {
	p := new(TaskSubmStatus)
	*p = x
	return p
}

func (x TaskSubmStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskSubmStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_deikstra_proto_enumTypes[0].Descriptor()
}

func (TaskSubmStatus) Type() protoreflect.EnumType {
	return &file_deikstra_proto_enumTypes[0]
}

func (x TaskSubmStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskSubmStatus.Descriptor instead.
func (TaskSubmStatus) EnumDescriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{0}
}

type TaskTestStatus int32

const (
	TaskTestStatus_TTS_TLE TaskTestStatus = 0  // Time Limit Exceeded
	TaskTestStatus_TTS_MLE TaskTestStatus = 1  // Memory Limit Exceeded
	TaskTestStatus_TTS_OK  TaskTestStatus = 2  // Accepted
	TaskTestStatus_TTS_PT  TaskTestStatus = 3  // Partial solution
	TaskTestStatus_TTS_WA  TaskTestStatus = 4  // Wrong Answer
	TaskTestStatus_TTS_RE  TaskTestStatus = 5  // Runtime Error
	TaskTestStatus_TTS_PE  TaskTestStatus = 6  // Presentation Error
	TaskTestStatus_TTS_ILE TaskTestStatus = 7  // Idleness limit exceeded
	TaskTestStatus_TTS_IG  TaskTestStatus = 8  // Ignored
	TaskTestStatus_TTS_SV  TaskTestStatus = 9  // Security violation
	TaskTestStatus_TTS_CF  TaskTestStatus = 10 // Check Failed
)

// Enum value maps for TaskTestStatus.
var (
	TaskTestStatus_name = map[int32]string{
		0:  "TTS_TLE",
		1:  "TTS_MLE",
		2:  "TTS_OK",
		3:  "TTS_PT",
		4:  "TTS_WA",
		5:  "TTS_RE",
		6:  "TTS_PE",
		7:  "TTS_ILE",
		8:  "TTS_IG",
		9:  "TTS_SV",
		10: "TTS_CF",
	}
	TaskTestStatus_value = map[string]int32{
		"TTS_TLE": 0,
		"TTS_MLE": 1,
		"TTS_OK":  2,
		"TTS_PT":  3,
		"TTS_WA":  4,
		"TTS_RE":  5,
		"TTS_PE":  6,
		"TTS_ILE": 7,
		"TTS_IG":  8,
		"TTS_SV":  9,
		"TTS_CF":  10,
	}
)

func (x TaskTestStatus) Enum() *TaskTestStatus {
	p := new(TaskTestStatus)
	*p = x
	return p
}

func (x TaskTestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskTestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_deikstra_proto_enumTypes[1].Descriptor()
}

func (TaskTestStatus) Type() protoreflect.EnumType {
	return &file_deikstra_proto_enumTypes[1]
}

func (x TaskTestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskTestStatus.Descriptor instead.
func (TaskTestStatus) EnumDescriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{1}
}

type JobAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JobAction) Reset() {
	*x = JobAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobAction) ProtoMessage() {}

func (x *JobAction) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobAction.ProtoReflect.Descriptor instead.
func (*JobAction) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{0}
}

type RegisterWorker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkerName string `protobuf:"bytes,1,opt,name=worker_name,json=workerName,proto3" json:"worker_name,omitempty"`
}

func (x *RegisterWorker) Reset() {
	*x = RegisterWorker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterWorker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterWorker) ProtoMessage() {}

func (x *RegisterWorker) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterWorker.ProtoReflect.Descriptor instead.
func (*RegisterWorker) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterWorker) GetWorkerName() string {
	if x != nil {
		return x.WorkerName
	}
	return ""
}

type TaskSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskCode    string `protobuf:"bytes,1,opt,name=task_code,json=taskCode,proto3" json:"task_code,omitempty"`
	TaskVersion int32  `protobuf:"varint,2,opt,name=task_version,json=taskVersion,proto3" json:"task_version,omitempty"`
	LangCode    string `protobuf:"bytes,3,opt,name=lang_code,json=langCode,proto3" json:"lang_code,omitempty"`
	SubmSrcCode string `protobuf:"bytes,4,opt,name=subm_src_code,json=submSrcCode,proto3" json:"subm_src_code,omitempty"`
}

func (x *TaskSubmission) Reset() {
	*x = TaskSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSubmission) ProtoMessage() {}

func (x *TaskSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSubmission.ProtoReflect.Descriptor instead.
func (*TaskSubmission) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{2}
}

func (x *TaskSubmission) GetTaskCode() string {
	if x != nil {
		return x.TaskCode
	}
	return ""
}

func (x *TaskSubmission) GetTaskVersion() int32 {
	if x != nil {
		return x.TaskVersion
	}
	return 0
}

func (x *TaskSubmission) GetLangCode() string {
	if x != nil {
		return x.LangCode
	}
	return ""
}

func (x *TaskSubmission) GetSubmSrcCode() string {
	if x != nil {
		return x.SubmSrcCode
	}
	return ""
}

type ExecSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LangCode    string `protobuf:"bytes,1,opt,name=lang_code,json=langCode,proto3" json:"lang_code,omitempty"`
	SubmSrcCode string `protobuf:"bytes,2,opt,name=subm_src_code,json=submSrcCode,proto3" json:"subm_src_code,omitempty"`
	StdInput    string `protobuf:"bytes,3,opt,name=std_input,json=stdInput,proto3" json:"std_input,omitempty"`
}

func (x *ExecSubmission) Reset() {
	*x = ExecSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecSubmission) ProtoMessage() {}

func (x *ExecSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecSubmission.ProtoReflect.Descriptor instead.
func (*ExecSubmission) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{3}
}

func (x *ExecSubmission) GetLangCode() string {
	if x != nil {
		return x.LangCode
	}
	return ""
}

func (x *ExecSubmission) GetSubmSrcCode() string {
	if x != nil {
		return x.SubmSrcCode
	}
	return ""
}

func (x *ExecSubmission) GetStdInput() string {
	if x != nil {
		return x.StdInput
	}
	return ""
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId int32 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Types that are assignable to Job:
	//
	//	*Job_TaskSubmission
	//	*Job_ExecSubmission
	Job isJob_Job `protobuf_oneof:"job"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{4}
}

func (x *Job) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (m *Job) GetJob() isJob_Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (x *Job) GetTaskSubmission() *TaskSubmission {
	if x, ok := x.GetJob().(*Job_TaskSubmission); ok {
		return x.TaskSubmission
	}
	return nil
}

func (x *Job) GetExecSubmission() *ExecSubmission {
	if x, ok := x.GetJob().(*Job_ExecSubmission); ok {
		return x.ExecSubmission
	}
	return nil
}

type isJob_Job interface {
	isJob_Job()
}

type Job_TaskSubmission struct {
	TaskSubmission *TaskSubmission `protobuf:"bytes,2,opt,name=task_submission,json=taskSubmission,proto3,oneof"`
}

type Job_ExecSubmission struct {
	ExecSubmission *ExecSubmission `protobuf:"bytes,3,opt,name=exec_submission,json=execSubmission,proto3,oneof"`
}

func (*Job_TaskSubmission) isJob_Job() {}

func (*Job_ExecSubmission) isJob_Job() {}

type TaskSubmResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmStatus TaskSubmStatus `protobuf:"varint,1,opt,name=subm_status,json=submStatus,proto3,enum=protofiles.TaskSubmStatus" json:"subm_status,omitempty"`
}

func (x *TaskSubmResult) Reset() {
	*x = TaskSubmResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskSubmResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskSubmResult) ProtoMessage() {}

func (x *TaskSubmResult) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskSubmResult.ProtoReflect.Descriptor instead.
func (*TaskSubmResult) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{5}
}

func (x *TaskSubmResult) GetSubmStatus() TaskSubmStatus {
	if x != nil {
		return x.SubmStatus
	}
	return TaskSubmStatus_TSS_IQS
}

type TaskTestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestStatus TaskTestStatus `protobuf:"varint,1,opt,name=test_status,json=testStatus,proto3,enum=protofiles.TaskTestStatus" json:"test_status,omitempty"`
}

func (x *TaskTestResult) Reset() {
	*x = TaskTestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskTestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskTestResult) ProtoMessage() {}

func (x *TaskTestResult) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskTestResult.ProtoReflect.Descriptor instead.
func (*TaskTestResult) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{6}
}

func (x *TaskTestResult) GetTestStatus() TaskTestStatus {
	if x != nil {
		return x.TestStatus
	}
	return TaskTestStatus_TTS_TLE
}

type ExecResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *ExecResult) Reset() {
	*x = ExecResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResult) ProtoMessage() {}

func (x *ExecResult) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResult.ProtoReflect.Descriptor instead.
func (*ExecResult) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{7}
}

func (x *ExecResult) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *ExecResult) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

type JobStatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId int32 `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Types that are assignable to Update:
	//
	//	*JobStatusUpdate_TaskRes
	//	*JobStatusUpdate_TestRes
	//	*JobStatusUpdate_ExecRes
	Update isJobStatusUpdate_Update `protobuf_oneof:"update"`
}

func (x *JobStatusUpdate) Reset() {
	*x = JobStatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deikstra_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobStatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobStatusUpdate) ProtoMessage() {}

func (x *JobStatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_deikstra_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobStatusUpdate.ProtoReflect.Descriptor instead.
func (*JobStatusUpdate) Descriptor() ([]byte, []int) {
	return file_deikstra_proto_rawDescGZIP(), []int{8}
}

func (x *JobStatusUpdate) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (m *JobStatusUpdate) GetUpdate() isJobStatusUpdate_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (x *JobStatusUpdate) GetTaskRes() *TaskSubmResult {
	if x, ok := x.GetUpdate().(*JobStatusUpdate_TaskRes); ok {
		return x.TaskRes
	}
	return nil
}

func (x *JobStatusUpdate) GetTestRes() *TaskTestResult {
	if x, ok := x.GetUpdate().(*JobStatusUpdate_TestRes); ok {
		return x.TestRes
	}
	return nil
}

func (x *JobStatusUpdate) GetExecRes() *ExecResult {
	if x, ok := x.GetUpdate().(*JobStatusUpdate_ExecRes); ok {
		return x.ExecRes
	}
	return nil
}

type isJobStatusUpdate_Update interface {
	isJobStatusUpdate_Update()
}

type JobStatusUpdate_TaskRes struct {
	TaskRes *TaskSubmResult `protobuf:"bytes,2,opt,name=task_res,json=taskRes,proto3,oneof"`
}

type JobStatusUpdate_TestRes struct {
	TestRes *TaskTestResult `protobuf:"bytes,3,opt,name=test_res,json=testRes,proto3,oneof"`
}

type JobStatusUpdate_ExecRes struct {
	ExecRes *ExecResult `protobuf:"bytes,4,opt,name=exec_res,json=execRes,proto3,oneof"`
}

func (*JobStatusUpdate_TaskRes) isJobStatusUpdate_Update() {}

func (*JobStatusUpdate_TestRes) isJobStatusUpdate_Update() {}

func (*JobStatusUpdate_ExecRes) isJobStatusUpdate_Update() {}

var File_deikstra_proto protoreflect.FileDescriptor

var file_deikstra_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x65, 0x69, 0x6b, 0x73, 0x74, 0x72, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x0b, 0x0a, 0x09,
	0x4a, 0x6f, 0x62, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x0e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x91, 0x01, 0x0a,
	0x0e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x75, 0x62, 0x6d, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x53, 0x72, 0x63, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x6e, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x6d, 0x5f, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x53, 0x72, 0x63, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x64, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x22, 0xb1, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12,
	0x45, 0x0a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x65,
	0x78, 0x65, 0x63, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x05, 0x0a,
	0x03, 0x6a, 0x6f, 0x62, 0x22, 0x4d, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x4d, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72,
	0x22, 0xd9, 0x01, 0x0a, 0x0f, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x08, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x75, 0x62, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x07, 0x74, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x00, 0x52, 0x07, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x12, 0x33, 0x0a,
	0x08, 0x65, 0x78, 0x65, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x07, 0x65, 0x78, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2a, 0xd6, 0x01, 0x0a,
	0x0e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x75, 0x62, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x53, 0x53, 0x5f, 0x49, 0x51, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x53, 0x53, 0x5f, 0x49, 0x43, 0x53, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x53, 0x53,
	0x5f, 0x49, 0x54, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x5f, 0x43, 0x45,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x53, 0x53, 0x5f, 0x54, 0x4c, 0x45, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x54, 0x53, 0x53, 0x5f, 0x4d, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x53, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x5f,
	0x50, 0x54, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x5f, 0x57, 0x41, 0x10, 0x08,
	0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x5f, 0x52, 0x45, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06,
	0x54, 0x53, 0x53, 0x5f, 0x50, 0x45, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x53, 0x53, 0x5f,
	0x49, 0x4c, 0x45, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x5f, 0x43, 0x46, 0x10,
	0x0c, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53, 0x5f, 0x53, 0x56, 0x10, 0x0d, 0x12, 0x0a, 0x0a,
	0x06, 0x54, 0x53, 0x53, 0x5f, 0x52, 0x4a, 0x10, 0x0e, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x53, 0x53,
	0x5f, 0x44, 0x51, 0x10, 0x0f, 0x2a, 0x97, 0x01, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x65,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x54, 0x53, 0x5f,
	0x54, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x54, 0x53, 0x5f, 0x4d, 0x4c, 0x45,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x54, 0x53, 0x5f, 0x50, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54,
	0x53, 0x5f, 0x57, 0x41, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54, 0x53, 0x5f, 0x52, 0x45,
	0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54, 0x53, 0x5f, 0x50, 0x45, 0x10, 0x06, 0x12, 0x0b,
	0x0a, 0x07, 0x54, 0x54, 0x53, 0x5f, 0x49, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x54,
	0x54, 0x53, 0x5f, 0x49, 0x47, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54, 0x53, 0x5f, 0x53,
	0x56, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x54, 0x53, 0x5f, 0x43, 0x46, 0x10, 0x0a, 0x32,
	0x92, 0x01, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x22, 0x00, 0x30, 0x01, 0x12, 0x49, 0x0a, 0x0f, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4b, 0x72, 0x69, 0x73, 0x6a, 0x61, 0x6e, 0x69, 0x73, 0x50, 0x2f, 0x64, 0x65,
	0x69, 0x6b, 0x73, 0x74, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_deikstra_proto_rawDescOnce sync.Once
	file_deikstra_proto_rawDescData = file_deikstra_proto_rawDesc
)

func file_deikstra_proto_rawDescGZIP() []byte {
	file_deikstra_proto_rawDescOnce.Do(func() {
		file_deikstra_proto_rawDescData = protoimpl.X.CompressGZIP(file_deikstra_proto_rawDescData)
	})
	return file_deikstra_proto_rawDescData
}

var file_deikstra_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_deikstra_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_deikstra_proto_goTypes = []interface{}{
	(TaskSubmStatus)(0),     // 0: protofiles.TaskSubmStatus
	(TaskTestStatus)(0),     // 1: protofiles.TaskTestStatus
	(*JobAction)(nil),       // 2: protofiles.JobAction
	(*RegisterWorker)(nil),  // 3: protofiles.RegisterWorker
	(*TaskSubmission)(nil),  // 4: protofiles.TaskSubmission
	(*ExecSubmission)(nil),  // 5: protofiles.ExecSubmission
	(*Job)(nil),             // 6: protofiles.Job
	(*TaskSubmResult)(nil),  // 7: protofiles.TaskSubmResult
	(*TaskTestResult)(nil),  // 8: protofiles.TaskTestResult
	(*ExecResult)(nil),      // 9: protofiles.ExecResult
	(*JobStatusUpdate)(nil), // 10: protofiles.JobStatusUpdate
}
var file_deikstra_proto_depIdxs = []int32{
	4,  // 0: protofiles.Job.task_submission:type_name -> protofiles.TaskSubmission
	5,  // 1: protofiles.Job.exec_submission:type_name -> protofiles.ExecSubmission
	0,  // 2: protofiles.TaskSubmResult.subm_status:type_name -> protofiles.TaskSubmStatus
	1,  // 3: protofiles.TaskTestResult.test_status:type_name -> protofiles.TaskTestStatus
	7,  // 4: protofiles.JobStatusUpdate.task_res:type_name -> protofiles.TaskSubmResult
	8,  // 5: protofiles.JobStatusUpdate.test_res:type_name -> protofiles.TaskTestResult
	9,  // 6: protofiles.JobStatusUpdate.exec_res:type_name -> protofiles.ExecResult
	3,  // 7: protofiles.Scheduler.GetJobs:input_type -> protofiles.RegisterWorker
	10, // 8: protofiles.Scheduler.ReportJobStatus:input_type -> protofiles.JobStatusUpdate
	6,  // 9: protofiles.Scheduler.GetJobs:output_type -> protofiles.Job
	2,  // 10: protofiles.Scheduler.ReportJobStatus:output_type -> protofiles.JobAction
	9,  // [9:11] is the sub-list for method output_type
	7,  // [7:9] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_deikstra_proto_init() }
func file_deikstra_proto_init() {
	if File_deikstra_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deikstra_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterWorker); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSubmission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecSubmission); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskSubmResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskTestResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deikstra_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobStatusUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_deikstra_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Job_TaskSubmission)(nil),
		(*Job_ExecSubmission)(nil),
	}
	file_deikstra_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*JobStatusUpdate_TaskRes)(nil),
		(*JobStatusUpdate_TestRes)(nil),
		(*JobStatusUpdate_ExecRes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deikstra_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_deikstra_proto_goTypes,
		DependencyIndexes: file_deikstra_proto_depIdxs,
		EnumInfos:         file_deikstra_proto_enumTypes,
		MessageInfos:      file_deikstra_proto_msgTypes,
	}.Build()
	File_deikstra_proto = out.File
	file_deikstra_proto_rawDesc = nil
	file_deikstra_proto_goTypes = nil
	file_deikstra_proto_depIdxs = nil
}

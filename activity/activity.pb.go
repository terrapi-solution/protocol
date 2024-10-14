// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: activity/activity.proto

package activity

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

type InsertActivityRequest_Pointer int32

const (
	InsertActivityRequest_stdout InsertActivityRequest_Pointer = 0
	InsertActivityRequest_stderr InsertActivityRequest_Pointer = 1
)

// Enum value maps for InsertActivityRequest_Pointer.
var (
	InsertActivityRequest_Pointer_name = map[int32]string{
		0: "stdout",
		1: "stderr",
	}
	InsertActivityRequest_Pointer_value = map[string]int32{
		"stdout": 0,
		"stderr": 1,
	}
)

func (x InsertActivityRequest_Pointer) Enum() *InsertActivityRequest_Pointer {
	p := new(InsertActivityRequest_Pointer)
	*p = x
	return p
}

func (x InsertActivityRequest_Pointer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InsertActivityRequest_Pointer) Descriptor() protoreflect.EnumDescriptor {
	return file_activity_activity_proto_enumTypes[0].Descriptor()
}

func (InsertActivityRequest_Pointer) Type() protoreflect.EnumType {
	return &file_activity_activity_proto_enumTypes[0]
}

func (x InsertActivityRequest_Pointer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InsertActivityRequest_Pointer.Descriptor instead.
func (InsertActivityRequest_Pointer) EnumDescriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{2, 0}
}

type Activity_Pointer int32

const (
	Activity_stdout Activity_Pointer = 0
	Activity_stderr Activity_Pointer = 1
)

// Enum value maps for Activity_Pointer.
var (
	Activity_Pointer_name = map[int32]string{
		0: "stdout",
		1: "stderr",
	}
	Activity_Pointer_value = map[string]int32{
		"stdout": 0,
		"stderr": 1,
	}
)

func (x Activity_Pointer) Enum() *Activity_Pointer {
	p := new(Activity_Pointer)
	*p = x
	return p
}

func (x Activity_Pointer) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Activity_Pointer) Descriptor() protoreflect.EnumDescriptor {
	return file_activity_activity_proto_enumTypes[1].Descriptor()
}

func (Activity_Pointer) Type() protoreflect.EnumType {
	return &file_activity_activity_proto_enumTypes[1]
}

func (x Activity_Pointer) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Activity_Pointer.Descriptor instead.
func (Activity_Pointer) EnumDescriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{3, 0}
}

type InsertActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertActivityResponse) Reset() {
	*x = InsertActivityResponse{}
	mi := &file_activity_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertActivityResponse) ProtoMessage() {}

func (x *InsertActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_activity_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertActivityResponse.ProtoReflect.Descriptor instead.
func (*InsertActivityResponse) Descriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{0}
}

func (x *InsertActivityResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListActivityRequest) Reset() {
	*x = ListActivityRequest{}
	mi := &file_activity_activity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActivityRequest) ProtoMessage() {}

func (x *ListActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_activity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActivityRequest.ProtoReflect.Descriptor instead.
func (*ListActivityRequest) Descriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{1}
}

func (x *ListActivityRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type InsertActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployment int32                         `protobuf:"varint,2,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Pointer    InsertActivityRequest_Pointer `protobuf:"varint,3,opt,name=pointer,proto3,enum=protocol.InsertActivityRequest_Pointer" json:"pointer,omitempty"`
	Message    string                        `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InsertActivityRequest) Reset() {
	*x = InsertActivityRequest{}
	mi := &file_activity_activity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InsertActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertActivityRequest) ProtoMessage() {}

func (x *InsertActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_activity_activity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertActivityRequest.ProtoReflect.Descriptor instead.
func (*InsertActivityRequest) Descriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{2}
}

func (x *InsertActivityRequest) GetDeployment() int32 {
	if x != nil {
		return x.Deployment
	}
	return 0
}

func (x *InsertActivityRequest) GetPointer() InsertActivityRequest_Pointer {
	if x != nil {
		return x.Pointer
	}
	return InsertActivityRequest_stdout
}

func (x *InsertActivityRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Deployment int32            `protobuf:"varint,2,opt,name=deployment,proto3" json:"deployment,omitempty"`
	Pointer    Activity_Pointer `protobuf:"varint,3,opt,name=pointer,proto3,enum=protocol.Activity_Pointer" json:"pointer,omitempty"`
	Message    string           `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp  int32            `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	mi := &file_activity_activity_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_activity_activity_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{3}
}

func (x *Activity) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Activity) GetDeployment() int32 {
	if x != nil {
		return x.Deployment
	}
	return 0
}

func (x *Activity) GetPointer() Activity_Pointer {
	if x != nil {
		return x.Pointer
	}
	return Activity_stdout
}

func (x *Activity) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Activity) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Activities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Activities []*Activity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
}

func (x *Activities) Reset() {
	*x = Activities{}
	mi := &file_activity_activity_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Activities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activities) ProtoMessage() {}

func (x *Activities) ProtoReflect() protoreflect.Message {
	mi := &file_activity_activity_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activities.ProtoReflect.Descriptor instead.
func (*Activities) Descriptor() ([]byte, []int) {
	return file_activity_activity_proto_rawDescGZIP(), []int{4}
}

func (x *Activities) GetActivities() []*Activity {
	if x != nil {
		return x.Activities
	}
	return nil
}

var File_activity_activity_proto protoreflect.FileDescriptor

var file_activity_activity_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xb7, 0x01, 0x0a,
	0x15, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0a,
	0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74,
	0x64, 0x65, 0x72, 0x72, 0x10, 0x01, 0x22, 0xcb, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x52, 0x07, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x21, 0x0a, 0x07, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0a, 0x0a, 0x06,
	0x73, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65,
	0x72, 0x72, 0x10, 0x01, 0x22, 0x40, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0x9f, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x49, 0x6e,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x61, 0x70, 0x69, 0x2d, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_activity_activity_proto_rawDescOnce sync.Once
	file_activity_activity_proto_rawDescData = file_activity_activity_proto_rawDesc
)

func file_activity_activity_proto_rawDescGZIP() []byte {
	file_activity_activity_proto_rawDescOnce.Do(func() {
		file_activity_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_activity_activity_proto_rawDescData)
	})
	return file_activity_activity_proto_rawDescData
}

var file_activity_activity_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_activity_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_activity_activity_proto_goTypes = []any{
	(InsertActivityRequest_Pointer)(0), // 0: protocol.InsertActivityRequest.Pointer
	(Activity_Pointer)(0),              // 1: protocol.Activity.Pointer
	(*InsertActivityResponse)(nil),     // 2: protocol.InsertActivityResponse
	(*ListActivityRequest)(nil),        // 3: protocol.ListActivityRequest
	(*InsertActivityRequest)(nil),      // 4: protocol.InsertActivityRequest
	(*Activity)(nil),                   // 5: protocol.Activity
	(*Activities)(nil),                 // 6: protocol.Activities
}
var file_activity_activity_proto_depIdxs = []int32{
	0, // 0: protocol.InsertActivityRequest.pointer:type_name -> protocol.InsertActivityRequest.Pointer
	1, // 1: protocol.Activity.pointer:type_name -> protocol.Activity.Pointer
	5, // 2: protocol.Activities.activities:type_name -> protocol.Activity
	4, // 3: protocol.ActivityService.Insert:input_type -> protocol.InsertActivityRequest
	3, // 4: protocol.ActivityService.List:input_type -> protocol.ListActivityRequest
	2, // 5: protocol.ActivityService.Insert:output_type -> protocol.InsertActivityResponse
	6, // 6: protocol.ActivityService.List:output_type -> protocol.Activities
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_activity_activity_proto_init() }
func file_activity_activity_proto_init() {
	if File_activity_activity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_activity_activity_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_activity_activity_proto_goTypes,
		DependencyIndexes: file_activity_activity_proto_depIdxs,
		EnumInfos:         file_activity_activity_proto_enumTypes,
		MessageInfos:      file_activity_activity_proto_msgTypes,
	}.Build()
	File_activity_activity_proto = out.File
	file_activity_activity_proto_rawDesc = nil
	file_activity_activity_proto_goTypes = nil
	file_activity_activity_proto_depIdxs = nil
}

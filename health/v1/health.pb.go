/*
 * Licensed to the TerrAPI under one or more contributor
 * license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The TerrAPI licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: health/v1/health.proto

package health

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

type CheckResponse_ServingStatus int32

const (
	CheckResponse_SERVING_STATUS_UNSPECIFIED     CheckResponse_ServingStatus = 0
	CheckResponse_SERVING_STATUS_UNKNOWN         CheckResponse_ServingStatus = 1
	CheckResponse_SERVING_STATUS_SERVING         CheckResponse_ServingStatus = 2
	CheckResponse_SERVING_STATUS_NOT_SERVING     CheckResponse_ServingStatus = 3
	CheckResponse_SERVING_STATUS_SERVICE_UNKNOWN CheckResponse_ServingStatus = 4
	CheckResponse_SERVING_STATUS_INTERNAL_ERROR  CheckResponse_ServingStatus = 5
)

// Enum value maps for CheckResponse_ServingStatus.
var (
	CheckResponse_ServingStatus_name = map[int32]string{
		0: "SERVING_STATUS_UNSPECIFIED",
		1: "SERVING_STATUS_UNKNOWN",
		2: "SERVING_STATUS_SERVING",
		3: "SERVING_STATUS_NOT_SERVING",
		4: "SERVING_STATUS_SERVICE_UNKNOWN",
		5: "SERVING_STATUS_INTERNAL_ERROR",
	}
	CheckResponse_ServingStatus_value = map[string]int32{
		"SERVING_STATUS_UNSPECIFIED":     0,
		"SERVING_STATUS_UNKNOWN":         1,
		"SERVING_STATUS_SERVING":         2,
		"SERVING_STATUS_NOT_SERVING":     3,
		"SERVING_STATUS_SERVICE_UNKNOWN": 4,
		"SERVING_STATUS_INTERNAL_ERROR":  5,
	}
)

func (x CheckResponse_ServingStatus) Enum() *CheckResponse_ServingStatus {
	p := new(CheckResponse_ServingStatus)
	*p = x
	return p
}

func (x CheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_health_v1_health_proto_enumTypes[0].Descriptor()
}

func (CheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_health_v1_health_proto_enumTypes[0]
}

func (x CheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckResponse_ServingStatus.Descriptor instead.
func (CheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{2, 0}
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	mi := &file_health_v1_health_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type CheckAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *CheckAllRequest) Reset() {
	*x = CheckAllRequest{}
	mi := &file_health_v1_health_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAllRequest) ProtoMessage() {}

func (x *CheckAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAllRequest.ProtoReflect.Descriptor instead.
func (*CheckAllRequest) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{1}
}

func (x *CheckAllRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status CheckResponse_ServingStatus `protobuf:"varint,2,opt,name=status,proto3,enum=health.v1.CheckResponse_ServingStatus" json:"status,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	mi := &file_health_v1_health_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{2}
}

func (x *CheckResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CheckResponse) GetStatus() CheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return CheckResponse_SERVING_STATUS_UNSPECIFIED
}

type CheckAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*CheckResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *CheckAllResponse) Reset() {
	*x = CheckAllResponse{}
	mi := &file_health_v1_health_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAllResponse) ProtoMessage() {}

func (x *CheckAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_v1_health_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAllResponse.ProtoReflect.Descriptor instead.
func (*CheckAllResponse) Descriptor() ([]byte, []int) {
	return file_health_v1_health_proto_rawDescGZIP(), []int{3}
}

func (x *CheckAllResponse) GetResults() []*CheckResponse {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_health_v1_health_proto protoreflect.FileDescriptor

var file_health_v1_health_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a,
	0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xb4, 0x02, 0x0a, 0x0d, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xce,
	0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x22, 0x0a, 0x1e, 0x53, 0x45, 0x52, 0x56,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x22,
	0x46, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x94, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_v1_health_proto_rawDescOnce sync.Once
	file_health_v1_health_proto_rawDescData = file_health_v1_health_proto_rawDesc
)

func file_health_v1_health_proto_rawDescGZIP() []byte {
	file_health_v1_health_proto_rawDescOnce.Do(func() {
		file_health_v1_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_v1_health_proto_rawDescData)
	})
	return file_health_v1_health_proto_rawDescData
}

var file_health_v1_health_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_health_v1_health_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_health_v1_health_proto_goTypes = []any{
	(CheckResponse_ServingStatus)(0), // 0: health.v1.CheckResponse.ServingStatus
	(*CheckRequest)(nil),             // 1: health.v1.CheckRequest
	(*CheckAllRequest)(nil),          // 2: health.v1.CheckAllRequest
	(*CheckResponse)(nil),            // 3: health.v1.CheckResponse
	(*CheckAllResponse)(nil),         // 4: health.v1.CheckAllResponse
}
var file_health_v1_health_proto_depIdxs = []int32{
	0, // 0: health.v1.CheckResponse.status:type_name -> health.v1.CheckResponse.ServingStatus
	3, // 1: health.v1.CheckAllResponse.results:type_name -> health.v1.CheckResponse
	1, // 2: health.v1.HealthService.Check:input_type -> health.v1.CheckRequest
	2, // 3: health.v1.HealthService.CheckAll:input_type -> health.v1.CheckAllRequest
	3, // 4: health.v1.HealthService.Check:output_type -> health.v1.CheckResponse
	4, // 5: health.v1.HealthService.CheckAll:output_type -> health.v1.CheckAllResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_health_v1_health_proto_init() }
func file_health_v1_health_proto_init() {
	if File_health_v1_health_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_health_v1_health_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_v1_health_proto_goTypes,
		DependencyIndexes: file_health_v1_health_proto_depIdxs,
		EnumInfos:         file_health_v1_health_proto_enumTypes,
		MessageInfos:      file_health_v1_health_proto_msgTypes,
	}.Build()
	File_health_v1_health_proto = out.File
	file_health_v1_health_proto_rawDesc = nil
	file_health_v1_health_proto_goTypes = nil
	file_health_v1_health_proto_depIdxs = nil
}

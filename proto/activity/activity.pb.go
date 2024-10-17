// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: proto/activity/activity.proto

package yuemnoi_activity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActivityLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	LogDetail          string                 `protobuf:"bytes,2,opt,name=logDetail,proto3" json:"logDetail,omitempty"`
	UserId             string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	FormattedTimestamp string                 `protobuf:"bytes,5,opt,name=formattedTimestamp,proto3" json:"formattedTimestamp,omitempty"`
}

func (x *ActivityLog) Reset() {
	*x = ActivityLog{}
	mi := &file_proto_activity_activity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActivityLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityLog) ProtoMessage() {}

func (x *ActivityLog) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_activity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityLog.ProtoReflect.Descriptor instead.
func (*ActivityLog) Descriptor() ([]byte, []int) {
	return file_proto_activity_activity_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityLog) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ActivityLog) GetLogDetail() string {
	if x != nil {
		return x.LogDetail
	}
	return ""
}

func (x *ActivityLog) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ActivityLog) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ActivityLog) GetFormattedTimestamp() string {
	if x != nil {
		return x.FormattedTimestamp
	}
	return ""
}

type CreateActivityLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogDetail string `protobuf:"bytes,1,opt,name=logDetail,proto3" json:"logDetail,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateActivityLogRequest) Reset() {
	*x = CreateActivityLogRequest{}
	mi := &file_proto_activity_activity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateActivityLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActivityLogRequest) ProtoMessage() {}

func (x *CreateActivityLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_activity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateActivityLogRequest.ProtoReflect.Descriptor instead.
func (*CreateActivityLogRequest) Descriptor() ([]byte, []int) {
	return file_proto_activity_activity_proto_rawDescGZIP(), []int{1}
}

func (x *CreateActivityLogRequest) GetLogDetail() string {
	if x != nil {
		return x.LogDetail
	}
	return ""
}

func (x *CreateActivityLogRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateActivityLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityLog *ActivityLog `protobuf:"bytes,1,opt,name=activityLog,proto3" json:"activityLog,omitempty"`
}

func (x *CreateActivityLogResponse) Reset() {
	*x = CreateActivityLogResponse{}
	mi := &file_proto_activity_activity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateActivityLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateActivityLogResponse) ProtoMessage() {}

func (x *CreateActivityLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_activity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateActivityLogResponse.ProtoReflect.Descriptor instead.
func (*CreateActivityLogResponse) Descriptor() ([]byte, []int) {
	return file_proto_activity_activity_proto_rawDescGZIP(), []int{2}
}

func (x *CreateActivityLogResponse) GetActivityLog() *ActivityLog {
	if x != nil {
		return x.ActivityLog
	}
	return nil
}

type ViewActivityHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ViewActivityHistoryRequest) Reset() {
	*x = ViewActivityHistoryRequest{}
	mi := &file_proto_activity_activity_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewActivityHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewActivityHistoryRequest) ProtoMessage() {}

func (x *ViewActivityHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_activity_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewActivityHistoryRequest.ProtoReflect.Descriptor instead.
func (*ViewActivityHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_activity_activity_proto_rawDescGZIP(), []int{3}
}

func (x *ViewActivityHistoryRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type ViewActivityHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityLog []*ActivityLog `protobuf:"bytes,1,rep,name=activityLog,proto3" json:"activityLog,omitempty"`
}

func (x *ViewActivityHistoryResponse) Reset() {
	*x = ViewActivityHistoryResponse{}
	mi := &file_proto_activity_activity_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ViewActivityHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewActivityHistoryResponse) ProtoMessage() {}

func (x *ViewActivityHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_activity_activity_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewActivityHistoryResponse.ProtoReflect.Descriptor instead.
func (*ViewActivityHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_activity_activity_proto_rawDescGZIP(), []int{4}
}

func (x *ViewActivityHistoryResponse) GetActivityLog() []*ActivityLog {
	if x != nil {
		return x.ActivityLog
	}
	return nil
}

var File_proto_activity_activity_proto protoreflect.FileDescriptor

var file_proto_activity_activity_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbe, 0x01, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x51, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x6f, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x4c, 0x6f, 0x67, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f,
	0x67, 0x22, 0x35, 0x0a, 0x1a, 0x56, 0x69, 0x65, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x1b, 0x56, 0x69, 0x65, 0x77,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x32, 0xb6, 0x01, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x4c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13,
	0x56, 0x69, 0x65, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x34,
	0x53, 0x61, 0x6f, 0x31, 0x53, 0x61, 0x64, 0x2f, 0x79, 0x75, 0x65, 0x6d, 0x6e, 0x6f, 0x69, 0x2d,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_activity_activity_proto_rawDescOnce sync.Once
	file_proto_activity_activity_proto_rawDescData = file_proto_activity_activity_proto_rawDesc
)

func file_proto_activity_activity_proto_rawDescGZIP() []byte {
	file_proto_activity_activity_proto_rawDescOnce.Do(func() {
		file_proto_activity_activity_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_activity_activity_proto_rawDescData)
	})
	return file_proto_activity_activity_proto_rawDescData
}

var file_proto_activity_activity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_activity_activity_proto_goTypes = []any{
	(*ActivityLog)(nil),                 // 0: ActivityLog
	(*CreateActivityLogRequest)(nil),    // 1: CreateActivityLogRequest
	(*CreateActivityLogResponse)(nil),   // 2: CreateActivityLogResponse
	(*ViewActivityHistoryRequest)(nil),  // 3: ViewActivityHistoryRequest
	(*ViewActivityHistoryResponse)(nil), // 4: ViewActivityHistoryResponse
	(*timestamppb.Timestamp)(nil),       // 5: google.protobuf.Timestamp
}
var file_proto_activity_activity_proto_depIdxs = []int32{
	5, // 0: ActivityLog.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: CreateActivityLogResponse.activityLog:type_name -> ActivityLog
	0, // 2: ViewActivityHistoryResponse.activityLog:type_name -> ActivityLog
	1, // 3: ActivityLogService.CreateActivityLog:input_type -> CreateActivityLogRequest
	3, // 4: ActivityLogService.ViewActivityHistory:input_type -> ViewActivityHistoryRequest
	2, // 5: ActivityLogService.CreateActivityLog:output_type -> CreateActivityLogResponse
	4, // 6: ActivityLogService.ViewActivityHistory:output_type -> ViewActivityHistoryResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_activity_activity_proto_init() }
func file_proto_activity_activity_proto_init() {
	if File_proto_activity_activity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_activity_activity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_activity_activity_proto_goTypes,
		DependencyIndexes: file_proto_activity_activity_proto_depIdxs,
		MessageInfos:      file_proto_activity_activity_proto_msgTypes,
	}.Build()
	File_proto_activity_activity_proto = out.File
	file_proto_activity_activity_proto_rawDesc = nil
	file_proto_activity_activity_proto_goTypes = nil
	file_proto_activity_activity_proto_depIdxs = nil
}

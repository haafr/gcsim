// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: protos/model/notification.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComputeFailedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DbId          string                 `protobuf:"bytes,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty" bson:"dbId,omitempty"`
	Config        string                 `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	Submitter     string                 `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"`
	Reason        string                 `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty" bson:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ComputeFailedEvent) Reset() {
	*x = ComputeFailedEvent{}
	mi := &file_protos_model_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComputeFailedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeFailedEvent) ProtoMessage() {}

func (x *ComputeFailedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeFailedEvent.ProtoReflect.Descriptor instead.
func (*ComputeFailedEvent) Descriptor() ([]byte, []int) {
	return file_protos_model_notification_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeFailedEvent) GetDbId() string {
	if x != nil {
		return x.DbId
	}
	return ""
}

func (x *ComputeFailedEvent) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *ComputeFailedEvent) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

func (x *ComputeFailedEvent) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ComputeCompletedEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DbId          string                 `protobuf:"bytes,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty" bson:"dbId,omitempty"`
	ShareId       string                 `protobuf:"bytes,2,opt,name=share_id,json=shareId,proto3" json:"share_id,omitempty" bson:"shareId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ComputeCompletedEvent) Reset() {
	*x = ComputeCompletedEvent{}
	mi := &file_protos_model_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComputeCompletedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeCompletedEvent) ProtoMessage() {}

func (x *ComputeCompletedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeCompletedEvent.ProtoReflect.Descriptor instead.
func (*ComputeCompletedEvent) Descriptor() ([]byte, []int) {
	return file_protos_model_notification_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeCompletedEvent) GetDbId() string {
	if x != nil {
		return x.DbId
	}
	return ""
}

func (x *ComputeCompletedEvent) GetShareId() string {
	if x != nil {
		return x.ShareId
	}
	return ""
}

type SubmissionDeleteEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DbId          string                 `protobuf:"bytes,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty" bson:"dbId,omitempty"`
	Config        string                 `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	Submitter     string                 `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty" bson:"submitter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmissionDeleteEvent) Reset() {
	*x = SubmissionDeleteEvent{}
	mi := &file_protos_model_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmissionDeleteEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionDeleteEvent) ProtoMessage() {}

func (x *SubmissionDeleteEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionDeleteEvent.ProtoReflect.Descriptor instead.
func (*SubmissionDeleteEvent) Descriptor() ([]byte, []int) {
	return file_protos_model_notification_proto_rawDescGZIP(), []int{2}
}

func (x *SubmissionDeleteEvent) GetDbId() string {
	if x != nil {
		return x.DbId
	}
	return ""
}

func (x *SubmissionDeleteEvent) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *SubmissionDeleteEvent) GetSubmitter() string {
	if x != nil {
		return x.Submitter
	}
	return ""
}

type EntryReplaceEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DbId          string                 `protobuf:"bytes,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty" bson:"dbId,omitempty"`
	Config        string                 `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty" bson:"config,omitempty"`
	OldConfig     string                 `protobuf:"bytes,3,opt,name=old_config,json=oldConfig,proto3" json:"old_config,omitempty" bson:"oldConfig,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EntryReplaceEvent) Reset() {
	*x = EntryReplaceEvent{}
	mi := &file_protos_model_notification_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntryReplaceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryReplaceEvent) ProtoMessage() {}

func (x *EntryReplaceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_notification_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryReplaceEvent.ProtoReflect.Descriptor instead.
func (*EntryReplaceEvent) Descriptor() ([]byte, []int) {
	return file_protos_model_notification_proto_rawDescGZIP(), []int{3}
}

func (x *EntryReplaceEvent) GetDbId() string {
	if x != nil {
		return x.DbId
	}
	return ""
}

func (x *EntryReplaceEvent) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *EntryReplaceEvent) GetOldConfig() string {
	if x != nil {
		return x.OldConfig
	}
	return ""
}

type DescReplaceEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DbId          string                 `protobuf:"bytes,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty" bson:"dbId,omitempty"`
	Desc          string                 `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" bson:"desc,omitempty"`
	OldDesc       string                 `protobuf:"bytes,3,opt,name=old_desc,json=oldDesc,proto3" json:"old_desc,omitempty" bson:"oldDesc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DescReplaceEvent) Reset() {
	*x = DescReplaceEvent{}
	mi := &file_protos_model_notification_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescReplaceEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescReplaceEvent) ProtoMessage() {}

func (x *DescReplaceEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_notification_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescReplaceEvent.ProtoReflect.Descriptor instead.
func (*DescReplaceEvent) Descriptor() ([]byte, []int) {
	return file_protos_model_notification_proto_rawDescGZIP(), []int{4}
}

func (x *DescReplaceEvent) GetDbId() string {
	if x != nil {
		return x.DbId
	}
	return ""
}

func (x *DescReplaceEvent) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *DescReplaceEvent) GetOldDesc() string {
	if x != nil {
		return x.OldDesc
	}
	return ""
}

var File_protos_model_notification_proto protoreflect.FileDescriptor

var file_protos_model_notification_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x77, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13,
	0x0a, 0x05, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x62, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x64, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x62, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x65, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x15, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x62, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x5f,
	0x0a, 0x11, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x62, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x56, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x6c, 0x64, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x6c, 0x64, 0x44, 0x65, 0x73, 0x63, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x6e, 0x73, 0x69, 0x6d,
	0x2f, 0x67, 0x63, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protos_model_notification_proto_rawDescOnce sync.Once
	file_protos_model_notification_proto_rawDescData []byte
)

func file_protos_model_notification_proto_rawDescGZIP() []byte {
	file_protos_model_notification_proto_rawDescOnce.Do(func() {
		file_protos_model_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_model_notification_proto_rawDesc), len(file_protos_model_notification_proto_rawDesc)))
	})
	return file_protos_model_notification_proto_rawDescData
}

var file_protos_model_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_model_notification_proto_goTypes = []any{
	(*ComputeFailedEvent)(nil),    // 0: model.ComputeFailedEvent
	(*ComputeCompletedEvent)(nil), // 1: model.ComputeCompletedEvent
	(*SubmissionDeleteEvent)(nil), // 2: model.SubmissionDeleteEvent
	(*EntryReplaceEvent)(nil),     // 3: model.EntryReplaceEvent
	(*DescReplaceEvent)(nil),      // 4: model.DescReplaceEvent
}
var file_protos_model_notification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_model_notification_proto_init() }
func file_protos_model_notification_proto_init() {
	if File_protos_model_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_model_notification_proto_rawDesc), len(file_protos_model_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_notification_proto_goTypes,
		DependencyIndexes: file_protos_model_notification_proto_depIdxs,
		MessageInfos:      file_protos_model_notification_proto_msgTypes,
	}.Build()
	File_protos_model_notification_proto = out.File
	file_protos_model_notification_proto_goTypes = nil
	file_protos_model_notification_proto_depIdxs = nil
}

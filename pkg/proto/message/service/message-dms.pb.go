// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.21.12
// source: proto/message/service/message-dms.proto

package insync_service

import (
	model1 "github.com/duongquoctue/insync/pkg/proto/common/model"
	model "github.com/duongquoctue/insync/pkg/proto/message/model"
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

type MessageSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *model.Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageSendRequest) Reset() {
	*x = MessageSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_message_dms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageSendRequest) ProtoMessage() {}

func (x *MessageSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_message_dms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageSendRequest.ProtoReflect.Descriptor instead.
func (*MessageSendRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_service_message_dms_proto_rawDescGZIP(), []int{0}
}

func (x *MessageSendRequest) GetMessage() *model.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_proto_message_service_message_dms_proto protoreflect.FileDescriptor

var file_proto_message_service_message_dms_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2d,
	0x64, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x73, 0x67, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x21,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x54,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x53, 0x65, 0x6e,
	0x64, 0x12, 0x27, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x6f, 0x6e, 0x67, 0x71, 0x75, 0x6f, 0x63, 0x74, 0x75, 0x65, 0x2f,
	0x69, 0x6e, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x3b, 0x69, 0x6e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_service_message_dms_proto_rawDescOnce sync.Once
	file_proto_message_service_message_dms_proto_rawDescData = file_proto_message_service_message_dms_proto_rawDesc
)

func file_proto_message_service_message_dms_proto_rawDescGZIP() []byte {
	file_proto_message_service_message_dms_proto_rawDescOnce.Do(func() {
		file_proto_message_service_message_dms_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_service_message_dms_proto_rawDescData)
	})
	return file_proto_message_service_message_dms_proto_rawDescData
}

var file_proto_message_service_message_dms_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_message_service_message_dms_proto_goTypes = []interface{}{
	(*MessageSendRequest)(nil), // 0: msg.message.service.MessageSendRequest
	(*model.Message)(nil),      // 1: msg.message.model.Message
	(*model1.Status)(nil),      // 2: msg.common.model.Status
}
var file_proto_message_service_message_dms_proto_depIdxs = []int32{
	1, // 0: msg.message.service.MessageSendRequest.message:type_name -> msg.message.model.Message
	0, // 1: msg.message.service.Message.Send:input_type -> msg.message.service.MessageSendRequest
	2, // 2: msg.message.service.Message.Send:output_type -> msg.common.model.Status
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_message_service_message_dms_proto_init() }
func file_proto_message_service_message_dms_proto_init() {
	if File_proto_message_service_message_dms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_service_message_dms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageSendRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_service_message_dms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_message_service_message_dms_proto_goTypes,
		DependencyIndexes: file_proto_message_service_message_dms_proto_depIdxs,
		MessageInfos:      file_proto_message_service_message_dms_proto_msgTypes,
	}.Build()
	File_proto_message_service_message_dms_proto = out.File
	file_proto_message_service_message_dms_proto_rawDesc = nil
	file_proto_message_service_message_dms_proto_goTypes = nil
	file_proto_message_service_message_dms_proto_depIdxs = nil
}

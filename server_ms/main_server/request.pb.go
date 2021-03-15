// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: main_server/main_server.proto

package main_server

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

type CommandPackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	ChatId  string `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *CommandPackRequest) Reset() {
	*x = CommandPackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandPackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandPackRequest) ProtoMessage() {}

func (x *CommandPackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandPackRequest.ProtoReflect.Descriptor instead.
func (*CommandPackRequest) Descriptor() ([]byte, []int) {
	return file_request_request_proto_rawDescGZIP(), []int{0}
}

func (x *CommandPackRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandPackRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type CommandPackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CommandPackResponse) Reset() {
	*x = CommandPackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandPackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandPackResponse) ProtoMessage() {}

func (x *CommandPackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_request_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandPackResponse.ProtoReflect.Descriptor instead.
func (*CommandPackResponse) Descriptor() ([]byte, []int) {
	return file_request_request_proto_rawDescGZIP(), []int{1}
}

func (x *CommandPackResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *CommandPackResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_request_request_proto protoreflect.FileDescriptor

var file_request_request_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x70, 0x61, 0x63, 0x6b, 0x22, 0x47, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x49, 0x0a,
	0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x68, 0x0a, 0x12, 0x54, 0x65, 0x6c, 0x65,
	0x67, 0x72, 0x61, 0x6d, 0x42, 0x6f, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x12, 0x1f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x6d, 0x6f, 0x6e, 0x6a, 0x6a, 0x75, 0x62, 0x6f, 0x74, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_request_request_proto_rawDescOnce sync.Once
	file_request_request_proto_rawDescData = file_request_request_proto_rawDesc
)

func file_request_request_proto_rawDescGZIP() []byte {
	file_request_request_proto_rawDescOnce.Do(func() {
		file_request_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_request_proto_rawDescData)
	})
	return file_request_request_proto_rawDescData
}

var file_request_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_request_request_proto_goTypes = []interface{}{
	(*CommandPackRequest)(nil),  // 0: commandpack.CommandPackRequest
	(*CommandPackResponse)(nil), // 1: commandpack.CommandPackResponse
}
var file_request_request_proto_depIdxs = []int32{
	0, // 0: commandpack.TelegramBotService.CommandPack:input_type -> commandpack.CommandPackRequest
	1, // 1: commandpack.TelegramBotService.CommandPack:output_type -> commandpack.CommandPackResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_request_request_proto_init() }
func file_request_request_proto_init() {
	if File_request_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_request_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandPackRequest); i {
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
		file_request_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandPackResponse); i {
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
			RawDescriptor: file_request_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_request_request_proto_goTypes,
		DependencyIndexes: file_request_request_proto_depIdxs,
		MessageInfos:      file_request_request_proto_msgTypes,
	}.Build()
	File_request_request_proto = out.File
	file_request_request_proto_rawDesc = nil
	file_request_request_proto_goTypes = nil
	file_request_request_proto_depIdxs = nil
}

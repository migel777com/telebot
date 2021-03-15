// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: database/db.proto

package database

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId int64  `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Subject  string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{0}
}

func (x *BookRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *BookRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type BigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookPacks []*BookPack `protobuf:"bytes,1,rep,name=bookPacks,proto3" json:"bookPacks,omitempty"`
}

func (x *BigResponse) Reset() {
	*x = BigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigResponse) ProtoMessage() {}

func (x *BigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigResponse.ProtoReflect.Descriptor instead.
func (*BigResponse) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{1}
}

func (x *BigResponse) GetBookPacks() []*BookPack {
	if x != nil {
		return x.BookPacks
	}
	return nil
}

type BookPack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId   int64  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	CourseId int64  `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Subject  string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	BookName string `protobuf:"bytes,4,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	BookLink string `protobuf:"bytes,5,opt,name=book_link,json=bookLink,proto3" json:"book_link,omitempty"`
}

func (x *BookPack) Reset() {
	*x = BookPack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookPack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookPack) ProtoMessage() {}

func (x *BookPack) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookPack.ProtoReflect.Descriptor instead.
func (*BookPack) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{2}
}

func (x *BookPack) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BookPack) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *BookPack) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *BookPack) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *BookPack) GetBookLink() string {
	if x != nil {
		return x.BookLink
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId    string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserEmail string `protobuf:"bytes,2,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	Vkey      string `protobuf:"bytes,3,opt,name=vkey,proto3" json:"vkey,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *RegisterRequest) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *RegisterRequest) GetVkey() string {
	if x != nil {
		return x.Vkey
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConfirmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vkey string `protobuf:"bytes,1,opt,name=vkey,proto3" json:"vkey,omitempty"`
}

func (x *ConfirmRequest) Reset() {
	*x = ConfirmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmRequest) ProtoMessage() {}

func (x *ConfirmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmRequest.ProtoReflect.Descriptor instead.
func (*ConfirmRequest) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmRequest) GetVkey() string {
	if x != nil {
		return x.Vkey
	}
	return ""
}

type VerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chatid string `protobuf:"bytes,1,opt,name=chatid,proto3" json:"chatid,omitempty"`
}

func (x *VerificationRequest) Reset() {
	*x = VerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationRequest) ProtoMessage() {}

func (x *VerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationRequest.ProtoReflect.Descriptor instead.
func (*VerificationRequest) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{6}
}

func (x *VerificationRequest) GetChatid() string {
	if x != nil {
		return x.Chatid
	}
	return ""
}

type VerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VerificationResponse) Reset() {
	*x = VerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationResponse) ProtoMessage() {}

func (x *VerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationResponse.ProtoReflect.Descriptor instead.
func (*VerificationResponse) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{7}
}

func (x *VerificationResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_databaseproto_db_proto protoreflect.FileDescriptor

var file_databaseproto_db_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x44, 0x0a,
	0x0b, 0x62, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x09,
	0x62, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x50, 0x61,
	0x63, 0x6b, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x63, 0x6b,
	0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x5d, 0x0a, 0x0f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6b, 0x65, 0x79, 0x22, 0x44, 0x0a, 0x10, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x24, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x76, 0x6b, 0x65, 0x79, 0x22, 0x2e, 0x0a, 0x14, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x68, 0x61, 0x74, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xe7, 0x02, 0x0a, 0x15, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x44, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60,
	0x0a, 0x11, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x19, 0x5a, 0x17, 0x6d, 0x6f, 0x6e, 0x6a, 0x6a, 0x75, 0x62, 0x6f, 0x74, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_databaseproto_db_proto_rawDescOnce sync.Once
	file_databaseproto_db_proto_rawDescData = file_databaseproto_db_proto_rawDesc
)

func file_databaseproto_db_proto_rawDescGZIP() []byte {
	file_databaseproto_db_proto_rawDescOnce.Do(func() {
		file_databaseproto_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_databaseproto_db_proto_rawDescData)
	})
	return file_databaseproto_db_proto_rawDescData
}

var file_databaseproto_db_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_databaseproto_db_proto_goTypes = []interface{}{
	(*BookRequest)(nil),          // 0: database.bookRequest
	(*BigResponse)(nil),          // 1: database.bigResponse
	(*BookPack)(nil),             // 2: database.BookPack
	(*RegisterRequest)(nil),      // 3: database.registerRequest
	(*RegisterResponse)(nil),     // 4: database.registerResponse
	(*ConfirmRequest)(nil),       // 5: database.confirmRequest
	(*VerificationRequest)(nil),  // 6: database.verification_request
	(*VerificationResponse)(nil), // 7: database.verification_response
}
var file_databaseproto_db_proto_depIdxs = []int32{
	2, // 0: database.bigResponse.bookPacks:type_name -> database.BookPack
	0, // 1: database.DatabaseAccessService.getBooks:input_type -> database.bookRequest
	3, // 2: database.DatabaseAccessService.registerUser:input_type -> database.registerRequest
	5, // 3: database.DatabaseAccessService.confirmRegister:input_type -> database.confirmRequest
	6, // 4: database.DatabaseAccessService.checkVerification:input_type -> database.verification_request
	1, // 5: database.DatabaseAccessService.getBooks:output_type -> database.bigResponse
	4, // 6: database.DatabaseAccessService.registerUser:output_type -> database.registerResponse
	4, // 7: database.DatabaseAccessService.confirmRegister:output_type -> database.registerResponse
	7, // 8: database.DatabaseAccessService.checkVerification:output_type -> database.verification_response
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_databaseproto_db_proto_init() }
func file_databaseproto_db_proto_init() {
	if File_databaseproto_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_databaseproto_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRequest); i {
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
		file_databaseproto_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigResponse); i {
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
		file_databaseproto_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookPack); i {
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
		file_databaseproto_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_databaseproto_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_databaseproto_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmRequest); i {
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
		file_databaseproto_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationRequest); i {
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
		file_databaseproto_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationResponse); i {
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
			RawDescriptor: file_databaseproto_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_databaseproto_db_proto_goTypes,
		DependencyIndexes: file_databaseproto_db_proto_depIdxs,
		MessageInfos:      file_databaseproto_db_proto_msgTypes,
	}.Build()
	File_databaseproto_db_proto = out.File
	file_databaseproto_db_proto_rawDesc = nil
	file_databaseproto_db_proto_goTypes = nil
	file_databaseproto_db_proto_depIdxs = nil
}

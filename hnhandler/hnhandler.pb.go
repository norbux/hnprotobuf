// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: hnhandler/hnhandler.proto

package hnprotobuf

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

// The User request message. Represents the parameter used for querying the user endpoint.
type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hnhandler_hnhandler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hnhandler_hnhandler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_hnhandler_hnhandler_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// The user response message. Represents the fields of the user object.
type UserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//google.protobuf.Timestamp created = 2;
	Created   int64   `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Karma     int32   `protobuf:"varint,3,opt,name=karma,proto3" json:"karma,omitempty"`
	About     string  `protobuf:"bytes,4,opt,name=about,proto3" json:"about,omitempty"`
	Submitted []int64 `protobuf:"varint,5,rep,packed,name=submitted,proto3" json:"submitted,omitempty"`
}

func (x *UserReply) Reset() {
	*x = UserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hnhandler_hnhandler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReply) ProtoMessage() {}

func (x *UserReply) ProtoReflect() protoreflect.Message {
	mi := &file_hnhandler_hnhandler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReply.ProtoReflect.Descriptor instead.
func (*UserReply) Descriptor() ([]byte, []int) {
	return file_hnhandler_hnhandler_proto_rawDescGZIP(), []int{1}
}

func (x *UserReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserReply) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *UserReply) GetKarma() int32 {
	if x != nil {
		return x.Karma
	}
	return 0
}

func (x *UserReply) GetAbout() string {
	if x != nil {
		return x.About
	}
	return ""
}

func (x *UserReply) GetSubmitted() []int64 {
	if x != nil {
		return x.Submitted
	}
	return nil
}

// The item request message. Represents the parameter used for querying the items endpoint.
type ItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ItemRequest) Reset() {
	*x = ItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hnhandler_hnhandler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRequest) ProtoMessage() {}

func (x *ItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hnhandler_hnhandler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRequest.ProtoReflect.Descriptor instead.
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return file_hnhandler_hnhandler_proto_rawDescGZIP(), []int{2}
}

func (x *ItemRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// The item response message. Represents the fiels of the item object.
type ItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Deleted     bool    `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Type        string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	By          string  `protobuf:"bytes,4,opt,name=by,proto3" json:"by,omitempty"`
	Time        int64   `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	Text        string  `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Dead        bool    `protobuf:"varint,7,opt,name=dead,proto3" json:"dead,omitempty"`
	Parent      int64   `protobuf:"varint,8,opt,name=parent,proto3" json:"parent,omitempty"`
	Poll        int64   `protobuf:"varint,9,opt,name=poll,proto3" json:"poll,omitempty"`
	Kids        []int64 `protobuf:"varint,10,rep,packed,name=kids,proto3" json:"kids,omitempty"`
	Url         string  `protobuf:"bytes,11,opt,name=url,proto3" json:"url,omitempty"`
	Score       int32   `protobuf:"varint,12,opt,name=score,proto3" json:"score,omitempty"`
	Title       string  `protobuf:"bytes,13,opt,name=title,proto3" json:"title,omitempty"`
	Parts       []int64 `protobuf:"varint,14,rep,packed,name=parts,proto3" json:"parts,omitempty"`
	Descendants int32   `protobuf:"varint,15,opt,name=descendants,proto3" json:"descendants,omitempty"`
}

func (x *ItemReply) Reset() {
	*x = ItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hnhandler_hnhandler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemReply) ProtoMessage() {}

func (x *ItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_hnhandler_hnhandler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemReply.ProtoReflect.Descriptor instead.
func (*ItemReply) Descriptor() ([]byte, []int) {
	return file_hnhandler_hnhandler_proto_rawDescGZIP(), []int{3}
}

func (x *ItemReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemReply) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *ItemReply) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ItemReply) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

func (x *ItemReply) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *ItemReply) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ItemReply) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

func (x *ItemReply) GetParent() int64 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *ItemReply) GetPoll() int64 {
	if x != nil {
		return x.Poll
	}
	return 0
}

func (x *ItemReply) GetKids() []int64 {
	if x != nil {
		return x.Kids
	}
	return nil
}

func (x *ItemReply) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ItemReply) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *ItemReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ItemReply) GetParts() []int64 {
	if x != nil {
		return x.Parts
	}
	return nil
}

func (x *ItemReply) GetDescendants() int32 {
	if x != nil {
		return x.Descendants
	}
	return 0
}

// The new stories request message. Represents the parameter used for querying
// the new stories endpoint. In this case it is an empty message type because
// the method would not require any input parameter. This is a way to avoid
// breaking code in case a parameter would be required for this service.
// Another option to achieve this would be to use:
//
//  import "google/protobuf/empty.proto";
//  ...
//  rpc SomeService(google.protobuf.Empty) returns (Whoami) {
//      ...
//  }
//
// Refer to: https://stackoverflow.com/questions/29687243/protobuf-rpc-service-method-without-parameters
type NewStoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewStoriesRequest) Reset() {
	*x = NewStoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hnhandler_hnhandler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewStoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewStoriesRequest) ProtoMessage() {}

func (x *NewStoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hnhandler_hnhandler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewStoriesRequest.ProtoReflect.Descriptor instead.
func (*NewStoriesRequest) Descriptor() ([]byte, []int) {
	return file_hnhandler_hnhandler_proto_rawDescGZIP(), []int{4}
}

// The new stories response message. Represents the collection of storie IDs.
// message NewStoriesItem {
//     int64 id = 1;
// }
type NewStoriesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// int32 val = 2;
	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *NewStoriesReply) Reset() {
	*x = NewStoriesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hnhandler_hnhandler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewStoriesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewStoriesReply) ProtoMessage() {}

func (x *NewStoriesReply) ProtoReflect() protoreflect.Message {
	mi := &file_hnhandler_hnhandler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewStoriesReply.ProtoReflect.Descriptor instead.
func (*NewStoriesReply) Descriptor() ([]byte, []int) {
	return file_hnhandler_hnhandler_proto_rawDescGZIP(), []int{5}
}

func (x *NewStoriesReply) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_hnhandler_hnhandler_proto protoreflect.FileDescriptor

var file_hnhandler_hnhandler_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x68, 0x6e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x68, 0x6e, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x22, 0x1d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6b, 0x61, 0x72, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6b, 0x61, 0x72,
	0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcb, 0x02, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x62,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x61,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x64,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x6b, 0x69, 0x64, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x74, 0x73,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x61,
	0x6e, 0x74, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x53,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x45, 0x0a, 0x0a, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x68, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68,
	0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x32, 0x45, 0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x2e, 0x68, 0x6e,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x5d, 0x0a, 0x10, 0x4e, 0x65, 0x77,
	0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x49, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x68, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x74,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68,
	0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x72, 0x74, 0x68, 0x2d, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x2f, 0x68, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hnhandler_hnhandler_proto_rawDescOnce sync.Once
	file_hnhandler_hnhandler_proto_rawDescData = file_hnhandler_hnhandler_proto_rawDesc
)

func file_hnhandler_hnhandler_proto_rawDescGZIP() []byte {
	file_hnhandler_hnhandler_proto_rawDescOnce.Do(func() {
		file_hnhandler_hnhandler_proto_rawDescData = protoimpl.X.CompressGZIP(file_hnhandler_hnhandler_proto_rawDescData)
	})
	return file_hnhandler_hnhandler_proto_rawDescData
}

var file_hnhandler_hnhandler_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hnhandler_hnhandler_proto_goTypes = []interface{}{
	(*UserRequest)(nil),       // 0: hnhandler.UserRequest
	(*UserReply)(nil),         // 1: hnhandler.UserReply
	(*ItemRequest)(nil),       // 2: hnhandler.ItemRequest
	(*ItemReply)(nil),         // 3: hnhandler.ItemReply
	(*NewStoriesRequest)(nil), // 4: hnhandler.NewStoriesRequest
	(*NewStoriesReply)(nil),   // 5: hnhandler.NewStoriesReply
}
var file_hnhandler_hnhandler_proto_depIdxs = []int32{
	0, // 0: hnhandler.UserGetter.GetUser:input_type -> hnhandler.UserRequest
	2, // 1: hnhandler.ItemGetter.GetItem:input_type -> hnhandler.ItemRequest
	4, // 2: hnhandler.NewStoriesGetter.GetNewStories:input_type -> hnhandler.NewStoriesRequest
	1, // 3: hnhandler.UserGetter.GetUser:output_type -> hnhandler.UserReply
	3, // 4: hnhandler.ItemGetter.GetItem:output_type -> hnhandler.ItemReply
	5, // 5: hnhandler.NewStoriesGetter.GetNewStories:output_type -> hnhandler.NewStoriesReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hnhandler_hnhandler_proto_init() }
func file_hnhandler_hnhandler_proto_init() {
	if File_hnhandler_hnhandler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hnhandler_hnhandler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_hnhandler_hnhandler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReply); i {
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
		file_hnhandler_hnhandler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRequest); i {
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
		file_hnhandler_hnhandler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemReply); i {
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
		file_hnhandler_hnhandler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewStoriesRequest); i {
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
		file_hnhandler_hnhandler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewStoriesReply); i {
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
			RawDescriptor: file_hnhandler_hnhandler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_hnhandler_hnhandler_proto_goTypes,
		DependencyIndexes: file_hnhandler_hnhandler_proto_depIdxs,
		MessageInfos:      file_hnhandler_hnhandler_proto_msgTypes,
	}.Build()
	File_hnhandler_hnhandler_proto = out.File
	file_hnhandler_hnhandler_proto_rawDesc = nil
	file_hnhandler_hnhandler_proto_goTypes = nil
	file_hnhandler_hnhandler_proto_depIdxs = nil
}

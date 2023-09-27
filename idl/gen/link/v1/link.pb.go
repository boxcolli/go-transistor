// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: link/v1/link.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`      // any length that ensures uniqueness within current network level
	Role int32  `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"` // 0: publisher, 1: subscriber, 2?: hybrid
	Addr string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`  // full grpc address with port
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{0}
}

func (x *Entity) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Entity) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

func (x *Entity) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type GetLevelRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLevelRangeRequest) Reset() {
	*x = GetLevelRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLevelRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLevelRangeRequest) ProtoMessage() {}

func (x *GetLevelRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLevelRangeRequest.ProtoReflect.Descriptor instead.
func (*GetLevelRangeRequest) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{1}
}

type GetLevelRangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int32 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int32 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetLevelRangeResponse) Reset() {
	*x = GetLevelRangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLevelRangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLevelRangeResponse) ProtoMessage() {}

func (x *GetLevelRangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLevelRangeResponse.ProtoReflect.Descriptor instead.
func (*GetLevelRangeResponse) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{2}
}

func (x *GetLevelRangeResponse) GetStart() int32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetLevelRangeResponse) GetEnd() int32 {
	if x != nil {
		return x.End
	}
	return 0
}

type GetTopologyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *GetTopologyRequest) Reset() {
	*x = GetTopologyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopologyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopologyRequest) ProtoMessage() {}

func (x *GetTopologyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopologyRequest.ProtoReflect.Descriptor instead.
func (*GetTopologyRequest) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{3}
}

func (x *GetTopologyRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type GetTopologyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"` // list
}

func (x *GetTopologyResponse) Reset() {
	*x = GetTopologyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopologyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopologyResponse) ProtoMessage() {}

func (x *GetTopologyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopologyResponse.ProtoReflect.Descriptor instead.
func (*GetTopologyResponse) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopologyResponse) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"` // This is not a global network level. 0: Same network tree level of current node
	Role  int32 `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{5}
}

func (x *JoinRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *JoinRequest) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*Entity `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"` // list
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{6}
}

func (x *JoinResponse) GetEntities() []*Entity {
	if x != nil {
		return x.Entities
	}
	return nil
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []string `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{7}
}

func (x *SubscribeRequest) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	TopicId   []byte                 `protobuf:"bytes,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	Data      *anypb.Any             `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_v1_link_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_link_v1_link_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_link_v1_link_proto_rawDescGZIP(), []int{8}
}

func (x *SubscribeResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SubscribeResponse) GetTopicId() []byte {
	if x != nil {
		return x.TopicId
	}
	return nil
}

func (x *SubscribeResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SubscribeResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_link_v1_link_proto protoreflect.FileDescriptor

var file_link_v1_link_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x06, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x22, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x37, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x3b, 0x0a,
	0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x32, 0xac, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x12, 0x1b, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x14, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_v1_link_proto_rawDescOnce sync.Once
	file_link_v1_link_proto_rawDescData = file_link_v1_link_proto_rawDesc
)

func file_link_v1_link_proto_rawDescGZIP() []byte {
	file_link_v1_link_proto_rawDescOnce.Do(func() {
		file_link_v1_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_v1_link_proto_rawDescData)
	})
	return file_link_v1_link_proto_rawDescData
}

var file_link_v1_link_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_link_v1_link_proto_goTypes = []interface{}{
	(*Entity)(nil),                // 0: link.v1.Entity
	(*GetLevelRangeRequest)(nil),  // 1: link.v1.GetLevelRangeRequest
	(*GetLevelRangeResponse)(nil), // 2: link.v1.GetLevelRangeResponse
	(*GetTopologyRequest)(nil),    // 3: link.v1.GetTopologyRequest
	(*GetTopologyResponse)(nil),   // 4: link.v1.GetTopologyResponse
	(*JoinRequest)(nil),           // 5: link.v1.JoinRequest
	(*JoinResponse)(nil),          // 6: link.v1.JoinResponse
	(*SubscribeRequest)(nil),      // 7: link.v1.SubscribeRequest
	(*SubscribeResponse)(nil),     // 8: link.v1.SubscribeResponse
	(*anypb.Any)(nil),             // 9: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_link_v1_link_proto_depIdxs = []int32{
	0,  // 0: link.v1.GetTopologyResponse.entities:type_name -> link.v1.Entity
	0,  // 1: link.v1.JoinResponse.entities:type_name -> link.v1.Entity
	9,  // 2: link.v1.SubscribeResponse.data:type_name -> google.protobuf.Any
	10, // 3: link.v1.SubscribeResponse.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 4: link.v1.LinkService.GetLevelRange:input_type -> link.v1.GetLevelRangeRequest
	3,  // 5: link.v1.LinkService.GetTopology:input_type -> link.v1.GetTopologyRequest
	5,  // 6: link.v1.LinkService.Join:input_type -> link.v1.JoinRequest
	7,  // 7: link.v1.LinkService.Subscribe:input_type -> link.v1.SubscribeRequest
	2,  // 8: link.v1.LinkService.GetLevelRange:output_type -> link.v1.GetLevelRangeResponse
	4,  // 9: link.v1.LinkService.GetTopology:output_type -> link.v1.GetTopologyResponse
	6,  // 10: link.v1.LinkService.Join:output_type -> link.v1.JoinResponse
	8,  // 11: link.v1.LinkService.Subscribe:output_type -> link.v1.SubscribeResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_link_v1_link_proto_init() }
func file_link_v1_link_proto_init() {
	if File_link_v1_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_link_v1_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_link_v1_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLevelRangeRequest); i {
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
		file_link_v1_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLevelRangeResponse); i {
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
		file_link_v1_link_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopologyRequest); i {
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
		file_link_v1_link_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopologyResponse); i {
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
		file_link_v1_link_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_link_v1_link_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_link_v1_link_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_link_v1_link_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
			RawDescriptor: file_link_v1_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_v1_link_proto_goTypes,
		DependencyIndexes: file_link_v1_link_proto_depIdxs,
		MessageInfos:      file_link_v1_link_proto_msgTypes,
	}.Build()
	File_link_v1_link_proto = out.File
	file_link_v1_link_proto_rawDesc = nil
	file_link_v1_link_proto_goTypes = nil
	file_link_v1_link_proto_depIdxs = nil
}

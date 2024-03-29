// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/comet.proto

package zyjgrpc

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

type PushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BroadcastReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BroadcastReq) Reset() {
	*x = BroadcastReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReq) ProtoMessage() {}

func (x *BroadcastReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReq.ProtoReflect.Descriptor instead.
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{1}
}

func (x *BroadcastReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BroadcastGroupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID string `protobuf:"bytes,1,opt,name=groupID,proto3" json:"groupID,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BroadcastGroupReq) Reset() {
	*x = BroadcastGroupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastGroupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastGroupReq) ProtoMessage() {}

func (x *BroadcastGroupReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastGroupReq.ProtoReflect.Descriptor instead.
func (*BroadcastGroupReq) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{2}
}

func (x *BroadcastGroupReq) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *BroadcastGroupReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GroupsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GroupsReq) Reset() {
	*x = GroupsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupsReq) ProtoMessage() {}

func (x *GroupsReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupsReq.ProtoReflect.Descriptor instead.
func (*GroupsReq) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{3}
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{4}
}

func (x *StreamRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PushMsgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMsgReply) Reset() {
	*x = PushMsgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReply) ProtoMessage() {}

func (x *PushMsgReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReply.ProtoReflect.Descriptor instead.
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{5}
}

type BroadcastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BroadcastReply) Reset() {
	*x = BroadcastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastReply) ProtoMessage() {}

func (x *BroadcastReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastReply.ProtoReflect.Descriptor instead.
func (*BroadcastReply) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{6}
}

func (x *BroadcastReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BroadcastGroupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastGroupReply) Reset() {
	*x = BroadcastGroupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastGroupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastGroupReply) ProtoMessage() {}

func (x *BroadcastGroupReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastGroupReply.ProtoReflect.Descriptor instead.
func (*BroadcastGroupReply) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{7}
}

type GroupsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group map[string]bool `protobuf:"bytes,1,rep,name=group,proto3" json:"group,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GroupsReply) Reset() {
	*x = GroupsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupsReply) ProtoMessage() {}

func (x *GroupsReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupsReply.ProtoReflect.Descriptor instead.
func (*GroupsReply) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{8}
}

func (x *GroupsReply) GetGroup() map[string]bool {
	if x != nil {
		return x.Group
	}
	return nil
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Event:
	//
	//	*StreamResponse_ClientLogin
	//	*StreamResponse_ClientLogout
	//	*StreamResponse_ClientMessage
	//	*StreamResponse_ServerShutdown
	Event isStreamResponse_Event `protobuf_oneof:"event"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{9}
}

func (x *StreamResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *StreamResponse) GetClientLogin() *StreamResponse_Login {
	if x, ok := x.GetEvent().(*StreamResponse_ClientLogin); ok {
		return x.ClientLogin
	}
	return nil
}

func (x *StreamResponse) GetClientLogout() *StreamResponse_Logout {
	if x, ok := x.GetEvent().(*StreamResponse_ClientLogout); ok {
		return x.ClientLogout
	}
	return nil
}

func (x *StreamResponse) GetClientMessage() *StreamResponse_Message {
	if x, ok := x.GetEvent().(*StreamResponse_ClientMessage); ok {
		return x.ClientMessage
	}
	return nil
}

func (x *StreamResponse) GetServerShutdown() *StreamResponse_Shutdown {
	if x, ok := x.GetEvent().(*StreamResponse_ServerShutdown); ok {
		return x.ServerShutdown
	}
	return nil
}

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_ClientLogin struct {
	ClientLogin *StreamResponse_Login `protobuf:"bytes,2,opt,name=client_login,json=clientLogin,proto3,oneof"`
}

type StreamResponse_ClientLogout struct {
	ClientLogout *StreamResponse_Logout `protobuf:"bytes,3,opt,name=client_logout,json=clientLogout,proto3,oneof"`
}

type StreamResponse_ClientMessage struct {
	ClientMessage *StreamResponse_Message `protobuf:"bytes,4,opt,name=client_message,json=clientMessage,proto3,oneof"`
}

type StreamResponse_ServerShutdown struct {
	ServerShutdown *StreamResponse_Shutdown `protobuf:"bytes,5,opt,name=server_shutdown,json=serverShutdown,proto3,oneof"`
}

func (*StreamResponse_ClientLogin) isStreamResponse_Event() {}

func (*StreamResponse_ClientLogout) isStreamResponse_Event() {}

func (*StreamResponse_ClientMessage) isStreamResponse_Event() {}

func (*StreamResponse_ServerShutdown) isStreamResponse_Event() {}

type StreamResponse_Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StreamResponse_Login) Reset() {
	*x = StreamResponse_Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse_Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse_Login) ProtoMessage() {}

func (x *StreamResponse_Login) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse_Login.ProtoReflect.Descriptor instead.
func (*StreamResponse_Login) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{9, 0}
}

func (x *StreamResponse_Login) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StreamResponse_Logout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *StreamResponse_Logout) Reset() {
	*x = StreamResponse_Logout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse_Logout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse_Logout) ProtoMessage() {}

func (x *StreamResponse_Logout) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse_Logout.ProtoReflect.Descriptor instead.
func (*StreamResponse_Logout) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{9, 1}
}

func (x *StreamResponse_Logout) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StreamResponse_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StreamResponse_Message) Reset() {
	*x = StreamResponse_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse_Message) ProtoMessage() {}

func (x *StreamResponse_Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse_Message.ProtoReflect.Descriptor instead.
func (*StreamResponse_Message) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{9, 2}
}

func (x *StreamResponse_Message) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamResponse_Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StreamResponse_Shutdown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamResponse_Shutdown) Reset() {
	*x = StreamResponse_Shutdown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_comet_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse_Shutdown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse_Shutdown) ProtoMessage() {}

func (x *StreamResponse_Shutdown) ProtoReflect() protoreflect.Message {
	mi := &file_proto_comet_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse_Shutdown.ProtoReflect.Descriptor instead.
func (*StreamResponse_Shutdown) Descriptor() ([]byte, []int) {
	return file_proto_comet_proto_rawDescGZIP(), []int{9, 3}
}

var File_proto_comet_proto protoreflect.FileDescriptor

var file_proto_comet_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x28, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x47, 0x0a, 0x11, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71,
	0x22, 0x29, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2a, 0x0a, 0x0e, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x79,
	0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a,
	0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a,
	0x38, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe1, 0x03, 0x0a, 0x0e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x40, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x43, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f,
	0x77, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x1a, 0x1b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x1a, 0x1c, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x37, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x9a, 0x02,
	0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2f,
	0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x35, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x06,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x6e, 0x72, 0x79, 0x73, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x79, 0x32, 0x30, 0x32, 0x32, 0x67, 0x6f,
	0x2f, 0x63, 0x68, 0x33, 0x37, 0x2f, 0x7a, 0x79, 0x6a, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_comet_proto_rawDescOnce sync.Once
	file_proto_comet_proto_rawDescData = file_proto_comet_proto_rawDesc
)

func file_proto_comet_proto_rawDescGZIP() []byte {
	file_proto_comet_proto_rawDescOnce.Do(func() {
		file_proto_comet_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_comet_proto_rawDescData)
	})
	return file_proto_comet_proto_rawDescData
}

var file_proto_comet_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_comet_proto_goTypes = []interface{}{
	(*PushMsgReq)(nil),              // 0: v1.PushMsgReq
	(*BroadcastReq)(nil),            // 1: v1.BroadcastReq
	(*BroadcastGroupReq)(nil),       // 2: v1.BroadcastGroupReq
	(*GroupsReq)(nil),               // 3: v1.GroupsReq
	(*StreamRequest)(nil),           // 4: v1.StreamRequest
	(*PushMsgReply)(nil),            // 5: v1.PushMsgReply
	(*BroadcastReply)(nil),          // 6: v1.BroadcastReply
	(*BroadcastGroupReply)(nil),     // 7: v1.BroadcastGroupReply
	(*GroupsReply)(nil),             // 8: v1.GroupsReply
	(*StreamResponse)(nil),          // 9: v1.StreamResponse
	nil,                             // 10: v1.GroupsReply.GroupEntry
	(*StreamResponse_Login)(nil),    // 11: v1.StreamResponse.Login
	(*StreamResponse_Logout)(nil),   // 12: v1.StreamResponse.Logout
	(*StreamResponse_Message)(nil),  // 13: v1.StreamResponse.Message
	(*StreamResponse_Shutdown)(nil), // 14: v1.StreamResponse.Shutdown
	(*timestamppb.Timestamp)(nil),   // 15: google.protobuf.Timestamp
}
var file_proto_comet_proto_depIdxs = []int32{
	10, // 0: v1.GroupsReply.group:type_name -> v1.GroupsReply.GroupEntry
	15, // 1: v1.StreamResponse.timestamp:type_name -> google.protobuf.Timestamp
	11, // 2: v1.StreamResponse.client_login:type_name -> v1.StreamResponse.Login
	12, // 3: v1.StreamResponse.client_logout:type_name -> v1.StreamResponse.Logout
	13, // 4: v1.StreamResponse.client_message:type_name -> v1.StreamResponse.Message
	14, // 5: v1.StreamResponse.server_shutdown:type_name -> v1.StreamResponse.Shutdown
	4,  // 6: v1.Comet.Stream:input_type -> v1.StreamRequest
	0,  // 7: v1.Comet.PushMsg:input_type -> v1.PushMsgReq
	1,  // 8: v1.Comet.Broadcast:input_type -> v1.BroadcastReq
	2,  // 9: v1.Comet.BroadcastGroup:input_type -> v1.BroadcastGroupReq
	3,  // 10: v1.Comet.Groups:input_type -> v1.GroupsReq
	9,  // 11: v1.Comet.Stream:output_type -> v1.StreamResponse
	5,  // 12: v1.Comet.PushMsg:output_type -> v1.PushMsgReply
	6,  // 13: v1.Comet.Broadcast:output_type -> v1.BroadcastReply
	7,  // 14: v1.Comet.BroadcastGroup:output_type -> v1.BroadcastGroupReply
	8,  // 15: v1.Comet.Groups:output_type -> v1.GroupsReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_comet_proto_init() }
func file_proto_comet_proto_init() {
	if File_proto_comet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_comet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReq); i {
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
		file_proto_comet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReq); i {
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
		file_proto_comet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastGroupReq); i {
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
		file_proto_comet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupsReq); i {
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
		file_proto_comet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_proto_comet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReply); i {
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
		file_proto_comet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastReply); i {
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
		file_proto_comet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastGroupReply); i {
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
		file_proto_comet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupsReply); i {
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
		file_proto_comet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
		file_proto_comet_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse_Login); i {
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
		file_proto_comet_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse_Logout); i {
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
		file_proto_comet_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse_Message); i {
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
		file_proto_comet_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse_Shutdown); i {
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
	file_proto_comet_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*StreamResponse_ClientLogin)(nil),
		(*StreamResponse_ClientLogout)(nil),
		(*StreamResponse_ClientMessage)(nil),
		(*StreamResponse_ServerShutdown)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_comet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_comet_proto_goTypes,
		DependencyIndexes: file_proto_comet_proto_depIdxs,
		MessageInfos:      file_proto_comet_proto_msgTypes,
	}.Build()
	File_proto_comet_proto = out.File
	file_proto_comet_proto_rawDesc = nil
	file_proto_comet_proto_goTypes = nil
	file_proto_comet_proto_depIdxs = nil
}

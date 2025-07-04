// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: cmd/cmd_chat.proto

package pb

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

type PrivateChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid uint32 `protobuf:"varint,7,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	// Types that are assignable to Content:
	//
	//	*PrivateChatReq_Text
	//	*PrivateChatReq_Icon
	Content isPrivateChatReq_Content `protobuf_oneof:"content"`
}

func (x *PrivateChatReq) Reset() {
	*x = PrivateChatReq{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateChatReq) ProtoMessage() {}

func (x *PrivateChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateChatReq.ProtoReflect.Descriptor instead.
func (*PrivateChatReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateChatReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (m *PrivateChatReq) GetContent() isPrivateChatReq_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *PrivateChatReq) GetText() string {
	if x, ok := x.GetContent().(*PrivateChatReq_Text); ok {
		return x.Text
	}
	return ""
}

func (x *PrivateChatReq) GetIcon() uint32 {
	if x, ok := x.GetContent().(*PrivateChatReq_Icon); ok {
		return x.Icon
	}
	return 0
}

type isPrivateChatReq_Content interface {
	isPrivateChatReq_Content()
}

type PrivateChatReq_Text struct {
	Text string `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type PrivateChatReq_Icon struct {
	Icon uint32 `protobuf:"varint,4,opt,name=icon,proto3,oneof"`
}

func (*PrivateChatReq_Text) isPrivateChatReq_Content() {}

func (*PrivateChatReq_Icon) isPrivateChatReq_Content() {}

type PrivateChatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatForbiddenEndtime uint32 `protobuf:"varint,12,opt,name=chat_forbidden_endtime,json=chatForbiddenEndtime,proto3" json:"chat_forbidden_endtime,omitempty"`
	Retcode              int32  `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PrivateChatRsp) Reset() {
	*x = PrivateChatRsp{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateChatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateChatRsp) ProtoMessage() {}

func (x *PrivateChatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateChatRsp.ProtoReflect.Descriptor instead.
func (*PrivateChatRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateChatRsp) GetChatForbiddenEndtime() uint32 {
	if x != nil {
		return x.ChatForbiddenEndtime
	}
	return 0
}

func (x *PrivateChatRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type PrivateChatNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo *ChatInfo `protobuf:"bytes,7,opt,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
}

func (x *PrivateChatNotify) Reset() {
	*x = PrivateChatNotify{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateChatNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateChatNotify) ProtoMessage() {}

func (x *PrivateChatNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateChatNotify.ProtoReflect.Descriptor instead.
func (*PrivateChatNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateChatNotify) GetChatInfo() *ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

type PullPrivateChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid    uint32 `protobuf:"varint,5,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	PullNum      uint32 `protobuf:"varint,7,opt,name=pull_num,json=pullNum,proto3" json:"pull_num,omitempty"`
	FromSequence uint32 `protobuf:"varint,12,opt,name=from_sequence,json=fromSequence,proto3" json:"from_sequence,omitempty"`
}

func (x *PullPrivateChatReq) Reset() {
	*x = PullPrivateChatReq{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullPrivateChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullPrivateChatReq) ProtoMessage() {}

func (x *PullPrivateChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullPrivateChatReq.ProtoReflect.Descriptor instead.
func (*PullPrivateChatReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{3}
}

func (x *PullPrivateChatReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *PullPrivateChatReq) GetPullNum() uint32 {
	if x != nil {
		return x.PullNum
	}
	return 0
}

func (x *PullPrivateChatReq) GetFromSequence() uint32 {
	if x != nil {
		return x.FromSequence
	}
	return 0
}

type PullPrivateChatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo []*ChatInfo `protobuf:"bytes,15,rep,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
	Retcode  int32       `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PullPrivateChatRsp) Reset() {
	*x = PullPrivateChatRsp{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullPrivateChatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullPrivateChatRsp) ProtoMessage() {}

func (x *PullPrivateChatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullPrivateChatRsp.ProtoReflect.Descriptor instead.
func (*PullPrivateChatRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{4}
}

func (x *PullPrivateChatRsp) GetChatInfo() []*ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

func (x *PullPrivateChatRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type PullRecentChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PullNum       uint32 `protobuf:"varint,6,opt,name=pull_num,json=pullNum,proto3" json:"pull_num,omitempty"`
	BeginSequence uint32 `protobuf:"varint,15,opt,name=begin_sequence,json=beginSequence,proto3" json:"begin_sequence,omitempty"`
}

func (x *PullRecentChatReq) Reset() {
	*x = PullRecentChatReq{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullRecentChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRecentChatReq) ProtoMessage() {}

func (x *PullRecentChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRecentChatReq.ProtoReflect.Descriptor instead.
func (*PullRecentChatReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{5}
}

func (x *PullRecentChatReq) GetPullNum() uint32 {
	if x != nil {
		return x.PullNum
	}
	return 0
}

func (x *PullRecentChatReq) GetBeginSequence() uint32 {
	if x != nil {
		return x.BeginSequence
	}
	return 0
}

type PullRecentChatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatInfo []*ChatInfo `protobuf:"bytes,15,rep,name=chat_info,json=chatInfo,proto3" json:"chat_info,omitempty"`
	Retcode  int32       `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *PullRecentChatRsp) Reset() {
	*x = PullRecentChatRsp{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullRecentChatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullRecentChatRsp) ProtoMessage() {}

func (x *PullRecentChatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullRecentChatRsp.ProtoReflect.Descriptor instead.
func (*PullRecentChatRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{6}
}

func (x *PullRecentChatRsp) GetChatInfo() []*ChatInfo {
	if x != nil {
		return x.ChatInfo
	}
	return nil
}

func (x *PullRecentChatRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type ReadPrivateChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetUid uint32 `protobuf:"varint,1,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *ReadPrivateChatReq) Reset() {
	*x = ReadPrivateChatReq{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadPrivateChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPrivateChatReq) ProtoMessage() {}

func (x *ReadPrivateChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPrivateChatReq.ProtoReflect.Descriptor instead.
func (*ReadPrivateChatReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{7}
}

func (x *ReadPrivateChatReq) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

type ReadPrivateChatRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *ReadPrivateChatRsp) Reset() {
	*x = ReadPrivateChatRsp{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadPrivateChatRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPrivateChatRsp) ProtoMessage() {}

func (x *ReadPrivateChatRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPrivateChatRsp.ProtoReflect.Descriptor instead.
func (*ReadPrivateChatRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{8}
}

func (x *ReadPrivateChatRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type ChatChannelUpdateNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId   uint32           `protobuf:"varint,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	IsCreate    bool             `protobuf:"varint,15,opt,name=is_create,json=isCreate,proto3" json:"is_create,omitempty"`
	ChannelInfo *ChatChannelInfo `protobuf:"bytes,14,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
}

func (x *ChatChannelUpdateNotify) Reset() {
	*x = ChatChannelUpdateNotify{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatChannelUpdateNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatChannelUpdateNotify) ProtoMessage() {}

func (x *ChatChannelUpdateNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatChannelUpdateNotify.ProtoReflect.Descriptor instead.
func (*ChatChannelUpdateNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{9}
}

func (x *ChatChannelUpdateNotify) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *ChatChannelUpdateNotify) GetIsCreate() bool {
	if x != nil {
		return x.IsCreate
	}
	return false
}

func (x *ChatChannelUpdateNotify) GetChannelInfo() *ChatChannelInfo {
	if x != nil {
		return x.ChannelInfo
	}
	return nil
}

type ChatChannelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsShield  bool   `protobuf:"varint,15,opt,name=is_shield,json=isShield,proto3" json:"is_shield,omitempty"`
	ChannelId uint32 `protobuf:"varint,8,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *ChatChannelInfo) Reset() {
	*x = ChatChannelInfo{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatChannelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatChannelInfo) ProtoMessage() {}

func (x *ChatChannelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatChannelInfo.ProtoReflect.Descriptor instead.
func (*ChatChannelInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{10}
}

func (x *ChatChannelInfo) GetIsShield() bool {
	if x != nil {
		return x.IsShield
	}
	return false
}

func (x *ChatChannelInfo) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type ChatChannelDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelList     []uint32           `protobuf:"varint,3,rep,packed,name=channel_list,json=channelList,proto3" json:"channel_list,omitempty"`
	ChannelInfoList []*ChatChannelInfo `protobuf:"bytes,7,rep,name=channel_info_list,json=channelInfoList,proto3" json:"channel_info_list,omitempty"`
}

func (x *ChatChannelDataNotify) Reset() {
	*x = ChatChannelDataNotify{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatChannelDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatChannelDataNotify) ProtoMessage() {}

func (x *ChatChannelDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatChannelDataNotify.ProtoReflect.Descriptor instead.
func (*ChatChannelDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{11}
}

func (x *ChatChannelDataNotify) GetChannelList() []uint32 {
	if x != nil {
		return x.ChannelList
	}
	return nil
}

func (x *ChatChannelDataNotify) GetChannelInfoList() []*ChatChannelInfo {
	if x != nil {
		return x.ChannelInfoList
	}
	return nil
}

type ChatChannelShieldNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsShield  bool   `protobuf:"varint,5,opt,name=is_shield,json=isShield,proto3" json:"is_shield,omitempty"`
	ChannelId uint32 `protobuf:"varint,14,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *ChatChannelShieldNotify) Reset() {
	*x = ChatChannelShieldNotify{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatChannelShieldNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatChannelShieldNotify) ProtoMessage() {}

func (x *ChatChannelShieldNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatChannelShieldNotify.ProtoReflect.Descriptor instead.
func (*ChatChannelShieldNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{12}
}

func (x *ChatChannelShieldNotify) GetIsShield() bool {
	if x != nil {
		return x.IsShield
	}
	return false
}

func (x *ChatChannelShieldNotify) GetChannelId() uint32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type ChatChannelInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelInfo *ChatChannelInfo `protobuf:"bytes,2,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
}

func (x *ChatChannelInfoNotify) Reset() {
	*x = ChatChannelInfoNotify{}
	mi := &file_cmd_cmd_chat_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatChannelInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatChannelInfoNotify) ProtoMessage() {}

func (x *ChatChannelInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_chat_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatChannelInfoNotify.ProtoReflect.Descriptor instead.
func (*ChatChannelInfoNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_chat_proto_rawDescGZIP(), []int{13}
}

func (x *ChatChannelInfoNotify) GetChannelInfo() *ChatChannelInfo {
	if x != nil {
		return x.ChannelInfo
	}
	return nil
}

var File_cmd_cmd_chat_proto protoreflect.FileDescriptor

var file_cmd_cmd_chat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6d, 0x64,
	0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x66, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x68, 0x61, 0x74,
	0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x45, 0x6e, 0x64, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x41, 0x0a, 0x11, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x2c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x73, 0x0a,
	0x12, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x75, 0x6c, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x22, 0x5c, 0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x63, 0x68,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x55, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x75, 0x6c, 0x6c, 0x4e, 0x75, 0x6d,
	0x12, 0x25, 0x0a, 0x0e, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x53,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x5b, 0x0a, 0x11, 0x50, 0x75, 0x6c, 0x6c, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x33, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x61,
	0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x17, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x0f,
	0x43, 0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x15, 0x43,
	0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x55, 0x0a, 0x17, 0x43,
	0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x22, 0x52, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70,
	0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_chat_proto_rawDescOnce sync.Once
	file_cmd_cmd_chat_proto_rawDescData = file_cmd_cmd_chat_proto_rawDesc
)

func file_cmd_cmd_chat_proto_rawDescGZIP() []byte {
	file_cmd_cmd_chat_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_chat_proto_rawDescData)
	})
	return file_cmd_cmd_chat_proto_rawDescData
}

var file_cmd_cmd_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_cmd_cmd_chat_proto_goTypes = []any{
	(*PrivateChatReq)(nil),          // 0: proto.PrivateChatReq
	(*PrivateChatRsp)(nil),          // 1: proto.PrivateChatRsp
	(*PrivateChatNotify)(nil),       // 2: proto.PrivateChatNotify
	(*PullPrivateChatReq)(nil),      // 3: proto.PullPrivateChatReq
	(*PullPrivateChatRsp)(nil),      // 4: proto.PullPrivateChatRsp
	(*PullRecentChatReq)(nil),       // 5: proto.PullRecentChatReq
	(*PullRecentChatRsp)(nil),       // 6: proto.PullRecentChatRsp
	(*ReadPrivateChatReq)(nil),      // 7: proto.ReadPrivateChatReq
	(*ReadPrivateChatRsp)(nil),      // 8: proto.ReadPrivateChatRsp
	(*ChatChannelUpdateNotify)(nil), // 9: proto.ChatChannelUpdateNotify
	(*ChatChannelInfo)(nil),         // 10: proto.ChatChannelInfo
	(*ChatChannelDataNotify)(nil),   // 11: proto.ChatChannelDataNotify
	(*ChatChannelShieldNotify)(nil), // 12: proto.ChatChannelShieldNotify
	(*ChatChannelInfoNotify)(nil),   // 13: proto.ChatChannelInfoNotify
	(*ChatInfo)(nil),                // 14: proto.ChatInfo
}
var file_cmd_cmd_chat_proto_depIdxs = []int32{
	14, // 0: proto.PrivateChatNotify.chat_info:type_name -> proto.ChatInfo
	14, // 1: proto.PullPrivateChatRsp.chat_info:type_name -> proto.ChatInfo
	14, // 2: proto.PullRecentChatRsp.chat_info:type_name -> proto.ChatInfo
	10, // 3: proto.ChatChannelUpdateNotify.channel_info:type_name -> proto.ChatChannelInfo
	10, // 4: proto.ChatChannelDataNotify.channel_info_list:type_name -> proto.ChatChannelInfo
	10, // 5: proto.ChatChannelInfoNotify.channel_info:type_name -> proto.ChatChannelInfo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_cmd_chat_proto_init() }
func file_cmd_cmd_chat_proto_init() {
	if File_cmd_cmd_chat_proto != nil {
		return
	}
	file_cmd_cmd_scene_proto_init()
	file_cmd_cmd_chat_proto_msgTypes[0].OneofWrappers = []any{
		(*PrivateChatReq_Text)(nil),
		(*PrivateChatReq_Icon)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_cmd_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_chat_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_chat_proto_depIdxs,
		MessageInfos:      file_cmd_cmd_chat_proto_msgTypes,
	}.Build()
	File_cmd_cmd_chat_proto = out.File
	file_cmd_cmd_chat_proto_rawDesc = nil
	file_cmd_cmd_chat_proto_goTypes = nil
	file_cmd_cmd_chat_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: server_only/cmd_mp.server.proto

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

type UpdateMpStatusNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                  uint32            `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	IsOnlyMpWithPsPlayer bool              `protobuf:"varint,2,opt,name=is_only_mp_with_ps_player,json=isOnlyMpWithPsPlayer,proto3" json:"is_only_mp_with_ps_player,omitempty"`
	PlatformType         PlatformType      `protobuf:"varint,3,opt,name=platform_type,json=platformType,proto3,enum=proto.PlatformType" json:"platform_type,omitempty"`
	OnlinePlayerInfo     *OnlinePlayerInfo `protobuf:"bytes,4,opt,name=online_player_info,json=onlinePlayerInfo,proto3" json:"online_player_info,omitempty"`
}

func (x *UpdateMpStatusNotify) Reset() {
	*x = UpdateMpStatusNotify{}
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateMpStatusNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMpStatusNotify) ProtoMessage() {}

func (x *UpdateMpStatusNotify) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMpStatusNotify.ProtoReflect.Descriptor instead.
func (*UpdateMpStatusNotify) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_mp_server_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateMpStatusNotify) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UpdateMpStatusNotify) GetIsOnlyMpWithPsPlayer() bool {
	if x != nil {
		return x.IsOnlyMpWithPsPlayer
	}
	return false
}

func (x *UpdateMpStatusNotify) GetPlatformType() PlatformType {
	if x != nil {
		return x.PlatformType
	}
	return PlatformType_EDITOR
}

func (x *UpdateMpStatusNotify) GetOnlinePlayerInfo() *OnlinePlayerInfo {
	if x != nil {
		return x.OnlinePlayerInfo
	}
	return nil
}

type DelMpStatusNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DelMpStatusNotify) Reset() {
	*x = DelMpStatusNotify{}
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelMpStatusNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelMpStatusNotify) ProtoMessage() {}

func (x *DelMpStatusNotify) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelMpStatusNotify.ProtoReflect.Descriptor instead.
func (*DelMpStatusNotify) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_mp_server_proto_rawDescGZIP(), []int{1}
}

func (x *DelMpStatusNotify) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetPlayerMpStatusListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPlayerMpStatusListReq) Reset() {
	*x = GetPlayerMpStatusListReq{}
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerMpStatusListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerMpStatusListReq) ProtoMessage() {}

func (x *GetPlayerMpStatusListReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerMpStatusListReq.ProtoReflect.Descriptor instead.
func (*GetPlayerMpStatusListReq) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_mp_server_proto_rawDescGZIP(), []int{2}
}

type GetPlayerMpStatusListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32               `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PlayerList []*OnlinePlayerInfo `protobuf:"bytes,2,rep,name=player_list,json=playerList,proto3" json:"player_list,omitempty"`
}

func (x *GetPlayerMpStatusListRsp) Reset() {
	*x = GetPlayerMpStatusListRsp{}
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerMpStatusListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerMpStatusListRsp) ProtoMessage() {}

func (x *GetPlayerMpStatusListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerMpStatusListRsp.ProtoReflect.Descriptor instead.
func (*GetPlayerMpStatusListRsp) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_mp_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerMpStatusListRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPlayerMpStatusListRsp) GetPlayerList() []*OnlinePlayerInfo {
	if x != nil {
		return x.PlayerList
	}
	return nil
}

type GetPlayerMpStatusInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsOnlineId bool `protobuf:"varint,1,opt,name=is_online_id,json=isOnlineId,proto3" json:"is_online_id,omitempty"`
	// Types that are assignable to PlayerId:
	//
	//	*GetPlayerMpStatusInfoReq_TargetUid
	//	*GetPlayerMpStatusInfoReq_OnlineId
	//	*GetPlayerMpStatusInfoReq_PsnId
	PlayerId isGetPlayerMpStatusInfoReq_PlayerId `protobuf_oneof:"player_id"`
}

func (x *GetPlayerMpStatusInfoReq) Reset() {
	*x = GetPlayerMpStatusInfoReq{}
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerMpStatusInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerMpStatusInfoReq) ProtoMessage() {}

func (x *GetPlayerMpStatusInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerMpStatusInfoReq.ProtoReflect.Descriptor instead.
func (*GetPlayerMpStatusInfoReq) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_mp_server_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerMpStatusInfoReq) GetIsOnlineId() bool {
	if x != nil {
		return x.IsOnlineId
	}
	return false
}

func (m *GetPlayerMpStatusInfoReq) GetPlayerId() isGetPlayerMpStatusInfoReq_PlayerId {
	if m != nil {
		return m.PlayerId
	}
	return nil
}

func (x *GetPlayerMpStatusInfoReq) GetTargetUid() uint32 {
	if x, ok := x.GetPlayerId().(*GetPlayerMpStatusInfoReq_TargetUid); ok {
		return x.TargetUid
	}
	return 0
}

func (x *GetPlayerMpStatusInfoReq) GetOnlineId() string {
	if x, ok := x.GetPlayerId().(*GetPlayerMpStatusInfoReq_OnlineId); ok {
		return x.OnlineId
	}
	return ""
}

func (x *GetPlayerMpStatusInfoReq) GetPsnId() string {
	if x, ok := x.GetPlayerId().(*GetPlayerMpStatusInfoReq_PsnId); ok {
		return x.PsnId
	}
	return ""
}

type isGetPlayerMpStatusInfoReq_PlayerId interface {
	isGetPlayerMpStatusInfoReq_PlayerId()
}

type GetPlayerMpStatusInfoReq_TargetUid struct {
	TargetUid uint32 `protobuf:"varint,2,opt,name=target_uid,json=targetUid,proto3,oneof"`
}

type GetPlayerMpStatusInfoReq_OnlineId struct {
	OnlineId string `protobuf:"bytes,3,opt,name=online_id,json=onlineId,proto3,oneof"`
}

type GetPlayerMpStatusInfoReq_PsnId struct {
	PsnId string `protobuf:"bytes,4,opt,name=psn_id,json=psnId,proto3,oneof"`
}

func (*GetPlayerMpStatusInfoReq_TargetUid) isGetPlayerMpStatusInfoReq_PlayerId() {}

func (*GetPlayerMpStatusInfoReq_OnlineId) isGetPlayerMpStatusInfoReq_PlayerId() {}

func (*GetPlayerMpStatusInfoReq_PsnId) isGetPlayerMpStatusInfoReq_PlayerId() {}

type GetPlayerMpStatusInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode          int32             `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TargetUid        uint32            `protobuf:"varint,2,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
	TargetPlayerInfo *OnlinePlayerInfo `protobuf:"bytes,3,opt,name=target_player_info,json=targetPlayerInfo,proto3" json:"target_player_info,omitempty"`
}

func (x *GetPlayerMpStatusInfoRsp) Reset() {
	*x = GetPlayerMpStatusInfoRsp{}
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPlayerMpStatusInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerMpStatusInfoRsp) ProtoMessage() {}

func (x *GetPlayerMpStatusInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_server_only_cmd_mp_server_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerMpStatusInfoRsp.ProtoReflect.Descriptor instead.
func (*GetPlayerMpStatusInfoRsp) Descriptor() ([]byte, []int) {
	return file_server_only_cmd_mp_server_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlayerMpStatusInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetPlayerMpStatusInfoRsp) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

func (x *GetPlayerMpStatusInfoRsp) GetTargetPlayerInfo() *OnlinePlayerInfo {
	if x != nil {
		return x.TargetPlayerInfo
	}
	return nil
}

var File_server_only_cmd_mp_server_proto protoreflect.FileDescriptor

var file_server_only_cmd_mp_server_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x2f, 0x63, 0x6d,
	0x64, 0x5f, 0x6d, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x37, 0x0a, 0x19, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x6d, 0x70, 0x5f,
	0x77, 0x69, 0x74, 0x68, 0x5f, 0x70, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x79, 0x4d, 0x70, 0x57, 0x69,
	0x74, 0x68, 0x50, 0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x12, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x25, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x4d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x6e,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa2,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4d, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1d,
	0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x06, 0x70, 0x73, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x70, 0x73, 0x6e, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4d, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x12, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34,
	0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_server_only_cmd_mp_server_proto_rawDescOnce sync.Once
	file_server_only_cmd_mp_server_proto_rawDescData = file_server_only_cmd_mp_server_proto_rawDesc
)

func file_server_only_cmd_mp_server_proto_rawDescGZIP() []byte {
	file_server_only_cmd_mp_server_proto_rawDescOnce.Do(func() {
		file_server_only_cmd_mp_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_only_cmd_mp_server_proto_rawDescData)
	})
	return file_server_only_cmd_mp_server_proto_rawDescData
}

var file_server_only_cmd_mp_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_server_only_cmd_mp_server_proto_goTypes = []any{
	(*UpdateMpStatusNotify)(nil),     // 0: proto.UpdateMpStatusNotify
	(*DelMpStatusNotify)(nil),        // 1: proto.DelMpStatusNotify
	(*GetPlayerMpStatusListReq)(nil), // 2: proto.GetPlayerMpStatusListReq
	(*GetPlayerMpStatusListRsp)(nil), // 3: proto.GetPlayerMpStatusListRsp
	(*GetPlayerMpStatusInfoReq)(nil), // 4: proto.GetPlayerMpStatusInfoReq
	(*GetPlayerMpStatusInfoRsp)(nil), // 5: proto.GetPlayerMpStatusInfoRsp
	(PlatformType)(0),                // 6: proto.PlatformType
	(*OnlinePlayerInfo)(nil),         // 7: proto.OnlinePlayerInfo
}
var file_server_only_cmd_mp_server_proto_depIdxs = []int32{
	6, // 0: proto.UpdateMpStatusNotify.platform_type:type_name -> proto.PlatformType
	7, // 1: proto.UpdateMpStatusNotify.online_player_info:type_name -> proto.OnlinePlayerInfo
	7, // 2: proto.GetPlayerMpStatusListRsp.player_list:type_name -> proto.OnlinePlayerInfo
	7, // 3: proto.GetPlayerMpStatusInfoRsp.target_player_info:type_name -> proto.OnlinePlayerInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_server_only_cmd_mp_server_proto_init() }
func file_server_only_cmd_mp_server_proto_init() {
	if File_server_only_cmd_mp_server_proto != nil {
		return
	}
	file_define_proto_init()
	file_server_only_cmd_mp_server_proto_msgTypes[4].OneofWrappers = []any{
		(*GetPlayerMpStatusInfoReq_TargetUid)(nil),
		(*GetPlayerMpStatusInfoReq_OnlineId)(nil),
		(*GetPlayerMpStatusInfoReq_PsnId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_only_cmd_mp_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_server_only_cmd_mp_server_proto_goTypes,
		DependencyIndexes: file_server_only_cmd_mp_server_proto_depIdxs,
		MessageInfos:      file_server_only_cmd_mp_server_proto_msgTypes,
	}.Build()
	File_server_only_cmd_mp_server_proto = out.File
	file_server_only_cmd_mp_server_proto_rawDesc = nil
	file_server_only_cmd_mp_server_proto_goTypes = nil
	file_server_only_cmd_mp_server_proto_depIdxs = nil
}

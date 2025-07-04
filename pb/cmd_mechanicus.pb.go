// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: cmd/cmd_mechanicus.proto

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

type MechanicusInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GearLevelPairList        []*Uint32Pair `protobuf:"bytes,14,rep,name=gear_level_pair_list,json=gearLevelPairList,proto3" json:"gear_level_pair_list,omitempty"`
	OpenSequenceIdList       []uint32      `protobuf:"varint,7,rep,packed,name=open_sequence_id_list,json=openSequenceIdList,proto3" json:"open_sequence_id_list,omitempty"`
	Coin                     uint32        `protobuf:"varint,8,opt,name=coin,proto3" json:"coin,omitempty"`
	PunishOverTime           uint32        `protobuf:"varint,12,opt,name=punish_over_time,json=punishOverTime,proto3" json:"punish_over_time,omitempty"`
	MechanicusId             uint32        `protobuf:"varint,10,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	FinishDifficultLevelList []uint32      `protobuf:"varint,13,rep,packed,name=finish_difficult_level_list,json=finishDifficultLevelList,proto3" json:"finish_difficult_level_list,omitempty"`
	IsFinishTeachDungeon     bool          `protobuf:"varint,4,opt,name=is_finish_teach_dungeon,json=isFinishTeachDungeon,proto3" json:"is_finish_teach_dungeon,omitempty"`
}

func (x *MechanicusInfo) Reset() {
	*x = MechanicusInfo{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusInfo) ProtoMessage() {}

func (x *MechanicusInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusInfo.ProtoReflect.Descriptor instead.
func (*MechanicusInfo) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{0}
}

func (x *MechanicusInfo) GetGearLevelPairList() []*Uint32Pair {
	if x != nil {
		return x.GearLevelPairList
	}
	return nil
}

func (x *MechanicusInfo) GetOpenSequenceIdList() []uint32 {
	if x != nil {
		return x.OpenSequenceIdList
	}
	return nil
}

func (x *MechanicusInfo) GetCoin() uint32 {
	if x != nil {
		return x.Coin
	}
	return 0
}

func (x *MechanicusInfo) GetPunishOverTime() uint32 {
	if x != nil {
		return x.PunishOverTime
	}
	return 0
}

func (x *MechanicusInfo) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *MechanicusInfo) GetFinishDifficultLevelList() []uint32 {
	if x != nil {
		return x.FinishDifficultLevelList
	}
	return nil
}

func (x *MechanicusInfo) GetIsFinishTeachDungeon() bool {
	if x != nil {
		return x.IsFinishTeachDungeon
	}
	return false
}

type GetMechanicusInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMechanicusInfoReq) Reset() {
	*x = GetMechanicusInfoReq{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMechanicusInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMechanicusInfoReq) ProtoMessage() {}

func (x *GetMechanicusInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMechanicusInfoReq.ProtoReflect.Descriptor instead.
func (*GetMechanicusInfoReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{1}
}

type GetMechanicusInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode        int32           `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MechanicusInfo *MechanicusInfo `protobuf:"bytes,15,opt,name=mechanicus_info,json=mechanicusInfo,proto3" json:"mechanicus_info,omitempty"`
}

func (x *GetMechanicusInfoRsp) Reset() {
	*x = GetMechanicusInfoRsp{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMechanicusInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMechanicusInfoRsp) ProtoMessage() {}

func (x *GetMechanicusInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMechanicusInfoRsp.ProtoReflect.Descriptor instead.
func (*GetMechanicusInfoRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{2}
}

func (x *GetMechanicusInfoRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMechanicusInfoRsp) GetMechanicusInfo() *MechanicusInfo {
	if x != nil {
		return x.MechanicusInfo
	}
	return nil
}

type MechanicusSequenceOpenNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusId uint32 `protobuf:"varint,8,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	SequenceId   uint32 `protobuf:"varint,7,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
}

func (x *MechanicusSequenceOpenNotify) Reset() {
	*x = MechanicusSequenceOpenNotify{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusSequenceOpenNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusSequenceOpenNotify) ProtoMessage() {}

func (x *MechanicusSequenceOpenNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusSequenceOpenNotify.ProtoReflect.Descriptor instead.
func (*MechanicusSequenceOpenNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{3}
}

func (x *MechanicusSequenceOpenNotify) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *MechanicusSequenceOpenNotify) GetSequenceId() uint32 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

type MechanicusCoinNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusId uint32 `protobuf:"varint,7,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	Coin         uint32 `protobuf:"varint,4,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (x *MechanicusCoinNotify) Reset() {
	*x = MechanicusCoinNotify{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusCoinNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusCoinNotify) ProtoMessage() {}

func (x *MechanicusCoinNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusCoinNotify.ProtoReflect.Descriptor instead.
func (*MechanicusCoinNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{4}
}

func (x *MechanicusCoinNotify) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *MechanicusCoinNotify) GetCoin() uint32 {
	if x != nil {
		return x.Coin
	}
	return 0
}

type MechanicusOpenNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusId uint32 `protobuf:"varint,2,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
}

func (x *MechanicusOpenNotify) Reset() {
	*x = MechanicusOpenNotify{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusOpenNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusOpenNotify) ProtoMessage() {}

func (x *MechanicusOpenNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusOpenNotify.ProtoReflect.Descriptor instead.
func (*MechanicusOpenNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{5}
}

func (x *MechanicusOpenNotify) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

type MechanicusCloseNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusId uint32 `protobuf:"varint,6,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
}

func (x *MechanicusCloseNotify) Reset() {
	*x = MechanicusCloseNotify{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusCloseNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusCloseNotify) ProtoMessage() {}

func (x *MechanicusCloseNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusCloseNotify.ProtoReflect.Descriptor instead.
func (*MechanicusCloseNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{6}
}

func (x *MechanicusCloseNotify) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

type MechanicusUnlockGearReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MechanicusId uint32 `protobuf:"varint,7,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	GearId       uint32 `protobuf:"varint,6,opt,name=gear_id,json=gearId,proto3" json:"gear_id,omitempty"`
}

func (x *MechanicusUnlockGearReq) Reset() {
	*x = MechanicusUnlockGearReq{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusUnlockGearReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusUnlockGearReq) ProtoMessage() {}

func (x *MechanicusUnlockGearReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusUnlockGearReq.ProtoReflect.Descriptor instead.
func (*MechanicusUnlockGearReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{7}
}

func (x *MechanicusUnlockGearReq) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *MechanicusUnlockGearReq) GetGearId() uint32 {
	if x != nil {
		return x.GearId
	}
	return 0
}

type MechanicusUnlockGearRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode      int32  `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MechanicusId uint32 `protobuf:"varint,8,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	GearId       uint32 `protobuf:"varint,14,opt,name=gear_id,json=gearId,proto3" json:"gear_id,omitempty"`
}

func (x *MechanicusUnlockGearRsp) Reset() {
	*x = MechanicusUnlockGearRsp{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusUnlockGearRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusUnlockGearRsp) ProtoMessage() {}

func (x *MechanicusUnlockGearRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusUnlockGearRsp.ProtoReflect.Descriptor instead.
func (*MechanicusUnlockGearRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{8}
}

func (x *MechanicusUnlockGearRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MechanicusUnlockGearRsp) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *MechanicusUnlockGearRsp) GetGearId() uint32 {
	if x != nil {
		return x.GearId
	}
	return 0
}

type MechanicusLevelupGearReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GearId       uint32 `protobuf:"varint,14,opt,name=gear_id,json=gearId,proto3" json:"gear_id,omitempty"`
	MechanicusId uint32 `protobuf:"varint,12,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
}

func (x *MechanicusLevelupGearReq) Reset() {
	*x = MechanicusLevelupGearReq{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusLevelupGearReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusLevelupGearReq) ProtoMessage() {}

func (x *MechanicusLevelupGearReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusLevelupGearReq.ProtoReflect.Descriptor instead.
func (*MechanicusLevelupGearReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{9}
}

func (x *MechanicusLevelupGearReq) GetGearId() uint32 {
	if x != nil {
		return x.GearId
	}
	return 0
}

func (x *MechanicusLevelupGearReq) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

type MechanicusLevelupGearRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GearId         uint32 `protobuf:"varint,7,opt,name=gear_id,json=gearId,proto3" json:"gear_id,omitempty"`
	MechanicusId   uint32 `protobuf:"varint,2,opt,name=mechanicus_id,json=mechanicusId,proto3" json:"mechanicus_id,omitempty"`
	AfterGearLevel uint32 `protobuf:"varint,12,opt,name=after_gear_level,json=afterGearLevel,proto3" json:"after_gear_level,omitempty"`
	Retcode        int32  `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MechanicusLevelupGearRsp) Reset() {
	*x = MechanicusLevelupGearRsp{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusLevelupGearRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusLevelupGearRsp) ProtoMessage() {}

func (x *MechanicusLevelupGearRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusLevelupGearRsp.ProtoReflect.Descriptor instead.
func (*MechanicusLevelupGearRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{10}
}

func (x *MechanicusLevelupGearRsp) GetGearId() uint32 {
	if x != nil {
		return x.GearId
	}
	return 0
}

func (x *MechanicusLevelupGearRsp) GetMechanicusId() uint32 {
	if x != nil {
		return x.MechanicusId
	}
	return 0
}

func (x *MechanicusLevelupGearRsp) GetAfterGearLevel() uint32 {
	if x != nil {
		return x.AfterGearLevel
	}
	return 0
}

func (x *MechanicusLevelupGearRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type EnterMechanicusDungeonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DifficultLevel uint32 `protobuf:"varint,7,opt,name=difficult_level,json=difficultLevel,proto3" json:"difficult_level,omitempty"`
}

func (x *EnterMechanicusDungeonReq) Reset() {
	*x = EnterMechanicusDungeonReq{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnterMechanicusDungeonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterMechanicusDungeonReq) ProtoMessage() {}

func (x *EnterMechanicusDungeonReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterMechanicusDungeonReq.ProtoReflect.Descriptor instead.
func (*EnterMechanicusDungeonReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{11}
}

func (x *EnterMechanicusDungeonReq) GetDifficultLevel() uint32 {
	if x != nil {
		return x.DifficultLevel
	}
	return 0
}

type EnterMechanicusDungeonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WrongUid       uint32 `protobuf:"varint,12,opt,name=wrong_uid,json=wrongUid,proto3" json:"wrong_uid,omitempty"`
	DifficultLevel uint32 `protobuf:"varint,13,opt,name=difficult_level,json=difficultLevel,proto3" json:"difficult_level,omitempty"`
	Retcode        int32  `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DungeonId      uint32 `protobuf:"varint,11,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
}

func (x *EnterMechanicusDungeonRsp) Reset() {
	*x = EnterMechanicusDungeonRsp{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnterMechanicusDungeonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterMechanicusDungeonRsp) ProtoMessage() {}

func (x *EnterMechanicusDungeonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterMechanicusDungeonRsp.ProtoReflect.Descriptor instead.
func (*EnterMechanicusDungeonRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{12}
}

func (x *EnterMechanicusDungeonRsp) GetWrongUid() uint32 {
	if x != nil {
		return x.WrongUid
	}
	return 0
}

func (x *EnterMechanicusDungeonRsp) GetDifficultLevel() uint32 {
	if x != nil {
		return x.DifficultLevel
	}
	return 0
}

func (x *EnterMechanicusDungeonRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EnterMechanicusDungeonRsp) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

type MechanicusCandidateTeamCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DifficultLevel uint32 `protobuf:"varint,6,opt,name=difficult_level,json=difficultLevel,proto3" json:"difficult_level,omitempty"`
}

func (x *MechanicusCandidateTeamCreateReq) Reset() {
	*x = MechanicusCandidateTeamCreateReq{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusCandidateTeamCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusCandidateTeamCreateReq) ProtoMessage() {}

func (x *MechanicusCandidateTeamCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusCandidateTeamCreateReq.ProtoReflect.Descriptor instead.
func (*MechanicusCandidateTeamCreateReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{13}
}

func (x *MechanicusCandidateTeamCreateReq) GetDifficultLevel() uint32 {
	if x != nil {
		return x.DifficultLevel
	}
	return 0
}

type MechanicusCandidateTeamCreateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DungeonId      uint32 `protobuf:"varint,1,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	Retcode        int32  `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DifficultLevel uint32 `protobuf:"varint,10,opt,name=difficult_level,json=difficultLevel,proto3" json:"difficult_level,omitempty"`
}

func (x *MechanicusCandidateTeamCreateRsp) Reset() {
	*x = MechanicusCandidateTeamCreateRsp{}
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MechanicusCandidateTeamCreateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MechanicusCandidateTeamCreateRsp) ProtoMessage() {}

func (x *MechanicusCandidateTeamCreateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mechanicus_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MechanicusCandidateTeamCreateRsp.ProtoReflect.Descriptor instead.
func (*MechanicusCandidateTeamCreateRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mechanicus_proto_rawDescGZIP(), []int{14}
}

func (x *MechanicusCandidateTeamCreateRsp) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *MechanicusCandidateTeamCreateRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MechanicusCandidateTeamCreateRsp) GetDifficultLevel() uint32 {
	if x != nil {
		return x.DifficultLevel
	}
	return 0
}

var File_cmd_cmd_mechanicus_proto protoreflect.FileDescriptor

var file_cmd_cmd_mechanicus_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe0, 0x02, 0x0a, 0x0e, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x42, 0x0a, 0x14, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x5f, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x11, 0x67, 0x65, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x50, 0x61,
	0x69, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x6f, 0x70, 0x65, 0x6e, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0x28, 0x0a,
	0x10, 0x70, 0x75, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x70, 0x75, 0x6e, 0x69, 0x73, 0x68, 0x4f,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x1b,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x18, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75,
	0x6c, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x17, 0x69,
	0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x5f, 0x64,
	0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x65, 0x61, 0x63, 0x68, 0x44, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x70, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0f,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x6d, 0x65,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x64, 0x0a, 0x1c,
	0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x4f, 0x0a, 0x14, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73,
	0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63,
	0x6f, 0x69, 0x6e, 0x22, 0x3b, 0x0a, 0x14, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d,
	0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x15, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64, 0x22, 0x57,
	0x0a, 0x17, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x55, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x47, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x67, 0x65, 0x61, 0x72, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x17, 0x4d, 0x65, 0x63, 0x68, 0x61,
	0x6e, 0x69, 0x63, 0x75, 0x73, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x47, 0x65, 0x61, 0x72, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x61, 0x72, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x18, 0x4d, 0x65,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x47,
	0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x18, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69,
	0x63, 0x75, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x75, 0x70, 0x47, 0x65, 0x61, 0x72, 0x52, 0x73,
	0x70, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x67, 0x65, 0x61, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65,
	0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x61, 0x72, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x47, 0x65, 0x61, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x44, 0x0a, 0x19, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x63, 0x68,
	0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x9a, 0x01, 0x0a, 0x19, 0x45, 0x6e,
	0x74, 0x65, 0x72, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63, 0x75, 0x73, 0x44, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x72, 0x6f, 0x6e, 0x67,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x72, 0x6f, 0x6e,
	0x67, 0x55, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e,
	0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x20, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e,
	0x69, 0x63, 0x75, 0x73, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61,
	0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69,
	0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x84, 0x01, 0x0a, 0x20, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x63,
	0x75, 0x73, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75,
	0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x64, 0x69, 0x66, 0x66,
	0x69, 0x63, 0x75, 0x6c, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d,
	0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_mechanicus_proto_rawDescOnce sync.Once
	file_cmd_cmd_mechanicus_proto_rawDescData = file_cmd_cmd_mechanicus_proto_rawDesc
)

func file_cmd_cmd_mechanicus_proto_rawDescGZIP() []byte {
	file_cmd_cmd_mechanicus_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_mechanicus_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_mechanicus_proto_rawDescData)
	})
	return file_cmd_cmd_mechanicus_proto_rawDescData
}

var file_cmd_cmd_mechanicus_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_cmd_cmd_mechanicus_proto_goTypes = []any{
	(*MechanicusInfo)(nil),                   // 0: proto.MechanicusInfo
	(*GetMechanicusInfoReq)(nil),             // 1: proto.GetMechanicusInfoReq
	(*GetMechanicusInfoRsp)(nil),             // 2: proto.GetMechanicusInfoRsp
	(*MechanicusSequenceOpenNotify)(nil),     // 3: proto.MechanicusSequenceOpenNotify
	(*MechanicusCoinNotify)(nil),             // 4: proto.MechanicusCoinNotify
	(*MechanicusOpenNotify)(nil),             // 5: proto.MechanicusOpenNotify
	(*MechanicusCloseNotify)(nil),            // 6: proto.MechanicusCloseNotify
	(*MechanicusUnlockGearReq)(nil),          // 7: proto.MechanicusUnlockGearReq
	(*MechanicusUnlockGearRsp)(nil),          // 8: proto.MechanicusUnlockGearRsp
	(*MechanicusLevelupGearReq)(nil),         // 9: proto.MechanicusLevelupGearReq
	(*MechanicusLevelupGearRsp)(nil),         // 10: proto.MechanicusLevelupGearRsp
	(*EnterMechanicusDungeonReq)(nil),        // 11: proto.EnterMechanicusDungeonReq
	(*EnterMechanicusDungeonRsp)(nil),        // 12: proto.EnterMechanicusDungeonRsp
	(*MechanicusCandidateTeamCreateReq)(nil), // 13: proto.MechanicusCandidateTeamCreateReq
	(*MechanicusCandidateTeamCreateRsp)(nil), // 14: proto.MechanicusCandidateTeamCreateRsp
	(*Uint32Pair)(nil),                       // 15: proto.Uint32Pair
}
var file_cmd_cmd_mechanicus_proto_depIdxs = []int32{
	15, // 0: proto.MechanicusInfo.gear_level_pair_list:type_name -> proto.Uint32Pair
	0,  // 1: proto.GetMechanicusInfoRsp.mechanicus_info:type_name -> proto.MechanicusInfo
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_cmd_mechanicus_proto_init() }
func file_cmd_cmd_mechanicus_proto_init() {
	if File_cmd_cmd_mechanicus_proto != nil {
		return
	}
	file_define_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_cmd_mechanicus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_mechanicus_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_mechanicus_proto_depIdxs,
		MessageInfos:      file_cmd_cmd_mechanicus_proto_msgTypes,
	}.Build()
	File_cmd_cmd_mechanicus_proto = out.File
	file_cmd_cmd_mechanicus_proto_rawDesc = nil
	file_cmd_cmd_mechanicus_proto_goTypes = nil
	file_cmd_cmd_mechanicus_proto_depIdxs = nil
}

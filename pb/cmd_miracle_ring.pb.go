// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: cmd/cmd_miracle_ring.proto

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

type UseMiracleRingReq_MiracleRingOpType int32

const (
	UseMiracleRingReq_MIRACLE_RING_OP_NONE    UseMiracleRingReq_MiracleRingOpType = 0
	UseMiracleRingReq_MIRACLE_RING_OP_PLACE   UseMiracleRingReq_MiracleRingOpType = 1
	UseMiracleRingReq_MIRACLE_RING_OP_RETRACT UseMiracleRingReq_MiracleRingOpType = 2
)

// Enum value maps for UseMiracleRingReq_MiracleRingOpType.
var (
	UseMiracleRingReq_MiracleRingOpType_name = map[int32]string{
		0: "MIRACLE_RING_OP_NONE",
		1: "MIRACLE_RING_OP_PLACE",
		2: "MIRACLE_RING_OP_RETRACT",
	}
	UseMiracleRingReq_MiracleRingOpType_value = map[string]int32{
		"MIRACLE_RING_OP_NONE":    0,
		"MIRACLE_RING_OP_PLACE":   1,
		"MIRACLE_RING_OP_RETRACT": 2,
	}
)

func (x UseMiracleRingReq_MiracleRingOpType) Enum() *UseMiracleRingReq_MiracleRingOpType {
	p := new(UseMiracleRingReq_MiracleRingOpType)
	*p = x
	return p
}

func (x UseMiracleRingReq_MiracleRingOpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UseMiracleRingReq_MiracleRingOpType) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_miracle_ring_proto_enumTypes[0].Descriptor()
}

func (UseMiracleRingReq_MiracleRingOpType) Type() protoreflect.EnumType {
	return &file_cmd_cmd_miracle_ring_proto_enumTypes[0]
}

func (x UseMiracleRingReq_MiracleRingOpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UseMiracleRingReq_MiracleRingOpType.Descriptor instead.
func (UseMiracleRingReq_MiracleRingOpType) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{0, 0}
}

type UseMiracleRingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleRingOpType uint32  `protobuf:"varint,13,opt,name=miracle_ring_op_type,json=miracleRingOpType,proto3" json:"miracle_ring_op_type,omitempty"`
	Pos               *Vector `protobuf:"bytes,8,opt,name=pos,proto3" json:"pos,omitempty"`
	Rot               *Vector `protobuf:"bytes,7,opt,name=rot,proto3" json:"rot,omitempty"`
}

func (x *UseMiracleRingReq) Reset() {
	*x = UseMiracleRingReq{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UseMiracleRingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseMiracleRingReq) ProtoMessage() {}

func (x *UseMiracleRingReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseMiracleRingReq.ProtoReflect.Descriptor instead.
func (*UseMiracleRingReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{0}
}

func (x *UseMiracleRingReq) GetMiracleRingOpType() uint32 {
	if x != nil {
		return x.MiracleRingOpType
	}
	return 0
}

func (x *UseMiracleRingReq) GetPos() *Vector {
	if x != nil {
		return x.Pos
	}
	return nil
}

func (x *UseMiracleRingReq) GetRot() *Vector {
	if x != nil {
		return x.Rot
	}
	return nil
}

type UseMiracleRingRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode           int32  `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MiracleRingOpType uint32 `protobuf:"varint,7,opt,name=miracle_ring_op_type,json=miracleRingOpType,proto3" json:"miracle_ring_op_type,omitempty"`
}

func (x *UseMiracleRingRsp) Reset() {
	*x = UseMiracleRingRsp{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UseMiracleRingRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UseMiracleRingRsp) ProtoMessage() {}

func (x *UseMiracleRingRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UseMiracleRingRsp.ProtoReflect.Descriptor instead.
func (*UseMiracleRingRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{1}
}

func (x *UseMiracleRingRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *UseMiracleRingRsp) GetMiracleRingOpType() uint32 {
	if x != nil {
		return x.MiracleRingOpType
	}
	return 0
}

type MiracleRingDataNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsGadgetCreated     bool   `protobuf:"varint,8,opt,name=is_gadget_created,json=isGadgetCreated,proto3" json:"is_gadget_created,omitempty"`
	LastTakeRewardTime  uint32 `protobuf:"varint,14,opt,name=last_take_reward_time,json=lastTakeRewardTime,proto3" json:"last_take_reward_time,omitempty"`
	GadgetEntityId      uint32 `protobuf:"varint,12,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
	LastDeliverItemTime uint32 `protobuf:"varint,10,opt,name=last_deliver_item_time,json=lastDeliverItemTime,proto3" json:"last_deliver_item_time,omitempty"`
	MiracleRingCd       uint32 `protobuf:"varint,7,opt,name=miracle_ring_cd,json=miracleRingCd,proto3" json:"miracle_ring_cd,omitempty"`
}

func (x *MiracleRingDataNotify) Reset() {
	*x = MiracleRingDataNotify{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingDataNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingDataNotify) ProtoMessage() {}

func (x *MiracleRingDataNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingDataNotify.ProtoReflect.Descriptor instead.
func (*MiracleRingDataNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{2}
}

func (x *MiracleRingDataNotify) GetIsGadgetCreated() bool {
	if x != nil {
		return x.IsGadgetCreated
	}
	return false
}

func (x *MiracleRingDataNotify) GetLastTakeRewardTime() uint32 {
	if x != nil {
		return x.LastTakeRewardTime
	}
	return 0
}

func (x *MiracleRingDataNotify) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

func (x *MiracleRingDataNotify) GetLastDeliverItemTime() uint32 {
	if x != nil {
		return x.LastDeliverItemTime
	}
	return 0
}

func (x *MiracleRingDataNotify) GetMiracleRingCd() uint32 {
	if x != nil {
		return x.MiracleRingCd
	}
	return 0
}

type MiracleRingTakeRewardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GadgetId       uint32 `protobuf:"varint,11,opt,name=gadget_id,json=gadgetId,proto3" json:"gadget_id,omitempty"`
	GadgetEntityId uint32 `protobuf:"varint,7,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
}

func (x *MiracleRingTakeRewardReq) Reset() {
	*x = MiracleRingTakeRewardReq{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingTakeRewardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingTakeRewardReq) ProtoMessage() {}

func (x *MiracleRingTakeRewardReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingTakeRewardReq.ProtoReflect.Descriptor instead.
func (*MiracleRingTakeRewardReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{3}
}

func (x *MiracleRingTakeRewardReq) GetGadgetId() uint32 {
	if x != nil {
		return x.GadgetId
	}
	return 0
}

func (x *MiracleRingTakeRewardReq) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

type MiracleRingTakeRewardRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode int32 `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MiracleRingTakeRewardRsp) Reset() {
	*x = MiracleRingTakeRewardRsp{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingTakeRewardRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingTakeRewardRsp) ProtoMessage() {}

func (x *MiracleRingTakeRewardRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingTakeRewardRsp.ProtoReflect.Descriptor instead.
func (*MiracleRingTakeRewardRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{4}
}

func (x *MiracleRingTakeRewardRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

type MiracleRingDropResultNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastTakeRewardTime int32 `protobuf:"varint,5,opt,name=last_take_reward_time,json=lastTakeRewardTime,proto3" json:"last_take_reward_time,omitempty"`
	DropResult         int32 `protobuf:"varint,9,opt,name=drop_result,json=dropResult,proto3" json:"drop_result,omitempty"`
}

func (x *MiracleRingDropResultNotify) Reset() {
	*x = MiracleRingDropResultNotify{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingDropResultNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingDropResultNotify) ProtoMessage() {}

func (x *MiracleRingDropResultNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingDropResultNotify.ProtoReflect.Descriptor instead.
func (*MiracleRingDropResultNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{5}
}

func (x *MiracleRingDropResultNotify) GetLastTakeRewardTime() int32 {
	if x != nil {
		return x.LastTakeRewardTime
	}
	return 0
}

func (x *MiracleRingDropResultNotify) GetDropResult() int32 {
	if x != nil {
		return x.DropResult
	}
	return 0
}

type MiracleRingDeliverItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpType             InterOpType  `protobuf:"varint,9,opt,name=op_type,json=opType,proto3,enum=proto.InterOpType" json:"op_type,omitempty"`
	ItemParamList      []*ItemParam `protobuf:"bytes,1,rep,name=item_param_list,json=itemParamList,proto3" json:"item_param_list,omitempty"`
	FoodWeaponGuidList []uint64     `protobuf:"varint,4,rep,packed,name=food_weapon_guid_list,json=foodWeaponGuidList,proto3" json:"food_weapon_guid_list,omitempty"`
	GadgetId           uint32       `protobuf:"varint,14,opt,name=gadget_id,json=gadgetId,proto3" json:"gadget_id,omitempty"`
	GadgetEntityId     uint32       `protobuf:"varint,5,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
}

func (x *MiracleRingDeliverItemReq) Reset() {
	*x = MiracleRingDeliverItemReq{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingDeliverItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingDeliverItemReq) ProtoMessage() {}

func (x *MiracleRingDeliverItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingDeliverItemReq.ProtoReflect.Descriptor instead.
func (*MiracleRingDeliverItemReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{6}
}

func (x *MiracleRingDeliverItemReq) GetOpType() InterOpType {
	if x != nil {
		return x.OpType
	}
	return InterOpType_INTER_OP_FINISH
}

func (x *MiracleRingDeliverItemReq) GetItemParamList() []*ItemParam {
	if x != nil {
		return x.ItemParamList
	}
	return nil
}

func (x *MiracleRingDeliverItemReq) GetFoodWeaponGuidList() []uint64 {
	if x != nil {
		return x.FoodWeaponGuidList
	}
	return nil
}

func (x *MiracleRingDeliverItemReq) GetGadgetId() uint32 {
	if x != nil {
		return x.GadgetId
	}
	return 0
}

func (x *MiracleRingDeliverItemReq) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

type MiracleRingDeliverItemRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InteractType   InteractType `protobuf:"varint,15,opt,name=interact_type,json=interactType,proto3,enum=proto.InteractType" json:"interact_type,omitempty"`
	Retcode        int32        `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OpType         InterOpType  `protobuf:"varint,14,opt,name=op_type,json=opType,proto3,enum=proto.InterOpType" json:"op_type,omitempty"`
	GadgetId       uint32       `protobuf:"varint,4,opt,name=gadget_id,json=gadgetId,proto3" json:"gadget_id,omitempty"`
	GadgetEntityId uint32       `protobuf:"varint,9,opt,name=gadget_entity_id,json=gadgetEntityId,proto3" json:"gadget_entity_id,omitempty"`
}

func (x *MiracleRingDeliverItemRsp) Reset() {
	*x = MiracleRingDeliverItemRsp{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingDeliverItemRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingDeliverItemRsp) ProtoMessage() {}

func (x *MiracleRingDeliverItemRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingDeliverItemRsp.ProtoReflect.Descriptor instead.
func (*MiracleRingDeliverItemRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{7}
}

func (x *MiracleRingDeliverItemRsp) GetInteractType() InteractType {
	if x != nil {
		return x.InteractType
	}
	return InteractType_INTERACT_NONE
}

func (x *MiracleRingDeliverItemRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MiracleRingDeliverItemRsp) GetOpType() InterOpType {
	if x != nil {
		return x.OpType
	}
	return InterOpType_INTER_OP_FINISH
}

func (x *MiracleRingDeliverItemRsp) GetGadgetId() uint32 {
	if x != nil {
		return x.GadgetId
	}
	return 0
}

func (x *MiracleRingDeliverItemRsp) GetGadgetEntityId() uint32 {
	if x != nil {
		return x.GadgetEntityId
	}
	return 0
}

type MiracleRingDestroyNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId uint32 `protobuf:"varint,7,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
}

func (x *MiracleRingDestroyNotify) Reset() {
	*x = MiracleRingDestroyNotify{}
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MiracleRingDestroyNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiracleRingDestroyNotify) ProtoMessage() {}

func (x *MiracleRingDestroyNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_miracle_ring_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiracleRingDestroyNotify.ProtoReflect.Descriptor instead.
func (*MiracleRingDestroyNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_miracle_ring_proto_rawDescGZIP(), []int{8}
}

func (x *MiracleRingDestroyNotify) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

var File_cmd_cmd_miracle_ring_proto protoreflect.FileDescriptor

var file_cmd_cmd_miracle_ring_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a,
	0x14, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x69, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x12,
	0x1f, 0x0a, 0x03, 0x72, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x72, 0x6f, 0x74,
	0x22, 0x65, 0x0a, 0x11, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x4f,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45,
	0x5f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x4d, 0x49, 0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x5f,
	0x4f, 0x50, 0x5f, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x49,
	0x52, 0x41, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x4f, 0x50, 0x5f, 0x52, 0x45,
	0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x02, 0x22, 0x5e, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e,
	0x67, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xfd, 0x01, 0x0a, 0x15, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73,
	0x47, 0x61, 0x64, 0x67, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a,
	0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6c, 0x61,
	0x73, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6c, 0x61, 0x73, 0x74,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x72, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x52, 0x69, 0x6e, 0x67, 0x43, 0x64, 0x22, 0x61, 0x0a, 0x18, 0x4d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x18, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x71, 0x0a, 0x1b, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x44,
	0x72, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12,
	0x31, 0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x6c, 0x61, 0x73, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x72, 0x6f, 0x70, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x19, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38,
	0x0a, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0d, 0x69, 0x74, 0x65, 0x6d, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x15, 0x66, 0x6f, 0x6f, 0x64,
	0x5f, 0x77, 0x65, 0x61, 0x70, 0x6f, 0x6e, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x12, 0x66, 0x6f, 0x6f, 0x64, 0x57, 0x65, 0x61,
	0x70, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67,
	0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x19, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70,
	0x12, 0x38, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x10, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x18, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x52, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x64, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b,
	0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_miracle_ring_proto_rawDescOnce sync.Once
	file_cmd_cmd_miracle_ring_proto_rawDescData = file_cmd_cmd_miracle_ring_proto_rawDesc
)

func file_cmd_cmd_miracle_ring_proto_rawDescGZIP() []byte {
	file_cmd_cmd_miracle_ring_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_miracle_ring_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_miracle_ring_proto_rawDescData)
	})
	return file_cmd_cmd_miracle_ring_proto_rawDescData
}

var file_cmd_cmd_miracle_ring_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_cmd_miracle_ring_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cmd_cmd_miracle_ring_proto_goTypes = []any{
	(UseMiracleRingReq_MiracleRingOpType)(0), // 0: proto.UseMiracleRingReq.MiracleRingOpType
	(*UseMiracleRingReq)(nil),                // 1: proto.UseMiracleRingReq
	(*UseMiracleRingRsp)(nil),                // 2: proto.UseMiracleRingRsp
	(*MiracleRingDataNotify)(nil),            // 3: proto.MiracleRingDataNotify
	(*MiracleRingTakeRewardReq)(nil),         // 4: proto.MiracleRingTakeRewardReq
	(*MiracleRingTakeRewardRsp)(nil),         // 5: proto.MiracleRingTakeRewardRsp
	(*MiracleRingDropResultNotify)(nil),      // 6: proto.MiracleRingDropResultNotify
	(*MiracleRingDeliverItemReq)(nil),        // 7: proto.MiracleRingDeliverItemReq
	(*MiracleRingDeliverItemRsp)(nil),        // 8: proto.MiracleRingDeliverItemRsp
	(*MiracleRingDestroyNotify)(nil),         // 9: proto.MiracleRingDestroyNotify
	(*Vector)(nil),                           // 10: proto.Vector
	(InterOpType)(0),                         // 11: proto.InterOpType
	(*ItemParam)(nil),                        // 12: proto.ItemParam
	(InteractType)(0),                        // 13: proto.InteractType
}
var file_cmd_cmd_miracle_ring_proto_depIdxs = []int32{
	10, // 0: proto.UseMiracleRingReq.pos:type_name -> proto.Vector
	10, // 1: proto.UseMiracleRingReq.rot:type_name -> proto.Vector
	11, // 2: proto.MiracleRingDeliverItemReq.op_type:type_name -> proto.InterOpType
	12, // 3: proto.MiracleRingDeliverItemReq.item_param_list:type_name -> proto.ItemParam
	13, // 4: proto.MiracleRingDeliverItemRsp.interact_type:type_name -> proto.InteractType
	11, // 5: proto.MiracleRingDeliverItemRsp.op_type:type_name -> proto.InterOpType
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_cmd_miracle_ring_proto_init() }
func file_cmd_cmd_miracle_ring_proto_init() {
	if File_cmd_cmd_miracle_ring_proto != nil {
		return
	}
	file_define_proto_init()
	file_cmd_cmd_gadget_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_cmd_miracle_ring_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_miracle_ring_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_miracle_ring_proto_depIdxs,
		EnumInfos:         file_cmd_cmd_miracle_ring_proto_enumTypes,
		MessageInfos:      file_cmd_cmd_miracle_ring_proto_msgTypes,
	}.Build()
	File_cmd_cmd_miracle_ring_proto = out.File
	file_cmd_cmd_miracle_ring_proto_rawDesc = nil
	file_cmd_cmd_miracle_ring_proto_goTypes = nil
	file_cmd_cmd_miracle_ring_proto_depIdxs = nil
}

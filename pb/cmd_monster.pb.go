// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: cmd/cmd_monster.proto

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

type MonsterSummonTagNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SummonTagMap    map[uint32]uint32 `protobuf:"bytes,1,rep,name=summon_tag_map,json=summonTagMap,proto3" json:"summon_tag_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MonsterEntityId uint32            `protobuf:"varint,8,opt,name=monster_entity_id,json=monsterEntityId,proto3" json:"monster_entity_id,omitempty"`
}

func (x *MonsterSummonTagNotify) Reset() {
	*x = MonsterSummonTagNotify{}
	mi := &file_cmd_cmd_monster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonsterSummonTagNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonsterSummonTagNotify) ProtoMessage() {}

func (x *MonsterSummonTagNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_monster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonsterSummonTagNotify.ProtoReflect.Descriptor instead.
func (*MonsterSummonTagNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_monster_proto_rawDescGZIP(), []int{0}
}

func (x *MonsterSummonTagNotify) GetSummonTagMap() map[uint32]uint32 {
	if x != nil {
		return x.SummonTagMap
	}
	return nil
}

func (x *MonsterSummonTagNotify) GetMonsterEntityId() uint32 {
	if x != nil {
		return x.MonsterEntityId
	}
	return 0
}

var File_cmd_cmd_monster_proto protoreflect.FileDescriptor

var file_cmd_cmd_monster_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc,
	0x01, 0x0a, 0x16, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e,
	0x54, 0x61, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x55, 0x0a, 0x0e, 0x73, 0x75, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x73, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4d, 0x61, 0x70,
	0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x6f, 0x6e,
	0x73, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a, 0x3f, 0x0a, 0x11,
	0x53, 0x75, 0x6d, 0x6d, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76,
	0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_monster_proto_rawDescOnce sync.Once
	file_cmd_cmd_monster_proto_rawDescData = file_cmd_cmd_monster_proto_rawDesc
)

func file_cmd_cmd_monster_proto_rawDescGZIP() []byte {
	file_cmd_cmd_monster_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_monster_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_monster_proto_rawDescData)
	})
	return file_cmd_cmd_monster_proto_rawDescData
}

var file_cmd_cmd_monster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_cmd_monster_proto_goTypes = []any{
	(*MonsterSummonTagNotify)(nil), // 0: proto.MonsterSummonTagNotify
	nil,                            // 1: proto.MonsterSummonTagNotify.SummonTagMapEntry
}
var file_cmd_cmd_monster_proto_depIdxs = []int32{
	1, // 0: proto.MonsterSummonTagNotify.summon_tag_map:type_name -> proto.MonsterSummonTagNotify.SummonTagMapEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_cmd_monster_proto_init() }
func file_cmd_cmd_monster_proto_init() {
	if File_cmd_cmd_monster_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_cmd_monster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_monster_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_monster_proto_depIdxs,
		MessageInfos:      file_cmd_cmd_monster_proto_msgTypes,
	}.Build()
	File_cmd_cmd_monster_proto = out.File
	file_cmd_cmd_monster_proto_rawDesc = nil
	file_cmd_cmd_monster_proto_goTypes = nil
	file_cmd_cmd_monster_proto_depIdxs = nil
}

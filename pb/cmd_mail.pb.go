// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc2
// source: cmd/cmd_mail.proto

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

type AuthkeySignType int32

const (
	AuthkeySignType_AUTHKEY_SIGN_TYPE_NONE    AuthkeySignType = 0
	AuthkeySignType_AUTHKEY_SIGN_TYPE_DEFAULT AuthkeySignType = 1
	AuthkeySignType_AUTHKEY_SIGN_TYPE_RSA     AuthkeySignType = 2
)

// Enum value maps for AuthkeySignType.
var (
	AuthkeySignType_name = map[int32]string{
		0: "AUTHKEY_SIGN_TYPE_NONE",
		1: "AUTHKEY_SIGN_TYPE_DEFAULT",
		2: "AUTHKEY_SIGN_TYPE_RSA",
	}
	AuthkeySignType_value = map[string]int32{
		"AUTHKEY_SIGN_TYPE_NONE":    0,
		"AUTHKEY_SIGN_TYPE_DEFAULT": 1,
		"AUTHKEY_SIGN_TYPE_RSA":     2,
	}
)

func (x AuthkeySignType) Enum() *AuthkeySignType {
	p := new(AuthkeySignType)
	*p = x
	return p
}

func (x AuthkeySignType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthkeySignType) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_cmd_mail_proto_enumTypes[0].Descriptor()
}

func (AuthkeySignType) Type() protoreflect.EnumType {
	return &file_cmd_cmd_mail_proto_enumTypes[0]
}

func (x AuthkeySignType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthkeySignType.Descriptor instead.
func (AuthkeySignType) EnumDescriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{0}
}

type MailChangeNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailList      []*MailData `protobuf:"bytes,14,rep,name=mail_list,json=mailList,proto3" json:"mail_list,omitempty"`
	DelMailIdList []uint32    `protobuf:"varint,8,rep,packed,name=del_mail_id_list,json=delMailIdList,proto3" json:"del_mail_id_list,omitempty"`
}

func (x *MailChangeNotify) Reset() {
	*x = MailChangeNotify{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MailChangeNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailChangeNotify) ProtoMessage() {}

func (x *MailChangeNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailChangeNotify.ProtoReflect.Descriptor instead.
func (*MailChangeNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{0}
}

func (x *MailChangeNotify) GetMailList() []*MailData {
	if x != nil {
		return x.MailList
	}
	return nil
}

func (x *MailChangeNotify) GetDelMailIdList() []uint32 {
	if x != nil {
		return x.DelMailIdList
	}
	return nil
}

type ReadMailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailIdList []uint32 `protobuf:"varint,2,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
}

func (x *ReadMailNotify) Reset() {
	*x = ReadMailNotify{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadMailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadMailNotify) ProtoMessage() {}

func (x *ReadMailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadMailNotify.ProtoReflect.Descriptor instead.
func (*ReadMailNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{1}
}

func (x *ReadMailNotify) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

type GetMailItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailIdList []uint32 `protobuf:"varint,6,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
}

func (x *GetMailItemReq) Reset() {
	*x = GetMailItemReq{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMailItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailItemReq) ProtoMessage() {}

func (x *GetMailItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailItemReq.ProtoReflect.Descriptor instead.
func (*GetMailItemReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{2}
}

func (x *GetMailItemReq) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

type GetMailItemRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32         `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MailIdList []uint32      `protobuf:"varint,3,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
	ItemList   []*EquipParam `protobuf:"bytes,2,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *GetMailItemRsp) Reset() {
	*x = GetMailItemRsp{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMailItemRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMailItemRsp) ProtoMessage() {}

func (x *GetMailItemRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMailItemRsp.ProtoReflect.Descriptor instead.
func (*GetMailItemRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{3}
}

func (x *GetMailItemRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMailItemRsp) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

func (x *GetMailItemRsp) GetItemList() []*EquipParam {
	if x != nil {
		return x.ItemList
	}
	return nil
}

type DelMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MailIdList []uint32 `protobuf:"varint,12,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
}

func (x *DelMailReq) Reset() {
	*x = DelMailReq{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelMailReq) ProtoMessage() {}

func (x *DelMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelMailReq.ProtoReflect.Descriptor instead.
func (*DelMailReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{4}
}

func (x *DelMailReq) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

type DelMailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    int32    `protobuf:"varint,11,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MailIdList []uint32 `protobuf:"varint,5,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
}

func (x *DelMailRsp) Reset() {
	*x = DelMailRsp{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DelMailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelMailRsp) ProtoMessage() {}

func (x *DelMailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelMailRsp.ProtoReflect.Descriptor instead.
func (*DelMailRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{5}
}

func (x *DelMailRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *DelMailRsp) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

type GetAuthkeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthAppid  string `protobuf:"bytes,14,opt,name=auth_appid,json=authAppid,proto3" json:"auth_appid,omitempty"`
	SignType   uint32 `protobuf:"varint,7,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	AuthkeyVer uint32 `protobuf:"varint,13,opt,name=authkey_ver,json=authkeyVer,proto3" json:"authkey_ver,omitempty"`
}

func (x *GetAuthkeyReq) Reset() {
	*x = GetAuthkeyReq{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthkeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthkeyReq) ProtoMessage() {}

func (x *GetAuthkeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthkeyReq.ProtoReflect.Descriptor instead.
func (*GetAuthkeyReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{6}
}

func (x *GetAuthkeyReq) GetAuthAppid() string {
	if x != nil {
		return x.AuthAppid
	}
	return ""
}

func (x *GetAuthkeyReq) GetSignType() uint32 {
	if x != nil {
		return x.SignType
	}
	return 0
}

func (x *GetAuthkeyReq) GetAuthkeyVer() uint32 {
	if x != nil {
		return x.AuthkeyVer
	}
	return 0
}

type GetAuthkeyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthAppid  string `protobuf:"bytes,4,opt,name=auth_appid,json=authAppid,proto3" json:"auth_appid,omitempty"`
	SignType   uint32 `protobuf:"varint,15,opt,name=sign_type,json=signType,proto3" json:"sign_type,omitempty"`
	Retcode    int32  `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AuthkeyVer uint32 `protobuf:"varint,9,opt,name=authkey_ver,json=authkeyVer,proto3" json:"authkey_ver,omitempty"`
	GameBiz    string `protobuf:"bytes,11,opt,name=game_biz,json=gameBiz,proto3" json:"game_biz,omitempty"`
	Authkey    string `protobuf:"bytes,3,opt,name=authkey,proto3" json:"authkey,omitempty"`
}

func (x *GetAuthkeyRsp) Reset() {
	*x = GetAuthkeyRsp{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAuthkeyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthkeyRsp) ProtoMessage() {}

func (x *GetAuthkeyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthkeyRsp.ProtoReflect.Descriptor instead.
func (*GetAuthkeyRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{7}
}

func (x *GetAuthkeyRsp) GetAuthAppid() string {
	if x != nil {
		return x.AuthAppid
	}
	return ""
}

func (x *GetAuthkeyRsp) GetSignType() uint32 {
	if x != nil {
		return x.SignType
	}
	return 0
}

func (x *GetAuthkeyRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAuthkeyRsp) GetAuthkeyVer() uint32 {
	if x != nil {
		return x.AuthkeyVer
	}
	return 0
}

func (x *GetAuthkeyRsp) GetGameBiz() string {
	if x != nil {
		return x.GameBiz
	}
	return ""
}

func (x *GetAuthkeyRsp) GetAuthkey() string {
	if x != nil {
		return x.Authkey
	}
	return ""
}

type ClientNewMailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotReadNum          uint32 `protobuf:"varint,7,opt,name=not_read_num,json=notReadNum,proto3" json:"not_read_num,omitempty"`
	NotGotAttachmentNum uint32 `protobuf:"varint,2,opt,name=not_got_attachment_num,json=notGotAttachmentNum,proto3" json:"not_got_attachment_num,omitempty"`
}

func (x *ClientNewMailNotify) Reset() {
	*x = ClientNewMailNotify{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClientNewMailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientNewMailNotify) ProtoMessage() {}

func (x *ClientNewMailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientNewMailNotify.ProtoReflect.Descriptor instead.
func (*ClientNewMailNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{8}
}

func (x *ClientNewMailNotify) GetNotReadNum() uint32 {
	if x != nil {
		return x.NotReadNum
	}
	return 0
}

func (x *ClientNewMailNotify) GetNotGotAttachmentNum() uint32 {
	if x != nil {
		return x.NotGotAttachmentNum
	}
	return 0
}

type GetAllMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCollected bool `protobuf:"varint,7,opt,name=is_collected,json=isCollected,proto3" json:"is_collected,omitempty"`
}

func (x *GetAllMailReq) Reset() {
	*x = GetAllMailReq{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailReq) ProtoMessage() {}

func (x *GetAllMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailReq.ProtoReflect.Descriptor instead.
func (*GetAllMailReq) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllMailReq) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

type GetAllMailRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     int32       `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MailList    []*MailData `protobuf:"bytes,14,rep,name=mail_list,json=mailList,proto3" json:"mail_list,omitempty"`
	IsCollected bool        `protobuf:"varint,1,opt,name=is_collected,json=isCollected,proto3" json:"is_collected,omitempty"`
	IsTruncated bool        `protobuf:"varint,2,opt,name=is_truncated,json=isTruncated,proto3" json:"is_truncated,omitempty"`
}

func (x *GetAllMailRsp) Reset() {
	*x = GetAllMailRsp{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllMailRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailRsp) ProtoMessage() {}

func (x *GetAllMailRsp) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailRsp.ProtoReflect.Descriptor instead.
func (*GetAllMailRsp) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{10}
}

func (x *GetAllMailRsp) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAllMailRsp) GetMailList() []*MailData {
	if x != nil {
		return x.MailList
	}
	return nil
}

func (x *GetAllMailRsp) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

func (x *GetAllMailRsp) GetIsTruncated() bool {
	if x != nil {
		return x.IsTruncated
	}
	return false
}

type ChangeMailStarNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsStar     bool     `protobuf:"varint,14,opt,name=is_star,json=isStar,proto3" json:"is_star,omitempty"`
	MailIdList []uint32 `protobuf:"varint,2,rep,packed,name=mail_id_list,json=mailIdList,proto3" json:"mail_id_list,omitempty"`
}

func (x *ChangeMailStarNotify) Reset() {
	*x = ChangeMailStarNotify{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeMailStarNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeMailStarNotify) ProtoMessage() {}

func (x *ChangeMailStarNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeMailStarNotify.ProtoReflect.Descriptor instead.
func (*ChangeMailStarNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{11}
}

func (x *ChangeMailStarNotify) GetIsStar() bool {
	if x != nil {
		return x.IsStar
	}
	return false
}

func (x *ChangeMailStarNotify) GetMailIdList() []uint32 {
	if x != nil {
		return x.MailIdList
	}
	return nil
}

type GetAllMailNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsCollected bool `protobuf:"varint,13,opt,name=is_collected,json=isCollected,proto3" json:"is_collected,omitempty"`
}

func (x *GetAllMailNotify) Reset() {
	*x = GetAllMailNotify{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllMailNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailNotify) ProtoMessage() {}

func (x *GetAllMailNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailNotify.ProtoReflect.Descriptor instead.
func (*GetAllMailNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{12}
}

func (x *GetAllMailNotify) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

type GetAllMailResultNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction    string      `protobuf:"bytes,9,opt,name=transaction,proto3" json:"transaction,omitempty"`
	MailList       []*MailData `protobuf:"bytes,5,rep,name=mail_list,json=mailList,proto3" json:"mail_list,omitempty"`
	PageIndex      uint32      `protobuf:"varint,11,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	TotalPageCount uint32      `protobuf:"varint,4,opt,name=total_page_count,json=totalPageCount,proto3" json:"total_page_count,omitempty"`
	IsCollected    bool        `protobuf:"varint,7,opt,name=is_collected,json=isCollected,proto3" json:"is_collected,omitempty"`
	Retcode        int32       `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetAllMailResultNotify) Reset() {
	*x = GetAllMailResultNotify{}
	mi := &file_cmd_cmd_mail_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllMailResultNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMailResultNotify) ProtoMessage() {}

func (x *GetAllMailResultNotify) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_cmd_mail_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMailResultNotify.ProtoReflect.Descriptor instead.
func (*GetAllMailResultNotify) Descriptor() ([]byte, []int) {
	return file_cmd_cmd_mail_proto_rawDescGZIP(), []int{13}
}

func (x *GetAllMailResultNotify) GetTransaction() string {
	if x != nil {
		return x.Transaction
	}
	return ""
}

func (x *GetAllMailResultNotify) GetMailList() []*MailData {
	if x != nil {
		return x.MailList
	}
	return nil
}

func (x *GetAllMailResultNotify) GetPageIndex() uint32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *GetAllMailResultNotify) GetTotalPageCount() uint32 {
	if x != nil {
		return x.TotalPageCount
	}
	return 0
}

func (x *GetAllMailResultNotify) GetIsCollected() bool {
	if x != nil {
		return x.IsCollected
	}
	return false
}

func (x *GetAllMailResultNotify) GetRetcode() int32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_cmd_cmd_mail_proto protoreflect.FileDescriptor

var file_cmd_cmd_mail_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6d, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x10, 0x4d, 0x61, 0x69,
	0x6c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2c, 0x0a,
	0x09, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x10, 0x64,
	0x65, 0x6c, 0x5f, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x61, 0x69, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61,
	0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61,
	0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x7c, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4d, 0x61, 0x69, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a,
	0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x41,
	0x70, 0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x56,
	0x65, 0x72, 0x22, 0xbb, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6b, 0x65,
	0x79, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x61, 0x70, 0x70,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x41, 0x70,
	0x70, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x62, 0x69, 0x7a, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67,
	0x61, 0x6d, 0x65, 0x42, 0x69, 0x7a, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79,
	0x22, 0x6c, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x65, 0x77, 0x4d, 0x61, 0x69,
	0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6e,
	0x6f, 0x74, 0x52, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x33, 0x0a, 0x16, 0x6e, 0x6f, 0x74,
	0x5f, 0x67, 0x6f, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6e, 0x6f, 0x74, 0x47, 0x6f,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x32,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69,
	0x6c, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c,
	0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x69, 0x6c,
	0x53, 0x74, 0x61, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x53,
	0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d,
	0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0xee, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x09, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x67, 0x0a,
	0x0f, 0x41, 0x75, 0x74, 0x68, 0x6b, 0x65, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x16, 0x41, 0x55, 0x54, 0x48, 0x4b, 0x45, 0x59, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19,
	0x41, 0x55, 0x54, 0x48, 0x4b, 0x45, 0x59, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x41,
	0x55, 0x54, 0x48, 0x4b, 0x45, 0x59, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x53, 0x41, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x79, 0x76, 0x61, 0x74, 0x2d, 0x68, 0x65, 0x6c, 0x70,
	0x65, 0x72, 0x2f, 0x68, 0x6b, 0x34, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_cmd_mail_proto_rawDescOnce sync.Once
	file_cmd_cmd_mail_proto_rawDescData = file_cmd_cmd_mail_proto_rawDesc
)

func file_cmd_cmd_mail_proto_rawDescGZIP() []byte {
	file_cmd_cmd_mail_proto_rawDescOnce.Do(func() {
		file_cmd_cmd_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_cmd_mail_proto_rawDescData)
	})
	return file_cmd_cmd_mail_proto_rawDescData
}

var file_cmd_cmd_mail_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cmd_cmd_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_cmd_cmd_mail_proto_goTypes = []any{
	(AuthkeySignType)(0),           // 0: proto.AuthkeySignType
	(*MailChangeNotify)(nil),       // 1: proto.MailChangeNotify
	(*ReadMailNotify)(nil),         // 2: proto.ReadMailNotify
	(*GetMailItemReq)(nil),         // 3: proto.GetMailItemReq
	(*GetMailItemRsp)(nil),         // 4: proto.GetMailItemRsp
	(*DelMailReq)(nil),             // 5: proto.DelMailReq
	(*DelMailRsp)(nil),             // 6: proto.DelMailRsp
	(*GetAuthkeyReq)(nil),          // 7: proto.GetAuthkeyReq
	(*GetAuthkeyRsp)(nil),          // 8: proto.GetAuthkeyRsp
	(*ClientNewMailNotify)(nil),    // 9: proto.ClientNewMailNotify
	(*GetAllMailReq)(nil),          // 10: proto.GetAllMailReq
	(*GetAllMailRsp)(nil),          // 11: proto.GetAllMailRsp
	(*ChangeMailStarNotify)(nil),   // 12: proto.ChangeMailStarNotify
	(*GetAllMailNotify)(nil),       // 13: proto.GetAllMailNotify
	(*GetAllMailResultNotify)(nil), // 14: proto.GetAllMailResultNotify
	(*MailData)(nil),               // 15: proto.MailData
	(*EquipParam)(nil),             // 16: proto.EquipParam
}
var file_cmd_cmd_mail_proto_depIdxs = []int32{
	15, // 0: proto.MailChangeNotify.mail_list:type_name -> proto.MailData
	16, // 1: proto.GetMailItemRsp.item_list:type_name -> proto.EquipParam
	15, // 2: proto.GetAllMailRsp.mail_list:type_name -> proto.MailData
	15, // 3: proto.GetAllMailResultNotify.mail_list:type_name -> proto.MailData
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cmd_cmd_mail_proto_init() }
func file_cmd_cmd_mail_proto_init() {
	if File_cmd_cmd_mail_proto != nil {
		return
	}
	file_define_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_cmd_mail_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_cmd_mail_proto_goTypes,
		DependencyIndexes: file_cmd_cmd_mail_proto_depIdxs,
		EnumInfos:         file_cmd_cmd_mail_proto_enumTypes,
		MessageInfos:      file_cmd_cmd_mail_proto_msgTypes,
	}.Build()
	File_cmd_cmd_mail_proto = out.File
	file_cmd_cmd_mail_proto_rawDesc = nil
	file_cmd_cmd_mail_proto_goTypes = nil
	file_cmd_cmd_mail_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/msg.proto

package protomsg

/*
包名，通过protoc生成时go文件时
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgBase struct {
	ModeType             string   `protobuf:"bytes,1,opt,name=ModeType,proto3" json:"ModeType,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	MsgType              string   `protobuf:"bytes,3,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	ConnectId            int32    `protobuf:"varint,4,opt,name=ConnectId,proto3" json:"ConnectId,omitempty"`
	Datas                []byte   `protobuf:"bytes,5,opt,name=Datas,proto3" json:"Datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgBase) Reset()         { *m = MsgBase{} }
func (m *MsgBase) String() string { return proto.CompactTextString(m) }
func (*MsgBase) ProtoMessage()    {}
func (*MsgBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{0}
}
func (m *MsgBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgBase.Unmarshal(m, b)
}
func (m *MsgBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgBase.Marshal(b, m, deterministic)
}
func (dst *MsgBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBase.Merge(dst, src)
}
func (m *MsgBase) XXX_Size() int {
	return xxx_messageInfo_MsgBase.Size(m)
}
func (m *MsgBase) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBase.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBase proto.InternalMessageInfo

func (m *MsgBase) GetModeType() string {
	if m != nil {
		return m.ModeType
	}
	return ""
}

func (m *MsgBase) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MsgBase) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MsgBase) GetConnectId() int32 {
	if m != nil {
		return m.ConnectId
	}
	return 0
}

func (m *MsgBase) GetDatas() []byte {
	if m != nil {
		return m.Datas
	}
	return nil
}

type MsgRegisterToGate struct {
	ModeType             string   `protobuf:"bytes,1,opt,name=ModeType,proto3" json:"ModeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgRegisterToGate) Reset()         { *m = MsgRegisterToGate{} }
func (m *MsgRegisterToGate) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterToGate) ProtoMessage()    {}
func (*MsgRegisterToGate) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{1}
}
func (m *MsgRegisterToGate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRegisterToGate.Unmarshal(m, b)
}
func (m *MsgRegisterToGate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRegisterToGate.Marshal(b, m, deterministic)
}
func (dst *MsgRegisterToGate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterToGate.Merge(dst, src)
}
func (m *MsgRegisterToGate) XXX_Size() int {
	return xxx_messageInfo_MsgRegisterToGate.Size(m)
}
func (m *MsgRegisterToGate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterToGate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterToGate proto.InternalMessageInfo

func (m *MsgRegisterToGate) GetModeType() string {
	if m != nil {
		return m.ModeType
	}
	return ""
}

type CS_MsgQuickLogin struct {
	Platform             string   `protobuf:"bytes,1,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Machineid            string   `protobuf:"bytes,2,opt,name=Machineid,proto3" json:"Machineid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_MsgQuickLogin) Reset()         { *m = CS_MsgQuickLogin{} }
func (m *CS_MsgQuickLogin) String() string { return proto.CompactTextString(m) }
func (*CS_MsgQuickLogin) ProtoMessage()    {}
func (*CS_MsgQuickLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{2}
}
func (m *CS_MsgQuickLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_MsgQuickLogin.Unmarshal(m, b)
}
func (m *CS_MsgQuickLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_MsgQuickLogin.Marshal(b, m, deterministic)
}
func (dst *CS_MsgQuickLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_MsgQuickLogin.Merge(dst, src)
}
func (m *CS_MsgQuickLogin) XXX_Size() int {
	return xxx_messageInfo_CS_MsgQuickLogin.Size(m)
}
func (m *CS_MsgQuickLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_MsgQuickLogin.DiscardUnknown(m)
}

var xxx_messageInfo_CS_MsgQuickLogin proto.InternalMessageInfo

func (m *CS_MsgQuickLogin) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *CS_MsgQuickLogin) GetMachineid() string {
	if m != nil {
		return m.Machineid
	}
	return ""
}

func (m *CS_MsgQuickLogin) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CS_Login struct {
	Phonenumber          string   `protobuf:"bytes,1,opt,name=Phonenumber,proto3" json:"Phonenumber,omitempty"`
	Platform             string   `protobuf:"bytes,2,opt,name=Platform,proto3" json:"Platform,omitempty"`
	Machineid            string   `protobuf:"bytes,3,opt,name=Machineid,proto3" json:"Machineid,omitempty"`
	WechatId             string   `protobuf:"bytes,4,opt,name=Wechat_id,json=WechatId,proto3" json:"Wechat_id,omitempty"`
	QQId                 string   `protobuf:"bytes,5,opt,name=QQ_id,json=QQId,proto3" json:"QQ_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Login) Reset()         { *m = CS_Login{} }
func (m *CS_Login) String() string { return proto.CompactTextString(m) }
func (*CS_Login) ProtoMessage()    {}
func (*CS_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{3}
}
func (m *CS_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Login.Unmarshal(m, b)
}
func (m *CS_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Login.Marshal(b, m, deterministic)
}
func (dst *CS_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Login.Merge(dst, src)
}
func (m *CS_Login) XXX_Size() int {
	return xxx_messageInfo_CS_Login.Size(m)
}
func (m *CS_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Login.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Login proto.InternalMessageInfo

func (m *CS_Login) GetPhonenumber() string {
	if m != nil {
		return m.Phonenumber
	}
	return ""
}

func (m *CS_Login) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *CS_Login) GetMachineid() string {
	if m != nil {
		return m.Machineid
	}
	return ""
}

func (m *CS_Login) GetWechatId() string {
	if m != nil {
		return m.WechatId
	}
	return ""
}

func (m *CS_Login) GetQQId() string {
	if m != nil {
		return m.QQId
	}
	return ""
}

// 玩家移动操作
type CS_PlayerMove struct {
	IDs                  []int32  `protobuf:"varint,1,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
	X                    float32  `protobuf:"fixed32,2,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,3,opt,name=Y,proto3" json:"Y,omitempty"`
	IsStart              bool     `protobuf:"varint,4,opt,name=IsStart,proto3" json:"IsStart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_PlayerMove) Reset()         { *m = CS_PlayerMove{} }
func (m *CS_PlayerMove) String() string { return proto.CompactTextString(m) }
func (*CS_PlayerMove) ProtoMessage()    {}
func (*CS_PlayerMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{4}
}
func (m *CS_PlayerMove) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_PlayerMove.Unmarshal(m, b)
}
func (m *CS_PlayerMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_PlayerMove.Marshal(b, m, deterministic)
}
func (dst *CS_PlayerMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_PlayerMove.Merge(dst, src)
}
func (m *CS_PlayerMove) XXX_Size() int {
	return xxx_messageInfo_CS_PlayerMove.Size(m)
}
func (m *CS_PlayerMove) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_PlayerMove.DiscardUnknown(m)
}

var xxx_messageInfo_CS_PlayerMove proto.InternalMessageInfo

func (m *CS_PlayerMove) GetIDs() []int32 {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *CS_PlayerMove) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CS_PlayerMove) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CS_PlayerMove) GetIsStart() bool {
	if m != nil {
		return m.IsStart
	}
	return false
}

// 玩家攻击操作
type CS_PlayerAttack struct {
	IDs                  []int32  `protobuf:"varint,1,rep,packed,name=IDs,proto3" json:"IDs,omitempty"`
	TargetUnitID         int32    `protobuf:"varint,2,opt,name=TargetUnitID,proto3" json:"TargetUnitID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_PlayerAttack) Reset()         { *m = CS_PlayerAttack{} }
func (m *CS_PlayerAttack) String() string { return proto.CompactTextString(m) }
func (*CS_PlayerAttack) ProtoMessage()    {}
func (*CS_PlayerAttack) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{5}
}
func (m *CS_PlayerAttack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_PlayerAttack.Unmarshal(m, b)
}
func (m *CS_PlayerAttack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_PlayerAttack.Marshal(b, m, deterministic)
}
func (dst *CS_PlayerAttack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_PlayerAttack.Merge(dst, src)
}
func (m *CS_PlayerAttack) XXX_Size() int {
	return xxx_messageInfo_CS_PlayerAttack.Size(m)
}
func (m *CS_PlayerAttack) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_PlayerAttack.DiscardUnknown(m)
}

var xxx_messageInfo_CS_PlayerAttack proto.InternalMessageInfo

func (m *CS_PlayerAttack) GetIDs() []int32 {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *CS_PlayerAttack) GetTargetUnitID() int32 {
	if m != nil {
		return m.TargetUnitID
	}
	return 0
}

// 玩家技能(包括道具)
type CS_PlayerSkill struct {
	IDs                  int32    `protobuf:"varint,1,opt,name=IDs,proto3" json:"IDs,omitempty"`
	SkillID              int32    `protobuf:"varint,2,opt,name=SkillID,proto3" json:"SkillID,omitempty"`
	X                    float32  `protobuf:"fixed32,3,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,4,opt,name=Y,proto3" json:"Y,omitempty"`
	TargetUnitID         int32    `protobuf:"varint,5,opt,name=TargetUnitID,proto3" json:"TargetUnitID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_PlayerSkill) Reset()         { *m = CS_PlayerSkill{} }
func (m *CS_PlayerSkill) String() string { return proto.CompactTextString(m) }
func (*CS_PlayerSkill) ProtoMessage()    {}
func (*CS_PlayerSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{6}
}
func (m *CS_PlayerSkill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_PlayerSkill.Unmarshal(m, b)
}
func (m *CS_PlayerSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_PlayerSkill.Marshal(b, m, deterministic)
}
func (dst *CS_PlayerSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_PlayerSkill.Merge(dst, src)
}
func (m *CS_PlayerSkill) XXX_Size() int {
	return xxx_messageInfo_CS_PlayerSkill.Size(m)
}
func (m *CS_PlayerSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_PlayerSkill.DiscardUnknown(m)
}

var xxx_messageInfo_CS_PlayerSkill proto.InternalMessageInfo

func (m *CS_PlayerSkill) GetIDs() int32 {
	if m != nil {
		return m.IDs
	}
	return 0
}

func (m *CS_PlayerSkill) GetSkillID() int32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

func (m *CS_PlayerSkill) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CS_PlayerSkill) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CS_PlayerSkill) GetTargetUnitID() int32 {
	if m != nil {
		return m.TargetUnitID
	}
	return 0
}

// 单位数据
type UnitDatas struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Level                int32    `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	HP                   int32    `protobuf:"varint,3,opt,name=HP,proto3" json:"HP,omitempty"`
	MP                   int32    `protobuf:"varint,4,opt,name=MP,proto3" json:"MP,omitempty"`
	X                    float32  `protobuf:"fixed32,5,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,6,opt,name=Y,proto3" json:"Y,omitempty"`
	ID                   int32    `protobuf:"varint,7,opt,name=ID,proto3" json:"ID,omitempty"`
	ModeType             string   `protobuf:"bytes,8,opt,name=ModeType,proto3" json:"ModeType,omitempty"`
	MaxHP                int32    `protobuf:"varint,9,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	MaxMP                int32    `protobuf:"varint,10,opt,name=MaxMP,proto3" json:"MaxMP,omitempty"`
	Experience           int32    `protobuf:"varint,11,opt,name=Experience,proto3" json:"Experience,omitempty"`
	MaxExperience        int32    `protobuf:"varint,12,opt,name=MaxExperience,proto3" json:"MaxExperience,omitempty"`
	ControlID            int32    `protobuf:"varint,13,opt,name=ControlID,proto3" json:"ControlID,omitempty"`
	AnimotorState        int32    `protobuf:"varint,14,opt,name=AnimotorState,proto3" json:"AnimotorState,omitempty"`
	AttackTime           float32  `protobuf:"fixed32,15,opt,name=AttackTime,proto3" json:"AttackTime,omitempty"`
	DirectionX           float32  `protobuf:"fixed32,16,opt,name=DirectionX,proto3" json:"DirectionX,omitempty"`
	DirectionY           float32  `protobuf:"fixed32,17,opt,name=DirectionY,proto3" json:"DirectionY,omitempty"`
	UnitType             int32    `protobuf:"varint,18,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
	AttackAcpabilities   int32    `protobuf:"varint,19,opt,name=AttackAcpabilities,proto3" json:"AttackAcpabilities,omitempty"`
	AttackMode           int32    `protobuf:"varint,20,opt,name=AttackMode,proto3" json:"AttackMode,omitempty"`
	IsMain               int32    `protobuf:"varint,21,opt,name=IsMain,proto3" json:"IsMain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnitDatas) Reset()         { *m = UnitDatas{} }
func (m *UnitDatas) String() string { return proto.CompactTextString(m) }
func (*UnitDatas) ProtoMessage()    {}
func (*UnitDatas) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{7}
}
func (m *UnitDatas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnitDatas.Unmarshal(m, b)
}
func (m *UnitDatas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnitDatas.Marshal(b, m, deterministic)
}
func (dst *UnitDatas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnitDatas.Merge(dst, src)
}
func (m *UnitDatas) XXX_Size() int {
	return xxx_messageInfo_UnitDatas.Size(m)
}
func (m *UnitDatas) XXX_DiscardUnknown() {
	xxx_messageInfo_UnitDatas.DiscardUnknown(m)
}

var xxx_messageInfo_UnitDatas proto.InternalMessageInfo

func (m *UnitDatas) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UnitDatas) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *UnitDatas) GetHP() int32 {
	if m != nil {
		return m.HP
	}
	return 0
}

func (m *UnitDatas) GetMP() int32 {
	if m != nil {
		return m.MP
	}
	return 0
}

func (m *UnitDatas) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *UnitDatas) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *UnitDatas) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UnitDatas) GetModeType() string {
	if m != nil {
		return m.ModeType
	}
	return ""
}

func (m *UnitDatas) GetMaxHP() int32 {
	if m != nil {
		return m.MaxHP
	}
	return 0
}

func (m *UnitDatas) GetMaxMP() int32 {
	if m != nil {
		return m.MaxMP
	}
	return 0
}

func (m *UnitDatas) GetExperience() int32 {
	if m != nil {
		return m.Experience
	}
	return 0
}

func (m *UnitDatas) GetMaxExperience() int32 {
	if m != nil {
		return m.MaxExperience
	}
	return 0
}

func (m *UnitDatas) GetControlID() int32 {
	if m != nil {
		return m.ControlID
	}
	return 0
}

func (m *UnitDatas) GetAnimotorState() int32 {
	if m != nil {
		return m.AnimotorState
	}
	return 0
}

func (m *UnitDatas) GetAttackTime() float32 {
	if m != nil {
		return m.AttackTime
	}
	return 0
}

func (m *UnitDatas) GetDirectionX() float32 {
	if m != nil {
		return m.DirectionX
	}
	return 0
}

func (m *UnitDatas) GetDirectionY() float32 {
	if m != nil {
		return m.DirectionY
	}
	return 0
}

func (m *UnitDatas) GetUnitType() int32 {
	if m != nil {
		return m.UnitType
	}
	return 0
}

func (m *UnitDatas) GetAttackAcpabilities() int32 {
	if m != nil {
		return m.AttackAcpabilities
	}
	return 0
}

func (m *UnitDatas) GetAttackMode() int32 {
	if m != nil {
		return m.AttackMode
	}
	return 0
}

func (m *UnitDatas) GetIsMain() int32 {
	if m != nil {
		return m.IsMain
	}
	return 0
}

type MsgUserEnterScene struct {
	Uid                  int32    `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	ConnectId            int32    `protobuf:"varint,2,opt,name=ConnectId,proto3" json:"ConnectId,omitempty"`
	SrcServerName        string   `protobuf:"bytes,3,opt,name=SrcServerName,proto3" json:"SrcServerName,omitempty"`
	DestServerName       string   `protobuf:"bytes,4,opt,name=DestServerName,proto3" json:"DestServerName,omitempty"`
	SceneName            string   `protobuf:"bytes,5,opt,name=SceneName,proto3" json:"SceneName,omitempty"`
	Datas                []byte   `protobuf:"bytes,6,opt,name=Datas,proto3" json:"Datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgUserEnterScene) Reset()         { *m = MsgUserEnterScene{} }
func (m *MsgUserEnterScene) String() string { return proto.CompactTextString(m) }
func (*MsgUserEnterScene) ProtoMessage()    {}
func (*MsgUserEnterScene) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{8}
}
func (m *MsgUserEnterScene) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgUserEnterScene.Unmarshal(m, b)
}
func (m *MsgUserEnterScene) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgUserEnterScene.Marshal(b, m, deterministic)
}
func (dst *MsgUserEnterScene) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUserEnterScene.Merge(dst, src)
}
func (m *MsgUserEnterScene) XXX_Size() int {
	return xxx_messageInfo_MsgUserEnterScene.Size(m)
}
func (m *MsgUserEnterScene) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUserEnterScene.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUserEnterScene proto.InternalMessageInfo

func (m *MsgUserEnterScene) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *MsgUserEnterScene) GetConnectId() int32 {
	if m != nil {
		return m.ConnectId
	}
	return 0
}

func (m *MsgUserEnterScene) GetSrcServerName() string {
	if m != nil {
		return m.SrcServerName
	}
	return ""
}

func (m *MsgUserEnterScene) GetDestServerName() string {
	if m != nil {
		return m.DestServerName
	}
	return ""
}

func (m *MsgUserEnterScene) GetSceneName() string {
	if m != nil {
		return m.SceneName
	}
	return ""
}

func (m *MsgUserEnterScene) GetDatas() []byte {
	if m != nil {
		return m.Datas
	}
	return nil
}

// 每帧更新单位数据
type SC_Update struct {
	CurFrame             int32        `protobuf:"varint,1,opt,name=CurFrame,proto3" json:"CurFrame,omitempty"`
	NewUnits             []*UnitDatas `protobuf:"bytes,2,rep,name=NewUnits,proto3" json:"NewUnits,omitempty"`
	OldUnits             []*UnitDatas `protobuf:"bytes,3,rep,name=OldUnits,proto3" json:"OldUnits,omitempty"`
	RemoveUnits          []int32      `protobuf:"varint,4,rep,packed,name=RemoveUnits,proto3" json:"RemoveUnits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SC_Update) Reset()         { *m = SC_Update{} }
func (m *SC_Update) String() string { return proto.CompactTextString(m) }
func (*SC_Update) ProtoMessage()    {}
func (*SC_Update) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{9}
}
func (m *SC_Update) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Update.Unmarshal(m, b)
}
func (m *SC_Update) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Update.Marshal(b, m, deterministic)
}
func (dst *SC_Update) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Update.Merge(dst, src)
}
func (m *SC_Update) XXX_Size() int {
	return xxx_messageInfo_SC_Update.Size(m)
}
func (m *SC_Update) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Update.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Update proto.InternalMessageInfo

func (m *SC_Update) GetCurFrame() int32 {
	if m != nil {
		return m.CurFrame
	}
	return 0
}

func (m *SC_Update) GetNewUnits() []*UnitDatas {
	if m != nil {
		return m.NewUnits
	}
	return nil
}

func (m *SC_Update) GetOldUnits() []*UnitDatas {
	if m != nil {
		return m.OldUnits
	}
	return nil
}

func (m *SC_Update) GetRemoveUnits() []int32 {
	if m != nil {
		return m.RemoveUnits
	}
	return nil
}

// 玩家进入新场景时的场景信息
type SC_NewScene struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	LogicFps             int32    `protobuf:"varint,2,opt,name=LogicFps,proto3" json:"LogicFps,omitempty"`
	CurFrame             int32    `protobuf:"varint,3,opt,name=CurFrame,proto3" json:"CurFrame,omitempty"`
	ServerName           string   `protobuf:"bytes,4,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_NewScene) Reset()         { *m = SC_NewScene{} }
func (m *SC_NewScene) String() string { return proto.CompactTextString(m) }
func (*SC_NewScene) ProtoMessage()    {}
func (*SC_NewScene) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{10}
}
func (m *SC_NewScene) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_NewScene.Unmarshal(m, b)
}
func (m *SC_NewScene) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_NewScene.Marshal(b, m, deterministic)
}
func (dst *SC_NewScene) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_NewScene.Merge(dst, src)
}
func (m *SC_NewScene) XXX_Size() int {
	return xxx_messageInfo_SC_NewScene.Size(m)
}
func (m *SC_NewScene) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_NewScene.DiscardUnknown(m)
}

var xxx_messageInfo_SC_NewScene proto.InternalMessageInfo

func (m *SC_NewScene) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SC_NewScene) GetLogicFps() int32 {
	if m != nil {
		return m.LogicFps
	}
	return 0
}

func (m *SC_NewScene) GetCurFrame() int32 {
	if m != nil {
		return m.CurFrame
	}
	return 0
}

func (m *SC_NewScene) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

type SC_Logined struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Uid                  int32    `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Logined) Reset()         { *m = SC_Logined{} }
func (m *SC_Logined) String() string { return proto.CompactTextString(m) }
func (*SC_Logined) ProtoMessage()    {}
func (*SC_Logined) Descriptor() ([]byte, []int) {
	return fileDescriptor_617aaba9bf0c07c7, []int{11}
}
func (m *SC_Logined) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Logined.Unmarshal(m, b)
}
func (m *SC_Logined) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Logined.Marshal(b, m, deterministic)
}
func (dst *SC_Logined) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Logined.Merge(dst, src)
}
func (m *SC_Logined) XXX_Size() int {
	return xxx_messageInfo_SC_Logined.Size(m)
}
func (m *SC_Logined) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Logined.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Logined proto.InternalMessageInfo

func (m *SC_Logined) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SC_Logined) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SC_Logined) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgBase)(nil), "protomsg.MsgBase")
	proto.RegisterType((*MsgRegisterToGate)(nil), "protomsg.MsgRegisterToGate")
	proto.RegisterType((*CS_MsgQuickLogin)(nil), "protomsg.CS_MsgQuickLogin")
	proto.RegisterType((*CS_Login)(nil), "protomsg.CS_Login")
	proto.RegisterType((*CS_PlayerMove)(nil), "protomsg.CS_PlayerMove")
	proto.RegisterType((*CS_PlayerAttack)(nil), "protomsg.CS_PlayerAttack")
	proto.RegisterType((*CS_PlayerSkill)(nil), "protomsg.CS_PlayerSkill")
	proto.RegisterType((*UnitDatas)(nil), "protomsg.UnitDatas")
	proto.RegisterType((*MsgUserEnterScene)(nil), "protomsg.MsgUserEnterScene")
	proto.RegisterType((*SC_Update)(nil), "protomsg.SC_Update")
	proto.RegisterType((*SC_NewScene)(nil), "protomsg.SC_NewScene")
	proto.RegisterType((*SC_Logined)(nil), "protomsg.SC_Logined")
}

func init() { proto.RegisterFile("protobuf/msg.proto", fileDescriptor_617aaba9bf0c07c7) }

var fileDescriptor_617aaba9bf0c07c7 = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5d, 0x6e, 0xe3, 0x36,
	0x10, 0x86, 0x64, 0x2b, 0xb1, 0x27, 0x89, 0x37, 0xcb, 0xa4, 0x05, 0xb1, 0x2d, 0x16, 0x86, 0xb0,
	0x28, 0xf2, 0x94, 0x00, 0xed, 0x09, 0x52, 0x29, 0xbb, 0x11, 0xb0, 0x4a, 0x65, 0x29, 0xc1, 0x26,
	0x4f, 0x2e, 0x23, 0x73, 0x15, 0x22, 0x96, 0x64, 0x90, 0x8c, 0x37, 0x01, 0xfa, 0xda, 0x3b, 0xf4,
	0x0e, 0x3d, 0x40, 0x4f, 0xd0, 0x7b, 0x15, 0x24, 0xf5, 0xeb, 0x1a, 0x79, 0x12, 0xbf, 0x6f, 0x86,
	0x33, 0xc3, 0x21, 0xe7, 0x13, 0xa0, 0x15, 0x2f, 0x65, 0x79, 0xff, 0xf4, 0xf5, 0x2c, 0x17, 0xd9,
	0xa9, 0x06, 0x68, 0xa4, 0x3f, 0xb9, 0xc8, 0xdc, 0x3f, 0x2d, 0xd8, 0x0d, 0x45, 0xf6, 0x2b, 0x11,
	0x14, 0xbd, 0x83, 0x51, 0x58, 0x2e, 0xe8, 0xf5, 0xcb, 0x8a, 0x62, 0x6b, 0x6a, 0x9d, 0x8c, 0xe3,
	0x06, 0xa3, 0x43, 0x18, 0xdc, 0xb0, 0x05, 0xb6, 0xa7, 0xd6, 0x89, 0x13, 0xab, 0x25, 0xc2, 0x7a,
	0xa3, 0x76, 0x1e, 0x68, 0xe7, 0x1a, 0xa2, 0x1f, 0x61, 0xec, 0x95, 0x45, 0x41, 0x53, 0x19, 0x2c,
	0xf0, 0x50, 0xef, 0x68, 0x09, 0x74, 0x0c, 0x8e, 0x4f, 0x24, 0x11, 0xd8, 0x99, 0x5a, 0x27, 0xfb,
	0xb1, 0x01, 0xee, 0x19, 0xbc, 0x0d, 0x45, 0x16, 0xd3, 0x8c, 0x09, 0x49, 0xf9, 0x75, 0xf9, 0x89,
	0xc8, 0x57, 0x0b, 0x72, 0x7f, 0x87, 0x43, 0x2f, 0x99, 0x87, 0x22, 0x9b, 0x3d, 0xb1, 0xf4, 0xf1,
	0x73, 0x99, 0xb1, 0x42, 0xf9, 0x47, 0x4b, 0x22, 0xbf, 0x96, 0x3c, 0xaf, 0xfd, 0x6b, 0xac, 0x8a,
	0x0a, 0x49, 0xfa, 0xc0, 0x0a, 0x5a, 0x1d, 0x63, 0x1c, 0xb7, 0x04, 0x42, 0x30, 0xbc, 0x22, 0x79,
	0x7d, 0x12, 0xbd, 0x76, 0xff, 0xb2, 0x60, 0xe4, 0x25, 0x73, 0x13, 0x7a, 0x0a, 0x7b, 0xd1, 0x43,
	0x59, 0xd0, 0xe2, 0x29, 0xbf, 0xa7, 0xbc, 0x8a, 0xde, 0xa5, 0x7a, 0xc9, 0xed, 0xd7, 0x92, 0x0f,
	0x36, 0x93, 0xff, 0x00, 0xe3, 0x2f, 0x34, 0x7d, 0x20, 0x72, 0xce, 0x4c, 0xbf, 0xc6, 0xf1, 0xc8,
	0x10, 0xc1, 0x02, 0x1d, 0x81, 0x33, 0x9b, 0x29, 0x83, 0x63, 0x4a, 0x9b, 0xcd, 0x82, 0x85, 0xfb,
	0x05, 0x0e, 0xbc, 0x64, 0x1e, 0x2d, 0xc9, 0x0b, 0xe5, 0x61, 0xb9, 0xd6, 0xd7, 0x13, 0xf8, 0x02,
	0x5b, 0xd3, 0x81, 0xba, 0x9e, 0xc0, 0x17, 0x68, 0x1f, 0xac, 0x5b, 0x5d, 0x87, 0x1d, 0x5b, 0xb7,
	0x0a, 0xdd, 0xe9, 0xc4, 0x76, 0x6c, 0xdd, 0xa9, 0xab, 0x0b, 0x44, 0x22, 0x09, 0x97, 0x3a, 0xdd,
	0x28, 0xae, 0xa1, 0xfb, 0x09, 0xde, 0x34, 0x81, 0xcf, 0xa5, 0x24, 0xe9, 0xe3, 0x96, 0xd0, 0x2e,
	0xec, 0x5f, 0x13, 0x9e, 0x51, 0x79, 0x53, 0x30, 0x19, 0xf8, 0xd5, 0xa3, 0xe8, 0x71, 0xee, 0x1f,
	0x30, 0x69, 0x02, 0x25, 0x8f, 0x6c, 0xb9, 0x6c, 0xe3, 0x58, 0x75, 0x1c, 0x0c, 0xbb, 0xda, 0xd4,
	0x84, 0xa8, 0xa1, 0x29, 0x7e, 0xd0, 0x2b, 0x7e, 0x58, 0x17, 0xbf, 0x99, 0xdd, 0xd9, 0x92, 0xfd,
	0x9f, 0x21, 0x8c, 0xd5, 0x52, 0xbf, 0xad, 0xe6, 0x72, 0xad, 0xf6, 0x72, 0xd5, 0x2b, 0xfc, 0x4c,
	0xd7, 0x74, 0x59, 0x65, 0x36, 0x00, 0x4d, 0xc0, 0xbe, 0x8c, 0x74, 0x62, 0x27, 0xb6, 0x2f, 0x23,
	0x85, 0xc3, 0xa8, 0x7a, 0xc2, 0x76, 0x18, 0x99, 0xba, 0x9c, 0x5e, 0x5d, 0x3b, 0x75, 0x5d, 0x13,
	0xb0, 0x03, 0x1f, 0xef, 0x1a, 0xdf, 0xc0, 0xef, 0x3d, 0xde, 0xd1, 0xc6, 0x34, 0x1d, 0x83, 0x13,
	0x92, 0xe7, 0xcb, 0x08, 0x8f, 0x4d, 0x76, 0x0d, 0x2a, 0x36, 0x8c, 0x30, 0x34, 0x6c, 0x18, 0xa1,
	0xf7, 0x00, 0x17, 0xcf, 0x2b, 0xca, 0x19, 0x2d, 0x52, 0x8a, 0xf7, 0xb4, 0xa9, 0xc3, 0xa0, 0x0f,
	0x70, 0x10, 0x92, 0xe7, 0x8e, 0xcb, 0xbe, 0x76, 0xe9, 0x93, 0xd5, 0x4c, 0x4a, 0x5e, 0xaa, 0x6e,
	0x1f, 0x34, 0x33, 0x69, 0x08, 0x15, 0xe3, 0xbc, 0x60, 0x79, 0x29, 0x4b, 0x9e, 0x48, 0x22, 0x29,
	0x9e, 0x98, 0x18, 0x3d, 0x52, 0x55, 0x62, 0xde, 0xc4, 0x35, 0xcb, 0x29, 0x7e, 0xa3, 0x0f, 0xde,
	0x61, 0x94, 0xdd, 0x67, 0x9c, 0xa6, 0x92, 0x95, 0xc5, 0x2d, 0x3e, 0x34, 0xf6, 0x96, 0xe9, 0xd9,
	0xef, 0xf0, 0xdb, 0x0d, 0xfb, 0x9d, 0xea, 0x98, 0xba, 0x34, 0xdd, 0x31, 0xa4, 0x0b, 0x68, 0x30,
	0x3a, 0x05, 0x64, 0x32, 0x9d, 0xa7, 0x2b, 0x72, 0xcf, 0x96, 0x4c, 0x32, 0x2a, 0xf0, 0x91, 0xf6,
	0xda, 0x62, 0x69, 0x6b, 0x55, 0x3d, 0xc7, 0xc7, 0xa6, 0x6b, 0x2d, 0x83, 0xbe, 0x87, 0x9d, 0x40,
	0x84, 0x84, 0x15, 0xf8, 0x3b, 0x6d, 0xab, 0x90, 0xfb, 0xaf, 0xa5, 0x85, 0xe8, 0x46, 0x50, 0x7e,
	0x51, 0x48, 0xca, 0x93, 0x94, 0x16, 0x8d, 0xfa, 0x59, 0xad, 0xfa, 0xf5, 0x34, 0xce, 0xde, 0xd4,
	0xb8, 0x0f, 0x70, 0x90, 0xf0, 0x34, 0xa1, 0x7c, 0x4d, 0x79, 0x47, 0x57, 0xfa, 0x24, 0xfa, 0x09,
	0x26, 0x3e, 0x15, 0xb2, 0xe3, 0x66, 0x86, 0x7f, 0x83, 0x55, 0xb9, 0x74, 0x19, 0xda, 0xc5, 0xc8,
	0x40, 0x4b, 0xb4, 0x7a, 0xba, 0xd3, 0xd5, 0xd3, 0xbf, 0x2d, 0x18, 0x27, 0xde, 0xfc, 0x66, 0xb5,
	0xa8, 0x84, 0xd4, 0x7b, 0xe2, 0x1f, 0x79, 0x3d, 0x05, 0x4e, 0xdc, 0x60, 0x74, 0x06, 0xa3, 0x2b,
	0xfa, 0x4d, 0x35, 0x5a, 0x60, 0x7b, 0x3a, 0x38, 0xd9, 0xfb, 0xf9, 0xe8, 0xb4, 0xfe, 0x3d, 0x9c,
	0x36, 0x43, 0x14, 0x37, 0x4e, 0x6a, 0xc3, 0x6f, 0xcb, 0x85, 0xd9, 0x30, 0x78, 0x65, 0x43, 0xed,
	0xa4, 0xb4, 0x33, 0xa6, 0x79, 0xb9, 0xa6, 0x66, 0xcf, 0x50, 0x2b, 0x49, 0x97, 0x72, 0x5f, 0x60,
	0x2f, 0xf1, 0xe6, 0x57, 0xf4, 0x9b, 0x69, 0xf7, 0xb6, 0x81, 0x7d, 0x07, 0x23, 0xa5, 0xc4, 0xe9,
	0xc7, 0x95, 0xa8, 0xfa, 0xdd, 0xe0, 0xde, 0xf1, 0x06, 0x1b, 0xc7, 0x7b, 0x0f, 0xf0, 0xbf, 0x06,
	0x77, 0x18, 0xf7, 0x12, 0x20, 0xf1, 0x8c, 0xc8, 0x53, 0xfd, 0x1f, 0xf0, 0xd4, 0x83, 0x31, 0x4d,
	0xd2, 0xeb, 0x2d, 0xbf, 0xbe, 0x63, 0x70, 0x2e, 0x38, 0x2f, 0x79, 0x75, 0xad, 0x06, 0xdc, 0xef,
	0xe8, 0x26, 0xfc, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0xcf, 0xe9, 0xee, 0x71, 0x07,
	0x00, 0x00,
}

// Copyright (c) 2024-present,  Teamgram Studio (https://teamgram.net).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: rpc_metadata.proto

package metadata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TakeoutMessageRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinId         int32                  `protobuf:"varint,1,opt,name=min_id,json=minId,proto3" json:"min_id,omitempty"`
	MaxId         int32                  `protobuf:"varint,2,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TakeoutMessageRange) Reset() {
	*x = TakeoutMessageRange{}
	mi := &file_rpc_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TakeoutMessageRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeoutMessageRange) ProtoMessage() {}

func (x *TakeoutMessageRange) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeoutMessageRange.ProtoReflect.Descriptor instead.
func (*TakeoutMessageRange) Descriptor() ([]byte, []int) {
	return file_rpc_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *TakeoutMessageRange) GetMinId() int32 {
	if x != nil {
		return x.MinId
	}
	return 0
}

func (x *TakeoutMessageRange) GetMaxId() int32 {
	if x != nil {
		return x.MaxId
	}
	return 0
}

type Takeout struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Range         *TakeoutMessageRange   `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Takeout) Reset() {
	*x = Takeout{}
	mi := &file_rpc_metadata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Takeout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Takeout) ProtoMessage() {}

func (x *Takeout) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_metadata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Takeout.ProtoReflect.Descriptor instead.
func (*Takeout) Descriptor() ([]byte, []int) {
	return file_rpc_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *Takeout) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Takeout) GetRange() *TakeoutMessageRange {
	if x != nil {
		return x.Range
	}
	return nil
}

type RpcMetadata struct {
	state       protoimpl.MessageState `protogen:"open.v1"`
	ServerId    string                 `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ClientAddr  string                 `protobuf:"bytes,3,opt,name=client_addr,json=clientAddr,proto3" json:"client_addr,omitempty"`
	AuthId      int64                  `protobuf:"varint,4,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	SessionId   int64                  `protobuf:"varint,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ReceiveTime int64                  `protobuf:"varint,8,opt,name=receive_time,json=receiveTime,proto3" json:"receive_time,omitempty"`
	// 用户ID
	UserId      int64 `protobuf:"varint,11,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientMsgId int64 `protobuf:"varint,12,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`
	IsBot       bool  `protobuf:"varint,13,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	// mtproto
	Layer  int32  `protobuf:"varint,14,opt,name=layer,proto3" json:"layer,omitempty"`
	Client string `protobuf:"bytes,15,opt,name=client,proto3" json:"client,omitempty"`
	// admin
	IsAdmin bool `protobuf:"varint,16,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// talkeout
	Takeout *Takeout `protobuf:"bytes,20,opt,name=takeout,proto3" json:"takeout,omitempty"`
	// langpack
	Langpack string `protobuf:"bytes,21,opt,name=langpack,proto3" json:"langpack,omitempty"`
	// perm_auth_key_id
	PermAuthKeyId int64 `protobuf:"varint,22,opt,name=perm_auth_key_id,json=permAuthKeyId,proto3" json:"perm_auth_key_id,omitempty"`
	// lang_code
	LangCode      string `protobuf:"bytes,23,opt,name=lang_code,json=langCode,proto3" json:"lang_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RpcMetadata) Reset() {
	*x = RpcMetadata{}
	mi := &file_rpc_metadata_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RpcMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcMetadata) ProtoMessage() {}

func (x *RpcMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_metadata_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcMetadata.ProtoReflect.Descriptor instead.
func (*RpcMetadata) Descriptor() ([]byte, []int) {
	return file_rpc_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *RpcMetadata) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *RpcMetadata) GetClientAddr() string {
	if x != nil {
		return x.ClientAddr
	}
	return ""
}

func (x *RpcMetadata) GetAuthId() int64 {
	if x != nil {
		return x.AuthId
	}
	return 0
}

func (x *RpcMetadata) GetSessionId() int64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *RpcMetadata) GetReceiveTime() int64 {
	if x != nil {
		return x.ReceiveTime
	}
	return 0
}

func (x *RpcMetadata) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RpcMetadata) GetClientMsgId() int64 {
	if x != nil {
		return x.ClientMsgId
	}
	return 0
}

func (x *RpcMetadata) GetIsBot() bool {
	if x != nil {
		return x.IsBot
	}
	return false
}

func (x *RpcMetadata) GetLayer() int32 {
	if x != nil {
		return x.Layer
	}
	return 0
}

func (x *RpcMetadata) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *RpcMetadata) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *RpcMetadata) GetTakeout() *Takeout {
	if x != nil {
		return x.Takeout
	}
	return nil
}

func (x *RpcMetadata) GetLangpack() string {
	if x != nil {
		return x.Langpack
	}
	return ""
}

func (x *RpcMetadata) GetPermAuthKeyId() int64 {
	if x != nil {
		return x.PermAuthKeyId
	}
	return 0
}

func (x *RpcMetadata) GetLangCode() string {
	if x != nil {
		return x.LangCode
	}
	return ""
}

var File_rpc_metadata_proto protoreflect.FileDescriptor

var file_rpc_metadata_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x43,
	0x0a, 0x13, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06,
	0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61,
	0x78, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x07, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33,
	0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x05, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x22, 0xd2, 0x03, 0x0a, 0x0b, 0x52, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x62, 0x6f, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x6f, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x74, 0x61, 0x6b,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x74,
	0x61, 0x6b, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x70, 0x61,
	0x63, 0x6b, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x70, 0x61,
	0x63, 0x6b, 0x12, 0x27, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x65,
	0x72, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x67, 0x72, 0x61, 0x6d, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_rpc_metadata_proto_rawDescOnce sync.Once
	file_rpc_metadata_proto_rawDescData []byte
)

func file_rpc_metadata_proto_rawDescGZIP() []byte {
	file_rpc_metadata_proto_rawDescOnce.Do(func() {
		file_rpc_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rpc_metadata_proto_rawDesc), len(file_rpc_metadata_proto_rawDesc)))
	})
	return file_rpc_metadata_proto_rawDescData
}

var file_rpc_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_metadata_proto_goTypes = []any{
	(*TakeoutMessageRange)(nil), // 0: metadata.TakeoutMessageRange
	(*Takeout)(nil),             // 1: metadata.Takeout
	(*RpcMetadata)(nil),         // 2: metadata.RpcMetadata
}
var file_rpc_metadata_proto_depIdxs = []int32{
	0, // 0: metadata.Takeout.range:type_name -> metadata.TakeoutMessageRange
	1, // 1: metadata.RpcMetadata.takeout:type_name -> metadata.Takeout
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_metadata_proto_init() }
func file_rpc_metadata_proto_init() {
	if File_rpc_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rpc_metadata_proto_rawDesc), len(file_rpc_metadata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_metadata_proto_goTypes,
		DependencyIndexes: file_rpc_metadata_proto_depIdxs,
		MessageInfos:      file_rpc_metadata_proto_msgTypes,
	}.Build()
	File_rpc_metadata_proto = out.File
	file_rpc_metadata_proto_goTypes = nil
	file_rpc_metadata_proto_depIdxs = nil
}

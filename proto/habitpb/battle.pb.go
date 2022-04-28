// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/battle.proto

package habitpb

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

type Battle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitOne *Habit `protobuf:"bytes,1,opt,name=habit_one,json=habitOne,proto3" json:"habit_one,omitempty"`
	HabitTwo *Habit `protobuf:"bytes,2,opt,name=habit_two,json=habitTwo,proto3" json:"habit_two,omitempty"`
	Code     string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Winner   string `protobuf:"bytes,4,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *Battle) Reset() {
	*x = Battle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_battle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Battle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Battle) ProtoMessage() {}

func (x *Battle) ProtoReflect() protoreflect.Message {
	mi := &file_proto_battle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Battle.ProtoReflect.Descriptor instead.
func (*Battle) Descriptor() ([]byte, []int) {
	return file_proto_battle_proto_rawDescGZIP(), []int{0}
}

func (x *Battle) GetHabitOne() *Habit {
	if x != nil {
		return x.HabitOne
	}
	return nil
}

func (x *Battle) GetHabitTwo() *Habit {
	if x != nil {
		return x.HabitTwo
	}
	return nil
}

func (x *Battle) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Battle) GetWinner() string {
	if x != nil {
		return x.Winner
	}
	return ""
}

type BattleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	HabitID  string `protobuf:"bytes,2,opt,name=habitID,proto3" json:"habitID,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *BattleRequest) Reset() {
	*x = BattleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_battle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleRequest) ProtoMessage() {}

func (x *BattleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_battle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleRequest.ProtoReflect.Descriptor instead.
func (*BattleRequest) Descriptor() ([]byte, []int) {
	return file_proto_battle_proto_rawDescGZIP(), []int{1}
}

func (x *BattleRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BattleRequest) GetHabitID() string {
	if x != nil {
		return x.HabitID
	}
	return ""
}

func (x *BattleRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type BattleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Ok      bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Pending bool   `protobuf:"varint,4,opt,name=pending,proto3" json:"pending,omitempty"`
}

func (x *BattleResponse) Reset() {
	*x = BattleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_battle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleResponse) ProtoMessage() {}

func (x *BattleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_battle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleResponse.ProtoReflect.Descriptor instead.
func (*BattleResponse) Descriptor() ([]byte, []int) {
	return file_proto_battle_proto_rawDescGZIP(), []int{2}
}

func (x *BattleResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BattleResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *BattleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BattleResponse) GetPending() bool {
	if x != nil {
		return x.Pending
	}
	return false
}

type BattleAssociationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitID  string `protobuf:"bytes,1,opt,name=habitID,proto3" json:"habitID,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *BattleAssociationsRequest) Reset() {
	*x = BattleAssociationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_battle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleAssociationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleAssociationsRequest) ProtoMessage() {}

func (x *BattleAssociationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_battle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleAssociationsRequest.ProtoReflect.Descriptor instead.
func (*BattleAssociationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_battle_proto_rawDescGZIP(), []int{3}
}

func (x *BattleAssociationsRequest) GetHabitID() string {
	if x != nil {
		return x.HabitID
	}
	return ""
}

func (x *BattleAssociationsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type BattleAssociationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Codes []string `protobuf:"bytes,2,rep,name=codes,proto3" json:"codes,omitempty"`
}

func (x *BattleAssociationsResponse) Reset() {
	*x = BattleAssociationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_battle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleAssociationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleAssociationsResponse) ProtoMessage() {}

func (x *BattleAssociationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_battle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleAssociationsResponse.ProtoReflect.Descriptor instead.
func (*BattleAssociationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_battle_proto_rawDescGZIP(), []int{4}
}

func (x *BattleAssociationsResponse) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_proto_battle_proto protoreflect.FileDescriptor

var file_proto_battle_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x06, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x09, 0x68, 0x61, 0x62, 0x69, 0x74, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x08, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x09, 0x68, 0x61, 0x62, 0x69, 0x74, 0x5f,
	0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x52, 0x08, 0x68, 0x61, 0x62, 0x69, 0x74, 0x54, 0x77, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x68, 0x61, 0x62, 0x69, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x68, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x51, 0x0a, 0x19,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x62,
	0x69, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x32, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x68, 0x61, 0x62, 0x69, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_battle_proto_rawDescOnce sync.Once
	file_proto_battle_proto_rawDescData = file_proto_battle_proto_rawDesc
)

func file_proto_battle_proto_rawDescGZIP() []byte {
	file_proto_battle_proto_rawDescOnce.Do(func() {
		file_proto_battle_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_battle_proto_rawDescData)
	})
	return file_proto_battle_proto_rawDescData
}

var file_proto_battle_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_battle_proto_goTypes = []interface{}{
	(*Battle)(nil),                     // 0: Battle
	(*BattleRequest)(nil),              // 1: BattleRequest
	(*BattleResponse)(nil),             // 2: BattleResponse
	(*BattleAssociationsRequest)(nil),  // 3: BattleAssociationsRequest
	(*BattleAssociationsResponse)(nil), // 4: BattleAssociationsResponse
	(*Habit)(nil),                      // 5: Habit
}
var file_proto_battle_proto_depIdxs = []int32{
	5, // 0: Battle.habit_one:type_name -> Habit
	5, // 1: Battle.habit_two:type_name -> Habit
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_battle_proto_init() }
func file_proto_battle_proto_init() {
	if File_proto_battle_proto != nil {
		return
	}
	file_proto_habit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_battle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Battle); i {
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
		file_proto_battle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleRequest); i {
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
		file_proto_battle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleResponse); i {
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
		file_proto_battle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleAssociationsRequest); i {
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
		file_proto_battle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleAssociationsResponse); i {
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
			RawDescriptor: file_proto_battle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_battle_proto_goTypes,
		DependencyIndexes: file_proto_battle_proto_depIdxs,
		MessageInfos:      file_proto_battle_proto_msgTypes,
	}.Build()
	File_proto_battle_proto = out.File
	file_proto_battle_proto_rawDesc = nil
	file_proto_battle_proto_goTypes = nil
	file_proto_battle_proto_depIdxs = nil
}

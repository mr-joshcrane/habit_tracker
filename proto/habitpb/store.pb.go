// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: store.proto

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

type Habit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streak        int32  `protobuf:"varint,1,opt,name=streak,proto3" json:"streak,omitempty"`
	LastPerformed int64  `protobuf:"varint,2,opt,name=last_performed,json=lastPerformed,proto3" json:"last_performed,omitempty"`
	HabitName     string `protobuf:"bytes,3,opt,name=habit_name,json=habitName,proto3" json:"habit_name,omitempty"`
	User          string `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Habit) Reset() {
	*x = Habit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habit) ProtoMessage() {}

func (x *Habit) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habit.ProtoReflect.Descriptor instead.
func (*Habit) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{0}
}

func (x *Habit) GetStreak() int32 {
	if x != nil {
		return x.Streak
	}
	return 0
}

func (x *Habit) GetLastPerformed() int64 {
	if x != nil {
		return x.LastPerformed
	}
	return 0
}

func (x *Habit) GetHabitName() string {
	if x != nil {
		return x.HabitName
	}
	return ""
}

func (x *Habit) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type Habits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habits map[string]*Habit `protobuf:"bytes,1,rep,name=habits,proto3" json:"habits,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Habits) Reset() {
	*x = Habits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habits) ProtoMessage() {}

func (x *Habits) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habits.ProtoReflect.Descriptor instead.
func (*Habits) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{1}
}

func (x *Habits) GetHabits() map[string]*Habit {
	if x != nil {
		return x.Habits
	}
	return nil
}

type UpdateHabitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habit *Habit `protobuf:"bytes,1,opt,name=habit,proto3" json:"habit,omitempty"`
}

func (x *UpdateHabitRequest) Reset() {
	*x = UpdateHabitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHabitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHabitRequest) ProtoMessage() {}

func (x *UpdateHabitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHabitRequest.ProtoReflect.Descriptor instead.
func (*UpdateHabitRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateHabitRequest) GetHabit() *Habit {
	if x != nil {
		return x.Habit
	}
	return nil
}

type UpdateHabitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateHabitResponse) Reset() {
	*x = UpdateHabitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHabitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHabitResponse) ProtoMessage() {}

func (x *UpdateHabitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHabitResponse.ProtoReflect.Descriptor instead.
func (*UpdateHabitResponse) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateHabitResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *UpdateHabitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetHabitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habitname string `protobuf:"bytes,1,opt,name=habitname,proto3" json:"habitname,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetHabitRequest) Reset() {
	*x = GetHabitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHabitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHabitRequest) ProtoMessage() {}

func (x *GetHabitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHabitRequest.ProtoReflect.Descriptor instead.
func (*GetHabitRequest) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{4}
}

func (x *GetHabitRequest) GetHabitname() string {
	if x != nil {
		return x.Habitname
	}
	return ""
}

func (x *GetHabitRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetHabitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Habit *Habit `protobuf:"bytes,2,opt,name=habit,proto3" json:"habit,omitempty"`
}

func (x *GetHabitResponse) Reset() {
	*x = GetHabitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHabitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHabitResponse) ProtoMessage() {}

func (x *GetHabitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_store_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHabitResponse.ProtoReflect.Descriptor instead.
func (*GetHabitResponse) Descriptor() ([]byte, []int) {
	return file_store_proto_rawDescGZIP(), []int{5}
}

func (x *GetHabitResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetHabitResponse) GetHabit() *Habit {
	if x != nil {
		return x.Habit
	}
	return nil
}

var File_store_proto protoreflect.FileDescriptor

var file_store_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a,
	0x05, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6b, 0x12, 0x25,
	0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x62, 0x69, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x78, 0x0a, 0x06, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x2e, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x61, 0x62, 0x69, 0x74, 0x73, 0x1a,
	0x41, 0x0a, 0x0b, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x32, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52,
	0x05, 0x68, 0x61, 0x62, 0x69, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52,
	0x05, 0x68, 0x61, 0x62, 0x69, 0x74, 0x32, 0x7d, 0x0a, 0x0c, 0x48, 0x61, 0x62, 0x69, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62,
	0x69, 0x74, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x13, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_proto_rawDescOnce sync.Once
	file_store_proto_rawDescData = file_store_proto_rawDesc
)

func file_store_proto_rawDescGZIP() []byte {
	file_store_proto_rawDescOnce.Do(func() {
		file_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_proto_rawDescData)
	})
	return file_store_proto_rawDescData
}

var file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_store_proto_goTypes = []interface{}{
	(*Habit)(nil),               // 0: Habit
	(*Habits)(nil),              // 1: Habits
	(*UpdateHabitRequest)(nil),  // 2: UpdateHabitRequest
	(*UpdateHabitResponse)(nil), // 3: UpdateHabitResponse
	(*GetHabitRequest)(nil),     // 4: GetHabitRequest
	(*GetHabitResponse)(nil),    // 5: GetHabitResponse
	nil,                         // 6: Habits.HabitsEntry
}
var file_store_proto_depIdxs = []int32{
	6, // 0: Habits.habits:type_name -> Habits.HabitsEntry
	0, // 1: UpdateHabitRequest.habit:type_name -> Habit
	0, // 2: GetHabitResponse.habit:type_name -> Habit
	0, // 3: Habits.HabitsEntry.value:type_name -> Habit
	4, // 4: HabitService.GetHabit:input_type -> GetHabitRequest
	2, // 5: HabitService.UpdateHabit:input_type -> UpdateHabitRequest
	5, // 6: HabitService.GetHabit:output_type -> GetHabitResponse
	3, // 7: HabitService.UpdateHabit:output_type -> UpdateHabitResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_proto_init() }
func file_store_proto_init() {
	if File_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Habit); i {
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
		file_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Habits); i {
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
		file_store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHabitRequest); i {
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
		file_store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHabitResponse); i {
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
		file_store_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHabitRequest); i {
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
		file_store_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHabitResponse); i {
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
			RawDescriptor: file_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_store_proto_goTypes,
		DependencyIndexes: file_store_proto_depIdxs,
		MessageInfos:      file_store_proto_msgTypes,
	}.Build()
	File_store_proto = out.File
	file_store_proto_rawDesc = nil
	file_store_proto_goTypes = nil
	file_store_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/store.proto

package habitpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_store_proto protoreflect.FileDescriptor

var file_proto_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc3, 0x02, 0x0a, 0x0c, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x48, 0x61, 0x62, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0a, 0x4a,
	0x6f, 0x69, 0x6e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x68, 0x61, 0x62, 0x69, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_store_proto_goTypes = []interface{}{
	(*PerformHabitRequest)(nil),        // 0: PerformHabitRequest
	(*ListHabitsRequest)(nil),          // 1: ListHabitsRequest
	(*BattleRequest)(nil),              // 2: BattleRequest
	(*BattleAssociationsRequest)(nil),  // 3: BattleAssociationsRequest
	(*PerformHabitResponse)(nil),       // 4: PerformHabitResponse
	(*ListHabitsResponse)(nil),         // 5: ListHabitsResponse
	(*BattleResponse)(nil),             // 6: BattleResponse
	(*BattleAssociationsResponse)(nil), // 7: BattleAssociationsResponse
}
var file_proto_store_proto_depIdxs = []int32{
	0, // 0: HabitService.PerformHabit:input_type -> PerformHabitRequest
	1, // 1: HabitService.DisplayHabits:input_type -> ListHabitsRequest
	2, // 2: HabitService.RegisterBattle:input_type -> BattleRequest
	2, // 3: HabitService.JoinBattle:input_type -> BattleRequest
	3, // 4: HabitService.GetBattleAssociations:input_type -> BattleAssociationsRequest
	4, // 5: HabitService.PerformHabit:output_type -> PerformHabitResponse
	5, // 6: HabitService.DisplayHabits:output_type -> ListHabitsResponse
	6, // 7: HabitService.RegisterBattle:output_type -> BattleResponse
	6, // 8: HabitService.JoinBattle:output_type -> BattleResponse
	7, // 9: HabitService.GetBattleAssociations:output_type -> BattleAssociationsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_store_proto_init() }
func file_proto_store_proto_init() {
	if File_proto_store_proto != nil {
		return
	}
	file_proto_battle_proto_init()
	file_proto_habit_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_store_proto_goTypes,
		DependencyIndexes: file_proto_store_proto_depIdxs,
	}.Build()
	File_proto_store_proto = out.File
	file_proto_store_proto_rawDesc = nil
	file_proto_store_proto_goTypes = nil
	file_proto_store_proto_depIdxs = nil
}

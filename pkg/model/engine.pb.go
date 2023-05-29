// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pb/model/engine.proto

package model

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

// Logic for how energy gets added to a target
type EnergyAdd int32

const (
	EnergyAdd_INVALID EnergyAdd = 0
	// The energy added can be increased by the target's ERR
	EnergyAdd_ENERGY_ADD_AMOUNT EnergyAdd = 1
	// The energy added is static and does not get modified by the target's ERR
	EnergyAdd_ENERGY_ADD_AMOUNT_FIXED EnergyAdd = 2
)

// Enum value maps for EnergyAdd.
var (
	EnergyAdd_name = map[int32]string{
		0: "INVALID",
		1: "ENERGY_ADD_AMOUNT",
		2: "ENERGY_ADD_AMOUNT_FIXED",
	}
	EnergyAdd_value = map[string]int32{
		"INVALID":                 0,
		"ENERGY_ADD_AMOUNT":       1,
		"ENERGY_ADD_AMOUNT_FIXED": 2,
	}
)

func (x EnergyAdd) Enum() *EnergyAdd {
	p := new(EnergyAdd)
	*p = x
	return p
}

func (x EnergyAdd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnergyAdd) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_model_engine_proto_enumTypes[0].Descriptor()
}

func (EnergyAdd) Type() protoreflect.EnumType {
	return &file_pb_model_engine_proto_enumTypes[0]
}

func (x EnergyAdd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnergyAdd.Descriptor instead.
func (EnergyAdd) EnumDescriptor() ([]byte, []int) {
	return file_pb_model_engine_proto_rawDescGZIP(), []int{0}
}

// Logic for how to modify the gauge of a target (turn advance/delay)
type ModifyGauge int32

const (
	ModifyGauge_INVALID_MODIFY_GAUGE ModifyGauge = 0
	// The modify amount is a percentage [0,1] that will modify gauge against the base gauge
	ModifyGauge_MODIFY_GAUGE_NORMALIZED ModifyGauge = 1
	// The modify amount is AV. The AV will be normalized by target speed
	ModifyGauge_MODIFY_GAUGE_AV ModifyGauge = 2
)

// Enum value maps for ModifyGauge.
var (
	ModifyGauge_name = map[int32]string{
		0: "INVALID_MODIFY_GAUGE",
		1: "MODIFY_GAUGE_NORMALIZED",
		2: "MODIFY_GAUGE_AV",
	}
	ModifyGauge_value = map[string]int32{
		"INVALID_MODIFY_GAUGE":    0,
		"MODIFY_GAUGE_NORMALIZED": 1,
		"MODIFY_GAUGE_AV":         2,
	}
)

func (x ModifyGauge) Enum() *ModifyGauge {
	p := new(ModifyGauge)
	*p = x
	return p
}

func (x ModifyGauge) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModifyGauge) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_model_engine_proto_enumTypes[1].Descriptor()
}

func (ModifyGauge) Type() protoreflect.EnumType {
	return &file_pb_model_engine_proto_enumTypes[1]
}

func (x ModifyGauge) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModifyGauge.Descriptor instead.
func (ModifyGauge) EnumDescriptor() ([]byte, []int) {
	return file_pb_model_engine_proto_rawDescGZIP(), []int{1}
}

type StatusType int32

const (
	StatusType_UNKNOWN_STATUS StatusType = 0
	StatusType_STATUS_BUFF    StatusType = 1
	StatusType_STATUS_DEBUFF  StatusType = 2
)

// Enum value maps for StatusType.
var (
	StatusType_name = map[int32]string{
		0: "UNKNOWN_STATUS",
		1: "STATUS_BUFF",
		2: "STATUS_DEBUFF",
	}
	StatusType_value = map[string]int32{
		"UNKNOWN_STATUS": 0,
		"STATUS_BUFF":    1,
		"STATUS_DEBUFF":  2,
	}
)

func (x StatusType) Enum() *StatusType {
	p := new(StatusType)
	*p = x
	return p
}

func (x StatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_model_engine_proto_enumTypes[2].Descriptor()
}

func (StatusType) Type() protoreflect.EnumType {
	return &file_pb_model_engine_proto_enumTypes[2]
}

func (x StatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusType.Descriptor instead.
func (StatusType) EnumDescriptor() ([]byte, []int) {
	return file_pb_model_engine_proto_rawDescGZIP(), []int{2}
}

var File_pb_model_engine_proto protoreflect.FileDescriptor

var file_pb_model_engine_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2a, 0x4c,
	0x0a, 0x09, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x41, 0x64, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4e, 0x45, 0x52,
	0x47, 0x59, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x4d, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x45, 0x52, 0x47, 0x59, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x41, 0x4d,
	0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x58, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x59, 0x0a, 0x0b,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x47, 0x61, 0x75, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x47, 0x41,
	0x55, 0x47, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f,
	0x47, 0x41, 0x55, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x49, 0x5a, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x47, 0x41, 0x55,
	0x47, 0x45, 0x5f, 0x41, 0x56, 0x10, 0x02, 0x2a, 0x44, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x42, 0x55, 0x46, 0x46, 0x10, 0x02, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x69,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x2f, 0x73, 0x72, 0x73, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_model_engine_proto_rawDescOnce sync.Once
	file_pb_model_engine_proto_rawDescData = file_pb_model_engine_proto_rawDesc
)

func file_pb_model_engine_proto_rawDescGZIP() []byte {
	file_pb_model_engine_proto_rawDescOnce.Do(func() {
		file_pb_model_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_model_engine_proto_rawDescData)
	})
	return file_pb_model_engine_proto_rawDescData
}

var file_pb_model_engine_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_pb_model_engine_proto_goTypes = []interface{}{
	(EnergyAdd)(0),   // 0: model.EnergyAdd
	(ModifyGauge)(0), // 1: model.ModifyGauge
	(StatusType)(0),  // 2: model.StatusType
}
var file_pb_model_engine_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_model_engine_proto_init() }
func file_pb_model_engine_proto_init() {
	if File_pb_model_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_model_engine_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_model_engine_proto_goTypes,
		DependencyIndexes: file_pb_model_engine_proto_depIdxs,
		EnumInfos:         file_pb_model_engine_proto_enumTypes,
	}.Build()
	File_pb_model_engine_proto = out.File
	file_pb_model_engine_proto_rawDesc = nil
	file_pb_model_engine_proto_goTypes = nil
	file_pb_model_engine_proto_depIdxs = nil
}
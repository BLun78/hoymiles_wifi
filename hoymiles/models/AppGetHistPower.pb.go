// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: protos/AppGetHistPower.proto

package models

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

type AppGetHistPowerReqDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControlPoint  int32  `protobuf:"varint,1,opt,name=control_point,json=controlPoint,proto3" json:"control_point,omitempty"`    // Control point identifier
	Offset        int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`                                    // Offset value
	RequestedTime uint32 `protobuf:"varint,3,opt,name=requested_time,json=requestedTime,proto3" json:"requested_time,omitempty"` // Time requested
	RequestedDay  uint32 `protobuf:"varint,4,opt,name=requested_day,json=requestedDay,proto3" json:"requested_day,omitempty"`    // Day requested
}

func (x *AppGetHistPowerReqDTO) Reset() {
	*x = AppGetHistPowerReqDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_AppGetHistPower_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppGetHistPowerReqDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppGetHistPowerReqDTO) ProtoMessage() {}

func (x *AppGetHistPowerReqDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_AppGetHistPower_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppGetHistPowerReqDTO.ProtoReflect.Descriptor instead.
func (*AppGetHistPowerReqDTO) Descriptor() ([]byte, []int) {
	return file_protos_AppGetHistPower_proto_rawDescGZIP(), []int{0}
}

func (x *AppGetHistPowerReqDTO) GetControlPoint() int32 {
	if x != nil {
		return x.ControlPoint
	}
	return 0
}

func (x *AppGetHistPowerReqDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AppGetHistPowerReqDTO) GetRequestedTime() uint32 {
	if x != nil {
		return x.RequestedTime
	}
	return 0
}

func (x *AppGetHistPowerReqDTO) GetRequestedDay() uint32 {
	if x != nil {
		return x.RequestedDay
	}
	return 0
}

type AppGetHistPowerResDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber  int64   `protobuf:"varint,1,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`      // Device serial number
	AccessPoint   int32   `protobuf:"varint,2,opt,name=access_point,json=accessPoint,proto3" json:"access_point,omitempty"`         // Access point identifier
	ControlPoint  int32   `protobuf:"varint,3,opt,name=control_point,json=controlPoint,proto3" json:"control_point,omitempty"`      // Control point identifier
	Offset        int32   `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`                                      // Offset value
	RequestTime   uint32  `protobuf:"varint,5,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`         // Timestamp of the request
	StartTime     uint32  `protobuf:"varint,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`               // Start timestamp
	LongTermStart uint32  `protobuf:"varint,7,opt,name=long_term_start,json=longTermStart,proto3" json:"long_term_start,omitempty"` // Long-term start timestamp
	AbsoluteStart uint32  `protobuf:"varint,8,opt,name=absolute_start,json=absoluteStart,proto3" json:"absolute_start,omitempty"`   // Absolute start timestamp
	StepTime      uint32  `protobuf:"varint,9,opt,name=step_time,json=stepTime,proto3" json:"step_time,omitempty"`                  // Step time
	RelativePower uint32  `protobuf:"varint,10,opt,name=relative_power,json=relativePower,proto3" json:"relative_power,omitempty"`  // Relative power
	TotalEnergy   uint32  `protobuf:"varint,11,opt,name=total_energy,json=totalEnergy,proto3" json:"total_energy,omitempty"`        // Total energy
	DailyEnergy   uint32  `protobuf:"varint,12,opt,name=daily_energy,json=dailyEnergy,proto3" json:"daily_energy,omitempty"`        // Daily energy
	PowerArray    []int32 `protobuf:"varint,13,rep,packed,name=power_array,json=powerArray,proto3" json:"power_array,omitempty"`    // Array of power values
	WarningNumber uint32  `protobuf:"varint,14,opt,name=warning_number,json=warningNumber,proto3" json:"warning_number,omitempty"`  // Warning number
}

func (x *AppGetHistPowerResDTO) Reset() {
	*x = AppGetHistPowerResDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_AppGetHistPower_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppGetHistPowerResDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppGetHistPowerResDTO) ProtoMessage() {}

func (x *AppGetHistPowerResDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_AppGetHistPower_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppGetHistPowerResDTO.ProtoReflect.Descriptor instead.
func (*AppGetHistPowerResDTO) Descriptor() ([]byte, []int) {
	return file_protos_AppGetHistPower_proto_rawDescGZIP(), []int{1}
}

func (x *AppGetHistPowerResDTO) GetSerialNumber() int64 {
	if x != nil {
		return x.SerialNumber
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetAccessPoint() int32 {
	if x != nil {
		return x.AccessPoint
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetControlPoint() int32 {
	if x != nil {
		return x.ControlPoint
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetRequestTime() uint32 {
	if x != nil {
		return x.RequestTime
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetLongTermStart() uint32 {
	if x != nil {
		return x.LongTermStart
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetAbsoluteStart() uint32 {
	if x != nil {
		return x.AbsoluteStart
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetStepTime() uint32 {
	if x != nil {
		return x.StepTime
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetRelativePower() uint32 {
	if x != nil {
		return x.RelativePower
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetTotalEnergy() uint32 {
	if x != nil {
		return x.TotalEnergy
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetDailyEnergy() uint32 {
	if x != nil {
		return x.DailyEnergy
	}
	return 0
}

func (x *AppGetHistPowerResDTO) GetPowerArray() []int32 {
	if x != nil {
		return x.PowerArray
	}
	return nil
}

func (x *AppGetHistPowerResDTO) GetWarningNumber() uint32 {
	if x != nil {
		return x.WarningNumber
	}
	return 0
}

var File_protos_AppGetHistPower_proto protoreflect.FileDescriptor

var file_protos_AppGetHistPower_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x41, 0x70, 0x70, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0,
	0x01, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x77,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x44, 0x54, 0x4f, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x79, 0x22, 0xff, 0x03, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x6f, 0x6e,
	0x67, 0x54, 0x65, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x62,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x74, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x69, 0x6c,
	0x79, 0x5f, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x64, 0x61, 0x69, 0x6c, 0x79, 0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x0a, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x25, 0x0a, 0x0e,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x42, 0x23, 0x5a, 0x0f, 0x68, 0x6f, 0x79, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x0f, 0x48, 0x6f, 0x79, 0x6d, 0x69, 0x6c, 0x65,
	0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_AppGetHistPower_proto_rawDescOnce sync.Once
	file_protos_AppGetHistPower_proto_rawDescData = file_protos_AppGetHistPower_proto_rawDesc
)

func file_protos_AppGetHistPower_proto_rawDescGZIP() []byte {
	file_protos_AppGetHistPower_proto_rawDescOnce.Do(func() {
		file_protos_AppGetHistPower_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_AppGetHistPower_proto_rawDescData)
	})
	return file_protos_AppGetHistPower_proto_rawDescData
}

var file_protos_AppGetHistPower_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_AppGetHistPower_proto_goTypes = []any{
	(*AppGetHistPowerReqDTO)(nil), // 0: AppGetHistPowerReqDTO
	(*AppGetHistPowerResDTO)(nil), // 1: AppGetHistPowerResDTO
}
var file_protos_AppGetHistPower_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_AppGetHistPower_proto_init() }
func file_protos_AppGetHistPower_proto_init() {
	if File_protos_AppGetHistPower_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_AppGetHistPower_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AppGetHistPowerReqDTO); i {
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
		file_protos_AppGetHistPower_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AppGetHistPowerResDTO); i {
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
			RawDescriptor: file_protos_AppGetHistPower_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_AppGetHistPower_proto_goTypes,
		DependencyIndexes: file_protos_AppGetHistPower_proto_depIdxs,
		MessageInfos:      file_protos_AppGetHistPower_proto_msgTypes,
	}.Build()
	File_protos_AppGetHistPower_proto = out.File
	file_protos_AppGetHistPower_proto_rawDesc = nil
	file_protos_AppGetHistPower_proto_goTypes = nil
	file_protos_AppGetHistPower_proto_depIdxs = nil
}

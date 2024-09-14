// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: protos/EventData.proto

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

type EventDataResDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset     int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`                            // Offset in the response
	Time       int32  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`                                // Timestamp of the response
	TimeYmdHms string `protobuf:"bytes,3,opt,name=time_ymd_hms,json=timeYmdHms,proto3" json:"time_ymd_hms,omitempty"` // Start timestamp in YYYY-MM-DD HH:MM:SS format
}

func (x *EventDataResDTO) Reset() {
	*x = EventDataResDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_EventData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDataResDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDataResDTO) ProtoMessage() {}

func (x *EventDataResDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_EventData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDataResDTO.ProtoReflect.Descriptor instead.
func (*EventDataResDTO) Descriptor() ([]byte, []int) {
	return file_protos_EventData_proto_rawDescGZIP(), []int{0}
}

func (x *EventDataResDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *EventDataResDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *EventDataResDTO) GetTimeYmdHms() string {
	if x != nil {
		return x.TimeYmdHms
	}
	return ""
}

type MIEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventCode      int32 `protobuf:"varint,1,opt,name=event_code,json=eventCode,proto3" json:"event_code,omitempty"`                 // Event code
	EventStatus    int32 `protobuf:"varint,2,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`           // Event status
	EventCount     int32 `protobuf:"varint,3,opt,name=event_count,json=eventCount,proto3" json:"event_count,omitempty"`              // Event count
	PvVoltage      int32 `protobuf:"varint,4,opt,name=pv_voltage,json=pvVoltage,proto3" json:"pv_voltage,omitempty"`                 // PV voltage
	GridVoltage    int32 `protobuf:"varint,5,opt,name=grid_voltage,json=gridVoltage,proto3" json:"grid_voltage,omitempty"`           // Grid voltage
	GridFrequency  int32 `protobuf:"varint,6,opt,name=grid_frequency,json=gridFrequency,proto3" json:"grid_frequency,omitempty"`     // Grid frequency
	GridPower      int32 `protobuf:"varint,7,opt,name=grid_power,json=gridPower,proto3" json:"grid_power,omitempty"`                 // Grid power
	Temperature    int32 `protobuf:"varint,8,opt,name=temperature,proto3" json:"temperature,omitempty"`                              // Temperature
	MiId           int64 `protobuf:"varint,9,opt,name=mi_id,json=miId,proto3" json:"mi_id,omitempty"`                                // Meter Interface (MI) ID
	StartTimestamp int32 `protobuf:"varint,10,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"` // Start timestamp
}

func (x *MIEvent) Reset() {
	*x = MIEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_EventData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MIEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MIEvent) ProtoMessage() {}

func (x *MIEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_EventData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MIEvent.ProtoReflect.Descriptor instead.
func (*MIEvent) Descriptor() ([]byte, []int) {
	return file_protos_EventData_proto_rawDescGZIP(), []int{1}
}

func (x *MIEvent) GetEventCode() int32 {
	if x != nil {
		return x.EventCode
	}
	return 0
}

func (x *MIEvent) GetEventStatus() int32 {
	if x != nil {
		return x.EventStatus
	}
	return 0
}

func (x *MIEvent) GetEventCount() int32 {
	if x != nil {
		return x.EventCount
	}
	return 0
}

func (x *MIEvent) GetPvVoltage() int32 {
	if x != nil {
		return x.PvVoltage
	}
	return 0
}

func (x *MIEvent) GetGridVoltage() int32 {
	if x != nil {
		return x.GridVoltage
	}
	return 0
}

func (x *MIEvent) GetGridFrequency() int32 {
	if x != nil {
		return x.GridFrequency
	}
	return 0
}

func (x *MIEvent) GetGridPower() int32 {
	if x != nil {
		return x.GridPower
	}
	return 0
}

func (x *MIEvent) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *MIEvent) GetMiId() int64 {
	if x != nil {
		return x.MiId
	}
	return 0
}

func (x *MIEvent) GetStartTimestamp() int32 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

type EventDataReqDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int32      `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`                    // Offset in the request
	Time     int32      `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`                        // Timestamp of the request
	MiEvents []*MIEvent `protobuf:"bytes,3,rep,name=mi_events,json=miEvents,proto3" json:"mi_events,omitempty"` // List of MI (Meter Interface) events
}

func (x *EventDataReqDTO) Reset() {
	*x = EventDataReqDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_EventData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventDataReqDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventDataReqDTO) ProtoMessage() {}

func (x *EventDataReqDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_EventData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventDataReqDTO.ProtoReflect.Descriptor instead.
func (*EventDataReqDTO) Descriptor() ([]byte, []int) {
	return file_protos_EventData_proto_rawDescGZIP(), []int{2}
}

func (x *EventDataReqDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *EventDataReqDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *EventDataReqDTO) GetMiEvents() []*MIEvent {
	if x != nil {
		return x.MiEvents
	}
	return nil
}

var File_protos_EventData_proto protoreflect.FileDescriptor

var file_protos_EventData_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x79, 0x6d, 0x64, 0x5f, 0x68, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x69, 0x6d, 0x65, 0x59, 0x6d, 0x64, 0x48, 0x6d, 0x73, 0x22, 0xd4, 0x02, 0x0a, 0x07, 0x4d, 0x49,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x76, 0x5f, 0x76,
	0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x76,
	0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x72, 0x69, 0x64, 0x5f,
	0x76, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67,
	0x72, 0x69, 0x64, 0x56, 0x6f, 0x6c, 0x74, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x67, 0x72,
	0x69, 0x64, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x67, 0x72, 0x69, 0x64, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x69, 0x64, 0x5f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x67, 0x72, 0x69, 0x64, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x6d, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x6d, 0x69, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x64, 0x0a, 0x0f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x44, 0x54, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x25, 0x0a, 0x09, 0x6d, 0x69, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x49, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x6d, 0x69,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x23, 0x5a, 0x0f, 0x68, 0x6f, 0x79, 0x6d, 0x69, 0x6c,
	0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0xaa, 0x02, 0x0f, 0x48, 0x6f, 0x79, 0x6d,
	0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_EventData_proto_rawDescOnce sync.Once
	file_protos_EventData_proto_rawDescData = file_protos_EventData_proto_rawDesc
)

func file_protos_EventData_proto_rawDescGZIP() []byte {
	file_protos_EventData_proto_rawDescOnce.Do(func() {
		file_protos_EventData_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_EventData_proto_rawDescData)
	})
	return file_protos_EventData_proto_rawDescData
}

var file_protos_EventData_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_EventData_proto_goTypes = []any{
	(*EventDataResDTO)(nil), // 0: EventDataResDTO
	(*MIEvent)(nil),         // 1: MIEvent
	(*EventDataReqDTO)(nil), // 2: EventDataReqDTO
}
var file_protos_EventData_proto_depIdxs = []int32{
	1, // 0: EventDataReqDTO.mi_events:type_name -> MIEvent
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_EventData_proto_init() }
func file_protos_EventData_proto_init() {
	if File_protos_EventData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_EventData_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EventDataResDTO); i {
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
		file_protos_EventData_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MIEvent); i {
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
		file_protos_EventData_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EventDataReqDTO); i {
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
			RawDescriptor: file_protos_EventData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_EventData_proto_goTypes,
		DependencyIndexes: file_protos_EventData_proto_depIdxs,
		MessageInfos:      file_protos_EventData_proto_msgTypes,
	}.Build()
	File_protos_EventData_proto = out.File
	file_protos_EventData_proto_rawDesc = nil
	file_protos_EventData_proto_goTypes = nil
	file_protos_EventData_proto_depIdxs = nil
}
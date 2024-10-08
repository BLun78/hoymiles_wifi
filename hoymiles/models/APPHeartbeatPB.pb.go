// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: protos/APPHeartbeatPB.proto

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

type HBResDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset             int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`                                                    // Offset value
	Time               int32  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`                                                        // Timestamp of the request
	Csq                int32  `protobuf:"varint,3,opt,name=csq,proto3" json:"csq,omitempty"`                                                          // Carrier Signal Quality (CSQ)
	DtuSerialNumber    string `protobuf:"bytes,4,opt,name=dtu_serial_number,json=dtuSerialNumber,proto3" json:"dtu_serial_number,omitempty"`          // DTU serial number
	DeviceSerialNumber string `protobuf:"bytes,5,opt,name=device_serial_number,json=deviceSerialNumber,proto3" json:"device_serial_number,omitempty"` // Device serial number
}

func (x *HBResDTO) Reset() {
	*x = HBResDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_APPHeartbeatPB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HBResDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HBResDTO) ProtoMessage() {}

func (x *HBResDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_APPHeartbeatPB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HBResDTO.ProtoReflect.Descriptor instead.
func (*HBResDTO) Descriptor() ([]byte, []int) {
	return file_protos_APPHeartbeatPB_proto_rawDescGZIP(), []int{0}
}

func (x *HBResDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *HBResDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *HBResDTO) GetCsq() int32 {
	if x != nil {
		return x.Csq
	}
	return 0
}

func (x *HBResDTO) GetDtuSerialNumber() string {
	if x != nil {
		return x.DtuSerialNumber
	}
	return ""
}

func (x *HBResDTO) GetDeviceSerialNumber() string {
	if x != nil {
		return x.DeviceSerialNumber
	}
	return ""
}

type HBReqDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset     int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`                            // Offset value
	Time       int32  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`                                // Timestamp of the response
	TimeYmdHms string `protobuf:"bytes,3,opt,name=time_ymd_hms,json=timeYmdHms,proto3" json:"time_ymd_hms,omitempty"` // Timestamp in the format YMD_HMS
}

func (x *HBReqDTO) Reset() {
	*x = HBReqDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_APPHeartbeatPB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HBReqDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HBReqDTO) ProtoMessage() {}

func (x *HBReqDTO) ProtoReflect() protoreflect.Message {
	mi := &file_protos_APPHeartbeatPB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HBReqDTO.ProtoReflect.Descriptor instead.
func (*HBReqDTO) Descriptor() ([]byte, []int) {
	return file_protos_APPHeartbeatPB_proto_rawDescGZIP(), []int{1}
}

func (x *HBReqDTO) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *HBReqDTO) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *HBReqDTO) GetTimeYmdHms() string {
	if x != nil {
		return x.TimeYmdHms
	}
	return ""
}

var File_protos_APPHeartbeatPB_proto protoreflect.FileDescriptor

var file_protos_APPHeartbeatPB_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x41, 0x50, 0x50, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x50, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01,
	0x0a, 0x08, 0x48, 0x42, 0x52, 0x65, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x71, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x63, 0x73, 0x71, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x74, 0x75, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x74, 0x75, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x08, 0x48, 0x42, 0x52, 0x65, 0x71, 0x44,
	0x54, 0x4f, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x79, 0x6d, 0x64, 0x5f, 0x68, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x59, 0x6d, 0x64, 0x48, 0x6d, 0x73,
	0x42, 0x23, 0x5a, 0x0f, 0x68, 0x6f, 0x79, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0xaa, 0x02, 0x0f, 0x48, 0x6f, 0x79, 0x6d, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_APPHeartbeatPB_proto_rawDescOnce sync.Once
	file_protos_APPHeartbeatPB_proto_rawDescData = file_protos_APPHeartbeatPB_proto_rawDesc
)

func file_protos_APPHeartbeatPB_proto_rawDescGZIP() []byte {
	file_protos_APPHeartbeatPB_proto_rawDescOnce.Do(func() {
		file_protos_APPHeartbeatPB_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_APPHeartbeatPB_proto_rawDescData)
	})
	return file_protos_APPHeartbeatPB_proto_rawDescData
}

var file_protos_APPHeartbeatPB_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_APPHeartbeatPB_proto_goTypes = []any{
	(*HBResDTO)(nil), // 0: HBResDTO
	(*HBReqDTO)(nil), // 1: HBReqDTO
}
var file_protos_APPHeartbeatPB_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_APPHeartbeatPB_proto_init() }
func file_protos_APPHeartbeatPB_proto_init() {
	if File_protos_APPHeartbeatPB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_APPHeartbeatPB_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HBResDTO); i {
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
		file_protos_APPHeartbeatPB_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*HBReqDTO); i {
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
			RawDescriptor: file_protos_APPHeartbeatPB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_APPHeartbeatPB_proto_goTypes,
		DependencyIndexes: file_protos_APPHeartbeatPB_proto_depIdxs,
		MessageInfos:      file_protos_APPHeartbeatPB_proto_msgTypes,
	}.Build()
	File_protos_APPHeartbeatPB_proto = out.File
	file_protos_APPHeartbeatPB_proto_rawDesc = nil
	file_protos_APPHeartbeatPB_proto_goTypes = nil
	file_protos_APPHeartbeatPB_proto_depIdxs = nil
}

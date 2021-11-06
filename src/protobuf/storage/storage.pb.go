// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: storage.proto

package storage

import (
	common "github.com/zikwall/monit/src/protobuf/common"
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

type HeatmapMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamID  string `protobuf:"bytes,1,opt,name=streamID,proto3" json:"streamID,omitempty"`
	UniqueID  string `protobuf:"bytes,2,opt,name=uniqueID,proto3" json:"uniqueID,omitempty"`
	Platform  string `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	App       string `protobuf:"bytes,4,opt,name=app,proto3" json:"app,omitempty"`
	Version   string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Os        string `protobuf:"bytes,6,opt,name=os,proto3" json:"os,omitempty"`
	Browser   string `protobuf:"bytes,7,opt,name=browser,proto3" json:"browser,omitempty"`
	Region    string `protobuf:"bytes,9,opt,name=region,proto3" json:"region,omitempty"`
	Host      string `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`
	UserAgent string `protobuf:"bytes,11,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
}

func (x *HeatmapMessage) Reset() {
	*x = HeatmapMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeatmapMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeatmapMessage) ProtoMessage() {}

func (x *HeatmapMessage) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeatmapMessage.ProtoReflect.Descriptor instead.
func (*HeatmapMessage) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{0}
}

func (x *HeatmapMessage) GetStreamID() string {
	if x != nil {
		return x.StreamID
	}
	return ""
}

func (x *HeatmapMessage) GetUniqueID() string {
	if x != nil {
		return x.UniqueID
	}
	return ""
}

func (x *HeatmapMessage) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *HeatmapMessage) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *HeatmapMessage) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HeatmapMessage) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *HeatmapMessage) GetBrowser() string {
	if x != nil {
		return x.Browser
	}
	return ""
}

func (x *HeatmapMessage) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *HeatmapMessage) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HeatmapMessage) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type HeatmapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *common.Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *HeatmapResponse) Reset() {
	*x = HeatmapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeatmapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeatmapResponse) ProtoMessage() {}

func (x *HeatmapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeatmapResponse.ProtoReflect.Descriptor instead.
func (*HeatmapResponse) Descriptor() ([]byte, []int) {
	return file_storage_proto_rawDescGZIP(), []int{1}
}

func (x *HeatmapResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_storage_proto protoreflect.FileDescriptor

var file_storage_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x20, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x02, 0x0a, 0x0e, 0x48,
	0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x61, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x22, 0x36, 0x0a, 0x0f, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x4e, 0x0a, 0x07, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x48, 0x65, 0x61,
	0x74, 0x6d, 0x61, 0x70, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x48,
	0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x74, 0x6d, 0x61, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x6b, 0x77, 0x61, 0x6c, 0x6c, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_storage_proto_rawDescOnce sync.Once
	file_storage_proto_rawDescData = file_storage_proto_rawDesc
)

func file_storage_proto_rawDescGZIP() []byte {
	file_storage_proto_rawDescOnce.Do(func() {
		file_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_proto_rawDescData)
	})
	return file_storage_proto_rawDescData
}

var file_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_storage_proto_goTypes = []interface{}{
	(*HeatmapMessage)(nil),  // 0: storage.HeatmapMessage
	(*HeatmapResponse)(nil), // 1: storage.HeatmapResponse
	(*common.Error)(nil),    // 2: common.Error
}
var file_storage_proto_depIdxs = []int32{
	2, // 0: storage.HeatmapResponse.error:type_name -> common.Error
	0, // 1: storage.Storage.WriteHeatmap:input_type -> storage.HeatmapMessage
	1, // 2: storage.Storage.WriteHeatmap:output_type -> storage.HeatmapResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_proto_init() }
func file_storage_proto_init() {
	if File_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeatmapMessage); i {
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
		file_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeatmapResponse); i {
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
			RawDescriptor: file_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_storage_proto_goTypes,
		DependencyIndexes: file_storage_proto_depIdxs,
		MessageInfos:      file_storage_proto_msgTypes,
	}.Build()
	File_storage_proto = out.File
	file_storage_proto_rawDesc = nil
	file_storage_proto_goTypes = nil
	file_storage_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/chart.proto

package charts

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

type LineChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *LineChartRequest) Reset() {
	*x = LineChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineChartRequest) ProtoMessage() {}

func (x *LineChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineChartRequest.ProtoReflect.Descriptor instead.
func (*LineChartRequest) Descriptor() ([]byte, []int) {
	return file_proto_chart_proto_rawDescGZIP(), []int{0}
}

func (x *LineChartRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type LineChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *LineChartResponse) Reset() {
	*x = LineChartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LineChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LineChartResponse) ProtoMessage() {}

func (x *LineChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LineChartResponse.ProtoReflect.Descriptor instead.
func (*LineChartResponse) Descriptor() ([]byte, []int) {
	return file_proto_chart_proto_rawDescGZIP(), []int{1}
}

func (x *LineChartResponse) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *LineChartResponse) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

type ProgressChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProgressChartRequest) Reset() {
	*x = ProgressChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressChartRequest) ProtoMessage() {}

func (x *ProgressChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressChartRequest.ProtoReflect.Descriptor instead.
func (*ProgressChartRequest) Descriptor() ([]byte, []int) {
	return file_proto_chart_proto_rawDescGZIP(), []int{2}
}

type ProgressChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ProgressChartResponse) Reset() {
	*x = ProgressChartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgressChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressChartResponse) ProtoMessage() {}

func (x *ProgressChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgressChartResponse.ProtoReflect.Descriptor instead.
func (*ProgressChartResponse) Descriptor() ([]byte, []int) {
	return file_proto_chart_proto_rawDescGZIP(), []int{3}
}

func (x *ProgressChartResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ProgressChartResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_proto_chart_proto protoreflect.FileDescriptor

var file_proto_chart_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x10, 0x4c,
	0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2f, 0x0a, 0x11, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x22, 0x16, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43,
	0x0a, 0x15, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0xb8, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x2e, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x2e,
	0x4c, 0x69, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_chart_proto_rawDescOnce sync.Once
	file_proto_chart_proto_rawDescData = file_proto_chart_proto_rawDesc
)

func file_proto_chart_proto_rawDescGZIP() []byte {
	file_proto_chart_proto_rawDescOnce.Do(func() {
		file_proto_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chart_proto_rawDescData)
	})
	return file_proto_chart_proto_rawDescData
}

var file_proto_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_chart_proto_goTypes = []interface{}{
	(*LineChartRequest)(nil),      // 0: charts.LineChartRequest
	(*LineChartResponse)(nil),     // 1: charts.LineChartResponse
	(*ProgressChartRequest)(nil),  // 2: charts.ProgressChartRequest
	(*ProgressChartResponse)(nil), // 3: charts.ProgressChartResponse
}
var file_proto_chart_proto_depIdxs = []int32{
	0, // 0: charts.ChartDataService.GetLineChartData:input_type -> charts.LineChartRequest
	2, // 1: charts.ChartDataService.GetProgressChartData:input_type -> charts.ProgressChartRequest
	1, // 2: charts.ChartDataService.GetLineChartData:output_type -> charts.LineChartResponse
	3, // 3: charts.ChartDataService.GetProgressChartData:output_type -> charts.ProgressChartResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_chart_proto_init() }
func file_proto_chart_proto_init() {
	if File_proto_chart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineChartRequest); i {
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
		file_proto_chart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LineChartResponse); i {
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
		file_proto_chart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressChartRequest); i {
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
		file_proto_chart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgressChartResponse); i {
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
			RawDescriptor: file_proto_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chart_proto_goTypes,
		DependencyIndexes: file_proto_chart_proto_depIdxs,
		MessageInfos:      file_proto_chart_proto_msgTypes,
	}.Build()
	File_proto_chart_proto = out.File
	file_proto_chart_proto_rawDesc = nil
	file_proto_chart_proto_goTypes = nil
	file_proto_chart_proto_depIdxs = nil
}

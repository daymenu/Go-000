// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: area.proto

package model

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 定义一个点
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

// 定义圆
type Circular struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dot    *Point  `protobuf:"bytes,1,opt,name=dot,proto3" json:"dot,omitempty"`
	Radius float64 `protobuf:"fixed64,2,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *Circular) Reset() {
	*x = Circular{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Circular) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circular) ProtoMessage() {}

func (x *Circular) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circular.ProtoReflect.Descriptor instead.
func (*Circular) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{1}
}

func (x *Circular) GetDot() *Point {
	if x != nil {
		return x.Dot
	}
	return nil
}

func (x *Circular) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

// 定义请求消息
type AreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string    `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Circular  *Circular `protobuf:"bytes,2,opt,name=circular,proto3" json:"circular,omitempty"`
}

func (x *AreaRequest) Reset() {
	*x = AreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaRequest) ProtoMessage() {}

func (x *AreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaRequest.ProtoReflect.Descriptor instead.
func (*AreaRequest) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{2}
}

func (x *AreaRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AreaRequest) GetCircular() *Circular {
	if x != nil {
		return x.Circular
	}
	return nil
}

// 定义响应消息
type AreaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Area float64 `protobuf:"fixed64,2,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *AreaResponse) Reset() {
	*x = AreaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_area_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AreaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AreaResponse) ProtoMessage() {}

func (x *AreaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_area_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AreaResponse.ProtoReflect.Descriptor instead.
func (*AreaResponse) Descriptor() ([]byte, []int) {
	return file_area_proto_rawDescGZIP(), []int{3}
}

func (x *AreaResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AreaResponse) GetArea() float64 {
	if x != nil {
		return x.Area
	}
	return 0
}

var File_area_proto protoreflect.FileDescriptor

var file_area_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x42, 0x0a, 0x08, 0x43, 0x69, 0x72, 0x63,
	0x75, 0x6c, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x03, 0x64, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x03, 0x64, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x0b,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x08, 0x63, 0x69,
	0x72, 0x63, 0x75, 0x6c, 0x61, 0x72, 0x22, 0x36, 0x0a, 0x0c, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x32, 0x46,
	0x0a, 0x0f, 0x43, 0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x47, 0x6f, 0x2d,
	0x30, 0x30, 0x30, 0x2f, 0x57, 0x65, 0x65, 0x6b, 0x30, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_area_proto_rawDescOnce sync.Once
	file_area_proto_rawDescData = file_area_proto_rawDesc
)

func file_area_proto_rawDescGZIP() []byte {
	file_area_proto_rawDescOnce.Do(func() {
		file_area_proto_rawDescData = protoimpl.X.CompressGZIP(file_area_proto_rawDescData)
	})
	return file_area_proto_rawDescData
}

var file_area_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_area_proto_goTypes = []interface{}{
	(*Point)(nil),        // 0: model.Point
	(*Circular)(nil),     // 1: model.Circular
	(*AreaRequest)(nil),  // 2: model.AreaRequest
	(*AreaResponse)(nil), // 3: model.AreaResponse
}
var file_area_proto_depIdxs = []int32{
	0, // 0: model.Circular.dot:type_name -> model.Point
	1, // 1: model.AreaRequest.circular:type_name -> model.Circular
	2, // 2: model.CircularService.Area:input_type -> model.AreaRequest
	3, // 3: model.CircularService.Area:output_type -> model.AreaResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_area_proto_init() }
func file_area_proto_init() {
	if File_area_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_area_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_area_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Circular); i {
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
		file_area_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaRequest); i {
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
		file_area_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AreaResponse); i {
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
			RawDescriptor: file_area_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_area_proto_goTypes,
		DependencyIndexes: file_area_proto_depIdxs,
		MessageInfos:      file_area_proto_msgTypes,
	}.Build()
	File_area_proto = out.File
	file_area_proto_rawDesc = nil
	file_area_proto_goTypes = nil
	file_area_proto_depIdxs = nil
}

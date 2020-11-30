// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: movie.proto

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

type Moive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ToStar string `protobuf:"bytes,3,opt,name=to_star,json=toStar,proto3" json:"to_star,omitempty"`
	Love   int64  `protobuf:"varint,4,opt,name=love,proto3" json:"love,omitempty"`
}

func (x *Moive) Reset() {
	*x = Moive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Moive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moive) ProtoMessage() {}

func (x *Moive) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moive.ProtoReflect.Descriptor instead.
func (*Moive) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Moive) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Moive) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Moive) GetToStar() string {
	if x != nil {
		return x.ToStar
	}
	return ""
}

func (x *Moive) GetLove() int64 {
	if x != nil {
		return x.Love
	}
	return 0
}

type SearchMoive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Moive     *Moive   `protobuf:"bytes,1,opt,name=moive,proto3" json:"moive,omitempty"`
	LoveMoive []*Moive `protobuf:"bytes,2,rep,name=loveMoive,proto3" json:"loveMoive,omitempty"`
}

func (x *SearchMoive) Reset() {
	*x = SearchMoive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoive) ProtoMessage() {}

func (x *SearchMoive) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoive.ProtoReflect.Descriptor instead.
func (*SearchMoive) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{1}
}

func (x *SearchMoive) GetMoive() *Moive {
	if x != nil {
		return x.Moive
	}
	return nil
}

func (x *SearchMoive) GetLoveMoive() []*Moive {
	if x != nil {
		return x.LoveMoive
	}
	return nil
}

type SearchMoiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *SearchMoiveRequest) Reset() {
	*x = SearchMoiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoiveRequest) ProtoMessage() {}

func (x *SearchMoiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoiveRequest.ProtoReflect.Descriptor instead.
func (*SearchMoiveRequest) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{2}
}

func (x *SearchMoiveRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type SearchMoiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SearchMoive *SearchMoive `protobuf:"bytes,2,opt,name=SearchMoive,proto3" json:"SearchMoive,omitempty"`
}

func (x *SearchMoiveResponse) Reset() {
	*x = SearchMoiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movie_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMoiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMoiveResponse) ProtoMessage() {}

func (x *SearchMoiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movie_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMoiveResponse.ProtoReflect.Descriptor instead.
func (*SearchMoiveResponse) Descriptor() ([]byte, []int) {
	return file_movie_proto_rawDescGZIP(), []int{3}
}

func (x *SearchMoiveResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SearchMoiveResponse) GetSearchMoive() *SearchMoive {
	if x != nil {
		return x.SearchMoive
	}
	return nil
}

var File_movie_proto protoreflect.FileDescriptor

var file_movie_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f,
	0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x6f, 0x76, 0x65, 0x22, 0x5d,
	0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x12, 0x22, 0x0a,
	0x05, 0x6d, 0x6f, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x69, 0x76,
	0x65, 0x12, 0x2a, 0x0a, 0x09, 0x6c, 0x6f, 0x76, 0x65, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x69,
	0x76, 0x65, 0x52, 0x09, 0x6c, 0x6f, 0x76, 0x65, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x22, 0x32, 0x0a,
	0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x22, 0x5f, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x34, 0x0a, 0x0b,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x4d, 0x6f, 0x69, 0x76, 0x65, 0x52, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x69,
	0x76, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x75, 0x2f, 0x47, 0x6f, 0x2d, 0x30, 0x30, 0x30, 0x2f,
	0x57, 0x65, 0x65, 0x6b, 0x30, 0x32, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movie_proto_rawDescOnce sync.Once
	file_movie_proto_rawDescData = file_movie_proto_rawDesc
)

func file_movie_proto_rawDescGZIP() []byte {
	file_movie_proto_rawDescOnce.Do(func() {
		file_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_movie_proto_rawDescData)
	})
	return file_movie_proto_rawDescData
}

var file_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_movie_proto_goTypes = []interface{}{
	(*Moive)(nil),               // 0: model.Moive
	(*SearchMoive)(nil),         // 1: model.SearchMoive
	(*SearchMoiveRequest)(nil),  // 2: model.SearchMoiveRequest
	(*SearchMoiveResponse)(nil), // 3: model.SearchMoiveResponse
}
var file_movie_proto_depIdxs = []int32{
	0, // 0: model.SearchMoive.moive:type_name -> model.Moive
	0, // 1: model.SearchMoive.loveMoive:type_name -> model.Moive
	1, // 2: model.SearchMoiveResponse.SearchMoive:type_name -> model.SearchMoive
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_movie_proto_init() }
func file_movie_proto_init() {
	if File_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Moive); i {
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
		file_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoive); i {
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
		file_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoiveRequest); i {
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
		file_movie_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMoiveResponse); i {
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
			RawDescriptor: file_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_movie_proto_goTypes,
		DependencyIndexes: file_movie_proto_depIdxs,
		MessageInfos:      file_movie_proto_msgTypes,
	}.Build()
	File_movie_proto = out.File
	file_movie_proto_rawDesc = nil
	file_movie_proto_goTypes = nil
	file_movie_proto_depIdxs = nil
}
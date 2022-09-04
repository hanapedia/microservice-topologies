// This proto contains gRPC definition for fanout_x service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: pb_fanout_3/fanout.proto

package pb_fanout_3

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

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fanout_3_fanout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fanout_3_fanout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_pb_fanout_3_fanout_proto_rawDescGZIP(), []int{0}
}

func (x *Res) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_fanout_3_fanout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_pb_fanout_3_fanout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_pb_fanout_3_fanout_proto_rawDescGZIP(), []int{1}
}

func (x *Req) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_pb_fanout_3_fanout_proto protoreflect.FileDescriptor

var file_pb_fanout_3_fanout_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x62, 0x5f, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x33, 0x2f, 0x66, 0x61,
	0x6e, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x66, 0x61, 0x6e, 0x6f,
	0x75, 0x74, 0x5f, 0x33, 0x22, 0x17, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x17, 0x0a,
	0x03, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32, 0x32, 0x0a, 0x08, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74,
	0x5f, 0x33, 0x12, 0x26, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x49, 0x64, 0x73, 0x12, 0x0d, 0x2e, 0x66,
	0x61, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x33, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x66, 0x61,
	0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x6e, 0x61, 0x70, 0x65, 0x64,
	0x69, 0x61, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x69, 0x65, 0x73, 0x2f, 0x66, 0x61, 0x6e, 0x6f, 0x75,
	0x74, 0x2f, 0x66, 0x61, 0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x33, 0x2f, 0x70, 0x62, 0x5f, 0x66, 0x61,
	0x6e, 0x6f, 0x75, 0x74, 0x5f, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_fanout_3_fanout_proto_rawDescOnce sync.Once
	file_pb_fanout_3_fanout_proto_rawDescData = file_pb_fanout_3_fanout_proto_rawDesc
)

func file_pb_fanout_3_fanout_proto_rawDescGZIP() []byte {
	file_pb_fanout_3_fanout_proto_rawDescOnce.Do(func() {
		file_pb_fanout_3_fanout_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_fanout_3_fanout_proto_rawDescData)
	})
	return file_pb_fanout_3_fanout_proto_rawDescData
}

var file_pb_fanout_3_fanout_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_fanout_3_fanout_proto_goTypes = []interface{}{
	(*Res)(nil), // 0: fanout_3.Res
	(*Req)(nil), // 1: fanout_3.Req
}
var file_pb_fanout_3_fanout_proto_depIdxs = []int32{
	1, // 0: fanout_3.fanout_3.GetIds:input_type -> fanout_3.Req
	0, // 1: fanout_3.fanout_3.GetIds:output_type -> fanout_3.Res
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_fanout_3_fanout_proto_init() }
func file_pb_fanout_3_fanout_proto_init() {
	if File_pb_fanout_3_fanout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_fanout_3_fanout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
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
		file_pb_fanout_3_fanout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
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
			RawDescriptor: file_pb_fanout_3_fanout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_fanout_3_fanout_proto_goTypes,
		DependencyIndexes: file_pb_fanout_3_fanout_proto_depIdxs,
		MessageInfos:      file_pb_fanout_3_fanout_proto_msgTypes,
	}.Build()
	File_pb_fanout_3_fanout_proto = out.File
	file_pb_fanout_3_fanout_proto_rawDesc = nil
	file_pb_fanout_3_fanout_proto_goTypes = nil
	file_pb_fanout_3_fanout_proto_depIdxs = nil
}

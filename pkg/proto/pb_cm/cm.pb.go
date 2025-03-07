// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.28.1
// source: pb_cm/cm.proto

package pb_cm

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	pb_enum "lark/pkg/proto/pb_enum"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CloudMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic    pb_enum.TOPIC         `protobuf:"varint,1,opt,name=topic,proto3,enum=pb_enum.TOPIC" json:"topic,omitempty"`
	SubTopic pb_enum.SUB_TOPIC     `protobuf:"varint,2,opt,name=sub_topic,json=subTopic,proto3,enum=pb_enum.SUB_TOPIC" json:"sub_topic,omitempty"`
	Member   []*CloudMessageMember `protobuf:"bytes,3,rep,name=member,proto3" json:"member,omitempty"`
	Body     []byte                `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CloudMessageReq) Reset() {
	*x = CloudMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cm_cm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudMessageReq) ProtoMessage() {}

func (x *CloudMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cm_cm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudMessageReq.ProtoReflect.Descriptor instead.
func (*CloudMessageReq) Descriptor() ([]byte, []int) {
	return file_pb_cm_cm_proto_rawDescGZIP(), []int{0}
}

func (x *CloudMessageReq) GetTopic() pb_enum.TOPIC {
	if x != nil {
		return x.Topic
	}
	return pb_enum.TOPIC(0)
}

func (x *CloudMessageReq) GetSubTopic() pb_enum.SUB_TOPIC {
	if x != nil {
		return x.SubTopic
	}
	return pb_enum.SUB_TOPIC(0)
}

func (x *CloudMessageReq) GetMember() []*CloudMessageMember {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *CloudMessageReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type CloudMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CloudMessageResp) Reset() {
	*x = CloudMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cm_cm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudMessageResp) ProtoMessage() {}

func (x *CloudMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cm_cm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudMessageResp.ProtoReflect.Descriptor instead.
func (*CloudMessageResp) Descriptor() ([]byte, []int) {
	return file_pb_cm_cm_proto_rawDescGZIP(), []int{1}
}

func (x *CloudMessageResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CloudMessageResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CloudMessageMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64                 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Platform pb_enum.PLATFORM_TYPE `protobuf:"varint,3,opt,name=platform,proto3,enum=pb_enum.PLATFORM_TYPE" json:"platform,omitempty"`
}

func (x *CloudMessageMember) Reset() {
	*x = CloudMessageMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_cm_cm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudMessageMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudMessageMember) ProtoMessage() {}

func (x *CloudMessageMember) ProtoReflect() protoreflect.Message {
	mi := &file_pb_cm_cm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudMessageMember.ProtoReflect.Descriptor instead.
func (*CloudMessageMember) Descriptor() ([]byte, []int) {
	return file_pb_cm_cm_proto_rawDescGZIP(), []int{2}
}

func (x *CloudMessageMember) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CloudMessageMember) GetPlatform() pb_enum.PLATFORM_TYPE {
	if x != nil {
		return x.Platform
	}
	return pb_enum.PLATFORM_TYPE(0)
}

var File_pb_cm_cm_proto protoreflect.FileDescriptor

var file_pb_cm_cm_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x5f, 0x63, 0x6d, 0x2f, 0x63, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x62, 0x5f, 0x63, 0x6d, 0x1a, 0x12, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x0f,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x24, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2f, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x53, 0x55, 0x42, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6d, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x38, 0x0a,
	0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x5a, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x32, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x50, 0x4c, 0x41, 0x54,
	0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x32, 0x4f, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x5f, 0x63, 0x6d, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x5f, 0x63, 0x6d, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x63, 0x6d, 0x3b,
	0x70, 0x62, 0x5f, 0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_cm_cm_proto_rawDescOnce sync.Once
	file_pb_cm_cm_proto_rawDescData = file_pb_cm_cm_proto_rawDesc
)

func file_pb_cm_cm_proto_rawDescGZIP() []byte {
	file_pb_cm_cm_proto_rawDescOnce.Do(func() {
		file_pb_cm_cm_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_cm_cm_proto_rawDescData)
	})
	return file_pb_cm_cm_proto_rawDescData
}

var file_pb_cm_cm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_cm_cm_proto_goTypes = []interface{}{
	(*CloudMessageReq)(nil),    // 0: pb_cm.CloudMessageReq
	(*CloudMessageResp)(nil),   // 1: pb_cm.CloudMessageResp
	(*CloudMessageMember)(nil), // 2: pb_cm.CloudMessageMember
	(pb_enum.TOPIC)(0),         // 3: pb_enum.TOPIC
	(pb_enum.SUB_TOPIC)(0),     // 4: pb_enum.SUB_TOPIC
	(pb_enum.PLATFORM_TYPE)(0), // 5: pb_enum.PLATFORM_TYPE
}
var file_pb_cm_cm_proto_depIdxs = []int32{
	3, // 0: pb_cm.CloudMessageReq.topic:type_name -> pb_enum.TOPIC
	4, // 1: pb_cm.CloudMessageReq.sub_topic:type_name -> pb_enum.SUB_TOPIC
	2, // 2: pb_cm.CloudMessageReq.member:type_name -> pb_cm.CloudMessageMember
	5, // 3: pb_cm.CloudMessageMember.platform:type_name -> pb_enum.PLATFORM_TYPE
	0, // 4: pb_cm.CloudMessage.CloudMessage:input_type -> pb_cm.CloudMessageReq
	1, // 5: pb_cm.CloudMessage.CloudMessage:output_type -> pb_cm.CloudMessageResp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_cm_cm_proto_init() }
func file_pb_cm_cm_proto_init() {
	if File_pb_cm_cm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_cm_cm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudMessageReq); i {
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
		file_pb_cm_cm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudMessageResp); i {
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
		file_pb_cm_cm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudMessageMember); i {
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
			RawDescriptor: file_pb_cm_cm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_cm_cm_proto_goTypes,
		DependencyIndexes: file_pb_cm_cm_proto_depIdxs,
		MessageInfos:      file_pb_cm_cm_proto_msgTypes,
	}.Build()
	File_pb_cm_cm_proto = out.File
	file_pb_cm_cm_proto_rawDesc = nil
	file_pb_cm_cm_proto_goTypes = nil
	file_pb_cm_cm_proto_depIdxs = nil
}

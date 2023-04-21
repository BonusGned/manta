// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: networksystem_protomessages.proto

package dota

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

type NetMessageSplitscreenUserChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot *uint32 `protobuf:"varint,1,opt,name=slot" json:"slot,omitempty"`
}

func (x *NetMessageSplitscreenUserChanged) Reset() {
	*x = NetMessageSplitscreenUserChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networksystem_protomessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessageSplitscreenUserChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessageSplitscreenUserChanged) ProtoMessage() {}

func (x *NetMessageSplitscreenUserChanged) ProtoReflect() protoreflect.Message {
	mi := &file_networksystem_protomessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessageSplitscreenUserChanged.ProtoReflect.Descriptor instead.
func (*NetMessageSplitscreenUserChanged) Descriptor() ([]byte, []int) {
	return file_networksystem_protomessages_proto_rawDescGZIP(), []int{0}
}

func (x *NetMessageSplitscreenUserChanged) GetSlot() uint32 {
	if x != nil && x.Slot != nil {
		return *x.Slot
	}
	return 0
}

type NetMessageConnectionClosed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason *uint32 `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
}

func (x *NetMessageConnectionClosed) Reset() {
	*x = NetMessageConnectionClosed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networksystem_protomessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessageConnectionClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessageConnectionClosed) ProtoMessage() {}

func (x *NetMessageConnectionClosed) ProtoReflect() protoreflect.Message {
	mi := &file_networksystem_protomessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessageConnectionClosed.ProtoReflect.Descriptor instead.
func (*NetMessageConnectionClosed) Descriptor() ([]byte, []int) {
	return file_networksystem_protomessages_proto_rawDescGZIP(), []int{1}
}

func (x *NetMessageConnectionClosed) GetReason() uint32 {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return 0
}

type NetMessageConnectionCrashed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason *uint32 `protobuf:"varint,1,opt,name=reason" json:"reason,omitempty"`
}

func (x *NetMessageConnectionCrashed) Reset() {
	*x = NetMessageConnectionCrashed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networksystem_protomessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessageConnectionCrashed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessageConnectionCrashed) ProtoMessage() {}

func (x *NetMessageConnectionCrashed) ProtoReflect() protoreflect.Message {
	mi := &file_networksystem_protomessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessageConnectionCrashed.ProtoReflect.Descriptor instead.
func (*NetMessageConnectionCrashed) Descriptor() ([]byte, []int) {
	return file_networksystem_protomessages_proto_rawDescGZIP(), []int{2}
}

func (x *NetMessageConnectionCrashed) GetReason() uint32 {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return 0
}

type NetMessagePacketStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetMessagePacketStart) Reset() {
	*x = NetMessagePacketStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networksystem_protomessages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessagePacketStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessagePacketStart) ProtoMessage() {}

func (x *NetMessagePacketStart) ProtoReflect() protoreflect.Message {
	mi := &file_networksystem_protomessages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessagePacketStart.ProtoReflect.Descriptor instead.
func (*NetMessagePacketStart) Descriptor() ([]byte, []int) {
	return file_networksystem_protomessages_proto_rawDescGZIP(), []int{3}
}

type NetMessagePacketEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NetMessagePacketEnd) Reset() {
	*x = NetMessagePacketEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_networksystem_protomessages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMessagePacketEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMessagePacketEnd) ProtoMessage() {}

func (x *NetMessagePacketEnd) ProtoReflect() protoreflect.Message {
	mi := &file_networksystem_protomessages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMessagePacketEnd.ProtoReflect.Descriptor instead.
func (*NetMessagePacketEnd) Descriptor() ([]byte, []int) {
	return file_networksystem_protomessages_proto_rawDescGZIP(), []int{4}
}

var File_networksystem_protomessages_proto protoreflect.FileDescriptor

var file_networksystem_protomessages_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x6f, 0x74, 0x61, 0x22, 0x36, 0x0a, 0x20, 0x4e, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x73, 0x63, 0x72, 0x65,
	0x65, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6c, 0x6f,
	0x74, 0x22, 0x34, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x1b, 0x4e, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x72, 0x61, 0x73, 0x68, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x17,
	0x0a, 0x15, 0x4e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x74,
	0x61, 0x62, 0x75, 0x66, 0x66, 0x2f, 0x6d, 0x61, 0x6e, 0x74, 0x61, 0x2f, 0x64, 0x6f, 0x74, 0x61,
	0x3b, 0x64, 0x6f, 0x74, 0x61,
}

var (
	file_networksystem_protomessages_proto_rawDescOnce sync.Once
	file_networksystem_protomessages_proto_rawDescData = file_networksystem_protomessages_proto_rawDesc
)

func file_networksystem_protomessages_proto_rawDescGZIP() []byte {
	file_networksystem_protomessages_proto_rawDescOnce.Do(func() {
		file_networksystem_protomessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_networksystem_protomessages_proto_rawDescData)
	})
	return file_networksystem_protomessages_proto_rawDescData
}

var file_networksystem_protomessages_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_networksystem_protomessages_proto_goTypes = []interface{}{
	(*NetMessageSplitscreenUserChanged)(nil), // 0: dota.NetMessageSplitscreenUserChanged
	(*NetMessageConnectionClosed)(nil),       // 1: dota.NetMessageConnectionClosed
	(*NetMessageConnectionCrashed)(nil),      // 2: dota.NetMessageConnectionCrashed
	(*NetMessagePacketStart)(nil),            // 3: dota.NetMessagePacketStart
	(*NetMessagePacketEnd)(nil),              // 4: dota.NetMessagePacketEnd
}
var file_networksystem_protomessages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_networksystem_protomessages_proto_init() }
func file_networksystem_protomessages_proto_init() {
	if File_networksystem_protomessages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_networksystem_protomessages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessageSplitscreenUserChanged); i {
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
		file_networksystem_protomessages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessageConnectionClosed); i {
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
		file_networksystem_protomessages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessageConnectionCrashed); i {
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
		file_networksystem_protomessages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessagePacketStart); i {
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
		file_networksystem_protomessages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMessagePacketEnd); i {
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
			RawDescriptor: file_networksystem_protomessages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_networksystem_protomessages_proto_goTypes,
		DependencyIndexes: file_networksystem_protomessages_proto_depIdxs,
		MessageInfos:      file_networksystem_protomessages_proto_msgTypes,
	}.Build()
	File_networksystem_protomessages_proto = out.File
	file_networksystem_protomessages_proto_rawDesc = nil
	file_networksystem_protomessages_proto_goTypes = nil
	file_networksystem_protomessages_proto_depIdxs = nil
}

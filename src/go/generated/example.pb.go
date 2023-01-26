// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.0
// source: example.proto

package generated

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// GetIoTDataReq defines message struct for client request
type GetIoTDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *IoTThing `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *GetIoTDataReq) Reset() {
	*x = GetIoTDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIoTDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIoTDataReq) ProtoMessage() {}

func (x *GetIoTDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIoTDataReq.ProtoReflect.Descriptor instead.
func (*GetIoTDataReq) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{0}
}

func (x *GetIoTDataReq) GetDevice() *IoTThing {
	if x != nil {
		return x.Device
	}
	return nil
}

// GetIoTDataRes defines message struct for server response
type GetIoTDataRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*IoTData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"` // repeated keyword is used to create array or slice (in go terminology) of a given structure
}

func (x *GetIoTDataRes) Reset() {
	*x = GetIoTDataRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIoTDataRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIoTDataRes) ProtoMessage() {}

func (x *GetIoTDataRes) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIoTDataRes.ProtoReflect.Descriptor instead.
func (*GetIoTDataRes) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{1}
}

func (x *GetIoTDataRes) GetData() []*IoTData {
	if x != nil {
		return x.Data
	}
	return nil
}

// IoTData shows example of message structure that can be used across multiple other messages
type IoTData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnyInt32Val  int32  `protobuf:"varint,1,opt,name=anyInt32Val,proto3" json:"anyInt32Val,omitempty"`
	AnyInt64Val  int64  `protobuf:"varint,2,opt,name=anyInt64Val,proto3" json:"anyInt64Val,omitempty"`
	AnyStringVal string `protobuf:"bytes,3,opt,name=anyStringVal,proto3" json:"anyStringVal,omitempty"`
	AnyBoolVal   bool   `protobuf:"varint,4,opt,name=anyBoolVal,proto3" json:"anyBoolVal,omitempty"`
}

func (x *IoTData) Reset() {
	*x = IoTData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IoTData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IoTData) ProtoMessage() {}

func (x *IoTData) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IoTData.ProtoReflect.Descriptor instead.
func (*IoTData) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{2}
}

func (x *IoTData) GetAnyInt32Val() int32 {
	if x != nil {
		return x.AnyInt32Val
	}
	return 0
}

func (x *IoTData) GetAnyInt64Val() int64 {
	if x != nil {
		return x.AnyInt64Val
	}
	return 0
}

func (x *IoTData) GetAnyStringVal() string {
	if x != nil {
		return x.AnyStringVal
	}
	return ""
}

func (x *IoTData) GetAnyBoolVal() bool {
	if x != nil {
		return x.AnyBoolVal
	}
	return false
}

// GetIoTDataStreamReq is a request message structure
type GetIoTDataStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device *IoTThing `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *GetIoTDataStreamReq) Reset() {
	*x = GetIoTDataStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIoTDataStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIoTDataStreamReq) ProtoMessage() {}

func (x *GetIoTDataStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIoTDataStreamReq.ProtoReflect.Descriptor instead.
func (*GetIoTDataStreamReq) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{3}
}

func (x *GetIoTDataStreamReq) GetDevice() *IoTThing {
	if x != nil {
		return x.Device
	}
	return nil
}

// GetIoTDataStreamReq is a response message structure
type GetIoTDataStreamRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *IoTData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetIoTDataStreamRes) Reset() {
	*x = GetIoTDataStreamRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIoTDataStreamRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIoTDataStreamRes) ProtoMessage() {}

func (x *GetIoTDataStreamRes) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIoTDataStreamRes.ProtoReflect.Descriptor instead.
func (*GetIoTDataStreamRes) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{4}
}

func (x *GetIoTDataStreamRes) GetData() *IoTData {
	if x != nil {
		return x.Data
	}
	return nil
}

// IoTThing is a message structure reused across multiple messages
// In this example it shows the structure for the Device
type IoTThing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IoTThing) Reset() {
	*x = IoTThing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IoTThing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IoTThing) ProtoMessage() {}

func (x *IoTThing) ProtoReflect() protoreflect.Message {
	mi := &file_example_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IoTThing.ProtoReflect.Descriptor instead.
func (*IoTThing) Descriptor() ([]byte, []int) {
	return file_example_proto_rawDescGZIP(), []int{5}
}

func (x *IoTThing) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_example_proto protoreflect.FileDescriptor

var file_example_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x22, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x12, 0x32, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6f, 0x54, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6f, 0x54, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x49, 0x6f, 0x54, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x6e, 0x79, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6e, 0x79, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x6e,
	0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6e,
	0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x61, 0x6e, 0x79, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x49, 0x6f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x12, 0x32, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6f, 0x54, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6f,
	0x54, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x08, 0x49,
	0x6f, 0x54, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xc1, 0x01, 0x0a, 0x07,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49, 0x6f,
	0x54, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x49, 0x6f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x25, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6f, 0x54, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x12, 0x5a, 0x10, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_proto_rawDescOnce sync.Once
	file_example_proto_rawDescData = file_example_proto_rawDesc
)

func file_example_proto_rawDescGZIP() []byte {
	file_example_proto_rawDescOnce.Do(func() {
		file_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_proto_rawDescData)
	})
	return file_example_proto_rawDescData
}

var file_example_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_example_proto_goTypes = []interface{}{
	(*GetIoTDataReq)(nil),       // 0: example.protobuf.GetIoTDataReq
	(*GetIoTDataRes)(nil),       // 1: example.protobuf.GetIoTDataRes
	(*IoTData)(nil),             // 2: example.protobuf.IoTData
	(*GetIoTDataStreamReq)(nil), // 3: example.protobuf.GetIoTDataStreamReq
	(*GetIoTDataStreamRes)(nil), // 4: example.protobuf.GetIoTDataStreamRes
	(*IoTThing)(nil),            // 5: example.protobuf.IoTThing
}
var file_example_proto_depIdxs = []int32{
	5, // 0: example.protobuf.GetIoTDataReq.device:type_name -> example.protobuf.IoTThing
	2, // 1: example.protobuf.GetIoTDataRes.data:type_name -> example.protobuf.IoTData
	5, // 2: example.protobuf.GetIoTDataStreamReq.device:type_name -> example.protobuf.IoTThing
	2, // 3: example.protobuf.GetIoTDataStreamRes.data:type_name -> example.protobuf.IoTData
	0, // 4: example.protobuf.Example.GetIoTData:input_type -> example.protobuf.GetIoTDataReq
	3, // 5: example.protobuf.Example.GetIoTDataStream:input_type -> example.protobuf.GetIoTDataStreamReq
	1, // 6: example.protobuf.Example.GetIoTData:output_type -> example.protobuf.GetIoTDataRes
	4, // 7: example.protobuf.Example.GetIoTDataStream:output_type -> example.protobuf.GetIoTDataStreamRes
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_example_proto_init() }
func file_example_proto_init() {
	if File_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIoTDataReq); i {
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
		file_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIoTDataRes); i {
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
		file_example_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IoTData); i {
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
		file_example_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIoTDataStreamReq); i {
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
		file_example_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIoTDataStreamRes); i {
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
		file_example_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IoTThing); i {
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
			RawDescriptor: file_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_proto_goTypes,
		DependencyIndexes: file_example_proto_depIdxs,
		MessageInfos:      file_example_proto_msgTypes,
	}.Build()
	File_example_proto = out.File
	file_example_proto_rawDesc = nil
	file_example_proto_goTypes = nil
	file_example_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	// GetIoTData example of unary call, where client sends one request and server responds back with a request and call is closed.
	GetIoTData(ctx context.Context, in *GetIoTDataReq, opts ...grpc.CallOption) (*GetIoTDataRes, error)
	// GetIoTDataStream example of server stream, where client sends a request and server responds back with stream of data. Call remains open as long as server is sending the data, after which server closes the stream with NIL or EOF.
	GetIoTDataStream(ctx context.Context, in *GetIoTDataStreamReq, opts ...grpc.CallOption) (Example_GetIoTDataStreamClient, error)
}

type exampleClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleClient(cc grpc.ClientConnInterface) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) GetIoTData(ctx context.Context, in *GetIoTDataReq, opts ...grpc.CallOption) (*GetIoTDataRes, error) {
	out := new(GetIoTDataRes)
	err := c.cc.Invoke(ctx, "/example.protobuf.Example/GetIoTData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) GetIoTDataStream(ctx context.Context, in *GetIoTDataStreamReq, opts ...grpc.CallOption) (Example_GetIoTDataStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Example_serviceDesc.Streams[0], "/example.protobuf.Example/GetIoTDataStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &exampleGetIoTDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Example_GetIoTDataStreamClient interface {
	Recv() (*GetIoTDataStreamRes, error)
	grpc.ClientStream
}

type exampleGetIoTDataStreamClient struct {
	grpc.ClientStream
}

func (x *exampleGetIoTDataStreamClient) Recv() (*GetIoTDataStreamRes, error) {
	m := new(GetIoTDataStreamRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	// GetIoTData example of unary call, where client sends one request and server responds back with a request and call is closed.
	GetIoTData(context.Context, *GetIoTDataReq) (*GetIoTDataRes, error)
	// GetIoTDataStream example of server stream, where client sends a request and server responds back with stream of data. Call remains open as long as server is sending the data, after which server closes the stream with NIL or EOF.
	GetIoTDataStream(*GetIoTDataStreamReq, Example_GetIoTDataStreamServer) error
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) GetIoTData(context.Context, *GetIoTDataReq) (*GetIoTDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIoTData not implemented")
}
func (*UnimplementedExampleServer) GetIoTDataStream(*GetIoTDataStreamReq, Example_GetIoTDataStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetIoTDataStream not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_GetIoTData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIoTDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).GetIoTData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.protobuf.Example/GetIoTData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).GetIoTData(ctx, req.(*GetIoTDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example_GetIoTDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetIoTDataStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExampleServer).GetIoTDataStream(m, &exampleGetIoTDataStreamServer{stream})
}

type Example_GetIoTDataStreamServer interface {
	Send(*GetIoTDataStreamRes) error
	grpc.ServerStream
}

type exampleGetIoTDataStreamServer struct {
	grpc.ServerStream
}

func (x *exampleGetIoTDataStreamServer) Send(m *GetIoTDataStreamRes) error {
	return x.ServerStream.SendMsg(m)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.protobuf.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIoTData",
			Handler:    _Example_GetIoTData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetIoTDataStream",
			Handler:       _Example_GetIoTDataStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example.proto",
}
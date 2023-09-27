// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: hello/v1/hello.proto

package pb

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

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_v1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_hello_v1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_hello_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *Hello) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Hello) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_hello_v1_hello_proto protoreflect.FileDescriptor

var file_hello_v1_hello_proto_rawDesc = []byte{
	0x0a, 0x14, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x76, 0x31,
	0x22, 0x2b, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hello_v1_hello_proto_rawDescOnce sync.Once
	file_hello_v1_hello_proto_rawDescData = file_hello_v1_hello_proto_rawDesc
)

func file_hello_v1_hello_proto_rawDescGZIP() []byte {
	file_hello_v1_hello_proto_rawDescOnce.Do(func() {
		file_hello_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_v1_hello_proto_rawDescData)
	})
	return file_hello_v1_hello_proto_rawDescData
}

var file_hello_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hello_v1_hello_proto_goTypes = []interface{}{
	(*Hello)(nil), // 0: hello.v1.Hello
}
var file_hello_v1_hello_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hello_v1_hello_proto_init() }
func file_hello_v1_hello_proto_init() {
	if File_hello_v1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_v1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
			RawDescriptor: file_hello_v1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hello_v1_hello_proto_goTypes,
		DependencyIndexes: file_hello_v1_hello_proto_depIdxs,
		MessageInfos:      file_hello_v1_hello_proto_msgTypes,
	}.Build()
	File_hello_v1_hello_proto = out.File
	file_hello_v1_hello_proto_rawDesc = nil
	file_hello_v1_hello_proto_goTypes = nil
	file_hello_v1_hello_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: author.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorRequest) Reset() {
	*x = AuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorRequest) ProtoMessage() {}

func (x *AuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorRequest.ProtoReflect.Descriptor instead.
func (*AuthorRequest) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AuthorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author []*Author `protobuf:"bytes,1,rep,name=author,proto3" json:"author,omitempty"`
}

func (x *AuthorList) Reset() {
	*x = AuthorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorList) ProtoMessage() {}

func (x *AuthorList) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorList.ProtoReflect.Descriptor instead.
func (*AuthorList) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorList) GetAuthor() []*Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type AuthorId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AuthorId) Reset() {
	*x = AuthorId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorId) ProtoMessage() {}

func (x *AuthorId) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorId.ProtoReflect.Descriptor instead.
func (*AuthorId) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AuthorUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AuthorUpdate) Reset() {
	*x = AuthorUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_author_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorUpdate) ProtoMessage() {}

func (x *AuthorUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_author_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorUpdate.ProtoReflect.Descriptor instead.
func (*AuthorUpdate) Descriptor() ([]byte, []int) {
	return file_author_proto_rawDescGZIP(), []int{4}
}

func (x *AuthorUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AuthorUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_author_proto protoreflect.FileDescriptor

var file_author_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x23, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x34, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x1a, 0x0a, 0x08,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x83, 0x02, 0x0a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x0e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2e,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x0e,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x32,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_author_proto_rawDescOnce sync.Once
	file_author_proto_rawDescData = file_author_proto_rawDesc
)

func file_author_proto_rawDescGZIP() []byte {
	file_author_proto_rawDescOnce.Do(func() {
		file_author_proto_rawDescData = protoimpl.X.CompressGZIP(file_author_proto_rawDescData)
	})
	return file_author_proto_rawDescData
}

var file_author_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_author_proto_goTypes = []any{
	(*Author)(nil),        // 0: author.Author
	(*AuthorRequest)(nil), // 1: author.AuthorRequest
	(*AuthorList)(nil),    // 2: author.AuthorList
	(*AuthorId)(nil),      // 3: author.AuthorId
	(*AuthorUpdate)(nil),  // 4: author.AuthorUpdate
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_author_proto_depIdxs = []int32{
	0, // 0: author.AuthorList.author:type_name -> author.Author
	5, // 1: author.AuthorService.Get:input_type -> google.protobuf.Empty
	1, // 2: author.AuthorService.Create:input_type -> author.AuthorRequest
	3, // 3: author.AuthorService.GetOne:input_type -> author.AuthorId
	4, // 4: author.AuthorService.Update:input_type -> author.AuthorUpdate
	3, // 5: author.AuthorService.Delete:input_type -> author.AuthorId
	2, // 6: author.AuthorService.Get:output_type -> author.AuthorList
	0, // 7: author.AuthorService.Create:output_type -> author.Author
	0, // 8: author.AuthorService.GetOne:output_type -> author.Author
	0, // 9: author.AuthorService.Update:output_type -> author.Author
	5, // 10: author.AuthorService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_author_proto_init() }
func file_author_proto_init() {
	if File_author_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_author_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Author); i {
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
		file_author_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AuthorRequest); i {
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
		file_author_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AuthorList); i {
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
		file_author_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AuthorId); i {
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
		file_author_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AuthorUpdate); i {
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
			RawDescriptor: file_author_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_author_proto_goTypes,
		DependencyIndexes: file_author_proto_depIdxs,
		MessageInfos:      file_author_proto_msgTypes,
	}.Build()
	File_author_proto = out.File
	file_author_proto_rawDesc = nil
	file_author_proto_goTypes = nil
	file_author_proto_depIdxs = nil
}

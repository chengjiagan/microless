// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/urlshorten.proto

package urlshorten

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	proto "microless/socialnetwork/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComposeUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []string `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ComposeUrlsRequest) Reset() {
	*x = ComposeUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urlshorten_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeUrlsRequest) ProtoMessage() {}

func (x *ComposeUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshorten_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeUrlsRequest.ProtoReflect.Descriptor instead.
func (*ComposeUrlsRequest) Descriptor() ([]byte, []int) {
	return file_proto_urlshorten_proto_rawDescGZIP(), []int{0}
}

func (x *ComposeUrlsRequest) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

type ComposeUrlsRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []*proto.Url `protobuf:"bytes,2,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *ComposeUrlsRespond) Reset() {
	*x = ComposeUrlsRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urlshorten_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComposeUrlsRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComposeUrlsRespond) ProtoMessage() {}

func (x *ComposeUrlsRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshorten_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComposeUrlsRespond.ProtoReflect.Descriptor instead.
func (*ComposeUrlsRespond) Descriptor() ([]byte, []int) {
	return file_proto_urlshorten_proto_rawDescGZIP(), []int{1}
}

func (x *ComposeUrlsRespond) GetUrls() []*proto.Url {
	if x != nil {
		return x.Urls
	}
	return nil
}

type GetExtendedUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShortenedUrls []string `protobuf:"bytes,1,rep,name=shortened_urls,json=shortenedUrls,proto3" json:"shortened_urls,omitempty"`
}

func (x *GetExtendedUrlsRequest) Reset() {
	*x = GetExtendedUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urlshorten_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExtendedUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExtendedUrlsRequest) ProtoMessage() {}

func (x *GetExtendedUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshorten_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExtendedUrlsRequest.ProtoReflect.Descriptor instead.
func (*GetExtendedUrlsRequest) Descriptor() ([]byte, []int) {
	return file_proto_urlshorten_proto_rawDescGZIP(), []int{2}
}

func (x *GetExtendedUrlsRequest) GetShortenedUrls() []string {
	if x != nil {
		return x.ShortenedUrls
	}
	return nil
}

type GetExtendedUrlsRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls []string `protobuf:"bytes,2,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *GetExtendedUrlsRespond) Reset() {
	*x = GetExtendedUrlsRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_urlshorten_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExtendedUrlsRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExtendedUrlsRespond) ProtoMessage() {}

func (x *GetExtendedUrlsRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_urlshorten_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExtendedUrlsRespond.ProtoReflect.Descriptor instead.
func (*GetExtendedUrlsRespond) Descriptor() ([]byte, []int) {
	return file_proto_urlshorten_proto_rawDescGZIP(), []int{3}
}

func (x *GetExtendedUrlsRespond) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

var File_proto_urlshorten_proto protoreflect.FileDescriptor

var file_proto_urlshorten_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x73, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x1a, 0x10, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28,
	0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x30,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73,
	0x22, 0x3f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x55,
	0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x72, 0x6c,
	0x73, 0x22, 0x2c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x32,
	0x9e, 0x02, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65,
	0x55, 0x72, 0x6c, 0x73, 0x12, 0x36, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73,
	0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x12, 0x89, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73,
	0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x75,
	0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x42, 0x2a, 0x5a, 0x28, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_urlshorten_proto_rawDescOnce sync.Once
	file_proto_urlshorten_proto_rawDescData = file_proto_urlshorten_proto_rawDesc
)

func file_proto_urlshorten_proto_rawDescGZIP() []byte {
	file_proto_urlshorten_proto_rawDescOnce.Do(func() {
		file_proto_urlshorten_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_urlshorten_proto_rawDescData)
	})
	return file_proto_urlshorten_proto_rawDescData
}

var file_proto_urlshorten_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_urlshorten_proto_goTypes = []interface{}{
	(*ComposeUrlsRequest)(nil),     // 0: microless.socialnetwork.urlshorten.ComposeUrlsRequest
	(*ComposeUrlsRespond)(nil),     // 1: microless.socialnetwork.urlshorten.ComposeUrlsRespond
	(*GetExtendedUrlsRequest)(nil), // 2: microless.socialnetwork.urlshorten.GetExtendedUrlsRequest
	(*GetExtendedUrlsRespond)(nil), // 3: microless.socialnetwork.urlshorten.GetExtendedUrlsRespond
	(*proto.Url)(nil),              // 4: microless.socialnetwork.Url
}
var file_proto_urlshorten_proto_depIdxs = []int32{
	4, // 0: microless.socialnetwork.urlshorten.ComposeUrlsRespond.urls:type_name -> microless.socialnetwork.Url
	0, // 1: microless.socialnetwork.urlshorten.UrlShortenService.ComposeUrls:input_type -> microless.socialnetwork.urlshorten.ComposeUrlsRequest
	2, // 2: microless.socialnetwork.urlshorten.UrlShortenService.GetExtendedUrls:input_type -> microless.socialnetwork.urlshorten.GetExtendedUrlsRequest
	1, // 3: microless.socialnetwork.urlshorten.UrlShortenService.ComposeUrls:output_type -> microless.socialnetwork.urlshorten.ComposeUrlsRespond
	3, // 4: microless.socialnetwork.urlshorten.UrlShortenService.GetExtendedUrls:output_type -> microless.socialnetwork.urlshorten.GetExtendedUrlsRespond
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_urlshorten_proto_init() }
func file_proto_urlshorten_proto_init() {
	if File_proto_urlshorten_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_urlshorten_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeUrlsRequest); i {
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
		file_proto_urlshorten_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComposeUrlsRespond); i {
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
		file_proto_urlshorten_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExtendedUrlsRequest); i {
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
		file_proto_urlshorten_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExtendedUrlsRespond); i {
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
			RawDescriptor: file_proto_urlshorten_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_urlshorten_proto_goTypes,
		DependencyIndexes: file_proto_urlshorten_proto_depIdxs,
		MessageInfos:      file_proto_urlshorten_proto_msgTypes,
	}.Build()
	File_proto_urlshorten_proto = out.File
	file_proto_urlshorten_proto_rawDesc = nil
	file_proto_urlshorten_proto_goTypes = nil
	file_proto_urlshorten_proto_depIdxs = nil
}

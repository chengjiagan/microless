// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/page.proto

package page

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	proto "microless/media/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReadPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId     string `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	ReviewStart int32  `protobuf:"varint,2,opt,name=review_start,json=reviewStart,proto3" json:"review_start,omitempty"`
	ReviewStop  int32  `protobuf:"varint,3,opt,name=review_stop,json=reviewStop,proto3" json:"review_stop,omitempty"`
}

func (x *ReadPageRequest) Reset() {
	*x = ReadPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPageRequest) ProtoMessage() {}

func (x *ReadPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPageRequest.ProtoReflect.Descriptor instead.
func (*ReadPageRequest) Descriptor() ([]byte, []int) {
	return file_proto_page_proto_rawDescGZIP(), []int{0}
}

func (x *ReadPageRequest) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *ReadPageRequest) GetReviewStart() int32 {
	if x != nil {
		return x.ReviewStart
	}
	return 0
}

func (x *ReadPageRequest) GetReviewStop() int32 {
	if x != nil {
		return x.ReviewStop
	}
	return 0
}

type ReadPageRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieInfo *proto.MovieInfo  `protobuf:"bytes,1,opt,name=movie_info,json=movieInfo,proto3" json:"movie_info,omitempty"`
	CastInfos []*proto.CastInfo `protobuf:"bytes,2,rep,name=cast_infos,json=castInfos,proto3" json:"cast_infos,omitempty"`
	Plot      string            `protobuf:"bytes,3,opt,name=plot,proto3" json:"plot,omitempty"`
	Reviews   []*proto.Review   `protobuf:"bytes,4,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *ReadPageRespond) Reset() {
	*x = ReadPageRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPageRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPageRespond) ProtoMessage() {}

func (x *ReadPageRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPageRespond.ProtoReflect.Descriptor instead.
func (*ReadPageRespond) Descriptor() ([]byte, []int) {
	return file_proto_page_proto_rawDescGZIP(), []int{1}
}

func (x *ReadPageRespond) GetMovieInfo() *proto.MovieInfo {
	if x != nil {
		return x.MovieInfo
	}
	return nil
}

func (x *ReadPageRespond) GetCastInfos() []*proto.CastInfo {
	if x != nil {
		return x.CastInfos
	}
	return nil
}

func (x *ReadPageRespond) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *ReadPageRespond) GetReviews() []*proto.Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

var File_proto_page_proto protoreflect.FileDescriptor

var file_proto_page_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x6f, 0x70, 0x22, 0xcd, 0x01, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x61, 0x73,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x43, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x63, 0x61, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x32, 0x88, 0x01, 0x0a, 0x0b, 0x50,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x08, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x64, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x7b, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x73, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_page_proto_rawDescOnce sync.Once
	file_proto_page_proto_rawDescData = file_proto_page_proto_rawDesc
)

func file_proto_page_proto_rawDescGZIP() []byte {
	file_proto_page_proto_rawDescOnce.Do(func() {
		file_proto_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_page_proto_rawDescData)
	})
	return file_proto_page_proto_rawDescData
}

var file_proto_page_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_page_proto_goTypes = []interface{}{
	(*ReadPageRequest)(nil), // 0: microless.media.page.ReadPageRequest
	(*ReadPageRespond)(nil), // 1: microless.media.page.ReadPageRespond
	(*proto.MovieInfo)(nil), // 2: microless.media.MovieInfo
	(*proto.CastInfo)(nil),  // 3: microless.media.CastInfo
	(*proto.Review)(nil),    // 4: microless.media.Review
}
var file_proto_page_proto_depIdxs = []int32{
	2, // 0: microless.media.page.ReadPageRespond.movie_info:type_name -> microless.media.MovieInfo
	3, // 1: microless.media.page.ReadPageRespond.cast_infos:type_name -> microless.media.CastInfo
	4, // 2: microless.media.page.ReadPageRespond.reviews:type_name -> microless.media.Review
	0, // 3: microless.media.page.PageService.ReadPage:input_type -> microless.media.page.ReadPageRequest
	1, // 4: microless.media.page.PageService.ReadPage:output_type -> microless.media.page.ReadPageRespond
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_page_proto_init() }
func file_proto_page_proto_init() {
	if File_proto_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPageRequest); i {
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
		file_proto_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPageRespond); i {
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
			RawDescriptor: file_proto_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_page_proto_goTypes,
		DependencyIndexes: file_proto_page_proto_depIdxs,
		MessageInfos:      file_proto_page_proto_msgTypes,
	}.Build()
	File_proto_page_proto = out.File
	file_proto_page_proto_rawDesc = nil
	file_proto_page_proto_goTypes = nil
	file_proto_page_proto_depIdxs = nil
}

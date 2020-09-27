// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.6.1
// source: binarizer.proto

package binarizer

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

type BinarizationMethod int32

const (
	BinarizationMethod_UNDEFINED BinarizationMethod = 0
	BinarizationMethod_SAUVOLA   BinarizationMethod = 1
	BinarizationMethod_NIBLACK   BinarizationMethod = 2
	BinarizationMethod_OTSU      BinarizationMethod = 3
)

// Enum value maps for BinarizationMethod.
var (
	BinarizationMethod_name = map[int32]string{
		0: "UNDEFINED",
		1: "SAUVOLA",
		2: "NIBLACK",
		3: "OTSU",
	}
	BinarizationMethod_value = map[string]int32{
		"UNDEFINED": 0,
		"SAUVOLA":   1,
		"NIBLACK":   2,
		"OTSU":      3,
	}
)

func (x BinarizationMethod) Enum() *BinarizationMethod {
	p := new(BinarizationMethod)
	*p = x
	return p
}

func (x BinarizationMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinarizationMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_binarizer_proto_enumTypes[0].Descriptor()
}

func (BinarizationMethod) Type() protoreflect.EnumType {
	return &file_binarizer_proto_enumTypes[0]
}

func (x BinarizationMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BinarizationMethod.Descriptor instead.
func (BinarizationMethod) EnumDescriptor() ([]byte, []int) {
	return file_binarizer_proto_rawDescGZIP(), []int{0}
}

type BinarizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images             [][]byte           `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	BinarizationMethod BinarizationMethod `protobuf:"varint,2,opt,name=binarization_method,json=binarizationMethod,proto3,enum=binarizer.BinarizationMethod" json:"binarization_method,omitempty"`
}

func (x *BinarizeRequest) Reset() {
	*x = BinarizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binarizer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinarizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinarizeRequest) ProtoMessage() {}

func (x *BinarizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_binarizer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinarizeRequest.ProtoReflect.Descriptor instead.
func (*BinarizeRequest) Descriptor() ([]byte, []int) {
	return file_binarizer_proto_rawDescGZIP(), []int{0}
}

func (x *BinarizeRequest) GetImages() [][]byte {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *BinarizeRequest) GetBinarizationMethod() BinarizationMethod {
	if x != nil {
		return x.BinarizationMethod
	}
	return BinarizationMethod_UNDEFINED
}

type BinarizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images  [][]byte `protobuf:"bytes,1,rep,name=images,proto3" json:"images,omitempty"`
	OutPath string   `protobuf:"bytes,2,opt,name=out_path,json=outPath,proto3" json:"out_path,omitempty"`
}

func (x *BinarizeResponse) Reset() {
	*x = BinarizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_binarizer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinarizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinarizeResponse) ProtoMessage() {}

func (x *BinarizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_binarizer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinarizeResponse.ProtoReflect.Descriptor instead.
func (*BinarizeResponse) Descriptor() ([]byte, []int) {
	return file_binarizer_proto_rawDescGZIP(), []int{1}
}

func (x *BinarizeResponse) GetImages() [][]byte {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *BinarizeResponse) GetOutPath() string {
	if x != nil {
		return x.OutPath
	}
	return ""
}

var File_binarizer_proto protoreflect.FileDescriptor

var file_binarizer_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x22, 0x79, 0x0a, 0x0f,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x13, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72,
	0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x12, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x45, 0x0a, 0x10, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x50, 0x61, 0x74, 0x68, 0x2a, 0x47,
	0x0a, 0x12, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x41, 0x55, 0x56, 0x4f, 0x4c, 0x41, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x49, 0x42, 0x4c, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x4f, 0x54, 0x53, 0x55, 0x10, 0x03, 0x32, 0x52, 0x0a, 0x09, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x69, 0x7a, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65,
	0x12, 0x1a, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e,
	0x61, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x50, 0x0a, 0x2f, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x65, 0x69, 0x74, 0x6f, 0x6c, 0x2e,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x42, 0x0e,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x0b, 0x2e, 0x3b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_binarizer_proto_rawDescOnce sync.Once
	file_binarizer_proto_rawDescData = file_binarizer_proto_rawDesc
)

func file_binarizer_proto_rawDescGZIP() []byte {
	file_binarizer_proto_rawDescOnce.Do(func() {
		file_binarizer_proto_rawDescData = protoimpl.X.CompressGZIP(file_binarizer_proto_rawDescData)
	})
	return file_binarizer_proto_rawDescData
}

var file_binarizer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_binarizer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_binarizer_proto_goTypes = []interface{}{
	(BinarizationMethod)(0),  // 0: binarizer.BinarizationMethod
	(*BinarizeRequest)(nil),  // 1: binarizer.BinarizeRequest
	(*BinarizeResponse)(nil), // 2: binarizer.BinarizeResponse
}
var file_binarizer_proto_depIdxs = []int32{
	0, // 0: binarizer.BinarizeRequest.binarization_method:type_name -> binarizer.BinarizationMethod
	1, // 1: binarizer.Binarizer.Binarize:input_type -> binarizer.BinarizeRequest
	2, // 2: binarizer.Binarizer.Binarize:output_type -> binarizer.BinarizeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_binarizer_proto_init() }
func file_binarizer_proto_init() {
	if File_binarizer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_binarizer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinarizeRequest); i {
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
		file_binarizer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinarizeResponse); i {
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
			RawDescriptor: file_binarizer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_binarizer_proto_goTypes,
		DependencyIndexes: file_binarizer_proto_depIdxs,
		EnumInfos:         file_binarizer_proto_enumTypes,
		MessageInfos:      file_binarizer_proto_msgTypes,
	}.Build()
	File_binarizer_proto = out.File
	file_binarizer_proto_rawDesc = nil
	file_binarizer_proto_goTypes = nil
	file_binarizer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BinarizerClient is the client API for Binarizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BinarizerClient interface {
	Binarize(ctx context.Context, in *BinarizeRequest, opts ...grpc.CallOption) (*BinarizeResponse, error)
}

type binarizerClient struct {
	cc grpc.ClientConnInterface
}

func NewBinarizerClient(cc grpc.ClientConnInterface) BinarizerClient {
	return &binarizerClient{cc}
}

func (c *binarizerClient) Binarize(ctx context.Context, in *BinarizeRequest, opts ...grpc.CallOption) (*BinarizeResponse, error) {
	out := new(BinarizeResponse)
	err := c.cc.Invoke(ctx, "/binarizer.Binarizer/Binarize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BinarizerServer is the server API for Binarizer service.
type BinarizerServer interface {
	Binarize(context.Context, *BinarizeRequest) (*BinarizeResponse, error)
}

// UnimplementedBinarizerServer can be embedded to have forward compatible implementations.
type UnimplementedBinarizerServer struct {
}

func (*UnimplementedBinarizerServer) Binarize(context.Context, *BinarizeRequest) (*BinarizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Binarize not implemented")
}

func RegisterBinarizerServer(s *grpc.Server, srv BinarizerServer) {
	s.RegisterService(&_Binarizer_serviceDesc, srv)
}

func _Binarizer_Binarize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinarizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BinarizerServer).Binarize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/binarizer.Binarizer/Binarize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BinarizerServer).Binarize(ctx, req.(*BinarizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Binarizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "binarizer.Binarizer",
	HandlerType: (*BinarizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Binarize",
			Handler:    _Binarizer_Binarize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "binarizer.proto",
}
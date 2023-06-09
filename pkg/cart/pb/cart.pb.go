// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/cart/pb/cart.proto

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

// Add to cart
type AddToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *AddToCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddToCartResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AddToCartResponce) Reset() {
	*x = AddToCartResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToCartResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartResponce) ProtoMessage() {}

func (x *AddToCartResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartResponce.ProtoReflect.Descriptor instead.
func (*AddToCartResponce) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{1}
}

func (x *AddToCartResponce) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddToCartResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId int64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	UsreId    int64 `protobuf:"varint,2,opt,name=usreId,proto3" json:"usreId,omitempty"`
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveFromCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *RemoveFromCartRequest) GetUsreId() int64 {
	if x != nil {
		return x.UsreId
	}
	return 0
}

type RemoveFromcartResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RemoveFromcartResponce) Reset() {
	*x = RemoveFromcartResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_cart_pb_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFromcartResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromcartResponce) ProtoMessage() {}

func (x *RemoveFromcartResponce) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cart_pb_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromcartResponce.ProtoReflect.Descriptor instead.
func (*RemoveFromcartResponce) Descriptor() ([]byte, []int) {
	return file_pkg_cart_pb_cart_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFromcartResponce) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RemoveFromcartResponce) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_cart_pb_cart_proto protoreflect.FileDescriptor

var file_pkg_cart_pb_cart_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x48,
	0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4d, 0x0a, 0x15, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x72, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x72, 0x65, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x63, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0x98, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x16, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x63, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_cart_pb_cart_proto_rawDescOnce sync.Once
	file_pkg_cart_pb_cart_proto_rawDescData = file_pkg_cart_pb_cart_proto_rawDesc
)

func file_pkg_cart_pb_cart_proto_rawDescGZIP() []byte {
	file_pkg_cart_pb_cart_proto_rawDescOnce.Do(func() {
		file_pkg_cart_pb_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_cart_pb_cart_proto_rawDescData)
	})
	return file_pkg_cart_pb_cart_proto_rawDescData
}

var file_pkg_cart_pb_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_cart_pb_cart_proto_goTypes = []interface{}{
	(*AddToCartRequest)(nil),       // 0: cart.AddToCartRequest
	(*AddToCartResponce)(nil),      // 1: cart.AddToCartResponce
	(*RemoveFromCartRequest)(nil),  // 2: cart.RemoveFromCartRequest
	(*RemoveFromcartResponce)(nil), // 3: cart.RemoveFromcartResponce
}
var file_pkg_cart_pb_cart_proto_depIdxs = []int32{
	0, // 0: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	2, // 1: cart.CartService.RemoveFromCart:input_type -> cart.RemoveFromCartRequest
	1, // 2: cart.CartService.AddToCart:output_type -> cart.AddToCartResponce
	3, // 3: cart.CartService.RemoveFromCart:output_type -> cart.RemoveFromcartResponce
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_cart_pb_cart_proto_init() }
func file_pkg_cart_pb_cart_proto_init() {
	if File_pkg_cart_pb_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_cart_pb_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToCartResponce); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFromCartRequest); i {
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
		file_pkg_cart_pb_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFromcartResponce); i {
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
			RawDescriptor: file_pkg_cart_pb_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_cart_pb_cart_proto_goTypes,
		DependencyIndexes: file_pkg_cart_pb_cart_proto_depIdxs,
		MessageInfos:      file_pkg_cart_pb_cart_proto_msgTypes,
	}.Build()
	File_pkg_cart_pb_cart_proto = out.File
	file_pkg_cart_pb_cart_proto_rawDesc = nil
	file_pkg_cart_pb_cart_proto_goTypes = nil
	file_pkg_cart_pb_cart_proto_depIdxs = nil
}

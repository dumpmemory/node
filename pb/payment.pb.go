// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.25.2
// source: pb/payment.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgreementID    string `protobuf:"bytes,1,opt,name=AgreementID,proto3" json:"AgreementID,omitempty"`
	AgreementTotal string `protobuf:"bytes,2,opt,name=AgreementTotal,proto3" json:"AgreementTotal,omitempty"`
	TransactorFee  string `protobuf:"bytes,3,opt,name=TransactorFee,proto3" json:"TransactorFee,omitempty"`
	Hashlock       string `protobuf:"bytes,4,opt,name=Hashlock,proto3" json:"Hashlock,omitempty"`
	Provider       string `protobuf:"bytes,5,opt,name=Provider,proto3" json:"Provider,omitempty"`
	ChainID        int64  `protobuf:"varint,6,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_pb_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_pb_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Invoice) GetAgreementID() string {
	if x != nil {
		return x.AgreementID
	}
	return ""
}

func (x *Invoice) GetAgreementTotal() string {
	if x != nil {
		return x.AgreementTotal
	}
	return ""
}

func (x *Invoice) GetTransactorFee() string {
	if x != nil {
		return x.TransactorFee
	}
	return ""
}

func (x *Invoice) GetHashlock() string {
	if x != nil {
		return x.Hashlock
	}
	return ""
}

func (x *Invoice) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Invoice) GetChainID() int64 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

type ExchangeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Promise        *Promise `protobuf:"bytes,1,opt,name=Promise,proto3" json:"Promise,omitempty"`
	AgreementID    string   `protobuf:"bytes,2,opt,name=AgreementID,proto3" json:"AgreementID,omitempty"`
	AgreementTotal string   `protobuf:"bytes,3,opt,name=AgreementTotal,proto3" json:"AgreementTotal,omitempty"`
	Provider       string   `protobuf:"bytes,4,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Signature      string   `protobuf:"bytes,5,opt,name=Signature,proto3" json:"Signature,omitempty"`
	HermesID       string   `protobuf:"bytes,6,opt,name=HermesID,proto3" json:"HermesID,omitempty"`
	ChainID        int64    `protobuf:"varint,7,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
}

func (x *ExchangeMessage) Reset() {
	*x = ExchangeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeMessage) ProtoMessage() {}

func (x *ExchangeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeMessage.ProtoReflect.Descriptor instead.
func (*ExchangeMessage) Descriptor() ([]byte, []int) {
	return file_pb_payment_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeMessage) GetPromise() *Promise {
	if x != nil {
		return x.Promise
	}
	return nil
}

func (x *ExchangeMessage) GetAgreementID() string {
	if x != nil {
		return x.AgreementID
	}
	return ""
}

func (x *ExchangeMessage) GetAgreementTotal() string {
	if x != nil {
		return x.AgreementTotal
	}
	return ""
}

func (x *ExchangeMessage) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ExchangeMessage) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *ExchangeMessage) GetHermesID() string {
	if x != nil {
		return x.HermesID
	}
	return ""
}

func (x *ExchangeMessage) GetChainID() int64 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

type Promise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelID []byte `protobuf:"bytes,1,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"`
	Amount    string `protobuf:"bytes,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Fee       string `protobuf:"bytes,3,opt,name=Fee,proto3" json:"Fee,omitempty"`
	Hashlock  []byte `protobuf:"bytes,4,opt,name=Hashlock,proto3" json:"Hashlock,omitempty"`
	R         []byte `protobuf:"bytes,5,opt,name=R,proto3" json:"R,omitempty"`
	ChainID   int64  `protobuf:"varint,6,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
	Signature []byte `protobuf:"bytes,7,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *Promise) Reset() {
	*x = Promise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promise) ProtoMessage() {}

func (x *Promise) ProtoReflect() protoreflect.Message {
	mi := &file_pb_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promise.ProtoReflect.Descriptor instead.
func (*Promise) Descriptor() ([]byte, []int) {
	return file_pb_payment_proto_rawDescGZIP(), []int{2}
}

func (x *Promise) GetChannelID() []byte {
	if x != nil {
		return x.ChannelID
	}
	return nil
}

func (x *Promise) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Promise) GetFee() string {
	if x != nil {
		return x.Fee
	}
	return ""
}

func (x *Promise) GetHashlock() []byte {
	if x != nil {
		return x.Hashlock
	}
	return nil
}

func (x *Promise) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *Promise) GetChainID() int64 {
	if x != nil {
		return x.ChainID
	}
	return 0
}

func (x *Promise) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_pb_payment_proto protoreflect.FileDescriptor

var file_pb_payment_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xcb, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x46, 0x65, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x46,
	0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x22, 0xf2, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6d,
	0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x49, 0x44, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x65, 0x72, 0x6d, 0x65, 0x73, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x22, 0xb3, 0x01, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x46,
	0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x65, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x48, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x52, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x52, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_payment_proto_rawDescOnce sync.Once
	file_pb_payment_proto_rawDescData = file_pb_payment_proto_rawDesc
)

func file_pb_payment_proto_rawDescGZIP() []byte {
	file_pb_payment_proto_rawDescOnce.Do(func() {
		file_pb_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_payment_proto_rawDescData)
	})
	return file_pb_payment_proto_rawDescData
}

var file_pb_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_payment_proto_goTypes = []interface{}{
	(*Invoice)(nil),         // 0: pb.Invoice
	(*ExchangeMessage)(nil), // 1: pb.ExchangeMessage
	(*Promise)(nil),         // 2: pb.Promise
}
var file_pb_payment_proto_depIdxs = []int32{
	2, // 0: pb.ExchangeMessage.Promise:type_name -> pb.Promise
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_payment_proto_init() }
func file_pb_payment_proto_init() {
	if File_pb_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_pb_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeMessage); i {
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
		file_pb_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Promise); i {
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
			RawDescriptor: file_pb_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_payment_proto_goTypes,
		DependencyIndexes: file_pb_payment_proto_depIdxs,
		MessageInfos:      file_pb_payment_proto_msgTypes,
	}.Build()
	File_pb_payment_proto = out.File
	file_pb_payment_proto_rawDesc = nil
	file_pb_payment_proto_goTypes = nil
	file_pb_payment_proto_depIdxs = nil
}

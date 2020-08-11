// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.0.0
// source: sales.proto

package model

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type OrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idproduct string `protobuf:"bytes,1,opt,name=idproduct,proto3" json:"idproduct,omitempty"`
}

func (x *OrderProduct) Reset() {
	*x = OrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProduct) ProtoMessage() {}

func (x *OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProduct.ProtoReflect.Descriptor instead.
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{0}
}

func (x *OrderProduct) GetIdproduct() string {
	if x != nil {
		return x.Idproduct
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idorder        string `protobuf:"bytes,1,opt,name=Idorder,proto3" json:"Idorder,omitempty"`
	Idproduct      string `protobuf:"bytes,2,opt,name=Idproduct,proto3" json:"Idproduct,omitempty"`
	Nameproduct    string `protobuf:"bytes,3,opt,name=Nameproduct,proto3" json:"Nameproduct,omitempty"`
	Bill           string `protobuf:"bytes,4,opt,name=Bill,proto3" json:"Bill,omitempty"`
	Status_Payment string `protobuf:"bytes,5,opt,name=Status_Payment,json=StatusPayment,proto3" json:"Status_Payment,omitempty"`
	No_Tf          string `protobuf:"bytes,6,opt,name=No_Tf,json=NoTf,proto3" json:"No_Tf,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetIdorder() string {
	if x != nil {
		return x.Idorder
	}
	return ""
}

func (x *Transaction) GetIdproduct() string {
	if x != nil {
		return x.Idproduct
	}
	return ""
}

func (x *Transaction) GetNameproduct() string {
	if x != nil {
		return x.Nameproduct
	}
	return ""
}

func (x *Transaction) GetBill() string {
	if x != nil {
		return x.Bill
	}
	return ""
}

func (x *Transaction) GetStatus_Payment() string {
	if x != nil {
		return x.Status_Payment
	}
	return ""
}

func (x *Transaction) GetNo_Tf() string {
	if x != nil {
		return x.No_Tf
	}
	return ""
}

type OrderList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderList []*Transaction `protobuf:"bytes,1,rep,name=OrderList,proto3" json:"OrderList,omitempty"`
}

func (x *OrderList) Reset() {
	*x = OrderList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderList) ProtoMessage() {}

func (x *OrderList) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderList.ProtoReflect.Descriptor instead.
func (*OrderList) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{2}
}

func (x *OrderList) GetOrderList() []*Transaction {
	if x != nil {
		return x.OrderList
	}
	return nil
}

type InputPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idorder    string `protobuf:"bytes,1,opt,name=Idorder,proto3" json:"Idorder,omitempty"`
	NoTransfer string `protobuf:"bytes,2,opt,name=NoTransfer,proto3" json:"NoTransfer,omitempty"`
}

func (x *InputPayment) Reset() {
	*x = InputPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputPayment) ProtoMessage() {}

func (x *InputPayment) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputPayment.ProtoReflect.Descriptor instead.
func (*InputPayment) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{3}
}

func (x *InputPayment) GetIdorder() string {
	if x != nil {
		return x.Idorder
	}
	return ""
}

func (x *InputPayment) GetNoTransfer() string {
	if x != nil {
		return x.NoTransfer
	}
	return ""
}

type OutputPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
	Message     string       `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *OutputPayment) Reset() {
	*x = OutputPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputPayment) ProtoMessage() {}

func (x *OutputPayment) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputPayment.ProtoReflect.Descriptor instead.
func (*OutputPayment) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{4}
}

func (x *OutputPayment) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *OutputPayment) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type OutTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=Transaction,proto3" json:"Transaction,omitempty"`
	Message     string       `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *OutTransaction) Reset() {
	*x = OutTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sales_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutTransaction) ProtoMessage() {}

func (x *OutTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_sales_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutTransaction.ProtoReflect.Descriptor instead.
func (*OutTransaction) Descriptor() ([]byte, []int) {
	return file_sales_proto_rawDescGZIP(), []int{5}
}

func (x *OutTransaction) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *OutTransaction) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_sales_proto protoreflect.FileDescriptor

var file_sales_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2c, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0xb7, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x49, 0x64, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x49, 0x64, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x64, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x64,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x61,
	0x6d, 0x65, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x69, 0x6c,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x05, 0x4e, 0x6f, 0x5f, 0x54, 0x66, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x6f, 0x54, 0x66, 0x22, 0x3d, 0x0a, 0x09, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x0c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x64, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x64, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4e, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x22, 0x5f, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x60, 0x0a, 0x0e, 0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x4f, 0x75, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sales_proto_rawDescOnce sync.Once
	file_sales_proto_rawDescData = file_sales_proto_rawDesc
)

func file_sales_proto_rawDescGZIP() []byte {
	file_sales_proto_rawDescOnce.Do(func() {
		file_sales_proto_rawDescData = protoimpl.X.CompressGZIP(file_sales_proto_rawDescData)
	})
	return file_sales_proto_rawDescData
}

var file_sales_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sales_proto_goTypes = []interface{}{
	(*OrderProduct)(nil),   // 0: model.OrderProduct
	(*Transaction)(nil),    // 1: model.Transaction
	(*OrderList)(nil),      // 2: model.OrderList
	(*InputPayment)(nil),   // 3: model.InputPayment
	(*OutputPayment)(nil),  // 4: model.OutputPayment
	(*OutTransaction)(nil), // 5: model.OutTransaction
	(*empty.Empty)(nil),    // 6: google.protobuf.Empty
}
var file_sales_proto_depIdxs = []int32{
	1, // 0: model.OrderList.OrderList:type_name -> model.Transaction
	1, // 1: model.OutputPayment.Transaction:type_name -> model.Transaction
	1, // 2: model.OutTransaction.Transaction:type_name -> model.Transaction
	0, // 3: model.SalesService.NewOrder:input_type -> model.OrderProduct
	6, // 4: model.SalesService.ListOrder:input_type -> google.protobuf.Empty
	3, // 5: model.SalesService.Payment:input_type -> model.InputPayment
	5, // 6: model.SalesService.NewOrder:output_type -> model.OutTransaction
	2, // 7: model.SalesService.ListOrder:output_type -> model.OrderList
	4, // 8: model.SalesService.Payment:output_type -> model.OutputPayment
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sales_proto_init() }
func file_sales_proto_init() {
	if File_sales_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sales_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProduct); i {
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
		file_sales_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_sales_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderList); i {
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
		file_sales_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputPayment); i {
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
		file_sales_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputPayment); i {
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
		file_sales_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutTransaction); i {
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
			RawDescriptor: file_sales_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sales_proto_goTypes,
		DependencyIndexes: file_sales_proto_depIdxs,
		MessageInfos:      file_sales_proto_msgTypes,
	}.Build()
	File_sales_proto = out.File
	file_sales_proto_rawDesc = nil
	file_sales_proto_goTypes = nil
	file_sales_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SalesServiceClient is the client API for SalesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SalesServiceClient interface {
	NewOrder(ctx context.Context, in *OrderProduct, opts ...grpc.CallOption) (*OutTransaction, error)
	ListOrder(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*OrderList, error)
	Payment(ctx context.Context, in *InputPayment, opts ...grpc.CallOption) (*OutputPayment, error)
}

type salesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSalesServiceClient(cc grpc.ClientConnInterface) SalesServiceClient {
	return &salesServiceClient{cc}
}

func (c *salesServiceClient) NewOrder(ctx context.Context, in *OrderProduct, opts ...grpc.CallOption) (*OutTransaction, error) {
	out := new(OutTransaction)
	err := c.cc.Invoke(ctx, "/model.SalesService/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesServiceClient) ListOrder(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*OrderList, error) {
	out := new(OrderList)
	err := c.cc.Invoke(ctx, "/model.SalesService/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *salesServiceClient) Payment(ctx context.Context, in *InputPayment, opts ...grpc.CallOption) (*OutputPayment, error) {
	out := new(OutputPayment)
	err := c.cc.Invoke(ctx, "/model.SalesService/Payment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SalesServiceServer is the server API for SalesService service.
type SalesServiceServer interface {
	NewOrder(context.Context, *OrderProduct) (*OutTransaction, error)
	ListOrder(context.Context, *empty.Empty) (*OrderList, error)
	Payment(context.Context, *InputPayment) (*OutputPayment, error)
}

// UnimplementedSalesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSalesServiceServer struct {
}

func (*UnimplementedSalesServiceServer) NewOrder(context.Context, *OrderProduct) (*OutTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (*UnimplementedSalesServiceServer) ListOrder(context.Context, *empty.Empty) (*OrderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (*UnimplementedSalesServiceServer) Payment(context.Context, *InputPayment) (*OutputPayment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payment not implemented")
}

func RegisterSalesServiceServer(s *grpc.Server, srv SalesServiceServer) {
	s.RegisterService(&_SalesService_serviceDesc, srv)
}

func _SalesService_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SalesService/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).NewOrder(ctx, req.(*OrderProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesService_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SalesService/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).ListOrder(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SalesService_Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputPayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SalesServiceServer).Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.SalesService/Payment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SalesServiceServer).Payment(ctx, req.(*InputPayment))
	}
	return interceptor(ctx, in, info, handler)
}

var _SalesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.SalesService",
	HandlerType: (*SalesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewOrder",
			Handler:    _SalesService_NewOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _SalesService_ListOrder_Handler,
		},
		{
			MethodName: "Payment",
			Handler:    _SalesService_Payment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sales.proto",
}

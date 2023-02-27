// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: api.proto

package grpc_api

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type SuccessorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *SuccessorResponse) Reset() {
	*x = SuccessorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessorResponse) ProtoMessage() {}

func (x *SuccessorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessorResponse.ProtoReflect.Descriptor instead.
func (*SuccessorResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *SuccessorResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SuccessorResponse) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type PredecessorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *PredecessorResponse) Reset() {
	*x = PredecessorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredecessorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredecessorResponse) ProtoMessage() {}

func (x *PredecessorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredecessorResponse.ProtoReflect.Descriptor instead.
func (*PredecessorResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *PredecessorResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PredecessorResponse) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type HandleNewPredecessorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *HandleNewPredecessorRequest) Reset() {
	*x = HandleNewPredecessorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleNewPredecessorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleNewPredecessorRequest) ProtoMessage() {}

func (x *HandleNewPredecessorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleNewPredecessorRequest.ProtoReflect.Descriptor instead.
func (*HandleNewPredecessorRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *HandleNewPredecessorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HandleNewPredecessorRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

type HandleNewPredecessorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *HandleNewPredecessorResponse) Reset() {
	*x = HandleNewPredecessorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleNewPredecessorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleNewPredecessorResponse) ProtoMessage() {}

func (x *HandleNewPredecessorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleNewPredecessorResponse.ProtoReflect.Descriptor instead.
func (*HandleNewPredecessorResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *HandleNewPredecessorResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type HandleNewSuccessorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Endpoint      string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	NSuccId       int64  `protobuf:"varint,3,opt,name=nSuccId,proto3" json:"nSuccId,omitempty"`
	NSuccEndpoint string `protobuf:"bytes,4,opt,name=nSuccEndpoint,proto3" json:"nSuccEndpoint,omitempty"`
}

func (x *HandleNewSuccessorRequest) Reset() {
	*x = HandleNewSuccessorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleNewSuccessorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleNewSuccessorRequest) ProtoMessage() {}

func (x *HandleNewSuccessorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleNewSuccessorRequest.ProtoReflect.Descriptor instead.
func (*HandleNewSuccessorRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *HandleNewSuccessorRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HandleNewSuccessorRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *HandleNewSuccessorRequest) GetNSuccId() int64 {
	if x != nil {
		return x.NSuccId
	}
	return 0
}

func (x *HandleNewSuccessorRequest) GetNSuccEndpoint() string {
	if x != nil {
		return x.NSuccEndpoint
	}
	return ""
}

type HandleNewSuccessorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *HandleNewSuccessorResponse) Reset() {
	*x = HandleNewSuccessorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandleNewSuccessorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandleNewSuccessorResponse) ProtoMessage() {}

func (x *HandleNewSuccessorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandleNewSuccessorResponse.ProtoReflect.Descriptor instead.
func (*HandleNewSuccessorResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *HandleNewSuccessorResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type QueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	StrKey string `protobuf:"bytes,2,opt,name=strKey,proto3" json:"strKey,omitempty"`
}

func (x *QueryRequest) Reset() {
	*x = QueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRequest) ProtoMessage() {}

func (x *QueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRequest.ProtoReflect.Descriptor instead.
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *QueryRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *QueryRequest) GetStrKey() string {
	if x != nil {
		return x.StrKey
	}
	return ""
}

type QueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data                    []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	ResponsibleNodeId       int64  `protobuf:"varint,2,opt,name=responsibleNodeId,proto3" json:"responsibleNodeId,omitempty"`
	ResponsibleNodeEndpoint string `protobuf:"bytes,3,opt,name=responsibleNodeEndpoint,proto3" json:"responsibleNodeEndpoint,omitempty"`
}

func (x *QueryResponse) Reset() {
	*x = QueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResponse) ProtoMessage() {}

func (x *QueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResponse.ProtoReflect.Descriptor instead.
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *QueryResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QueryResponse) GetResponsibleNodeId() int64 {
	if x != nil {
		return x.ResponsibleNodeId
	}
	return 0
}

func (x *QueryResponse) GetResponsibleNodeEndpoint() string {
	if x != nil {
		return x.ResponsibleNodeEndpoint
	}
	return ""
}

type RepSaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Value  []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	StrKey string `protobuf:"bytes,3,opt,name=strKey,proto3" json:"strKey,omitempty"`
}

func (x *RepSaveRequest) Reset() {
	*x = RepSaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepSaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepSaveRequest) ProtoMessage() {}

func (x *RepSaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepSaveRequest.ProtoReflect.Descriptor instead.
func (*RepSaveRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *RepSaveRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *RepSaveRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *RepSaveRequest) GetStrKey() string {
	if x != nil {
		return x.StrKey
	}
	return ""
}

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	StrKey string `protobuf:"bytes,3,opt,name=strKey,proto3" json:"strKey,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

func (x *SaveRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *SaveRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SaveRequest) GetStrKey() string {
	if x != nil {
		return x.StrKey
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	StrKey string `protobuf:"bytes,2,opt,name=strKey,proto3" json:"strKey,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *DeleteRequest) GetStrKey() string {
	if x != nil {
		return x.StrKey
	}
	return ""
}

type OwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    int64  `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	StrKey string `protobuf:"bytes,2,opt,name=strKey,proto3" json:"strKey,omitempty"`
}

func (x *OwnerRequest) Reset() {
	*x = OwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerRequest) ProtoMessage() {}

func (x *OwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerRequest.ProtoReflect.Descriptor instead.
func (*OwnerRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{12}
}

func (x *OwnerRequest) GetKey() int64 {
	if x != nil {
		return x.Key
	}
	return 0
}

func (x *OwnerRequest) GetStrKey() string {
	if x != nil {
		return x.StrKey
	}
	return ""
}

type OwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerNodeId       int64  `protobuf:"varint,1,opt,name=ownerNodeId,proto3" json:"ownerNodeId,omitempty"`
	OwnerNodeEndpoint string `protobuf:"bytes,2,opt,name=ownerNodeEndpoint,proto3" json:"ownerNodeEndpoint,omitempty"`
}

func (x *OwnerResponse) Reset() {
	*x = OwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OwnerResponse) ProtoMessage() {}

func (x *OwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OwnerResponse.ProtoReflect.Descriptor instead.
func (*OwnerResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{13}
}

func (x *OwnerResponse) GetOwnerNodeId() int64 {
	if x != nil {
		return x.OwnerNodeId
	}
	return 0
}

func (x *OwnerResponse) GetOwnerNodeEndpoint() string {
	if x != nil {
		return x.OwnerNodeEndpoint
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x70, 0x69, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3f,
	0x0a, 0x11, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22,
	0x41, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x22, 0x49, 0x0a, 0x1b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65, 0x77, 0x50,
	0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x2e, 0x0a,
	0x1c, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x87, 0x01,
	0x0a, 0x19, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x53, 0x75, 0x63, 0x63,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x1a, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x38, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x22,
	0x8b, 0x01, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x11, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x50, 0x0a,
	0x0e, 0x52, 0x65, 0x70, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x22,
	0x4b, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x22, 0x39, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x22, 0x38, 0x0a, 0x0c, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x4b, 0x65,
	0x79, 0x22, 0x5f, 0x0a, 0x0d, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64,
	0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x32, 0x95, 0x06, 0x0a, 0x07, 0x44, 0x48, 0x54, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x64, 0x65,
	0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x14, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x12, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x65, 0x64, 0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x72, 0x65, 0x64,
	0x65, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x61, 0x0a, 0x12, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65, 0x77, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4e, 0x65,
	0x77, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x53,
	0x61, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x70, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x12, 0x42, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a,
	0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6f, 0x6e, 0x69, 0x73, 0x6d,
	0x61, 0x6e, 0x65, 0x6f, 0x74, 0x6f, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x48, 0x54,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_proto_goTypes = []interface{}{
	(*Empty)(nil),                        // 0: grpc_api.Empty
	(*SuccessorResponse)(nil),            // 1: grpc_api.SuccessorResponse
	(*PredecessorResponse)(nil),          // 2: grpc_api.PredecessorResponse
	(*HandleNewPredecessorRequest)(nil),  // 3: grpc_api.HandleNewPredecessorRequest
	(*HandleNewPredecessorResponse)(nil), // 4: grpc_api.HandleNewPredecessorResponse
	(*HandleNewSuccessorRequest)(nil),    // 5: grpc_api.HandleNewSuccessorRequest
	(*HandleNewSuccessorResponse)(nil),   // 6: grpc_api.HandleNewSuccessorResponse
	(*QueryRequest)(nil),                 // 7: grpc_api.QueryRequest
	(*QueryResponse)(nil),                // 8: grpc_api.QueryResponse
	(*RepSaveRequest)(nil),               // 9: grpc_api.RepSaveRequest
	(*SaveRequest)(nil),                  // 10: grpc_api.SaveRequest
	(*DeleteRequest)(nil),                // 11: grpc_api.DeleteRequest
	(*OwnerRequest)(nil),                 // 12: grpc_api.OwnerRequest
	(*OwnerResponse)(nil),                // 13: grpc_api.OwnerResponse
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: grpc_api.DHTNode.Ping:input_type -> grpc_api.Empty
	0,  // 1: grpc_api.DHTNode.Successor:input_type -> grpc_api.Empty
	0,  // 2: grpc_api.DHTNode.Predecessor:input_type -> grpc_api.Empty
	3,  // 3: grpc_api.DHTNode.HandleNewPredecessor:input_type -> grpc_api.HandleNewPredecessorRequest
	5,  // 4: grpc_api.DHTNode.HandleNewSuccessor:input_type -> grpc_api.HandleNewSuccessorRequest
	7,  // 5: grpc_api.DHTNode.Query:input_type -> grpc_api.QueryRequest
	10, // 6: grpc_api.DHTNode.Save:input_type -> grpc_api.SaveRequest
	11, // 7: grpc_api.DHTNode.Delete:input_type -> grpc_api.DeleteRequest
	9,  // 8: grpc_api.DHTNode.RepSave:input_type -> grpc_api.RepSaveRequest
	10, // 9: grpc_api.DHTNode.SaveStream:input_type -> grpc_api.SaveRequest
	7,  // 10: grpc_api.DHTNode.QueryStream:input_type -> grpc_api.QueryRequest
	12, // 11: grpc_api.DHTNode.Owner:input_type -> grpc_api.OwnerRequest
	0,  // 12: grpc_api.DHTNode.Ping:output_type -> grpc_api.Empty
	1,  // 13: grpc_api.DHTNode.Successor:output_type -> grpc_api.SuccessorResponse
	2,  // 14: grpc_api.DHTNode.Predecessor:output_type -> grpc_api.PredecessorResponse
	4,  // 15: grpc_api.DHTNode.HandleNewPredecessor:output_type -> grpc_api.HandleNewPredecessorResponse
	6,  // 16: grpc_api.DHTNode.HandleNewSuccessor:output_type -> grpc_api.HandleNewSuccessorResponse
	8,  // 17: grpc_api.DHTNode.Query:output_type -> grpc_api.QueryResponse
	0,  // 18: grpc_api.DHTNode.Save:output_type -> grpc_api.Empty
	0,  // 19: grpc_api.DHTNode.Delete:output_type -> grpc_api.Empty
	0,  // 20: grpc_api.DHTNode.RepSave:output_type -> grpc_api.Empty
	0,  // 21: grpc_api.DHTNode.SaveStream:output_type -> grpc_api.Empty
	8,  // 22: grpc_api.DHTNode.QueryStream:output_type -> grpc_api.QueryResponse
	13, // 23: grpc_api.DHTNode.Owner:output_type -> grpc_api.OwnerResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessorResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredecessorResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleNewPredecessorRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleNewPredecessorResponse); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleNewSuccessorRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandleNewSuccessorResponse); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRequest); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryResponse); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepSaveRequest); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerRequest); i {
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
		file_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OwnerResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

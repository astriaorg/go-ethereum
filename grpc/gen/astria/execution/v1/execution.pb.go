// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: astria/execution/v1/execution.proto

package executionv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DoBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrevBlockHash []byte                 `protobuf:"bytes,1,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	Transactions  [][]byte               `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *DoBlockRequest) Reset() {
	*x = DoBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoBlockRequest) ProtoMessage() {}

func (x *DoBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoBlockRequest.ProtoReflect.Descriptor instead.
func (*DoBlockRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1_execution_proto_rawDescGZIP(), []int{0}
}

func (x *DoBlockRequest) GetPrevBlockHash() []byte {
	if x != nil {
		return x.PrevBlockHash
	}
	return nil
}

func (x *DoBlockRequest) GetTransactions() [][]byte {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *DoBlockRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type DoBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *DoBlockResponse) Reset() {
	*x = DoBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1_execution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoBlockResponse) ProtoMessage() {}

func (x *DoBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1_execution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoBlockResponse.ProtoReflect.Descriptor instead.
func (*DoBlockResponse) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1_execution_proto_rawDescGZIP(), []int{1}
}

func (x *DoBlockResponse) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

type FinalizeBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *FinalizeBlockRequest) Reset() {
	*x = FinalizeBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1_execution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeBlockRequest) ProtoMessage() {}

func (x *FinalizeBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1_execution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeBlockRequest.ProtoReflect.Descriptor instead.
func (*FinalizeBlockRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1_execution_proto_rawDescGZIP(), []int{2}
}

func (x *FinalizeBlockRequest) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

type FinalizeBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinalizeBlockResponse) Reset() {
	*x = FinalizeBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1_execution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalizeBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalizeBlockResponse) ProtoMessage() {}

func (x *FinalizeBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1_execution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalizeBlockResponse.ProtoReflect.Descriptor instead.
func (*FinalizeBlockResponse) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1_execution_proto_rawDescGZIP(), []int{3}
}

type InitStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitStateRequest) Reset() {
	*x = InitStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1_execution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitStateRequest) ProtoMessage() {}

func (x *InitStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1_execution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitStateRequest.ProtoReflect.Descriptor instead.
func (*InitStateRequest) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1_execution_proto_rawDescGZIP(), []int{4}
}

type InitStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *InitStateResponse) Reset() {
	*x = InitStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_astria_execution_v1_execution_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitStateResponse) ProtoMessage() {}

func (x *InitStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_astria_execution_v1_execution_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitStateResponse.ProtoReflect.Descriptor instead.
func (*InitStateResponse) Descriptor() ([]byte, []int) {
	return file_astria_execution_v1_execution_proto_rawDescGZIP(), []int{5}
}

func (x *InitStateResponse) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

var File_astria_execution_v1_execution_proto protoreflect.FileDescriptor

var file_astria_execution_v1_execution_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x0e,
	0x44, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x30, 0x0a, 0x0f, 0x44, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x35, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x17, 0x0a,
	0x15, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x11, 0x49, 0x6e,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x32, 0xac,
	0x02, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61,
	0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x07, 0x44, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x61, 0x73, 0x74,
	0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x29, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xdd, 0x01,
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x73, 0x74, 0x72, 0x69, 0x61, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x13, 0x41, 0x73, 0x74, 0x72, 0x69, 0x61,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13,
	0x41, 0x73, 0x74, 0x72, 0x69, 0x61, 0x5c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x41, 0x73, 0x74, 0x72, 0x69, 0x61, 0x5c, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x41, 0x73, 0x74, 0x72, 0x69, 0x61, 0x3a, 0x3a,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_astria_execution_v1_execution_proto_rawDescOnce sync.Once
	file_astria_execution_v1_execution_proto_rawDescData = file_astria_execution_v1_execution_proto_rawDesc
)

func file_astria_execution_v1_execution_proto_rawDescGZIP() []byte {
	file_astria_execution_v1_execution_proto_rawDescOnce.Do(func() {
		file_astria_execution_v1_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_astria_execution_v1_execution_proto_rawDescData)
	})
	return file_astria_execution_v1_execution_proto_rawDescData
}

var file_astria_execution_v1_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_astria_execution_v1_execution_proto_goTypes = []interface{}{
	(*DoBlockRequest)(nil),        // 0: astria.execution.v1.DoBlockRequest
	(*DoBlockResponse)(nil),       // 1: astria.execution.v1.DoBlockResponse
	(*FinalizeBlockRequest)(nil),  // 2: astria.execution.v1.FinalizeBlockRequest
	(*FinalizeBlockResponse)(nil), // 3: astria.execution.v1.FinalizeBlockResponse
	(*InitStateRequest)(nil),      // 4: astria.execution.v1.InitStateRequest
	(*InitStateResponse)(nil),     // 5: astria.execution.v1.InitStateResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_astria_execution_v1_execution_proto_depIdxs = []int32{
	6, // 0: astria.execution.v1.DoBlockRequest.timestamp:type_name -> google.protobuf.Timestamp
	4, // 1: astria.execution.v1.ExecutionService.InitState:input_type -> astria.execution.v1.InitStateRequest
	0, // 2: astria.execution.v1.ExecutionService.DoBlock:input_type -> astria.execution.v1.DoBlockRequest
	2, // 3: astria.execution.v1.ExecutionService.FinalizeBlock:input_type -> astria.execution.v1.FinalizeBlockRequest
	5, // 4: astria.execution.v1.ExecutionService.InitState:output_type -> astria.execution.v1.InitStateResponse
	1, // 5: astria.execution.v1.ExecutionService.DoBlock:output_type -> astria.execution.v1.DoBlockResponse
	3, // 6: astria.execution.v1.ExecutionService.FinalizeBlock:output_type -> astria.execution.v1.FinalizeBlockResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_astria_execution_v1_execution_proto_init() }
func file_astria_execution_v1_execution_proto_init() {
	if File_astria_execution_v1_execution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_astria_execution_v1_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoBlockRequest); i {
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
		file_astria_execution_v1_execution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoBlockResponse); i {
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
		file_astria_execution_v1_execution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeBlockRequest); i {
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
		file_astria_execution_v1_execution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalizeBlockResponse); i {
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
		file_astria_execution_v1_execution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitStateRequest); i {
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
		file_astria_execution_v1_execution_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitStateResponse); i {
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
			RawDescriptor: file_astria_execution_v1_execution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_astria_execution_v1_execution_proto_goTypes,
		DependencyIndexes: file_astria_execution_v1_execution_proto_depIdxs,
		MessageInfos:      file_astria_execution_v1_execution_proto_msgTypes,
	}.Build()
	File_astria_execution_v1_execution_proto = out.File
	file_astria_execution_v1_execution_proto_rawDesc = nil
	file_astria_execution_v1_execution_proto_goTypes = nil
	file_astria_execution_v1_execution_proto_depIdxs = nil
}

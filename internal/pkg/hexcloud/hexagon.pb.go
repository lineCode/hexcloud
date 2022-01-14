// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/hexagon.proto

package hexcloud

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
		mi := &file_api_hexagon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[0]
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
	return file_api_hexagon_proto_rawDescGZIP(), []int{0}
}

type Hex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X         int64  `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"` //
	Y         int64  `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z         int64  `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
	Type      string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Direction string `protobuf:"bytes,5,opt,name=Direction,proto3" json:"Direction,omitempty"`
}

func (x *Hex) Reset() {
	*x = Hex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hex) ProtoMessage() {}

func (x *Hex) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hex.ProtoReflect.Descriptor instead.
func (*Hex) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{1}
}

func (x *Hex) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Hex) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Hex) GetZ() int64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *Hex) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Hex) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type HexAxial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	U         int64  `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	V         int64  `protobuf:"varint,2,opt,name=V,proto3" json:"V,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Direction string `protobuf:"bytes,4,opt,name=Direction,proto3" json:"Direction,omitempty"`
}

func (x *HexAxial) Reset() {
	*x = HexAxial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexAxial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexAxial) ProtoMessage() {}

func (x *HexAxial) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexAxial.ProtoReflect.Descriptor instead.
func (*HexAxial) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{2}
}

func (x *HexAxial) GetU() int64 {
	if x != nil {
		return x.U
	}
	return 0
}

func (x *HexAxial) GetV() int64 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *HexAxial) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *HexAxial) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

type ConversionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hc []*Hex `protobuf:"bytes,1,rep,name=hc,proto3" json:"hc,omitempty"`
}

func (x *ConversionRequest) Reset() {
	*x = ConversionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConversionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionRequest) ProtoMessage() {}

func (x *ConversionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionRequest.ProtoReflect.Descriptor instead.
func (*ConversionRequest) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{3}
}

func (x *ConversionRequest) GetHc() []*Hex {
	if x != nil {
		return x.Hc
	}
	return nil
}

type HexagonRingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ha     *Hex  `protobuf:"bytes,1,opt,name=ha,proto3" json:"ha,omitempty"`
	Radius int64 `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
	Fill   bool  `protobuf:"varint,3,opt,name=fill,proto3" json:"fill,omitempty"`
}

func (x *HexagonRingRequest) Reset() {
	*x = HexagonRingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexagonRingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexagonRingRequest) ProtoMessage() {}

func (x *HexagonRingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexagonRingRequest.ProtoReflect.Descriptor instead.
func (*HexagonRingRequest) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{4}
}

func (x *HexagonRingRequest) GetHa() *Hex {
	if x != nil {
		return x.Ha
	}
	return nil
}

func (x *HexagonRingRequest) GetRadius() int64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *HexagonRingRequest) GetFill() bool {
	if x != nil {
		return x.Fill
	}
	return false
}

type HexAxialResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ha []*HexAxial `protobuf:"bytes,1,rep,name=ha,proto3" json:"ha,omitempty"`
}

func (x *HexAxialResponse) Reset() {
	*x = HexAxialResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexAxialResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexAxialResponse) ProtoMessage() {}

func (x *HexAxialResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexAxialResponse.ProtoReflect.Descriptor instead.
func (*HexAxialResponse) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{5}
}

func (x *HexAxialResponse) GetHa() []*HexAxial {
	if x != nil {
		return x.Ha
	}
	return nil
}

type HexCubeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hc []*Hex `protobuf:"bytes,1,rep,name=hc,proto3" json:"hc,omitempty"`
}

func (x *HexCubeResponse) Reset() {
	*x = HexCubeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexCubeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexCubeResponse) ProtoMessage() {}

func (x *HexCubeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexCubeResponse.ProtoReflect.Descriptor instead.
func (*HexCubeResponse) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{6}
}

func (x *HexCubeResponse) GetHc() []*Hex {
	if x != nil {
		return x.Hc
	}
	return nil
}

type HexAxialList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ha []*HexAxial `protobuf:"bytes,1,rep,name=ha,proto3" json:"ha,omitempty"`
}

func (x *HexAxialList) Reset() {
	*x = HexAxialList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexAxialList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexAxialList) ProtoMessage() {}

func (x *HexAxialList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexAxialList.ProtoReflect.Descriptor instead.
func (*HexAxialList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{7}
}

func (x *HexAxialList) GetHa() []*HexAxial {
	if x != nil {
		return x.Ha
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{8}
}

func (x *Status) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_hexagon_proto protoreflect.FileDescriptor

var file_api_hexagon_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68,
	0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x61, 0x0a, 0x03, 0x48, 0x65, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12, 0x0c,
	0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01,
	0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x5a, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x58, 0x0a, 0x08,
	0x48, 0x65, 0x78, 0x41, 0x78, 0x69, 0x61, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x55, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x55, 0x12, 0x0c, 0x0a, 0x01, 0x56, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x56, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x02, 0x68,
	0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x52, 0x02, 0x68, 0x63, 0x22, 0x72, 0x0a,
	0x12, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x52, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x02, 0x68, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65,
	0x78, 0x52, 0x02, 0x68, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x69, 0x6c,
	0x6c, 0x22, 0x49, 0x0a, 0x10, 0x48, 0x65, 0x78, 0x41, 0x78, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x02, 0x68, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65,
	0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x48, 0x65, 0x78, 0x41, 0x78, 0x69, 0x61, 0x6c, 0x52, 0x02, 0x68, 0x61, 0x22, 0x43, 0x0a, 0x0f,
	0x48, 0x65, 0x78, 0x43, 0x75, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x02, 0x68, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x52, 0x02, 0x68,
	0x63, 0x22, 0x45, 0x0a, 0x0c, 0x48, 0x65, 0x78, 0x41, 0x78, 0x69, 0x61, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x02, 0x68, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x41,
	0x78, 0x69, 0x61, 0x6c, 0x52, 0x02, 0x68, 0x61, 0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0xb7, 0x02, 0x0a, 0x0e, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x78, 0x61, 0x67, 0x6f, 0x6e, 0x52, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x52,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x43, 0x75, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x41, 0x78, 0x69, 0x61, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b,
	0x5a, 0x09, 0x2f, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_hexagon_proto_rawDescOnce sync.Once
	file_api_hexagon_proto_rawDescData = file_api_hexagon_proto_rawDesc
)

func file_api_hexagon_proto_rawDescGZIP() []byte {
	file_api_hexagon_proto_rawDescOnce.Do(func() {
		file_api_hexagon_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hexagon_proto_rawDescData)
	})
	return file_api_hexagon_proto_rawDescData
}

var file_api_hexagon_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_hexagon_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: endpoints.hexworld.hexcloud.Empty
	(*Hex)(nil),                // 1: endpoints.hexworld.hexcloud.Hex
	(*HexAxial)(nil),           // 2: endpoints.hexworld.hexcloud.HexAxial
	(*ConversionRequest)(nil),  // 3: endpoints.hexworld.hexcloud.ConversionRequest
	(*HexagonRingRequest)(nil), // 4: endpoints.hexworld.hexcloud.HexagonRingRequest
	(*HexAxialResponse)(nil),   // 5: endpoints.hexworld.hexcloud.HexAxialResponse
	(*HexCubeResponse)(nil),    // 6: endpoints.hexworld.hexcloud.HexCubeResponse
	(*HexAxialList)(nil),       // 7: endpoints.hexworld.hexcloud.HexAxialList
	(*Status)(nil),             // 8: endpoints.hexworld.hexcloud.Status
}
var file_api_hexagon_proto_depIdxs = []int32{
	1, // 0: endpoints.hexworld.hexcloud.ConversionRequest.hc:type_name -> endpoints.hexworld.hexcloud.Hex
	1, // 1: endpoints.hexworld.hexcloud.HexagonRingRequest.ha:type_name -> endpoints.hexworld.hexcloud.Hex
	2, // 2: endpoints.hexworld.hexcloud.HexAxialResponse.ha:type_name -> endpoints.hexworld.hexcloud.HexAxial
	1, // 3: endpoints.hexworld.hexcloud.HexCubeResponse.hc:type_name -> endpoints.hexworld.hexcloud.Hex
	2, // 4: endpoints.hexworld.hexcloud.HexAxialList.ha:type_name -> endpoints.hexworld.hexcloud.HexAxial
	4, // 5: endpoints.hexworld.hexcloud.HexagonService.GetHexagonRing:input_type -> endpoints.hexworld.hexcloud.HexagonRingRequest
	7, // 6: endpoints.hexworld.hexcloud.HexagonService.StoreHexagons:input_type -> endpoints.hexworld.hexcloud.HexAxialList
	0, // 7: endpoints.hexworld.hexcloud.HexagonService.GetStatus:input_type -> endpoints.hexworld.hexcloud.Empty
	6, // 8: endpoints.hexworld.hexcloud.HexagonService.GetHexagonRing:output_type -> endpoints.hexworld.hexcloud.HexCubeResponse
	0, // 9: endpoints.hexworld.hexcloud.HexagonService.StoreHexagons:output_type -> endpoints.hexworld.hexcloud.Empty
	8, // 10: endpoints.hexworld.hexcloud.HexagonService.GetStatus:output_type -> endpoints.hexworld.hexcloud.Status
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_hexagon_proto_init() }
func file_api_hexagon_proto_init() {
	if File_api_hexagon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hexagon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_hexagon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hex); i {
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
		file_api_hexagon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexAxial); i {
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
		file_api_hexagon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConversionRequest); i {
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
		file_api_hexagon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexagonRingRequest); i {
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
		file_api_hexagon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexAxialResponse); i {
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
		file_api_hexagon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexCubeResponse); i {
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
		file_api_hexagon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexAxialList); i {
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
		file_api_hexagon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_api_hexagon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hexagon_proto_goTypes,
		DependencyIndexes: file_api_hexagon_proto_depIdxs,
		MessageInfos:      file_api_hexagon_proto_msgTypes,
	}.Build()
	File_api_hexagon_proto = out.File
	file_api_hexagon_proto_rawDesc = nil
	file_api_hexagon_proto_goTypes = nil
	file_api_hexagon_proto_depIdxs = nil
}

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

type Direction int32

const (
	Direction_N  Direction = 0
	Direction_NE Direction = 1
	Direction_E  Direction = 2
	Direction_SE Direction = 3
	Direction_S  Direction = 4
	Direction_SW Direction = 5
	Direction_W  Direction = 6
	Direction_NW Direction = 7
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "N",
		1: "NE",
		2: "E",
		3: "SE",
		4: "S",
		5: "SW",
		6: "W",
		7: "NW",
	}
	Direction_value = map[string]int32{
		"N":  0,
		"NE": 1,
		"E":  2,
		"SE": 3,
		"S":  4,
		"SW": 5,
		"W":  6,
		"NW": 7,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_api_hexagon_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_api_hexagon_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{0}
}

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

	X         int64     `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y         int64     `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z         int64     `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
	Direction Direction `protobuf:"varint,4,opt,name=Direction,proto3,enum=endpoints.hexworld.hexcloud.Direction" json:"Direction,omitempty"`
	Reference string    `protobuf:"bytes,5,opt,name=Reference,proto3" json:"Reference,omitempty"`
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

func (x *Hex) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_N
}

func (x *Hex) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

type HexReference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref string `protobuf:"bytes,1,opt,name=Ref,proto3" json:"Ref,omitempty"`
}

func (x *HexReference) Reset() {
	*x = HexReference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexReference) ProtoMessage() {}

func (x *HexReference) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HexReference.ProtoReflect.Descriptor instead.
func (*HexReference) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{2}
}

func (x *HexReference) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

type HexRefList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref []*HexReference `protobuf:"bytes,1,rep,name=ref,proto3" json:"ref,omitempty"`
}

func (x *HexRefList) Reset() {
	*x = HexRefList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexRefList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexRefList) ProtoMessage() {}

func (x *HexRefList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HexRefList.ProtoReflect.Descriptor instead.
func (*HexRefList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{3}
}

func (x *HexRefList) GetRef() []*HexReference {
	if x != nil {
		return x.Ref
	}
	return nil
}

type HexList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hex []*Hex `protobuf:"bytes,1,rep,name=hex,proto3" json:"hex,omitempty"`
}

func (x *HexList) Reset() {
	*x = HexList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexList) ProtoMessage() {}

func (x *HexList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HexList.ProtoReflect.Descriptor instead.
func (*HexList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{4}
}

func (x *HexList) GetHex() []*Hex {
	if x != nil {
		return x.Hex
	}
	return nil
}

type HexagonGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hex    *Hex  `protobuf:"bytes,1,opt,name=hex,proto3" json:"hex,omitempty"`
	Radius int64 `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
	Fill   bool  `protobuf:"varint,3,opt,name=fill,proto3" json:"fill,omitempty"`
}

func (x *HexagonGetRequest) Reset() {
	*x = HexagonGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexagonGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexagonGetRequest) ProtoMessage() {}

func (x *HexagonGetRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use HexagonGetRequest.ProtoReflect.Descriptor instead.
func (*HexagonGetRequest) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{5}
}

func (x *HexagonGetRequest) GetHex() *Hex {
	if x != nil {
		return x.Hex
	}
	return nil
}

func (x *HexagonGetRequest) GetRadius() int64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *HexagonGetRequest) GetFill() bool {
	if x != nil {
		return x.Fill
	}
	return false
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
		mi := &file_api_hexagon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{7}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_hexagon_proto protoreflect.FileDescriptor

var file_api_hexagon_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68,
	0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x93, 0x01, 0x0a, 0x03, 0x48, 0x65,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x58, 0x12,
	0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x59, 0x12, 0x0c, 0x0a,
	0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x5a, 0x12, 0x44, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22,
	0x20, 0x0a, 0x0c, 0x48, 0x65, 0x78, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x52, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x52, 0x65,
	0x66, 0x22, 0x49, 0x0a, 0x0a, 0x48, 0x65, 0x78, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x3b, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x3d, 0x0a, 0x07,
	0x48, 0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x03, 0x68, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x52, 0x03, 0x68, 0x65, 0x78, 0x22, 0x73, 0x0a, 0x11, 0x48,
	0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x32, 0x0a, 0x03, 0x68, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x52,
	0x03, 0x68, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x6c,
	0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x2a, 0x47, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x0a,
	0x01, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01,
	0x45, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x45, 0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x53,
	0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x57, 0x10, 0x05, 0x12, 0x05, 0x0a, 0x01, 0x57, 0x10,
	0x06, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x57, 0x10, 0x07, 0x32, 0xc4, 0x06, 0x0a, 0x0e, 0x48, 0x65,
	0x78, 0x61, 0x67, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0f,
	0x52, 0x65, 0x70, 0x6f, 0x41, 0x64, 0x64, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x73, 0x12,
	0x27, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65,
	0x78, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65,
	0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x5f, 0x0a,
	0x0f, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65, 0x6c, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x73,
	0x12, 0x27, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48,
	0x65, 0x78, 0x52, 0x65, 0x66, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4f,
	0x0a, 0x06, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64, 0x12, 0x20, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65,
	0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x1a, 0x23, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x5e, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x56, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x24, 0x2e, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x69,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68,
	0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x51, 0x0a, 0x0b, 0x48, 0x65, 0x78, 0x61, 0x67,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x1a, 0x20, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65,
	0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x12, 0x5a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x23, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65,
	0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23,
	0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x5b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x68, 0x65, 0x78,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x68, 0x65, 0x78, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_hexagon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_hexagon_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_hexagon_proto_goTypes = []interface{}{
	(Direction)(0),            // 0: endpoints.hexworld.hexcloud.Direction
	(*Empty)(nil),             // 1: endpoints.hexworld.hexcloud.Empty
	(*Hex)(nil),               // 2: endpoints.hexworld.hexcloud.Hex
	(*HexReference)(nil),      // 3: endpoints.hexworld.hexcloud.HexReference
	(*HexRefList)(nil),        // 4: endpoints.hexworld.hexcloud.HexRefList
	(*HexList)(nil),           // 5: endpoints.hexworld.hexcloud.HexList
	(*HexagonGetRequest)(nil), // 6: endpoints.hexworld.hexcloud.HexagonGetRequest
	(*Status)(nil),            // 7: endpoints.hexworld.hexcloud.Status
	(*Result)(nil),            // 8: endpoints.hexworld.hexcloud.Result
}
var file_api_hexagon_proto_depIdxs = []int32{
	0,  // 0: endpoints.hexworld.hexcloud.Hex.Direction:type_name -> endpoints.hexworld.hexcloud.Direction
	3,  // 1: endpoints.hexworld.hexcloud.HexRefList.ref:type_name -> endpoints.hexworld.hexcloud.HexReference
	2,  // 2: endpoints.hexworld.hexcloud.HexList.hex:type_name -> endpoints.hexworld.hexcloud.Hex
	2,  // 3: endpoints.hexworld.hexcloud.HexagonGetRequest.hex:type_name -> endpoints.hexworld.hexcloud.Hex
	4,  // 4: endpoints.hexworld.hexcloud.HexagonService.RepoAddHexagons:input_type -> endpoints.hexworld.hexcloud.HexRefList
	4,  // 5: endpoints.hexworld.hexcloud.HexagonService.RepoDelHexagons:input_type -> endpoints.hexworld.hexcloud.HexRefList
	2,  // 6: endpoints.hexworld.hexcloud.HexagonService.MapAdd:input_type -> endpoints.hexworld.hexcloud.Hex
	6,  // 7: endpoints.hexworld.hexcloud.HexagonService.MapGet:input_type -> endpoints.hexworld.hexcloud.HexagonGetRequest
	5,  // 8: endpoints.hexworld.hexcloud.HexagonService.MapRemove:input_type -> endpoints.hexworld.hexcloud.HexList
	2,  // 9: endpoints.hexworld.hexcloud.HexagonService.HexagonInfo:input_type -> endpoints.hexworld.hexcloud.Hex
	1,  // 10: endpoints.hexworld.hexcloud.HexagonService.GetStatusServer:input_type -> endpoints.hexworld.hexcloud.Empty
	1,  // 11: endpoints.hexworld.hexcloud.HexagonService.GetStatusStorage:input_type -> endpoints.hexworld.hexcloud.Empty
	1,  // 12: endpoints.hexworld.hexcloud.HexagonService.GetStatusClients:input_type -> endpoints.hexworld.hexcloud.Empty
	8,  // 13: endpoints.hexworld.hexcloud.HexagonService.RepoAddHexagons:output_type -> endpoints.hexworld.hexcloud.Result
	8,  // 14: endpoints.hexworld.hexcloud.HexagonService.RepoDelHexagons:output_type -> endpoints.hexworld.hexcloud.Result
	8,  // 15: endpoints.hexworld.hexcloud.HexagonService.MapAdd:output_type -> endpoints.hexworld.hexcloud.Result
	5,  // 16: endpoints.hexworld.hexcloud.HexagonService.MapGet:output_type -> endpoints.hexworld.hexcloud.HexList
	8,  // 17: endpoints.hexworld.hexcloud.HexagonService.MapRemove:output_type -> endpoints.hexworld.hexcloud.Result
	2,  // 18: endpoints.hexworld.hexcloud.HexagonService.HexagonInfo:output_type -> endpoints.hexworld.hexcloud.Hex
	7,  // 19: endpoints.hexworld.hexcloud.HexagonService.GetStatusServer:output_type -> endpoints.hexworld.hexcloud.Status
	7,  // 20: endpoints.hexworld.hexcloud.HexagonService.GetStatusStorage:output_type -> endpoints.hexworld.hexcloud.Status
	7,  // 21: endpoints.hexworld.hexcloud.HexagonService.GetStatusClients:output_type -> endpoints.hexworld.hexcloud.Status
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
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
			switch v := v.(*HexReference); i {
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
			switch v := v.(*HexRefList); i {
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
			switch v := v.(*HexList); i {
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
			switch v := v.(*HexagonGetRequest); i {
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
		file_api_hexagon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hexagon_proto_goTypes,
		DependencyIndexes: file_api_hexagon_proto_depIdxs,
		EnumInfos:         file_api_hexagon_proto_enumTypes,
		MessageInfos:      file_api_hexagon_proto_msgTypes,
	}.Build()
	File_api_hexagon_proto = out.File
	file_api_hexagon_proto_rawDesc = nil
	file_api_hexagon_proto_goTypes = nil
	file_api_hexagon_proto_depIdxs = nil
}

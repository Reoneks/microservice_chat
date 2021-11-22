// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: rooms.proto

package proto

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

type RoomID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RoomID) Reset() {
	*x = RoomID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomID) ProtoMessage() {}

func (x *RoomID) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomID.ProtoReflect.Descriptor instead.
func (*RoomID) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{0}
}

func (x *RoomID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type RoomStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomInfo []byte `protobuf:"bytes,1,opt,name=RoomInfo,proto3" json:"RoomInfo,omitempty"`
}

func (x *RoomStruct) Reset() {
	*x = RoomStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStruct) ProtoMessage() {}

func (x *RoomStruct) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStruct.ProtoReflect.Descriptor instead.
func (*RoomStruct) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{1}
}

func (x *RoomStruct) GetRoomInfo() []byte {
	if x != nil {
		return x.RoomInfo
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomIDs []string `protobuf:"bytes,1,rep,name=RoomIDs,proto3" json:"RoomIDs,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetRoomIDs() []string {
	if x != nil {
		return x.RoomIDs
	}
	return nil
}

type GetRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rooms  []*RoomStruct `protobuf:"bytes,1,rep,name=Rooms,proto3" json:"Rooms,omitempty"`
	Status *Status       `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetRoomsResponse) Reset() {
	*x = GetRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomsResponse) ProtoMessage() {}

func (x *GetRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomsResponse.ProtoReflect.Descriptor instead.
func (*GetRoomsResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoomsResponse) GetRooms() []*RoomStruct {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *GetRoomsResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type RoomStructResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room   *RoomStruct `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`
	Status *Status     `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *RoomStructResponse) Reset() {
	*x = RoomStructResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomStructResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomStructResponse) ProtoMessage() {}

func (x *RoomStructResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomStructResponse.ProtoReflect.Descriptor instead.
func (*RoomStructResponse) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{4}
}

func (x *RoomStructResponse) GetRoom() *RoomStruct {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *RoomStructResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID string `protobuf:"bytes,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
	UserID string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[5]
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
	return file_rooms_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *DeleteRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok    bool   `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[6]
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
	return file_rooms_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Status) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room   *RoomStruct `protobuf:"bytes,1,opt,name=Room,proto3" json:"Room,omitempty"`
	UserID string      `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateRequest) GetRoom() *RoomStruct {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *UpdateRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type AddUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomID  string   `protobuf:"bytes,1,opt,name=RoomID,proto3" json:"RoomID,omitempty"`
	UserID  string   `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserIDs []string `protobuf:"bytes,3,rep,name=UserIDs,proto3" json:"UserIDs,omitempty"`
}

func (x *AddUsersRequest) Reset() {
	*x = AddUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rooms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUsersRequest) ProtoMessage() {}

func (x *AddUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rooms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUsersRequest.ProtoReflect.Descriptor instead.
func (*AddUsersRequest) Descriptor() ([]byte, []int) {
	return file_rooms_proto_rawDescGZIP(), []int{8}
}

func (x *AddUsersRequest) GetRoomID() string {
	if x != nil {
		return x.RoomID
	}
	return ""
}

func (x *AddUsersRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AddUsersRequest) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

var File_rooms_proto protoreflect.FileDescriptor

var file_rooms_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72,
	0x6f, 0x6f, 0x6d, 0x73, 0x22, 0x18, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x28,
	0x0a, 0x0a, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x22, 0x62, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x62, 0x0a, 0x12, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x25, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4e, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x5b, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x32, 0xdd, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x35, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x1a, 0x19, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x12,
	0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x17,
	0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x11, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x14, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x14, 0x2e, 0x72, 0x6f, 0x6f,
	0x6d, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rooms_proto_rawDescOnce sync.Once
	file_rooms_proto_rawDescData = file_rooms_proto_rawDesc
)

func file_rooms_proto_rawDescGZIP() []byte {
	file_rooms_proto_rawDescOnce.Do(func() {
		file_rooms_proto_rawDescData = protoimpl.X.CompressGZIP(file_rooms_proto_rawDescData)
	})
	return file_rooms_proto_rawDescData
}

var file_rooms_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_rooms_proto_goTypes = []interface{}{
	(*RoomID)(nil),             // 0: rooms.RoomID
	(*RoomStruct)(nil),         // 1: rooms.RoomStruct
	(*Filter)(nil),             // 2: rooms.Filter
	(*GetRoomsResponse)(nil),   // 3: rooms.GetRoomsResponse
	(*RoomStructResponse)(nil), // 4: rooms.RoomStructResponse
	(*DeleteRequest)(nil),      // 5: rooms.DeleteRequest
	(*Status)(nil),             // 6: rooms.Status
	(*UpdateRequest)(nil),      // 7: rooms.UpdateRequest
	(*AddUsersRequest)(nil),    // 8: rooms.AddUsersRequest
}
var file_rooms_proto_depIdxs = []int32{
	1,  // 0: rooms.GetRoomsResponse.Rooms:type_name -> rooms.RoomStruct
	6,  // 1: rooms.GetRoomsResponse.Status:type_name -> rooms.Status
	1,  // 2: rooms.RoomStructResponse.Room:type_name -> rooms.RoomStruct
	6,  // 3: rooms.RoomStructResponse.Status:type_name -> rooms.Status
	1,  // 4: rooms.UpdateRequest.Room:type_name -> rooms.RoomStruct
	0,  // 5: rooms.Rooms.GetRoom:input_type -> rooms.RoomID
	2,  // 6: rooms.Rooms.GetRooms:input_type -> rooms.Filter
	1,  // 7: rooms.Rooms.CreateRoom:input_type -> rooms.RoomStruct
	5,  // 8: rooms.Rooms.DeleteRoom:input_type -> rooms.DeleteRequest
	7,  // 9: rooms.Rooms.UpdateRoom:input_type -> rooms.UpdateRequest
	8,  // 10: rooms.Rooms.AddUsers:input_type -> rooms.AddUsersRequest
	4,  // 11: rooms.Rooms.GetRoom:output_type -> rooms.RoomStructResponse
	3,  // 12: rooms.Rooms.GetRooms:output_type -> rooms.GetRoomsResponse
	4,  // 13: rooms.Rooms.CreateRoom:output_type -> rooms.RoomStructResponse
	6,  // 14: rooms.Rooms.DeleteRoom:output_type -> rooms.Status
	4,  // 15: rooms.Rooms.UpdateRoom:output_type -> rooms.RoomStructResponse
	6,  // 16: rooms.Rooms.AddUsers:output_type -> rooms.Status
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rooms_proto_init() }
func file_rooms_proto_init() {
	if File_rooms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rooms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomID); i {
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
		file_rooms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStruct); i {
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
		file_rooms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_rooms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomsResponse); i {
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
		file_rooms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomStructResponse); i {
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
		file_rooms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_rooms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_rooms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_rooms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUsersRequest); i {
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
			RawDescriptor: file_rooms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rooms_proto_goTypes,
		DependencyIndexes: file_rooms_proto_depIdxs,
		MessageInfos:      file_rooms_proto_msgTypes,
	}.Build()
	File_rooms_proto = out.File
	file_rooms_proto_rawDesc = nil
	file_rooms_proto_goTypes = nil
	file_rooms_proto_depIdxs = nil
}
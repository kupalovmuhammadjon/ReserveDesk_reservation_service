// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: reservation.proto

package reservations

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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MenuRespons struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ItemType     string `protobuf:"bytes,2,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price        int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Description  string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	RestaurantId string `protobuf:"bytes,6,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	CreatedAt    string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *MenuRespons) Reset() {
	*x = MenuRespons{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuRespons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRespons) ProtoMessage() {}

func (x *MenuRespons) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRespons.ProtoReflect.Descriptor instead.
func (*MenuRespons) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *MenuRespons) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MenuRespons) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *MenuRespons) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuRespons) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MenuRespons) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MenuRespons) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *MenuRespons) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MenuRespons) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Exists struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=Exists,proto3" json:"Exists,omitempty"`
}

func (x *Exists) Reset() {
	*x = Exists{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exists) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exists) ProtoMessage() {}

func (x *Exists) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exists.ProtoReflect.Descriptor instead.
func (*Exists) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *Exists) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{3}
}

type ReservationUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestaurantId  string `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	ArrivingTime  string `protobuf:"bytes,3,opt,name=arriving_time,json=arrivingTime,proto3" json:"arriving_time,omitempty"`
	UserId        string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NumberOfSeats int32  `protobuf:"varint,5,opt,name=number_of_seats,json=numberOfSeats,proto3" json:"number_of_seats,omitempty"`
}

func (x *ReservationUpdate) Reset() {
	*x = ReservationUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationUpdate) ProtoMessage() {}

func (x *ReservationUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationUpdate.ProtoReflect.Descriptor instead.
func (*ReservationUpdate) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{4}
}

func (x *ReservationUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationUpdate) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *ReservationUpdate) GetArrivingTime() string {
	if x != nil {
		return x.ArrivingTime
	}
	return ""
}

func (x *ReservationUpdate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationUpdate) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

type ReservationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RestaurantId  string `protobuf:"bytes,2,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	ArrivingTime  string `protobuf:"bytes,3,opt,name=arriving_time,json=arrivingTime,proto3" json:"arriving_time,omitempty"`
	UserId        string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NumberOfSeats int32  `protobuf:"varint,5,opt,name=number_of_seats,json=numberOfSeats,proto3" json:"number_of_seats,omitempty"`
	CreatedAt     string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ReservationInfo) Reset() {
	*x = ReservationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationInfo) ProtoMessage() {}

func (x *ReservationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationInfo.ProtoReflect.Descriptor instead.
func (*ReservationInfo) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{5}
}

func (x *ReservationInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReservationInfo) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *ReservationInfo) GetArrivingTime() string {
	if x != nil {
		return x.ArrivingTime
	}
	return ""
}

func (x *ReservationInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationInfo) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

func (x *ReservationInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ReservationInfo) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Reservations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*ReservationInfo `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *Reservations) Reset() {
	*x = Reservations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservations) ProtoMessage() {}

func (x *Reservations) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservations.ProtoReflect.Descriptor instead.
func (*Reservations) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{6}
}

func (x *Reservations) GetReservations() []*ReservationInfo {
	if x != nil {
		return x.Reservations
	}
	return nil
}

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId  string `protobuf:"bytes,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	ArrivingTime  string `protobuf:"bytes,2,opt,name=arriving_time,json=arrivingTime,proto3" json:"arriving_time,omitempty"`
	UserId        string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NumberOfSeats int32  `protobuf:"varint,4,opt,name=number_of_seats,json=numberOfSeats,proto3" json:"number_of_seats,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{7}
}

func (x *Reservation) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *Reservation) GetArrivingTime() string {
	if x != nil {
		return x.ArrivingTime
	}
	return ""
}

func (x *Reservation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Reservation) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

type ReservationFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt     string `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	NumberOfSeats int32  `protobuf:"varint,2,opt,name=number_of_seats,json=numberOfSeats,proto3" json:"number_of_seats,omitempty"`
	RestaurantId  string `protobuf:"bytes,3,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	ArrivingTime  string `protobuf:"bytes,4,opt,name=arriving_time,json=arrivingTime,proto3" json:"arriving_time,omitempty"`
	UserId        string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit         int32  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32  `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ReservationFilter) Reset() {
	*x = ReservationFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationFilter) ProtoMessage() {}

func (x *ReservationFilter) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationFilter.ProtoReflect.Descriptor instead.
func (*ReservationFilter) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{8}
}

func (x *ReservationFilter) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ReservationFilter) GetNumberOfSeats() int32 {
	if x != nil {
		return x.NumberOfSeats
	}
	return 0
}

func (x *ReservationFilter) GetRestaurantId() string {
	if x != nil {
		return x.RestaurantId
	}
	return ""
}

func (x *ReservationFilter) GetArrivingTime() string {
	if x != nil {
		return x.ArrivingTime
	}
	return ""
}

func (x *ReservationFilter) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ReservationFilter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_reservation_proto protoreflect.FileDescriptor

var file_reservation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe9, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x6f, 0x66, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x22, 0xea, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f,
	0x66, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x50, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x98, 0x01, 0x0a,
	0x0b, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x61,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x4f, 0x66, 0x53, 0x65, 0x61, 0x74, 0x73, 0x22, 0xeb, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x53,
	0x65, 0x61, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x72, 0x72,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0xb8, 0x03, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x11, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64,
	0x1a, 0x11, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x1c, 0x2e, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x15,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x1a, 0x19, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00,
	0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_reservation_proto_rawDescOnce sync.Once
	file_reservation_proto_rawDescData = file_reservation_proto_rawDesc
)

func file_reservation_proto_rawDescGZIP() []byte {
	file_reservation_proto_rawDescOnce.Do(func() {
		file_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_proto_rawDescData)
	})
	return file_reservation_proto_rawDescData
}

var file_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_reservation_proto_goTypes = []any{
	(*Id)(nil),                // 0: reservation.Id
	(*MenuRespons)(nil),       // 1: reservation.MenuRespons
	(*Exists)(nil),            // 2: reservation.Exists
	(*Void)(nil),              // 3: reservation.Void
	(*ReservationUpdate)(nil), // 4: reservation.ReservationUpdate
	(*ReservationInfo)(nil),   // 5: reservation.ReservationInfo
	(*Reservations)(nil),      // 6: reservation.Reservations
	(*Reservation)(nil),       // 7: reservation.Reservation
	(*ReservationFilter)(nil), // 8: reservation.ReservationFilter
}
var file_reservation_proto_depIdxs = []int32{
	5, // 0: reservation.Reservations.reservations:type_name -> reservation.ReservationInfo
	7, // 1: reservation.ReservationService.CreateReservation:input_type -> reservation.Reservation
	4, // 2: reservation.ReservationService.UpdateReservation:input_type -> reservation.ReservationUpdate
	0, // 3: reservation.ReservationService.DeleteReservation:input_type -> reservation.Id
	0, // 4: reservation.ReservationService.GetReservationById:input_type -> reservation.Id
	0, // 5: reservation.ReservationService.ValidateReservationId:input_type -> reservation.Id
	8, // 6: reservation.ReservationService.GetAllReservations:input_type -> reservation.ReservationFilter
	3, // 7: reservation.ReservationService.CreateReservation:output_type -> reservation.Void
	3, // 8: reservation.ReservationService.UpdateReservation:output_type -> reservation.Void
	3, // 9: reservation.ReservationService.DeleteReservation:output_type -> reservation.Void
	5, // 10: reservation.ReservationService.GetReservationById:output_type -> reservation.ReservationInfo
	2, // 11: reservation.ReservationService.ValidateReservationId:output_type -> reservation.Exists
	6, // 12: reservation.ReservationService.GetAllReservations:output_type -> reservation.Reservations
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_reservation_proto_init() }
func file_reservation_proto_init() {
	if File_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Id); i {
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
		file_reservation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MenuRespons); i {
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
		file_reservation_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Exists); i {
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
		file_reservation_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Void); i {
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
		file_reservation_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ReservationUpdate); i {
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
		file_reservation_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ReservationInfo); i {
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
		file_reservation_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Reservations); i {
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
		file_reservation_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Reservation); i {
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
		file_reservation_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ReservationFilter); i {
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
			RawDescriptor: file_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_proto_goTypes,
		DependencyIndexes: file_reservation_proto_depIdxs,
		MessageInfos:      file_reservation_proto_msgTypes,
	}.Build()
	File_reservation_proto = out.File
	file_reservation_proto_rawDesc = nil
	file_reservation_proto_goTypes = nil
	file_reservation_proto_depIdxs = nil
}

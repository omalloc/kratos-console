// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: console/resource/zone.proto

package resource

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

// 可用区
type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 可用区名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 可用区代码
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// 地区名称
	RegionName string `protobuf:"bytes,4,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	// 地区代码
	RegionCode string `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// 环境
	Env string `protobuf:"bytes,6,opt,name=env,proto3" json:"env,omitempty"`
	// 创建时间
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{0}
}

func (x *Zone) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Zone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Zone) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Zone) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *Zone) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *Zone) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Zone) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Zone) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetZoneListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 页码
	Current int32 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	// 每页数量
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// 排序字段
	OrderBy string `protobuf:"bytes,3,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// 排序方式
	Order string `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetZoneListRequest) Reset() {
	*x = GetZoneListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneListRequest) ProtoMessage() {}

func (x *GetZoneListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneListRequest.ProtoReflect.Descriptor instead.
func (*GetZoneListRequest) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{1}
}

func (x *GetZoneListRequest) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *GetZoneListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetZoneListRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *GetZoneListRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type GetZoneListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区列表
	Data []*Zone `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// 总数
	Total int32 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetZoneListReply) Reset() {
	*x = GetZoneListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneListReply) ProtoMessage() {}

func (x *GetZoneListReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneListReply.ProtoReflect.Descriptor instead.
func (*GetZoneListReply) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{2}
}

func (x *GetZoneListReply) GetData() []*Zone {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetZoneListReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetZoneRequest) Reset() {
	*x = GetZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneRequest) ProtoMessage() {}

func (x *GetZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneRequest.ProtoReflect.Descriptor instead.
func (*GetZoneRequest) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{3}
}

func (x *GetZoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetZoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区
	Data *Zone `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetZoneReply) Reset() {
	*x = GetZoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneReply) ProtoMessage() {}

func (x *GetZoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneReply.ProtoReflect.Descriptor instead.
func (*GetZoneReply) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{4}
}

func (x *GetZoneReply) GetData() *Zone {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 可用区代码
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// 地区名称
	RegionName string `protobuf:"bytes,3,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	// 地区代码
	RegionCode string `protobuf:"bytes,4,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
}

func (x *CreateZoneRequest) Reset() {
	*x = CreateZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneRequest) ProtoMessage() {}

func (x *CreateZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneRequest.ProtoReflect.Descriptor instead.
func (*CreateZoneRequest) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{5}
}

func (x *CreateZoneRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateZoneRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateZoneRequest) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *CreateZoneRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

type CreateZoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区
	Data *Zone `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateZoneReply) Reset() {
	*x = CreateZoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneReply) ProtoMessage() {}

func (x *CreateZoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneReply.ProtoReflect.Descriptor instead.
func (*CreateZoneReply) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{6}
}

func (x *CreateZoneReply) GetData() *Zone {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 可用区名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 可用区代码
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	// 地区名称
	RegionName string `protobuf:"bytes,4,opt,name=region_name,json=regionName,proto3" json:"region_name,omitempty"`
	// 地区代码
	RegionCode string `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
}

func (x *UpdateZoneRequest) Reset() {
	*x = UpdateZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateZoneRequest) ProtoMessage() {}

func (x *UpdateZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateZoneRequest.ProtoReflect.Descriptor instead.
func (*UpdateZoneRequest) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateZoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateZoneRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateZoneRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateZoneRequest) GetRegionName() string {
	if x != nil {
		return x.RegionName
	}
	return ""
}

func (x *UpdateZoneRequest) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

type UpdateZoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateZoneReply) Reset() {
	*x = UpdateZoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateZoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateZoneReply) ProtoMessage() {}

func (x *UpdateZoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateZoneReply.ProtoReflect.Descriptor instead.
func (*UpdateZoneReply) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{8}
}

type DisableZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 可用区ID
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 备注
	Remark string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *DisableZoneRequest) Reset() {
	*x = DisableZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableZoneRequest) ProtoMessage() {}

func (x *DisableZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableZoneRequest.ProtoReflect.Descriptor instead.
func (*DisableZoneRequest) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{9}
}

func (x *DisableZoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DisableZoneRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type DisableZoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisableZoneReply) Reset() {
	*x = DisableZoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableZoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableZoneReply) ProtoMessage() {}

func (x *DisableZoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableZoneReply.ProtoReflect.Descriptor instead.
func (*DisableZoneReply) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{10}
}

type DeleteZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteZoneRequest) Reset() {
	*x = DeleteZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteZoneRequest) ProtoMessage() {}

func (x *DeleteZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteZoneRequest.ProtoReflect.Descriptor instead.
func (*DeleteZoneRequest) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteZoneRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteZoneReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteZoneReply) Reset() {
	*x = DeleteZoneReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_resource_zone_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteZoneReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteZoneReply) ProtoMessage() {}

func (x *DeleteZoneReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_resource_zone_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteZoneReply.ProtoReflect.Descriptor instead.
func (*DeleteZoneReply) Descriptor() ([]byte, []int) {
	return file_console_resource_zone_proto_rawDescGZIP(), []int{12}
}

var File_console_resource_zone_proto protoreflect.FileDescriptor

var file_console_resource_zone_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x04, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x58, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x20,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x7d, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x41, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3c, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x22, 0x12, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x59, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x2f, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_console_resource_zone_proto_rawDescOnce sync.Once
	file_console_resource_zone_proto_rawDescData = file_console_resource_zone_proto_rawDesc
)

func file_console_resource_zone_proto_rawDescGZIP() []byte {
	file_console_resource_zone_proto_rawDescOnce.Do(func() {
		file_console_resource_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_console_resource_zone_proto_rawDescData)
	})
	return file_console_resource_zone_proto_rawDescData
}

var file_console_resource_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_console_resource_zone_proto_goTypes = []interface{}{
	(*Zone)(nil),               // 0: api.console.resource.Zone
	(*GetZoneListRequest)(nil), // 1: api.console.resource.GetZoneListRequest
	(*GetZoneListReply)(nil),   // 2: api.console.resource.GetZoneListReply
	(*GetZoneRequest)(nil),     // 3: api.console.resource.GetZoneRequest
	(*GetZoneReply)(nil),       // 4: api.console.resource.GetZoneReply
	(*CreateZoneRequest)(nil),  // 5: api.console.resource.CreateZoneRequest
	(*CreateZoneReply)(nil),    // 6: api.console.resource.CreateZoneReply
	(*UpdateZoneRequest)(nil),  // 7: api.console.resource.UpdateZoneRequest
	(*UpdateZoneReply)(nil),    // 8: api.console.resource.UpdateZoneReply
	(*DisableZoneRequest)(nil), // 9: api.console.resource.DisableZoneRequest
	(*DisableZoneReply)(nil),   // 10: api.console.resource.DisableZoneReply
	(*DeleteZoneRequest)(nil),  // 11: api.console.resource.DeleteZoneRequest
	(*DeleteZoneReply)(nil),    // 12: api.console.resource.DeleteZoneReply
}
var file_console_resource_zone_proto_depIdxs = []int32{
	0, // 0: api.console.resource.GetZoneListReply.data:type_name -> api.console.resource.Zone
	0, // 1: api.console.resource.GetZoneReply.data:type_name -> api.console.resource.Zone
	0, // 2: api.console.resource.CreateZoneReply.data:type_name -> api.console.resource.Zone
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_console_resource_zone_proto_init() }
func file_console_resource_zone_proto_init() {
	if File_console_resource_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_console_resource_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		file_console_resource_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneListRequest); i {
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
		file_console_resource_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneListReply); i {
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
		file_console_resource_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneRequest); i {
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
		file_console_resource_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneReply); i {
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
		file_console_resource_zone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneRequest); i {
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
		file_console_resource_zone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneReply); i {
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
		file_console_resource_zone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateZoneRequest); i {
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
		file_console_resource_zone_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateZoneReply); i {
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
		file_console_resource_zone_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableZoneRequest); i {
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
		file_console_resource_zone_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableZoneReply); i {
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
		file_console_resource_zone_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteZoneRequest); i {
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
		file_console_resource_zone_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteZoneReply); i {
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
			RawDescriptor: file_console_resource_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_console_resource_zone_proto_goTypes,
		DependencyIndexes: file_console_resource_zone_proto_depIdxs,
		MessageInfos:      file_console_resource_zone_proto_msgTypes,
	}.Build()
	File_console_resource_zone_proto = out.File
	file_console_resource_zone_proto_rawDesc = nil
	file_console_resource_zone_proto_goTypes = nil
	file_console_resource_zone_proto_depIdxs = nil
}

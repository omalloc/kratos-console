// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: console/discovery/disocvery.proto

package discovery

import (
	protobuf "github.com/omalloc/contrib/protobuf"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一标识
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// 主机名
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// 服务名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 版本
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// 服务地址
	Endpoints []string `protobuf:"bytes,5,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// 集群
	Cluster string `protobuf:"bytes,6,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// 挂起状态 true 为挂起
	Hang bool `protobuf:"varint,7,opt,name=hang,proto3" json:"hang,omitempty"`
	// 元信息
	Metadata map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 命名空间
	Namespace string `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Service) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *Service) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Service) GetHang() bool {
	if x != nil {
		return x.Hang
	}
	return false
}

func (x *Service) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Service) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ServiceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 唯一标识
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// 主机名
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// 服务名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 版本
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// 服务地址
	Endpoints []string `protobuf:"bytes,5,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// 集群
	Clusters []string `protobuf:"bytes,6,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// 同服务，不同实例的数组信息
	Children []*Service `protobuf:"bytes,7,rep,name=children,proto3" json:"children,omitempty"`
	// 命名空间
	Namespace string `protobuf:"bytes,8,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 数据类型
	IsGroup bool `protobuf:"varint,9,opt,name=is_group,json=isGroup,proto3" json:"is_group,omitempty"`
}

func (x *ServiceGroup) Reset() {
	*x = ServiceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceGroup) ProtoMessage() {}

func (x *ServiceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceGroup.ProtoReflect.Descriptor instead.
func (*ServiceGroup) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceGroup) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ServiceGroup) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ServiceGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceGroup) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceGroup) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ServiceGroup) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *ServiceGroup) GetChildren() []*Service {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *ServiceGroup) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ServiceGroup) GetIsGroup() bool {
	if x != nil {
		return x.IsGroup
	}
	return false
}

type OnlineServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pagination filter
	Pagination *protobuf.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	// service name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *OnlineServiceRequest) Reset() {
	*x = OnlineServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineServiceRequest) ProtoMessage() {}

func (x *OnlineServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineServiceRequest.ProtoReflect.Descriptor instead.
func (*OnlineServiceRequest) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{2}
}

func (x *OnlineServiceRequest) GetPagination() *protobuf.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *OnlineServiceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type OnlineServiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *protobuf.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Data       []*ServiceGroup      `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *OnlineServiceReply) Reset() {
	*x = OnlineServiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlineServiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlineServiceReply) ProtoMessage() {}

func (x *OnlineServiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlineServiceReply.ProtoReflect.Descriptor instead.
func (*OnlineServiceReply) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{3}
}

func (x *OnlineServiceReply) GetPagination() *protobuf.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *OnlineServiceReply) GetData() []*ServiceGroup {
	if x != nil {
		return x.Data
	}
	return nil
}

type KVListClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KVListClustersRequest) Reset() {
	*x = KVListClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVListClustersRequest) ProtoMessage() {}

func (x *KVListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVListClustersRequest.ProtoReflect.Descriptor instead.
func (*KVListClustersRequest) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{4}
}

type KVListClustersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []string `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *KVListClustersReply) Reset() {
	*x = KVListClustersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVListClustersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVListClustersReply) ProtoMessage() {}

func (x *KVListClustersReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVListClustersReply.ProtoReflect.Descriptor instead.
func (*KVListClustersReply) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{5}
}

func (x *KVListClustersReply) GetClusters() []string {
	if x != nil {
		return x.Clusters
	}
	return nil
}

type KVListKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *KVListKeysRequest) Reset() {
	*x = KVListKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVListKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVListKeysRequest) ProtoMessage() {}

func (x *KVListKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVListKeysRequest.ProtoReflect.Descriptor instead.
func (*KVListKeysRequest) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{6}
}

func (x *KVListKeysRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type KVListKeysReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *KVListKeysReply) Reset() {
	*x = KVListKeysReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVListKeysReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVListKeysReply) ProtoMessage() {}

func (x *KVListKeysReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVListKeysReply.ProtoReflect.Descriptor instead.
func (*KVListKeysReply) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{7}
}

func (x *KVListKeysReply) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type KVGetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
}

func (x *KVGetValueRequest) Reset() {
	*x = KVGetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVGetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGetValueRequest) ProtoMessage() {}

func (x *KVGetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGetValueRequest.ProtoReflect.Descriptor instead.
func (*KVGetValueRequest) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{8}
}

func (x *KVGetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVGetValueRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

type KVGetValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVGetValueReply) Reset() {
	*x = KVGetValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_console_discovery_disocvery_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVGetValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVGetValueReply) ProtoMessage() {}

func (x *KVGetValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_console_discovery_disocvery_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVGetValueReply.ProtoReflect.Descriptor instead.
func (*KVGetValueReply) Descriptor() ([]byte, []int) {
	return file_console_discovery_disocvery_proto_rawDescGZIP(), []int{9}
}

func (x *KVGetValueReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_console_discovery_disocvery_proto protoreflect.FileDescriptor

var file_console_discovery_disocvery_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x64, 0x69, 0x73, 0x6f, 0x63, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x68, 0x61, 0x6e, 0x67, 0x12, 0x48, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a,
	0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x99, 0x02, 0x0a,
	0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x60, 0x0a, 0x14, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x34, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x17, 0x0a, 0x15, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x4b, 0x56, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x2d, 0x0a, 0x11,
	0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0f, 0x4b,
	0x56, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65,
	0x79, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x0f, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb6, 0x04, 0x0a,
	0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x9a, 0x01, 0x0a, 0x0e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x2d,
	0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x0e, 0x4b, 0x56, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x6b, 0x76, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x7c, 0x0a, 0x0a, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x56, 0x4c,
	0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x56, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x6b, 0x76, 0x2f,
	0x6b, 0x65, 0x79, 0x73, 0x12, 0x7f, 0x0a, 0x0a, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x56, 0x47, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x4b, 0x56, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x6b, 0x76, 0x2f, 0x2d, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x5c, 0x0a, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x01,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6d, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x73,
	0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_console_discovery_disocvery_proto_rawDescOnce sync.Once
	file_console_discovery_disocvery_proto_rawDescData = file_console_discovery_disocvery_proto_rawDesc
)

func file_console_discovery_disocvery_proto_rawDescGZIP() []byte {
	file_console_discovery_disocvery_proto_rawDescOnce.Do(func() {
		file_console_discovery_disocvery_proto_rawDescData = protoimpl.X.CompressGZIP(file_console_discovery_disocvery_proto_rawDescData)
	})
	return file_console_discovery_disocvery_proto_rawDescData
}

var file_console_discovery_disocvery_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_console_discovery_disocvery_proto_goTypes = []interface{}{
	(*Service)(nil),               // 0: api.console.discovery.Service
	(*ServiceGroup)(nil),          // 1: api.console.discovery.ServiceGroup
	(*OnlineServiceRequest)(nil),  // 2: api.console.discovery.OnlineServiceRequest
	(*OnlineServiceReply)(nil),    // 3: api.console.discovery.OnlineServiceReply
	(*KVListClustersRequest)(nil), // 4: api.console.discovery.KVListClustersRequest
	(*KVListClustersReply)(nil),   // 5: api.console.discovery.KVListClustersReply
	(*KVListKeysRequest)(nil),     // 6: api.console.discovery.KVListKeysRequest
	(*KVListKeysReply)(nil),       // 7: api.console.discovery.KVListKeysReply
	(*KVGetValueRequest)(nil),     // 8: api.console.discovery.KVGetValueRequest
	(*KVGetValueReply)(nil),       // 9: api.console.discovery.KVGetValueReply
	nil,                           // 10: api.console.discovery.Service.MetadataEntry
	(*protobuf.Pagination)(nil),   // 11: protobuf.Pagination
}
var file_console_discovery_disocvery_proto_depIdxs = []int32{
	10, // 0: api.console.discovery.Service.metadata:type_name -> api.console.discovery.Service.MetadataEntry
	0,  // 1: api.console.discovery.ServiceGroup.children:type_name -> api.console.discovery.Service
	11, // 2: api.console.discovery.OnlineServiceRequest.pagination:type_name -> protobuf.Pagination
	11, // 3: api.console.discovery.OnlineServiceReply.pagination:type_name -> protobuf.Pagination
	1,  // 4: api.console.discovery.OnlineServiceReply.data:type_name -> api.console.discovery.ServiceGroup
	2,  // 5: api.console.discovery.Discovery.OnlineServices:input_type -> api.console.discovery.OnlineServiceRequest
	4,  // 6: api.console.discovery.Discovery.KVListClusters:input_type -> api.console.discovery.KVListClustersRequest
	6,  // 7: api.console.discovery.Discovery.KVListKeys:input_type -> api.console.discovery.KVListKeysRequest
	8,  // 8: api.console.discovery.Discovery.KVGetValue:input_type -> api.console.discovery.KVGetValueRequest
	3,  // 9: api.console.discovery.Discovery.OnlineServices:output_type -> api.console.discovery.OnlineServiceReply
	5,  // 10: api.console.discovery.Discovery.KVListClusters:output_type -> api.console.discovery.KVListClustersReply
	7,  // 11: api.console.discovery.Discovery.KVListKeys:output_type -> api.console.discovery.KVListKeysReply
	9,  // 12: api.console.discovery.Discovery.KVGetValue:output_type -> api.console.discovery.KVGetValueReply
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_console_discovery_disocvery_proto_init() }
func file_console_discovery_disocvery_proto_init() {
	if File_console_discovery_disocvery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_console_discovery_disocvery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_console_discovery_disocvery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceGroup); i {
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
		file_console_discovery_disocvery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineServiceRequest); i {
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
		file_console_discovery_disocvery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlineServiceReply); i {
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
		file_console_discovery_disocvery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVListClustersRequest); i {
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
		file_console_discovery_disocvery_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVListClustersReply); i {
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
		file_console_discovery_disocvery_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVListKeysRequest); i {
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
		file_console_discovery_disocvery_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVListKeysReply); i {
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
		file_console_discovery_disocvery_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVGetValueRequest); i {
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
		file_console_discovery_disocvery_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVGetValueReply); i {
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
			RawDescriptor: file_console_discovery_disocvery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_console_discovery_disocvery_proto_goTypes,
		DependencyIndexes: file_console_discovery_disocvery_proto_depIdxs,
		MessageInfos:      file_console_discovery_disocvery_proto_msgTypes,
	}.Build()
	File_console_discovery_disocvery_proto = out.File
	file_console_discovery_disocvery_proto_rawDesc = nil
	file_console_discovery_disocvery_proto_goTypes = nil
	file_console_discovery_disocvery_proto_depIdxs = nil
}

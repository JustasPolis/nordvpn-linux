// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: service_response.proto

package pb

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

// ServiceErrorCode defines a set of error codes which handling
// does not depend on any specific command used.
type ServiceErrorCode int32

const (
	ServiceErrorCode_NOT_LOGGED_IN  ServiceErrorCode = 0
	ServiceErrorCode_API_FAILURE    ServiceErrorCode = 1
	ServiceErrorCode_CONFIG_FAILURE ServiceErrorCode = 2
)

// Enum value maps for ServiceErrorCode.
var (
	ServiceErrorCode_name = map[int32]string{
		0: "NOT_LOGGED_IN",
		1: "API_FAILURE",
		2: "CONFIG_FAILURE",
	}
	ServiceErrorCode_value = map[string]int32{
		"NOT_LOGGED_IN":  0,
		"API_FAILURE":    1,
		"CONFIG_FAILURE": 2,
	}
)

func (x ServiceErrorCode) Enum() *ServiceErrorCode {
	p := new(ServiceErrorCode)
	*p = x
	return p
}

func (x ServiceErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_service_response_proto_enumTypes[0].Descriptor()
}

func (ServiceErrorCode) Type() protoreflect.EnumType {
	return &file_service_response_proto_enumTypes[0]
}

func (x ServiceErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceErrorCode.Descriptor instead.
func (ServiceErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_service_response_proto_rawDescGZIP(), []int{0}
}

// MeshnetErrorCode defines a set of meshnet specific error codes.
type MeshnetErrorCode int32

const (
	MeshnetErrorCode_NOT_REGISTERED   MeshnetErrorCode = 0
	MeshnetErrorCode_LIB_FAILURE      MeshnetErrorCode = 1
	MeshnetErrorCode_ALREADY_ENABLED  MeshnetErrorCode = 3
	MeshnetErrorCode_ALREADY_DISABLED MeshnetErrorCode = 4
	MeshnetErrorCode_NOT_ENABLED      MeshnetErrorCode = 5
	MeshnetErrorCode_TECH_FAILURE     MeshnetErrorCode = 6
	MeshnetErrorCode_TUNNEL_CLOSED    MeshnetErrorCode = 7
)

// Enum value maps for MeshnetErrorCode.
var (
	MeshnetErrorCode_name = map[int32]string{
		0: "NOT_REGISTERED",
		1: "LIB_FAILURE",
		3: "ALREADY_ENABLED",
		4: "ALREADY_DISABLED",
		5: "NOT_ENABLED",
		6: "TECH_FAILURE",
		7: "TUNNEL_CLOSED",
	}
	MeshnetErrorCode_value = map[string]int32{
		"NOT_REGISTERED":   0,
		"LIB_FAILURE":      1,
		"ALREADY_ENABLED":  3,
		"ALREADY_DISABLED": 4,
		"NOT_ENABLED":      5,
		"TECH_FAILURE":     6,
		"TUNNEL_CLOSED":    7,
	}
)

func (x MeshnetErrorCode) Enum() *MeshnetErrorCode {
	p := new(MeshnetErrorCode)
	*p = x
	return p
}

func (x MeshnetErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MeshnetErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_service_response_proto_enumTypes[1].Descriptor()
}

func (MeshnetErrorCode) Type() protoreflect.EnumType {
	return &file_service_response_proto_enumTypes[1]
}

func (x MeshnetErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MeshnetErrorCode.Descriptor instead.
func (MeshnetErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_service_response_proto_rawDescGZIP(), []int{1}
}

// MeshnetErrorCode is one of the:
// - Empty response
// - Service error
// - Meshnet error
type MeshnetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*MeshnetResponse_Empty
	//	*MeshnetResponse_ServiceError
	//	*MeshnetResponse_MeshnetError
	Response isMeshnetResponse_Response `protobuf_oneof:"response"`
}

func (x *MeshnetResponse) Reset() {
	*x = MeshnetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshnetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshnetResponse) ProtoMessage() {}

func (x *MeshnetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshnetResponse.ProtoReflect.Descriptor instead.
func (*MeshnetResponse) Descriptor() ([]byte, []int) {
	return file_service_response_proto_rawDescGZIP(), []int{0}
}

func (m *MeshnetResponse) GetResponse() isMeshnetResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *MeshnetResponse) GetEmpty() *Empty {
	if x, ok := x.GetResponse().(*MeshnetResponse_Empty); ok {
		return x.Empty
	}
	return nil
}

func (x *MeshnetResponse) GetServiceError() ServiceErrorCode {
	if x, ok := x.GetResponse().(*MeshnetResponse_ServiceError); ok {
		return x.ServiceError
	}
	return ServiceErrorCode_NOT_LOGGED_IN
}

func (x *MeshnetResponse) GetMeshnetError() MeshnetErrorCode {
	if x, ok := x.GetResponse().(*MeshnetResponse_MeshnetError); ok {
		return x.MeshnetError
	}
	return MeshnetErrorCode_NOT_REGISTERED
}

type isMeshnetResponse_Response interface {
	isMeshnetResponse_Response()
}

type MeshnetResponse_Empty struct {
	Empty *Empty `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

type MeshnetResponse_ServiceError struct {
	ServiceError ServiceErrorCode `protobuf:"varint,2,opt,name=service_error,json=serviceError,proto3,enum=meshpb.ServiceErrorCode,oneof"`
}

type MeshnetResponse_MeshnetError struct {
	MeshnetError MeshnetErrorCode `protobuf:"varint,3,opt,name=meshnet_error,json=meshnetError,proto3,enum=meshpb.MeshnetErrorCode,oneof"`
}

func (*MeshnetResponse_Empty) isMeshnetResponse_Response() {}

func (*MeshnetResponse_ServiceError) isMeshnetResponse_Response() {}

func (*MeshnetResponse_MeshnetError) isMeshnetResponse_Response() {}

// ServiceResponse is either an empty response or a service error
type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*ServiceResponse_Empty
	//	*ServiceResponse_ErrorCode
	Response isServiceResponse_Response `protobuf_oneof:"response"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_service_response_proto_rawDescGZIP(), []int{1}
}

func (m *ServiceResponse) GetResponse() isServiceResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ServiceResponse) GetEmpty() *Empty {
	if x, ok := x.GetResponse().(*ServiceResponse_Empty); ok {
		return x.Empty
	}
	return nil
}

func (x *ServiceResponse) GetErrorCode() ServiceErrorCode {
	if x, ok := x.GetResponse().(*ServiceResponse_ErrorCode); ok {
		return x.ErrorCode
	}
	return ServiceErrorCode_NOT_LOGGED_IN
}

type isServiceResponse_Response interface {
	isServiceResponse_Response()
}

type ServiceResponse_Empty struct {
	Empty *Empty `protobuf:"bytes,1,opt,name=empty,proto3,oneof"`
}

type ServiceResponse_ErrorCode struct {
	ErrorCode ServiceErrorCode `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=meshpb.ServiceErrorCode,oneof"`
}

func (*ServiceResponse_Empty) isServiceResponse_Response() {}

func (*ServiceResponse_ErrorCode) isServiceResponse_Response() {}

// ServiceBoolResponse is either a bool response or a service error
type ServiceBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*ServiceBoolResponse_Value
	//	*ServiceBoolResponse_ErrorCode
	Response isServiceBoolResponse_Response `protobuf_oneof:"response"`
}

func (x *ServiceBoolResponse) Reset() {
	*x = ServiceBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceBoolResponse) ProtoMessage() {}

func (x *ServiceBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceBoolResponse.ProtoReflect.Descriptor instead.
func (*ServiceBoolResponse) Descriptor() ([]byte, []int) {
	return file_service_response_proto_rawDescGZIP(), []int{2}
}

func (m *ServiceBoolResponse) GetResponse() isServiceBoolResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *ServiceBoolResponse) GetValue() bool {
	if x, ok := x.GetResponse().(*ServiceBoolResponse_Value); ok {
		return x.Value
	}
	return false
}

func (x *ServiceBoolResponse) GetErrorCode() ServiceErrorCode {
	if x, ok := x.GetResponse().(*ServiceBoolResponse_ErrorCode); ok {
		return x.ErrorCode
	}
	return ServiceErrorCode_NOT_LOGGED_IN
}

type isServiceBoolResponse_Response interface {
	isServiceBoolResponse_Response()
}

type ServiceBoolResponse_Value struct {
	Value bool `protobuf:"varint,1,opt,name=value,proto3,oneof"`
}

type ServiceBoolResponse_ErrorCode struct {
	ErrorCode ServiceErrorCode `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3,enum=meshpb.ServiceErrorCode,oneof"`
}

func (*ServiceBoolResponse_Value) isServiceBoolResponse_Response() {}

func (*ServiceBoolResponse_ErrorCode) isServiceBoolResponse_Response() {}

var File_service_response_proto protoreflect.FileDescriptor

var file_service_response_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x62,
	0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01,
	0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48,
	0x00, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x18, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x65, 0x73,
	0x68, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x6e, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x65,
	0x73, 0x68, 0x6e, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7f, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x70,
	0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x39, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x48, 0x00,
	0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x74, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x48, 0x00, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x4a, 0x0a,
	0x10, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f,
	0x49, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x49, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x2a, 0x98, 0x01, 0x0a, 0x10, 0x4d, 0x65,
	0x73, 0x68, 0x6e, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x0e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x49, 0x42, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45,
	0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0f,
	0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x10, 0x0a, 0x0c, 0x54, 0x45, 0x43, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x06, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x55, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x44, 0x10, 0x07, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f,
	0x6e, 0x6f, 0x72, 0x64, 0x76, 0x70, 0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x6d, 0x65,
	0x73, 0x68, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_response_proto_rawDescOnce sync.Once
	file_service_response_proto_rawDescData = file_service_response_proto_rawDesc
)

func file_service_response_proto_rawDescGZIP() []byte {
	file_service_response_proto_rawDescOnce.Do(func() {
		file_service_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_response_proto_rawDescData)
	})
	return file_service_response_proto_rawDescData
}

var file_service_response_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_service_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_response_proto_goTypes = []interface{}{
	(ServiceErrorCode)(0),       // 0: meshpb.ServiceErrorCode
	(MeshnetErrorCode)(0),       // 1: meshpb.MeshnetErrorCode
	(*MeshnetResponse)(nil),     // 2: meshpb.MeshnetResponse
	(*ServiceResponse)(nil),     // 3: meshpb.ServiceResponse
	(*ServiceBoolResponse)(nil), // 4: meshpb.ServiceBoolResponse
	(*Empty)(nil),               // 5: meshpb.Empty
}
var file_service_response_proto_depIdxs = []int32{
	5, // 0: meshpb.MeshnetResponse.empty:type_name -> meshpb.Empty
	0, // 1: meshpb.MeshnetResponse.service_error:type_name -> meshpb.ServiceErrorCode
	1, // 2: meshpb.MeshnetResponse.meshnet_error:type_name -> meshpb.MeshnetErrorCode
	5, // 3: meshpb.ServiceResponse.empty:type_name -> meshpb.Empty
	0, // 4: meshpb.ServiceResponse.error_code:type_name -> meshpb.ServiceErrorCode
	0, // 5: meshpb.ServiceBoolResponse.error_code:type_name -> meshpb.ServiceErrorCode
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_service_response_proto_init() }
func file_service_response_proto_init() {
	if File_service_response_proto != nil {
		return
	}
	file_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshnetResponse); i {
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
		file_service_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_service_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceBoolResponse); i {
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
	file_service_response_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MeshnetResponse_Empty)(nil),
		(*MeshnetResponse_ServiceError)(nil),
		(*MeshnetResponse_MeshnetError)(nil),
	}
	file_service_response_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ServiceResponse_Empty)(nil),
		(*ServiceResponse_ErrorCode)(nil),
	}
	file_service_response_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ServiceBoolResponse_Value)(nil),
		(*ServiceBoolResponse_ErrorCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_response_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_response_proto_goTypes,
		DependencyIndexes: file_service_response_proto_depIdxs,
		EnumInfos:         file_service_response_proto_enumTypes,
		MessageInfos:      file_service_response_proto_msgTypes,
	}.Build()
	File_service_response_proto = out.File
	file_service_response_proto_rawDesc = nil
	file_service_response_proto_goTypes = nil
	file_service_response_proto_depIdxs = nil
}

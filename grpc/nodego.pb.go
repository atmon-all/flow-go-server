// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0--rc2
// source: nodego.proto

package grpc

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

// Execute next node request.
type ExecuteNextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Flow ID.
	FlowId string `protobuf:"bytes,1,opt,name=flowId,proto3" json:"flowId,omitempty"`
	// Flow version.
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Current node.
	Current *FlowNode `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	// Next node.
	Next *FlowNode `protobuf:"bytes,4,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *ExecuteNextRequest) Reset() {
	*x = ExecuteNextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodego_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteNextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteNextRequest) ProtoMessage() {}

func (x *ExecuteNextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nodego_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteNextRequest.ProtoReflect.Descriptor instead.
func (*ExecuteNextRequest) Descriptor() ([]byte, []int) {
	return file_nodego_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteNextRequest) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *ExecuteNextRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ExecuteNextRequest) GetCurrent() *FlowNode {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *ExecuteNextRequest) GetNext() *FlowNode {
	if x != nil {
		return x.Next
	}
	return nil
}

// Execute next node response.
type ExecuteNextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Responsed code from data agent.
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// Responsed message from data agent.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ExecuteNextResponse) Reset() {
	*x = ExecuteNextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodego_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteNextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteNextResponse) ProtoMessage() {}

func (x *ExecuteNextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nodego_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteNextResponse.ProtoReflect.Descriptor instead.
func (*ExecuteNextResponse) Descriptor() ([]byte, []int) {
	return file_nodego_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteNextResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ExecuteNextResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Flow node.
type FlowNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Flow node ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Flow node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Flow node description.
	Desc string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// Flow node server type, values are 'Golang','Java' or 'NodeJS'.
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Flow context define.
	Context string `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *FlowNode) Reset() {
	*x = FlowNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodego_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowNode) ProtoMessage() {}

func (x *FlowNode) ProtoReflect() protoreflect.Message {
	mi := &file_nodego_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowNode.ProtoReflect.Descriptor instead.
func (*FlowNode) Descriptor() ([]byte, []int) {
	return file_nodego_proto_rawDescGZIP(), []int{2}
}

func (x *FlowNode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FlowNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlowNode) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *FlowNode) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FlowNode) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

var File_nodego_proto protoreflect.FileDescriptor

var file_nodego_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x67, 0x6f, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x22, 0x43, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x08, 0x46, 0x6c, 0x6f, 0x77, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x32, 0x52, 0x0a, 0x06, 0x4e, 0x6f, 0x64, 0x65,
	0x67, 0x6f, 0x12, 0x48, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x65, 0x78,
	0x74, 0x12, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x4e, 0x65,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x74, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x50, 0x01, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x74, 0x6d, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodego_proto_rawDescOnce sync.Once
	file_nodego_proto_rawDescData = file_nodego_proto_rawDesc
)

func file_nodego_proto_rawDescGZIP() []byte {
	file_nodego_proto_rawDescOnce.Do(func() {
		file_nodego_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodego_proto_rawDescData)
	})
	return file_nodego_proto_rawDescData
}

var file_nodego_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_nodego_proto_goTypes = []interface{}{
	(*ExecuteNextRequest)(nil),  // 0: nodego.ExecuteNextRequest
	(*ExecuteNextResponse)(nil), // 1: nodego.ExecuteNextResponse
	(*FlowNode)(nil),            // 2: nodego.FlowNode
}
var file_nodego_proto_depIdxs = []int32{
	2, // 0: nodego.ExecuteNextRequest.current:type_name -> nodego.FlowNode
	2, // 1: nodego.ExecuteNextRequest.next:type_name -> nodego.FlowNode
	0, // 2: nodego.Nodego.ExecuteNext:input_type -> nodego.ExecuteNextRequest
	1, // 3: nodego.Nodego.ExecuteNext:output_type -> nodego.ExecuteNextResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nodego_proto_init() }
func file_nodego_proto_init() {
	if File_nodego_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodego_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteNextRequest); i {
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
		file_nodego_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteNextResponse); i {
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
		file_nodego_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowNode); i {
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
			RawDescriptor: file_nodego_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodego_proto_goTypes,
		DependencyIndexes: file_nodego_proto_depIdxs,
		MessageInfos:      file_nodego_proto_msgTypes,
	}.Build()
	File_nodego_proto = out.File
	file_nodego_proto_rawDesc = nil
	file_nodego_proto_goTypes = nil
	file_nodego_proto_depIdxs = nil
}

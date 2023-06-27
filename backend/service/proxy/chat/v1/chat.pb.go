// Code generated by github.com/tinkler/mqttadmin; DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: chat/v1/chat.proto

package chat_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages  []*ChatCompletionMessage  `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	MaxTokens int32                     `protobuf:"varint,2,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	Model     string                    `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Stoped    bool                      `protobuf:"varint,4,opt,name=stoped,proto3" json:"stoped,omitempty"`
	Functions []*ChatCompletionFunction `protobuf:"bytes,5,rep,name=functions,proto3" json:"functions,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetMessages() []*ChatCompletionMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Chat) GetMaxTokens() int32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *Chat) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Chat) GetStoped() bool {
	if x != nil {
		return x.Stoped
	}
	return false
}

func (x *Chat) GetFunctions() []*ChatCompletionFunction {
	if x != nil {
		return x.Functions
	}
	return nil
}

type ChatCompletionMessageFunctionCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Arguments string `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *ChatCompletionMessageFunctionCall) Reset() {
	*x = ChatCompletionMessageFunctionCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionMessageFunctionCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionMessageFunctionCall) ProtoMessage() {}

func (x *ChatCompletionMessageFunctionCall) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionMessageFunctionCall.ProtoReflect.Descriptor instead.
func (*ChatCompletionMessageFunctionCall) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatCompletionMessageFunctionCall) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatCompletionMessageFunctionCall) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

type ChatCompletionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role         string                             `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Content      string                             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Name         string                             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	FunctionCall *ChatCompletionMessageFunctionCall `protobuf:"bytes,4,opt,name=function_call,json=functionCall,proto3" json:"function_call,omitempty"`
	FinishReason string                             `protobuf:"bytes,5,opt,name=finish_reason,json=finishReason,proto3" json:"finish_reason,omitempty"`
}

func (x *ChatCompletionMessage) Reset() {
	*x = ChatCompletionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionMessage) ProtoMessage() {}

func (x *ChatCompletionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionMessage.ProtoReflect.Descriptor instead.
func (*ChatCompletionMessage) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatCompletionMessage) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ChatCompletionMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatCompletionMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatCompletionMessage) GetFunctionCall() *ChatCompletionMessageFunctionCall {
	if x != nil {
		return x.FunctionCall
	}
	return nil
}

func (x *ChatCompletionMessage) GetFinishReason() string {
	if x != nil {
		return x.FinishReason
	}
	return ""
}

type ChatCompletionFunctionParametersProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Enum        []string `protobuf:"bytes,3,rep,name=enum,proto3" json:"enum,omitempty"`
}

func (x *ChatCompletionFunctionParametersProperty) Reset() {
	*x = ChatCompletionFunctionParametersProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionFunctionParametersProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionFunctionParametersProperty) ProtoMessage() {}

func (x *ChatCompletionFunctionParametersProperty) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionFunctionParametersProperty.ProtoReflect.Descriptor instead.
func (*ChatCompletionFunctionParametersProperty) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatCompletionFunctionParametersProperty) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChatCompletionFunctionParametersProperty) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChatCompletionFunctionParametersProperty) GetEnum() []string {
	if x != nil {
		return x.Enum
	}
	return nil
}

type ChatCompletionFunctionParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string                                               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Properties map[string]*ChatCompletionFunctionParametersProperty `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Required   []string                                             `protobuf:"bytes,3,rep,name=required,proto3" json:"required,omitempty"`
}

func (x *ChatCompletionFunctionParameters) Reset() {
	*x = ChatCompletionFunctionParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionFunctionParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionFunctionParameters) ProtoMessage() {}

func (x *ChatCompletionFunctionParameters) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionFunctionParameters.ProtoReflect.Descriptor instead.
func (*ChatCompletionFunctionParameters) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ChatCompletionFunctionParameters) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ChatCompletionFunctionParameters) GetProperties() map[string]*ChatCompletionFunctionParametersProperty {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *ChatCompletionFunctionParameters) GetRequired() []string {
	if x != nil {
		return x.Required
	}
	return nil
}

type ChatCompletionFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string                            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Parameters  *ChatCompletionFunctionParameters `protobuf:"bytes,3,opt,name=parameters,proto3" json:"parameters,omitempty"`
}

func (x *ChatCompletionFunction) Reset() {
	*x = ChatCompletionFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_v1_chat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatCompletionFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatCompletionFunction) ProtoMessage() {}

func (x *ChatCompletionFunction) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatCompletionFunction.ProtoReflect.Descriptor instead.
func (*ChatCompletionFunction) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{5}
}

func (x *ChatCompletionFunction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatCompletionFunction) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChatCompletionFunction) GetParameters() *ChatCompletionFunctionParameters {
	if x != nil {
		return x.Parameters
	}
	return nil
}

var File_chat_v1_chat_proto protoreflect.FileDescriptor

var file_chat_v1_chat_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x3a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x70, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x70, 0x65, 0x64, 0x12, 0x3d, 0x0a, 0x09, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x55, 0x0a, 0x21, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xcf, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x0d,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x74, 0x0a, 0x28, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x22, 0x9f, 0x02, 0x0a, 0x20, 0x43, 0x68, 0x61,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x59, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x1a, 0x70, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x47, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x99, 0x01, 0x0a, 0x16, 0x43,
	0x68, 0x61, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x32, 0x8a, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x47,
	0x73, 0x72, 0x76, 0x12, 0x39, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6e, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x12, 0x43,
	0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x6f, 0x0a, 0x21, 0x69, 0x6e, 0x6b, 0x2e, 0x73, 0x66, 0x73, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x71, 0x74, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x63, 0x68, 0x61, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x73, 0x6d, 0x73, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_v1_chat_proto_rawDescOnce sync.Once
	file_chat_v1_chat_proto_rawDescData = file_chat_v1_chat_proto_rawDesc
)

func file_chat_v1_chat_proto_rawDescGZIP() []byte {
	file_chat_v1_chat_proto_rawDescOnce.Do(func() {
		file_chat_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_v1_chat_proto_rawDescData)
	})
	return file_chat_v1_chat_proto_rawDescData
}

var file_chat_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chat_v1_chat_proto_goTypes = []interface{}{
	(*Chat)(nil), // 0: chat.v1.Chat
	(*ChatCompletionMessageFunctionCall)(nil),        // 1: chat.v1.ChatCompletionMessageFunctionCall
	(*ChatCompletionMessage)(nil),                    // 2: chat.v1.ChatCompletionMessage
	(*ChatCompletionFunctionParametersProperty)(nil), // 3: chat.v1.ChatCompletionFunctionParametersProperty
	(*ChatCompletionFunctionParameters)(nil),         // 4: chat.v1.ChatCompletionFunctionParameters
	(*ChatCompletionFunction)(nil),                   // 5: chat.v1.ChatCompletionFunction
	nil,                                              // 6: chat.v1.ChatCompletionFunctionParameters.PropertiesEntry
	(*anypb.Any)(nil),                                // 7: google.protobuf.Any
}
var file_chat_v1_chat_proto_depIdxs = []int32{
	2, // 0: chat.v1.Chat.messages:type_name -> chat.v1.ChatCompletionMessage
	5, // 1: chat.v1.Chat.functions:type_name -> chat.v1.ChatCompletionFunction
	1, // 2: chat.v1.ChatCompletionMessage.function_call:type_name -> chat.v1.ChatCompletionMessageFunctionCall
	6, // 3: chat.v1.ChatCompletionFunctionParameters.properties:type_name -> chat.v1.ChatCompletionFunctionParameters.PropertiesEntry
	4, // 4: chat.v1.ChatCompletionFunction.parameters:type_name -> chat.v1.ChatCompletionFunctionParameters
	3, // 5: chat.v1.ChatCompletionFunctionParameters.PropertiesEntry.value:type_name -> chat.v1.ChatCompletionFunctionParametersProperty
	7, // 6: chat.v1.ChatGsrv.ChatEnquire:input_type -> google.protobuf.Any
	7, // 7: chat.v1.ChatGsrv.ChatEnquireStream:input_type -> google.protobuf.Any
	7, // 8: chat.v1.ChatGsrv.ChatEnquire:output_type -> google.protobuf.Any
	7, // 9: chat.v1.ChatGsrv.ChatEnquireStream:output_type -> google.protobuf.Any
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chat_v1_chat_proto_init() }
func file_chat_v1_chat_proto_init() {
	if File_chat_v1_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_v1_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_chat_v1_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionMessageFunctionCall); i {
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
		file_chat_v1_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionMessage); i {
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
		file_chat_v1_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionFunctionParametersProperty); i {
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
		file_chat_v1_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionFunctionParameters); i {
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
		file_chat_v1_chat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatCompletionFunction); i {
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
			RawDescriptor: file_chat_v1_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_v1_chat_proto_goTypes,
		DependencyIndexes: file_chat_v1_chat_proto_depIdxs,
		MessageInfos:      file_chat_v1_chat_proto_msgTypes,
	}.Build()
	File_chat_v1_chat_proto = out.File
	file_chat_v1_chat_proto_rawDesc = nil
	file_chat_v1_chat_proto_goTypes = nil
	file_chat_v1_chat_proto_depIdxs = nil
}

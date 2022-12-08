// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/actions/sdk/v2/validation_results.proto

package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Wrapper for repeated validation result.
type ValidationResults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Multiple validation results.
	Results []*ValidationResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ValidationResults) Reset() {
	*x = ValidationResults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_validation_results_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResults) ProtoMessage() {}

func (x *ValidationResults) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_validation_results_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResults.ProtoReflect.Descriptor instead.
func (*ValidationResults) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_validation_results_proto_rawDescGZIP(), []int{0}
}

func (x *ValidationResults) GetResults() []*ValidationResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// Represents a validation result associated with the app content.
type ValidationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Holds the validation message.
	ValidationMessage string `protobuf:"bytes,1,opt,name=validation_message,json=validationMessage,proto3" json:"validation_message,omitempty"`
	// Context to identify the resource the validation message relates to.
	ValidationContext *ValidationResult_ValidationContext `protobuf:"bytes,2,opt,name=validation_context,json=validationContext,proto3" json:"validation_context,omitempty"`
}

func (x *ValidationResult) Reset() {
	*x = ValidationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_validation_results_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResult) ProtoMessage() {}

func (x *ValidationResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_validation_results_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResult.ProtoReflect.Descriptor instead.
func (*ValidationResult) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_validation_results_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationResult) GetValidationMessage() string {
	if x != nil {
		return x.ValidationMessage
	}
	return ""
}

func (x *ValidationResult) GetValidationContext() *ValidationResult_ValidationContext {
	if x != nil {
		return x.ValidationContext
	}
	return nil
}

// Context to identify the resource the validation message relates to.
type ValidationResult_ValidationContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Language code of the lozalized resource.
	// Empty if the error is for non-localized resource.
	// See the list of supported languages in
	// https://developers.google.com/assistant/console/languages-locales
	LanguageCode string `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
}

func (x *ValidationResult_ValidationContext) Reset() {
	*x = ValidationResult_ValidationContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_validation_results_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResult_ValidationContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResult_ValidationContext) ProtoMessage() {}

func (x *ValidationResult_ValidationContext) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_validation_results_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResult_ValidationContext.ProtoReflect.Descriptor instead.
func (*ValidationResult_ValidationContext) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_validation_results_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ValidationResult_ValidationContext) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

var File_google_actions_sdk_v2_validation_results_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_validation_results_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x22, 0x56, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73,
	0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22,
	0xe5, 0x01, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x68, 0x0a, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x11, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x38, 0x0a,
	0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x6f, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x42, 0x16, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x64, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_validation_results_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_validation_results_proto_rawDescData = file_google_actions_sdk_v2_validation_results_proto_rawDesc
)

func file_google_actions_sdk_v2_validation_results_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_validation_results_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_validation_results_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_validation_results_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_validation_results_proto_rawDescData
}

var file_google_actions_sdk_v2_validation_results_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_actions_sdk_v2_validation_results_proto_goTypes = []interface{}{
	(*ValidationResults)(nil),                  // 0: google.actions.sdk.v2.ValidationResults
	(*ValidationResult)(nil),                   // 1: google.actions.sdk.v2.ValidationResult
	(*ValidationResult_ValidationContext)(nil), // 2: google.actions.sdk.v2.ValidationResult.ValidationContext
}
var file_google_actions_sdk_v2_validation_results_proto_depIdxs = []int32{
	1, // 0: google.actions.sdk.v2.ValidationResults.results:type_name -> google.actions.sdk.v2.ValidationResult
	2, // 1: google.actions.sdk.v2.ValidationResult.validation_context:type_name -> google.actions.sdk.v2.ValidationResult.ValidationContext
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_validation_results_proto_init() }
func file_google_actions_sdk_v2_validation_results_proto_init() {
	if File_google_actions_sdk_v2_validation_results_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_validation_results_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResults); i {
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
		file_google_actions_sdk_v2_validation_results_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResult); i {
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
		file_google_actions_sdk_v2_validation_results_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidationResult_ValidationContext); i {
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
			RawDescriptor: file_google_actions_sdk_v2_validation_results_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_validation_results_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_validation_results_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_validation_results_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_validation_results_proto = out.File
	file_google_actions_sdk_v2_validation_results_proto_rawDesc = nil
	file_google_actions_sdk_v2_validation_results_proto_goTypes = nil
	file_google_actions_sdk_v2_validation_results_proto_depIdxs = nil
}
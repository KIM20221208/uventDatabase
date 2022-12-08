// Copyright 2021 Google LLC
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
// source: google/cloud/retail/v2beta/completion_service.proto

package retail

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Auto-complete parameters.
type CompleteQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Catalog for which the completion is performed.
	//
	// Full resource name of catalog, such as
	// `projects/*/locations/global/catalogs/default_catalog`.
	Catalog string `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
	// Required. The query used to generate suggestions.
	//
	// The maximum number of allowed characters is 255.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// Required field. A unique identifier for tracking visitors. For example,
	// this could be implemented with an HTTP cookie, which should be able to
	// uniquely identify a visitor on a single device. This unique identifier
	// should not change if the visitor logs in or out of the website.
	//
	// The field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	VisitorId string `protobuf:"bytes,7,opt,name=visitor_id,json=visitorId,proto3" json:"visitor_id,omitempty"`
	// Note that this field applies for `user-data` dataset only. For requests
	// with `cloud-retail` dataset, setting this field has no effect.
	//
	// The language filters applied to the output suggestions. If set, it should
	// contain the language of the query. If not set, suggestions are returned
	// without considering language restrictions. This is the BCP-47 language
	// code, such as "en-US" or "sr-Latn". For more information, see [Tags for
	// Identifying Languages](https://tools.ietf.org/html/bcp47). The maximum
	// number of language codes is 3.
	LanguageCodes []string `protobuf:"bytes,3,rep,name=language_codes,json=languageCodes,proto3" json:"language_codes,omitempty"`
	// The device type context for completion suggestions.
	// It is useful to apply different suggestions on different device types, e.g.
	// `DESKTOP`, `MOBILE`. If it is empty, the suggestions are across all device
	// types.
	//
	// Supported formats:
	//
	// * `UNKNOWN_DEVICE_TYPE`
	//
	// * `DESKTOP`
	//
	// * `MOBILE`
	//
	// * A customized string starts with `OTHER_`, e.g. `OTHER_IPHONE`.
	DeviceType string `protobuf:"bytes,4,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// Determines which dataset to use for fetching completion. "user-data" will
	// use the imported dataset through
	// [CompletionService.ImportCompletionData][google.cloud.retail.v2beta.CompletionService.ImportCompletionData].
	// "cloud-retail" will use the dataset generated by cloud retail based on user
	// events. If leave empty, it will use the "user-data".
	//
	// Current supported values:
	//
	// * user-data
	//
	// * cloud-retail:
	//   This option requires enabling auto-learning function first. See
	//   [guidelines](https://cloud.google.com/retail/docs/completion-overview#generated-completion-dataset).
	Dataset string `protobuf:"bytes,6,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Completion max suggestions. If left unset or set to 0, then will fallback
	// to the configured value
	// [CompletionConfig.max_suggestions][google.cloud.retail.v2beta.CompletionConfig.max_suggestions].
	//
	// The maximum allowed max suggestions is 20. If it is set higher, it will be
	// capped by 20.
	MaxSuggestions int32 `protobuf:"varint,5,opt,name=max_suggestions,json=maxSuggestions,proto3" json:"max_suggestions,omitempty"`
}

func (x *CompleteQueryRequest) Reset() {
	*x = CompleteQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryRequest) ProtoMessage() {}

func (x *CompleteQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryRequest.ProtoReflect.Descriptor instead.
func (*CompleteQueryRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_completion_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteQueryRequest) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *CompleteQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *CompleteQueryRequest) GetVisitorId() string {
	if x != nil {
		return x.VisitorId
	}
	return ""
}

func (x *CompleteQueryRequest) GetLanguageCodes() []string {
	if x != nil {
		return x.LanguageCodes
	}
	return nil
}

func (x *CompleteQueryRequest) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *CompleteQueryRequest) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *CompleteQueryRequest) GetMaxSuggestions() int32 {
	if x != nil {
		return x.MaxSuggestions
	}
	return 0
}

// Response of the auto-complete query.
type CompleteQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Results of the matching suggestions. The result list is ordered and the
	// first result is top suggestion.
	CompletionResults []*CompleteQueryResponse_CompletionResult `protobuf:"bytes,1,rep,name=completion_results,json=completionResults,proto3" json:"completion_results,omitempty"`
	// A unique complete token. This should be included in the
	// [UserEvent.completion_detail][google.cloud.retail.v2beta.UserEvent.completion_detail]
	// for search events resulting from this completion, which enables accurate
	// attribution of complete model performance.
	AttributionToken string `protobuf:"bytes,2,opt,name=attribution_token,json=attributionToken,proto3" json:"attribution_token,omitempty"`
	// Matched recent searches of this user. The maximum number of recent searches
	// is 10. This field is a restricted feature. Contact Retail Search support
	// team if you are interested in enabling it.
	//
	// This feature is only available when
	// [CompleteQueryRequest.visitor_id][google.cloud.retail.v2beta.CompleteQueryRequest.visitor_id]
	// field is set and [UserEvent][google.cloud.retail.v2beta.UserEvent] is
	// imported. The recent searches satisfy the follow rules:
	//
	//  * They are ordered from latest to oldest.
	//
	//  * They are matched with
	//  [CompleteQueryRequest.query][google.cloud.retail.v2beta.CompleteQueryRequest.query]
	//  case insensitively.
	//
	//  * They are transformed to lower case.
	//
	//  * They are UTF-8 safe.
	//
	// Recent searches are deduplicated. More recent searches will be reserved
	// when duplication happens.
	RecentSearchResults []*CompleteQueryResponse_RecentSearchResult `protobuf:"bytes,3,rep,name=recent_search_results,json=recentSearchResults,proto3" json:"recent_search_results,omitempty"`
}

func (x *CompleteQueryResponse) Reset() {
	*x = CompleteQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse) ProtoMessage() {}

func (x *CompleteQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_completion_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteQueryResponse) GetCompletionResults() []*CompleteQueryResponse_CompletionResult {
	if x != nil {
		return x.CompletionResults
	}
	return nil
}

func (x *CompleteQueryResponse) GetAttributionToken() string {
	if x != nil {
		return x.AttributionToken
	}
	return ""
}

func (x *CompleteQueryResponse) GetRecentSearchResults() []*CompleteQueryResponse_RecentSearchResult {
	if x != nil {
		return x.RecentSearchResults
	}
	return nil
}

// Resource that represents completion results.
type CompleteQueryResponse_CompletionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The suggestion for the query.
	Suggestion string `protobuf:"bytes,1,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	// Custom attributes for the suggestion term.
	//
	// * For "user-data", the attributes are additional custom attributes
	// ingested through BigQuery.
	//
	// * For "cloud-retail", the attributes are product attributes generated
	// by Cloud Retail. It requires
	// [UserEvent.product_details][google.cloud.retail.v2beta.UserEvent.product_details]
	// is imported properly.
	Attributes map[string]*CustomAttribute `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CompleteQueryResponse_CompletionResult) Reset() {
	*x = CompleteQueryResponse_CompletionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse_CompletionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse_CompletionResult) ProtoMessage() {}

func (x *CompleteQueryResponse_CompletionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse_CompletionResult.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse_CompletionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_completion_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CompleteQueryResponse_CompletionResult) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

func (x *CompleteQueryResponse_CompletionResult) GetAttributes() map[string]*CustomAttribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// Recent search of this user.
type CompleteQueryResponse_RecentSearchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The recent search query.
	RecentSearch string `protobuf:"bytes,1,opt,name=recent_search,json=recentSearch,proto3" json:"recent_search,omitempty"`
}

func (x *CompleteQueryResponse_RecentSearchResult) Reset() {
	*x = CompleteQueryResponse_RecentSearchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse_RecentSearchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse_RecentSearchResult) ProtoMessage() {}

func (x *CompleteQueryResponse_RecentSearchResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse_RecentSearchResult.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse_RecentSearchResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2beta_completion_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CompleteQueryResponse_RecentSearchResult) GetRecentSearch() string {
	if x != nil {
		return x.RecentSearch
	}
	return ""
}

var File_google_cloud_retail_v2beta_completion_service_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2beta_completion_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x07, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xe0, 0x41,
	0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x19, 0x0a, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x78, 0x5f,
	0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x81, 0x05, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x11, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x2b,
	0x0a, 0x11, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x78, 0x0a, 0x15, 0x72,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x92, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x52,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x6a,
	0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x12, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x32, 0xce, 0x04, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x30, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x49, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x43, 0x12, 0x41, 0x2f, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x3a,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0xab, 0x02,
	0x0a, 0x14, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xba,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x22, 0x48, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x3a, 0x01, 0x2a, 0xca, 0x41, 0x64, 0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x49, 0xca, 0x41, 0x15,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f,
	0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xdf, 0x01, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x42, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02,
	0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x1a, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2beta_completion_service_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2beta_completion_service_proto_rawDescData = file_google_cloud_retail_v2beta_completion_service_proto_rawDesc
)

func file_google_cloud_retail_v2beta_completion_service_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2beta_completion_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2beta_completion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2beta_completion_service_proto_rawDescData)
	})
	return file_google_cloud_retail_v2beta_completion_service_proto_rawDescData
}

var file_google_cloud_retail_v2beta_completion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_retail_v2beta_completion_service_proto_goTypes = []interface{}{
	(*CompleteQueryRequest)(nil),                     // 0: google.cloud.retail.v2beta.CompleteQueryRequest
	(*CompleteQueryResponse)(nil),                    // 1: google.cloud.retail.v2beta.CompleteQueryResponse
	(*CompleteQueryResponse_CompletionResult)(nil),   // 2: google.cloud.retail.v2beta.CompleteQueryResponse.CompletionResult
	(*CompleteQueryResponse_RecentSearchResult)(nil), // 3: google.cloud.retail.v2beta.CompleteQueryResponse.RecentSearchResult
	nil,                                 // 4: google.cloud.retail.v2beta.CompleteQueryResponse.CompletionResult.AttributesEntry
	(*CustomAttribute)(nil),             // 5: google.cloud.retail.v2beta.CustomAttribute
	(*ImportCompletionDataRequest)(nil), // 6: google.cloud.retail.v2beta.ImportCompletionDataRequest
	(*longrunning.Operation)(nil),       // 7: google.longrunning.Operation
}
var file_google_cloud_retail_v2beta_completion_service_proto_depIdxs = []int32{
	2, // 0: google.cloud.retail.v2beta.CompleteQueryResponse.completion_results:type_name -> google.cloud.retail.v2beta.CompleteQueryResponse.CompletionResult
	3, // 1: google.cloud.retail.v2beta.CompleteQueryResponse.recent_search_results:type_name -> google.cloud.retail.v2beta.CompleteQueryResponse.RecentSearchResult
	4, // 2: google.cloud.retail.v2beta.CompleteQueryResponse.CompletionResult.attributes:type_name -> google.cloud.retail.v2beta.CompleteQueryResponse.CompletionResult.AttributesEntry
	5, // 3: google.cloud.retail.v2beta.CompleteQueryResponse.CompletionResult.AttributesEntry.value:type_name -> google.cloud.retail.v2beta.CustomAttribute
	0, // 4: google.cloud.retail.v2beta.CompletionService.CompleteQuery:input_type -> google.cloud.retail.v2beta.CompleteQueryRequest
	6, // 5: google.cloud.retail.v2beta.CompletionService.ImportCompletionData:input_type -> google.cloud.retail.v2beta.ImportCompletionDataRequest
	1, // 6: google.cloud.retail.v2beta.CompletionService.CompleteQuery:output_type -> google.cloud.retail.v2beta.CompleteQueryResponse
	7, // 7: google.cloud.retail.v2beta.CompletionService.ImportCompletionData:output_type -> google.longrunning.Operation
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2beta_completion_service_proto_init() }
func file_google_cloud_retail_v2beta_completion_service_proto_init() {
	if File_google_cloud_retail_v2beta_completion_service_proto != nil {
		return
	}
	file_google_cloud_retail_v2beta_common_proto_init()
	file_google_cloud_retail_v2beta_import_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteQueryRequest); i {
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
		file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteQueryResponse); i {
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
		file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteQueryResponse_CompletionResult); i {
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
		file_google_cloud_retail_v2beta_completion_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteQueryResponse_RecentSearchResult); i {
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
			RawDescriptor: file_google_cloud_retail_v2beta_completion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_retail_v2beta_completion_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2beta_completion_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2beta_completion_service_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2beta_completion_service_proto = out.File
	file_google_cloud_retail_v2beta_completion_service_proto_rawDesc = nil
	file_google_cloud_retail_v2beta_completion_service_proto_goTypes = nil
	file_google_cloud_retail_v2beta_completion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompletionServiceClient is the client API for CompletionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompletionServiceClient interface {
	// Completes the specified prefix with keyword suggestions.
	//
	// This feature is only available for users who have Retail Search enabled.
	// Please enable Retail Search on Cloud Console before using this feature.
	CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error)
	// Bulk import of processed completion dataset.
	//
	// Request processing is asynchronous. Partial updating is not supported.
	//
	// The operation is successfully finished only after the imported suggestions
	// are indexed successfully and ready for serving. The process takes hours.
	//
	// This feature is only available for users who have Retail Search enabled.
	// Please enable Retail Search on Cloud Console before using this feature.
	ImportCompletionData(ctx context.Context, in *ImportCompletionDataRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type completionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletionServiceClient(cc grpc.ClientConnInterface) CompletionServiceClient {
	return &completionServiceClient{cc}
}

func (c *completionServiceClient) CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error) {
	out := new(CompleteQueryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2beta.CompletionService/CompleteQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *completionServiceClient) ImportCompletionData(ctx context.Context, in *ImportCompletionDataRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2beta.CompletionService/ImportCompletionData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompletionServiceServer is the server API for CompletionService service.
type CompletionServiceServer interface {
	// Completes the specified prefix with keyword suggestions.
	//
	// This feature is only available for users who have Retail Search enabled.
	// Please enable Retail Search on Cloud Console before using this feature.
	CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error)
	// Bulk import of processed completion dataset.
	//
	// Request processing is asynchronous. Partial updating is not supported.
	//
	// The operation is successfully finished only after the imported suggestions
	// are indexed successfully and ready for serving. The process takes hours.
	//
	// This feature is only available for users who have Retail Search enabled.
	// Please enable Retail Search on Cloud Console before using this feature.
	ImportCompletionData(context.Context, *ImportCompletionDataRequest) (*longrunning.Operation, error)
}

// UnimplementedCompletionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompletionServiceServer struct {
}

func (*UnimplementedCompletionServiceServer) CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteQuery not implemented")
}
func (*UnimplementedCompletionServiceServer) ImportCompletionData(context.Context, *ImportCompletionDataRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportCompletionData not implemented")
}

func RegisterCompletionServiceServer(s *grpc.Server, srv CompletionServiceServer) {
	s.RegisterService(&_CompletionService_serviceDesc, srv)
}

func _CompletionService_CompleteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).CompleteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2beta.CompletionService/CompleteQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).CompleteQuery(ctx, req.(*CompleteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompletionService_ImportCompletionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportCompletionDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).ImportCompletionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2beta.CompletionService/ImportCompletionData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).ImportCompletionData(ctx, req.(*ImportCompletionDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompletionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.retail.v2beta.CompletionService",
	HandlerType: (*CompletionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteQuery",
			Handler:    _CompletionService_CompleteQuery_Handler,
		},
		{
			MethodName: "ImportCompletionData",
			Handler:    _CompletionService_ImportCompletionData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/retail/v2beta/completion_service.proto",
}

// Copyright 2025 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: teleport/workloadidentity/v1/revocation_service.proto

package workloadidentityv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request for CreateWorkloadIdentityX509Revocation.
type CreateWorkloadIdentityX509RevocationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The X509 workload identity revocation to create.
	WorkloadIdentityX509Revocation *WorkloadIdentityX509Revocation `protobuf:"bytes,1,opt,name=workload_identity_x509_revocation,json=workloadIdentityX509Revocation,proto3" json:"workload_identity_x509_revocation,omitempty"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *CreateWorkloadIdentityX509RevocationRequest) Reset() {
	*x = CreateWorkloadIdentityX509RevocationRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWorkloadIdentityX509RevocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkloadIdentityX509RevocationRequest) ProtoMessage() {}

func (x *CreateWorkloadIdentityX509RevocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkloadIdentityX509RevocationRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkloadIdentityX509RevocationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWorkloadIdentityX509RevocationRequest) GetWorkloadIdentityX509Revocation() *WorkloadIdentityX509Revocation {
	if x != nil {
		return x.WorkloadIdentityX509Revocation
	}
	return nil
}

// The request for UpdateWorkloadIdentityX509Revocation.
type UpdateWorkloadIdentityX509RevocationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The X509 workload identity revocation to update.
	WorkloadIdentityX509Revocation *WorkloadIdentityX509Revocation `protobuf:"bytes,1,opt,name=workload_identity_x509_revocation,json=workloadIdentityX509Revocation,proto3" json:"workload_identity_x509_revocation,omitempty"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *UpdateWorkloadIdentityX509RevocationRequest) Reset() {
	*x = UpdateWorkloadIdentityX509RevocationRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateWorkloadIdentityX509RevocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWorkloadIdentityX509RevocationRequest) ProtoMessage() {}

func (x *UpdateWorkloadIdentityX509RevocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWorkloadIdentityX509RevocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateWorkloadIdentityX509RevocationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateWorkloadIdentityX509RevocationRequest) GetWorkloadIdentityX509Revocation() *WorkloadIdentityX509Revocation {
	if x != nil {
		return x.WorkloadIdentityX509Revocation
	}
	return nil
}

// The request for UpsertWorkloadIdentityX509Revocation.
type UpsertWorkloadIdentityX509RevocationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The X509 workload identity revocation to upsert.
	WorkloadIdentityX509Revocation *WorkloadIdentityX509Revocation `protobuf:"bytes,1,opt,name=workload_identity_x509_revocation,json=workloadIdentityX509Revocation,proto3" json:"workload_identity_x509_revocation,omitempty"`
	unknownFields                  protoimpl.UnknownFields
	sizeCache                      protoimpl.SizeCache
}

func (x *UpsertWorkloadIdentityX509RevocationRequest) Reset() {
	*x = UpsertWorkloadIdentityX509RevocationRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertWorkloadIdentityX509RevocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertWorkloadIdentityX509RevocationRequest) ProtoMessage() {}

func (x *UpsertWorkloadIdentityX509RevocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertWorkloadIdentityX509RevocationRequest.ProtoReflect.Descriptor instead.
func (*UpsertWorkloadIdentityX509RevocationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertWorkloadIdentityX509RevocationRequest) GetWorkloadIdentityX509Revocation() *WorkloadIdentityX509Revocation {
	if x != nil {
		return x.WorkloadIdentityX509Revocation
	}
	return nil
}

// The request for GetWorkloadIdentityX509Revocation.
type GetWorkloadIdentityX509RevocationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the X509 workload identity revocation to retrieve.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorkloadIdentityX509RevocationRequest) Reset() {
	*x = GetWorkloadIdentityX509RevocationRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkloadIdentityX509RevocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkloadIdentityX509RevocationRequest) ProtoMessage() {}

func (x *GetWorkloadIdentityX509RevocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkloadIdentityX509RevocationRequest.ProtoReflect.Descriptor instead.
func (*GetWorkloadIdentityX509RevocationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetWorkloadIdentityX509RevocationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request for DeleteWorkloadIdentityX509Revocation.
type DeleteWorkloadIdentityX509RevocationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the workload identity to delete.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteWorkloadIdentityX509RevocationRequest) Reset() {
	*x = DeleteWorkloadIdentityX509RevocationRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteWorkloadIdentityX509RevocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkloadIdentityX509RevocationRequest) ProtoMessage() {}

func (x *DeleteWorkloadIdentityX509RevocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkloadIdentityX509RevocationRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkloadIdentityX509RevocationRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteWorkloadIdentityX509RevocationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request for ListWorkloadIdentityX509Revocations.
type ListWorkloadIdentityX509RevocationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	// The server may impose a different page size at its discretion.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page_token value returned from a previous ListWorkloadIdentities request, if any.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListWorkloadIdentityX509RevocationsRequest) Reset() {
	*x = ListWorkloadIdentityX509RevocationsRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWorkloadIdentityX509RevocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkloadIdentityX509RevocationsRequest) ProtoMessage() {}

func (x *ListWorkloadIdentityX509RevocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkloadIdentityX509RevocationsRequest.ProtoReflect.Descriptor instead.
func (*ListWorkloadIdentityX509RevocationsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListWorkloadIdentityX509RevocationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListWorkloadIdentityX509RevocationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response for ListWorkloadIdentityX509Revocations.
type ListWorkloadIdentityX509RevocationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The page of workload identities that matched the request.
	WorkloadIdentityX509Revocations []*WorkloadIdentityX509Revocation `protobuf:"bytes,1,rep,name=workload_identity_x509_revocations,json=workloadIdentityX509Revocations,proto3" json:"workload_identity_x509_revocations,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListWorkloadIdentityX509RevocationsResponse) Reset() {
	*x = ListWorkloadIdentityX509RevocationsResponse{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListWorkloadIdentityX509RevocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWorkloadIdentityX509RevocationsResponse) ProtoMessage() {}

func (x *ListWorkloadIdentityX509RevocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWorkloadIdentityX509RevocationsResponse.ProtoReflect.Descriptor instead.
func (*ListWorkloadIdentityX509RevocationsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListWorkloadIdentityX509RevocationsResponse) GetWorkloadIdentityX509Revocations() []*WorkloadIdentityX509Revocation {
	if x != nil {
		return x.WorkloadIdentityX509Revocations
	}
	return nil
}

func (x *ListWorkloadIdentityX509RevocationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The request for StreamSignedCRL.
type StreamSignedCRLRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamSignedCRLRequest) Reset() {
	*x = StreamSignedCRLRequest{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamSignedCRLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSignedCRLRequest) ProtoMessage() {}

func (x *StreamSignedCRLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSignedCRLRequest.ProtoReflect.Descriptor instead.
func (*StreamSignedCRLRequest) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{7}
}

// The response for StreamSignedCRL.
type StreamSignedCRLResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The signed Certificate Revocation List (CRL).
	//
	// The syntax of the CRL is defined at https://www.rfc-editor.org/rfc/rfc5280.html#section-5
	// This field is encoded in DER ASN.1 without any PEM wrapping.
	//
	// When a new signed CRL is available, the full new CRL will be sent to the
	// client again using this field.
	Crl           []byte `protobuf:"bytes,1,opt,name=crl,proto3" json:"crl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamSignedCRLResponse) Reset() {
	*x = StreamSignedCRLResponse{}
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamSignedCRLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamSignedCRLResponse) ProtoMessage() {}

func (x *StreamSignedCRLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamSignedCRLResponse.ProtoReflect.Descriptor instead.
func (*StreamSignedCRLResponse) Descriptor() ([]byte, []int) {
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP(), []int{8}
}

func (x *StreamSignedCRLResponse) GetCrl() []byte {
	if x != nil {
		return x.Crl
	}
	return nil
}

var File_teleport_workloadidentity_v1_revocation_service_proto protoreflect.FileDescriptor

const file_teleport_workloadidentity_v1_revocation_service_proto_rawDesc = "" +
	"\n" +
	"5teleport/workloadidentity/v1/revocation_service.proto\x12\x1cteleport.workloadidentity.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a6teleport/workloadidentity/v1/revocation_resource.proto\"\xb7\x01\n" +
	"+CreateWorkloadIdentityX509RevocationRequest\x12\x87\x01\n" +
	"!workload_identity_x509_revocation\x18\x01 \x01(\v2<.teleport.workloadidentity.v1.WorkloadIdentityX509RevocationR\x1eworkloadIdentityX509Revocation\"\xb7\x01\n" +
	"+UpdateWorkloadIdentityX509RevocationRequest\x12\x87\x01\n" +
	"!workload_identity_x509_revocation\x18\x01 \x01(\v2<.teleport.workloadidentity.v1.WorkloadIdentityX509RevocationR\x1eworkloadIdentityX509Revocation\"\xb7\x01\n" +
	"+UpsertWorkloadIdentityX509RevocationRequest\x12\x87\x01\n" +
	"!workload_identity_x509_revocation\x18\x01 \x01(\v2<.teleport.workloadidentity.v1.WorkloadIdentityX509RevocationR\x1eworkloadIdentityX509Revocation\">\n" +
	"(GetWorkloadIdentityX509RevocationRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"A\n" +
	"+DeleteWorkloadIdentityX509RevocationRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"h\n" +
	"*ListWorkloadIdentityX509RevocationsRequest\x12\x1b\n" +
	"\tpage_size\x18\x01 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tR\tpageToken\"\xe1\x01\n" +
	"+ListWorkloadIdentityX509RevocationsResponse\x12\x89\x01\n" +
	"\"workload_identity_x509_revocations\x18\x01 \x03(\v2<.teleport.workloadidentity.v1.WorkloadIdentityX509RevocationR\x1fworkloadIdentityX509Revocations\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\x18\n" +
	"\x16StreamSignedCRLRequest\"+\n" +
	"\x17StreamSignedCRLResponse\x12\x10\n" +
	"\x03crl\x18\x01 \x01(\fR\x03crl2\xb1\t\n" +
	"!WorkloadIdentityRevocationService\x12\xaf\x01\n" +
	"$CreateWorkloadIdentityX509Revocation\x12I.teleport.workloadidentity.v1.CreateWorkloadIdentityX509RevocationRequest\x1a<.teleport.workloadidentity.v1.WorkloadIdentityX509Revocation\x12\xaf\x01\n" +
	"$UpsertWorkloadIdentityX509Revocation\x12I.teleport.workloadidentity.v1.UpsertWorkloadIdentityX509RevocationRequest\x1a<.teleport.workloadidentity.v1.WorkloadIdentityX509Revocation\x12\xaf\x01\n" +
	"$UpdateWorkloadIdentityX509Revocation\x12I.teleport.workloadidentity.v1.UpdateWorkloadIdentityX509RevocationRequest\x1a<.teleport.workloadidentity.v1.WorkloadIdentityX509Revocation\x12\xa9\x01\n" +
	"!GetWorkloadIdentityX509Revocation\x12F.teleport.workloadidentity.v1.GetWorkloadIdentityX509RevocationRequest\x1a<.teleport.workloadidentity.v1.WorkloadIdentityX509Revocation\x12\x89\x01\n" +
	"$DeleteWorkloadIdentityX509Revocation\x12I.teleport.workloadidentity.v1.DeleteWorkloadIdentityX509RevocationRequest\x1a\x16.google.protobuf.Empty\x12\xba\x01\n" +
	"#ListWorkloadIdentityX509Revocations\x12H.teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsRequest\x1aI.teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsResponse\x12\x80\x01\n" +
	"\x0fStreamSignedCRL\x124.teleport.workloadidentity.v1.StreamSignedCRLRequest\x1a5.teleport.workloadidentity.v1.StreamSignedCRLResponse0\x01BdZbgithub.com/gravitational/teleport/api/gen/proto/go/teleport/workloadidentity/v1;workloadidentityv1b\x06proto3"

var (
	file_teleport_workloadidentity_v1_revocation_service_proto_rawDescOnce sync.Once
	file_teleport_workloadidentity_v1_revocation_service_proto_rawDescData []byte
)

func file_teleport_workloadidentity_v1_revocation_service_proto_rawDescGZIP() []byte {
	file_teleport_workloadidentity_v1_revocation_service_proto_rawDescOnce.Do(func() {
		file_teleport_workloadidentity_v1_revocation_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_teleport_workloadidentity_v1_revocation_service_proto_rawDesc), len(file_teleport_workloadidentity_v1_revocation_service_proto_rawDesc)))
	})
	return file_teleport_workloadidentity_v1_revocation_service_proto_rawDescData
}

var file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_teleport_workloadidentity_v1_revocation_service_proto_goTypes = []any{
	(*CreateWorkloadIdentityX509RevocationRequest)(nil), // 0: teleport.workloadidentity.v1.CreateWorkloadIdentityX509RevocationRequest
	(*UpdateWorkloadIdentityX509RevocationRequest)(nil), // 1: teleport.workloadidentity.v1.UpdateWorkloadIdentityX509RevocationRequest
	(*UpsertWorkloadIdentityX509RevocationRequest)(nil), // 2: teleport.workloadidentity.v1.UpsertWorkloadIdentityX509RevocationRequest
	(*GetWorkloadIdentityX509RevocationRequest)(nil),    // 3: teleport.workloadidentity.v1.GetWorkloadIdentityX509RevocationRequest
	(*DeleteWorkloadIdentityX509RevocationRequest)(nil), // 4: teleport.workloadidentity.v1.DeleteWorkloadIdentityX509RevocationRequest
	(*ListWorkloadIdentityX509RevocationsRequest)(nil),  // 5: teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsRequest
	(*ListWorkloadIdentityX509RevocationsResponse)(nil), // 6: teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsResponse
	(*StreamSignedCRLRequest)(nil),                      // 7: teleport.workloadidentity.v1.StreamSignedCRLRequest
	(*StreamSignedCRLResponse)(nil),                     // 8: teleport.workloadidentity.v1.StreamSignedCRLResponse
	(*WorkloadIdentityX509Revocation)(nil),              // 9: teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	(*emptypb.Empty)(nil),                               // 10: google.protobuf.Empty
}
var file_teleport_workloadidentity_v1_revocation_service_proto_depIdxs = []int32{
	9,  // 0: teleport.workloadidentity.v1.CreateWorkloadIdentityX509RevocationRequest.workload_identity_x509_revocation:type_name -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	9,  // 1: teleport.workloadidentity.v1.UpdateWorkloadIdentityX509RevocationRequest.workload_identity_x509_revocation:type_name -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	9,  // 2: teleport.workloadidentity.v1.UpsertWorkloadIdentityX509RevocationRequest.workload_identity_x509_revocation:type_name -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	9,  // 3: teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsResponse.workload_identity_x509_revocations:type_name -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	0,  // 4: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.CreateWorkloadIdentityX509Revocation:input_type -> teleport.workloadidentity.v1.CreateWorkloadIdentityX509RevocationRequest
	2,  // 5: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.UpsertWorkloadIdentityX509Revocation:input_type -> teleport.workloadidentity.v1.UpsertWorkloadIdentityX509RevocationRequest
	1,  // 6: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.UpdateWorkloadIdentityX509Revocation:input_type -> teleport.workloadidentity.v1.UpdateWorkloadIdentityX509RevocationRequest
	3,  // 7: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.GetWorkloadIdentityX509Revocation:input_type -> teleport.workloadidentity.v1.GetWorkloadIdentityX509RevocationRequest
	4,  // 8: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.DeleteWorkloadIdentityX509Revocation:input_type -> teleport.workloadidentity.v1.DeleteWorkloadIdentityX509RevocationRequest
	5,  // 9: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.ListWorkloadIdentityX509Revocations:input_type -> teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsRequest
	7,  // 10: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.StreamSignedCRL:input_type -> teleport.workloadidentity.v1.StreamSignedCRLRequest
	9,  // 11: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.CreateWorkloadIdentityX509Revocation:output_type -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	9,  // 12: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.UpsertWorkloadIdentityX509Revocation:output_type -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	9,  // 13: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.UpdateWorkloadIdentityX509Revocation:output_type -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	9,  // 14: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.GetWorkloadIdentityX509Revocation:output_type -> teleport.workloadidentity.v1.WorkloadIdentityX509Revocation
	10, // 15: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.DeleteWorkloadIdentityX509Revocation:output_type -> google.protobuf.Empty
	6,  // 16: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.ListWorkloadIdentityX509Revocations:output_type -> teleport.workloadidentity.v1.ListWorkloadIdentityX509RevocationsResponse
	8,  // 17: teleport.workloadidentity.v1.WorkloadIdentityRevocationService.StreamSignedCRL:output_type -> teleport.workloadidentity.v1.StreamSignedCRLResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_teleport_workloadidentity_v1_revocation_service_proto_init() }
func file_teleport_workloadidentity_v1_revocation_service_proto_init() {
	if File_teleport_workloadidentity_v1_revocation_service_proto != nil {
		return
	}
	file_teleport_workloadidentity_v1_revocation_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_teleport_workloadidentity_v1_revocation_service_proto_rawDesc), len(file_teleport_workloadidentity_v1_revocation_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_workloadidentity_v1_revocation_service_proto_goTypes,
		DependencyIndexes: file_teleport_workloadidentity_v1_revocation_service_proto_depIdxs,
		MessageInfos:      file_teleport_workloadidentity_v1_revocation_service_proto_msgTypes,
	}.Build()
	File_teleport_workloadidentity_v1_revocation_service_proto = out.File
	file_teleport_workloadidentity_v1_revocation_service_proto_goTypes = nil
	file_teleport_workloadidentity_v1_revocation_service_proto_depIdxs = nil
}

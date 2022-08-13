//-
// ‌
// Hedera Mirror Node
// ​
// Copyright (C) 2019-2022 Hedera Hashgraph, LLC
// ​
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
// ‍

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: mirror/mirror_network_service.proto

package mirror

import (
	services "github.com/hashgraph/hedera-protobufs-go/services"
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

//*
// Request object to query an address book for its list of nodes
type AddressBookQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The ID of the address book file on the network. Can be either 0.0.101 or 0.0.102.
	FileId *services.FileID `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	//*
	// The maximum number of node addresses to receive before stopping. If not set or set to zero it will return all node addresses in the database.
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *AddressBookQuery) Reset() {
	*x = AddressBookQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mirror_mirror_network_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBookQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBookQuery) ProtoMessage() {}

func (x *AddressBookQuery) ProtoReflect() protoreflect.Message {
	mi := &file_mirror_mirror_network_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBookQuery.ProtoReflect.Descriptor instead.
func (*AddressBookQuery) Descriptor() ([]byte, []int) {
	return file_mirror_mirror_network_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddressBookQuery) GetFileId() *services.FileID {
	if x != nil {
		return x.FileId
	}
	return nil
}

func (x *AddressBookQuery) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_mirror_mirror_network_service_proto protoreflect.FileDescriptor

var file_mirror_mirror_network_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0x61, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x67, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x2e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x6f, 0x6f, 0x6b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x30, 0x01, 0x42, 0x1f, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x2e, 0x6d, 0x69, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mirror_mirror_network_service_proto_rawDescOnce sync.Once
	file_mirror_mirror_network_service_proto_rawDescData = file_mirror_mirror_network_service_proto_rawDesc
)

func file_mirror_mirror_network_service_proto_rawDescGZIP() []byte {
	file_mirror_mirror_network_service_proto_rawDescOnce.Do(func() {
		file_mirror_mirror_network_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_mirror_mirror_network_service_proto_rawDescData)
	})
	return file_mirror_mirror_network_service_proto_rawDescData
}

var file_mirror_mirror_network_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mirror_mirror_network_service_proto_goTypes = []interface{}{
	(*AddressBookQuery)(nil),     // 0: com.hedera.mirror.api.proto.AddressBookQuery
	(*services.FileID)(nil),      // 1: proto.FileID
	(*services.NodeAddress)(nil), // 2: proto.NodeAddress
}
var file_mirror_mirror_network_service_proto_depIdxs = []int32{
	1, // 0: com.hedera.mirror.api.proto.AddressBookQuery.file_id:type_name -> proto.FileID
	0, // 1: com.hedera.mirror.api.proto.NetworkService.getNodes:input_type -> com.hedera.mirror.api.proto.AddressBookQuery
	2, // 2: com.hedera.mirror.api.proto.NetworkService.getNodes:output_type -> proto.NodeAddress
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mirror_mirror_network_service_proto_init() }
func file_mirror_mirror_network_service_proto_init() {
	if File_mirror_mirror_network_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mirror_mirror_network_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBookQuery); i {
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
			RawDescriptor: file_mirror_mirror_network_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mirror_mirror_network_service_proto_goTypes,
		DependencyIndexes: file_mirror_mirror_network_service_proto_depIdxs,
		MessageInfos:      file_mirror_mirror_network_service_proto_msgTypes,
	}.Build()
	File_mirror_mirror_network_service_proto = out.File
	file_mirror_mirror_network_service_proto_rawDesc = nil
	file_mirror_mirror_network_service_proto_goTypes = nil
	file_mirror_mirror_network_service_proto_depIdxs = nil
}

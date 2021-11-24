// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: protos/items.proto

package protos

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

// message in ItemQuery
type QueryOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptName string `protobuf:"bytes,1,opt,name=optName,proto3" json:"optName,omitempty"`
	Upper   int32  `protobuf:"varint,2,opt,name=upper,proto3" json:"upper,omitempty"`
	Under   int32  `protobuf:"varint,3,opt,name=under,proto3" json:"under,omitempty"`
}

func (x *QueryOption) Reset() {
	*x = QueryOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOption) ProtoMessage() {}

func (x *QueryOption) ProtoReflect() protoreflect.Message {
	mi := &file_protos_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOption.ProtoReflect.Descriptor instead.
func (*QueryOption) Descriptor() ([]byte, []int) {
	return file_protos_items_proto_rawDescGZIP(), []int{0}
}

func (x *QueryOption) GetOptName() string {
	if x != nil {
		return x.OptName
	}
	return ""
}

func (x *QueryOption) GetUpper() int32 {
	if x != nil {
		return x.Upper
	}
	return 0
}

func (x *QueryOption) GetUnder() int32 {
	if x != nil {
		return x.Under
	}
	return 0
}

// request message
type ItemQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CostUpper int32          `protobuf:"varint,2,opt,name=costUpper,proto3" json:"costUpper,omitempty"`
	CostUnder int32          `protobuf:"varint,3,opt,name=costUnder,proto3" json:"costUnder,omitempty"`
	QueryOpt  []*QueryOption `protobuf:"bytes,4,rep,name=queryOpt,proto3" json:"queryOpt,omitempty"`
}

func (x *ItemQuery) Reset() {
	*x = ItemQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemQuery) ProtoMessage() {}

func (x *ItemQuery) ProtoReflect() protoreflect.Message {
	mi := &file_protos_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemQuery.ProtoReflect.Descriptor instead.
func (*ItemQuery) Descriptor() ([]byte, []int) {
	return file_protos_items_proto_rawDescGZIP(), []int{1}
}

func (x *ItemQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemQuery) GetCostUpper() int32 {
	if x != nil {
		return x.CostUpper
	}
	return 0
}

func (x *ItemQuery) GetCostUnder() int32 {
	if x != nil {
		return x.CostUnder
	}
	return 0
}

func (x *ItemQuery) GetQueryOpt() []*QueryOption {
	if x != nil {
		return x.QueryOpt
	}
	return nil
}

// message in ItemSpec
type ItemOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptName string `protobuf:"bytes,1,opt,name=optName,proto3" json:"optName,omitempty"`
	Value   int32  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ItemOption) Reset() {
	*x = ItemOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemOption) ProtoMessage() {}

func (x *ItemOption) ProtoReflect() protoreflect.Message {
	mi := &file_protos_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemOption.ProtoReflect.Descriptor instead.
func (*ItemOption) Descriptor() ([]byte, []int) {
	return file_protos_items_proto_rawDescGZIP(), []int{2}
}

func (x *ItemOption) GetOptName() string {
	if x != nil {
		return x.OptName
	}
	return ""
}

func (x *ItemOption) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// response message
type ItemSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cost    int32         `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	ItemOpt []*ItemOption `protobuf:"bytes,4,rep,name=itemOpt,proto3" json:"itemOpt,omitempty"`
}

func (x *ItemSpec) Reset() {
	*x = ItemSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemSpec) ProtoMessage() {}

func (x *ItemSpec) ProtoReflect() protoreflect.Message {
	mi := &file_protos_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemSpec.ProtoReflect.Descriptor instead.
func (*ItemSpec) Descriptor() ([]byte, []int) {
	return file_protos_items_proto_rawDescGZIP(), []int{3}
}

func (x *ItemSpec) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemSpec) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *ItemSpec) GetItemOpt() []*ItemOption {
	if x != nil {
		return x.ItemOpt
	}
	return nil
}

type ItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList []*ItemSpec `protobuf:"bytes,1,rep,name=itemList,proto3" json:"itemList,omitempty"`
}

func (x *ItemList) Reset() {
	*x = ItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemList) ProtoMessage() {}

func (x *ItemList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemList.ProtoReflect.Descriptor instead.
func (*ItemList) Descriptor() ([]byte, []int) {
	return file_protos_items_proto_rawDescGZIP(), []int{4}
}

func (x *ItemList) GetItemList() []*ItemSpec {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_protos_items_proto protoreflect.FileDescriptor

var file_protos_items_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x53, 0x0a, 0x0b, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x70, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x70, 0x70, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x6e, 0x64, 0x65, 0x72,
	0x22, 0x8b, 0x01, 0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x22, 0x3c,
	0x0a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x6f, 0x0a, 0x08,
	0x49, 0x74, 0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x4f, 0x70, 0x74, 0x22, 0x37, 0x0a,
	0x08, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x69, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x36, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x2d, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x10, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x77, 0x61,
	0x6e, 0x6b, 0x69, 0x6d, 0x31, 0x32, 0x33, 0x2f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2d, 0x67, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_items_proto_rawDescOnce sync.Once
	file_protos_items_proto_rawDescData = file_protos_items_proto_rawDesc
)

func file_protos_items_proto_rawDescGZIP() []byte {
	file_protos_items_proto_rawDescOnce.Do(func() {
		file_protos_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_items_proto_rawDescData)
	})
	return file_protos_items_proto_rawDescData
}

var file_protos_items_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_items_proto_goTypes = []interface{}{
	(*QueryOption)(nil), // 0: items.QueryOption
	(*ItemQuery)(nil),   // 1: items.ItemQuery
	(*ItemOption)(nil),  // 2: items.ItemOption
	(*ItemSpec)(nil),    // 3: items.ItemSpec
	(*ItemList)(nil),    // 4: items.ItemList
}
var file_protos_items_proto_depIdxs = []int32{
	0, // 0: items.ItemQuery.queryOpt:type_name -> items.QueryOption
	2, // 1: items.ItemSpec.itemOpt:type_name -> items.ItemOption
	3, // 2: items.ItemList.itemList:type_name -> items.ItemSpec
	1, // 3: items.Items.GetAll:input_type -> items.ItemQuery
	4, // 4: items.Items.GetAll:output_type -> items.ItemList
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_items_proto_init() }
func file_protos_items_proto_init() {
	if File_protos_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOption); i {
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
		file_protos_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemQuery); i {
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
		file_protos_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemOption); i {
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
		file_protos_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemSpec); i {
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
		file_protos_items_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemList); i {
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
			RawDescriptor: file_protos_items_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_items_proto_goTypes,
		DependencyIndexes: file_protos_items_proto_depIdxs,
		MessageInfos:      file_protos_items_proto_msgTypes,
	}.Build()
	File_protos_items_proto = out.File
	file_protos_items_proto_rawDesc = nil
	file_protos_items_proto_goTypes = nil
	file_protos_items_proto_depIdxs = nil
}

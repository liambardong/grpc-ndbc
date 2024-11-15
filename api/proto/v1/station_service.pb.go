// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: station_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListStationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations []*Station `protobuf:"bytes,1,rep,name=stations,proto3" json:"stations,omitempty"`
}

func (x *ListStationsResponse) Reset() {
	*x = ListStationsResponse{}
	mi := &file_station_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStationsResponse) ProtoMessage() {}

func (x *ListStationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStationsResponse.ProtoReflect.Descriptor instead.
func (*ListStationsResponse) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListStationsResponse) GetStations() []*Station {
	if x != nil {
		return x.Stations
	}
	return nil
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StationId string  `protobuf:"bytes,1,opt,name=stationId,proto3" json:"stationId,omitempty"`
	Owner     string  `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Type      string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Hull      string  `protobuf:"bytes,4,opt,name=hull,proto3" json:"hull,omitempty"`
	Name      string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Payload   string  `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Latitude  float64 `protobuf:"fixed64,7,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Timezone  string  `protobuf:"bytes,10,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Forecast  string  `protobuf:"bytes,11,opt,name=forecast,proto3" json:"forecast,omitempty"`
	Note      string  `protobuf:"bytes,12,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	mi := &file_station_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_station_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_station_service_proto_rawDescGZIP(), []int{1}
}

func (x *Station) GetStationId() string {
	if x != nil {
		return x.StationId
	}
	return ""
}

func (x *Station) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Station) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Station) GetHull() string {
	if x != nil {
		return x.Hull
	}
	return ""
}

func (x *Station) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Station) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Station) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Station) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Station) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Station) GetForecast() string {
	if x != nil {
		return x.Forecast
	}
	return ""
}

func (x *Station) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

var File_station_service_proto protoreflect.FileDescriptor

var file_station_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x99, 0x02, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x75, 0x6c, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x75, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x32, 0x58, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x46, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x61, 0x6d, 0x62, 0x61, 0x72, 0x64, 0x6f, 0x6e,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6e, 0x64, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_station_service_proto_rawDescOnce sync.Once
	file_station_service_proto_rawDescData = file_station_service_proto_rawDesc
)

func file_station_service_proto_rawDescGZIP() []byte {
	file_station_service_proto_rawDescOnce.Do(func() {
		file_station_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_station_service_proto_rawDescData)
	})
	return file_station_service_proto_rawDescData
}

var file_station_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_station_service_proto_goTypes = []any{
	(*ListStationsResponse)(nil), // 0: api.v1.ListStationsResponse
	(*Station)(nil),              // 1: api.v1.Station
	(*emptypb.Empty)(nil),        // 2: google.protobuf.Empty
}
var file_station_service_proto_depIdxs = []int32{
	1, // 0: api.v1.ListStationsResponse.stations:type_name -> api.v1.Station
	2, // 1: api.v1.StationService.ListStations:input_type -> google.protobuf.Empty
	0, // 2: api.v1.StationService.ListStations:output_type -> api.v1.ListStationsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_station_service_proto_init() }
func file_station_service_proto_init() {
	if File_station_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_station_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_station_service_proto_goTypes,
		DependencyIndexes: file_station_service_proto_depIdxs,
		MessageInfos:      file_station_service_proto_msgTypes,
	}.Build()
	File_station_service_proto = out.File
	file_station_service_proto_rawDesc = nil
	file_station_service_proto_goTypes = nil
	file_station_service_proto_depIdxs = nil
}

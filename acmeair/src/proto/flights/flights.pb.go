// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/flights.proto

package flights

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	proto "microless/acmeair/proto"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTripFlightsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAirport  string                 `protobuf:"bytes,1,opt,name=from_airport,json=fromAirport,proto3" json:"from_airport,omitempty"`
	ToAirport    string                 `protobuf:"bytes,2,opt,name=to_airport,json=toAirport,proto3" json:"to_airport,omitempty"`
	FromDate     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=from_date,json=fromDate,proto3" json:"from_date,omitempty"`
	ReturnDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=return_date,json=returnDate,proto3" json:"return_date,omitempty"`
	OneWayFlight bool                   `protobuf:"varint,5,opt,name=one_way_flight,json=oneWayFlight,proto3" json:"one_way_flight,omitempty"`
}

func (x *GetTripFlightsRequest) Reset() {
	*x = GetTripFlightsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flights_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripFlightsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripFlightsRequest) ProtoMessage() {}

func (x *GetTripFlightsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flights_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripFlightsRequest.ProtoReflect.Descriptor instead.
func (*GetTripFlightsRequest) Descriptor() ([]byte, []int) {
	return file_proto_flights_proto_rawDescGZIP(), []int{0}
}

func (x *GetTripFlightsRequest) GetFromAirport() string {
	if x != nil {
		return x.FromAirport
	}
	return ""
}

func (x *GetTripFlightsRequest) GetToAirport() string {
	if x != nil {
		return x.ToAirport
	}
	return ""
}

func (x *GetTripFlightsRequest) GetFromDate() *timestamppb.Timestamp {
	if x != nil {
		return x.FromDate
	}
	return nil
}

func (x *GetTripFlightsRequest) GetReturnDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReturnDate
	}
	return nil
}

func (x *GetTripFlightsRequest) GetOneWayFlight() bool {
	if x != nil {
		return x.OneWayFlight
	}
	return false
}

type GetTripFlightsRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToFlights     []*proto.FlightInfo `protobuf:"bytes,1,rep,name=to_flights,json=toFlights,proto3" json:"to_flights,omitempty"`
	ReturnFlights []*proto.FlightInfo `protobuf:"bytes,2,rep,name=return_flights,json=returnFlights,proto3" json:"return_flights,omitempty"`
	OneWayFlight  bool                `protobuf:"varint,3,opt,name=one_way_flight,json=oneWayFlight,proto3" json:"one_way_flight,omitempty"`
}

func (x *GetTripFlightsRespond) Reset() {
	*x = GetTripFlightsRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flights_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTripFlightsRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTripFlightsRespond) ProtoMessage() {}

func (x *GetTripFlightsRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flights_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTripFlightsRespond.ProtoReflect.Descriptor instead.
func (*GetTripFlightsRespond) Descriptor() ([]byte, []int) {
	return file_proto_flights_proto_rawDescGZIP(), []int{1}
}

func (x *GetTripFlightsRespond) GetToFlights() []*proto.FlightInfo {
	if x != nil {
		return x.ToFlights
	}
	return nil
}

func (x *GetTripFlightsRespond) GetReturnFlights() []*proto.FlightInfo {
	if x != nil {
		return x.ReturnFlights
	}
	return nil
}

func (x *GetTripFlightsRespond) GetOneWayFlight() bool {
	if x != nil {
		return x.OneWayFlight
	}
	return false
}

type BrowseFlightsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAirport  string `protobuf:"bytes,1,opt,name=from_airport,json=fromAirport,proto3" json:"from_airport,omitempty"`
	ToAirport    string `protobuf:"bytes,2,opt,name=to_airport,json=toAirport,proto3" json:"to_airport,omitempty"`
	OneWayFlight bool   `protobuf:"varint,3,opt,name=one_way_flight,json=oneWayFlight,proto3" json:"one_way_flight,omitempty"`
}

func (x *BrowseFlightsRequest) Reset() {
	*x = BrowseFlightsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flights_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowseFlightsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowseFlightsRequest) ProtoMessage() {}

func (x *BrowseFlightsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flights_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowseFlightsRequest.ProtoReflect.Descriptor instead.
func (*BrowseFlightsRequest) Descriptor() ([]byte, []int) {
	return file_proto_flights_proto_rawDescGZIP(), []int{2}
}

func (x *BrowseFlightsRequest) GetFromAirport() string {
	if x != nil {
		return x.FromAirport
	}
	return ""
}

func (x *BrowseFlightsRequest) GetToAirport() string {
	if x != nil {
		return x.ToAirport
	}
	return ""
}

func (x *BrowseFlightsRequest) GetOneWayFlight() bool {
	if x != nil {
		return x.OneWayFlight
	}
	return false
}

type BrowseFlightsRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ToFlights     []*proto.FlightInfo `protobuf:"bytes,1,rep,name=to_flights,json=toFlights,proto3" json:"to_flights,omitempty"`
	ReturnFlights []*proto.FlightInfo `protobuf:"bytes,2,rep,name=return_flights,json=returnFlights,proto3" json:"return_flights,omitempty"`
	OneWayFlight  bool                `protobuf:"varint,3,opt,name=one_way_flight,json=oneWayFlight,proto3" json:"one_way_flight,omitempty"`
}

func (x *BrowseFlightsRespond) Reset() {
	*x = BrowseFlightsRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flights_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowseFlightsRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowseFlightsRespond) ProtoMessage() {}

func (x *BrowseFlightsRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flights_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowseFlightsRespond.ProtoReflect.Descriptor instead.
func (*BrowseFlightsRespond) Descriptor() ([]byte, []int) {
	return file_proto_flights_proto_rawDescGZIP(), []int{3}
}

func (x *BrowseFlightsRespond) GetToFlights() []*proto.FlightInfo {
	if x != nil {
		return x.ToFlights
	}
	return nil
}

func (x *BrowseFlightsRespond) GetReturnFlights() []*proto.FlightInfo {
	if x != nil {
		return x.ReturnFlights
	}
	return nil
}

func (x *BrowseFlightsRespond) GetOneWayFlight() bool {
	if x != nil {
		return x.OneWayFlight
	}
	return false
}

type GetFlightByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightId string `protobuf:"bytes,1,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
}

func (x *GetFlightByIdRequest) Reset() {
	*x = GetFlightByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flights_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightByIdRequest) ProtoMessage() {}

func (x *GetFlightByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flights_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightByIdRequest.ProtoReflect.Descriptor instead.
func (*GetFlightByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_flights_proto_rawDescGZIP(), []int{4}
}

func (x *GetFlightByIdRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

type GetFlightByIdRespond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flight *proto.FlightInfo `protobuf:"bytes,1,opt,name=flight,proto3" json:"flight,omitempty"`
}

func (x *GetFlightByIdRespond) Reset() {
	*x = GetFlightByIdRespond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_flights_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightByIdRespond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightByIdRespond) ProtoMessage() {}

func (x *GetFlightByIdRespond) ProtoReflect() protoreflect.Message {
	mi := &file_proto_flights_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightByIdRespond.ProtoReflect.Descriptor instead.
func (*GetFlightByIdRespond) Descriptor() ([]byte, []int) {
	return file_proto_flights_proto_rawDescGZIP(), []int{5}
}

func (x *GetFlightByIdRespond) GetFlight() *proto.FlightInfo {
	if x != nil {
		return x.Flight
	}
	return nil
}

var File_proto_flights_proto protoreflect.FileDescriptor

var file_proto_flights_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73,
	0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x0a,
	0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x5f, 0x77, 0x61, 0x79, 0x5f, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6f, 0x6e, 0x65,
	0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x69, 0x70, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x6f, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x12, 0x44, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x5f, 0x77,
	0x61, 0x79, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x7e, 0x0a,
	0x14, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x69,
	0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x72, 0x6f,
	0x6d, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x6e, 0x65, 0x5f, 0x77,
	0x61, 0x79, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0xc0, 0x01,
	0x0a, 0x14, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x6f, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72,
	0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x6e,
	0x65, 0x5f, 0x77, 0x61, 0x79, 0x5f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x6f, 0x6e, 0x65, 0x57, 0x61, 0x79, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x35, 0x0a,
	0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69,
	0x72, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x32, 0xc1, 0x03, 0x0a, 0x0e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9d, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x69, 0x70, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x30, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72,
	0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x22, 0x27,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9b, 0x01, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x66,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x22, 0x28, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x66, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x73, 0x2e, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2e, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x42, 0x21, 0x5a, 0x1f, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x63, 0x6d, 0x65, 0x61, 0x69, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_flights_proto_rawDescOnce sync.Once
	file_proto_flights_proto_rawDescData = file_proto_flights_proto_rawDesc
)

func file_proto_flights_proto_rawDescGZIP() []byte {
	file_proto_flights_proto_rawDescOnce.Do(func() {
		file_proto_flights_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_flights_proto_rawDescData)
	})
	return file_proto_flights_proto_rawDescData
}

var file_proto_flights_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_flights_proto_goTypes = []interface{}{
	(*GetTripFlightsRequest)(nil), // 0: microless.acmeair.flights.GetTripFlightsRequest
	(*GetTripFlightsRespond)(nil), // 1: microless.acmeair.flights.GetTripFlightsRespond
	(*BrowseFlightsRequest)(nil),  // 2: microless.acmeair.flights.BrowseFlightsRequest
	(*BrowseFlightsRespond)(nil),  // 3: microless.acmeair.flights.BrowseFlightsRespond
	(*GetFlightByIdRequest)(nil),  // 4: microless.acmeair.flights.GetFlightByIdRequest
	(*GetFlightByIdRespond)(nil),  // 5: microless.acmeair.flights.GetFlightByIdRespond
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*proto.FlightInfo)(nil),      // 7: microless.acmeair.FlightInfo
}
var file_proto_flights_proto_depIdxs = []int32{
	6,  // 0: microless.acmeair.flights.GetTripFlightsRequest.from_date:type_name -> google.protobuf.Timestamp
	6,  // 1: microless.acmeair.flights.GetTripFlightsRequest.return_date:type_name -> google.protobuf.Timestamp
	7,  // 2: microless.acmeair.flights.GetTripFlightsRespond.to_flights:type_name -> microless.acmeair.FlightInfo
	7,  // 3: microless.acmeair.flights.GetTripFlightsRespond.return_flights:type_name -> microless.acmeair.FlightInfo
	7,  // 4: microless.acmeair.flights.BrowseFlightsRespond.to_flights:type_name -> microless.acmeair.FlightInfo
	7,  // 5: microless.acmeair.flights.BrowseFlightsRespond.return_flights:type_name -> microless.acmeair.FlightInfo
	7,  // 6: microless.acmeair.flights.GetFlightByIdRespond.flight:type_name -> microless.acmeair.FlightInfo
	0,  // 7: microless.acmeair.flights.FlightsService.GetTripFlights:input_type -> microless.acmeair.flights.GetTripFlightsRequest
	2,  // 8: microless.acmeair.flights.FlightsService.BrowseFlights:input_type -> microless.acmeair.flights.BrowseFlightsRequest
	4,  // 9: microless.acmeair.flights.FlightsService.GetFlightById:input_type -> microless.acmeair.flights.GetFlightByIdRequest
	1,  // 10: microless.acmeair.flights.FlightsService.GetTripFlights:output_type -> microless.acmeair.flights.GetTripFlightsRespond
	3,  // 11: microless.acmeair.flights.FlightsService.BrowseFlights:output_type -> microless.acmeair.flights.BrowseFlightsRespond
	5,  // 12: microless.acmeair.flights.FlightsService.GetFlightById:output_type -> microless.acmeair.flights.GetFlightByIdRespond
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_flights_proto_init() }
func file_proto_flights_proto_init() {
	if File_proto_flights_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_flights_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripFlightsRequest); i {
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
		file_proto_flights_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTripFlightsRespond); i {
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
		file_proto_flights_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowseFlightsRequest); i {
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
		file_proto_flights_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowseFlightsRespond); i {
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
		file_proto_flights_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightByIdRequest); i {
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
		file_proto_flights_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightByIdRespond); i {
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
			RawDescriptor: file_proto_flights_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_flights_proto_goTypes,
		DependencyIndexes: file_proto_flights_proto_depIdxs,
		MessageInfos:      file_proto_flights_proto_msgTypes,
	}.Build()
	File_proto_flights_proto = out.File
	file_proto_flights_proto_rawDesc = nil
	file_proto_flights_proto_goTypes = nil
	file_proto_flights_proto_depIdxs = nil
}
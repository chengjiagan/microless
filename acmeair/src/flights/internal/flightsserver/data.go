package flightsserver

import (
	"microless/acmeair/proto"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Flight struct {
	FlightId               primitive.ObjectID `json:"flight_id" bson:"_id"`
	ScheduledDepartureTime time.Time          `json:"scheduled_departure_time" bson:"scheduled_departure_time"`
	ScheduledArrivalTime   time.Time          `json:"scheduled_arrival_time" bson:"scheduled_arrival_time"`
	FirstClassBaseCost     int64              `json:"first_class_base_cost" bson:"first_class_base_cost"`
	EconomyClassBaseCost   int64              `json:"economy_class_base_cost" bson:"economy_class_base_cost"`
	NumFirstClassSeats     int                `json:"num_first_class_seats" bson:"num_first_class_seats"`
	NumEconomyClassSeats   int                `json:"num_economy_class_seats" bson:"num_economy_class_seats"`
	AirplaneTypeId         string             `json:"airplane_type_id" bson:"airplane_type_id"`
	FlightSegment          *FlightSegment     `json:"flight_segment" bson:"flight_segment"`
}

type FlightSegment struct {
	FlightName    string `json:"flight_name" bson:"flight_name"`
	FlightSegment string `json:"flight_segment" bson:"flight_segment"`
	OriginPort    string `json:"origin_port" bson:"origin_port"`
	DestPort      string `json:"dest_port" bson:"dest_port"`
	Miles         int    `json:"miles" bson:"miles"`
}

func (f *Flight) toProto() *proto.FlightInfo {
	return &proto.FlightInfo{
		FlightId:               f.FlightId.Hex(),
		ScheduledDepartureTime: timestamppb.New(f.ScheduledDepartureTime),
		ScheduledArrivalTime:   timestamppb.New(f.ScheduledArrivalTime),
		FirstClassBaseCost:     f.FirstClassBaseCost,
		EconomyClassBaseCost:   f.EconomyClassBaseCost,
		NumFirstClassSeats:     int32(f.NumFirstClassSeats),
		NumEconomyClassSeats:   int32(f.NumEconomyClassSeats),
		AirplaneTypeId:         f.AirplaneTypeId,
		FlightSegment:          f.FlightSegment.toProto(),
	}
}

func (f *FlightSegment) toProto() *proto.FlightSegmentInfo {
	return &proto.FlightSegmentInfo{
		FlightName:    f.FlightName,
		FlightSegment: f.FlightSegment,
		OriginPort:    f.OriginPort,
		DestPort:      f.DestPort,
		Miles:         int32(f.Miles),
	}
}

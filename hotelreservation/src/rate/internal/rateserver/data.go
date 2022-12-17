package rateserver

import (
	pb "microless/hotelreservation/proto/rate"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatePlan struct {
	HotelOid primitive.ObjectID `json:"hotel_id" bson:"hotel_id"`
	UserOid  primitive.ObjectID `json:"user_id" bson:"user_id"`
	InDate   time.Time          `json:"in_date" bson:"in_date"`
	OutDate  time.Time          `json:"out_date" bson:"out_date"`
	Rate     int32              `json:"rate" bson:"rate"`
}

type HotelRate struct {
	HotelOid  primitive.ObjectID `json:"hotel_id" bson:"hotel_id"`
	TotalRate int64              `json:"total_rate" bson:"total_rate"`
	NumRate   int32              `json:"num_rate" bson:"num_rate"`
}

func (r *HotelRate) toProto() *pb.HotelRate {
	return &pb.HotelRate{
		HotelId: r.HotelOid.Hex(),
		Rate:    float64(r.TotalRate) / float64(r.NumRate),
	}
}

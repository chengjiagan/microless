package bookingsserver

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	BookingId      primitive.ObjectID `json:"booking_id" bson:"_id,omitempty"`
	OneWayFlight   bool               `json:"one_way_flight" bson:"one_way_flight"`
	ToFlightId     primitive.ObjectID `json:"to_flight_id" bson:"to_flight_id"`
	ReturnFlightId primitive.ObjectID `json:"return_flight_id" bson:"return_flight_id"`
	CustomerId     primitive.ObjectID `json:"customer_id" bson:"customer_id"`
}

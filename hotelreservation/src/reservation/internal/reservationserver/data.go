package reservationserver

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reservation struct {
	HotelOid   primitive.ObjectID `json:"hotel_id" bson:"hotel_id"`
	UserOid    primitive.ObjectID `json:"user_id" bson:"user_id"`
	InDate     time.Time          `json:"in_date" bson:"in_date"`
	OutDate    time.Time          `json:"out_date" bson:"out_date"`
	RoomNumber int                `json:"room_number" bson:"room_number"`
}

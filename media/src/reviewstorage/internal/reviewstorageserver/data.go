package reviewstorageserver

import (
	"microless/media/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Review struct {
	ReviewOid primitive.ObjectID `json:"review_id" bson:"_id,omitempty"`
	UserOid   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Text      string             `json:"text" bson:"text"`
	MovieOid  primitive.ObjectID `json:"movie_id" bson:"movie_id"`
	Rating    int32              `json:"rating" bson:"rating"`
}

func (review *Review) toProto() *proto.Review {
	timestamp := review.ReviewOid.Timestamp()
	return &proto.Review{
		ReviewId:  review.ReviewOid.Hex(),
		UserId:    review.UserOid.Hex(),
		Text:      review.Text,
		MovieId:   review.MovieOid.Hex(),
		Rating:    review.Rating,
		Timestamp: timestamppb.New(timestamp),
	}
}

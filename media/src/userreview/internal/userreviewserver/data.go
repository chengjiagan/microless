package userreviewserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserReview struct {
	UserOid    primitive.ObjectID   `bson:"user_id"`
	ReviewOids []primitive.ObjectID `bson:"review_ids"`
}

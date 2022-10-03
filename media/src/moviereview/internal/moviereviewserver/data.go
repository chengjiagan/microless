package moviereviewserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type MovieReview struct {
	MovieOid   primitive.ObjectID   `bson:"movie_id"`
	ReviewOids []primitive.ObjectID `bson:"review_ids"`
}

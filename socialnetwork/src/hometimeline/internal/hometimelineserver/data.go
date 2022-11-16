package hometimelineserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type HomeTimeline struct {
	UserId  primitive.ObjectID   `bson:"user_id"`
	PostIds []primitive.ObjectID `bson:"post_ids"`
}

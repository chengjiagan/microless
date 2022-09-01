package usertimelineserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserTimeline struct {
	UserId  primitive.ObjectID   `bson:"user_id"`
	PostIds []primitive.ObjectID `bson:"post_ids"`
}

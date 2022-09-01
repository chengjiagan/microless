package socialgraphserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSocialGraph struct {
	UserId    primitive.ObjectID   `bson:"user_id"`
	Followers []primitive.ObjectID `bson:"followers"`
	Followees []primitive.ObjectID `bson:"followees"`
}

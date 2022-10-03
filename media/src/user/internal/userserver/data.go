package userserver

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserOid   primitive.ObjectID `json:"user_oid,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Username  string             `json:"username" bson:"username"`
	Salt      string             `json:"salt" bson:"salt"`
	Password  string             `json:"password" bson:"password"`
}

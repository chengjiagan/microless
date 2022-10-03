package castinfoserver

import (
	"microless/media/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CastInfo struct {
	CastInfoOid primitive.ObjectID `json:"cast_info_id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Gender      bool               `json:"gender" bson:"gender"`
	Intro       string             `json:"intro" bson:"intro"`
}

func (info *CastInfo) toProto() *proto.CastInfo {
	return &proto.CastInfo{
		CastInfoId: info.CastInfoOid.Hex(),
		Name:       info.Name,
		Gender:     info.Gender,
		Intro:      info.Intro,
	}
}

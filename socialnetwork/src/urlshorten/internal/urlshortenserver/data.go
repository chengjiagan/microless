package urlshortenserver

import "microless/socialnetwork/proto"

type Url struct {
	ExpandedUrl  string `json:"expanded_url" bson:"expanded_url"`
	ShortenedUrl string `json:"shortened_url" bson:"shortened_url"`
}

func (u *Url) toProto() *proto.Url {
	return &proto.Url{
		ExpandedUrl:  u.ExpandedUrl,
		ShortenedUrl: u.ShortenedUrl,
	}
}

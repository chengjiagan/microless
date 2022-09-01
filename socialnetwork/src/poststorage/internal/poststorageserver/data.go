package poststorageserver

import (
	"microless/socialnetwork/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostType string

const (
	POST   = "POST"
	REPOST = "REPOST"
	REPLY  = "REPLY"
	DM     = "DM"
)

type Post struct {
	PostOid      primitive.ObjectID `json:"post_id" bson:"_id,omitempty"`
	Text         string             `json:"text" bson:"text"`
	Creator      User               `json:"creator" bson:"creator"`
	UserMentions []User             `json:"user_mentions" bson:"user_mentions"`
	Media        []Media            `json:"media" bson:"media"`
	Urls         []Url              `json:"urls" bson:"urls"`
	PostType     PostType           `json:"post_type" bson:"post_type"`
}

type User struct {
	UserId   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Username string             `json:"username" bson:"username"`
}

type Media struct {
	MediaId   int64  `json:"media_id" bson:"media_id"`
	MediaType string `json:"media_type" bson:"media_type"`
}

type Url struct {
	ExpandedUrl  string `json:"expanded_url" bson:"expanded_url"`
	ShortenedUrl string `json:"shortened_url" bson:"shortened_url"`
}

func postFromProto(pb *proto.Post) *Post {
	creatorOid, _ := primitive.ObjectIDFromHex(pb.Creator.UserId)
	creator := User{
		UserId:   creatorOid,
		Username: pb.Creator.Username,
	}

	mentions := make([]User, len(pb.UserMentions))
	for i, m := range pb.UserMentions {
		oid, _ := primitive.ObjectIDFromHex(m.UserId)
		mentions[i] = User{
			UserId:   oid,
			Username: m.Username,
		}
	}

	media := make([]Media, len(pb.Media))
	for i, m := range pb.Media {
		media[i] = Media{
			MediaId:   m.MediaId,
			MediaType: m.MediaType,
		}
	}

	urls := make([]Url, len(pb.Urls))
	for i, u := range pb.Urls {
		urls[i] = Url{
			ShortenedUrl: u.ShortenedUrl,
			ExpandedUrl:  u.ExpandedUrl,
		}
	}

	var postType PostType
	switch pb.PostType {
	case proto.PostType_POST:
		postType = POST
	case proto.PostType_REPOST:
		postType = REPOST
	case proto.PostType_REPLY:
		postType = REPLY
	case proto.PostType_DM:
		postType = DM
	}

	return &Post{
		Creator:      creator,
		Text:         pb.Text,
		UserMentions: mentions,
		Media:        media,
		Urls:         urls,
		PostType:     postType,
	}
}

func (post *Post) toProto() *proto.Post {
	creator := &proto.Creator{
		UserId:   post.Creator.UserId.Hex(),
		Username: post.Creator.Username,
	}

	mentions := make([]*proto.UserMention, len(post.UserMentions))
	for i, m := range post.UserMentions {
		mentions[i] = &proto.UserMention{
			UserId:   m.UserId.Hex(),
			Username: m.Username,
		}
	}

	media := make([]*proto.Media, len(post.Media))
	for i, m := range post.Media {
		media[i] = &proto.Media{
			MediaId:   m.MediaId,
			MediaType: m.MediaType,
		}
	}

	urls := make([]*proto.Url, len(post.Urls))
	for i, u := range post.Urls {
		urls[i] = &proto.Url{
			ShortenedUrl: u.ShortenedUrl,
			ExpandedUrl:  u.ExpandedUrl,
		}
	}

	timestamp := timestamppb.New(post.PostOid.Timestamp())

	var postType proto.PostType
	switch post.PostType {
	case POST:
		postType = proto.PostType_POST
	case REPOST:
		postType = proto.PostType_REPOST
	case REPLY:
		postType = proto.PostType_REPLY
	case DM:
		postType = proto.PostType_DM
	}

	return &proto.Post{
		Creator:      creator,
		Text:         post.Text,
		UserMentions: mentions,
		Media:        media,
		Urls:         urls,
		Timestamp:    timestamp,
		PostType:     postType,
	}
}

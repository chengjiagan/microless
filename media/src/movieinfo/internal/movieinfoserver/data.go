package movieinfoserver

import (
	"microless/media/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieInfo struct {
	MovieOid     primitive.ObjectID `json:"movie_id" bson:"_id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Casts        []*Cast            `json:"casts" bson:"casts"`
	PlotOid      primitive.ObjectID `json:"plot_id" bson:"plot_id"`
	ThumbnailIds []string           `json:"thumbnail_ids" bson:"thumbnail_ids"`
	PhotoIds     []string           `json:"photo_ids" bson:"photo_ids"`
	VideoIds     []string           `json:"video_ids" bson:"video_ids"`
	AvgRating    float64            `json:"avg_rating" bson:"avg_rating"`
	NumRating    int32              `json:"num_rating" bson:"num_rating"`
}

func (info *MovieInfo) toProto() *proto.MovieInfo {
	casts := make([]*proto.Cast, len(info.Casts))
	for i, cast := range info.Casts {
		casts[i] = cast.toProto()
	}

	return &proto.MovieInfo{
		MovieId:      info.MovieOid.Hex(),
		Title:        info.Title,
		Casts:        casts,
		PlotId:       info.PlotOid.Hex(),
		ThumbnailIds: info.ThumbnailIds,
		PhotoIds:     info.PhotoIds,
		VideoIds:     info.VideoIds,
		AvgRating:    info.AvgRating,
		NumRating:    info.NumRating,
	}
}

type Cast struct {
	CastId      int32              `json:"cast_id" bson:"cast_id"`
	Character   string             `json:"character" bson:"character"`
	CastInfoOid primitive.ObjectID `json:"cast_info_id" bson:"cast_info_id"`
}

func castFromProto(cast *proto.Cast) *Cast {
	oid, _ := primitive.ObjectIDFromHex(cast.CastInfoId)

	return &Cast{
		CastId:      cast.CastId,
		Character:   cast.Character,
		CastInfoOid: oid,
	}
}

func (cast *Cast) toProto() *proto.Cast {
	return &proto.Cast{
		CastId:     cast.CastId,
		Character:  cast.Character,
		CastInfoId: cast.CastInfoOid.Hex(),
	}
}

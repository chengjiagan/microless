package movieinfoserver

import (
	"context"

	pb "microless/media/proto/movieinfo"
	"microless/media/proto/moviereview"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MovieInfoServer) WriteMovieInfo(ctx context.Context, req *pb.WriteMovieInfoRequest) (*pb.WriteMovieInfoRespond, error) {
	// convert protobuf to MongoDB data model
	casts := make([]*Cast, len(req.Casts))
	for i, cast := range req.Casts {
		casts[i] = castFromProto(cast)
	}
	plotOid, _ := primitive.ObjectIDFromHex(req.PlotId)
	info := &MovieInfo{
		Title:        req.Title,
		Casts:        casts,
		PlotOid:      plotOid,
		ThumbnailIds: req.ThumbnailIds,
		PhotoIds:     req.PhotoIds,
		VideoIds:     req.VideoIds,
		AvgRating:    req.AvgRating,
		NumRating:    req.NumRating,
	}

	// insert new movie info into MongoDB
	s.logger.Info("Insert new movie info")
	result, err := s.mongodb.InsertOne(ctx, info)
	if err != nil {
		s.logger.Warnw("Failed to insert movie info into MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	oid := result.InsertedID.(primitive.ObjectID)

	// create movie in movie-review-service
	s.logger.Info("Create movie in movie-review-service")
	reviewReq := &moviereview.CreateMovieRequest{
		MovieId: oid.Hex(),
	}
	_, err = s.moviereviewClient.CreateMovie(ctx, reviewReq)
	if err != nil {
		s.logger.Warnw("Failed to create movie in movie-review-service", "err", err)
		return nil, err
	}

	return &pb.WriteMovieInfoRespond{MovieId: oid.Hex()}, nil
}

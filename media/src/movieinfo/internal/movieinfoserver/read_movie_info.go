package movieinfoserver

import (
	"context"
	"encoding/json"
	"microless/media/proto"
	pb "microless/media/proto/movieinfo"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MovieInfoServer) ReadMovieInfo(ctx context.Context, req *pb.ReadMovieInfoRequest) (*proto.MovieInfo, error) {
	movie, err := s.getMovieInfo(ctx, req.MovieId)
	if err != nil {
		return nil, err
	}

	// update memcached
	s.logger.Info("Update movie info in memcached")
	infoJson, _ := json.Marshal(movie)
	err = s.memcached.WithContext(ctx).Set(&memcache.Item{
		Key:   req.MovieId,
		Value: infoJson,
	})
	if err != nil {
		s.logger.Warnw("Failed to update Memcached", "movie_id", req.MovieId, "err", err)
	}

	return movie.toProto(), nil
}

func (s *MovieInfoServer) getMovieInfo(ctx context.Context, movieId string) (*MovieInfo, error) {
	movie := new(MovieInfo)

	// get movie info from memcached
	s.logger.Info("Read movie info from Memcached")
	resMc, err := s.memcached.WithContext(ctx).Get(movieId)
	if err != nil {
		s.logger.Warnw("Failed to get movie info from Memcached", "movie_id", movieId, "err", err)
	} else {
		// cache hit
		json.Unmarshal(resMc.Value, movie)
		return movie, nil
	}

	// cache miss, get movie info from mongodb
	s.logger.Info("Read movie info from MongoDB")
	oid, _ := primitive.ObjectIDFromHex(movieId)
	query := bson.M{"_id": oid}
	err = s.mongodb.FindOne(ctx, query).Decode(movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.logger.Warnw("Movie info not found", "movie_id", movieId, "err", err)
			return nil, status.Errorf(codes.NotFound, "Movie info %v not found", movieId)
		} else {
			s.logger.Warnw("Failed to get movie info from MongoDB", "movie_id", movieId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}
	return movie, nil
}

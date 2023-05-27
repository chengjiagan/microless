package moviereviewserver

import (
	"context"
	"microless/media/proto"
	pb "microless/media/proto/moviereview"
	"microless/media/proto/reviewstorage"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *MovieReviewService) ReadMovieReviews(ctx context.Context, req *pb.ReadMovieReviewsRequest) (*pb.ReadMovieReviewsRespond, error) {
	if req.Stop <= req.Start || req.Start < 0 {
		s.logger.Warnw("Invalid arguments", "movie_id", req.MovieId, "start", req.Start, "stop", req.Stop)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid arguments for ReadMovieReviews")
	}

	// get movie reviews from redis
	s.logger.Info("Get movie reviews from Redis")
	reviewIds, err := s.rdb.ZRevRange(ctx, req.MovieId, int64(req.Start), int64(req.Stop-1)).Result()
	if err != nil {
		s.logger.Warnw("Failed to get movie reviews from Redis", "movie_id", req.MovieId, "err", err)
	} else {
		// get all from redis
		if len(reviewIds) == int(req.Stop-req.Start) {
			s.logger.Info("Get reviews from review-storage-service")
			reviewReq := &reviewstorage.ReadReviewsRequest{ReviewIds: reviewIds}
			reviewResp, err := s.reviewstorageClient.ReadReviews(ctx, reviewReq)
			if err != nil {
				s.logger.Warnw("Failed to get reviews from review-storage-service", "err", err)
				return nil, err
			}
			return &pb.ReadMovieReviewsRespond{Reviews: reviewResp.Reviews}, nil
		}
	}

	// find in mongodb
	s.logger.Info("Get movie reviews from MongoDB")
	movieOid, _ := primitive.ObjectIDFromHex(req.MovieId)
	query := bson.M{"movie_id": movieOid}
	opts := options.FindOne().SetProjection(
		bson.M{
			"review_ids": bson.M{"$slice": bson.A{0, req.Stop}},
		},
	)
	doc := new(MovieReview)
	err = s.mongodb.FindOne(ctx, query, opts).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "movie_id: %v doesn't exit in MongoDB", req.MovieId)
		} else {
			s.logger.Warnw("Failed to get movie reviews from MongoDB", "movie_id", req.MovieId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}
	reviewOids := doc.ReviewOids

	g, ctx := errgroup.WithContext(ctx)
	// update redis
	g.Go(func() error {
		s.logger.Info("Update movie reviews in Redis")
		redisUpdate := make([]*redis.Z, len(reviewOids))
		for i, oid := range reviewOids {
			redisUpdate[i] = &redis.Z{
				Score:  float64(oid.Timestamp().Unix()),
				Member: oid.Hex(),
			}
		}

		_, err := s.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Del(ctx, req.MovieId)
			p.ZAdd(ctx, req.MovieId, redisUpdate...)
			return nil
		})
		if err != nil {
			s.logger.Warnw("Failed to update movie reviews in Redis", "err", err)
		}
		return nil
	})

	// get reviews from review-storage
	var reviews []*proto.Review
	g.Go(func() error {
		reviewIds = make([]string, int(req.Stop-req.Start))
		for i := range reviewIds {
			reviewIds[i] = reviewOids[i+int(req.Start)].Hex()
		}

		s.logger.Info("Get reviews from review-storage-service")
		reviewReq := &reviewstorage.ReadReviewsRequest{ReviewIds: reviewIds}
		reviewResp, err := s.reviewstorageClient.ReadReviews(ctx, reviewReq)
		if err != nil {
			s.logger.Error("Failed to get reviews from review-storage-service")
			return err
		}
		reviews = reviewResp.Reviews
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, err
	}
	return &pb.ReadMovieReviewsRespond{Reviews: reviews}, nil
}

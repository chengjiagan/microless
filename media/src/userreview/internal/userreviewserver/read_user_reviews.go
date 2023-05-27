package userreviewserver

import (
	"context"
	"microless/media/proto"
	"microless/media/proto/reviewstorage"
	pb "microless/media/proto/userreview"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserReviewService) ReadUserReviews(ctx context.Context, req *pb.ReadUserReviewsRequest) (*pb.ReadUserReviewsRespond, error) {
	if req.Stop <= req.Start || req.Start < 0 {
		s.logger.Warnw("Invalid arguments", "user_id", req.UserId, "start", req.Start, "stop", req.Stop)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid arguments for ReadUserReviews")
	}

	// get user reviews from redis
	s.logger.Info("Get user reviews from Redis")
	reviewIds, err := s.rdb.ZRevRange(ctx, req.UserId, int64(req.Start), int64(req.Stop-1)).Result()
	if err != nil {
		s.logger.Warnw("Failed to get user reviews from Redis", "user_id", req.UserId, "err", err)
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
			return &pb.ReadUserReviewsRespond{Reviews: reviewResp.Reviews}, nil
		}
	}

	// find in mongodb
	s.logger.Info("Get user reviews from MongoDB")
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	query := bson.M{"user_id": userOid}
	opts := options.FindOne().SetProjection(
		bson.M{
			"review_ids": bson.M{"$slice": bson.A{0, req.Stop}},
		},
	)
	doc := new(UserReview)
	err = s.mongodb.FindOne(ctx, query, opts).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exit in MongoDB", req.UserId)
		} else {
			s.logger.Warnw("Failed to get user reviews from MongoDB", "user_id", req.UserId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}
	reviewOids := doc.ReviewOids

	g, ctx := errgroup.WithContext(ctx)
	// update redis
	g.Go(func() error {
		s.logger.Info("Update user reviews in Redis")
		redisUpdate := make([]*redis.Z, len(reviewOids))
		for i, oid := range reviewOids {
			redisUpdate[i] = &redis.Z{
				Score:  float64(oid.Timestamp().Unix()),
				Member: oid.Hex(),
			}
		}

		_, err := s.rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
			p.Del(ctx, req.UserId)
			p.ZAdd(ctx, req.UserId, redisUpdate...)
			return nil
		})
		if err != nil {
			s.logger.Warnw("Failed to update user reviews in Redis", "err", err)
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
	return &pb.ReadUserReviewsRespond{Reviews: reviews}, nil
}

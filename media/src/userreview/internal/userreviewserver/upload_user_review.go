package userreviewserver

import (
	"context"
	pb "microless/media/proto/userreview"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserReviewService) UploadUserReview(ctx context.Context, req *pb.UploadUserReviewRequest) (*emptypb.Empty, error) {
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	reviewOid, _ := primitive.ObjectIDFromHex(req.ReviewId)

	// update user reviews in mongodb
	s.logger.Info("Update user review in MongoDB")
	query := bson.M{"user_id": userOid}
	update := bson.M{
		"$push": bson.M{
			"review_ids": bson.M{
				"$each": bson.A{reviewOid},
				"$sort": -1,
			},
		},
	}
	res, err := s.mongodb.UpdateOne(ctx, query, update)
	if err != nil {
		s.logger.Errorw("Failed to update user reviews", "user_id", req.UserId, "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	if res.MatchedCount < 1 {
		s.logger.Errorw("Unknown user", "user_id", req.UserId)
		return nil, status.Errorf(codes.NotFound, "user_id: %v doesn't exist", req.UserId)
	}

	// invalidate cache in redis
	s.logger.Info("Delete user reviews in Redis")
	err = s.rdb.Del(ctx, req.UserId).Err()
	if err != nil {
		s.logger.Warnw("Failed to delete user review in Redis", "err", err)
	}

	return &emptypb.Empty{}, nil
}

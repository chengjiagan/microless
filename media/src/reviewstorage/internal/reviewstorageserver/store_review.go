package reviewstorageserver

import (
	"context"
	pb "microless/media/proto/reviewstorage"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ReviewStorageService) StoreReview(ctx context.Context, req *pb.StoreReviewRequest) (*pb.StoreReviewRespond, error) {
	// generate the review object
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	movieOid, _ := primitive.ObjectIDFromHex(req.MovieId)
	review := &Review{
		UserOid:  userOid,
		Text:     req.Text,
		MovieOid: movieOid,
		Rating:   req.Rating,
	}

	// insert new review into mongodb
	result, err := s.mongodb.InsertOne(ctx, review)
	if err != nil {
		s.logger.Errorw("Failed to insert review into MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	s.logger.Info("Insert new review")
	oid := result.InsertedID.(primitive.ObjectID)
	respond := &pb.StoreReviewRespond{
		ReviewId:  oid.Hex(),
		Timestamp: timestamppb.New(oid.Timestamp()),
	}
	return respond, nil
}

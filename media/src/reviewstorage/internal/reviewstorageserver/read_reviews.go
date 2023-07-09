package reviewstorageserver

import (
	"context"
	"encoding/json"
	"microless/media/proto"
	pb "microless/media/proto/reviewstorage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ReviewStorageService) ReadReviews(ctx context.Context, req *pb.ReadReviewsRequest) (*pb.ReadReviewsRespond, error) {
	// ignore empty request
	if len(req.ReviewIds) == 0 {
		return &pb.ReadReviewsRespond{}, nil
	}

	reviews := make(map[string]*Review, len(req.ReviewIds))

	// get reviews from redis
	s.logger.Info("Get reviews from Redis")
	reviewsCache, err := s.rdb.MGet(ctx, req.ReviewIds...).Result()
	if err != nil {
		s.logger.Warnw("Failed to get reviews from Redis", "review_ids", req.ReviewIds, "err", err)
	} else {
		for i, v := range reviewsCache {
			if v != nil {
				review := new(Review)
				json.Unmarshal(v.([]byte), review)
				reviews[req.ReviewIds[i]] = review
			}
		}
	}

	// get all reviews from redis
	if len(reviews) == len(req.ReviewIds) {
		pbReviews := make([]*proto.Review, len(req.ReviewIds))
		for i, id := range req.ReviewIds {
			pbReviews[i] = reviews[id].toProto()
		}
		return &pb.ReadReviewsRespond{Reviews: pbReviews}, nil
	}

	// get reviews from mongodb
	s.logger.Info("Get reviews from MongoDB")
	oids := make([]primitive.ObjectID, 0, len(req.ReviewIds)-len(reviews))
	for _, id := range req.ReviewIds {
		if _, ok := reviews[id]; !ok {
			oid, _ := primitive.ObjectIDFromHex(id)
			oids = append(oids, oid)
		}
	}
	query := bson.M{"_id": bson.M{"$in": oids}}
	cursor, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Warnw("Failed to find reviews from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// decode from mongodb
	var reviewsMongo []*Review
	err = cursor.All(ctx, &reviewsMongo)
	if err != nil {
		s.logger.Warnw("Failed to find reviews from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// update redis
	reviewsMiss := make([]interface{}, 0, len(reviewsMongo))
	for _, review := range reviewsMongo {
		id := review.ReviewOid.Hex()
		reviews[id] = review
		reviewJson, _ := json.Marshal(review)
		reviewsMiss = append(reviewsMiss, id, reviewJson)
	}
	_, err = s.rdb.MSet(ctx, reviewsMiss...).Result()
	if err != nil {
		s.logger.Warnw("Failed to set reviews to Redis", "err", err)
	}

	// still unknown review_id exists
	if len(reviews) != len(req.ReviewIds) {
		s.logger.Warn("Unknown review_id")
		return nil, status.Error(codes.NotFound, "Unknown review_id")
	}

	pbReviews := make([]*proto.Review, len(req.ReviewIds))
	for i, id := range req.ReviewIds {
		pbReviews[i] = reviews[id].toProto()
	}
	return &pb.ReadReviewsRespond{Reviews: pbReviews}, nil
}

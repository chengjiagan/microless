package reviewstorageserver

import (
	"context"
	"encoding/json"
	"microless/media/proto"
	pb "microless/media/proto/reviewstorage"

	"github.com/bradfitz/gomemcache/memcache"
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

	// get reviews from memcached
	s.logger.Info("Get reviews from Memcached")
	reviewsMc, err := s.memcached.WithContext(ctx).GetMulti(req.ReviewIds)
	if err != nil {
		s.logger.Warnw("Failed to get reviews from Memcached", "review_ids", req.ReviewIds, "err", err)
	} else {
		for _, item := range reviewsMc {
			review := new(Review)
			json.Unmarshal(item.Value, review)
			reviews[item.Key] = review
		}
	}

	// get all reviews from memcached
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
	for _, review := range reviewsMongo {
		id := review.ReviewOid.Hex()
		reviews[id] = review

		// upload infos to memcached
		reviewJson, _ := json.Marshal(review)
		err = s.memcached.WithContext(ctx).Set(&memcache.Item{
			Key:   id,
			Value: reviewJson,
		})
		if err != nil {
			s.logger.Warnw("Failed to update Memcached", "review_id", id, "err", err)
		}
	}

	// still unknown cast_info_id exists
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

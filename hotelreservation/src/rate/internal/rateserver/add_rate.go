package rateserver

import (
	"context"

	pb "microless/hotelreservation/proto/rate"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO: we need two mongo collections: rate_plan and hotel_rate

func (s *RateService) AddRate(ctx context.Context, req *pb.AddRateRequest) (*emptypb.Empty, error) {
	hotelOid, _ := primitive.ObjectIDFromHex(req.HotelId)
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)

	// get rate plan from request
	ratePlan := RatePlan{
		HotelOid: hotelOid,
		UserOid:  userOid,
		InDate:   req.InDate.AsTime(),
		OutDate:  req.OutDate.AsTime(),
		Rate:     req.Rate,
	}

	// insert rate plan into database
	s.logger.Info("Insert rate plan into MongoDB")
	_, err := s.ratePlanDb.InsertOne(ctx, ratePlan)
	if err != nil {
		s.logger.Warnw("Failed to insert rate plan into database", "err", err)
		return nil, err
	}

	// update hotel rate
	s.logger.Info("Update hotel rate")
	query := bson.M{"hotel_id": hotelOid}
	update := bson.M{
		"$inc": bson.M{
			"total_rate": int64(req.Rate),
			"num_rate":   1,
		},
	}
	res, err := s.hotelRateDb.UpdateOne(ctx, query, update)
	if err != nil {
		s.logger.Warnw("Failed to update hotel rate", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	if res.MatchedCount == 0 {
		s.logger.Warnw("Unknown hotel_id", "hotel_id", req.HotelId)
		return nil, status.Errorf(codes.NotFound, "Hotel %v not found", req.HotelId)
	}

	// delete rate from memcached
	s.logger.Info("Delete rate from Memcached")
	err = s.memcached.WithContext(ctx).Delete(req.HotelId)
	if err != nil {
		s.logger.Warnw("Failed to delete rate from Memcached", "err", err)
	}

	return &emptypb.Empty{}, nil
}

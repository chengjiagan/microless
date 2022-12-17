package geoserver

import (
	"context"
	pb "microless/hotelreservation/proto/geo"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GeoService) AddHotel(ctx context.Context, req *pb.AddHotelRequest) (*emptypb.Empty, error) {
	// insert new location to mongodb
	s.logger.Info("Insert new hotel location to mongodb")
	hotelOid, _ := primitive.ObjectIDFromHex(req.HotelId)
	hotel := &HotelLocation{
		HotelId: hotelOid,
		Location: GeoJson{
			Type:        "Point",
			Coordinates: []float64{req.Lon, req.Lat},
		},
	}
	_, err := s.mongodb.InsertOne(ctx, hotel)
	if err != nil {
		s.logger.Errorw("Failed to insert new hotel location to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// delete old query caches from redis
	s.logger.Info("Delete old query caches from redis")
	// get query caches around the new hotel location
	opts := &redis.GeoRadiusQuery{
		Radius: maxSearchRadius,
		Unit:   "km",
	}
	result, err := s.rdb.GeoRadius(ctx, queryCacheKey, req.Lon, req.Lat, opts).Result()
	if err != nil {
		s.logger.Warnw("Failed to delete old query caches from Redis", "err", err)
		return &emptypb.Empty{}, nil
	}
	// delete query caches
	p := s.rdb.Pipeline()
	for _, item := range result {
		p.Del(ctx, item.Name)
	}
	_, err = p.Exec(ctx)
	if err != nil {
		s.logger.Warnw("Failed to delete old query caches from Redis", "err", err)
	}

	return &emptypb.Empty{}, nil
}

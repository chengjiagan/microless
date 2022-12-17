package geoserver

import (
	"context"
	"fmt"
	pb "microless/hotelreservation/proto/geo"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GeoService) Nearby(ctx context.Context, req *pb.NearbyRequest) (*pb.NearbyRespond, error) {
	// get naerby hotels from redis
	s.logger.Info("Get nearby hotels from redis")
	key := fmt.Sprintf("%.3f:%.3f", req.Lon, req.Lat)
	// check existance first, since ZRange will return empty slice if key does not exist
	if s.rdb.Exists(ctx, key).Val() == 1 {
		hotelIds, err := s.rdb.ZRange(ctx, key, 0, -1).Result()
		if err != nil {
			s.logger.Warnw("Failed to get nearby hotels from Redis", "err", err)
		} else {
			return &pb.NearbyRespond{HotelIds: hotelIds}, nil
		}
	}

	// get nearby hotels from mongodb
	s.logger.Info("Get nearby hotels from mongodb")
	query := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{req.Lon, req.Lat},
				},
				"$maxDistance": maxSearchRadius * 1000,
			},
		},
	}
	opts := options.Find().SetLimit(maxSearchResults)
	cur, err := s.mongodb.Find(ctx, query, opts)
	if err != nil {
		s.logger.Errorw("Failed to get nearby hotels from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	hotels := make([]*HotelLocation, 0)
	err = cur.All(ctx, &hotels)
	if err != nil {
		s.logger.Errorw("Failed to get nearby hotels from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// get hotel ids
	hotelIds := make([]string, len(hotels))
	for i, hotel := range hotels {
		hotelIds[i] = hotel.HotelId.Hex()
	}

	// update redis
	s.logger.Info("Update redis")
	p := s.rdb.Pipeline()
	p.GeoAdd(ctx, queryCacheKey, &redis.GeoLocation{
		Name:      key,
		Longitude: req.Lon,
		Latitude:  req.Lat,
	})
	for i, hotelId := range hotelIds {
		p.ZAdd(ctx, key, &redis.Z{
			Score:  float64(i),
			Member: hotelId,
		})
	}
	_, err = p.Exec(ctx)
	if err != nil {
		s.logger.Warnw("Failed to update redis", "err", err)
	}

	return &pb.NearbyRespond{HotelIds: hotelIds}, nil
}

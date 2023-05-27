package flightsserver

import (
	"context"
	"encoding/json"
	"fmt"
	pb "microless/acmeair/proto/flights"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightsService) GetFlightById(ctx context.Context, req *pb.GetFlightByIdRequest) (*pb.GetFlightByIdRespond, error) {
	flight := new(Flight)

	// get flight from memcached
	s.logger.Info("Get flight from memcached")
	key := fmt.Sprintf("flight:%s", req.FlightId)
	flightMc, err := s.memcached.WithContext(ctx).Get(key)
	if err == nil {
		// cache hit
		json.Unmarshal(flightMc.Value, flight)
		return &pb.GetFlightByIdRespond{Flight: flight.toProto()}, nil
	} else if err != memcache.ErrCacheMiss {
		s.logger.Warnw("Failed to get flight from memcached", "err", err)
	}

	// cache miss, get flight from mongodb
	s.logger.Info("Get flight from MongoDB")
	flightOid, _ := primitive.ObjectIDFromHex(req.FlightId)
	query := bson.M{"_id": flightOid}
	err = s.mongodb.FindOne(ctx, query).Decode(flight)
	if err != nil {
		s.logger.Warnw("Failed to get flight from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB err: %v", err)
	}

	// update memcached
	s.logger.Info("Update memcached")
	flightJson, _ := json.Marshal(flight)
	err = s.memcached.WithContext(ctx).Set(
		&memcache.Item{
			Key:   key,
			Value: flightJson,
		},
	)
	if err != nil {
		s.logger.Warnw("Failed to update memcached", "err", err)
	}

	return &pb.GetFlightByIdRespond{Flight: flight.toProto()}, nil
}

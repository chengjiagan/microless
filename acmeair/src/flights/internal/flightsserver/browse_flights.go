package flightsserver

import (
	"context"
	"encoding/json"
	"fmt"
	"microless/acmeair/proto"
	pb "microless/acmeair/proto/flights"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightsService) BrowseFlights(ctx context.Context, req *pb.BrowseFlightsRequest) (*pb.BrowseFlightsRespond, error) {
	// get to flights
	toFlights, err := s.getFlightByPort(ctx, req.FromAirport, req.ToAirport)
	if err != nil {
		return nil, err
	}
	toFlightsPb := make([]*proto.FlightInfo, len(toFlights))
	for i, f := range toFlights {
		toFlightsPb[i] = f.toProto()
	}

	if req.OneWayFlight {
		// don't get return flights
		return &pb.BrowseFlightsRespond{
			ToFlights:    toFlightsPb,
			OneWayFlight: req.OneWayFlight,
		}, nil
	}

	// get return flights
	retFlights, err := s.getFlightByPort(ctx, req.ToAirport, req.FromAirport)
	if err != nil {
		return nil, err
	}
	retFlightsPb := make([]*proto.FlightInfo, len(retFlights))
	for i, f := range retFlights {
		retFlightsPb[i] = f.toProto()
	}

	return &pb.BrowseFlightsRespond{
		ToFlights:     toFlightsPb,
		ReturnFlights: retFlightsPb,
		OneWayFlight:  req.OneWayFlight,
	}, nil
}

func (s *FlightsService) getFlightByPort(ctx context.Context, origPort, destPort string) ([]*Flight, error) {
	flights := make([]*Flight, 0)

	// get flight from memcached
	s.logger.Info("Get flights by port from memcached")
	key := fmt.Sprintf("orig_port:%s:dest_port:%s", origPort, destPort)
	flightMc, err := s.memcached.WithContext(ctx).Get(key)
	if err == nil {
		// cache hit
		json.Unmarshal(flightMc.Value, &flights)
		return flights, nil
	} else if err != memcache.ErrCacheMiss {
		s.logger.Warnw("Failed to get flights by port from memcached", "err", err)
	}

	// cache miss, get flight from mongodb
	s.logger.Info("Get flights by port from MongoDB")
	query := bson.M{
		"flight_segment.origin_port": origPort,
		"flight_segment.dest_port":   destPort,
	}
	cur, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Errorw("Failed to get flights by port from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB err: %v", err)
	}
	err = cur.All(ctx, &flights)
	if err != nil {
		s.logger.Errorw("Failed to get flights by port from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB err: %v", err)
	}

	// update memcached
	s.logger.Info("Update memcached")
	flightJson, _ := json.Marshal(flights)
	err = s.memcached.WithContext(ctx).Set(
		&memcache.Item{
			Key:   key,
			Value: flightJson,
		},
	)
	if err != nil {
		s.logger.Warnw("Failed to update memcached", "err", err)
	}

	return flights, nil
}

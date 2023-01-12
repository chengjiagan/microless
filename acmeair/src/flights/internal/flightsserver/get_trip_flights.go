package flightsserver

import (
	"context"
	"encoding/json"
	"fmt"
	"microless/acmeair/proto"
	pb "microless/acmeair/proto/flights"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *FlightsService) GetTripFlights(ctx context.Context, req *pb.GetTripFlightsRequest) (*pb.GetTripFlightsRespond, error) {
	// get to flights
	toFlights, err := s.getFlightByPortAndDate(ctx, req.FromAirport, req.ToAirport, req.FromDate.AsTime())
	if err != nil {
		return nil, err
	}
	toFlightsPb := make([]*proto.FlightInfo, len(toFlights))
	for i, f := range toFlights {
		toFlightsPb[i] = f.toProto()
	}

	if req.OneWayFlight {
		// don't get return flights
		return &pb.GetTripFlightsRespond{
			ToFlights:    toFlightsPb,
			OneWayFlight: req.OneWayFlight,
		}, nil
	}

	// get return flights
	retFlights, err := s.getFlightByPortAndDate(ctx, req.ToAirport, req.FromAirport, req.ReturnDate.AsTime())
	if err != nil {
		return nil, err
	}
	retFlightsPb := make([]*proto.FlightInfo, len(retFlights))
	for i, f := range retFlights {
		retFlightsPb[i] = f.toProto()
	}

	return &pb.GetTripFlightsRespond{
		ToFlights:     toFlightsPb,
		ReturnFlights: retFlightsPb,
		OneWayFlight:  req.OneWayFlight,
	}, nil
}

func (s *FlightsService) getFlightByPortAndDate(ctx context.Context, origPort, destPort string, date time.Time) ([]*Flight, error) {
	var flights []*Flight

	// get flight from memcached
	s.logger.Info("Get flights by port from memcached")
	key := fmt.Sprintf("orig_port:%s:dest_port:%s:date:%v", origPort, destPort, date)
	flightMc, err := s.memcached.WithContext(ctx).Get(key)
	if err == nil {
		// cache hit
		json.Unmarshal(flightMc.Value, &flights)
		return flights, nil
	} else if err != memcache.ErrCacheMiss {
		s.logger.Warnw("Failed to get flights by port and date from memcached", "err", err)
	}

	// cache miss, get flight from mongodb
	s.logger.Info("Get flights by port from MongoDB")
	query := bson.M{
		"flight_segment.origin_port": origPort,
		"flight_segment.dest_port":   destPort,
		"scheduled_departure_time":   date,
	}
	cur, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Errorw("Failed to get flights by port and date from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB err: %v", err)
	}
	err = cur.All(ctx, &flights)
	if err != nil {
		s.logger.Errorw("Failed to get flights by port and date from MongoDB", "err", err)
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

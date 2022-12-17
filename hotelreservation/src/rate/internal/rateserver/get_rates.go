package rateserver

import (
	"context"
	pb "microless/hotelreservation/proto/rate"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func (s *RateService) GetRates(ctx context.Context, req *pb.GetRatesRequest) (*pb.GetRatesRespond, error) {
	// return empty respond when request is empty
	if len(req.HotelIds) == 0 {
		return &pb.GetRatesRespond{}, nil
	}

	rates := make(map[string]*pb.HotelRate)

	// get rates from memcached
	s.logger.Info("Get rates from Memcached")
	ratesMc, err := s.memcached.WithContext(ctx).GetMulti(req.HotelIds)
	if err != nil {
		s.logger.Warnw("Failed to get rates from Memcached", "err", err)
	} else {
		for _, item := range ratesMc {
			rate := new(pb.HotelRate)
			protojson.Unmarshal(item.Value, rate)
			rates[item.Key] = rate
		}
	}

	// get all rate from memcached
	if len(rates) == len(req.HotelIds) {
		result := make([]*pb.HotelRate, 0, len(rates))
		for _, rate := range rates {
			result = append(result, rate)
		}
		return &pb.GetRatesRespond{Rates: result}, nil
	}

	// get rates from mongodb
	s.logger.Info("Get rates from MongoDB")
	oids := make([]primitive.ObjectID, 0, len(req.HotelIds)-len(rates))
	// get uncached hotel ids
	for _, id := range req.HotelIds {
		if _, ok := rates[id]; !ok {
			oid, _ := primitive.ObjectIDFromHex(id)
			oids = append(oids, oid)
		}
	}
	query := bson.M{"hotel_id": bson.M{"$in": oids}}
	cur, err := s.hotelRateDb.Find(ctx, query)
	if err != nil {
		s.logger.Errorw("Failed to find rates from MongoDB")
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// decode from cursor
	var ratesMongo []*HotelRate
	if err := cur.All(ctx, &ratesMongo); err != nil {
		s.logger.Errorw("Failed to decode rates from MongoDB")
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	for _, r := range ratesMongo {
		id := r.HotelOid.Hex()
		rates[id] = r.toProto()

		// update memcached
		rateJson, _ := protojson.Marshal(rates[id])
		err = s.memcached.WithContext(ctx).Set(&memcache.Item{
			Key:   id,
			Value: rateJson,
		})
		if err != nil {
			s.logger.Warnw("Failed to update rates to Memcached", "err", err)
		}
	}

	// still have unknown rates
	if len(rates) != len(req.HotelIds) {
		s.logger.Warn("Unknown hotel_id")
		return nil, status.Errorf(codes.NotFound, "Unknown hotel_id")
	}

	result := make([]*pb.HotelRate, 0, len(rates))
	for _, rate := range rates {
		result = append(result, rate)
	}
	return &pb.GetRatesRespond{Rates: result}, nil
}

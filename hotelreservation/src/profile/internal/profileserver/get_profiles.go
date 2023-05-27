package profileserver

import (
	"context"
	"encoding/json"
	"microless/hotelreservation/proto"
	pb "microless/hotelreservation/proto/profile"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProfileService) GetProfiles(ctx context.Context, req *pb.GetProfilesRequest) (*pb.GetProfilesRespond, error) {
	// return empty respond when request is empty
	if len(req.HotelIds) == 0 {
		return &pb.GetProfilesRespond{}, nil
	}

	hotels := make(map[string]*Hotel)

	// get hotels from memcached
	s.logger.Info("Get hotels from Memcached")
	hotelsMc, err := s.memcached.WithContext(ctx).GetMulti(req.HotelIds)
	if err != nil {
		s.logger.Warnw("Failed to get hotels from Memcached", "err", err)
	} else {
		for _, item := range hotelsMc {
			hotel := new(Hotel)
			json.Unmarshal(item.Value, hotel)
			hotels[item.Key] = hotel
		}
	}

	// get all hotel from memcached
	if len(hotels) == len(req.HotelIds) {
		hotelsPb := make([]*proto.Hotel, len(req.HotelIds))
		for i, id := range req.HotelIds {
			hotelsPb[i] = hotels[id].toProto()
		}
		return &pb.GetProfilesRespond{Hotels: hotelsPb}, nil
	}

	// get hotels from mongodb
	s.logger.Info("Get hotels from MongoDB")
	oids := make([]primitive.ObjectID, 0, len(req.HotelIds)-len(hotels))
	for _, id := range req.HotelIds {
		if _, ok := hotels[id]; !ok {
			oid, _ := primitive.ObjectIDFromHex(id)
			oids = append(oids, oid)
		}
	}
	query := bson.M{"_id": bson.M{"$in": oids}}
	cur, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Warnw("Failed to find hotels from MongoDB")
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// decode from cursor
	var hotelsMongo []*Hotel
	err = cur.All(ctx, &hotelsMongo)
	if err != nil {
		s.logger.Warnw("Failed to find hotels from MongoDB")
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	for _, h := range hotelsMongo {
		id := h.HotelId.Hex()
		hotels[id] = h

		// upload hotels to memcached
		hotelJson, _ := json.Marshal(h)
		err = s.memcached.WithContext(ctx).Set(&memcache.Item{
			Key:   id,
			Value: hotelJson,
		})
		if err != nil {
			s.logger.Warnw("Failed to update Memcahced", "hotel_id", id, "err", err)
		}
	}

	// still have unknown hotels
	if len(hotels) != len(req.HotelIds) {
		s.logger.Warn("Unknown hotel_id")
		return nil, status.Errorf(codes.NotFound, "Unknown hotel_id")
	}

	hotelsPb := make([]*proto.Hotel, len(req.HotelIds))
	for i, id := range req.HotelIds {
		hotelsPb[i] = hotels[id].toProto()
	}
	return &pb.GetProfilesRespond{Hotels: hotelsPb}, nil
}

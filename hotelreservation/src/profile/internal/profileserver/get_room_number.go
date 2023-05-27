package profileserver

import (
	"context"
	"encoding/json"
	pb "microless/hotelreservation/proto/profile"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProfileService) GetRoomNumber(ctx context.Context, req *pb.GetRoomNumberRequest) (*pb.GetRoomNumberRespond, error) {
	hotel := new(Hotel)

	// get hotel from memcached
	s.logger.Info("Get hotel from Memcached")
	item, err := s.memcached.WithContext(ctx).Get(req.HotelId)
	if err != nil {
		s.logger.Warnw("Failed to get hotel from Memcached", "err", err)
	} else {
		json.Unmarshal(item.Value, hotel)
		return &pb.GetRoomNumberRespond{RoomNumber: hotel.RoomNumber}, nil
	}

	// get hotel from mongodb
	s.logger.Info("Get hotel from MongoDB")
	hotelOid, _ := primitive.ObjectIDFromHex(req.HotelId)
	query := bson.M{"_id": hotelOid}
	err = s.mongodb.FindOne(ctx, query).Decode(hotel)
	if err != nil {
		s.logger.Warnw("Failed to find hotel from MongoDB")
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	return &pb.GetRoomNumberRespond{RoomNumber: hotel.RoomNumber}, nil
}

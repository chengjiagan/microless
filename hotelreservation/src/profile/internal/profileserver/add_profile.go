package profileserver

import (
	"context"
	"microless/hotelreservation/proto/geo"
	pb "microless/hotelreservation/proto/profile"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProfileService) AddProfile(ctx context.Context, req *pb.AddProfileRequest) (*pb.AddProfileRespond, error) {
	// insert new hotel profile
	s.logger.Info("Insert new hotel profile")
	hotel := hotelFromProto(req.Hotel)
	res, err := s.mongodb.InsertOne(ctx, hotel)
	if err != nil {
		s.logger.Errorw("Failed to insert new hotel profile to MongoDB", "error", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	hotelId := res.InsertedID.(primitive.ObjectID).Hex()

	// add hotel location to geo database
	s.logger.Info("Add hotel location to geo database")
	geoRequest := &geo.AddHotelRequest{
		HotelId: hotelId,
		Lat:     req.Hotel.Address.Lat,
		Lon:     req.Hotel.Address.Lon,
	}
	_, err = s.geoClient.AddHotel(ctx, geoRequest)
	if err != nil {
		s.logger.Errorw("Failed to add hotel location to geo database", "err", err)
		return nil, err
	}

	return &pb.AddProfileRespond{HotelId: hotelId}, nil
}

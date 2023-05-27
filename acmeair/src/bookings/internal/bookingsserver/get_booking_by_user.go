package bookingsserver

import (
	"context"
	"encoding/json"
	"fmt"
	"microless/acmeair/proto"
	pb "microless/acmeair/proto/bookings"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// get booking by user
func (s *BookingsService) GetBookingByUser(ctx context.Context, req *pb.GetBookingByUserRequest) (*pb.GetBookingByUserRespond, error) {
	var bookings []*Booking

	// get bookings from memcached
	s.logger.Info("Get bookings by user from memcached")
	key := fmt.Sprintf("user:%s", req.CustomerId)
	bookingsMc, err := s.memcached.WithContext(ctx).Get(key)
	if err == nil {
		// cache hit
		json.Unmarshal(bookingsMc.Value, &bookings)
		bookingsPb := make([]*proto.BookingInfo, len(bookings))
		for i, b := range bookings {
			bPb, err := s.getBookingInfoFromBooking(ctx, b)
			if err != nil {
				return nil, err
			}
			bookingsPb[i] = bPb
		}
		return &pb.GetBookingByUserRespond{Bookings: bookingsPb}, nil
	}

	// cache miss, get bookings from mongodb
	s.logger.Info("Get bookings by user from mongodb")
	customerOid, _ := primitive.ObjectIDFromHex(req.CustomerId)
	query := bson.M{"customer_id": customerOid}
	cur, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Warnw("Failed to get bookings from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	err = cur.All(ctx, &bookings)
	if err != nil {
		s.logger.Warnw("Failed to get bookings from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// update memcached
	s.logger.Info("Update bookings to memcached")
	bookingsJson, _ := json.Marshal(bookings)
	err = s.memcached.WithContext(ctx).Set(
		&memcache.Item{
			Key:   key,
			Value: bookingsJson,
		},
	)
	if err != nil {
		s.logger.Warnw("Failed to update bookings to memcached", "err", err)
	}

	bookingsPb := make([]*proto.BookingInfo, len(bookings))
	for i, b := range bookings {
		bPb, err := s.getBookingInfoFromBooking(ctx, b)
		if err != nil {
			return nil, err
		}
		bookingsPb[i] = bPb
	}
	return &pb.GetBookingByUserRespond{Bookings: bookingsPb}, nil
}

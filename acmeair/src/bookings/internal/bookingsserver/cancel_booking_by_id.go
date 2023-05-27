package bookingsserver

import (
	"context"
	"fmt"
	pb "microless/acmeair/proto/bookings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// cancel booking by number
func (s *BookingsService) CancelBookingById(ctx context.Context, req *pb.CancelBookingByIdRequest) (*emptypb.Empty, error) {
	// delete booking in mongodb
	s.logger.Info("Delete booking from MongoDB")
	bookingOid, _ := primitive.ObjectIDFromHex(req.BookingId)
	query := bson.M{"_id": bookingOid}
	res, err := s.mongodb.DeleteOne(ctx, query)
	if err != nil {
		s.logger.Warnw("Failed to delete booking from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Booking not found")
	}

	// delete booking by id in memcached
	s.logger.Info("Delete booking by number from memcached")
	key := fmt.Sprintf("booking:%s", req.BookingId)
	err = s.memcached.WithContext(ctx).Delete(key)
	if err != nil {
		s.logger.Warnw("Failed to delete booking from memcached", "err", err)
	}

	// delete booking by user in memcached
	s.logger.Info("Delete booking by user from memcached")
	key = fmt.Sprintf("user:%s", req.CustomerId)
	err = s.memcached.WithContext(ctx).Delete(key)
	if err != nil {
		s.logger.Warnw("Failed to delete booking from memcached", "err", err)
	}

	return &emptypb.Empty{}, nil
}

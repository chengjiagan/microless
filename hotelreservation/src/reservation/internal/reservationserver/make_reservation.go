package reservationserver

import (
	"context"
	"fmt"
	pb "microless/hotelreservation/proto/reservation"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ReservationService) MakeReservation(ctx context.Context, req *pb.MakeReservationRequest) (*emptypb.Empty, error) {
	inDate := req.InDate.AsTime()
	outDate := req.OutDate.AsTime()

	// check availability
	available, err := s.check(ctx, req.HotelId, inDate, outDate, int(req.RoomNumber))
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, status.Errorf(codes.FailedPrecondition, "Hotel %v is not available between %v and %v", req.HotelId, inDate, outDate)
	}

	// make reservation
	hotelOid, _ := primitive.ObjectIDFromHex(req.HotelId)
	userOid, _ := primitive.ObjectIDFromHex(req.UserId)
	reservation := &Reservation{
		HotelOid:   hotelOid,
		UserOid:    userOid,
		RoomNumber: int(req.RoomNumber),
		InDate:     inDate,
		OutDate:    outDate,
	}
	_, err = s.mongodb.InsertOne(ctx, reservation)
	if err != nil {
		s.logger.Errorw("Failed to insert reservation into MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// delete cache in memcached
	for day := inDate; day.Before(outDate); day = day.AddDate(0, 0, 1) {
		key := fmt.Sprintf("count:%s-%s", req.HotelId, day.Format("2006-01-02"))
		err := s.memcached.WithContext(ctx).Delete(key)
		if err != nil {
			s.logger.Warnw("Failed to delete cache in memcached", "err", err)
		}
	}

	return &emptypb.Empty{}, nil
}

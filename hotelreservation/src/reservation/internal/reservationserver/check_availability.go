package reservationserver

import (
	"context"
	"fmt"
	"microless/hotelreservation/proto/profile"
	pb "microless/hotelreservation/proto/reservation"
	"strconv"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ReservationService) CheckAvailability(ctx context.Context, req *pb.CheckAvailabilityRequest) (*pb.CheckAvailabilityRespond, error) {
	inDate := req.InDate.AsTime()
	outDate := req.OutDate.AsTime()
	availableHotelIds := make([]string, 0)

	// check every hotels
	for _, hotelId := range req.HotelIds {
		available, err := s.check(ctx, hotelId, inDate, outDate, int(req.RoomNumber))
		if err != nil {
			return nil, err
		}

		if available {
			availableHotelIds = append(availableHotelIds, hotelId)
		}
	}

	return &pb.CheckAvailabilityRespond{HotelIds: availableHotelIds}, nil
}

func (s *ReservationService) check(ctx context.Context, hotelId string, inDate time.Time, outDate time.Time, number int) (bool, error) {
	numberOfRoom, err := s.getNumberOfRoom(ctx, hotelId)
	if err != nil {
		return false, err
	}

	// check that there are enough rooms between inDate and outDate
	for day := inDate; day.Before(outDate); day = day.AddDate(0, 0, 1) {
		count, err := s.getReservationCount(ctx, hotelId, day)
		if err != nil {
			return false, err
		}

		// return false early if there are not enough rooms
		if count+number > numberOfRoom {
			return false, nil
		}
	}

	return true, nil
}

func (s *ReservationService) getNumberOfRoom(ctx context.Context, hotelId string) (int, error) {
	// get from profile-service
	s.logger.Info("Get number of room from profile-service")
	profileRequest := &profile.GetRoomNumberRequest{HotelId: hotelId}
	profileRespond, err := s.profileClient.GetRoomNumber(ctx, profileRequest)
	if err != nil {
		s.logger.Errorw("Failed to get number of room from profile-service", "err", err)
		return 0, err
	}
	return int(profileRespond.RoomNumber), nil
}

func (s *ReservationService) getReservationCount(ctx context.Context, hotelId string, date time.Time) (int, error) {
	// get from memcached
	s.logger.Info("Getting reservation count from memcached")
	key := fmt.Sprintf("count:%s-%s", hotelId, date.Format("2006-01-02"))
	resultMc, err := s.memcached.WithContext(ctx).Get(key)
	if err == nil {
		value, _ := strconv.Atoi(string(resultMc.Value))
		return value, nil
	} else if err != memcache.ErrCacheMiss {
		s.logger.Warnw("Failed to get reservation count from memcached", "err", err)
	}

	// get from MongoDB
	s.logger.Info("Getting reservation count from MongoDB")
	hotelOid, _ := primitive.ObjectIDFromHex(hotelId)
	query := bson.M{
		"hotel_id": hotelOid,
		"in_date":  bson.M{"$lte": date},
		"out_date": bson.M{"$gt": date},
	}
	cur, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Errorw("Failed to get reservations from MongoDB", "err", err)
		return 0, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	var reservations []*Reservation
	if err := cur.All(ctx, &reservations); err != nil {
		s.logger.Errorw("Failed to decode reservations from MongoDB", "err", err)
		return 0, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	count := 0
	for _, reservation := range reservations {
		count += reservation.RoomNumber
	}

	// update memcached
	s.logger.Info("Updating reservation count to memcached")
	err = s.memcached.WithContext(ctx).Set(&memcache.Item{
		Key:   key,
		Value: []byte(strconv.Itoa(count)),
	})
	if err != nil {
		s.logger.Warnw("Failed to update reservation count in memcached", "err", err)
	}

	return count, nil
}

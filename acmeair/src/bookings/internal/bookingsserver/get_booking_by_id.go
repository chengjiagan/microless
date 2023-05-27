package bookingsserver

import (
	"context"
	"encoding/json"
	"fmt"
	"microless/acmeair/proto"
	pb "microless/acmeair/proto/bookings"
	"microless/acmeair/proto/customer"
	"microless/acmeair/proto/flights"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// get booking by booking id
func (s *BookingsService) GetBookingById(ctx context.Context, req *pb.GetBookingByIdRequest) (*pb.GetBookingByIdRespond, error) {
	booking := new(Booking)
	// get booking from memcached
	s.logger.Info("Get booking from memcached")
	key := fmt.Sprintf("booking:%s", req.BookingId)
	bookingMc, err := s.memcached.WithContext(ctx).Get(key)
	if err == nil {
		// cache hit
		json.Unmarshal(bookingMc.Value, booking)
		bookingPb, err := s.getBookingInfoFromBooking(ctx, booking)
		if err != nil {
			return nil, err
		}
		return &pb.GetBookingByIdRespond{Booking: bookingPb}, nil
	} else if err != memcache.ErrCacheMiss {
		s.logger.Warnw("Failed to get booking from memcached", "err", err)
	}

	// cache miss, get booking from mongodb
	s.logger.Info("Get booking from MongoDB")
	bookingOid, _ := primitive.ObjectIDFromHex(req.BookingId)
	query := bson.M{"_id": bookingOid}
	err = s.mongodb.FindOne(ctx, query).Decode(booking)
	if err != nil {
		s.logger.Warnw("Failed to get booking from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	// update memcached
	s.logger.Info("Update booking to memcached")
	bookingJson, _ := json.Marshal(booking)
	err = s.memcached.WithContext(ctx).Set(
		&memcache.Item{
			Key:   key,
			Value: bookingJson,
		},
	)
	if err != nil {
		s.logger.Warnw("Failed to set booking to memcached", "err", err)
	}

	bookingPb, err := s.getBookingInfoFromBooking(ctx, booking)
	if err != nil {
		return nil, err
	}
	return &pb.GetBookingByIdRespond{Booking: bookingPb}, nil
}

func (s *BookingsService) getBookingInfoFromBooking(ctx context.Context, booking *Booking) (*proto.BookingInfo, error) {
	// get customer info from customer-service
	customerReq := &customer.GetCustomerRequest{
		CustomerId: booking.CustomerId.Hex(),
	}
	customerResp, err := s.customerClient.GetCustomer(ctx, customerReq)
	if err != nil {
		s.logger.Warnw("Failed to get customer info from customer-service", "err", err)
		return nil, err
	}

	// get to flight info from flights-service
	toFlightReq := &flights.GetFlightByIdRequest{
		FlightId: booking.ToFlightId.Hex(),
	}
	toFlightResp, err := s.flightsClient.GetFlightById(ctx, toFlightReq)
	if err != nil {
		s.logger.Warnw("Failed to get to flight info from flights-service", "err", err)
		return nil, err
	}

	if booking.OneWayFlight {
		return &proto.BookingInfo{
			BookingId:     booking.BookingId.Hex(),
			OneWayFlight:  booking.OneWayFlight,
			ToFlight:      toFlightResp.Flight,
			Customer:      customerResp.Customer,
			DateOfBooking: timestamppb.New(booking.BookingId.Timestamp()),
		}, nil
	}

	// get return flight info from flights-service
	retFlightReq := &flights.GetFlightByIdRequest{
		FlightId: booking.ReturnFlightId.Hex(),
	}
	retFlightResp, err := s.flightsClient.GetFlightById(ctx, retFlightReq)
	if err != nil {
		s.logger.Warnw("Failed to get return flight info from flights-service", "err", err)
		return nil, err
	}

	return &proto.BookingInfo{
		BookingId:     booking.BookingId.Hex(),
		OneWayFlight:  booking.OneWayFlight,
		ToFlight:      toFlightResp.Flight,
		ReturnFlight:  retFlightResp.Flight,
		Customer:      customerResp.Customer,
		DateOfBooking: timestamppb.New(booking.BookingId.Timestamp()),
	}, nil
}

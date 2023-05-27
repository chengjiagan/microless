package bookingsserver

import (
	"context"
	"microless/acmeair/proto"
	pb "microless/acmeair/proto/bookings"
	"microless/acmeair/proto/customer"
	"microless/acmeair/proto/flights"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *BookingsService) BookFlights(ctx context.Context, req *pb.BookFlightsRequest) (*pb.BookFlightsRespond, error) {
	// checking booking infomations
	// get customer info from customer-service
	customerReq := &customer.GetCustomerRequest{
		CustomerId: req.CustomerId,
	}
	customerResp, err := s.customerClient.GetCustomer(ctx, customerReq)
	if err != nil {
		s.logger.Warnw("Failed to get customer info from customer-service", "err", err)
		return nil, err
	}

	// get to flight info from flights-service
	toFlightReq := &flights.GetFlightByIdRequest{
		FlightId: req.ToFlightId,
	}
	toFlightResp, err := s.flightsClient.GetFlightById(ctx, toFlightReq)
	if err != nil {
		s.logger.Warnw("Failed to get to flight info from flights-service", "err", err)
		return nil, err
	}

	var retFlightResp *flights.GetFlightByIdRespond
	// check return flight if not one-way-flight
	if !req.OneWayFlight {
		// get return flight info from flights-service
		retFlightReq := &flights.GetFlightByIdRequest{
			FlightId: req.RetFlightId,
		}
		retFlightResp, err = s.flightsClient.GetFlightById(ctx, retFlightReq)
		if err != nil {
			s.logger.Warnw("Failed to get return flight info from flights-service", "err", err)
			return nil, err
		}
	}

	toFlightOid, _ := primitive.ObjectIDFromHex(req.ToFlightId)
	customerOid, _ := primitive.ObjectIDFromHex(req.CustomerId)
	booking := &Booking{
		OneWayFlight: req.OneWayFlight,
		ToFlightId:   toFlightOid,
		CustomerId:   customerOid,
	}
	// set return flight id if not one-way-flight
	if !req.OneWayFlight {
		booking.ReturnFlightId, _ = primitive.ObjectIDFromHex(req.RetFlightId)
	}

	// insert booking into mongodb
	res, err := s.mongodb.InsertOne(ctx, booking)
	if err != nil {
		s.logger.Warnw("Failed to insert booking", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	bookingOid := res.InsertedID.(primitive.ObjectID)

	bookingPb := &proto.BookingInfo{
		BookingId:     bookingOid.Hex(),
		OneWayFlight:  req.OneWayFlight,
		ToFlight:      toFlightResp.Flight,
		Customer:      customerResp.Customer,
		DateOfBooking: timestamppb.New(bookingOid.Timestamp()),
	}
	// set return flight info if not one-way-flight
	if !req.OneWayFlight {
		bookingPb.ReturnFlight = retFlightResp.Flight
	}
	return &pb.BookFlightsRespond{Booking: bookingPb}, nil
}

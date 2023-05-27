package searchserver

import (
	"context"
	"microless/hotelreservation/proto"
	"microless/hotelreservation/proto/geo"
	"microless/hotelreservation/proto/profile"
	"microless/hotelreservation/proto/rate"
	"microless/hotelreservation/proto/reservation"
	pb "microless/hotelreservation/proto/search"
	"sort"

	"golang.org/x/sync/errgroup"
)

func (s *SearchService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchRespond, error) {
	// get nearby hotels from geo-service
	geoRequest := &geo.NearbyRequest{
		Lat: req.Lat,
		Lon: req.Lon,
	}
	geoRespond, err := s.geoClient.Nearby(ctx, geoRequest)
	if err != nil {
		s.logger.Warnw("Failed to get nearby hotels from geo-service", "err", err)
		return nil, err
	}

	// filter out unavailable hotels
	reservationRequest := &reservation.CheckAvailabilityRequest{
		HotelIds:   geoRespond.HotelIds,
		InDate:     req.InDate,
		OutDate:    req.OutDate,
		RoomNumber: req.RoomNumber,
	}
	reservationRespond, err := s.reservationClient.CheckAvailability(ctx, reservationRequest)
	if err != nil {
		s.logger.Warnw("Failed to check availability from reservation-service", "err", err)
		return nil, err
	}

	g, ctx := errgroup.WithContext(ctx)

	// sort hotels by rate
	hotelIds := make([]string, len(reservationRespond.HotelIds))
	g.Go(func() error {
		// get rates from rate-service
		rateRequest := &rate.GetRatesRequest{HotelIds: reservationRespond.HotelIds}
		rateRespond, err := s.rateClient.GetRates(ctx, rateRequest)
		if err != nil {
			s.logger.Warnw("Failed to get rates from rate-service", "err", err)
			return err
		}

		// sort rates in descending order
		sort.Sort(byRate(rateRespond.Rates))
		for i, rate := range rateRespond.Rates {
			hotelIds[i] = rate.HotelId
		}

		return nil
	})

	// get hotel profiles from profile-service
	profiles := make(map[string]*proto.Hotel, len(reservationRespond.HotelIds))
	g.Go(func() error {
		profileRequest := &profile.GetProfilesRequest{HotelIds: reservationRespond.HotelIds}
		profileRespond, err := s.profileClient.GetProfiles(ctx, profileRequest)
		if err != nil {
			s.logger.Warnw("Failed to get hotels from profile-service", "err", err)
			return err
		}

		for _, hotel := range profileRespond.Hotels {
			profiles[hotel.HotelId] = hotel
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// build respond
	hotels := make([]*proto.Hotel, len(hotelIds))
	for i, hotelId := range hotelIds {
		hotels[i] = profiles[hotelId]
	}
	return &pb.SearchRespond{Hotels: hotels}, nil
}

// for sorting in descending order
type byRate []*rate.HotelRate

func (a byRate) Len() int           { return len(a) }
func (a byRate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byRate) Less(i, j int) bool { return a[i].Rate > a[j].Rate }

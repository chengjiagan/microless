package composereviewserver

import (
	"context"
	pb "microless/media/proto/composereview"
	"microless/media/proto/moviereview"
	"microless/media/proto/rating"
	"microless/media/proto/reviewstorage"
	"microless/media/proto/userreview"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ComposeReviewService) ComposeReview(ctx context.Context, req *pb.ComposeReviewRequest) (*emptypb.Empty, error) {
	// store review in review-storage
	s.logger.Info("Store review in review-storage-service")
	reviewReq := &reviewstorage.StoreReviewRequest{
		UserId:  req.UserId,
		MovieId: req.MovieId,
		Text:    req.Text,
		Rating:  req.Rating,
	}
	reviewResp, err := s.reviewstorageClient.StoreReview(ctx, reviewReq)
	if err != nil {
		s.logger.Errorw("Failed to store review to review-storage-service", "err", err)
		return nil, err
	}

	g, ctx := errgroup.WithContext(ctx)
	// add review to user reviews
	g.Go(func() error {
		s.logger.Info("Add review to user-review-service")
		userReq := &userreview.UploadUserReviewRequest{
			UserId:    req.UserId,
			ReviewId:  reviewResp.ReviewId,
			Timestamp: reviewResp.Timestamp,
		}
		_, err := s.userreviewClient.UploadUserReview(ctx, userReq)
		if err != nil {
			s.logger.Errorw("Failed to add review to user-review-service", "err", err)
			return err
		}
		return nil
	})

	// add review to movie reviews
	g.Go(func() error {
		s.logger.Info("Add review to movie-review-service")
		movieReq := &moviereview.UploadMovieReviewRequest{
			MovieId:   req.MovieId,
			ReviewId:  reviewResp.ReviewId,
			Timestamp: reviewResp.Timestamp,
		}
		_, err := s.moviereviewClient.UploadMovieReview(ctx, movieReq)
		if err != nil {
			s.logger.Errorw("Failed to add review to movie-review-service", "err", err)
			return err
		}
		return nil
	})

	// upload rating to rating service
	g.Go(func() error {
		s.logger.Info("Update movie rating in movie-info-service")
		ratingReq := &rating.UploadRatingRequest{
			MovieId: req.MovieId,
			Rating:  req.Rating,
		}
		_, err := s.ratingClient.UploadRating(ctx, ratingReq)
		if err != nil {
			s.logger.Errorw("Failed to upload rating in rating-service", "err", err)
			return err
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

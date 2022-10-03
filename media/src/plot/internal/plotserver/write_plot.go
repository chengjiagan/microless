package plotserver

import (
	"context"

	pb "microless/media/proto/plot"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlotService) WritePlot(ctx context.Context, req *pb.WritePlotRequest) (*pb.WritePlotRespond, error) {
	plot := &Plot{Plot: req.Plot}
	result, err := s.mongodb.InsertOne(ctx, plot)
	if err != nil {
		s.logger.Errorw("Failed to insert plot into MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	s.logger.Info("Insert new plot")
	oid := result.InsertedID.(primitive.ObjectID)
	return &pb.WritePlotRespond{PlotId: oid.Hex()}, nil
}

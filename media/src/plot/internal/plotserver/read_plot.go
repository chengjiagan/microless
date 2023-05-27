package plotserver

import (
	"context"
	"encoding/json"

	pb "microless/media/proto/plot"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *PlotService) ReadPlot(ctx context.Context, req *pb.ReadPlotRequest) (*pb.ReadPlotRespond, error) {
	plot := new(Plot)

	// get plot from memcached
	s.logger.Info("Read plot from Memcached")
	result, err := s.memcached.WithContext(ctx).Get(req.PlotId)
	if err != nil {
		s.logger.Warnw("Failed to get plot from memcached", "plot_id", req.PlotId, "err", err)
	} else {
		json.Unmarshal(result.Value, plot)
		return &pb.ReadPlotRespond{Plot: plot.Plot}, nil
	}

	// cache miss, get plot from mongodb
	s.logger.Info("Read plot from MongoDB")
	oid, _ := primitive.ObjectIDFromHex(req.PlotId)
	query := bson.M{"_id": oid}
	err = s.mongodb.FindOne(ctx, query).Decode(plot)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.logger.Warnw("Plot not found", "plot_id", req.PlotId)
			return nil, status.Errorf(codes.NotFound, "Plot %v not found", req.PlotId)
		} else {
			s.logger.Warnw("Failed to get plot from MongoDB", "plot_id", req.PlotId, "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}

	// update memcached
	s.logger.Info("Update plot in Memcached")
	plotJson, _ := json.Marshal(plot)
	err = s.memcached.Set(&memcache.Item{
		Key:   req.PlotId,
		Value: plotJson,
	})
	if err != nil {
		s.logger.Warnw("Failed to update plot in Memcached", "plot_id", req.PlotId, "err", err)
	}

	return &pb.ReadPlotRespond{Plot: plot.Plot}, nil
}

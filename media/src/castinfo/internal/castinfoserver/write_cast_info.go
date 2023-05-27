package castinfoserver

import (
	"context"
	pb "microless/media/proto/castinfo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CastInfoService) WriteCastInfo(ctx context.Context, req *pb.WriteCastInfoRequest) (*pb.WriteCastInfoRespond, error) {
	info := &CastInfo{
		Name:   req.Name,
		Gender: req.Gender,
		Intro:  req.Intro,
	}

	result, err := s.mongodb.InsertOne(ctx, info)
	if err != nil {
		s.logger.Warnw("Failed to insert cast info to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}

	s.logger.Info("Insert new cast info")
	oid := result.InsertedID.(primitive.ObjectID)
	return &pb.WriteCastInfoRespond{CastInfoId: oid.Hex()}, nil
}

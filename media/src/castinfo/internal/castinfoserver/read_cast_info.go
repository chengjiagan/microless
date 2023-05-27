package castinfoserver

import (
	"context"
	"encoding/json"
	"microless/media/proto"
	pb "microless/media/proto/castinfo"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CastInfoService) ReadCastInfo(ctx context.Context, req *pb.ReadCastInfoRequest) (*pb.ReadCastInfoRespond, error) {
	if len(req.CastIds) == 0 {
		s.logger.Info("Request no CastId")
		return &pb.ReadCastInfoRespond{}, nil
	}

	// get cast infos from memcached
	s.logger.Info("Read CastInfos from Memcached")
	infos := make(map[string]*CastInfo, len(req.CastIds))
	infosMc, err := s.memcached.WithContext(ctx).GetMulti(req.CastIds)
	if err != nil {
		s.logger.Warnw("Failed to get cast infos from Memcached", "cast_info_ids", req.CastIds, "err", err)
	} else {
		for k, v := range infosMc {
			info := new(CastInfo)
			json.Unmarshal(v.Value, info)
			infos[k] = info
		}
	}

	// got all cast info from memcached
	if len(infos) == len(req.CastIds) {
		pbInfos := make([]*proto.CastInfo, len(infos))
		for i, id := range req.CastIds {
			pbInfos[i] = infos[id].toProto()
		}
		return &pb.ReadCastInfoRespond{CastInfos: pbInfos}, nil
	}

	// get cast infos from mongodb
	s.logger.Info("Read CastInfos from MongoDB")
	oids := make([]primitive.ObjectID, 0, len(req.CastIds)-len(infos))
	for _, id := range req.CastIds {
		if _, ok := infos[id]; !ok {
			oid, _ := primitive.ObjectIDFromHex(id)
			oids = append(oids, oid)
		}
	}
	query := bson.M{"_id": bson.M{"$in": oids}}
	cursor, err := s.mongodb.Find(ctx, query)
	if err != nil {
		s.logger.Warnw("Failed to find CastInfo from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	// decode from mongodb
	var infosMongo []*CastInfo
	err = cursor.All(ctx, &infosMongo)
	if err != nil {
		s.logger.Warnw("Failed to find CastInfo from MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	for _, info := range infosMongo {
		id := info.CastInfoOid.Hex()
		infos[id] = info

		// upload infos to memcached
		infoJson, _ := json.Marshal(info)
		err = s.memcached.WithContext(ctx).Set(&memcache.Item{
			Key:   id,
			Value: infoJson,
		})
		if err != nil {
			s.logger.Warnw("Failed to update Memcached", "cast_info_id", id, "err", err)
		}
	}

	// still unknown cast_info_id exists
	if len(infos) != len(req.CastIds) {
		s.logger.Warn("Unknown cast_info_id")
		return nil, status.Error(codes.NotFound, "Unknown cast_info_id")
	}

	pbInfos := make([]*proto.CastInfo, len(infos))
	for i, id := range req.CastIds {
		pbInfos[i] = infos[id].toProto()
	}
	return &pb.ReadCastInfoRespond{CastInfos: pbInfos}, nil
}

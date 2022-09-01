package userserver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"

	"microless/socialnetwork/proto/socialgraph"
	pb "microless/socialnetwork/proto/user"
	"microless/socialnetwork/proto/usertimeline"
	"microless/socialnetwork/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserRespond, error) {
	count, err := s.mongodb.CountDocuments(ctx, bson.M{"username": req.Username})
	if err != nil {
		s.logger.Errorw("Failed to count users", "err", err, "username", req.Username)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	if count != 0 {
		s.logger.Debugw("Register existed user", "username", req.Username)
		return nil, status.Errorf(codes.FailedPrecondition, "Register existed user with username %v", req.Username)
	}

	// insert the new user into mongodb
	salt := utils.GetRandString(32)
	h := sha256.New()
	io.WriteString(h, req.Password+salt)
	hashedPswd := hex.EncodeToString(h.Sum(nil))
	user := User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Salt:      salt,
		Password:  hashedPswd,
	}
	result, err := s.mongodb.InsertOne(ctx, user)
	if err != nil {
		s.logger.Errorw("Failed to insert new user", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	s.logger.Debugw("Insert new user", "username", req.Username)
	userId := result.InsertedID.(primitive.ObjectID).Hex()

	// insert new user in social graph and user timeline
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		req := &socialgraph.InsertUserRequest{UserId: userId}
		_, err := s.socialgraphClient.InsertUser(ctx, req)
		if err != nil {
			s.logger.Errorw("Failed to insert user in social graph", "err", err)
			return err
		}
		return nil
	})
	g.Go(func() error {
		req := &usertimeline.InsertUserResquest{UserId: userId}
		_, err := s.usertimelineClient.InsertUser(ctx, req)
		if err != nil {
			s.logger.Errorw("Failed to insert user in user timeline", "err", err)
			return err
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserRespond{UserId: userId}, nil
}

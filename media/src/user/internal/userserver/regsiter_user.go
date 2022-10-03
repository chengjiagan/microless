package userserver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"

	pb "microless/media/proto/user"
	"microless/media/proto/userreview"
	"microless/media/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	s.logger.Info("Insert new user")
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
		s.logger.Errorw("Failed to insert new user to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	userId := result.InsertedID.(primitive.ObjectID).Hex()

	// create user in user-review-service
	s.logger.Info("Create user in user-review-service")
	reviewReq := &userreview.CreateUserRequest{UserId: userId}
	_, err = s.userreviewClient.CreateUser(ctx, reviewReq)
	if err != nil {
		s.logger.Errorw("Failed to create user in user-review-service", "err", err)
		return nil, err
	}

	return &pb.RegisterUserRespond{UserId: userId}, nil
}

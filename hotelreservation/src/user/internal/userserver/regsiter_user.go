package userserver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"

	pb "microless/hotelreservation/proto/user"
	"microless/hotelreservation/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserRespond, error) {
	count, err := s.mongodb.CountDocuments(ctx, bson.M{"username": req.Username})
	if err != nil {
		s.logger.Warnw("Failed to count users", "err", err, "username", req.Username)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	if count != 0 {
		s.logger.Debugw("Register existed user", "username", req.Username)
		return nil, status.Errorf(codes.FailedPrecondition, "Register existed user with username %v", req.Username)
	}

	// calculate the hashed password
	salt := utils.GetRandString(32)
	h := sha256.New()
	io.WriteString(h, req.Password+salt)
	hashedPswd := hex.EncodeToString(h.Sum(nil))

	// insert the new user into mongodb
	s.logger.Info("Insert new user")
	user := User{
		Username: req.Username,
		Salt:     salt,
		Password: hashedPswd,
	}
	result, err := s.mongodb.InsertOne(ctx, user)
	if err != nil {
		s.logger.Warnw("Failed to insert new user to MongoDB", "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
	}
	userId := result.InsertedID.(primitive.ObjectID).Hex()

	return &pb.RegisterUserRespond{UserId: userId}, nil
}

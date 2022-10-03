package userserver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"strconv"
	"time"

	pb "microless/media/proto/user"

	"github.com/bradfitz/gomemcache/memcache"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginRespond, error) {
	user, err := s.getUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	// check password
	h := sha256.New()
	io.WriteString(h, req.Password+user.Salt)
	hashedPswd := hex.EncodeToString(h.Sum(nil))
	if hashedPswd != user.Password {
		return nil, status.Errorf(codes.Unauthenticated, "Incorrect username or password")
	}

	// generate login token
	claims := jwt.MapClaims{
		"user_id":   user.UserOid.Hex(),
		"username":  req.Username,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		"ttl":       "3600",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString([]byte(s.secret))
	return &pb.LoginRespond{Token: tokenStr}, nil
}

func (s *UserService) getUser(ctx context.Context, username string) (*User, error) {
	user := new(User)

	keyMc := username + ":login"
	userMc, err := s.memcached.WithContext(ctx).Get(keyMc)
	if err != nil && err != memcache.ErrCacheMiss {
		s.logger.Errorw("Failed to get from Memcached", "err", err)
	}
	if userMc != nil {
		s.logger.Debugw("User cache hit from memcached", "username", username)
		json.Unmarshal(userMc.Value, user)
		return user, nil
	}

	// cache miss
	s.logger.Debugw("user_id cache miss from memcached", "username", username)
	err = s.mongodb.FindOne(ctx, bson.M{"username": username}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			s.logger.Warnw("User doesn't exist in MongoDB", "username", username)
			return nil, status.Errorf(codes.NotFound, "username: %v doesn't exist in MongoDB", username)
		} else {
			s.logger.Errorw("Failed to find user from MongoDB", "err", err)
			return nil, status.Errorf(codes.Internal, "MongoDB Err: %v", err)
		}
	}

	s.logger.Debugw("User found in MongoDB", "username", username)
	userJson, _ := json.Marshal(user)
	err = s.memcached.WithContext(ctx).Set(&memcache.Item{
		Key:   keyMc,
		Value: userJson,
	})
	if err != nil {
		s.logger.Errorw("Failed to set to Memcached", "err", err)
	}

	return user, nil
}

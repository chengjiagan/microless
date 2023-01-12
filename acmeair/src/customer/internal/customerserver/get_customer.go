package customerserver

import (
	"context"
	"encoding/json"
	pb "microless/acmeair/proto/customer"

	"github.com/bradfitz/gomemcache/memcache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CustomerService) GetCustomer(ctx context.Context, req *pb.GetCustomerRequest) (*pb.GetCustomerRespond, error) {
	customer := new(Customer)

	// get customerMc from memcached
	s.logger.Info("Get customer from Memcached")
	customerMc, err := s.memcached.WithContext(ctx).Get(req.CustomerId)
	if err == nil {
		// cache hit
		json.Unmarshal(customerMc.Value, customer)
		return &pb.GetCustomerRespond{Customer: customer.toProto()}, nil
	} else if err != memcache.ErrCacheMiss {
		s.logger.Warnw("Failed to get customer from Memcached", "customer_id", req.CustomerId, "err", err)
	}

	// cache miss, get customer from mongodb
	s.logger.Info("Get customer from MongoDB")
	customerOid, _ := primitive.ObjectIDFromHex(req.CustomerId)
	query := bson.M{"_id": customerOid}
	err = s.mongodb.FindOne(ctx, query).Decode(customer)
	if err != nil {
		s.logger.Errorw("Failed to get customer from MongoDB", "customer_id", req.CustomerId, "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB err: %v", err)
	}

	// update memcached
	s.logger.Info("Update customer to Memcached")
	customerJson, _ := json.Marshal(customer)
	err = s.memcached.WithContext(ctx).Set(
		&memcache.Item{
			Key:   req.CustomerId,
			Value: customerJson,
		},
	)
	if err != nil {
		s.logger.Warnw("Failed to update customer to Memcached", "customer_id", req.CustomerId, "err", err)
	}

	return &pb.GetCustomerRespond{Customer: customer.toProto()}, nil
}

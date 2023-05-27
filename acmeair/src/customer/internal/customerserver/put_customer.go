package customerserver

import (
	"context"
	pb "microless/acmeair/proto/customer"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *CustomerService) PutCustomer(ctx context.Context, req *pb.PutCustomerRequest) (*emptypb.Empty, error) {
	customerOid, _ := primitive.ObjectIDFromHex(req.CustomerId)
	customer := CustomerFromProto(req.Customer)

	// update mongodb
	s.logger.Info("Update customer to MongoDB")
	query := bson.M{"_id": customerOid}
	res, err := s.mongodb.ReplaceOne(ctx, query, customer)
	if err != nil {
		s.logger.Warnw("Failed to update customer to MongoDB", "customer_id", req.CustomerId, "err", err)
		return nil, status.Errorf(codes.Internal, "MongoDB err: %v", err)
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Customer %v not found", req.CustomerId)
	}

	// delete old customer from memcached
	s.logger.Info("Delete customer from Memcached")
	err = s.memcached.WithContext(ctx).Delete(req.CustomerId)
	if err != nil {
		s.logger.Warnw("Failed to delete customer from Memcached", "customer_id", req.CustomerId, "err", err)
	}

	return &emptypb.Empty{}, nil
}

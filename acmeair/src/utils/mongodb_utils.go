package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.uber.org/zap"
)

func NewMongodbClient(ctx context.Context, url string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(url).SetMonitor(otelmongo.NewMonitor())
	mongodb, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return mongodb, nil
}

func CreateGeoIndex(ctx context.Context, mongodb *mongo.Collection, key string) error {
	exist, err := checkIndex(ctx, mongodb, key)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	model := mongo.IndexModel{
		Keys:    bson.M{key: "2dsphere"},
		Options: options.Index().SetName(key),
	}
	_, err = mongodb.Indexes().CreateOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func CreateIndex(ctx context.Context, mongodb *mongo.Collection, key string, unique bool) error {
	exist, err := checkIndex(ctx, mongodb, key)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	model := mongo.IndexModel{
		Keys:    bson.M{key: 1},
		Options: options.Index().SetName(key).SetUnique(unique),
	}
	_, err = mongodb.Indexes().CreateOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

func checkIndex(ctx context.Context, mongodb *mongo.Collection, key string) (bool, error) {
	index := mongodb.Indexes()
	indexSpecs, err := index.ListSpecifications(ctx)
	if err != nil {
		return false, err
	}

	for _, spec := range indexSpecs {
		if spec.Name == key {
			return true, nil
		}
	}

	return false, nil
}

func ShutdownMongodb(mongodb *mongo.Client, ctx context.Context, logger *zap.Logger) {
	// Do not make the application hang when it is shutdown.
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	// Disconnect from MongoDB
	if err := mongodb.Disconnect(ctx); err != nil {
		logger.Fatal(err.Error())
	}
}

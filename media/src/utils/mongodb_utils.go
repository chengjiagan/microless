package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func NewMongodbClient(ctx context.Context, url string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(url).SetMonitor(otelmongo.NewMonitor())
	mongodb, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return mongodb, nil
}

func CreateIndex(ctx context.Context, mongodb *mongo.Collection, key string) error {
	index := mongodb.Indexes()
	name := fmt.Sprintf("%s_1", key)
	indexSpecs, err := index.ListSpecifications(ctx)
	if err == nil {
		for _, spec := range indexSpecs {
			if spec.Name == name {
				// index have been created, so just skip creating
				return nil
			}
		}
	}

	model := mongo.IndexModel{
		Keys:    bson.M{key: 1},
		Options: options.Index().SetName(name).SetUnique(true),
	}
	_, err = index.CreateOne(ctx, model)
	if err != nil {
		return err
	}

	return nil
}

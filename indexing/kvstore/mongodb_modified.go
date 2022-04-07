package kvstore

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	collection *mongo.Collection
	client     *mongo.Client
	db         *mongo.Database
}

// NewMongoDB TODO 实现mongoDB连接
func NewMongoDB(path string, db string, collection string) (*MongoDB, error) {
	client, err := mongo.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	database := client.Database(db)
	coll := database.Collection(collection)
	return &MongoDB{collection: coll, client: client, db: database}, nil
}

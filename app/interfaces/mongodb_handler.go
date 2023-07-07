package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBHandler belongs to the inteface layer
type MongoDBHandler interface {
	MongoDatabase() *mongo.Database
	Collection(name string, opts ...*options.CollectionOptions) (collection *mongo.Collection)
}

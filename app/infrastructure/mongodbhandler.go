package infrastructure

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBHandler belong to the infrastructure layer.
type MongoDBHandler struct {
	*mongo.Client
}

// NewMongoDBHandler returns connection and methods which is related to database handling.
func NewMongoDBHandler() (*MongoDBHandler, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_DATABASE_URL")))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	MongoDBHandler := MongoDBHandler{Client: client}

	return &MongoDBHandler, nil
}

// MongoDatabase is Mongo Database wrapper
func (mh *MongoDBHandler) MongoDatabase() *mongo.Database {
	return mh.Client.Database(os.Getenv("MONGO_DATABASE_NAME"))
}

// Collection is Mongo Collection wrapper
func (mh *MongoDBHandler) Collection(name string, opts ...*options.CollectionOptions) (collection *mongo.Collection) {
	return mh.MongoDatabase().Collection(name, opts...)
}

package interfaces

import (
	"context"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const positions_collection string = "positions"

// A PositionMongoRepository belong to the inteface layer
type PositionMongoRepository struct {
	MongoDBHandler MongoDBHandler
}

// FindAll is returns the number of entities.
func (pr *PositionMongoRepository) FindAll() (positions domain.Positions, err error) {
	positionsCollection := pr.MongoDBHandler.Collection(positions_collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := positionsCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(ctx, &positions); err != nil {
		panic(err)
	}

	return
}

// FindByID returns the entity identified by the given id.
func (pr *PositionMongoRepository) FindByID(positionID string) (position domain.Position, err error) {
	objectID, err := primitive.ObjectIDFromHex(positionID)

	if err != nil {
		return
	}

	result := pr.MongoDBHandler.Collection(positions_collection).FindOne(context.Background(), bson.M{"_id": objectID})

	result.Decode(&position)

	return
}

// Save is saves the given entity
func (pr *PositionMongoRepository) Save(p domain.Position) (err error) {
	positionsCollection := pr.MongoDBHandler.Collection(positions_collection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	currentTimestamp := time.Now()

	_, err = positionsCollection.InsertOne(ctx, bson.D{
		{Key: "name", Value: p.Name},
		{Key: "salary", Value: p.Salary},
		{Key: "created_at", Value: currentTimestamp},
		{Key: "updated_at", Value: currentTimestamp},
	})

	if err != nil {
		return
	}

	return
}

// DeleteByID is deletes the entity identified by the given id.
func (pr *PositionMongoRepository) DeleteByID(positionID string) (err error) {
	objectID, err := primitive.ObjectIDFromHex(positionID)

	if err != nil {
		return
	}

	_, err = pr.MongoDBHandler.Collection(positions_collection).DeleteOne(context.Background(), bson.M{"_id": objectID})

	if err != nil {
		return
	}

	return
}

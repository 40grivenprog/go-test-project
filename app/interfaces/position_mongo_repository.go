package interfaces

import (
	"context"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const positionsCollection string = "positions"

// A PositionMongoRepository belong to the inteface layer
type PositionMongoRepository struct {
	MongoDBHandler MongoDBHandler
}

// FindAll is returns the number of entities.
func (pr *PositionMongoRepository) FindAll() (positions domain.Positions, err error) {
	positionsCollection := pr.MongoDBHandler.Collection(positionsCollection)
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
	positionIDHex, err := primitive.ObjectIDFromHex(positionID)

	if err != nil {
		err = NewBadRequestError("position id", positionID)
		return
	}

	result := pr.MongoDBHandler.Collection(positionsCollection).FindOne(context.Background(), bson.M{"_id": positionIDHex})

	if err = result.Decode(&position); err != nil {
		err = NewRecordNotFoundError(positionIDHex)
		return
	}

	return
}

// Save is saves the given entity
func (pr *PositionMongoRepository) Save(p domain.Position) (err error) {
	positionsCollection := pr.MongoDBHandler.Collection(positionsCollection)
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
	positionIDHex, err := primitive.ObjectIDFromHex(positionID)

	if err != nil {
		err = NewBadRequestError("position id", positionID)
		return
	}

	result, err := pr.MongoDBHandler.Collection(positionsCollection).DeleteOne(context.Background(), bson.M{"_id": positionIDHex})

	if err != nil {
		return
	}

	if result.DeletedCount == 0 {
		err = NewRecordNotFoundError(positionIDHex)
		return
	}

	return
}

package interfaces

import (
	"context"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// An EmployeeMongoRepository belong to the inteface layer
type EmployeeMongoRepository struct {
	MongoDBHandler MongoDBHandler
}

const employeesCollection string = "employees"

// FindAllByPositionID returns all entities by position id.
func (er *EmployeeMongoRepository) FindAllByPositionID(positionID string) (employees domain.Employees, err error) {
	positionIDHex, err := primitive.ObjectIDFromHex(positionID)

	if err != nil {
		err = NewBadRequestError("position id", positionID)
		return
	}

	positionsCollection := er.MongoDBHandler.Collection(employeesCollection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := positionsCollection.Find(context.TODO(), bson.M{"position": positionIDHex})

	if err != nil {
		return
	}

	defer cursor.Close(context.TODO())

	if !cursor.Next(ctx) {
		err = NewRecordNotFoundError(positionIDHex)
		return
	}

	if err = cursor.All(ctx, &employees); err != nil {
		return
	}

	return
}

// FindByID returns the entity identified by the given id.
func (er *EmployeeMongoRepository) FindByID(employeeID string) (employee domain.Employee, err error) {
	employeeIDHex, err := primitive.ObjectIDFromHex(employeeID)

	if err != nil {
		err = NewBadRequestError("employee id", employeeID)
		return
	}

	result := er.MongoDBHandler.Collection(employeesCollection).FindOne(context.Background(), bson.M{"_id": employeeIDHex})

	if err = result.Decode(&employee); err != nil {
		err = NewRecordNotFoundError(employeeIDHex)
		return
	}

	return
}

// DeleteByID is deletes the entity identified by the given id.
func (er *EmployeeMongoRepository) DeleteByID(employeeID string) (err error) {
	employeeIDHex, err := primitive.ObjectIDFromHex(employeeID)

	if err != nil {
		err = NewBadRequestError("employee id", employeeID)
		return
	}

	result, err := er.MongoDBHandler.Collection(employeesCollection).DeleteOne(context.Background(), bson.M{"_id": employeeIDHex})

	if err != nil {
		return
	}

	if result.DeletedCount == 0 {
		err = NewRecordNotFoundError(employeeIDHex)
		return
	}

	return
}

// Save is saves the given entity
func (er *EmployeeMongoRepository) Save(e domain.Employee) (err error) {
	positionIDStr, ok := e.PositionID.(string)

	if !ok {
		err = NewBadRequestError("employee position", e.PositionID)
		return
	}

	positionIDHex, err := primitive.ObjectIDFromHex(positionIDStr)

	if err != nil {
		err = NewBadRequestError("employee position", positionIDStr)
		return
	}

	employeesCollection := er.MongoDBHandler.Collection(employeesCollection)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	currentTimestamp := time.Now()

	_, err = employeesCollection.InsertOne(ctx, bson.D{
		{Key: "first_name", Value: e.FirstName},
		{Key: "last_name", Value: e.LastName},
		{Key: "position", Value: positionIDHex},
		{Key: "created_at", Value: currentTimestamp},
		{Key: "updated_at", Value: currentTimestamp},
	})

	if err != nil {
		return
	}

	return
}

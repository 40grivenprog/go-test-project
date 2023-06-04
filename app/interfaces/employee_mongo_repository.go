package interfaces

import (
	"context"
	"errors"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// An EmployeeMongoRepository belong to the inteface layer
type EmployeeMongoRepository struct {
	MongoDBHandler MongoDBHandler
}

const employees_collection string = "employees"

// FindAllByPositionID returns all entities by position id.
func (er *EmployeeMongoRepository) FindAllByPositionID(positionID string) (employees domain.Employees, err error) {

	return
}

// FindByID returns the entity identified by the given id.
func (er *EmployeeMongoRepository) FindByID(employeeID string) (employee domain.Employee, err error) {


	return
}

// DeleteByID is deletes the entity identified by the given id.
func (er *EmployeeMongoRepository) DeleteByID(employeeID string) (err error) {
	return
}

// Save is saves the given entity
func (er *EmployeeMongoRepository) Save(e domain.Employee) (err error) {
	positionIDStr, ok := e.PositionID.(string)
	if !ok {
		return errors.New("Invalid Position Id type")
	}

	positionIDHex, err := primitive.ObjectIDFromHex(positionIDStr)

	if err != nil {
		return
	}

	employeesCollection := er.MongoDBHandler.Collection(employees_collection)
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

package interfaces

import (
	"context"
	"time"

	"github.com/bmf-san/go-clean-architecture-web-application-boilerplate/app/domain"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// UserMongoRepository belongs to interfaces layer
type UserMongoRepository struct {
	MongoDBHandler MongoDBHandler
}

const usersCollection string = "users"

// Save creates new record of User
func (ur *UserMongoRepository) Save(user domain.User) (err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	user.Password = string(hashedPassword)

	usersCollection := ur.MongoDBHandler.Collection(usersCollection)
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	currentTimestamp := time.Now()

	_, err = usersCollection.InsertOne(ctx, bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
		{Key: "created_at", Value: currentTimestamp},
		{Key: "updated_at", Value: currentTimestamp},
	})

	if err != nil {
		return
	}

	return
}

// FindByEmail returns user by email
func (ur *UserMongoRepository) FindByEmail(email string) (user domain.User, err error) {
	result := ur.MongoDBHandler.Collection(usersCollection).FindOne(context.Background(), bson.M{"email": email})

	if err = result.Decode(&user); err != nil {
		err = NewRecordNotFoundError(email)
		return
	}

	return
}

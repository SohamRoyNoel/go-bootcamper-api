package helpers

import (
	"context"
	"fmt"

	"bootcamp.com/server/services"

	// "github.com/go-playground/validator/v10"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var usersCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
// var validate = validator.New()

func RegistrationHelper(email string, password string, ctx context.Context, cancel context.CancelFunc) error {
	checkIfUserExists := services.CheckIfUserExists(email, ctx, cancel)
	fmt.Println("Data From Helper >>>>> ", checkIfUserExists)
	return checkIfUserExists
}

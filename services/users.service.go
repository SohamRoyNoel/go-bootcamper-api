package services

import (
	"context"
	"fmt"

	"bootcamp.com/server/database"
	"bootcamp.com/server/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var usersCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var userModel models.Users

// Check if users exists in DB
func CheckIfUserExists(emailId string, ctx context.Context, cancel context.CancelFunc)  error {
	userExists := usersCollection.FindOne(ctx, bson.M{"email": "billi@gmail.com"}).Decode(&userModel)
	fmt.Println("Data From Services >>>>> ", userExists);
	return userExists
}
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstace() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("warning: unable to found .env file")
	}

	MongoDb := os.Getenv("MONGODB_URI")
	if MongoDb == "" {
		log.Fatal("MONGODB_URI not set!")
	}

	fmt.Println("MongoDB URI: ", MongoDb)

	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(clientOptions)

	if err != nil {
		// TODO: handle return err
		return nil
	}
	return client
}

var Client *mongo.Client = DBInstace()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("warning: unable to found .env file")
	}
	databaseName := os.Getenv("DATABASE_NAME")

	fmt.Println("DATABASE NAME: ", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		// TODO: handle return err
		return nil
	}

	return collection
}

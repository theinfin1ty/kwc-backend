package configs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	opts := options.Client().ApplyURI(GetEnvVariable("MONGO_URI"))

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var DB = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(GetEnvVariable("DATABASE_NAME")).Collection(collectionName)
	return collection
}

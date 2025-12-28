package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoUri))
	if err != nil {
		log.Fatal("Error on MongoDB: ", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		fmt.Println("Error on connect in MongoDB")
		fmt.Println("Try again in 5 seconds")
		time.Sleep(time.Second * 5)
		ConnectMongo()
	}

	MongoClient = client
	log.Println("Connected to MongoDB!")
}

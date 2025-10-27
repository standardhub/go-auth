package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("mongo connect error: %v", err)
	}

	// simple ping
	if err := c.Ping(ctx, nil); err != nil {
		log.Fatalf("mongo ping error: %v", err)
	}

	return c
}

func getUserCollection() *mongo.Collection {
	return client.Database(mongoDatabase).Collection("users")
}
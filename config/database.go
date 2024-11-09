package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB *mongo.Database

func ConnectDB() {
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        log.Fatal("MONGODB_URI environment variable not set")
    }

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ping the MongoDB server to ensure a connection is established
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Ping(ctx, readpref.Primary())
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    fmt.Println("Successfully connected to MongoDB")

    // Set database to be used (use os.Getenv to make it configurable)
    dbName := os.Getenv("MONGODB_DATABASE")
    if dbName == "" {
        log.Fatal("MONGODB_DATABASE environment variable not set")
    }

    DB = client.Database(dbName)
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Collection(collectionName)
}

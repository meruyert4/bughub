package graph

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Mongo *mongo.Collection
}

func NewResolver() *Resolver {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Load environment variables first
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get MongoDB connection string from environment or use default
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("error connecting to mongo", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("ping failed", err)
	}

	dbName := os.Getenv("DATABASE_NAME")
	collName := os.Getenv("COLLECTION_NAME")

	if dbName == "" || collName == "" {
		log.Fatal("DATABASE_NAME or COLLECTION_NAME is empty")
	}

	db := client.Database(dbName)
	collection := db.Collection(collName)

	return &Resolver{
		Mongo: collection,
	}
}

package config

import (
	"log/slog"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func InitDatabase() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	uri := Envs.Database.URI
	docs := "https://www.mongodb.com/docs/mongodb-shell/write-scripts/env-variables/"

	if uri == "" {
		logger.Error("'MONGODB_URI' environment variable missing", "docs", docs)
		os.Exit(1)
	}

	var err error
	client, err = mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		logger.Error("Failed to connect to MongoDB", "error", err)
		os.Exit(1)
	}

	logger.Info("Connected to MongoDB successfully")
}

// MongoClient returns the MongoDB client instance.
func MongoClient() *mongo.Client {
	return client
}

// GetDatabase returns the database instance.
func GetDatabase() *mongo.Database {
	return client.Database(Envs.Database.Database)
}

// GetCollection returns a collection from the database.
func GetCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

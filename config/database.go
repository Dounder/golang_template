package config

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func InitDatabase() {
	uri := Envs.Database.URI
	docs := "www.mongodb.com/docs/drivers/go/current"

	if uri == "" {
		log.Fatal("'MONGODB_URI' environment variable missing. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	var err error
	client, err = mongo.Connect(options.Client().ApplyURI(uri))

	if err != nil {
		log.Panicf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")
}

// MongoClient returns the MongoDB client instance
func MongoClient() *mongo.Client {
	return client
}

// GetDatabase returns the database instance
func GetDatabase() *mongo.Database {
	return client.Database(Envs.Database.Database)
}

// GetCollection returns a collection from the database
func GetCollection(name string) *mongo.Collection {
	return GetDatabase().Collection(name)
}

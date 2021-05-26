package storage

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoClient will hold the structure for connecting to mongodb
type MongoClient struct {
	MongoDB       *mongo.Client
	ClientContext context.Context
}

//NewClient Create a new Instance of the Client
func NewClient(DatabaseURI string) (MongoClient, error) {

	var mclient MongoClient
	var err error
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI(DatabaseURI)
	mclient.ClientContext = ctx
	mclient.MongoDB, err = mongo.Connect(mclient.ClientContext, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mclient.MongoDB.Ping(mclient.ClientContext, nil)
	if err != nil {
		log.Fatal(err)
	}
	return mclient, err
}

package storage

import (
	"context"
	"epyphite/space/v1/SatTrackerGateway/pkg/models"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Client will hold the structure for connecting to mongodb
type Client struct {
	MongoDB       *mongo.Client
	ClientContext context.Context
}

//NewClient Create a new Instance of the Client
func NewClient(DatabaseURI string) (Client, error) {

	var mclient Client
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

//SaveSatObject will save a SatTrackObject
func (cl *Client) SaveSatObject(sattracker models.SatTrackStandard) (string, error) {
	database := cl.MongoDB.Database("SpaceOne")
	satTrackerCollection := database.Collection("SatTracker")

	insertResult, err := satTrackerCollection.InsertOne(cl.ClientContext, sattracker)
	if err != nil {
		panic(err)
	}

	oid, _ := insertResult.InsertedID.(primitive.ObjectID)
	return oid.String(), err
}

//SaveSatObjectBasic will save a SatTrackObject
func (cl *Client) SaveSatObjectBasic(sattracker models.SatTrackBasic) (string, error) {
	database := cl.MongoDB.Database("SpaceOne")
	satTrackerCollection := database.Collection("SatTrackerBasic")

	insertResult, err := satTrackerCollection.InsertOne(cl.ClientContext, sattracker)
	if err != nil {
		panic(err)
	}

	oid, _ := insertResult.InsertedID.(primitive.ObjectID)
	return oid.String(), err
}

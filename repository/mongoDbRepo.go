package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Element struct {
	ID primitive.ObjectID
	StackElement string
	InsertionTime time.Time
}

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	collection = "Stack-Collection"
	database = "Stack-Database"
)

func openMongoDBConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	var mongoURI string = "mongodb+srv://nokia:nokia@cluster0.wf4jo.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if  err!=nil {
		log.Fatal("Failed to create MongoDB client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = mongoClient.Connect(ctx)
	if err!=nil {
		log.Printf("*** Connection to MongoDB FAILED: %v ***", err)
	}

	fmt.Println("Connected to MongoDB !!!")
	return mongoClient, ctx, cancel
}

func Create(element *Element) (primitive.ObjectID, error){
	mongoClient, ctx, cancel := openMongoDBConnection()
	defer cancel()
	defer mongoClient.Disconnect(ctx)
	element.ID = primitive.NewObjectID()

	insertedElement,err := mongoClient.Database(database).Collection(collection).InsertOne(ctx, element)
	if err != nil {
		log.Printf("Could not push element into the stack: %v", err)
		return primitive.NilObjectID, err
	}
	elementId := insertedElement.InsertedID.(primitive.ObjectID)
	return elementId, nil
}
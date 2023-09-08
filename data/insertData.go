package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define structure of documents in the people collection
type Job struct {	
	JobName     string
	Description string
	Tags        []string
	DateCreated time.Time
}

func Insert(uri string) {

	// Replace the following with your Atlas connection string

	// Connect to your Atlas cluster
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	// Reference the database and collection to use
	collection := client.Database("opojobs").Collection("jobsOpen")

	// Create a new document
	newJob := Job{
		JobName:     "FullStack React",
		Description: "Create Web Responsives websites with dinasm",
		Tags:        []string{"React", "TypeScript", "Golang"},
	}

	// Insert the document into the specified collection
	collection.InsertOne(context.TODO(), newJob)

	// Find and return the document
	collection = client.Database("opojobs").Collection("jobsOpen")

	filter := bson.M{"Tags": "React"}

	var result Job
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document Found:\n%+v\n", result)
}

package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ShowAll(uri string) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("opojobs").Collection("jobsOpen")

	filter := bson.M{"tags": "React"}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result Job
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Documento Encontrado:\n%+v\n", result)
	}

	if err := cursor.Err(); err != nil {
		panic(err)
	}
}

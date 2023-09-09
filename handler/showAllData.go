package handler

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShowAllResponse struct {
	Error    error
	Response []Job
}

func ShowAll(uri string) ShowAllResponse {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return ShowAllResponse{Error: err}
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("opojobs").Collection("jobsOpen")

	// Remova o filtro ou use um filtro vazio para buscar todas as vagas
	filter := bson.M{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return ShowAllResponse{Error: err}
	}
	defer cursor.Close(context.Background())

	var results []Job

	for cursor.Next(context.Background()) {
		var result Job
		err := cursor.Decode(&result)
		if err != nil {
			return ShowAllResponse{Error: err}
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return ShowAllResponse{Error: err}
	}

	return ShowAllResponse{Response: results}
}

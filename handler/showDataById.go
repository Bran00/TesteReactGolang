package handler

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShowDataResponse struct {
	Data  interface{} // Use o tipo apropriado para os dados (por exemplo, Job)
    ID    string
	Error error // Erro, se ocorrer algum
}

func ShowDataByID(uri string, documentID string) (ShowDataResponse, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return ShowDataResponse{}, fmt.Errorf("falha ao conectar ao MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Obtém uma referência à coleção específica (substitua pelo nome da sua coleção)
	collection := client.Database("opojobs").Collection("jobsOpen")

	// Converte o ID em um formato adequado (geralmente um ObjectID)
	objectID, err := primitive.ObjectIDFromHex(documentID)

	if err != nil {
		log.Printf("Erro ao converter ID para ObjectID: %v", err)
		return ShowDataResponse{}, fmt.Errorf("falha ao converter ID: %v", err)
	}

	// Define um filtro com base no ID
	filter := bson.M{"_id": objectID}
	// Execute a consulta para encontrar o documento pelo ID
	var result Job
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Printf("Erro ao encontrar documento: %v", err)
		return ShowDataResponse{}, fmt.Errorf("falha ao encontrar documento: %v", err)
	}

	// Retorne os dados na estrutura de resposta
	return ShowDataResponse{Data: result, ID: string(documentID)}, nil
}

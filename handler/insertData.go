package handler

import (
	"context"
	"fmt"
	"time"

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

type InsertResponse struct {
	Message string `json:"message"`
}

type InsertRequest struct {
	JobName     string   `json:"jobName"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

func Insert(uri string, requestData InsertRequest) (InsertResponse, error) {
	// Conecte-se ao seu cluster Atlas
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return InsertResponse{}, fmt.Errorf("falha ao conectar ao MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Obtenha uma referência ao banco de dados e à coleção a ser usada
	collection := client.Database("opojobs").Collection("jobsOpen")

	// Crie um novo documento com base nos dados da solicitação
	newJob := Job{
		JobName:     requestData.JobName,
		Description: requestData.Description,
		Tags:        requestData.Tags,
		DateCreated: time.Now(),
	}

	// Insira o documento na coleção especificada
	_, err = collection.InsertOne(context.TODO(), newJob)
	if err != nil {
		return InsertResponse{}, fmt.Errorf("falha ao inserir documento: %v", err)
	}

	// Retorne uma resposta de sucesso
	response := InsertResponse{Message: "Documento inserido com sucesso"}
	return response, nil
}

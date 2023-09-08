package data

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteDocumentByID(uri string, documentID string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("falha ao conectar ao MongoDB: %v", err)
	}
	defer client.Disconnect(context.TODO())

	// Obtém uma referência à coleção específica (substitua pelo nome da sua coleção)
	collection := client.Database("opojobs").Collection("jobsOpen")

	// Converte o ID em um formato adequado (geralmente um ObjectID)
	objectID, err := primitive.ObjectIDFromHex(documentID)
	if err != nil {
		log.Printf("Erro ao converter ID para ObjectID: %v", err)
		return fmt.Errorf("falha ao converter ID: %v", err)
	}

	// Define um filtro com base no ID
	filter := bson.M{"_id": objectID}

	// Executa a exclusão do documento
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Printf("Erro ao excluir documento: %v", err)
		return fmt.Errorf("falha ao excluir documento: %v", err)
	}

	log.Printf("Documento com ID %s excluído com sucesso.", documentID)

	return nil
}

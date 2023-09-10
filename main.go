package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/gin-contrib/cors" // Comente ou remova esta linha
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/bran00/TesteReactGolang/router"
)

func main() {

	mongoPassword := os.Getenv("MONGO_DB_PASSWORD")
	mongoUriString := os.Getenv("MONGO_URI_STRING")

	mongoURI := fmt.Sprintf(mongoUriString, mongoPassword)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}

	r := router.Initialize(mongoURI)

	// Iniciar o servidor na porta desejada
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil)) // Substitua pela porta desejada
}

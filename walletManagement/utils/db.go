package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/spenmo-jamboree/walletManagement/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbConnection interface {
	Connection() *mongo.Client
}

type dbConnectionImpl struct{}

var (
	DbConnectionInstance dbConnection = dbConnectionImpl{}
	DbConnection                      = DbConnectionInstance.Connection()
)

func (dbConnection dbConnectionImpl) Connection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(common.DB_CONNECTION_STRING)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

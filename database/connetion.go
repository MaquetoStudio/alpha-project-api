package database

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
	err            error
)

func NewDbInstance() *mongo.Database {
	return getMongoClient().Database("maquetodb")
}

func getMongoClient() *mongo.Client {
	clientOnce.Do(func() {
		// Votre URI de connexion
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		// Connectez-vous à MongoDB
		clientInstance, err = mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		// Vérifiez la connexion
		err = clientInstance.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}
	})

	return clientInstance
}

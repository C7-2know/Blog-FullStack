package bootstrap

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDatabase(env *Env) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := env.DBUri
	clientOpts:= options.Client().ApplyURI(uri)
	client,err:= mongo.Connect(ctx,clientOpts)
	if err != nil {
		log.Fatalf("Error creating mongo client: %v", err)
	}

	err=client.Ping(ctx,readpref.Primary())
	if err != nil {
		log.Fatalf("Error pinging mongo: %v", err)
	}
	return client
}


func CheckDatabaseConnection(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

func CloseMongoDBConnection(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Error disconnecting from mongo: %v", err)
	}
	log.Println("Disconnected from MongoDB")
}
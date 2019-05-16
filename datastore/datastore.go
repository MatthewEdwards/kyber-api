package datastore

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Session *mongo.Client
}

// NewDBConnection will return a connection to the database
func NewDBConnection() (connection *MongoDB) {
	db := &MongoDB{
		Session: newDBClient(),
	}

	if db.Session == nil {
		log.Fatal("Unable to connect to the database")
	}
	return db
}

func newDBClient() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Error("Unable to connect to the database", err)
		return nil
	}

	return client
}

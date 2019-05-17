package datastore

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB stores a mongo client session
type MongoDB struct {
	Session *mongo.Client
}

// NewDBConnection will return a connection to the database
func NewDBConnection() (connection *MongoDB) {
	db := &MongoDB{
		Session: newClient(),
	}

	if db.Session == nil {
		log.Fatal("Unable to connect to the database")
	}
	return db
}

func newClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	defer cancel()

	if err != nil {
		log.Error("Unable to connect to the database", err)
		return nil
	}

	return client
}

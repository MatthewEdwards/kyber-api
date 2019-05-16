package controllers

import (
	"kyber-api/datastore"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// DStore is used to store a database connection
type DStore struct {
	connection *datastore.MgoConnection
}

// DatastoreConnect will get a connection to the database
func DatastoreConnect() *DStore {
	log.Info("[API] DatastoreConnect")
	DS := &DStore{
		connection: datastore.NewDBConnection(),
	}

	return DS
}

// NewAPIRouter will setup the API routes
func NewAPIRouter() *mux.Router {
	log.Info("[API] NewAPIRouter")
	DS := DatastoreConnect()

	r := mux.NewRouter().StrictSlash(true)

	/*******************************
	* 		API V1 Routes
	*******************************/
	// Articles
	r.HandleFunc("/api/v1/articles", DS.GetArticlesHandler).Methods("GET")
	r.HandleFunc("/api/v1/articles", DS.PostArticlesHandler).Methods("POST")

	return r
}

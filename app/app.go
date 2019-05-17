package app

import (
	"kyber-api/datastore"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type application struct {
	router *mux.Router
	db     *datastore.MongoDB
}

func (app *application) start() {
	if app.router != nil {
		log.Info("Starting server")
		log.Info(http.ListenAndServe(":8080", app.router))
	} else {
		log.Fatal("Unable to start server")
	}
}

// NewApplication will initliaze the appplication and start the server
func NewApplication() {
	log.Info("Setting up a new application")
	app := &application{
		router: mux.NewRouter().StrictSlash(true),
		db:     datastore.NewDBConnection(),
	}
	app.routes()
	app.start()
}

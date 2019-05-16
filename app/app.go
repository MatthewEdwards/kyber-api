package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type application struct {
	router *mux.Router
}

func (app *application) routes() {
	//app.router = mux.NewRouter().StrictSlash(true)
	//app.router.HandleFunc("/api/v1/articles", app.handlePostAricle).Methods("POST")

	app.router.HandleFunc("/api/v1/articles", app.handleGetAricles()).Methods("GET")
}

func newApplication() *application {
	return &application{
		router: mux.NewRouter(),
	}
}

func Start() {
	app := newApplication()
	app.newServer()
}

func (app *application) newServer() {

	app.routes()

	if app.router != nil {
		log.Info("Starting API")
		log.Info(http.ListenAndServe(":8080", app.router))
	}
}

func response(w http.ResponseWriter, responseData interface{}, status int) {
	log.Info("Retuning a response")

	jsn, err := json.MarshalIndent(responseData, "", "    ")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	w.Write(jsn)

}

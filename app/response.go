package app

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type response struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func respond(w http.ResponseWriter, responseData interface{}, status int) {
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

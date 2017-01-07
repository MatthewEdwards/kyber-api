package cmd

import (
	"kyber-api/api"
	"net/http"
	log "github.com/Sirupsen/logrus"
)

// LaunchAPI launchs the API server
func LaunchAPI() {
	log.Info("[API] LaunchAPI")

	router := api.NewAPIRouter()

	if router != nil{
		log.Info("[LaunchAPI] Launching API server")
		log.Info(http.ListenAndServe(":8000", router))
	}
}
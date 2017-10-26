/**
 * Kyber API Server
 */
package main

import (
	"kyber-api/controllers"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

// LaunchAPI launchs the API server
func LaunchAPI() {
	log.Info("[API] LaunchAPI")

	router := controllers.NewAPIRouter()

	if router != nil {
		log.Info("[LaunchAPI] Launching API server")
		log.Info(http.ListenAndServe(":8000", router))
	}
}

func main() {
	app := &cli.App{
		Name:  "Kyber",
		Usage: "News Dashboard",
		Commands: []*cli.Command{
			{
				// This will launch the API server
				Name:        "run",
				Aliases:     []string{"r"},
				Usage:       "Launch Kyber",
				Description: "Launch the API server",
				Action: func(c *cli.Context) error {
					log.Info("[APP] Launching API server")
					LaunchAPI()
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}

/**
 * Kyber API Server
 */
package main

import (
	"kyber-api/cmd"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:    "Kyber",
		Usage:   "News Dashboard",
		Commands: []*cli.Command{
			{
				// This will launch the API server
				Name:        "api",
				Aliases:     []string{"a"},
				Usage:       "Launch Kyber",
				Description: "Launch the API server",
				Action: func(c *cli.Context) error {
					log.Info("[APP] Launching API server")
					cmd.LaunchAPI()
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}

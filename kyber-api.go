package main

import (
	"kyber-api/app"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

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
					app.NewApplication()
					return nil
				},
			},
		},
	}
	app.Run(os.Args)
}

package main

import (
	"os"
	"log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
    Name: "source-pagerduty",
    Usage: "An Airbyte Source for the PagerDuty API.",
    Commands: []*cli.Command{
			{
				Name: "spec",
				Usage: "Returns the connector specification.",
				Action: func(c *cli.Context) error {
					log.Println("Hello, spec!")
					return nil
				},
			},
			{
				Name: "check",
				Usage: "The check command attempts to connect to the API.",
				Action: func(c *cli.Context) error {
					log.Println("Hello, check!")
					return nil
				},
			},
			{
				Name: "discover",
				Usage: "This command detects the structure of the data.",
				Action: func(c *cli.Context) error {
					log.Println("Hello, discover!")
					return nil
				},
			},
			{
				Name: "read",
				Usage: "This command reads data from the underlying data source.",
				Action: func(c *cli.Context) error {
					log.Println("Hello, read!")
					return nil
				},
			},
		},
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}

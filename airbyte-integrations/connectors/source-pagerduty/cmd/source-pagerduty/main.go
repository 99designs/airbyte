package main

import (
	_ "embed"
	"encoding/json"
	"github.com/urfave/cli/v2"
	"github.com/airbytehq/airbyte-integrations/connectors/source-pagerduty/internal"
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

func spec(c *cli.Context) error {
	message, err := internal.Specification()

	if err != nil {
		panic(err)
	}

	mBytes, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(mBytes))

	return nil
}

func check(c *cli.Context) error {
	bytes, err := ioutil.ReadFile(c.String("config"))
	log.Println(string(bytes))

	if err != nil {
		panic(err)
	}

	var config internal.Config

	err = json.Unmarshal(bytes, &config)

	if err != nil {
		panic(err)
	}

	log.Println(config)

	airbyteStatus, err := internal.Check(config)

	if err != nil {
		panic(err)
	}

	statusBytes, err := json.Marshal(airbyteStatus)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(statusBytes))

	return nil
}

func main() {
	app := &cli.App{
		Name:  "source-pagerduty",
		Usage: "An Airbyte Source for the PagerDuty API.",
		Commands: []*cli.Command{
			{
				Name:  "spec",
				Usage: "Returns the connector specification.",
				Action: spec,
			},
			{
				Name:  "check",
				Usage: "The check command attempts to connect to the API.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "config",
						Usage: "The filepath of a json file containing the config.",
					},
				},
				Action: check,
			},
			{
				Name:  "discover",
				Usage: "This command detects the structure of the data.",
				Action: func(c *cli.Context) error {
					log.Println("Hello, discover!")
					return nil
				},
			},
			{
				Name:  "read",
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

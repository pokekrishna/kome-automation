	package main

	import (
		"github.com/pokekrishna/kome-automation/config"
		"github.com/urfave/cli/v2"
		"log"
		"net/http"
		"os"
	)

const (
	appName = "kome-automation-server"
)

func rootHandler(response http.ResponseWriter, request *http.Request) {

}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "This is a web application used for Kome"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Usage:    "Path to configuration file",
			Value:    "./config.yaml",
			Required: true,
		},
	}
	app.Action = initialize

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Cannot run the app: ", err)
	}
}

func initialize(ctx *cli.Context) error {
	// First, check if config file is valid and then load
	_, err := config.LoadConf(ctx.String("config"))
	if err != nil {
		// if there is any problem in getting the conf, exit
		log.Fatal("Cannot load Config File: ", err)
	}
	return nil
}

package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"github.com/pokekrishna/kome-automation/server"
)

const (
	appName     = "kome-automation-server"
)

func rootHandler(response http.ResponseWriter, request *http.Request) {

}



func main() {
	app := cli.NewApp()
	app.Name =  appName
	app.Usage = "This is a web application used for Kome"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Usage: "Path to configuration file",
			Value: "./config.yaml",
			Required: true,

		},
	}
	app.Action = server.StartServer

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Cannot run the app: ", err)
	}
}

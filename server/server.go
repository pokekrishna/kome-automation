package server

import (
	"github.com/pokekrishna/kome-automation/config"
	"github.com/urfave/cli/v2"
	"log"
)

func StartServer(ctx *cli.Context) error{
	// First, check if config file is valid
	if err := config.ValidateConfFile(ctx.String("config")); err != nil{
		log.Fatal("Config File Invalid: ", err)
	}




	//http.HandleFunc("/", rootHandler)
	//err := http.ListenAndServe(
	//	fmt.Sprintf("%s:%s", bind_address, bind_port),
	//	nil,
	//	)
	//if err != nil{
	//	log.Fatal("Error in listening and serving")
	//}

	return nil
}
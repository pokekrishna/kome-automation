package server

import (
	"github.com/pokekrishna/kome-automation/config"
	"github.com/urfave/cli/v2"
	"log"
)

func StartServer(ctx *cli.Context) error {
	// First, check if config file is valid
	_, err := config.GetConf(ctx.String("config"))
	if err != nil {
		// if there is any problem in getting the conf, exit
		log.Fatal("Config File Invalid: ", err)
	}

	// TODO: think about what should be replacement name of get conf
	// since server doesnt need to have the conf file, it would just need the
	// address and other params needed to start up
	//
	// idea 1
	// get the needed params and load the yaml via init() in the package

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

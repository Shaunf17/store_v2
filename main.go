package main

import (
	"fmt"
	"store/configs"
	"store/server"
	"store/utils/logger"
)

func main() {
	// Configure and initialise logger
	logcfg := configs.LoggerConfigs{
		On:             true,
		Filename:       "store.log",
		WriteToConsole: true,
		WriteToFile:    false,
	}
	logerr := logger.Init(logcfg)
	if logerr != nil {
		switch logerr {
		case logger.ErrCantOpenFile:
			fmt.Println(logerr)
		case logger.ErrFileAlreadyOpen:
			fmt.Println(logerr)
		}
	}

	// Configure and start server
	srvcfg := configs.ServerConfigs{
		Host: "localhost",
		Port: ":8080",
		Root: "/",
	}
	server.Start(srvcfg)
}

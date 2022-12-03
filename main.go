package main

import (
	"log"
	"os"

	"github.com/so5dz/wx433/settings"
	"github.com/so5dz/wx433/storage"
)

type OperationMode string

const (
	OperationMode_Serve OperationMode = "serve"
	OperationMode_Read  OperationMode = "read"
)

func main() {
	operationMode := OperationMode(os.Args[1])
	settingsFile := os.Args[2]

	var wx433Settings settings.WX433Settings
	err := wx433Settings.Parse(settingsFile)
	if err != nil {
		log.Fatalln(err)
	}

	var wx433Storage storage.WX433Storage
	err = wx433Storage.Open(wx433Settings.DatabaseFile)
	if err != nil {
		log.Fatalln(err)
	}

	if operationMode == OperationMode_Serve {
		runServer(wx433Settings, wx433Storage)
	} else if operationMode == OperationMode_Read {
		readWeather(wx433Settings, wx433Storage)
	} else {
		log.Fatalln("unrecognized operation mode", operationMode)
	}
}

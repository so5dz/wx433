package main

import (
	"encoding/json"
	"log"

	"github.com/so5dz/wx433/model"
	"github.com/so5dz/wx433/settings"
	"github.com/so5dz/wx433/storage"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func runServer(wx433Settings settings.WX433Settings, wx433Storage storage.WX433Storage) error {
	log.Println("running in server mode")

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUDP(wx433Settings.ListenPath())
	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			messageString := logParts["message"].(string)
			var messageObj model.Message
			err := json.Unmarshal([]byte(messageString), &messageObj)
			if err != nil {
				log.Println("unable to deserialize", messageString)
				continue
			}
			if wx433Settings.IsWeatherDevice(messageObj.Model) {
				log.Println("temperature", messageObj.Temperature, "'C\t humidity", messageObj.Humidity, "%")
				err = wx433Storage.Persist(messageObj)
				if err != nil {
					log.Println("unable to persist reading", err)
					continue
				}
			}
		}
	}(channel)

	server.Wait()
	return nil
}

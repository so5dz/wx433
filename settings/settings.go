package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type WX433Settings struct {
	ListenHost      string       `json:"listenHost"`
	ListenPort      int          `json:"listenPort"`
	WeatherDevices  []string     `json:"weatherDevices"`
	DatabaseFile    string       `json:"databaseFile"`
	AveragingPeriod int          `json:"averagingPeriod"`
	APRS            APRSSettings `json:"aprs"`
}

func (ws *WX433Settings) Parse(filePath string) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileBytes, ws)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WX433Settings) IsWeatherDevice(deviceName string) bool {
	for _, weatherDeviceName := range ws.WeatherDevices {
		if weatherDeviceName == deviceName {
			return true
		}
	}
	return false
}

func (ws *WX433Settings) ListenPath() string {
	return fmt.Sprintf("%s:%d", ws.ListenHost, ws.ListenPort)
}

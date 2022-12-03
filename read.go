package main

import (
	"fmt"
	"math"

	"github.com/so5dz/wx433/settings"
	"github.com/so5dz/wx433/storage"
)

func readWeather(wx433Settings settings.WX433Settings, wx433Storage storage.WX433Storage) {
	reports, err := wx433Storage.Fetch(wx433Settings.AveragingPeriod)
	if err != nil {
		return
	}
	if len(reports) == 0 {
		return
	}

	temperatureSum := 0.0
	humiditySum := 0.0
	humidityReadings := 0
	for _, report := range reports {
		temperatureSum += report.Temperature
		if report.Humidity > 0 {
			humiditySum += report.Humidity
			humidityReadings++
		}
	}
	temperature := temperatureSum / float64(len(reports))
	humidity := humiditySum / float64(humidityReadings)

	fmt.Printf("!%s%c%s%c.../...t%sh%s %s\n",
		wx433Settings.APRS.EncodedLatitude(),
		wx433Settings.APRS.SymbolTable(),
		wx433Settings.APRS.EncodedLongitude(),
		wx433Settings.APRS.SymbolCode(),
		encodeTemperatureValue(temperature),
		encodeHumidityValue(humidity),
		wx433Settings.APRS.Comment,
	)
}

func encodeTemperatureValue(celsius float64) string {
	fahrenheit := int(math.Round(celsius*1.8 + 32))
	if fahrenheit < 0 {
		return fmt.Sprintf("-%02d", -fahrenheit)
	} else {
		return fmt.Sprintf("%03d", fahrenheit)
	}
}

func encodeHumidityValue(percentage float64) string {
	return fmt.Sprintf("%02.0f", percentage)
}

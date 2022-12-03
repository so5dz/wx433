package settings

import (
	"fmt"
	"math"
)

type APRSSettings struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Symbol    string  `json:"symbol"`
	Comment   string  `json:"comment"`
}

func (as *APRSSettings) EncodedLatitude() string {
	latitude := math.Mod(as.Latitude, 90.0)
	absoluteLatitude := math.Abs(latitude)
	direction := "N"
	if latitude < 0.0 {
		direction = "S"
	}
	fullDegrees := math.Floor(absoluteLatitude)
	fractionalDegrees := absoluteLatitude - fullDegrees
	fractionalMinutes := fractionalDegrees * 60.0
	return fmt.Sprintf("%2.0f%05.2f%s", fullDegrees, fractionalMinutes, direction)
}

func (as *APRSSettings) EncodedLongitude() string {
	longitude := math.Mod(as.Longitude, 180.0)
	absoluteLongitude := math.Abs(longitude)
	direction := "E"
	if longitude < 0.0 {
		direction = "W"
	}
	fullDegrees := math.Floor(absoluteLongitude)
	fractionalDegrees := absoluteLongitude - fullDegrees
	fractionalMinutes := fractionalDegrees * 60.0
	return fmt.Sprintf("%3.0f%05.2f%s", fullDegrees, fractionalMinutes, direction)
}

func (as *APRSSettings) SymbolTable() byte {
	return as.Symbol[0]
}

func (as *APRSSettings) SymbolCode() byte {
	return as.Symbol[1]
}

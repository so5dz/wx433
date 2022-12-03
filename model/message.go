package model

type Message struct {
	Time        Time    `json:"time"`
	Model       string  `json:"model"`
	Temperature float64 `json:"temperature_C"`
	Humidity    float64 `json:"humidity"`
}

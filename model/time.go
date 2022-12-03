package model

import (
	"encoding/json"
	"strings"
	"time"
)

type Time time.Time

const TimeFormat = "2006-01-02 15:04:05"

func (j *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(TimeFormat, s)
	if err != nil {
		return err
	}
	*j = Time(t)
	return nil
}

func (j Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j Time) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

package storage

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/so5dz/wx433/model"
)

const _DBDriver = "sqlite3"
const _WeatherCreateQuery = "CREATE TABLE IF NOT EXISTS weather (id INTEGER NOT NULL PRIMARY KEY, time DATETIME NOT NULL, temperature REAL, humidity REAL);"
const _WeatherInsertQuery = "INSERT INTO weather(time, temperature, humidity) VALUES (?, ?, ?)"
const _WeatherSelectQuery = "SELECT time, temperature, humidity FROM weather WHERE (CAST(strftime('%s', CURRENT_TIMESTAMP) AS INTEGER)-time) < ?"

type WX433Storage struct {
	db *sql.DB
}

func (ws *WX433Storage) Open(filePath string) error {
	var err error
	if ws.db, err = sql.Open(_DBDriver, filePath); err != nil {
		return err
	}
	if _, err = ws.db.Exec(_WeatherCreateQuery); err != nil {
		return err
	}
	return nil
}

func (ws *WX433Storage) Persist(message model.Message) error {
	transaction, err := ws.db.Begin()
	if err != nil {
		return err
	}
	statement, err := transaction.Prepare(_WeatherInsertQuery)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(time.Time(message.Time).Unix(), message.Temperature, message.Humidity)
	if err != nil {
		return err
	}
	err = transaction.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ws *WX433Storage) Fetch(lastSeconds int) ([]model.Message, error) {
	messages := make([]model.Message, 0)
	transaction, err := ws.db.Begin()
	if err != nil {
		return messages, err
	}
	statement, err := transaction.Prepare(_WeatherSelectQuery)
	if err != nil {
		return messages, err
	}
	defer statement.Close()
	rows, err := statement.Query(lastSeconds)
	if err != nil {
		return messages, err
	}
	defer rows.Close()
	for rows.Next() {
		var time time.Time
		var temp float64
		var humi float64
		err = rows.Scan(&time, &temp, &humi)
		if err != nil {
			return messages, err
		}
		messages = append(messages, model.Message{
			Time:        model.Time(time),
			Temperature: temp,
			Humidity:    humi,
		})
	}

	if err = rows.Err(); err != nil {
		return messages, err
	}
	if err = transaction.Commit(); err != nil {
		return messages, err
	}
	return messages, nil
}

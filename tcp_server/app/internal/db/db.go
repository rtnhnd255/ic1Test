package db

import (
	"log"

	"github.com/rtnhnd255/ic1Test/app/internal/config"
	"github.com/rtnhnd255/ic1Test/app/internal/model"

	"github.com/jackc/pgx"
)

//TODO: write migrations and method like "CHECK DOES IT EXIST AND IF ITS NOT CREATE IT YOOOO"

type Db struct {
	Config     pgx.ConnConfig
	Connection *pgx.Conn
}

func NewDBConn(config *config.Config) *Db {
	return &Db{
		Config: pgx.ConnConfig{
			User:     config.DbUser,
			Port:     config.DbPort,
			Host:     config.DbHost,
			Password: config.DbPassword,
			Database: config.DbName,
		},
	}
}

func (db *Db) open() error {
	var err error
	db.Connection, err = pgx.Connect(db.Config)

	if err != nil {
		log.Println("Postrges Open() problem", err)
		return err
	}
	return nil
}

func (db *Db) CreateRecord(rec model.RecordDTO) error {
	err := db.open()
	if err != nil {
		log.Println("Error while connecting pg:", err)
		return err
	}

	res := db.Connection.QueryRow(`
	INSERT INTO points (device_id, point_time, latitude, longitude)
	VALUES ($1, $2, $3, $4)
	`, rec.DeviceID, rec.PointTime, rec.Latitude, rec.Longitude)

	var id int
	err = res.Scan(&id)
	if err != nil {
		log.Println("Error while writing to pg:", err)
	}
	return nil
}

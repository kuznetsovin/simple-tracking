package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DbStore interface {
	LastPositionGateway
	TrackGateway
	GeoObjectGateway
	VehicleGateway
	ReportGateway
}

type DB struct {
	*sql.DB
}

func NewDataStore(dbConnection string) (*DB, error) {
	db, err := sql.Open("postgres", dbConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, err
}

/*
File describe gateway for working with vehicle data in database.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/
package models

type VehicleGateway interface {
	GetVehicleDict() (Vehicles, error)
	AddVehicle(VehicleRec) error
}

type VehicleRec struct {
	GpsID     int    `json:"gps_id"`
	GosNumber string `json:"gos_number"`
}

type Vehicles []VehicleRec

func (db *DB) GetVehicleDict() (Vehicles, error) {
	result := Vehicles{}

	rows, err := db.Query(`select gps_code, gos_number from vehicle`)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := VehicleRec{}
		if err := rows.Scan(&r.GpsID, &r.GosNumber); err != nil {
			return result, err
		}
		result = append(result, r)
	}
	err = rows.Err()

	return result, err
}

func (db *DB) AddVehicle(v VehicleRec) error {
	_, err := db.Exec(`insert into vehicle (gos_number, gps_code) values ($1, $2)`, v.GosNumber, v.GpsID)
	return err
}

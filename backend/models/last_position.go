/*
File describe gateway for working with last points in database.

Author: Igor Kuznetsov
Email: me@swe-notes.ru

(c) Copyright by Igor Kuznetsov.
*/
package models

type LastPositionGateway interface {
	GetLastVehiclePosition() (LastPositions, error)
}

type LastPositions []Point

func (db *DB) GetLastVehiclePosition() (LastPositions, error) {
	result := LastPositions{}

	rows, err := db.Query(`select id, navigate_date, packet_id, ST_X(ST_Transform(geom, 3857)), 
       ST_Y(ST_Transform(geom, 3857)), course from last_point.clients;`)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := Point{}
		if err := rows.Scan(&r.Client, &r.NavigateDate, &r.PacketID, &r.Longitude, &r.Latitude, &r.Course); err != nil {
			return result, err
		}
		result = append(result, r)
	}
	err = rows.Err()

	return result, err
}

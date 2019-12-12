package models

import "time"

type TrackGateway interface {
	GetTrackByClient(client uint32, dateStart, dateEnd time.Time) (Track, error)
}

type Track [][]float64

func (db *DB) GetTrackByClient(clientId uint32, dateStart, dateEnd time.Time) (Track, error) {
	result := Track{}

	rows, err := db.Query(`select ST_X(ST_Transform(geom, 3857)), ST_Y(ST_Transform(geom, 3857)) 
from vts.track 
where client = $1 and navigate_date between $2 and $3 
  and longitude between 37.32660 and 37.380005 
  and latitude between 55.675275 and 55.71167 
order by navigate_date`,
		clientId, dateStart, dateEnd)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := make([]float64, 2)
		if err := rows.Scan(&r[0], &r[1]); err != nil {
			return result, err
		}
		result = append(result, r)
	}
	err = rows.Err()

	return result, err
}

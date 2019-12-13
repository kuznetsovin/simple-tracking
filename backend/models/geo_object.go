package models

import (
	"github.com/kuznetsovin/go.geojson"
)

type GeoObjectGateway interface {
	GetGeoObjects() (GeoObjects, error)
	AddObject(GeoObject) error
}

type GeoObjectCommon struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type GeoObject struct {
	GeoObjectCommon
	Geom *geojson.Geometry `json:"geom,omitempty"`
}
type GeoObjects []GeoObject

func (db *DB) GetGeoObjects() (GeoObjects, error) {
	result := GeoObjects{}

	rows, err := db.Query(`SELECT id, name, ST_AsGeoJSON(ST_Transform(geom, 3857)) geom FROM geo_object`)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := GeoObject{}
		if err := rows.Scan(&r.Id, &r.Name, &r.Geom); err != nil {
			return result, err
		}
		result = append(result, r)
	}
	err = rows.Err()

	return result, err
}

func (db *DB) AddObject(g GeoObject) error {
	_, err := db.Exec(`insert into geo_object (name, geom) values ($1, st_setsrid(st_geomfromgeojson($2), 3857))`,
		g.Name, g.Geom)
	return err
}

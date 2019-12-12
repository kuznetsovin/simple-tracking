package models

import "time"

type ReportGateway interface {
	GetObjectReport(client uint32, dateStart, dateEnd time.Time) (ObjectReport, error)
}

type ObjectReportRow struct {
	GeoObjectCommon
	FirstPointTimestamp time.Time `json:"first_point_timestamp"`
	LastPointTimestamp  time.Time `json:"last_point_timestamp"`
	Mileage             float32   `json:"mileage"`
}

type ObjectReport []ObjectReportRow

func (db *DB) GetObjectReport(clientId uint32, dateStart, dateEnd time.Time) (ObjectReport, error) {
	result := ObjectReport{}

	rows, err := db.Query(`with uniq_point as (select distinct packet_id,
                                    navigate_date,
                                    geom,
                                    geom::geography geo_geom
                    from vts.track
                    where client = $1
                      and navigate_date between $2 and $3
                      and longitude between 37.32660 and 37.380005
                      and latitude between 55.675275 and 55.71167
                    order by navigate_date
),
     dist_point as (
         select packet_id,
                navigate_date,
                st_transform(geom, 3857)                                                           geom,
                st_distance(geo_geom, lag(geo_geom, 1, geo_geom) OVER (ORDER BY navigate_date)) as dist
         from uniq_point
     ),
     dist_by_obj as (select p.navigate_date,
                            go.id as odh_id,
                            p.dist,
                            go.id    obj_id,
                            go.name  obj_name
                     from geo_object go
                              inner join dist_point p on st_intersects(p.geom, go.geom)
                     order by 1)
select obj_id,
       obj_name,
       min(navigate_date) first_point_timestamp,
       max(navigate_date) last_point_timestamp,
       round((sum(dist) / 1000)::numeric, 2)   mileage
from dist_by_obj
group by obj_id, obj_name
order by obj_id`,
		clientId, dateStart, dateEnd)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := ObjectReportRow{}
		if err := rows.Scan(&r.Id, &r.Name, &r.FirstPointTimestamp, &r.LastPointTimestamp, &r.Mileage); err != nil {
			return result, err
		}
		result = append(result, r)
	}
	err = rows.Err()

	return result, err
}

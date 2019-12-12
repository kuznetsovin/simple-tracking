package models

type VehicleGateway interface {
	GetVehicleDict() (VehicleDict, error)
}

type VehicleDictRec struct {
	GpsID     int    `json:"gps_id"`
	GosNumber string `json:"gos_number"`
}

type VehicleDict []VehicleDictRec

func (db *DB) GetVehicleDict() (VehicleDict, error) {
	result := VehicleDict{}

	rows, err := db.Query(`select gps_code, gos_number from vehicle`)

	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := VehicleDictRec{}
		if err := rows.Scan(&r.GpsID, &r.GosNumber); err != nil {
			return result, err
		}
		result = append(result, r)
	}
	err = rows.Err()

	return result, err
}

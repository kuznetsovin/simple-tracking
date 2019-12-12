package utils

import (
	"encoding/json"
	"github.com/go-spatial/proj"
	"github.com/kuznetsovin/go.geojson"
	"time"
)

type BrokerMsg struct {
	Client              uint32  `json:"client"`
	PacketID            uint32  `json:"packet_id"`
	NavigationTimestamp int64   `json:"navigation_unix_time"`
	ReceivedTimestamp   int64   `json:"received_unix_time"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Speed               uint16  `json:"speed"`
	Nsat                uint8   `json:"nsat"`
	Course              uint8   `json:"course"`
}

func (m *BrokerMsg) FromBytes(msg []byte) error {
	return json.Unmarshal(msg, m)
}

func (m *BrokerMsg) ToGeoFeature() (geojson.Feature, error) {

	result := geojson.Feature{}

	lonlat, err := proj.Convert(proj.EPSG3857, []float64{m.Longitude, m.Latitude})
	if err != nil {
		return result, err
	}

	result.Geometry = geojson.NewPointGeometry(lonlat)
	result.Properties = map[string]interface{}{
		"client":    m.Client,
		"packet_id": m.PacketID,
		"nav_time":  time.Unix(m.NavigationTimestamp, 0).UTC(),
		"course":    m.Course,
	}

	return result, err
}

package utils

import (
	"github.com/kuznetsovin/go.geojson"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBrokerMsg_FromBytes(t *testing.T) {
	msg := []byte(`{"client":445,"packet_id":22989,"navigation_unix_time":1573116514,"received_unix_time":1573116516,"latitude":55.697471661422746,"longitude":37.3470249812461,"speed":30,"pdop":0,"hdop":0,"vdop":0,"nsat":17,"ns":0,"course":219,"kbm_sensors":[{"sensor_number":1,"value":1,"type":"din"},{"sensor_number":2,"value":0,"type":"din"},{"sensor_number":3,"value":0,"type":"din"},{"sensor_number":4,"value":0,"type":"din"},{"sensor_number":5,"value":0,"type":"din"},{"sensor_number":6,"value":0,"type":"din"},{"sensor_number":7,"value":0,"type":"din"},{"sensor_number":8,"value":0,"type":"din"}],"liquid_sensors":[{"sensor_number":1,"error_flag":"0","value_mm":2343,"value_l":0}]}`)

	brokerMessage := BrokerMsg{
		Client:              445,
		PacketID:            22989,
		NavigationTimestamp: 1573116514,
		ReceivedTimestamp:   1573116516,
		Latitude:            55.697471661422746,
		Longitude:           37.3470249812461,
		Speed:               30,
		Nsat:                17,
		Course:              219,
	}

	m := BrokerMsg{}

	if assert.NoError(t, m.FromBytes(msg)) {
		assert.Equal(t, brokerMessage, m)
	}
}

func Test_transform(t *testing.T) {
	m := BrokerMsg{
		Client:              445,
		PacketID:            22989,
		NavigationTimestamp: 1573116514,
		ReceivedTimestamp:   1573116516,
		Latitude:            55.697471661422746,
		Longitude:           37.3470249812461,
		Speed:               30,
		Nsat:                17,
		Course:              219,
	}

	geoFeature := geojson.Feature{
		Geometry: geojson.NewPointGeometry([]float64{4157451.803555983, 7498425.043314084}),
		Properties: map[string]interface{}{
			"client":    m.Client,
			"packet_id": m.PacketID,
			"nav_time":  time.Date(2019, time.November, 07, 8, 48, 34, 0, time.UTC),
			"course":    m.Course,
		},
	}

	r, err := m.ToGeoFeature()
	if assert.NoError(t, err) {
		assert.Equal(t, geoFeature, r)
	}
}

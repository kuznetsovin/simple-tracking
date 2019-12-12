package handlers

import (
	"github.com/kuznetsovin/go.geojson"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"simple-tracking/backend/models"
	"time"
)

func createReq(httpMethod, endpoint, urlParams string) *http.Request {
	target := endpoint
	if urlParams != "" {
		target += "?" + urlParams
	}
	req := httptest.NewRequest(httpMethod, target, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return req
}

type mockStore struct{}

func (m *mockStore) GetLastVehiclePosition() (models.LastPositions, error) {
	r := models.LastPositions{
		{
			Client:       445,
			NavigateDate: time.Date(2019, time.November, 19, 13, 47, 0, 0, time.UTC),
			PacketID:     7242,
			Longitude:    37.37615997,
			Latitude:     55.69934332,
			Course:       53,
		},
		{
			Client:       156,
			NavigateDate: time.Date(2019, time.November, 19, 13, 35, 1, 0, time.UTC),
			PacketID:     6980,
			Longitude:    37.37656997,
			Latitude:     55.69921665,
			Course:       0,
		},
	}

	return r, nil
}

func (m *mockStore) GetTrackByClient(client uint32, dateStart, dateEnd time.Time) (models.Track, error) {
	r := models.Track{
		{37.37656997, 55.69921665},
		{37.37656102, 55.69921671},
	}

	return r, nil
}

func (m *mockStore) GetGeoObjects() (models.GeoObjects, error) {
	geom1, _ := geojson.UnmarshalGeometry([]byte(`{"type":"Polygon","coordinates":[[[4156985.87172938,7495035.97343491],[4157066.06721206,7495009.08972197]]]}`))
	geom2, _ := geojson.UnmarshalGeometry([]byte(`{"type":"MultiPolygon","coordinates":[[[[4158501.3758332,7496734.60581653],[4158501.34448372,7496734.57589504]]]]}`))
	r := models.GeoObjects{
		{models.GeoObjectCommon{Id: 1, Name: "Объект №1"}, geom1},
		{models.GeoObjectCommon{Id: 2, Name: "Объект №2"}, geom2},
	}

	return r, nil
}

func (m *mockStore) GetVehicleDict() (models.VehicleDict, error) {
	r := models.VehicleDict{
		{GpsID: 1, GosNumber: "A777AA750"},
		{GpsID: 2, GosNumber: "B777BB750"},
	}
	return r, nil
}

func (m *mockStore) GetObjectReport(client uint32, dateStart, dateEnd time.Time) (models.ObjectReport, error) {
	r := models.ObjectReport{
		{
			GeoObjectCommon:     models.GeoObjectCommon{Id: 1, Name: "object 1"},
			FirstPointTimestamp: time.Date(2019, time.November, 19, 13, 35, 1, 0, time.UTC),
			LastPointTimestamp:  time.Date(2019, time.November, 19, 14, 35, 1, 0, time.UTC),
			Mileage:             30.5,
		},
	}
	return r, nil
}

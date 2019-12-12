package handlers

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_GeoObjects(t *testing.T) {
	endpoint := "/last-positions"

	e := echo.New()
	rec := httptest.NewRecorder()
	req := createReq(http.MethodGet, endpoint, "")
	c := e.NewContext(req, rec)
	h := &Handler{
		DB: &mockStore{},
	}

	if assert.NoError(t, h.GeoObjects(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[4156985.87172938,7495035.97343491],[4157066.06721206,7495009.08972197]]]},"properties":{"id":1,"name":"Объект №1"}},{"type":"Feature","geometry":{"type":"MultiPolygon","coordinates":[[[[4158501.3758332,7496734.60581653],[4158501.34448372,7496734.57589504]]]]},"properties":{"id":2,"name":"Объект №2"}}]}`,
			strings.Trim(rec.Body.String(), "\n"))
	}
}

package handlers

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_VehicleDict(t *testing.T) {
	endpoint := "/vehicle-dict"

	e := echo.New()
	rec := httptest.NewRecorder()
	req := createReq(http.MethodGet, endpoint, "", nil)
	c := e.NewContext(req, rec)
	h := &Handler{
		DB: &mockStore{},
	}

	if assert.NoError(t, h.GetVehicles(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `[{"gps_id":1,"gos_number":"A777AA750"},{"gps_id":2,"gos_number":"B777BB750"}]`,
			strings.Trim(rec.Body.String(), "\n"))
	}
}

func TestHandler_AddVehicle(t *testing.T) {
	endpoint := "/vehicle-dict"

	data := `{"gps_id":1,"gos_number":"A777AA750"}`

	e := echo.New()
	rec := httptest.NewRecorder()
	req := createReq(http.MethodPost, endpoint, "", strings.NewReader(data))
	c := e.NewContext(req, rec)

	mock := mockStore{}
	h := &Handler{
		DB: &mock,
	}

	if assert.NoError(t, h.AddVehicle(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, `{"gps_id":1,"gos_number":"A777AA750"}`, mock.value.(string))
	}
}

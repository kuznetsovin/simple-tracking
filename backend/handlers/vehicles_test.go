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
	req := createReq(http.MethodGet, endpoint, "")
	c := e.NewContext(req, rec)
	h := &Handler{
		DB: &mockStore{},
	}

	if assert.NoError(t, h.VehicleDict(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `[{"gps_id":1,"gos_number":"A777AA750"},{"gps_id":2,"gos_number":"B777BB750"}]`,
			strings.Trim(rec.Body.String(), "\n"))
	}
}

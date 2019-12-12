package handlers

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_LastPosition(t *testing.T) {
	endpoint := "/last-positions"

	e := echo.New()
	rec := httptest.NewRecorder()
	req := createReq(http.MethodGet, endpoint, "", nil)
	c := e.NewContext(req, rec)
	h := &Handler{
		DB: &mockStore{},
	}

	if assert.NoError(t, h.LastPosition(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[37.37615997,55.69934332]},"properties":{"client":445,"course":53,"nav_time":"2019-11-19T13:47:00Z","packet_id":7242}},{"type":"Feature","geometry":{"type":"Point","coordinates":[37.37656997,55.69921665]},"properties":{"client":156,"course":0,"nav_time":"2019-11-19T13:35:01Z","packet_id":6980}}]}`,
			strings.Trim(rec.Body.String(), "\n"))
	}
}

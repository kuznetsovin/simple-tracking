package handlers

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHandler_ReportObjectDist(t *testing.T) {
	q := make(url.Values)
	endpoint := "/report/object-dist"

	h := &Handler{DB: &mockStore{}}
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	rec := httptest.NewRecorder()

	req := createReq(http.MethodGet, endpoint, q.Encode(), nil)
	c := e.NewContext(req, rec)

	assert.EqualError(t, h.ReportObjectDist(c), "code=400, message=Incorrect client number")

	req = createReq(http.MethodGet, endpoint, q.Encode(), nil)
	c = e.NewContext(req, rec)
	c.SetParamNames("client")
	c.SetParamValues("156")
	assert.EqualError(t, h.ReportObjectDist(c), "code=400, message=Incorrect start date format")

	q.Set("start_date", "2019-11-19T15:00:01Z")
	req = createReq(http.MethodGet, endpoint, q.Encode(), nil)
	c = e.NewContext(req, rec)
	c.SetParamNames("client")
	c.SetParamValues("156")

	assert.EqualError(t, h.ReportObjectDist(c), "code=400, message=Incorrect end date format")

	q.Set("end_date", "2019-11-20T15:00:01Z")
	req = createReq(http.MethodGet, endpoint, q.Encode(), nil)
	c = e.NewContext(req, rec)
	c.SetParamNames("client")
	c.SetParamValues("156")

	if assert.NoError(t, h.ReportObjectDist(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"report":[{"id":1,"name":"object 1","first_point_timestamp":"2019-11-19T13:35:01Z","last_point_timestamp":"2019-11-19T14:35:01Z","mileage":30.5}]}`,
			strings.Trim(rec.Body.String(), "\n"))
	}
}

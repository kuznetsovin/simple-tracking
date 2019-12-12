package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) ReportObjectDist(c echo.Context) error {
	client, err := strconv.Atoi(c.Param("client"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect client number")
	}

	startDate, err := time.Parse(time.RFC3339, c.QueryParam("start_date"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect start date format")
	}

	endDate, err := time.Parse(time.RFC3339, c.QueryParam("end_date"))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Incorrect end date format")
	}

	result := map[string]interface{}{}
	if result["report"], err = h.DB.GetObjectReport(uint32(client), startDate, endDate); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

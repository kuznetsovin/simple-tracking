package handlers

import (
	"github.com/kuznetsovin/go.geojson"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) LastPosition(c echo.Context) error {
	lastPoints, err := h.DB.GetLastVehiclePosition()
	if err != nil {
		return err
	}

	result := geojson.NewFeatureCollection()

	for _, p := range lastPoints {
		f := geojson.NewPointFeature([]float64{p.Longitude, p.Latitude})
		f.Properties = map[string]interface{}{
			"client":    p.Client,
			"packet_id": p.PacketID,
			"nav_time":  p.NavigateDate,
			"course":    p.Course,
		}
		result.AddFeature(f)
	}
	return c.JSON(http.StatusOK, result)
}

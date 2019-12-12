package handlers

import (
	"github.com/kuznetsovin/go.geojson"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GeoObjects(c echo.Context) error {
	objects, err := h.DB.GetGeoObjects()
	if err != nil {
		return err
	}

	result := geojson.NewFeatureCollection()

	for _, o := range objects {
		f := geojson.NewFeature(o.Geom)
		f.Properties = map[string]interface{}{
			"id":   o.Id,
			"name": o.Name,
		}
		result.AddFeature(f)
	}
	return c.JSON(http.StatusOK, result)
}

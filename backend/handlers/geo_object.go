package handlers

import (
	"github.com/kuznetsovin/go.geojson"
	"github.com/labstack/echo"
	"net/http"
	"simple-tracking/backend/models"
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

func (h *Handler) AddObject(c echo.Context) error {
	body := models.GeoObject{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := h.DB.AddObject(body); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

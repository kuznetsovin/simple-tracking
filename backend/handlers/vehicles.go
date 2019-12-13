package handlers

import (
	"github.com/labstack/echo"
	"net/http"
	"simple-tracking/backend/models"
)

func (h *Handler) GetVehicles(c echo.Context) error {
	vehicles, err := h.DB.GetVehicleDict()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, vehicles)
}

func (h *Handler) AddVehicle(c echo.Context) error {
	body := models.VehicleRec{}

	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := h.DB.AddVehicle(body); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, nil)
}

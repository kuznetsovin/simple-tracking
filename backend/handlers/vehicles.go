package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) VehicleDict(c echo.Context) error {
	vehicles, err := h.DB.GetVehicleDict()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, vehicles)
}

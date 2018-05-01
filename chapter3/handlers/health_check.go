package handlers

import (
	"net/http"

	"github.com/PacktPublishing/Echo-Essentials/chapter3/renderings"
	"github.com/labstack/echo"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}

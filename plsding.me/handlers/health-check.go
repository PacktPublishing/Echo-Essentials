package handlers

import (
	"net/http"

	"github.com/PacktPublishing/Echo-Essentials/plsding.me/middlewares"
	"github.com/PacktPublishing/Echo-Essentials/plsding.me/renderings"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}

package handlers

import (
	"net/http"

	"github.com/PacktPublishing/Echo-Essentials/chapter6/middlewares"
	"github.com/PacktPublishing/Echo-Essentials/chapter6/renderings"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	if reqID, ok := c.Get(middlewares.RequestIDContextKey).(uuid.UUID); ok {
		c.Logger().Debugf("RequestID: %s", reqID.String())
	}
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}

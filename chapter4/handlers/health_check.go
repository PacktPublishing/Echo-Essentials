package handlers

import (
	"net/http"

	"github.com/PacktPublishing/Echo-Essentials/chapter4/middlewares"
	"github.com/PacktPublishing/Echo-Essentials/chapter4/renderings"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	c.Logger().Infof("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}

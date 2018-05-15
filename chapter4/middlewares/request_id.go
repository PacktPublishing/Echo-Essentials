package middlewares

import (
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

const (
	RequestIDContextKey = "request_id_context_key"
)

func RequestIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		requestID := uuid.NewV4()
		c.Logger().Infof("RequestID: %s", requestID)
		c.Set(RequestIDContextKey, requestID)
		return next(c)
	})
}

package handlers

import (
	"errors"

	"github.com/labstack/echo"
)

// Error - Example Error Handler
func Error(c echo.Context) error {
	return errors.New("failure!")
}

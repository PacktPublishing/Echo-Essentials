package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/", handler)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

// handler - Simple handler to make sure environment is setup
func handler(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status
	return c.String(http.StatusOK, "Hello World")
}

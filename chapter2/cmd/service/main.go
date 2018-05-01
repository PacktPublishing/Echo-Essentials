package main

import (
	"github.com/PacktPublishing/Echo-Essentials/chapter2/handlers"
	"github.com/PacktPublishing/Echo-Essentials/chapter2/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// Signing Key for our auth middleware
	var signingKey = []byte("superdupersecret!")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.SigningContextKey, signingKey)
			return next(c)
		}
	})

	reminderGroup := e.Group("/reminder")
	reminderGroup.Use(middleware.JWT(signingKey))
	reminderGroup.POST("", handlers.CreateReminder)

	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
	// Authentication routes
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout)
	g := e.Group("/v1")
	g.POST("/login", handlers.Login)
	g.POST("/logout", handlers.Logout)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

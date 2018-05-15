package main

import (
	"html/template"

	"github.com/PacktPublishing/Echo-Essentials/chapter8/bindings"
	"github.com/PacktPublishing/Echo-Essentials/chapter8/handlers"
	"github.com/PacktPublishing/Echo-Essentials/chapter8/middlewares"
	"github.com/PacktPublishing/Echo-Essentials/chapter8/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	StopTestServer = make(chan bool)
	TestRun        = false
)

func main() {
	// create a new echo instance
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Validator = new(bindings.Validator)

	t, err := template.New("reminders").Parse(handlers.RemindersTmpl)
	if err != nil {
		panic(err.Error())
	}

	e.Renderer = &handlers.CustomTemplate{t}

	e.Pre(middlewares.RequestIDMiddleware)

	e.Use(middleware.Logger())  // logger middleware will “wrap” recovery
	e.Use(middleware.Recover()) // as it is enumerated before in the Use calls

	if TestRun {
		e.POST("/stop-test-server", func(ctx echo.Context) error {
			StopTestServer <- true
			return nil
		})
	}

	e.File("/", "static/index.html")

	// Signing Key for our auth middleware
	var signingKey = []byte("superdupersecret!")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.SigningContextKey, signingKey)
			return next(c)
		}
	})

	// in order to serve static assets
	e.Static("/static", "static")

	// reminder handler group
	reminderGroup := e.Group("/reminder")
	reminderGroup.Use(middleware.JWT(signingKey))
	reminderGroup.POST("", handlers.CreateReminder)

	e.GET("/render", handlers.RenderReminders)
	e.GET("/render-more", handlers.RenderMoreReminders)

	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
	e.GET("/error", handlers.Error)
	// Authentication routes
	e.POST("/login", handlers.Login).Name = "login"
	e.POST("/logout", handlers.Logout)

	// V1 Routes
	v1 := e.Group("/v1")
	// V1 Authentication routes
	v1.POST("/login", handlers.Login)
	v1.POST("/logout", handlers.Logout)
	// V1 Reminder Routes
	v1Reminders := v1.Group("/reminder", middleware.JWT(signingKey))
	v1Reminders.Use(middleware.JWT(signingKey))
	v1Reminders.POST("", handlers.CreateReminder)
	v1Reminders.GET("/:id", handlers.GetReminder)
	// /v1/reminder/:id

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}

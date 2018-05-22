package main

import (
	"database/sql"

	"github.com/PacktPublishing/Echo-Essentials/chapter5/bindings"
	"github.com/PacktPublishing/Echo-Essentials/chapter5/handlers"
	"github.com/PacktPublishing/Echo-Essentials/chapter5/middlewares"
	"github.com/PacktPublishing/Echo-Essentials/chapter5/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// create a new echo instance
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Validator = new(bindings.Validator)

	e.Pre(middlewares.RequestIDMiddleware)

	e.Use(middleware.Logger())  // logger middleware will “wrap” recovery
	e.Use(middleware.Recover()) // as it is enumerated before in the Use calls

	e.GET("/", HandlerFunction, Middleware1, Middleware2, Middleware3)
	// RouteHandler = Middleware1(Middleware2(Middleware3(HandlerFunction)

	// Signing Key for our auth middleware
	var signingKey = []byte("superdupersecret!")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.SigningContextKey, signingKey)
			return next(c)
		}
	})

	// add database to context
	db, err := sql.Open("sqlite3", "./service.db")
	if err != nil {
		log.Fatalf("error opening database: %v\n", err)
	}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(models.DBContextKey, db)
			return next(c)
		}
	})

	// in order to serve static assets
	e.Static("/static", "static")

	// reminder handler group
	reminderGroup := e.Group("/reminder")
	reminderGroup.Use(middleware.JWT(signingKey))
	reminderGroup.POST("", handlers.CreateReminder)

	// Route / to handler function
	e.GET("/health-check", handlers.HealthCheck)
	// Authentication routes
	e.POST("/login", handlers.Login)
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

func HandlerFunction(ctx echo.Context) error {
	ctx.Logger().Info("in handler function")
	return nil
}

func middlewareFunc(i int) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Logger().Infof("middleware #%d start!", i)
			next(ctx)
			ctx.Logger().Infof("middleware #%d end!", i)
			return nil
		}
	}
}

var (
	Middleware1 = middlewareFunc(1)
	Middleware2 = middlewareFunc(2)
	Middleware3 = middlewareFunc(3)
)

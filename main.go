package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "knit2024/docs"
	"knit2024/infrastructure/controller"
	"knit2024/infrastructure/custom_error"
	"knit2024/infrastructure/log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	e := echo.New()

	logger := log.NewLogger("INFO")

	e.Logger = logger

	// Middlewares
	registerMiddlewares(e)

	// Controller
	controller.NewKnitController(e)

	// HealthCheck
	registerHealthCheck(e)

	// Swagger
	registerSwaggerRedirect(e)

	e.HTTPErrorHandler = custom_error.CustomEchoHTTPErrorHandler

	// Start Server
	go func() {
		if err := e.Start(":8070"); err != nil {
			if !strings.Contains(err.Error(), "client: Server closed") {
				logger.Fatal(err)
			}
		}
	}()
	serverChannel := make(chan struct{})

	// Stop Server
	go func() {
		sigint := make(chan os.Signal, 1)
		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		close(serverChannel)
	}()
	<-serverChannel
}

func registerSwaggerRedirect(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/swagger/index.html")
	})
}

func registerHealthCheck(e *echo.Echo) {
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
}

func registerMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:*"},
		AllowMethods: []string{echo.POST},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
}

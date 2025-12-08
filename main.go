package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	healthPath := os.Getenv("HEALTH_PATH")
	if healthPath == "" {
		healthPath = "/health"
	}
	slog.Info("Health check path set to", "path", healthPath)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	slog.Info("Server starting", "port", port, "healthPath", healthPath)

	e := echo.New()
	e.GET(healthPath, func(c echo.Context) error {
		return c.String(200, "OK")
	})
	e.Logger.Fatal(e.Start(":" + port))
}

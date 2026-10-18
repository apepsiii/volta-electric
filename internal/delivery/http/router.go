package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterRoutes registers all middleware and endpoints to the Echo instance
func RegisterRoutes(e *echo.Echo, h *Handler, staticDir string) {
	// Global Middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} (${latency_human})\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// Static Assets
	e.Static("/static", staticDir)
	e.File("/favicon.ico", staticDir+"/img/voltra-logo.svg")

	// Page Routes
	e.GET("/", h.HandleHome)
	e.GET("/health", h.HandleHealth)

	// API Routes
	api := e.Group("/api")
	{
		api.GET("/products", h.HandleGetProducts)
		api.GET("/products/:id", h.HandleGetProductDetail)
		api.POST("/inquiry", h.HandleSubmitInquiry)
	}
}


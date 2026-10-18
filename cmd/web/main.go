package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"voltra-electric/config"
	delivery "voltra-electric/internal/delivery/http"
	"voltra-electric/internal/repository"
)

func main() {
	// 1. Load Application Configuration
	cfg := config.Load()

	// 2. Initialize Echo instance
	e := echo.New()
	e.HideBanner = true

	// 3. Initialize Template Renderer
	renderer, err := delivery.NewTemplateRenderer("web/templates")
	if err != nil {
		e.Logger.Fatalf("Failed to initialize template renderer: %v", err)
	}
	e.Renderer = renderer

	// 4. Initialize Repository and Handlers
	repo := repository.NewMemoryRepository()
	handler := delivery.NewHandler(cfg, repo)

	// 5. Register Routes & Middlewares
	delivery.RegisterRoutes(e, handler, "web/static")

	// 6. Start Server in Goroutine with Graceful Shutdown
	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		fmt.Printf("⚡ [VOLTRA ELECTRIC] Server running on http://localhost%s (%s)\n", addr, cfg.Domain)
		fmt.Printf("🏢 %s - Brand: %s\n", cfg.CompanyName, cfg.BrandName)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatalf("Shutting down the server: %v", err)
		}
	}()

	// 7. Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("Server exited cleanly.")
}


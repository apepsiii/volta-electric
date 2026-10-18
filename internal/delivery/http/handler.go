package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"voltra-electric/config"
	"voltra-electric/internal/domain"
	"voltra-electric/internal/repository"
)

// Handler handles HTTP requests
type Handler struct {
	cfg  *config.Config
	repo repository.Repository
}

// NewHandler creates a new HTTP handler instance
func NewHandler(cfg *config.Config, repo repository.Repository) *Handler {
	return &Handler{
		cfg:  cfg,
		repo: repo,
	}
}

// HomePageData represents the ViewModel passed to the index template
type HomePageData struct {
	Config     *config.Config
	Company    domain.CompanyProfile
	Metrics    []domain.MetricStat
	Products   []domain.Product
	Components []domain.ComponentDetail
	Testing    []domain.TestingProcedure
	PageTitle  string
}

// HandleHome renders the main landing page
func (h *Handler) HandleHome(c echo.Context) error {
	data := HomePageData{
		Config:     h.cfg,
		Company:    h.repo.GetCompanyProfile(),
		Metrics:    h.repo.GetMetrics(),
		Products:   h.repo.GetProducts(),
		Components: h.repo.GetComponents(),
		Testing:    h.repo.GetTestingProcedures(),
		PageTitle:  "VOLTRA - Solusi Jaringan Distribusi Trafo Berkualitas | PT Mega Karya Perkasa",
	}

	return c.Render(http.StatusOK, "index.html", data)
}

// HandleGetProducts returns the products as JSON
func (h *Handler) HandleGetProducts(c echo.Context) error {
	products := h.repo.GetProducts()
	return c.JSON(http.StatusOK, echo.Map{
		"status":   "success",
		"products": products,
	})
}

// HandleGetProductDetail returns a specific product detail as JSON
func (h *Handler) HandleGetProductDetail(c echo.Context) error {
	id := c.Param("id")
	product, found := h.repo.GetProductByID(id)
	if !found {
		return c.JSON(http.StatusNotFound, echo.Map{
			"status":  "error",
			"message": "Product not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status":  "success",
		"product": product,
	})
}

// HandleSubmitInquiry processes an inquiry form and returns the WhatsApp URL
func (h *Handler) HandleSubmitInquiry(c echo.Context) error {
	req := new(domain.InquiryRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  "error",
			"message": "Invalid request parameters",
		})
	}

	waURL := req.GenerateWhatsAppURL(h.cfg.ContactWA)

	// If requested as JSON (AJAX)
	if c.Request().Header.Get("Accept") == "application/json" {
		return c.JSON(http.StatusOK, echo.Map{
			"status":       "success",
			"whatsapp_url": waURL,
		})
	}

	// Otherwise redirect directly to WhatsApp
	return c.Redirect(http.StatusSeeOther, waURL)
}

// HandleHealth checks application health status
func (h *Handler) HandleHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":      "ok",
		"app":         "voltra-electric",
		"company":     h.cfg.CompanyName,
		"brand":       h.cfg.BrandName,
		"environment": h.cfg.Environment,
	})
}


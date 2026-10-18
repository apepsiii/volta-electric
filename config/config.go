package config

import (
	"os"
)

// Config holds all server and application settings
type Config struct {
	Port        string
	Domain      string
	Environment string
	CompanyName string
	BrandName   string
	ContactWA   string
	ContactTelp string
	Address     string
}

// Load reads configuration from environment variables or defaults
func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		Port:        port,
		Domain:      "voltra-electric.id",
		Environment: env,
		CompanyName: "PT. MEGA KARYA PERKASA",
		BrandName:   "VOLTRA",
		ContactWA:   "6281386698222",
		ContactTelp: "0813-8669-8222",
		Address:     "Indonesia",
	}
}


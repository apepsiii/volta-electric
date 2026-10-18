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

	contactWA := os.Getenv("CONTACT_WA")
	if contactWA == "" {
		contactWA = "18732768961"
	}

	contactTelp := os.Getenv("CONTACT_TELP")
	if contactTelp == "" {
		contactTelp = "+1 (873) 276-8961"
	}

	return &Config{
		Port:        port,
		Domain:      "voltra-electric.id",
		Environment: env,
		CompanyName: "PT. MEGA KARYA PERKASA",
		BrandName:   "VOLTRA",
		ContactWA:   contactWA,
		ContactTelp: contactTelp,
		Address:     "Indonesia",
	}
}


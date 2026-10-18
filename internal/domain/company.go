package domain

// CompanyProfile represents corporate information from the PDF profile
type CompanyProfile struct {
	LegalName      string   `json:"legal_name"`
	BrandName      string   `json:"brand_name"`
	EstablishedYear int     `json:"established_year"`
	CoreBusiness   string   `json:"core_business"`
	Standards      []string `json:"standards"`
	Mission        string   `json:"mission"`
	QualitySystem  string   `json:"quality_system"`
}

// MetricStat represents a numeric highlight in the UI
type MetricStat struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Sublabel    string `json:"sublabel"`
	Highlight   string `json:"highlight"`
}

// ComponentDetail represents a technical transformer component
type ComponentDetail struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	KeyPoints   []string `json:"key_points"`
	ImageURLs   []string `json:"image_urls"`
}


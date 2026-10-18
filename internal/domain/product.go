package domain

// Product represents a transformer model manufactured by PT Mega Karya Perkasa
type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Series      string            `json:"series"`
	Tag         string            `json:"tag"`
	Category    string            `json:"category"`
	Subtitle    string            `json:"subtitle"`
	Description string            `json:"description"`
	Capacity    string            `json:"capacity"`
	ImageURL    string            `json:"image_url"`
	Badge       string            `json:"badge"`
	Features    []string          `json:"features"`
	Specs       map[string]string `json:"specs"`
}

// TestingProcedure represents the 9 stages of factory acceptance tests from PDF
type TestingProcedure struct {
	StepNumber  int    `json:"step_number"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Standard    string `json:"standard"`
	IsCrucial   bool   `json:"is_crucial"`
}


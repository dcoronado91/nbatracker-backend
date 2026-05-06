package models

type Team struct {
	Name          string `json:"name"`
	City          string `json:"city"`
	Abbreviation  string `json:"abbreviation"`
	Championships int    `json:"championships"`
	LogoURL       string `json:"logo_url"`
	Conference    string `json:"conference"`
	Division      string `json:"division"`
	CreatedAt     string `json:"created_at"`
}

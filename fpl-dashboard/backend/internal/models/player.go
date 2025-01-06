package models

type Player struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Team           int     `json:"team"`
	Position       string  `json:"position"`
	CurrentPrice   float64 `json:"current_price"`
	TotalPoints    int     `json:"total_points"`
	SelectedBy     float64 `json:"selected_by_percent"`
	Form           float64 `json:"form"`
	ExpectedPoints float64 `json:"expected_points"`
}

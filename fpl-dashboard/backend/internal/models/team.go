package models

import "time"

type Team struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Players      []Player  `json:"players"`
	TotalValue   float64   `json:"total_value"`
	TransferCost int       `json:"transfer_cost"`
	LastUpdated  time.Time `json:"last_updated"`
}

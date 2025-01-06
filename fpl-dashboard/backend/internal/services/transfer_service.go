package services

import (
	"sort"

	"github.com/fpl-dashboard/internal/client"
	"github.com/fpl-dashboard/internal/models"
)

type TransferService struct {
	fplClient *client.FPLClient
}

type TransferSuggestion struct {
	PlayerID          int     `json:"player_id"`
	PlayerName        string  `json:"player_name"`
	CurrentTeam       int     `json:"current_team"`
	TransferScore     float64 `json:"transfer_score"`
	RecommendedAction string  `json:"recommended_action"`
}

func NewTransferService(client *client.FPLClient) *TransferService {
	return &TransferService{
		fplClient: client,
	}
}

func (s *TransferService) AnalyzeTransfers(userTeam []models.Player) ([]TransferSuggestion, error) {
	suggestions := []TransferSuggestion{}

	for _, player := range userTeam {
		suggestion := TransferSuggestion{
			PlayerID:   player.ID,
			PlayerName: player.Name,
		}

		// Calculate transfer score based on multiple factors
		suggestion.TransferScore = calculateTransferScore(player)

		// Determine recommended action
		if suggestion.TransferScore > 7.5 {
			suggestion.RecommendedAction = "STRONG_BUY"
		} else if suggestion.TransferScore > 5 {
			suggestion.RecommendedAction = "CONSIDER"
		} else {
			suggestion.RecommendedAction = "HOLD"
		}

		suggestions = append(suggestions, suggestion)
	}

	// Sort suggestions by transfer score
	sort.Slice(suggestions, func(i, j int) bool {
		return suggestions[i].TransferScore > suggestions[j].TransferScore
	})

	return suggestions, nil
}

func calculateTransferScore(player models.Player) float64 {
	// Complex scoring mechanism
	score := 0.0

	// Form factor
	score += player.Form * 1.5

	// Points factor
	score += float64(player.TotalPoints) / 10

	// Price changes consideration
	if player.CurrentPrice < 6.0 {
		score += 1.0
	}

	// Popularity factor (inverse of selected by percentage)
	score += (100 - player.SelectedBy) / 10

	return score
}

package services

import (
	"sort"
	"time"

	"github.com/fpl-dashboard/internal/client"
	"github.com/fpl-dashboard/internal/models"
)

type PerformanceService struct {
	fplClient *client.FPLClient
}

type PlayerPerformance struct {
	Player            models.Player    `json:"player"`
	Form              float64          `json:"form"`
	ExpectedPoints    float64          `json:"expected_points"`
	UpcomingFixtures  []models.Fixture `json:"upcoming_fixtures"`
	RecentPerformance []struct {
		Gameweek int     `json:"gameweek"`
		Points   int     `json:"points"`
		Price    float64 `json:"price"`
	} `json:"recent_performance"`
}

type TeamAnalysis struct {
	Team             models.Team         `json:"team"`
	TopPerformers    []PlayerPerformance `json:"top_performers"`
	UnderPerformers  []PlayerPerformance `json:"under_performers"`
	TeamValue        float64             `json:"team_value"`
	PointsPerMillion float64             `json:"points_per_million"`
	LastUpdated      time.Time           `json:"last_updated"`
}

func NewPerformanceService(client *client.FPLClient) *PerformanceService {
	return &PerformanceService{
		fplClient: client,
	}
}

func (s *PerformanceService) AnalyzeTeamPerformance(teamID int) (*TeamAnalysis, error) {
	// Get team details
	team, err := s.fplClient.GetTeamDetails(teamID)
	if err != nil {
		return nil, err
	}

	// Get player performances
	var performances []PlayerPerformance
	for _, player := range team.Players {
		perf, err := s.getPlayerPerformance(player.ID)
		if err != nil {
			continue // Skip failed player analyses
		}
		performances = append(performances, *perf)
	}

	// Sort by form
	sort.Slice(performances, func(i, j int) bool {
		return performances[i].Form > performances[j].Form
	})

	// Calculate team metrics
	totalPoints := 0
	totalValue := 0.0
	for _, p := range performances {
		totalPoints += p.Player.TotalPoints
		totalValue += p.Player.CurrentPrice
	}

	analysis := &TeamAnalysis{
		Team:             *team,
		TopPerformers:    performances[:3],                   // Top 3 performers
		UnderPerformers:  performances[len(performances)-3:], // Bottom 3 performers
		TeamValue:        totalValue,
		PointsPerMillion: float64(totalPoints) / totalValue,
		LastUpdated:      time.Now(),
	}

	return analysis, nil
}

func (s *PerformanceService) getPlayerPerformance(playerID int) (*PlayerPerformance, error) {
	player, err := s.fplClient.GetPlayerDetails(playerID)
	if err != nil {
		return nil, err
	}

	fixtures, err := s.fplClient.GetUpcomingFixtures()
	if err != nil {
		return nil, err
	}

	// Filter fixtures for this player's team
	var playerFixtures []models.Fixture
	for _, fixture := range fixtures {
		if fixture.HomeTeam == player.Team || fixture.AwayTeam == player.Team {
			playerFixtures = append(playerFixtures, fixture)
		}
	}

	performance := &PlayerPerformance{
		Player:           *player,
		Form:             player.Form,
		ExpectedPoints:   player.ExpectedPoints,
		UpcomingFixtures: playerFixtures[:5], // Next 5 fixtures
	}

	return performance, nil
}

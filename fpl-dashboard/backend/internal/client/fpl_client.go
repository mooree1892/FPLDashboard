package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fpl-dashboard/internal/models"
)

type FPLClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewFPLClient() *FPLClient {
	return &FPLClient{
		baseURL: "https://fantasy.premierleague.com/api",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *FPLClient) GetPlayerDetails(playerID int) (*models.Player, error) {
	url := fmt.Sprintf("%s/element-summary/%d/", c.baseURL, playerID)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch player details: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var player models.Player
	if err := json.Unmarshal(body, &player); err != nil {
		return nil, fmt.Errorf("failed to parse player data: %v", err)
	}

	return &player, nil
}

func (c *FPLClient) GetUpcomingFixtures() ([]models.Fixture, error) {
	url := fmt.Sprintf("%s/fixtures", c.baseURL)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fixtures: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var fixtures []models.Fixture
	if err := json.Unmarshal(body, &fixtures); err != nil {
		return nil, fmt.Errorf("failed to parse fixtures data: %v", err)
	}

	return fixtures, nil
}

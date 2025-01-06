package models

import "time"

type Fixture struct {
	ID            int       `json:"id"`
	GameweekID    int       `json:"gameweek"`
	HomeTeam      int       `json:"home_team"`
	AwayTeam      int       `json:"away_team"`
	HomeTeamScore int       `json:"home_team_score,omitempty"`
	AwayTeamScore int       `json:"away_team_score,omitempty"`
	Difficulty    int       `json:"difficulty"`
	KickoffTime   time.Time `json:"kickoff_time"`
	IsFinished    bool      `json:"is_finished"`
}

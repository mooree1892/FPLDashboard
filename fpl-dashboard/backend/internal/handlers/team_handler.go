package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/fpl-dashboard/internal/models"
	"github.com/fpl-dashboard/internal/services"
	"github.com/gorilla/mux"
)

type TeamHandler struct {
	performanceService *services.PerformanceService
}

func NewTeamHandler(service *services.PerformanceService) *TeamHandler {
	return &TeamHandler{
		performanceService: service,
	}
}

func (h *TeamHandler) GetTeamDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	analysis, err := h.performanceService.AnalyzeTeamPerformance(teamID)
	if err != nil {
		http.Error(w, "Failed to analyze team performance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}

func (h *TeamHandler) GetTeamPlayers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamID, err := strconv.Atoi(vars["teamId"])
	if err != nil {
		http.Error(w, "Invalid team ID", http.StatusBadRequest)
		return
	}

	// Get query parameters for filtering
	position := r.URL.Query().Get("position")
	sortBy := r.URL.Query().Get("sort")

	analysis, err := h.performanceService.AnalyzeTeamPerformance(teamID)
	if err != nil {
		http.Error(w, "Failed to get team players: "+err.Error(), http.StatusInternalServerError)
		return
	}

	players := analysis.Team.Players

	// Apply filters
	if position != "" {
		var filtered []models.Player
		for _, p := range players {
			if p.Position == position {
				filtered = append(filtered, p)
			}
		}
		players = filtered
	}

	// Apply sorting
	switch sortBy {
	case "points":
		sort.Slice(players, func(i, j int) bool {
			return players[i].TotalPoints > players[j].TotalPoints
		})
	case "form":
		sort.Slice(players, func(i, j int) bool {
			return players[i].Form > players[j].Form
		})
	case "price":
		sort.Slice(players, func(i, j int) bool {
			return players[i].CurrentPrice > players[j].CurrentPrice
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

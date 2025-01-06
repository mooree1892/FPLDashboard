package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fpl-dashboard/internal/models"
	"github.com/fpl-dashboard/internal/services"
)

type TransferHandler struct {
	transferService *services.TransferService
}

func NewTransferHandler(service *services.TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: service,
	}
}

func (h *TransferHandler) GetTransferSuggestions(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, you'd parse the user's team from the request
	// For this example, we'll use a mock team
	mockTeam := []models.Player{
		// Populate with some player data
	}

	suggestions, err := h.transferService.AnalyzeTransfers(mockTeam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}

func (h *TransferHandler) GetTransferAnalysis(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	budget, err := strconv.ParseFloat(r.URL.Query().Get("budget"), 64)
	if err != nil {
		budget = 0.0 // Default to 0 if not specified
	}

	freeTransfers, err := strconv.Atoi(r.URL.Query().Get("free_transfers"))
	if err != nil {
		freeTransfers = 1 // Default to 1 if not specified
	}

	// For now, using mock team like in GetTransferSuggestions
	mockTeam := []models.Player{
		// Populate with some player data
	}

	analysis, err := h.transferService.AnalyzeOptimalTransfers(mockTeam, services.TransferOptions{
		Budget:        budget,
		FreeTransfers: freeTransfers,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}

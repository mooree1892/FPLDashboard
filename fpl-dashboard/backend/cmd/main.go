package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/fpl-dashboard/internal/client"
	"github.com/fpl-dashboard/internal/handlers"
	"github.com/fpl-dashboard/internal/services"
)

func main() {
	// Initialize router
	router := mux.NewRouter()

	// Initialize services and handlers
	fplClient := client.NewFPLClient()
	transferService := services.NewTransferService(fplClient)
	performanceService := services.NewPerformanceService(fplClient)

	// Initialize handlers
	transferHandler := handlers.NewTransferHandler(transferService)
	teamHandler := handlers.NewTeamHandler(performanceService)

	// API routes
	api := router.PathPrefix("/api").Subrouter()

	// Team routes
	api.HandleFunc("/team/{teamId}", teamHandler.GetTeamDetails).Methods("GET")
	api.HandleFunc("/team/{teamId}/players", teamHandler.GetTeamPlayers).Methods("GET")

	// Transfer routes
	api.HandleFunc("/transfers/suggestions", transferHandler.GetTransferSuggestions).Methods("GET")
	api.HandleFunc("/transfers/analysis", transferHandler.GetTransferAnalysis).Methods("GET")

	// CORS configuration
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Create server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsOpts.Handler(router),
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

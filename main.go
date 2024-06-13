package main

import (
	"foundation-app/internal/api/handlers"
	models "foundation-app/internal/models"
	inmem "foundation-app/pkg/store"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

/*
func main method is the entry point that initializes in-memory stores
and orchestrates API server with the initialised stores
*/
func main() {
	// Initialize in-memory stores
	emailStore := inmem.NewEmailStore()
	foundationStore := inmem.NewFoundationStore()
	nonprofitStore := inmem.NewNonprofitStore()

	// Create sample foundations
	sampleFoundations := []models.Foundation{
		{
			ID:    uuid.New(),
			Email: "foundation1@example.com",
			Name:  "Foundation 1",
		},
		{
			ID:    uuid.New(),
			Email: "foundation2@example.com",
			Name:  "Foundation 2",
		},
		{
			ID:    uuid.New(),
			Email: "foundation3@example.com",
			Name:  "Foundation 3",
		},
	}

	// Add sample foundations to the store
	for _, foundation := range sampleFoundations {
		_, err := foundationStore.CreateFoundation(foundation)
		if err != nil {
			log.Fatalf("Failed to create foundation: %v", err)
		}
	}

	// Create the router
	r := mux.NewRouter()

	// Create API handlers
	apiHandlers := handlers.NewAPIHandlers(emailStore, nonprofitStore)

	// Define API routes
	r.HandleFunc("/nonprofits", apiHandlers.CreateNonprofit).Methods(http.MethodPost)
	r.HandleFunc("/emails", apiHandlers.BulkSendEmails).Methods(http.MethodPost)
	r.HandleFunc("/emails", apiHandlers.GetAllEmails).Methods(http.MethodGet)

	// Start the server
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

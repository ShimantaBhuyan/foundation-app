// This package contains the API handlers for the Foundation App. The handlers
// serve various HTTP endpoints for interaction between users and the
// Foundation App backend.
package handlers

import (
	"foundation-app/internal/store"
)

type APIHandlers struct {
	emailStore     store.EmailStore
	nonprofitStore store.NonprofitStore
}

func NewAPIHandlers(emailStore store.EmailStore, nonprofitStore store.NonprofitStore) *APIHandlers {
	return &APIHandlers{
		emailStore:     emailStore,
		nonprofitStore: nonprofitStore,
	}
}

package store

import (
	model "github.com/ShimantaBhuyan/foundation-app/internal/models"

	"github.com/google/uuid"
)

// NonprofitStore interface defines the necessary functions for non-profit-related
// data interactions.
type NonprofitStore interface {
	// CreateNonprofit creates a new non-profit in the system by taking in
	// a NonProfit model as a parameter and returning a new UUID for
	// the non-profit.
	CreateNonprofit(nonprofit model.NonProfit) (uuid.UUID, error)

	// GetNonprofits returns all non-profits present in the system.
	GetNonprofits() ([]model.NonProfit, error)

	// GetNonprofitsByEmail returns all non-profits present in the system matching
	// the given email addresses.
	GetNonprofitsByEmail(emails []string) ([]model.NonProfit, error)
}

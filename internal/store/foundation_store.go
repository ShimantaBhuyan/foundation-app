package store

import (
	model "github.com/ShimantaBhuyan/foundation-app/internal/models"

	"github.com/google/uuid"
)

// FoundationStore interface defines the necessary functions for the
// foundation-related data interactions.
type FoundationStore interface {
	// CreateFoundation creates a new foundation in the system by taking in a
	// Foundation model as a parameter and returning a new UUID for
	// the foundation.
	CreateFoundation(foundation model.Foundation) (uuid.UUID, error)
}

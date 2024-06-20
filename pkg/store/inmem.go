package externalstore

import (
	"sync"

	models "github.com/ShimantaBhuyan/foundation-app/internal/models"
	"github.com/ShimantaBhuyan/foundation-app/internal/store"
	emailClient "github.com/ShimantaBhuyan/foundation-app/pkg/email"

	"github.com/google/uuid"
)

// InMemEmailStore is an in-memory implementation of the EmailStore interface
type InMemEmailStore struct {
	sentEmails map[uuid.UUID]models.SentEmails
	emails     map[uuid.UUID]models.Email
	mu         sync.RWMutex
	emailCli   *emailClient.Client // Email client instance
}

// NewEmailStore creates a new in-memory email store
func NewEmailStore() store.EmailStore {
	return &InMemEmailStore{
		sentEmails: make(map[uuid.UUID]models.SentEmails),
		emails:     make(map[uuid.UUID]models.Email),
		emailCli:   emailClient.NewClient(), // Initialize email client
	}
}

// BulkSendEmails sends emails to a list of nonprofits
func (s *InMemEmailStore) BulkSendEmails(emails map[string]models.Email) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for emailAddr, email := range emails {
		// Send email using the email client
		err := s.emailCli.SendEmail(email, emailAddr)
		if err != nil {
			return err
		}
		id := uuid.New()
		s.sentEmails[id] = models.SentEmails{
			EmailID:   id,
			Recipient: emailAddr,
			Cc:        email.Cc,
			Bcc:       email.Bcc,
		}
		s.emails[id] = email

	}

	return nil
}

// GetAllEmails returns all the emails sent to nonprofits
func (s *InMemEmailStore) GetAllEmails() ([]models.Email, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	emails := make([]models.Email, 0, len(s.emails))
	for _, email := range s.emails {
		emails = append(emails, email)
	}

	return emails, nil
}

// InMemFoundationStore is an in-memory implementation of the FoundationStore interface
type InMemFoundationStore struct {
	foundations map[uuid.UUID]models.Foundation
	mu          sync.RWMutex
}

// NewFoundationStore creates a new in-memory foundation store
func NewFoundationStore() store.FoundationStore {
	return &InMemFoundationStore{
		foundations: make(map[uuid.UUID]models.Foundation),
	}
}

// CreateFoundation creates a new foundation
func (s *InMemFoundationStore) CreateFoundation(foundation models.Foundation) (uuid.UUID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New()
	s.foundations[id] = foundation

	return id, nil
}

// InMemNonprofitStore is an in-memory implementation of the NonprofitStore interface
type InMemNonprofitStore struct {
	nonprofits map[uuid.UUID]models.NonProfit
	mu         sync.RWMutex
}

// NewNonprofitStore creates a new in-memory nonprofit store
func NewNonprofitStore() store.NonprofitStore {
	return &InMemNonprofitStore{
		nonprofits: make(map[uuid.UUID]models.NonProfit),
	}
}

// CreateNonprofit creates a new nonprofit
func (s *InMemNonprofitStore) CreateNonprofit(nonprofit models.NonProfit) (uuid.UUID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New()
	s.nonprofits[id] = nonprofit

	return id, nil
}

// GetNonprofits retrieves all the nonprofits
func (s *InMemNonprofitStore) GetNonprofits() ([]models.NonProfit, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nonprofits := make([]models.NonProfit, 0, len(s.nonprofits))
	for _, nonprofit := range s.nonprofits {
		nonprofits = append(nonprofits, nonprofit)
	}

	return nonprofits, nil
}

// GetNonprofitsByEmail retrieves all the nonprofits by email address
func (s *InMemNonprofitStore) GetNonprofitsByEmail(emails []string) ([]models.NonProfit, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	nonprofits := make([]models.NonProfit, 0, len(s.nonprofits))

	emailToNonprofit := make(map[string]models.NonProfit, len(s.nonprofits))
	for _, nonprofit := range s.nonprofits {
		emailToNonprofit[nonprofit.Email] = nonprofit
	}

	for _, email := range emails {
		if nonprofit, ok := emailToNonprofit[email]; ok {
			nonprofits = append(nonprofits, nonprofit)
		}
	}

	return nonprofits, nil
}

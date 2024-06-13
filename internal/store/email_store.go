package store

import (
	model "foundation-app/internal/models"
)

/*
EmailStore interface represents the necessary functions to interact with the email-related data.
*/
type EmailStore interface {
	// BulkSendEmails sends emails to multiple recipients
	BulkSendEmails(emails map[string]model.Email) error

	// GetAllEmails returns all emails present in the system.
	GetAllEmails() ([]model.Email, error)
}

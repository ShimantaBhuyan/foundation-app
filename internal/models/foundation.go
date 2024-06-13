package model

import "github.com/google/uuid"

// Foundation model represents the structure of a Foundation in our system
type Foundation struct {
	ID    uuid.UUID
	Email string
	Name  string
}

// Summary:
//   A Foundation object, which contains the ID, Name, and Email
//	for the Foundation.
//
// Fields:
//	- ID (uuid.UUID) : A unique identifier generated as a UUID for the Foundation.
//	- Email (string) : The email address associated with the Foundation.
//	- Name (string) : The name given to the Foundation.
//

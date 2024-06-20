package model

import "github.com/google/uuid"

// Email model represents an email in our system
type Email struct {
	ID      uuid.UUID `json:"id"`
	Subject string
	Body    string
	Cc      string
	Bcc     string
}

// Summary:
//   An Email object, which contains the ID, Subject, and Body for an Email
//	message. The ID value uniquely represents the Email message.
//
// Fields:
//	- ID (uuid.UUID) : A unique identifier generated as a UUID
//	  for the Email message.
//	- Subject (string): A short summary or topic of the Email message.
//	- Body (string): The main content or text of the Email message.
//

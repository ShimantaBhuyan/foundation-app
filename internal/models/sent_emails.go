package model

import "github.com/google/uuid"

// SentEmails struct represents the information of sent emails in our system
type SentEmails struct {
	EmailID   uuid.UUID
	Recipient string
	Cc        string
	Bcc       string
}

// Summary:
//   SentEmails struct represents the summary information about an email
//   that has been previously sent.
//
// Fields:
//	- EmailID (uuid.UUID) : A unique identifier generated as a UUID for the sent email which refers the Email model.
//	- Recipient (string) : The email address of the recipient.
//

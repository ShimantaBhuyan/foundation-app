// internal/api/dtos/email_dto.go
package dtos

// BulkSendEmailsDTO data transfer object
/**
  Summary:
    The BulkSendEmailsDTO struct contains the information required for sending emails
    to a list of email recipients.

  Fields:
    Subject - The subject of the email
    TemplateString - The raw body template of the email
    Recipients - The list of email addresses that should receive the email
*/
type BulkSendEmailsDTO struct {
	Subject        string   `json:"subject"`
	TemplateString string   `json:"templateString"`
	Recipients     []string `json:"recipients"`
	CcRecipients   string   `json:"cc"`
	BccRecipients  string   `json:"bcc"`
}

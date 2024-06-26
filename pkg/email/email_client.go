package email

import (
	"fmt"

	model "github.com/ShimantaBhuyan/foundation-app/internal/models"
)

// Client is a dummy email client implementation
//
// Client struct is a simple and dummy implementation of the EmailStore
type Client struct{}

// NewClient returns a new instance of the email client
func NewClient() *Client {
	return &Client{}
}

// SendEmail sends an email to the specified email address
func (c *Client) SendEmail(email model.Email, to string) error {
	fmt.Printf("EMAIL CLIENT | TO: %s\nBODY: %s\n", to, email.Body)
	fmt.Printf("CC'ed TO: %s\n", email.Cc)
	fmt.Printf("BCC'ed TO: %s\n", email.Bcc)

	return nil
}

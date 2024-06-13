package handlers

import (
	"encoding/json"
	"net/http"
)

// GetAllEmails retrieves all the emails sent to nonprofits
/**
  Summary:
    The GetAllEmails function handles the request for getting all the emails
    sent to nonprofits. It uses the email store to get all the emails and returns
    them as a JSON response to the client.

  Returns:
    Returns the list of all emails sent to nonprofits as a JSON response.
*/
func (h *APIHandlers) GetAllEmails(w http.ResponseWriter, r *http.Request) {
	emails, err := h.emailStore.GetAllEmails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

package handlers

import (
	"encoding/json"
	"net/http"
)

// GetAllEmails handles the request for getting all the emails sent to nonprofits.
//
// It uses the email store to get all the emails and returns them as a JSON
// response to the client.
//
// Returns:
//
//	Success: true
//	Data:    list of all emails sent to nonprofits
//	Error:   ""
//
// Error Responses:
//
//	500 - internal server error
func (h *APIHandlers) GetAllEmails(w http.ResponseWriter, r *http.Request) {
	emails, err := h.emailStore.GetAllEmails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := APIResponse{
		Success: true,
		Data:    emails,
		Error:   "",
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"foundation-app/internal/dtos"
	models "foundation-app/internal/models"

	"github.com/google/uuid"
)

type BulkSendEmailResponse struct {
	Message string `json:"message"`
}

// BulkSendEmails sends emails to a list of nonprofits
/**
  Summary:
    This function handles the request for sending emails to a list of nonprofits. It receives
    a list of email addresses with the associated template values, maps them to models.Email,
    and sends the email using the email store.

  Returns:
    A JSON response containing the following:
    - Message: A success message stating all the emails have been sent.
*/
func (h *APIHandlers) BulkSendEmails(w http.ResponseWriter, r *http.Request) {
	var bulkSendRequestDTO dtos.BulkSendEmailsDTO
	err := json.NewDecoder(r.Body).Decode(&bulkSendRequestDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nonprofits, err := h.nonprofitStore.GetNonprofitsByEmail(bulkSendRequestDTO.Recipients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templateVariables := make(map[string](map[string]string), len(nonprofits))
	for _, nonprofit := range nonprofits {
		templateVariable := map[string]string{}
		templateVariable["name"] = nonprofit.Name
		templateVariable["address"] = nonprofit.Address.Street + ", " + nonprofit.Address.City + ", " + nonprofit.Address.State + ", " + nonprofit.Address.Country
		templateVariables[nonprofit.Email] = templateVariable
	}

	emails := FromBulkSendEmailsDTO(bulkSendRequestDTO, templateVariables)

	err = h.emailStore.BulkSendEmails(emails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := BulkSendEmailResponse{Message: "Emails sent successfully"}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// This function maps the individual DTO structure to models.Email,
// making necessary internal changes, and saving them in a new map.
// It is used as a bridge to communicate between the input and expected format in the system.
// Returns the map of email addresses and models.Email.
func FromBulkSendEmailsDTO(bulkSendBody dtos.BulkSendEmailsDTO, templateVariables map[string]map[string]string) map[string]models.Email {
	emails := make(map[string]models.Email, len(bulkSendBody.Recipients))
	for i, emailAddr := range bulkSendBody.Recipients {
		email := bulkSendBody.Recipients[i]
		message := replaceVariables(bulkSendBody.TemplateString, templateVariables[email])

		emails[emailAddr] = models.Email{
			ID:      uuid.New(),
			Subject: bulkSendBody.Subject,
			Body:    message,
		}
	}

	return emails
}

// replaceVariables replaces the template variables with the provided values
/**
  Summary:
    This function accepts the email template and a map of variable values,
    processes and replaces the template variables with the provided
    values and returns the reformed email template.

  Parameters:
    template - Template string containing the raw template
    variables - Map containing the key-value pairs of variables to replace

  Returns:
    The reformed template string.
*/
func replaceVariables(template string, variables map[string]string) string {
	var buf bytes.Buffer
	for {
		start := bytes.IndexByte([]byte(template), '{')
		if start == -1 {
			break
		}

		end := bytes.IndexByte([]byte(template[start+1:]), '}')
		if end == -1 {
			break
		}
		end += start + 1

		buf.Write([]byte(template[:start]))
		key := template[start+1 : end]
		value, ok := variables[key]
		if ok {
			buf.WriteString(value)
		}

		template = template[end+1:]
	}
	buf.WriteString(template)

	return buf.String()
}

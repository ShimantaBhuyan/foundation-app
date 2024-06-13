package handlers

import (
	"encoding/json"
	"net/http"

	"foundation-app/internal/dtos"
	models "foundation-app/internal/models"

	"github.com/google/uuid"
)

type CreateNonprofitResponse struct {
	ID uuid.UUID `json:"id"`
}

// CreateNonprofit creates a new nonprofit
/**
  Summary:
    This function handles the request to create a new nonprofit in the system. It receives a JSON payload
    containing the required information for the nonprofit, validates it, and saves it in a nonprofit store.

  Returns:
    The newly created nonprofit's ID is sent back as a JSON response.
*/
func (h *APIHandlers) CreateNonprofit(w http.ResponseWriter, r *http.Request) {
	var nonprofitDTO dtos.CreateNonprofitDTO
	err := json.NewDecoder(r.Body).Decode(&nonprofitDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nonprofit := FromCreateNonprofitDTO(nonprofitDTO)

	_, err = h.nonprofitStore.CreateNonprofit(nonprofit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := CreateNonprofitResponse{
		ID: uuid.New(),
	}

	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// This function maps the CreateNonprofitDTO struct to the models.NonProfit
// struct. It takes care of translating between them and extracting the
// necessary information for the models.NonProfit struct.
func FromCreateNonprofitDTO(dto dtos.CreateNonprofitDTO) models.NonProfit {
	addressModel := models.NonprofitAddress{
		Street:  dto.Street,
		City:    dto.City,
		State:   dto.State,
		Zipcode: dto.Zipcode,
	}
	return models.NonProfit{
		Name:    dto.Name,
		Email:   dto.Email,
		Address: addressModel,
	}
}

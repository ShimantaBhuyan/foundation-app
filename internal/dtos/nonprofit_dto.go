// internal/api/dtos/nonprofit_dto.go
package dtos

// CreateNonprofitDTO data transfer object
/**
  Summary:
    The CreateNonprofitDTO struct contains the information required for creating a
    Nonprofit organization.

  Fields:
    Name - The name of the Nonprofit organization
    Email - The email address of the Nonprofit organization
    Street - The street address of the Nonprofit organization
    City - The city where the Nonprofit organization is located
    State - The state where the Nonprofit organization is located
    Zipcode - The zip code of the location of the Nonprofit organization
    Country - The country where the Nonprofit organization is located
*/
type CreateNonprofitDTO struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Street  string `json:"street"`
	City    string `json:"city" `
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	Country string `json:"country"`
}

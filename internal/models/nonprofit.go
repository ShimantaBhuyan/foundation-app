package model

// NonProfit model represents the details of a Non-Profit Organization in our system
type NonProfit struct {
	Email   string
	Name    string
	Address NonprofitAddress
}

// NonprofitAddress struct holds the address information of a Non-Profit in our system
type NonprofitAddress struct {
	Street  string
	City    string
	State   string
	Zipcode string
	Country string
}

// Summary:
//   NonProfit struct represents the overall information about a Non-Profit
//   organization including the organization's name, email, and full mailing
//   address.
//
// Fields:
//	- Email (string) : The contact email address associated with the Non-Profit.
//	- Name (string) : The official name of the Non-Profit organization.
//	- Address (NonprofitAddress) : An embedded struct representing the
//   mailing address of the Non-Profit organization.
//
// NonprofitAddress struct holds the street address information of a Non-Profit
// organization, including the city, state, zipcode, and country of the address
//
// Fields:
//	- Street (string): The street address of the Non-Profit organization.
//	- City (string): The name of the city where the Non-Profit organization
//	  is located.
//	- State (string): The state or province where the Non-Profit organization
//	  is located.
//	- Zipcode (string): The postal or zip code of the Non-Profit organization.
//	- Country (string): The country where the Non-Profit organization is located.
//

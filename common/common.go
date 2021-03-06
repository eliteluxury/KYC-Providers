package common

// User data provided to KYC in order to check an individual.
type UserData struct {

	// Id
	Id string `json:"id"`

	// First name
	FirstName string `json:"firstName"`

	// Last name
	LastName string `json:"lastName"`

	// Address
	Address string `json:"address"`

	// City
	City string `json:"city"`

	// State
	State string `json:"state"`

	// Zip code
	Zip string `json:"zip"`

	// etc.

}

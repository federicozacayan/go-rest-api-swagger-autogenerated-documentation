package people

type Address struct {
	// City name
	City string `json:"city,omitempty"`
	// State name
	State string `json:"state,omitempty"`
}

// swagger:model Person
type Person struct {
	// Identifier
	// required: true
	Id string `json:"id"`
	// First name
	// required: true
	FirstName string `json:"firstname"`
	// Last name
	// required: true
	LastName string `json:"lastname"`
	// A JSON that represents the Addres
	Addres *Address `json:"addres,omitempty"`
}

var people []Person

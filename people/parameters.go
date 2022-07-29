package people

// swagger:parameters CreatePerson DeletePerson GetPerson
type peopleIDParameterWraper struct {
	// The id of the person to delete
	// in: path
	// required: true
	Id int `json:"id"`
}

// swagger:parameters CreatePerson
type personWraper struct {
	// The id of the person to delete
	// in: query
	// required: true
	person1 *Person
}

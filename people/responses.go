package people

// OK
// swagger:response personResponse
type personResponse struct {
	// in: body
	person *Person
}

// OK
// swagger:response peopleResponse
type peopleResponse struct {
	// in: body
	Body []Person
}

// Created
// swagger:response personCreatedResponse
type personCreatedResponse struct {
	// in: body
	person *Person
}

// Accepted
// swagger:response acceptedResponse
type acceptedResponse struct {
}

// Internal Server Error
// swagger:response internalServerError
type InternalServerError struct {
}

// Unauthorized
// swagger:response unauthorizedError
type unauthorizedError struct {
}

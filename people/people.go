// Package classification People API.
//
// Documentation for **Person Resftull API**.
//
// This is a tiny project that pretends to provide an easy approach to getting an documented RestFul API following **openAPI Initiative**.
//
// More information about **Person API** can be found here: https://www.openapis.org/about
//
// ### Usage
//
// Scroll doown to see the diferent endpoints or search them in the top left searching box of this website.
//
//
//
// Schemes: http
// BasePath: /
// Version: 0.0.1
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Security:
// - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: Access
//          in: header
//
// swagger:meta
package people

func Init() {
	populateDatabase()
	initRouter()
}

package people

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

func getDocumentationHandler() http.Handler {
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	return middleware.Redoc(options, nil)
}

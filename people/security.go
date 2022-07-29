package people

import (
	"net/http"
)

var api_key = "1234"

func validateKey(w http.ResponseWriter, r *http.Request) bool {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] == api_key {
			return true
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	return false
}

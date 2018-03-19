package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/palantir/stacktrace"
)

// ResponseJSON .
func ResponseJSON(w http.ResponseWriter, statusCode int, value interface{}) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return stacktrace.Propagate(err, "cannot encode json")
	}
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(encoded))
	return nil
}

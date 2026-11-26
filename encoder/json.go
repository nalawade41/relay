package encoder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Encode writes v as a JSON response with the given status code.
func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// Decode reads the JSON request body into a value of type T.
func Decode[T any](r *http.Request) (v T, err error) {
	if err = json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode josn: %w", err)
	}
	return v, nil
}

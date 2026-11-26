package handlers

import (
	"net/http"

	"github.com/nalawade41/relay/encoder"
)

// type errorResponse struct {
// 	Error string `json:"error"`
// }

// HandleHealthz needs no dependencies, so it takes none.
func HandleHealthz() http.Handler {
	type response struct {
		Status string `json:"status"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = encoder.Encode(w, r, http.StatusOK, response{Status: "ok"})
	})
}

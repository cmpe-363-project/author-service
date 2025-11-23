package routes

import (
	restapiutils "author-service/internal/restapi/utils"
	"net/http"
)

// HandleGetVersion
// /api/version
func HandleGetVersion(version string) http.HandlerFunc {
	type Response struct {
		Version string `json:"version"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		restapiutils.WriteJSONResponse(w, http.StatusOK, Response{Version: version})
	}
}

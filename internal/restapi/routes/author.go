package routes

import (
	"author-service/internal/repository"
	restapiutils "author-service/internal/restapi/utils"
	"author-service/pkg/logger"
	"net/http"
	"strconv"
	"strings"
)

// HandleGetAuthors
// /api/authors/by-id
func HandleGetAuthors(logger logger.Logger, repo repository.Repository) http.HandlerFunc {
	type ResponseItem struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type Response struct {
		Items []ResponseItem `json:"items"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ids := r.URL.Query().Get("id")

		idsIntList := []int{}
		for _, idStr := range strings.Split(ids, ",") {
			idInt, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid author ID: "+idStr, http.StatusBadRequest)
				return
			}
			idsIntList = append(idsIntList, idInt)
		}

		authors, err := repo.GetAuthorsByIDs(idsIntList)
		if err != nil {
			logger.ErrorWithCtx(r.Context(), "GetAuthorsByIDs query failed", "error", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respItems := []ResponseItem{}
		for _, author := range authors {
			respItems = append(respItems, ResponseItem{
				ID:   author.ID,
				Name: author.Name,
			})
		}

		resp := Response{
			Items: respItems,
		}
		restapiutils.WriteJSONResponse(w, http.StatusOK, resp)
	}
}

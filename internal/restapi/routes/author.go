package routes

import (
	"author-service/internal/repository"
	restapiutils "author-service/internal/restapi/utils"
	"author-service/pkg/logger"
	"errors"
	"net/http"
	"strconv"
)

// HandleGetAuthor
// /api/authors/{id}
func HandleGetAuthor(logger logger.Logger, repo repository.MysqlRepository) http.HandlerFunc {
	type Response struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		author, err := repo.GetAuthorByID(idInt)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			logger.ErrorWithCtx(r.Context(), "GetAuthorByID query failed", "error", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := Response{
			ID:   author.ID,
			Name: author.Name,
		}
		restapiutils.WriteJSONResponse(w, http.StatusOK, resp)
	}
}

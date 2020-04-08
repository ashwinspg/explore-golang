package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ashwinspg/explore-golang/db"
	"github.com/ashwinspg/explore-golang/services"
	"github.com/ashwinspg/explore-golang/utils"
)

//GetMovieHandler - get movie details
func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	l := utils.LogEntryWithRef()

	uuid := chi.URLParam(r, "id")

	movieDTO, err := services.NewMovie(db.GetDB()).GetMovie(uuid)

	if err != nil {
		l.WithError(err).Error("Failed to get movie information")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get movie information"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movieDTO.Info)
}

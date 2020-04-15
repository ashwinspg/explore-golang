package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ashwinspg/explore-golang/db"
	"github.com/ashwinspg/explore-golang/services"
	"github.com/ashwinspg/explore-golang/utils"

	"github.com/go-chi/chi"
)

func setMovieRoutes(router chi.Router) {
	router.Get("/movies/{id}", GetMovieHandler)
}

//GetMovieHandler - get movie details
func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	l := utils.LogEntryWithRef()
	uuid := chi.URLParam(r, "id")

	movieDTO, err := services.NewMovie(db.GetDB(), l).GetMovie(uuid)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movieDTO.Info)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed while processing movie information"))
	}
}

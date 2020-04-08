package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RealImage/moviebuff-sdk-go"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/ashwinspg/explore-golang/config"
)

//GetMovieHandler - get movie details
func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	UUID := chi.URLParam(r, "id")

	moviebuffObj := moviebuff.New(moviebuff.Config{
		HostURL:     config.MOVIEBUFF_URL,
		StaticToken: config.MOVIEBUFF_TOKEN,
	})

	movieDetail, err := moviebuffObj.GetMovie(UUID)
	if err != nil {
		logrus.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieDetail)
}

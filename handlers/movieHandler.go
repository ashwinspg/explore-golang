package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/RealImage/moviebuff-sdk-go"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/daos"
	"github.com/ashwinspg/explore-golang/db"
	"github.com/ashwinspg/explore-golang/dtos"
	"github.com/ashwinspg/explore-golang/utils"
)

//GetMovieHandler - get movie details
func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	l := utils.LogEntryWithRef()

	uuid := chi.URLParam(r, "id")

	dbObj, err := db.GetPostgresDB()
	if err != nil {
		l.WithError(err).Fatal("Failed to get DB connection")
	}

	movieInfoDAO := daos.NewMovieInfo(dbObj)
	movieInfoDTO, err := getMovie(uuid, movieInfoDAO)

	if err != nil {
		l.WithError(err).Error("Failed to get movie information")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get movie information"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movieInfoDTO.Info)
}

func getMovie(uuid string, movieInfoDAO *daos.MovieInfo) (dtos.Movie, error) {
	movieInfoDTO, err := movieInfoDAO.FindByID(uuid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			movieInfoDTO, err = getMovieFromMovieBuff(uuid)
			if err != nil {
				return dtos.Movie{}, err
			}

			movieInfoDAO.Save(movieInfoDTO)
		default:
			return dtos.Movie{}, err
		}
	} else {
		logrus.Info("Fetched Movie Information from local DB")
	}

	return movieInfoDTO, nil
}

func getMovieFromMovieBuff(uuid string) (dtos.Movie, error) {
	moviebuffObj := moviebuff.New(moviebuff.Config{
		HostURL:     config.MOVIEBUFF_URL,
		StaticToken: config.MOVIEBUFF_TOKEN,
	})

	movieDetail, err := moviebuffObj.GetMovie(uuid)

	if err != nil {
		return dtos.Movie{}, err
	}

	var movieInfoDTO dtos.Movie
	movieInfoDTO.UUID = movieDetail.UUID
	movieInfoDTO.Info, err = utils.TransformToPropertyMap(movieDetail)

	if err != nil {
		return dtos.Movie{}, err
	}

	logrus.Info("Fetched Movie Information from MovieBuff SDK")

	return movieInfoDTO, nil
}

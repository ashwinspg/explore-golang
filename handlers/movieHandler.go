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
	uuid := chi.URLParam(r, "id")

	dbObj, _ := db.GetPostgresDB()

	movieInfoDAO := daos.NewMovieInfo(dbObj)
	movieInfoDTO, _ := getMovie(uuid, movieInfoDAO)

	// movieInfoDTO, _ := getMovieFromMovieBuff(uuid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieInfoDTO.Info)
}

func getMovie(uuid string, movieInfoDAO *daos.MovieInfo) (dtos.MovieInfo, error) {
	movieInfoDTO, err := movieInfoDAO.FindByID(uuid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			movieInfoDTO, _ = getMovieFromMovieBuff(uuid)
			movieInfoDAO.Save(movieInfoDTO)
		default:
			return dtos.MovieInfo{}, err
		}
	} else {
		logrus.Info("Fetching Movie Information from local DB")
	}

	return movieInfoDTO, nil
}

func getMovieFromMovieBuff(uuid string) (dtos.MovieInfo, error) {
	logrus.Info("Fetching Movie Information from MovieBuff SDK")
	moviebuffObj := moviebuff.New(moviebuff.Config{
		HostURL:     config.MOVIEBUFF_URL,
		StaticToken: config.MOVIEBUFF_TOKEN,
	})

	movieDetail, err := moviebuffObj.GetMovie(uuid)

	if err != nil {
		return dtos.MovieInfo{}, err
	}

	var movieInfoDTO dtos.MovieInfo
	movieInfoDTO.Movie_UUID = movieDetail.UUID
	movieInfoDTO.Info = make(utils.PropertyMap)
	j, _ := json.Marshal(movieDetail)
	json.Unmarshal(j, &movieInfoDTO.Info)

	return movieInfoDTO, nil
}

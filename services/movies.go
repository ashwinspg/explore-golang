package services

import (
	"database/sql"
	"github.com/RealImage/moviebuff-sdk-go"
	"github.com/sirupsen/logrus"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/daos"
	"github.com/ashwinspg/explore-golang/dtos"
	"github.com/ashwinspg/explore-golang/utils"
)

//Movie - service
type Movie struct {
	dao *daos.Movie
}

//NewMovie - instance creation
func NewMovie(db *sql.DB) *Movie {
	return &Movie{
		dao: daos.NewMovie(db),
	}
}

func (m *Movie) GetMovie(uuid string) (dtos.Movie, error) {
	movieDTO, err := m.dao.FindByID(uuid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			movieDTO, err = m.getMovieFromMovieBuff(uuid)
			if err != nil {
				return dtos.Movie{}, err
			}

			m.dao.Save(movieDTO)
		default:
			return dtos.Movie{}, err
		}
	} else {
		logrus.Info("Fetched Movie Information from local DB")
	}

	return movieDTO, nil
}

func (m *Movie) getMovieFromMovieBuff(uuid string) (dtos.Movie, error) {
	moviebuffObj := moviebuff.New(moviebuff.Config{
		HostURL:     config.MOVIEBUFF_URL,
		StaticToken: config.MOVIEBUFF_TOKEN,
	})

	movieDetail, err := moviebuffObj.GetMovie(uuid)

	if err != nil {
		return dtos.Movie{}, err
	}

	var movieDTO dtos.Movie
	movieDTO.UUID = movieDetail.UUID
	movieDTO.Info, err = utils.TransformToPropertyMap(movieDetail)

	if err != nil {
		return dtos.Movie{}, err
	}

	logrus.Info("Fetched Movie Information from MovieBuff SDK")

	return movieDTO, nil
}

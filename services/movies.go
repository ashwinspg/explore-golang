package services

import (
	"database/sql"

	"github.com/RealImage/moviebuff-sdk-go"
	"github.com/sirupsen/logrus"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/constants"
	"github.com/ashwinspg/explore-golang/daos"
	"github.com/ashwinspg/explore-golang/dtos"
	"github.com/ashwinspg/explore-golang/utils"
)

//Movie - service
type Movie struct {
	dao *daos.Movie
	l   *logrus.Entry
}

//NewMovie - instance creation
func NewMovie(db *sql.DB, l *logrus.Entry) *Movie {
	return &Movie{
		dao: daos.NewMovie(db),
		l:   l,
	}
}

//GetMovie - get movie based on uuid
func (m *Movie) GetMovie(uuid string) (movieDTO dtos.Movie, err error) {
	movieDTO, err = m.dao.FindByID(uuid)
	switch err {
	case nil:
		return
	case constants.ErrMovieNotFoundInDB:
		movieDTO, err = m.getMovieFromMovieBuff(uuid)
		if err != nil {
			m.l.WithError(err).Errorln("Failed to get Movie Information from MovieBuff")
			return
		}

		err = m.dao.Save(movieDTO)
		if err != nil {
			m.l.WithError(err).Errorln("Failed to save Movie Information from MovieBuff in Database")
		}
		return
	default:
		m.l.WithError(err).Errorln("Failed to get Movie Information from Database")
		return dtos.Movie{}, err
	}
}

func (m *Movie) getMovieFromMovieBuff(uuid string) (movieDTO dtos.Movie, err error) {
	moviebuffObj := moviebuff.New(moviebuff.Config{
		HostURL:     config.MOVIEBUFF_URL,
		StaticToken: config.MOVIEBUFF_TOKEN,
	})
	movieDetail, err := moviebuffObj.GetMovie(uuid)
	if err != nil {
		return dtos.Movie{}, err
	}

	movieDTO.UUID = movieDetail.UUID
	movieDTO.Info, err = utils.TransformToPropertyMap(movieDetail)
	if err != nil {
		return dtos.Movie{}, err
	}

	return
}

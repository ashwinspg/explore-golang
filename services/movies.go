package services

import (
	"database/sql"

	"github.com/RealImage/moviebuff-sdk-go"
	"github.com/sirupsen/logrus"

	"github.com/ashwinspg/explore-golang/config"
	"github.com/ashwinspg/explore-golang/daos"
	"github.com/ashwinspg/explore-golang/dtos"
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
	case daos.ErrMovieNotFound:
		movieDTO, err = m.getMovieFromMovieBuff(uuid)
		if err != nil {
			m.l.WithError(err).Errorln("Failed to get Movie Information from MovieBuff")
			return
		}

		err = m.dao.Save(movieDTO)
		if err != nil {
			m.l.WithError(err).Errorln("Failed to save Movie Information from MovieBuff")
		}
		return
	default:
		m.l.WithError(err).Errorln("Failed to get Movie Information from Database")
		return
	}
}

func (m *Movie) getMovieFromMovieBuff(uuid string) (movieDTO dtos.Movie, err error) {
	moviebuffObj := moviebuff.New(moviebuff.Config{
		HostURL:     config.MOVIEBUFF_URL,
		StaticToken: config.MOVIEBUFF_TOKEN,
	})
	movieDetail, err := moviebuffObj.GetMovie(uuid)
	if err != nil {
		return
	}

	movieDTO.UUID = movieDetail.UUID
	movieDTO.Info, err = dtos.TransformToMovieInfo(movieDetail)

	return
}

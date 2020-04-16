package services

import (
	"database/sql"
	"errors"

	"github.com/ashwinspg/explore-golang/daos"
	"github.com/ashwinspg/explore-golang/dtos"
	"github.com/ashwinspg/explore-golang/utils"

	"github.com/sirupsen/logrus"
)

var (
	ErrInvalidMovieUUID = errors.New("Invalid value for Movie UUID")
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
	if !utils.IsValidUUID(uuid) {
		err = ErrInvalidMovieUUID
		return
	}

	movieDTO, err = m.dao.FindByID(uuid)
	if err == nil {
		return
	} else if err != daos.ErrMovieNotFound {
		m.l.WithError(err).Errorln("Failed to get Movie Information from Database")
		return
	}

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
}

func (m *Movie) getMovieFromMovieBuff(uuid string) (movieDTO dtos.Movie, err error) {
	moviebuffObj, err := GetMovieBuff()
	if err != nil {
		return
	}

	movieDetail, err := moviebuffObj.GetMovie(uuid)
	if err != nil {
		return
	}

	movieDTO.UUID = movieDetail.UUID
	movieDTO.Info, err = dtos.TransformToMovieInfo(movieDetail)

	return
}

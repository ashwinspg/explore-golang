package daos

import (
	"database/sql"

	"github.com/ashwinspg/explore-golang/dtos"
)

//MovieInfo - DAO
type MovieInfo struct {
	db *sql.DB
}

//NewMovieInfo - instance creation
func NewMovieInfo(db *sql.DB) *MovieInfo {
	return &MovieInfo{db}
}

//Save - add new MovieInfo data
func (movieInfoDAO *MovieInfo) Save(movieInfoDTO dtos.MovieInfo) error {
	query := `INSERT INTO "` + dtos.MovieInfoTable + `"("movie_uuid", "info") VALUES($1, $2)`

	statement, err := movieInfoDAO.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(movieInfoDTO.Movie_UUID, movieInfoDTO.Info)

	if err != nil {
		return err
	}

	return nil
}

//FindByID - find movie_info through id
func (movieInfoDAO *MovieInfo) FindByID(uuid string) (dtos.MovieInfo, error) {
	query := `SELECT * FROM ` + dtos.MovieInfoTable + ` WHERE "movie_uuid" = $1`

	var movieInfoDTO dtos.MovieInfo

	statement, err := movieInfoDAO.db.Prepare(query)

	if err != nil {
		return dtos.MovieInfo{}, err
	}

	defer statement.Close()

	err = statement.QueryRow(uuid).Scan(&movieInfoDTO.Movie_UUID, &movieInfoDTO.Info)

	if err != nil {
		return dtos.MovieInfo{}, err
	}

	return movieInfoDTO, nil
}

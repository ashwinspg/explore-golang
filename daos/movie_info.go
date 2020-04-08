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
func (movieInfoDAO *MovieInfo) Save(movieInfoDTO dtos.Movie) error {
	query := `INSERT INTO "` + dtos.MovieTable + `"("uuid", "info") VALUES($1, $2)`

	statement, err := movieInfoDAO.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(movieInfoDTO.UUID, movieInfoDTO.Info)

	if err != nil {
		return err
	}

	return nil
}

//FindByID - find movie_info through id
func (movieInfoDAO *MovieInfo) FindByID(uuid string) (dtos.Movie, error) {
	query := `SELECT * FROM ` + dtos.MovieTable + ` WHERE "uuid" = $1`

	var movieInfoDTO dtos.Movie

	statement, err := movieInfoDAO.db.Prepare(query)

	if err != nil {
		return dtos.Movie{}, err
	}

	defer statement.Close()

	err = statement.QueryRow(uuid).Scan(&movieInfoDTO.UUID, &movieInfoDTO.Info)

	if err != nil {
		return dtos.Movie{}, err
	}

	return movieInfoDTO, nil
}

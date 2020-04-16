package daos

import (
	"database/sql"
	"errors"

	"github.com/ashwinspg/explore-golang/dtos"
)

var (
	ErrMovieNotFound = errors.New("Movie Information for given UUID is not found")
)

//Movie - DAO
type Movie struct {
	db *sql.DB
}

//NewMovie - instance creation
func NewMovie(db *sql.DB) *Movie {
	return &Movie{db}
}

//Save - add new Movie data
func (movieDAO *Movie) Save(movieDTO dtos.Movie) (err error) {
	query := `INSERT INTO "` + dtos.MovieTable + `"("uuid", "info") VALUES($1, $2);`
	statement, err := movieDAO.db.Prepare(query)
	if err != nil {
		return
	}

	defer statement.Close()

	_, err = statement.Exec(movieDTO.UUID, movieDTO.Info)
	return
}

//FindByID - find movie through uuid
func (movieDAO *Movie) FindByID(uuid string) (movieDTO dtos.Movie, err error) {
	query := `SELECT * FROM ` + dtos.MovieTable + ` WHERE "uuid" = $1`
	statement, err := movieDAO.db.Prepare(query)
	if err != nil {
		return
	}

	defer statement.Close()

	err = statement.QueryRow(uuid).Scan(&movieDTO.UUID, &movieDTO.Info)
	if err == sql.ErrNoRows {
		err = ErrMovieNotFound
	}

	return
}

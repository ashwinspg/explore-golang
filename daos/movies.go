package daos

import (
	"database/sql"

	"github.com/ashwinspg/explore-golang/dtos"
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
func (movieDAO *Movie) Save(movieDTO dtos.Movie) error {
	query := `INSERT INTO "` + dtos.MovieTable + `"("uuid", "info") VALUES($1, $2);`
	statement, err := movieDAO.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(movieDTO.UUID, movieDTO.Info)

	return err
}

//FindByID - find movie through uuid
func (movieDAO *Movie) FindByID(uuid string) (dtos.Movie, error) {
	var movieDTO dtos.Movie

	query := `SELECT * FROM ` + dtos.MovieTable + ` WHERE "uuid" = '` + uuid + `';`
	statement, err := movieDAO.db.Prepare(query)

	if err != nil {
		return dtos.Movie{}, err
	}

	defer statement.Close()

	err = statement.QueryRow().Scan(&movieDTO.UUID, &movieDTO.Info)

	return movieDTO, err
}

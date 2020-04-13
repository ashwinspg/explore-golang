package dtos

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/RealImage/moviebuff-sdk-go"
)

//MovieInfo - wrapper for moviebuff.Movie
type MovieInfo moviebuff.Movie

//Movie - DTO
type Movie struct {
	UUID string    `json:"uuid"`
	Info MovieInfo `json:"info"`
}

const MovieTable = "movies"

//Value - simply returns the JSON-encoded representation of the struct.
func (m MovieInfo) Value() (driver.Value, error) {
	return json.Marshal(m)
}

//Scan - simply decodes a JSON-encoded value into the struct fields.
func (m *MovieInfo) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		fmt.Println(ok)
		return errors.New("Type assertion .([]byte) failed.")
	}

	return json.Unmarshal(source, &m)
}

//TransformToMovieInfo - Value v transforms to MovieInfo
func TransformToMovieInfo(v interface{}) (des MovieInfo, err error) {
	j, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = json.Unmarshal(j, &des)
	return
}

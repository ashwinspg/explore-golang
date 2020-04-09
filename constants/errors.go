package constants

import "errors"

var (
	ErrMovieNotFoundInDB = errors.New("Movie Information for given UUID is not found in Database")
)

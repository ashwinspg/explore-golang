package dtos

import (
	"github.com/ashwinspg/explore-golang/utils"
)

//MovieInfo - DTO
type MovieInfo struct {
	Movie_UUID string            `json:"movie_uuid"`
	Info       utils.PropertyMap `json:"info"`
}

const MovieInfoTable string = "movie_info"

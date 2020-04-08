package dtos

import (
	"github.com/ashwinspg/explore-golang/utils"
)

//Movie - DTO
type Movie struct {
	UUID string            `json:"uuid"`
	Info utils.PropertyMap `json:"info"`
}

const MovieTable = "movies"

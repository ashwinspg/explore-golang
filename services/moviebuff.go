package services

import (
	"fmt"

	"github.com/ashwinspg/explore-golang/config"

	mbSDK "github.com/RealImage/moviebuff-sdk-go"
)

//GetMovieBuff - get moviebuff object
func GetMovieBuff() (mbSDK.Moviebuff, error) {
	switch config.ENV {
	case "dev":
		return mbSDK.New(mbSDK.Config{
			HostURL:     config.MOVIEBUFF_URL,
			StaticToken: config.MOVIEBUFF_TOKEN,
		}), nil
	case "test":
		return NewMockMoviebuff(), nil
	}

	return nil, fmt.Errorf("%s: Not a valid environment", config.ENV)
}

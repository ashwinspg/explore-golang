package services

import (
	"testing"

	"github.com/ashwinspg/explore-golang/test"

	"github.com/stretchr/testify/suite"
)

type MovieTestSuite struct {
	suite.Suite
	env     test.Env
	service *Movie
}

func (t *MovieTestSuite) SetupTest() {
	t.T()
	t.env = test.SetupTestEnv()
	t.service = NewMovie(t.env.DBConn, t.env.L)
}

func (t *MovieTestSuite) TestGetMovie() {
	movieDTO, err := t.service.GetMovie("1d3095da-3243-4a34-a7d6-2cb570446ffe")
	t.NotEmpty(movieDTO)
	t.NoError(err)
}

func TestMovieSuite(t *testing.T) {
	suite.Run(t, new(MovieTestSuite))
}

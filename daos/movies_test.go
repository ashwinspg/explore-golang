package daos

import (
	"testing"

	"github.com/ashwinspg/explore-golang/dtos"
	"github.com/ashwinspg/explore-golang/test"

	"github.com/stretchr/testify/suite"
)

type MovieTestSuite struct {
	suite.Suite
	env test.Env
	dao *Movie
}

func (t *MovieTestSuite) SetupTest() {
	t.env = test.SetupTestEnv()
	t.dao = NewMovie(t.env.DBConn)
}

func (t *MovieTestSuite) TestSave() {
	movieDTO := dtos.Movie{
		UUID: "976a6bc7-c5e6-4544-bcad-127db7de2e87",
		Info: dtos.MovieInfo{URL: "fast-and-furious-presents-hobbs-and-shaw-2019-hindi"},
	}
	err := t.dao.Save(movieDTO)
	t.NoError(err)

	result, err := t.dao.FindByID(movieDTO.UUID)
	t.NotNil(result)
	t.NoError(err, "Error finding movie by given UUID")

	result, err = t.dao.FindByID("976a6bc7-c5e6-4544-bcad-127db7de2111")
	t.Empty(result)
	t.Error(err)
}

func TestMovieSuite(t *testing.T) {
	suite.Run(t, new(MovieTestSuite))
}

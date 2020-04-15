package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ashwinspg/explore-golang/test"
)

func TestGetMovie(t *testing.T) {
	test.SetupTestEnv()
	request := httptest.NewRequest(http.MethodGet, "/movies/0815d6a2-67ba-4487-a529-142f28f4d21c", nil)
	w := httptest.NewRecorder()
	GetMovieHandler(w, request)
	response := w.Result()
	fmt.Println(response)
	assert.New(t).Equal(http.StatusOK, response.StatusCode)
}

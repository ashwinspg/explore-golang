package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashwinspg/explore-golang/test"

	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	test.SetupTestEnv()
	request := httptest.NewRequest(http.MethodGet, "/movies/0815d6a2-67ba-4487-a529-142f28f4d21c", nil)
	w := httptest.NewRecorder()
	GetRouter().ServeHTTP(w, request)
	response := w.Result()
	assert.New(t).Equal(http.StatusOK, response.StatusCode)
}

func TestGetMovieBadRequest(t *testing.T) {
	test.SetupTestEnv()
	request := httptest.NewRequest(http.MethodGet, "/movies/0815d6a2-67ba-4487-a529-142f28f4d21csfsfdsd", nil)
	w := httptest.NewRecorder()
	GetRouter().ServeHTTP(w, request)
	response := w.Result()
	assert.New(t).Equal(http.StatusBadRequest, response.StatusCode)
}

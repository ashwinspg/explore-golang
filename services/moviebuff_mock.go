package services

import (
	mbSDK "github.com/RealImage/moviebuff-sdk-go"
)

type mockMoviebuff struct{}

//NewMockMoviebuff - get moviebuff mock object
func NewMockMoviebuff() *mockMoviebuff {
	return new(mockMoviebuff)
}

//GetCertifications - mock implementation
func (mb *mockMoviebuff) GetCertifications(countryID string) (cert []mbSDK.Certification, err error) {
	return
}

//GetMovie - mock implementation
func (mb *mockMoviebuff) GetMovie(id string) (m *mbSDK.Movie, err error) {
	m = &mbSDK.Movie{
		UUID:     id,
		Name:     "Test Movie",
		Language: "English",
	}
	return
}

//GetPerson - mock implementation
func (mb *mockMoviebuff) GetPerson(id string) (p *mbSDK.Person, err error) { return }

//GetEntity - mock implementation
func (mb *mockMoviebuff) GetEntity(id string) (e *mbSDK.Entity, err error) { return }

//GetResources - mock implementation
func (mb *mockMoviebuff) GetResources(resourceType mbSDK.ResourceType, limit, page int) (r *mbSDK.Resources, err error) {
	return
}

//GetHolidayCalendar - mock implementation
func (mb *mockMoviebuff) GetHolidayCalendar(countryID string) (c *mbSDK.Calendar, err error) { return }

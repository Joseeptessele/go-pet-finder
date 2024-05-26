package main

import (
	"go-pet-finder/configuration"
	"go-pet-finder/models"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	testBackend := &TestBackend{}
	testAdapter := &RemoteService{Remote: testBackend}
	testApp = application{
		App: configuration.New(nil),
		CatService:  testAdapter,
	}

	os.Exit(m.Run())
}

type TestBackend struct{}

func (tb *TestBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
	breeds := []*models.CatBreed{
		&models.CatBreed{ID: 1, Breed: "tomcat", Details: "some detail"},
	}
	return breeds, nil
}

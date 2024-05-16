package pets

import "go-pet-finder/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet {
		Species: species,
		Breed: "",
		MinWeight: 0,
		MaxWeight: 0,
		Description: "no description",
		LifeSpan: 0,
	}

	return &pet
}

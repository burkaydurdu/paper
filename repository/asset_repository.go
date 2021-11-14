package repository

import "paper.com/paper/model"

type AssetRepository interface {
	GetAllColors() ([]model.Color, error)
	GetAllImages() ([]model.Image, error)
}

type assRepository struct {
	prefix   string
	isRandom bool
}

func NewAssetRepository() AssetRepository {
	return &assRepository{
		prefix:   "#",
		isRandom: false,
	}
}

func (c *assRepository) GetAllColors() ([]model.Color, error) {
	colorOne := model.Color{}

	colorOne.ID = 1
	colorOne.Name = "Black"
	colorOne.Random = c.isRandom

	colorTwo := model.Color{ID: 2, Name: "White", Random: c.isRandom}

	return []model.Color{colorOne, colorTwo}, nil
}

func (c *assRepository) GetAllImages() ([]model.Image, error) {
	imageOne := model.Image{}

	imageOne.ID = 1
	imageOne.URL = "/imageOne.png"
	imageOne.Random = c.isRandom

	imageTwo := model.Image{ID: 2, URL: "/imageTwo.jpg", Random: c.isRandom}

	return []model.Image{imageOne, imageTwo}, nil
}

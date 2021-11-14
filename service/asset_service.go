package service

import (
	"paper.com/paper/model"
	"paper.com/paper/repository"
)

type AssetService interface {
	GetAllColors() ([]model.Color, error)
	GetAllImages() ([]model.Image, error)
}

type assService struct {
	Repo repository.AssetRepository
}

func NewAssetService(ar repository.AssetRepository) AssetService {
	return &assService{Repo: ar}
}

func (a *assService) GetAllColors() ([]model.Color, error) {
	// ... Logic ... \\
	return a.Repo.GetAllColors()
}

func (a *assService) GetAllImages() ([]model.Image, error) {
	// ... Logic ... \\
	return a.Repo.GetAllImages()
}

package service

import (
	"galaxy-weather/model"
	"galaxy-weather/repository"
)

type IPlanetService interface {
	Save(p model.Planet) error
	GetAll() (model.Planet, model.Planet, model.Planet)
	GetList() []model.Planet
}

type planetService struct {
	repo repository.IPlanetRepo
}

func newPlanetService() IPlanetService {
	return &planetService{
		repository.NewPlanetRepo(),
	}
}

var PlanetService = newPlanetService()

func (ps planetService) Save(p model.Planet) error {
	return ps.repo.Save(p)
}

func (ps planetService) GetAll() (model.Planet, model.Planet, model.Planet) {
	return ps.repo.GetAll()
}

func (ps planetService) GetList() []model.Planet {
	f, b, v := ps.repo.GetAll()
	return []model.Planet{f, b, v}
}

package repository

import (
	"galaxy-weather/database"
	"galaxy-weather/model"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type IPlanetRepo interface {
	Save(planet model.Planet) error
	GetAll() (model.Planet, model.Planet, model.Planet)
}

type planetRepo struct {
	Driver *gorm.DB
}

func NewPlanetRepo() IPlanetRepo {
	return &planetRepo{
		Driver: database.Client(),
	}
}

func (pr planetRepo) Save(planet model.Planet) error {
	result := pr.Driver.Where(model.Planet{ID: planet.ID}).Attrs(planet).FirstOrCreate(&planet)

	if result.Error != nil {
		logrus.Errorf("[PlanetRepo.Save] (%s) Error: %s", planet.Name, result.Error)
		return result.Error
	}
	return nil
}

func (pr planetRepo) GetAll() (model.Planet, model.Planet, model.Planet) {
	ferengi := model.Planet{
		ID:              1,
		Name:            "Ferengi",
		AngularVelocity: 1,
		Distance:        500,
		Direction:       model.Clockwise,
		Radio:           80,
	}

	betasoide := model.Planet{
		ID:              2,
		Name:            "Betasoide",
		AngularVelocity: 3,
		Distance:        2000,
		Direction:       model.Clockwise,
		Radio:           30,
	}

	vulcano := model.Planet{
		ID:              3,
		Name:            "Vulcano",
		AngularVelocity: 5,
		Distance:        1000,
		Direction:       model.CounterClockwise,
		Radio:           50,
	}

	return ferengi, betasoide, vulcano
}

package repository

import (
	"galaxy-weather/database"
	"galaxy-weather/model"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)

type IWeatherRepo interface {
	GetByDay(day uint) (*model.Weather, error)
	GetAll() (*[]model.Weather, error)
	Save(weather model.Weather) error
	GetByRange(offset uint, limit uint) (*[]model.Weather, error)
}

type weatherRepo struct {
	Driver *gorm.DB
}

func NewWeatherRepo() IWeatherRepo {
	return &weatherRepo{
		Driver: database.Client(),
	}
}

func (r weatherRepo) GetByDay(day uint) (*model.Weather, error) {
	weatherData := new(model.Weather)
	err := r.Driver.Where(model.Weather{Day: day}).First(&weatherData).Error
	if err != nil {
		logrus.Errorf("[WeatherRepo.GetByDay] (%d) Error: %s", day, err.Error())
		return nil, err
	}

	return weatherData, nil
}

func (r weatherRepo) GetAll() (*[]model.Weather, error) {
	var weathers = make([]model.Weather, 0)
	err := r.Driver.Order("day").Find(&weathers).Error
	if err != nil {
		logrus.Errorf("[WeatherRepo.GetAll] Error: %s", err.Error())
		return nil, err
	}
	return &weathers, nil
}

func (r weatherRepo) Save(w model.Weather) error {
	err := r.Driver.
		Where(model.Weather{Day: w.Day}).
		Attrs(w).
		FirstOrCreate(&w).Error

	if err != nil {
		logrus.Errorf("[WeatherRepo.Save] (%d, %s) Error: %s", w.Day, w.WeatherType.String(), err.Error())
		return err
	}

	return nil
}

func (r weatherRepo) GetByRange(offset uint, limit uint) (*[]model.Weather, error) {
	var weathers = make([]model.Weather, 0)
	err := r.Driver.Offset(offset).Limit(limit).Order("day").Find(&weathers).Error
	if err != nil {
		logrus.Errorf("[WeatherRepo.GetByRange] Offset: %d, Limit: %d | Error: %s", offset, limit, err.Error())
		return nil, err
	}

	return &weathers, nil
}

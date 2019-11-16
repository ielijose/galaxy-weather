package service

import (
	"galaxy-weather/model"
	"galaxy-weather/utils"
	"math"
)

const MaxPerimeter = 6262.300354

type IGalaxyService interface {
	PredictWeather(uint) (*model.Weather, error)
	GetDaysPerYear() uint
}

type galaxyService struct {
	Ferengi     model.Planet
	Betasoide   model.Planet
	Vulcano     model.Planet
	Sun         model.Position
	DaysPerYear uint
}

func newGalaxyService() IGalaxyService {
	ferengi, betasoides, vulcanos := PlanetService.GetAll()
	return &galaxyService{
		Ferengi:     ferengi,
		Betasoide:   betasoides,
		Vulcano:     vulcanos,
		Sun:         model.Position{X: 0.0, Y: 0.0},
		DaysPerYear: vulcanos.GetDaysPerYear(),
	}
}

var GalaxyService = newGalaxyService()

func (gs galaxyService) PredictWeather(day uint) (*model.Weather, error) {
	var w = model.Weather{
		Day:         day,
		WeatherType: model.Unknown,
	}

	p1 := gs.Ferengi.GetPointByDay(day)
	p2 := gs.Betasoide.GetPointByDay(day)
	p3 := gs.Vulcano.GetPointByDay(day)

	if gs.areAligned(p1, p2, p3) {
		w.WeatherType = model.OptimalTemperature
	}

	if gs.areAlignedWithTheSun(p1, p2, p3) {
		w.WeatherType = model.Drought
	}

	isInside, rainType := gs.sunInsidePlanets(p1, p2, p3)
	if isInside {
		w.WeatherType = rainType
	}

	err := save(w, p1, p2, p3)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (gs galaxyService) GetDaysPerYear() uint {
	return gs.DaysPerYear
}

func (gs galaxyService) areAligned(p3, p1, p2 model.Position) bool {
	a := p2.X - p1.X
	b := p2.Y - p1.Y
	c := p3.X - p1.X
	d := p3.Y - p1.Y
	abcd := math.Abs(a*d - b*c)

	if (abcd < a*d/15) || (abcd < b*c/15) {
		return true
	}

	return utils.FloatEquals(abcd, 0.0)
}

func (gs galaxyService) areAlignedWithTheSun(p1, p2, p3 model.Position) bool {
	a := p2.X - p1.X
	b := p2.Y - p1.Y
	c := p3.X - p1.X
	d := p3.Y - p1.Y
	e := gs.Sun.X - p1.X
	f := gs.Sun.Y - p1.Y

	abcd := math.Abs(a*d - b*c)
	abef := math.Abs(a*f - b*e)

	return utils.FloatEquals(abcd, 0.0) && utils.FloatEquals(abef, 0.0)
}

func area(p1, p2, p3 model.Position) float64 {
	return math.Abs((p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y)) / 2.0)
}

func (gs galaxyService) sunInsidePlanets(p1, p2, p3 model.Position) (bool, model.Type) {
	a := area(p1, p2, p3)

	a1 := area(gs.Sun, p2, p3)
	a2 := area(p1, gs.Sun, p3)
	a3 := area(p1, p2, gs.Sun)

	if utils.FloatEquals(a, a1+a2+a3) && !utils.FloatEquals(a, 0.0) {
		if isMaxPerimeter(p1, p2, p3) {
			return true, model.HeavyRain
		}
		return true, model.Rain
	}

	return false, model.Unknown
}

func isMaxPerimeter(p1, p2, p3 model.Position) bool {
	perimeter := perimeter(p1, p2, p3)
	return perimeter >= MaxPerimeter
}

func perimeter(p1, p2, p3 model.Position) float64 {
	d1 := p1.DistanceOf(p2)
	d2 := p2.DistanceOf(p3)
	d3 := p3.DistanceOf(p1)

	return d1 + d2 + d3
}

func save(w model.Weather, p1 model.Position, p2 model.Position, p3 model.Position) error {
	err := WeatherService.Save(w)
	if err != nil {
		return err
	}

	err = PositionService.Save(p1)
	if err != nil {
		return err
	}

	err = PositionService.Save(p2)
	if err != nil {
		return err
	}

	err = PositionService.Save(p3)
	if err != nil {
		return err
	}

	return nil
}

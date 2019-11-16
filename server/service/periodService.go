package service

import (
	"galaxy-weather/model"

	"github.com/sirupsen/logrus"
)

type IPeriodService interface {
	GetAll() (*[]model.Period, error)
	GetStats() (interface{}, error)
}

type periodService struct{}

func newPeriodService() IPeriodService {
	return &periodService{}
}

var PeriodService = newPeriodService()

func (ps periodService) GetAll() (*[]model.Period, error) {
	var periods = make([]model.Period, 0)
	weatherList, err := WeatherService.GetAll()
	if err != nil {
		logrus.Errorf("[PeriodService.GetAll] Error getting periods: %s", err.Error())
		return nil, err
	}

	var period *model.Period

	for _, day := range *weatherList {
		day.WeatherType = day.ToType()

		if day.WeatherType == model.Unknown {
			continue
		}

		if period == nil {
			period = fromDay(day)
			continue
		}

		if period.WeatherType == day.WeatherType {
			period.End = day.Day
			continue
		} else {
			if period.WeatherType == model.Rain && day.WeatherType == model.HeavyRain {
				period.Peak = new(uint)
				*period.Peak = day.Day
				continue
			}
			periods = append(periods, *period)

			period = fromDay(day)
		}
	}

	return &periods, nil
}

func fromDay(day model.Weather) *model.Period {
	period := new(model.Period)
	period.Start = day.Day
	period.End = day.Day
	period.WeatherType = day.WeatherType
	period.Weather = day.Weather
	return period
}

func (ps periodService) GetStats() (interface{}, error) {

	stats := model.Stats{
		Drought:  0,
		Rain:     0,
		RainPeaks: []uint{},
		Optimal:  0,
	}

	periods, err := ps.GetAll()
	if err != nil {
		return nil, err
	}

	for _, period := range *periods {
		if period.WeatherType == model.Drought {
			stats.Drought++
		}
		if period.WeatherType == model.OptimalTemperature {
			stats.Optimal++
		}
		if period.WeatherType == model.Rain {
			stats.Rain++
			if period.Peak != nil && *period.Peak > 0 {
				stats.RainPeaks = append(stats.RainPeaks, *period.Peak)
			}
		}
	}
	return stats, nil
}

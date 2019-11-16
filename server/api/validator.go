package api

import "gopkg.in/go-playground/validator.v9"

type GalaxyWeatherValidator struct {
	validator *validator.Validate
}

func (cv *GalaxyWeatherValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type GalaxyWeatherError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

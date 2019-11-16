package samples

import (
	"fmt"
	"galaxy-weather/service"
	"net/http"

	"github.com/labstack/echo"
)

func sampleSeedHandler(c echo.Context) error {
	var response []string
	// Planets
	planets := service.PlanetService.GetList()

	for _, planet := range planets {
		err := service.PlanetService.Save(planet)
		if err == nil {
			response = append(response, fmt.Sprintf("Planet created successfully: %s", planet.Name))
		}
	}

	// Weather
	for day := uint(0); day < 10*service.GalaxyService.GetDaysPerYear(); day++ {
		w, err := service.GalaxyService.PredictWeather(day)
		if err == nil && w != nil {
			response = append(response, fmt.Sprintf("Weather for day %d predicted successfully: %s", w.Day, w.WeatherType.String()))
		}
	}
	return c.JSON(http.StatusOK, response)
}

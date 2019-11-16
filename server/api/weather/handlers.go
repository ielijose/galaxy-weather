package weather

import (
	"errors"
	"galaxy-weather/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func weatherByDayHandler(c echo.Context) error {
	param := c.Param("day")
	if param == "" {
		return errors.New("day is required")
	}

	d, _ := strconv.Atoi(param)
	if d < 0 {
		return errors.New("invalid day")
	}

	day := uint(d)
	response, err := service.WeatherService.GetByDay(day)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

func weatherByYearHandler(c echo.Context) error {
	param := c.Param("year")
	if param == "" {
		return errors.New("year is required")
	}

	y, _ := strconv.Atoi(param)
	if y < 1 {
		return errors.New("invalid year")
	}

	year := uint(y)
	response, err := service.WeatherService.GetByYear(year)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, response)
}

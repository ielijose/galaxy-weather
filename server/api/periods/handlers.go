package periods

import (
	"galaxy-weather/service"
	"net/http"

	"github.com/labstack/echo"
)

func periodsHandlers(c echo.Context) error {
	periods, err := service.PeriodService.GetAll()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, periods)
}


func periodStatsHandlers(c echo.Context) error {
	stats, err := service.PeriodService.GetStats()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, stats)
}
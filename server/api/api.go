package api

import (
	healthCheck "galaxy-weather/api/health-check"
	"galaxy-weather/api/periods"
	"galaxy-weather/api/samples"
	"galaxy-weather/api/weather"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func Init() *echo.Echo {
	server := echo.New()
	server.HTTPErrorHandler = errorHandler
	server.Validator = &GalaxyWeatherValidator{validator: validator.New()}

	// Middleware
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] Request: method=${method}, uri=${uri}, status=${status}\n",
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.Gzip())

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))

	// Routes
	api := server.Group("/api/galaxy-weather")

	// /api/galaxy-weather/health-check
	hcg := api.Group("/health-check")
	healthCheck.Init(hcg)

	// /api/galaxy-weather/weather
	wg := api.Group("/weather")
	weather.Init(wg)

	// api/galaxy-weather/periods
	pg := api.Group("/periods")
	periods.Init(pg)

	// api/galaxy-weather/samples
	sg := api.Group("/samples")
	samples.Init(sg)

	return server
}

func errorHandler(err error, c echo.Context) {
	log.Errorf("API Error | uri=%s | Message: %s", c.Path(), err.Error())

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	galaxyWeatherError := &GalaxyWeatherError{
		Status:  "error",
		Message: err.Error(),
	}
	_ = c.JSON(code, galaxyWeatherError)
}

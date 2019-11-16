package health_check

import "github.com/labstack/echo"

func healthCheckHandler(c echo.Context) error {
	return c.JSON(200, "Galaxy Weather")
}

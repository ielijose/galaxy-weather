package health_check

import "github.com/labstack/echo"

func Init(g *echo.Group) {
	g.GET("", healthCheckHandler)
}

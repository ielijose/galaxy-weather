package periods

import "github.com/labstack/echo"

func Init(g *echo.Group) {
	g.GET("", periodsHandlers)
	g.GET("/stats", periodStatsHandlers)
}


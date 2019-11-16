package weather

import "github.com/labstack/echo"

func Init(g *echo.Group) {
	g.GET("/day/:day", weatherByDayHandler)
	g.GET("/year/:year", weatherByYearHandler)
}

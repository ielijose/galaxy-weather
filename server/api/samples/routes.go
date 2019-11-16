package samples

import "github.com/labstack/echo"

func Init(g *echo.Group) {
	g.GET("", sampleSeedHandler)
}

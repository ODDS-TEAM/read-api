package main

import (
	"net/http"

	"github.com/ODDS-TEAM/read-api/config"
	"github.com/ODDS-TEAM/read-api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {

	e := echo.New()
	s := config.Spec()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
	)

	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OddsRead OK!")
	})

	routes.Init(e)
	e.Logger.Fatal(e.Start(s.APIPort))
}

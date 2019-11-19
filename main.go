package main

import (
	// "net/http"

	"net/http"

	"github.com/ODDS-TEAM/read-api/config"
	"github.com/ODDS-TEAM/read-api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// fmt.Println("hello")

	// Use labstack/echo for rich routing.
	// See https://echo.labstack.com/
	e := echo.New()
	s := config.Spec()

	// Middleware
	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
		// middleware.JWTWithConfig(middleware.JWTConfig{
		// 	SigningKey: []byte("sMJuczqQPYzocl1s6SLj"),
		// 	Skipper: func(c echo.Context) bool {
		// 		// Skip authentication for and login requests
		// 		if c.Path() == "/login" || c.Path() == "/_ah/health" || c.Path() == "/book" {
		// 			return true
		// 		}
		// 		return false
		// 	},
		// }),
	)
	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OddsRead OK!")
	})

	routes.Init(e)
	e.Logger.Fatal(e.Start(s.APIPort))
}

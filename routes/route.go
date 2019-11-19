package routes

import (
	"github.com/ODDS-TEAM/read-api/api"
	"github.com/labstack/echo"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	// Database connection.
	db, err := api.NewMongoDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	a := &api.MongoDB{
		Conn: db.Conn,
		BCol: db.BCol,
		TCol: db.TCol,
	}

	e.POST("/postbook", a.PostBook)
	e.POST("/posttag", a.PostTag)
	e.GET("/book", a.GetBook)
}

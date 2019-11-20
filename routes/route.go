package routes

import (
	api "github.com/ODDS-TEAM/read-api/controller"
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

	e.POST("/tags", a.PostTag)
	e.GET("/tags", a.GetTag)
	e.POST("/books", a.PostBook)
	e.GET("/books", a.GetBook)

	e.GET("/checkisbn/:isbn", a.CheckISBN)
	e.POST("/mocktag", a.MockTag)
}

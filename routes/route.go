package routes

import (
	controller "github.com/ODDS-TEAM/read-api/controller"
	"github.com/labstack/echo"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	// Database connection.
	db, err := controller.NewMongoDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	a := &controller.MongoDB{
		Conn: db.Conn,
		BCol: db.BCol,
		TCol: db.TCol,
	}

	api := e.Group("/api")
	api.POST("/tags", a.PostTag)
	api.GET("/tags", a.GetTag)
	api.POST("/books", a.PostBook)
	api.GET("/books", a.GetBook)

	//api for function
	api.GET("/checkisbn/:isbn", a.CheckISBN)
	api.POST("/mocktag", a.MockTag)
}

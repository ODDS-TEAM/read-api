package routes

import (
	controller "github.com/ODDS-TEAM/read-api/controller"
	"github.com/labstack/echo"
)

// Init initialize api routes and set up a connection.
func Init(e *echo.Echo) {
	api := "/api"
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

	e.POST(api+"/tags", a.PostTag)
	e.GET(api+"/tags", a.GetTag)
	e.POST(api+"/books", a.PostBook)
	e.GET(api+"/books", a.GetBook)

	e.GET(api+"/checkisbn/:isbn", a.CheckISBN)
	e.POST(api+"/mocktag", a.MockTag)
}

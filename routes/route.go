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

	api := "/api"

	e.POST(api+"/tags", a.PostTag)
	e.GET(api+"/tags", a.GetTag)
	e.POST(api+"/books", a.PostBook)
	e.GET(api+"/books", a.GetBook)

	e.GET(api+"/checkisbn/:isbn", a.CheckISBN)
	e.POST(api+"/mocktag", a.MockTag)
}

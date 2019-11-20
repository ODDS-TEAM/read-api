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

	url := "/api"

	e.POST(url+"/tags", a.PostTag)
	e.GET(url+"/tags", a.GetTag)
	e.POST(url+"/books", a.PostBook)
	e.GET(url+"/books", a.GetBook)

	e.GET(url+"/checkisbn/:isbn", a.CheckISBN)
	e.POST(url+"/mocktag", a.MockTag)
}

package api

import (
	"fmt"
	"net/http"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//PostBook ..
func (db *MongoDB) PostBook(c echo.Context) error {

	book := &model.Book{}
	if err := c.Bind(book); err != nil {
		fmt.Println("In c.Bind Error ", err)
		return err
	}

	book.BookID = bson.NewObjectId()
	if err := db.BCol.Insert(book); err != nil {
		fmt.Println("In Insert Error", err)
		return err
	}

	return c.JSON(http.StatusOK, book)
}

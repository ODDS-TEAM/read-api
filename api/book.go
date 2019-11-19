package api

import (
	"fmt"
	"net/http"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//PostBook Function
func (db *MongoDB) PostBook(c echo.Context) error {

	//recive data from POST
	book := &model.Book{}
	if err := c.Bind(book); err != nil {
		fmt.Println("In c.Bind Error ", err)
		return err
	}

	fmt.Println("===============", book.Tag)
	//create ObjectID book
	book.BookID = bson.NewObjectId()

	if err := db.BCol.Insert(book); err != nil {
		fmt.Println("In Insert Error", err)
		return err
	}

	return c.JSON(http.StatusOK, book)
}

func (db *MongoDB) GetBook(c echo.Context) error {

	books := []bson.M{}

	pipeline := []bson.M{
		bson.M{"$lookup": bson.M{"from": "Tag",
			"localField":   "tag",
			"foreignField": "_id",
			"as":           "tag"}},
	}

	if err := db.BCol.Pipe(pipeline).All(&books); err != nil {
		fmt.Println("======Error in GetBook======")
		return err
	}

	return c.JSON(http.StatusOK, books)
}

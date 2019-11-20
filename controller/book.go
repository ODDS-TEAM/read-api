package controller

import (
	"fmt"
	"net/http"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PostBook Function
func (db *MongoDB) PostBook(c echo.Context) error {

	//recive data from POST
	book := &model.Book{}
	if err := c.Bind(book); err != nil {
		fmt.Println("In c.Bind Error ", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// make ISBN unique field
	index := mgo.Index{
		Key:    []string{"isbn"},
		Unique: true,
	}
	err := db.BCol.EnsureIndex(index)
	if err != nil {
		return err
	}

	fmt.Println("===============", book.Tags)
	//create ObjectID book
	book.BookID = bson.NewObjectId()

	if err := db.BCol.Insert(book); err != nil {
		fmt.Println("In Insert Error", err)
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, book)
}

//GetBook Function
func (db *MongoDB) GetBook(c echo.Context) error {

	books := []bson.M{}

	pipeline := []bson.M{
		bson.M{"$lookup": bson.M{"from": "Tag",
			"localField":   "tags",
			"foreignField": "_id",
			"as":           "tags"}},
	}

	if err := db.BCol.Pipe(pipeline).All(&books); err != nil {
		fmt.Println("======Error in GetBook======")
		return err
	}

	return c.JSON(http.StatusOK, books)
}

// CheckISBN function
func (db *MongoDB) CheckISBN(c echo.Context) error {

	isbn := c.Param("isbn")
	book := &model.Book{}

	if err := c.Bind(book); err != nil {
		fmt.Println("In CheckISBN Error ", err)
		return err
	}

	if err := db.BCol.Find(bson.M{"isbn": isbn}).One(book); err != nil {
		fmt.Println("======Error in CheckISBN======", err)
		return c.String(http.StatusOK, "false")
	}

	return c.String(http.StatusOK, "true")
}

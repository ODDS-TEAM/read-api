package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//PostBook Function
func (db *MongoDB) PostBook(c echo.Context) error {
	var temptrue []string
	var tempfalse []string

	books := &model.Book{}
	if err := c.Bind(books); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for i := range books.Tags {
		lowerTags := strings.ToLower(books.Tags[i])
		result := &model.Tag{}
		err := db.TCol.Find(bson.M{"tagName": lowerTags}).One(&result)
		if err != nil {
			result = db.CreateTag(lowerTags)
			tempfalse = append(tempfalse, result.TagName)
		} else {
			temptrue = append(temptrue, lowerTags)
		}
	}

	temptrue = append(temptrue, tempfalse...)
	books.BookID = bson.NewObjectId()
	books.Tags = temptrue
	bookUpload, isUpload, _ := UploadImgs(c)

	if isUpload == true {
		books.ImgURL = bookUpload.ImgURL
	}

	if err := db.BCol.Insert(books); err != nil {
		fmt.Println("In Insert Error", err)
		return c.JSON(http.StatusConflict, err)
	}

	return c.JSON(http.StatusCreated, books)
}

//GetBook Function
func (db *MongoDB) GetBook(c echo.Context) error {

	books := []model.Book{}

	// pipeline := []bson.M{
	// 	bson.M{"$lookup": bson.M{"from": "Tag",
	// 		"localField":   "tags",
	// 		"foreignField": "_id",
	// 		"as":           "tags_out"}},
	// }

	// if err := db.BCol.Pipe(pipeline).All(&books); err != nil {
	// 	fmt.Println("Error in GetBook", err)
	// 	return err
	// }

	if err := db.BCol.Find(bson.M{}).All(&books); err != nil {
		fmt.Println("Error in GetBook", err)
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
		fmt.Println("Error in CheckISBN", err)
		return c.String(http.StatusOK, "false")
	}

	return c.String(http.StatusOK, "true")
}

//CreateTag ...
func (db *MongoDB) CreateTag(tag string) *model.Tag {
	tags := &model.Tag{
		TagID:   bson.NewObjectId(),
		TagName: tag,
	}
	if err := db.TCol.Insert(tags); err != nil {
		fmt.Println("Error in CreateTag", err)
		return tags
	}
	return tags
}

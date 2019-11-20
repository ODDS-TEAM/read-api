package controller

import (
	"fmt"
	"net/http"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PostTag Function
func (db *MongoDB) PostTag(c echo.Context) error {

	tag := &model.Tag{}
	if err := c.Bind(tag); err != nil {
		fmt.Println("In c.Bind Error ", err)
		return err
	}

	// make tagName unique field
	index := mgo.Index{
		Key:    []string{"tagName"},
		Unique: true,
	}
	err := db.TCol.EnsureIndex(index)
	if err != nil {
		return err
	}

	tag.TagID = bson.NewObjectId()
	if err := db.TCol.Insert(tag); err != nil {

		fmt.Println("Error in PostTag", err)
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, tag)
}

//GetTag Function
func (db *MongoDB) GetTag(c echo.Context) error {

	tags := []model.Tag{}

	if err := db.TCol.Find(bson.M{}).All(&tags); err != nil {
		fmt.Println("Error in GetTag")
		return err
	}
	return c.JSON(http.StatusOK, tags)
}

func (db *MongoDB) MockTag(c echo.Context) error {

	tag := &model.Tag{}
	tag.TagID = bson.NewObjectId()
	tag.TagName = "Computer & Technology"
	if err := db.TCol.Insert(tag); err != nil {
		fmt.Println("Error in MockTag", err)
		return err
	}

	tag2 := &model.Tag{}
	tag2.TagID = bson.NewObjectId()
	tag2.TagName = "Cookbook & Food"
	if err := db.TCol.Insert(tag2); err != nil {
		fmt.Println("Error in MockTag2", err)
		return err
	}

	tag3 := &model.Tag{}
	tag3.TagID = bson.NewObjectId()
	tag3.TagName = "Humor & Entertainment"
	if err := db.TCol.Insert(tag3); err != nil {
		fmt.Println("Error in MockTag3", err)
		return err
	}
	return nil
}

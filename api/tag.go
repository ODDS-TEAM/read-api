package api

import (
	"fmt"
	"net/http"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//PostTag PostMethod
func (db *MongoDB) PostTag(c echo.Context) error {

	tag := &model.Tag{}
	if err := c.Bind(tag); err != nil {
		fmt.Println("In c.Bind Error ", err)
		return err
	}

	tag.TagID = bson.NewObjectId()
	if err := db.TCol.Insert(tag); err != nil {
		fmt.Println("In Insert Error", err)
		return err
	}

	return c.JSON(http.StatusOK, tag)
}

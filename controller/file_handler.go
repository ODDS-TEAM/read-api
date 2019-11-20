package controller

import (
	"io"
	"net/http"
	"os"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

//Upload ddd
func (db *MongoDB) UploadImgs(c echo.Context) error {
	books := &model.Book{
		BookID: bson.NewObjectId(),
	}

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()

	//Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	books.ImgURL = file.Filename

	return c.JSON(http.StatusOK, books)
}

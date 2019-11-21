package controller

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/ODDS-TEAM/read-api/model"
	"github.com/gofrs/uuid"
	"github.com/labstack/echo"
)

//UploadImgs function
func UploadImgs(c echo.Context) (*model.Book, bool, error) {
	books := &model.Book{}

	c.Bind(books)

	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r)
		}
	}()

	//read file
	file, err := c.FormFile("image")
	if err != nil {
		return books, false, nil
	}

	//source
	src, err := file.Open()
	if err != nil {
		return books, false, nil
	}
	defer src.Close()

	uniqueFileName := uuid.Must(uuid.NewV4()).String() + path.Ext(file.Filename)
	filePath := "./asset/images/" + uniqueFileName

	//destination
	dst, err := os.Create(filePath)
	if err != nil {
		return books, false, nil
	}

	defer dst.Close()

	//copy
	if _, err = io.Copy(dst, src); err != nil {
		return books, false, nil
	}

	books.ImgURL = uniqueFileName
	return books, true, nil
}

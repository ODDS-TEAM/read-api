package controller

import (
	"fmt"
	"io"
	"os"

	"github.com/ODDS-TEAM/read-api/model"
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
	filePath := "./assets/img/" + file.Filename

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

	books.ImgURL = file.Filename
	return books, true, nil
}

package tests_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ODDS-TEAM/read-api/model"
	"gopkg.in/mgo.v2/bson"

	"github.com/ODDS-TEAM/read-api/config"
	"github.com/ODDS-TEAM/read-api/controller"
	"github.com/ODDS-TEAM/read-api/routes"

	"github.com/labstack/echo"
)

func TestListBook_ListOfBooks(t *testing.T) {
	spec := config.Specification{
		DBHost:    "localhost:27017",
		DBName:    "oddreads_test",
		DBBookCol: "Book",
		DBTagCol:  "Tag",
	}
	book := &model.Book{
		BookID:     bson.NewObjectId(),
		Isbn:       "testIsbn",
		Title:      "testTitle",
		Authors:    []string{"testAuthors"},
		Publishers: []string{"testPublishers"},
		ImgURL:     "testImgURL",
		Tags:       []string{"testTags"},
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
	rec := httptest.NewRecorder()
	db, _ := controller.NewMongoDB(&spec)
	routes.Init(e, &spec)
	db.BCol.RemoveAll(nil)
	db.BCol.Insert(book)
	e.ServeHTTP(rec, req)
	if rec.Code != 200 {
		t.Errorf("want rec.Code = 200 but got %d", rec.Code)
	}
	books := []model.Book{}
	dec := json.NewDecoder(rec.Body)
	if err := dec.Decode(&books); err != nil {
		t.Fatal(err)
	}
	if len(books) != 1 {
		t.Errorf("want len(tags) = 1 but got %d", len(books))
	}
	db.TCol.RemoveAll(nil)
}

func TestPostBook_InsertBook(t *testing.T) {
	book := `{"isbn":"testIsbn","title":"testTitle","authors":["testAuthors1","testAuthors2"],"publishers":["testPublishers"],"imgURL":"testImgURL","tags":["testTags1","testTags2"]}`
	req := httptest.NewRequest(http.MethodPost, "/api/books", strings.NewReader(book))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	spec := config.Specification{
		DBHost:    "localhost:27017",
		DBName:    "oddreads_test",
		DBBookCol: "Book",
		DBTagCol:  "Tag",
	}

	db, _ := controller.NewMongoDB(&spec)

	routes.Init(e, &spec)
	db.BCol.RemoveAll(nil)
	books := []model.Book{}
	db.BCol.Find(bson.M{}).All(&books)
	e.ServeHTTP(rec, req)

	if rec.Code != 201 {
		t.Errorf("want rec.Code = 201 but got %d", rec.Code)
	}
	db.BCol.RemoveAll(nil)
}

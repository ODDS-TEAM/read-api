package tests_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ODDS-TEAM/read-api/config"
	controller "github.com/ODDS-TEAM/read-api/controller"
	"github.com/ODDS-TEAM/read-api/model"
	"github.com/ODDS-TEAM/read-api/routes"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func TestListTag_ListOfTags(t *testing.T) {

	tag1 := &model.Tag{
		TagID:   bson.NewObjectId(),
		TagName: "First",
	}
	tag2 := &model.Tag{
		TagID:   bson.NewObjectId(),
		TagName: "Two",
	}

	req := httptest.NewRequest(http.MethodGet, "/api/tags", nil)
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

	db.TCol.RemoveAll(nil)

	db.TCol.Insert(tag1)
	db.TCol.Insert(tag2)

	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("want rec.Code = 200 but got %d", rec.Code)
	}

	tags := []model.Tag{}
	dec := json.NewDecoder(rec.Body)
	if err := dec.Decode(&tags); err != nil {
		t.Fatal(err)
	}
	if len(tags) != 2 {
		t.Errorf("want len(tags) = 2 but got %d", len(tags))
	}

	db.TCol.RemoveAll(nil)
}

func TestPostTag_InsertTag(t *testing.T) {
	tag := `{"tagName":"testTagName"}`
	req := httptest.NewRequest(http.MethodPost, "/api/tags", strings.NewReader(tag))
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
	db.TCol.RemoveAll(nil)
	tags := []model.Tag{}
	db.TCol.Find(bson.M{}).All(&tags)
	e.ServeHTTP(rec, req)

	if rec.Code != 200 {
		t.Errorf("want rec.Code = 200 but got %d", rec.Code)
	}
	db.TCol.RemoveAll(nil)
}

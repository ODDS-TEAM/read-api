package model

import "gopkg.in/mgo.v2/bson"

// Book Model
type Book struct {
	BookID    bson.ObjectId   `json:"_id" bson:"_id"`
	Isbn      string          `json:"isbn" bson:"isbn"`
	Title     string          `json:"title" bson:"title"`
	Author    []string        `json:"author" bson:"author"`
	Publisher string          `json:"publisher" bson:"publisher"`
	ImgURL    string          `json:"imgURL" bson:"imgURL"`
	Tags      []bson.ObjectId `json:"tags" bson:"tags"`
}

package model

import "gopkg.in/mgo.v2/bson"

// Book Model
type Book struct {
	BookID     bson.ObjectId `json:"_id" bson:"_id"`
	Isbn       string        `json:"isbn" bson:"isbn"`
	Title      string        `json:"title" bson:"title"`
	Authors    []string      `json:"authors" bson:"authors"`
	Publishers []string      `json:"publishers" bson:"publishers"`
	ImgURL     string        `json:"imgURL" bson:"imgURL"`
	Tags       []string      `json:"tags" bson:"tags"`
}

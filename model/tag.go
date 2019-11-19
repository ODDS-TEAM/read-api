package model

import "gopkg.in/mgo.v2/bson"

//Tag Model
type Tag struct {
	TagID    	bson.ObjectId `json:"_id" bson:"_id"`
	TagName      	string    `json:"tagName" bson:"tagName"`

}

package model

import "gopkg.in/mgo.v2/bson"

//Publisher Model
type Publisher struct {
	PublisherID    	bson.ObjectId `json:"_id" bson:"_id"`
	Publisher      	string    `json:"publisher" bson:"publisher"`

}

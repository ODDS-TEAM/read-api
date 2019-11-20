package controller

import (
	"fmt"

	"github.com/ODDS-TEAM/read-api/config"
	"gopkg.in/mgo.v2"
)

// MongoDB holds metadata about session database and collections name.
type (
	MongoDB struct {
		Conn *mgo.Session
		BCol *mgo.Collection
		TCol *mgo.Collection
	}
)

// NewMongoDB creates a new macOddsTeamDB backed by a given Mongo server.
func NewMongoDB() (*MongoDB, error) {
	s := config.Spec()
	conn, err := mgo.Dial(s.DBHost)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	return &MongoDB{
		Conn: conn,
		BCol: conn.DB(s.DBName).C(s.DBBookCol),
		TCol: conn.DB(s.DBName).C(s.DBTagCol),
	}, nil
}

// Close closes the database.
func (db *MongoDB) Close() {
	db.Conn.Close()
}

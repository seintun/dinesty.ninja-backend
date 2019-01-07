package models

import "gopkg.in/mgo.v2/bson"

// Item struct represents a basic information of restaurant business
type Item struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	BizID       string        `json:"bizID"`
	Name        string        `json:"name"`
	Price       int           `json:"price"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Active      bool          `json:"active"`
}

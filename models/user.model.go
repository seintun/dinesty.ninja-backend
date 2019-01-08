package models

import "gopkg.in/mgo.v2/bson"

// User struct represents a basic information of user
type User struct {
	ID bson.ObjectId `bson:"_id" json:"id"`

	Name     string        `json:"name"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
	Phone    string        `json:"phone"`
	Orders   []interface{} `json:"orders"`
}

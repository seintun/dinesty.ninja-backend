package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Order struct represents a basic information of Order
type Order struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	BizID     string        `json:"bizID"`
	UserID    string        `json:"userID"`
	BizName   string        `json:"bizName"`
	UserName  string        `json:"userName"`
	Guests    int           `json:"guests"`
	Date      time.Time     `json:"date"`
	MenuItems []Item        `json:"menuItems"`
	Paid      bool          `json:"paid"`
	Cancelled bool          `json:"cancelled"`
	Total     int           `json:"total"`
	Tax       int           `json:"tax"`
	Tip       int           `json:"tip"`
}

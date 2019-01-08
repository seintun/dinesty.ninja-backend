package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Order struct represents a basic information of Order
type Order struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	BizID     int64         `json:"bizID"`
	UserID    int64         `json:"userID"`
	BizName   string        `json:"bizName"`
	UserName  string        `json:"userName"`
	Guests    int           `json:"guests"`
	Date      time.Time     `json:"date"`
	MenuItems []interface{} `json:"menuItems"`
}

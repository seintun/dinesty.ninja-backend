package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	OCOLLECTION = "order"
)

// Queries

// CreateOrder insert a user into database
func (b *BizDAO) CreateOrder(o Order) error {
	err := db.C(OCOLLECTION).Insert(&o)
	return err
}

// FindOrderByID return specified user
func (b *BizDAO) FindOrderByID(id string) (Order, error) {
	query := bson.ObjectIdHex(id)
	var o Order
	err := db.C(OCOLLECTION).FindId(query).One(&o)
	return o, err
}

// DeleteOrderByID an existing user
func (b *BizDAO) DeleteOrderByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(OCOLLECTION).Remove(query)
	return err
}

// AddItemtoMenu an existing menuItem
func (b *BizDAO) AddItemtoMenu(id string, item Item) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	insert := bson.M{"$push": bson.M{"menuitems": &item}}
	err := db.C(OCOLLECTION).Update(query, insert)
	return err
}

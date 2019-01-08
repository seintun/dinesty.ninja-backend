package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	OCOLLECTION = "order"
)

// Queries

// Insert a user into database
func (b *BizDAO) CreateOrder(o Order) error {
	err := db.C(OCOLLECTION).Insert(&o)
	return err
}

// FindOrderByID return specified user
func (b *BizDAO) FindOrderByID(id string) (Order, error) {
	var o Order
	err := db.C(OCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&o)
	return o, err
}

// DeleteOrderByID an existing user
func (b *BizDAO) DeleteOrderByID(id string) error {
	err := db.C(OCOLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

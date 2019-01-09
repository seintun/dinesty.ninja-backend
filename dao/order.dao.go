package dao

import (
	"fmt"

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

// AddItemtoCart an existing menuItem
func (b *BizDAO) AddItemtoCart(id string, item Item) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	insert := bson.M{"$push": bson.M{"cart": &item}}
	err := db.C(OCOLLECTION).Update(query, insert)
	return err
}

// DeleteItemfromCart an existing menuItem
func (b *BizDAO) DeleteItemfromCart(id string, cid string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	item := bson.M{"cart": bson.M{"_id": bson.ObjectIdHex(cid)}}
	itemtoRemove := bson.M{"$pull": item}
	fmt.Println(id)
	fmt.Println(cid)
	err := db.C(OCOLLECTION).Update(query, itemtoRemove)
	return err
}

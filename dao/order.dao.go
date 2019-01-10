package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	OCOLLECTION = "order"
)

// Queries

// CreateOrder insert new order into order database, and push orderID into user and biz orders[]
func (b *BizDAO) CreateOrder(o Order) error {
	uQuery := bson.M{"_id": bson.ObjectIdHex(o.UserID)}
	bQuery := bson.M{"_id": bson.ObjectIdHex(o.BizID)}
	uInsert := bson.M{"$push": bson.M{"orders": o.ID.Hex()}}
	bInsert := bson.M{"$push": bson.M{"orders": o.ID.Hex()}}
	err := db.C(OCOLLECTION).Insert(&o)
	db.C(UCOLLECTION).Update(uQuery, uInsert)
	db.C(BCOLLECTION).Update(bQuery, bInsert)
	return err
}

// FindOrderByID return specified order
func (b *BizDAO) FindOrderByID(id string) (Order, error) {
	query := bson.ObjectIdHex(id)
	var o Order
	err := db.C(OCOLLECTION).FindId(query).One(&o)
	return o, err
}

// FetchOrdersByuserID return all orders of specified user
func (b *BizDAO) FetchOrdersByuserID(id string) ([]Order, error) {
	query := bson.M{"userid": id}
	var o []Order
	err := db.C(OCOLLECTION).Find(query).All(&o)
	return o, err
}

// FetchOrdersBybizID return all orders of specified user
func (b *BizDAO) FetchOrdersBybizID(id string) ([]Order, error) {
	query := bson.M{"bizid": id}
	var o []Order
	err := db.C(OCOLLECTION).Find(query).All(&o)
	return o, err
}

// UpdateOrderByID an existing order
func (b *BizDAO) UpdateOrderByID(id string, o Order) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(OCOLLECTION).Update(query, &o)
	return err
}

// CancelOrderByID an existing biz
func (b *BizDAO) CancelOrderByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	cancel := bson.M{"$set": bson.M{"cancelled": true}}
	err := db.C(OCOLLECTION).Update(query, cancel)
	return err
}

// DeleteOrderByID an existing order
func (b *BizDAO) DeleteOrderByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(OCOLLECTION).Remove(query)
	return err
}

// AddItemtoCart an existing cart
func (b *BizDAO) AddItemtoCart(id string, item Item) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	insert := bson.M{"$push": bson.M{"cart": &item}}
	err := db.C(OCOLLECTION).Update(query, insert)
	return err
}

// DeleteItemfromCart an existing cart
func (b *BizDAO) DeleteItemfromCart(id string, cid string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	item := bson.M{"cart": bson.M{"_id": bson.ObjectIdHex(cid)}}
	itemtoRemove := bson.M{"$pull": item}
	err := db.C(OCOLLECTION).Update(query, itemtoRemove)
	return err
}

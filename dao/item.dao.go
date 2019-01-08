package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	MCOLLECTION = "menu"
)

// Queries

// CreateItem insert menuItem into database
func (b *BizDAO) CreateItem(item Item) error {
	err := db.C(MCOLLECTION).Insert(&item)
	return err
}

// FetchItems return array of menuItems
func (b *BizDAO) FetchItems(id string) ([]Item, error) {
	var items []Item
	err := db.C(MCOLLECTION).Find(bson.M{"bizid": id}).All(&items)
	return items, err
}

// FindItemByID return specified menuItem
func (b *BizDAO) FindItemByID(mid string) (Item, error) {
	var item Item
	err := db.C(MCOLLECTION).FindId(bson.ObjectIdHex(mid)).One(&item)
	return item, err
}

// UpdateItemByID an existing menuItem
func (b *BizDAO) UpdateItemByID(mid string, item Item) error {
	err := db.C(MCOLLECTION).Update(bson.M{"_id": bson.ObjectIdHex(mid)}, &item)
	return err
}

// DeleteItemByID an existing menuItem
func (b *BizDAO) DeleteItemByID(mid string) error {
	err := db.C(MCOLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(mid)})
	return err
}

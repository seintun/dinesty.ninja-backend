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
	query := bson.M{"bizid": id}
	var items []Item
	err := db.C(MCOLLECTION).Find(query).All(&items)
	return items, err
}

// FindItemByID return specified menuItem
func (b *BizDAO) FindItemByID(mid string) (Item, error) {
	query := bson.ObjectIdHex(mid)
	var item Item
	err := db.C(MCOLLECTION).FindId(query).One(&item)
	return item, err
}

// UpdateItemByID an existing menuItem
func (b *BizDAO) UpdateItemByID(mid string, item Item) error {
	query := bson.M{"_id": bson.ObjectIdHex(mid)}
	err := db.C(MCOLLECTION).Update(query, &item)
	return err
}

// DeleteItemByID an existing menuItem
func (b *BizDAO) DeleteItemByID(mid string) error {
	query := bson.M{"_id": bson.ObjectIdHex(mid)}
	err := db.C(MCOLLECTION).Remove(query)
	return err
}

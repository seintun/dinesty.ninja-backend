package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	MCOLLECTION = "menu"
)

// Queries

// Insert a biz into database
func (b *BizDAO) CreateItem(item Item) error {
	err := db.C(MCOLLECTION).Insert(&item)
	return err
}

// FetchItems return array of menuItems from specified Biz
func (b *BizDAO) FetchItems(id string) ([]Item, error) {
	var item []Item
	err := db.C(MCOLLECTION).Find(bson.M{}).All(&item)
	return item, err
}

package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
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

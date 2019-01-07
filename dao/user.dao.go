package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
)

const (
	UCOLLECTION = "user"
)

// Queries

// Insert a biz into database
func (b *BizDAO) CreateUser(u User) error {
	err := db.C(UCOLLECTION).Insert(&u)
	return err
}

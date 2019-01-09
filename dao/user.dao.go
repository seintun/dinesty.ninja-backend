package dao

import (
	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	UCOLLECTION = "user"
)

// Queries

// Insert a user into database
func (b *BizDAO) CreateUser(u User) error {
	err := db.C(UCOLLECTION).Insert(&u)
	return err
}

// FindUserByID return specified user
func (b *BizDAO) FindUserByID(id string) (User, error) {
	query := bson.ObjectIdHex(id)
	var u User
	err := db.C(UCOLLECTION).FindId(query).One(&u)
	return u, err
}

// UpdateUserByID an existing user
func (b *BizDAO) UpdateUserByID(id string, u User) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(UCOLLECTION).Update(query, &u)
	return err
}

// DeleteUserByID an existing user
func (b *BizDAO) DeleteUserByID(id string) error {
	query := bson.M{"_id": bson.ObjectIdHex(id)}
	err := db.C(UCOLLECTION).Remove(query)
	return err
}

package dao

import (
	"log"

	. "github.com/seintun/dinesty.ninja-backend/models"

	mgo "gopkg.in/mgo.v2"
)

type BizDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "business"
)

// Connect Establish a connection to database
func (b *BizDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(b.Database)
}

// Queries

// Insert a biz into database
func (b *BizDAO) Insert(biz Biz) error {
	err := db.C(COLLECTION).Insert(&biz)
	return err
}

package dao

import (
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
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

// FindAll list of bizs
func (b *BizDAO) FetchBiz() ([]Biz, error) {
	var bizs []Biz
	err := db.C(COLLECTION).Find(bson.M{}).All(&bizs)
	return bizs, err
}

// // FindById a biz by its id
// func (b *BizDAO) FindById(id string) (Biz, error) {
// 	var biz Biz
// 	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&biz)
// 	return biz, err
// }

// Insert a biz into database
func (b *BizDAO) Insert(biz Biz) error {
	err := db.C(COLLECTION).Insert(&biz)
	return err
}

// // Delete an existing biz
// func (b *BizDAO) Delete(biz Biz) error {
// 	err := db.C(COLLECTION).Remove(&biz)
// 	return err
// }

// // Update an existing biz
// func (b *BizDAO) Update(biz Biz) error {
// 	err := db.C(COLLECTION).UpdateId(biz.ID, &biz)
// 	return err
// }

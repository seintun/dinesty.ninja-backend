package dao

import (
	"fmt"
	"log"

	. "github.com/seintun/dinesty.ninja-backend/models"
	"gopkg.in/mgo.v2/bson"

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

// FindBizByID return specified Biz
func (b *BizDAO) FindBizByID(id string) (Biz, error) {
	var biz Biz
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&biz)
	return biz, err
}

// Insert a biz into database
func (b *BizDAO) Insert(biz Biz) error {
	err := db.C(COLLECTION).Insert(&biz)
	return err
}

// DeleteBizByID an existing biz
func (b *BizDAO) DeleteBizByID(id string) error {
	fmt.Println(bson.ObjectId(id))
	err := db.C(COLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// // Update an existing biz
// func (b *BizDAO) Update(biz Biz) error {
// 	err := db.C(COLLECTION).UpdateId(biz.ID, &biz)
// 	return err
// }
